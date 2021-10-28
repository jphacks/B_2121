package api

import (
	"database/sql"
	"net/http"

	"github.com/jphacks/B_2121_server/models"
	"github.com/jphacks/B_2121_server/openapi"
	"github.com/jphacks/B_2121_server/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

func (h handler) NewCommunity(ctx echo.Context) error {
	info := session.GetAuthInfo(ctx)
	if !info.Authenticated {
		return echo.ErrUnauthorized
	}

	var req openapi.CreateCommunityRequest
	err := ctx.Bind(&req)
	if err != nil {
		ctx.Logger().Errorf("failed to bind request: %v", err)
		return echo.ErrBadRequest
	}

	loc := models.FromOpenApiLocation(req.Location)
	community, err := h.communityUseCase.NewCommunity(ctx.Request().Context(), info.UserId, req.Name, req.Description, *loc)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, community.ToOpenApiCommunity())
}

func (h handler) SearchCommunities(ctx echo.Context, params openapi.SearchCommunitiesParams) error {
	panic("implement me")
}

func (h handler) GetCommunityById(ctx echo.Context, id int) error {
	community, err := h.communityUseCase.GetCommunity(ctx.Request().Context(), int64(id))
	if err != nil {
		if xerrors.Is(err, sql.ErrNoRows) {
			return echo.ErrNotFound
		}
		return err
	}

	return ctx.JSON(http.StatusOK, community.ToOpenApiCommunityDetail())
}

func (h handler) ListCommunityRestaurants(ctx echo.Context, id int, params openapi.ListCommunityRestaurantsParams) error {
	panic("implement me")
}