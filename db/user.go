package db

import (
	"context"
	"net/url"
	"path"

	"github.com/jmoiron/sqlx"
	"github.com/jphacks/B_2121_server/models"
	"github.com/jphacks/B_2121_server/models_gen"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/xerrors"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) models.UserRepository {
	return &userRepository{db: db}
}

func (u userRepository) GetUserById(ctx context.Context, id int64, profileImageBase url.URL) (*models.User, error) {
	user, err := models_gen.FindUser(ctx, u.db, id)
	if err != nil {
		return nil, xerrors.Errorf("failed to find a user by ID: %w", err)
	}

	return fromGenUser(user, profileImageBase), err
}

func (u userRepository) GetUserDetailById(ctx context.Context, id int64, profileImageBase url.URL) (*models.UserDetail, error) {
	user, err := models_gen.FindUser(ctx, u.db, id)
	if err != nil {
		return nil, xerrors.Errorf("failed to find a user by ID: %w", err)
	}

	communityCount, err := user.Affiliations().Count(ctx, u.db)
	if err != nil {
		return nil, xerrors.Errorf("failed to count joined communities: %w", err)
	}
	bookmarkCount, err := user.Bookmarks().Count(ctx, u.db)
	if err != nil {
		return nil, xerrors.Errorf("failed to count bookmarked communities: %w", err)
	}

	modelUser := fromGenUser(user, profileImageBase)
	return &models.UserDetail{
		User:           *modelUser,
		BookmarkCount:  int(bookmarkCount),
		CommunityCount: int(communityCount),
	}, err
}

func (u userRepository) NewUser(ctx context.Context, userName string) (*models.User, error) {
	result, err := u.db.ExecContext(ctx, "INSERT INTO users(`name`) VALUES (?)", userName)
	if err != nil {
		return nil, xerrors.Errorf("failed to insert a new user: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, xerrors.Errorf("failed to get last insert id: %w", err)
	}
	return &models.User{
		Id:   id,
		Name: userName,
	}, nil
}

func (u userRepository) UpdateProfileImage(ctx context.Context, userId int64, fileName string) error {
	user, err := models_gen.FindUser(ctx, u.db, userId)
	if err != nil {
		return xerrors.Errorf("failed to find a user by ID: %w", err)
	}
	user.ProfileImageFile.SetValid(fileName)
	_, err = user.Update(ctx, u.db, boil.Infer())
	if err != nil {
		return xerrors.Errorf("failed to update database: %w", err)
	}
	return nil
}

func (u userRepository) ListUserCommunity(ctx context.Context, userId int64) ([]*models.Community, error) {
	community, err := models_gen.Communities(
		qm.InnerJoin("affiliation ON affiliation.community_id = communities.id"),
		qm.Where("user_id=?", userId),
	).All(ctx, u.db)
	if err != nil {
		return nil, xerrors.Errorf("failed to get communities: %w", err)
	}
	ret := make([]*models.Community, 0)
	for _, c := range community {
		ret = append(ret, &models.Community{Community: *c})
	}
	return ret, nil
}

func fromGenUser(u *models_gen.User, imageUrlBase url.URL) *models.User {
	if u.ProfileImageFile.Valid {
		imageUrlBase.Path = path.Join(imageUrlBase.Path, u.ProfileImageFile.String)
		return &models.User{
			Id:              u.ID,
			Name:            u.Name,
			ProfileImageUrl: imageUrlBase.String(),
		}
	}

	return &models.User{
		Id:              u.ID,
		Name:            u.Name,
		ProfileImageUrl: "",
	}
}