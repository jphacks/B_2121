// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.3 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

const (
	TokenScopes = "token.Scopes"
)

// Defines values for AuthVendor.
const (
	AuthVendorAnonymous AuthVendor = "Anonymous"

	AuthVendorApple AuthVendor = "Apple"

	AuthVendorGoogle AuthVendor = "Google"
)

// Add a restaurant to a community
type AddRestaurantRequest struct {
	RestaurantId Long `json:"restaurant_id"`
}

// AuthInfo defines model for authInfo.
type AuthInfo struct {
	Token  string     `json:"token"`
	Vendor AuthVendor `json:"vendor"`
}

// AuthVendor defines model for authVendor.
type AuthVendor string

// Private comments for a restaurant
type Comment struct {
	Body         *string `json:"body,omitempty"`
	CommunityId  *Long   `json:"community_id,omitempty"`
	RestaurantId *Long   `json:"restaurant_id,omitempty"`

	// Updated date and time
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Goyotashi community
type Community struct {
	Description   string   `json:"description"`
	Id            Long     `json:"id"`
	ImageUrls     []string `json:"imageUrls"`
	Location      Location `json:"location"`
	Name          string   `json:"name"`
	NumRestaurant int      `json:"num_restaurant"`
	NumUser       int      `json:"num_user"`
}

// CreateCommunityRequest defines model for createCommunityRequest.
type CreateCommunityRequest struct {
	Description string   `json:"description"`
	Location    Location `json:"location"`
	Name        string   `json:"name"`
}

// CreateUserRequest defines model for createUserRequest.
type CreateUserRequest struct {
	Name   string     `json:"name"`
	Vendor AuthVendor `json:"vendor"`
}

// CreateUserResponse defines model for createUserResponse.
type CreateUserResponse struct {
	AuthInfo AuthInfo `json:"auth_info"`

	// Reperesents user
	User User `json:"user"`
}

// GetCommunityIdTokenResponse defines model for getCommunityIdTokenResponse.
type GetCommunityIdTokenResponse struct {
	// Token dulation (seconds)
	ExpiresIn   int    `json:"expires_in"`
	InviteToken string `json:"invite_token"`
}

// JoinCommunityRequest defines model for joinCommunityRequest.
type JoinCommunityRequest struct {
	InviteToken string `json:"invite_token"`
}

// ListCommunityRestaurantsResponse defines model for listCommunityRestaurantsResponse.
type ListCommunityRestaurantsResponse struct {
	// Embedded struct due to allOf(#/components/schemas/pageInfo)
	PageInfo `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Restaurants *[]Restaurant `json:"restaurants,omitempty"`
}

// ListCommunityUsersResponse defines model for listCommunityUsersResponse.
type ListCommunityUsersResponse struct {
	// Embedded struct due to allOf(#/components/schemas/user)
	User `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Users *[]User `json:"users,omitempty"`
}

// ListUserBookmarkResponse defines model for listUserBookmarkResponse.
type ListUserBookmarkResponse struct {
	// Embedded struct due to allOf(#/components/schemas/pageInfo)
	PageInfo `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Communities *[]Community `json:"communities,omitempty"`
}

// ListUserCommunityResponse defines model for listUserCommunityResponse.
type ListUserCommunityResponse struct {
	// Embedded fields due to inline allOf schema
	Communities *[]Community `json:"communities,omitempty"`
	// Embedded struct due to allOf(#/components/schemas/pageInfo)
	PageInfo `yaml:",inline"`
}

// Location defines model for location.
type Location struct {
	// latitude
	Lat float64 `json:"lat"`

	// longitude
	Lng float64 `json:"lng"`
}

// Long defines model for long.
type Long int64

// PageInfo defines model for pageInfo.
type PageInfo struct {
	BeginCursor *int  `json:"begin_cursor,omitempty"`
	EndCursor   *int  `json:"end_cursor,omitempty"`
	HasNext     *bool `json:"has_next,omitempty"`
	HasPrevious *bool `json:"has_previous,omitempty"`
}

// Restaurant
type Restaurant struct {
	Id       Long     `json:"id"`
	ImageUrl *string  `json:"image_url,omitempty"`
	Location Location `json:"location"`
	Name     string   `json:"name"`
}

// SearchCommunityResponse defines model for searchCommunityResponse.
type SearchCommunityResponse struct {
	// Embedded struct due to allOf(#/components/schemas/pageInfo)
	PageInfo `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Communities *[]Community `json:"communities,omitempty"`
}

// SearchRestaurantResponse defines model for searchRestaurantResponse.
type SearchRestaurantResponse struct {
	// Embedded struct due to allOf(#/components/schemas/pageInfo)
	PageInfo `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Restaurants *[]Restaurant `json:"restaurants,omitempty"`
}

// Update private comments for a restaurant
type UpdateCommentRequest struct {
	Body *string `json:"body,omitempty"`
}

// UploadImageProfileResponse defines model for uploadImageProfileResponse.
type UploadImageProfileResponse struct {
	ImageUrl string `json:"imageUrl"`
}

// Reperesents user
type User struct {
	Id              Long    `json:"id"`
	Name            string  `json:"name"`
	ProfileImageUrl *string `json:"profile_image_url,omitempty"`
}

// UserDetail defines model for userDetail.
type UserDetail struct {
	// Embedded struct due to allOf(#/components/schemas/user)
	User `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	BookmarkCount  int `json:"bookmark_count"`
	CommunityCount int `json:"community_count"`
}

// PageQuery defines model for pageQuery.
type PageQuery Long

// NewCommunityJSONBody defines parameters for NewCommunity.
type NewCommunityJSONBody CreateCommunityRequest

// SearchCommunitiesParams defines parameters for SearchCommunities.
type SearchCommunitiesParams struct {
	After     *PageQuery `json:"after,omitempty"`
	Keyword   string     `json:"keyword"`
	CenterLat *float64   `json:"center_lat,omitempty"`
	CenterLng *float64   `json:"center_lng,omitempty"`
}

// ListCommunityRestaurantsParams defines parameters for ListCommunityRestaurants.
type ListCommunityRestaurantsParams struct {
	After *PageQuery `json:"after,omitempty"`
}

// AddRestaurantToCommunityJSONBody defines parameters for AddRestaurantToCommunity.
type AddRestaurantToCommunityJSONBody AddRestaurantRequest

// UpdateRestaurantCommentJSONBody defines parameters for UpdateRestaurantComment.
type UpdateRestaurantCommentJSONBody UpdateCommentRequest

// ListUsersOfCommunityParams defines parameters for ListUsersOfCommunity.
type ListUsersOfCommunityParams struct {
	After *PageQuery `json:"after,omitempty"`
}

// SearchRestaurantsParams defines parameters for SearchRestaurants.
type SearchRestaurantsParams struct {
	After     *PageQuery `json:"after,omitempty"`
	Keyword   string     `json:"keyword"`
	CenterLat *float64   `json:"center_lat,omitempty"`
	CenterLng *float64   `json:"center_lng,omitempty"`
}

// NewUserJSONBody defines parameters for NewUser.
type NewUserJSONBody CreateUserRequest

// PostUserMeCommunitiesJSONBody defines parameters for PostUserMeCommunities.
type PostUserMeCommunitiesJSONBody JoinCommunityRequest

// PostUserIdBookmarkJSONBody defines parameters for PostUserIdBookmark.
type PostUserIdBookmarkJSONBody struct {
	CommunityId Long `json:"community_id"`
}

// ListUserCommunitiesParams defines parameters for ListUserCommunities.
type ListUserCommunitiesParams struct {
	After *PageQuery `json:"after,omitempty"`
}

// NewCommunityJSONRequestBody defines body for NewCommunity for application/json ContentType.
type NewCommunityJSONRequestBody NewCommunityJSONBody

// AddRestaurantToCommunityJSONRequestBody defines body for AddRestaurantToCommunity for application/json ContentType.
type AddRestaurantToCommunityJSONRequestBody AddRestaurantToCommunityJSONBody

// UpdateRestaurantCommentJSONRequestBody defines body for UpdateRestaurantComment for application/json ContentType.
type UpdateRestaurantCommentJSONRequestBody UpdateRestaurantCommentJSONBody

// NewUserJSONRequestBody defines body for NewUser for application/json ContentType.
type NewUserJSONRequestBody NewUserJSONBody

// PostUserMeCommunitiesJSONRequestBody defines body for PostUserMeCommunities for application/json ContentType.
type PostUserMeCommunitiesJSONRequestBody PostUserMeCommunitiesJSONBody

// PostUserIdBookmarkJSONRequestBody defines body for PostUserIdBookmark for application/json ContentType.
type PostUserIdBookmarkJSONRequestBody PostUserIdBookmarkJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new community
	// (POST /community)
	NewCommunity(ctx echo.Context) error
	// Search communities using keyword and location
	// (GET /community/search)
	SearchCommunities(ctx echo.Context, params SearchCommunitiesParams) error
	// Get a community by id
	// (GET /community/{id})
	GetCommunityById(ctx echo.Context, id int) error
	// List restaurants in a community
	// (GET /community/{id}/restaurants)
	ListCommunityRestaurants(ctx echo.Context, id int, params ListCommunityRestaurantsParams) error
	// Add a restaurant to a community
	// (POST /community/{id}/restaurants)
	AddRestaurantToCommunity(ctx echo.Context, id int) error
	// Remove a restrant from the specified community
	// (DELETE /community/{id}/restaurants/{restaurant_id})
	RemoveRestaurantFromCommunity(ctx echo.Context, id int64, restaurantId int64) error
	// Get private comments for a restaurant
	// (GET /community/{id}/restaurants/{restaurant_id}/comments)
	GetRestaurantComment(ctx echo.Context, id int, restaurantId int) error
	// Update comment of the restaurant
	// (PUT /community/{id}/restaurants/{restaurant_id}/comments)
	UpdateRestaurantComment(ctx echo.Context, id int, restaurantId int) error
	// Get an invite token
	// (GET /community/{id}/token)
	GetCommunityIdToken(ctx echo.Context, id int) error
	// List users in a community
	// (GET /community/{id}/users)
	ListUsersOfCommunity(ctx echo.Context, id int, params ListUsersOfCommunityParams) error
	// Search restaurants using keyword and location
	// (GET /restaurant/search)
	SearchRestaurants(ctx echo.Context, params SearchRestaurantsParams) error
	// Get information about the speicifed restaurant.
	// (GET /restaurant/{id})
	GetRestaurantId(ctx echo.Context, id int64) error
	// Create a new user
	// (POST /user)
	NewUser(ctx echo.Context) error
	// Get my profile in detail
	// (GET /user/me)
	GetMyProfile(ctx echo.Context) error
	// Join a community
	// (POST /user/me/communities)
	PostUserMeCommunities(ctx echo.Context) error

	// (POST /user/profile)
	UploadProfileImage(ctx echo.Context) error
	// Get bookmarking list of the specified user
	// (GET /user/{id}/bookmark)
	GetUserIdBookmark(ctx echo.Context, id Long) error
	// Create a new bookmark
	// (POST /user/{id}/bookmark)
	PostUserIdBookmark(ctx echo.Context, id Long) error
	// Delete bookmark from the specified user
	// (DELETE /user/{id}/bookmark/{community_id})
	DeleteUserIdBookmarkCommunityId(ctx echo.Context, id Long, communityId Long) error
	// Get communities where the specified user joins
	// (GET /user/{id}/communities)
	ListUserCommunities(ctx echo.Context, id Long, params ListUserCommunitiesParams) error
	// Leave a community
	// (DELETE /user/{id}/communities/{community_id})
	DeleteUserIdCommunitiesCommunityId(ctx echo.Context, id Long, communityId Long) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// NewCommunity converts echo context to params.
func (w *ServerInterfaceWrapper) NewCommunity(ctx echo.Context) error {
	var err error

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.NewCommunity(ctx)
	return err
}

// SearchCommunities converts echo context to params.
func (w *ServerInterfaceWrapper) SearchCommunities(ctx echo.Context) error {
	var err error

	ctx.Set(TokenScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchCommunitiesParams
	// ------------- Optional query parameter "after" -------------

	err = runtime.BindQueryParameter("form", true, false, "after", ctx.QueryParams(), &params.After)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter after: %s", err))
	}

	// ------------- Required query parameter "keyword" -------------

	err = runtime.BindQueryParameter("form", true, true, "keyword", ctx.QueryParams(), &params.Keyword)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter keyword: %s", err))
	}

	// ------------- Optional query parameter "center_lat" -------------

	err = runtime.BindQueryParameter("form", true, false, "center_lat", ctx.QueryParams(), &params.CenterLat)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter center_lat: %s", err))
	}

	// ------------- Optional query parameter "center_lng" -------------

	err = runtime.BindQueryParameter("form", true, false, "center_lng", ctx.QueryParams(), &params.CenterLng)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter center_lng: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SearchCommunities(ctx, params)
	return err
}

// GetCommunityById converts echo context to params.
func (w *ServerInterfaceWrapper) GetCommunityById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCommunityById(ctx, id)
	return err
}

// ListCommunityRestaurants converts echo context to params.
func (w *ServerInterfaceWrapper) ListCommunityRestaurants(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(TokenScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListCommunityRestaurantsParams
	// ------------- Optional query parameter "after" -------------

	err = runtime.BindQueryParameter("form", true, false, "after", ctx.QueryParams(), &params.After)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter after: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListCommunityRestaurants(ctx, id, params)
	return err
}

// AddRestaurantToCommunity converts echo context to params.
func (w *ServerInterfaceWrapper) AddRestaurantToCommunity(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddRestaurantToCommunity(ctx, id)
	return err
}

// RemoveRestaurantFromCommunity converts echo context to params.
func (w *ServerInterfaceWrapper) RemoveRestaurantFromCommunity(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "restaurant_id" -------------
	var restaurantId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "restaurant_id", runtime.ParamLocationPath, ctx.Param("restaurant_id"), &restaurantId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter restaurant_id: %s", err))
	}

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RemoveRestaurantFromCommunity(ctx, id, restaurantId)
	return err
}

// GetRestaurantComment converts echo context to params.
func (w *ServerInterfaceWrapper) GetRestaurantComment(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "restaurant_id" -------------
	var restaurantId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "restaurant_id", runtime.ParamLocationPath, ctx.Param("restaurant_id"), &restaurantId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter restaurant_id: %s", err))
	}

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRestaurantComment(ctx, id, restaurantId)
	return err
}

// UpdateRestaurantComment converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateRestaurantComment(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "restaurant_id" -------------
	var restaurantId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "restaurant_id", runtime.ParamLocationPath, ctx.Param("restaurant_id"), &restaurantId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter restaurant_id: %s", err))
	}

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateRestaurantComment(ctx, id, restaurantId)
	return err
}

// GetCommunityIdToken converts echo context to params.
func (w *ServerInterfaceWrapper) GetCommunityIdToken(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCommunityIdToken(ctx, id)
	return err
}

// ListUsersOfCommunity converts echo context to params.
func (w *ServerInterfaceWrapper) ListUsersOfCommunity(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(TokenScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListUsersOfCommunityParams
	// ------------- Optional query parameter "after" -------------

	err = runtime.BindQueryParameter("form", true, false, "after", ctx.QueryParams(), &params.After)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter after: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListUsersOfCommunity(ctx, id, params)
	return err
}

// SearchRestaurants converts echo context to params.
func (w *ServerInterfaceWrapper) SearchRestaurants(ctx echo.Context) error {
	var err error

	ctx.Set(TokenScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchRestaurantsParams
	// ------------- Optional query parameter "after" -------------

	err = runtime.BindQueryParameter("form", true, false, "after", ctx.QueryParams(), &params.After)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter after: %s", err))
	}

	// ------------- Required query parameter "keyword" -------------

	err = runtime.BindQueryParameter("form", true, true, "keyword", ctx.QueryParams(), &params.Keyword)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter keyword: %s", err))
	}

	// ------------- Optional query parameter "center_lat" -------------

	err = runtime.BindQueryParameter("form", true, false, "center_lat", ctx.QueryParams(), &params.CenterLat)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter center_lat: %s", err))
	}

	// ------------- Optional query parameter "center_lng" -------------

	err = runtime.BindQueryParameter("form", true, false, "center_lng", ctx.QueryParams(), &params.CenterLng)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter center_lng: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SearchRestaurants(ctx, params)
	return err
}

// GetRestaurantId converts echo context to params.
func (w *ServerInterfaceWrapper) GetRestaurantId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRestaurantId(ctx, id)
	return err
}

// NewUser converts echo context to params.
func (w *ServerInterfaceWrapper) NewUser(ctx echo.Context) error {
	var err error

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.NewUser(ctx)
	return err
}

// GetMyProfile converts echo context to params.
func (w *ServerInterfaceWrapper) GetMyProfile(ctx echo.Context) error {
	var err error

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetMyProfile(ctx)
	return err
}

// PostUserMeCommunities converts echo context to params.
func (w *ServerInterfaceWrapper) PostUserMeCommunities(ctx echo.Context) error {
	var err error

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostUserMeCommunities(ctx)
	return err
}

// UploadProfileImage converts echo context to params.
func (w *ServerInterfaceWrapper) UploadProfileImage(ctx echo.Context) error {
	var err error

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UploadProfileImage(ctx)
	return err
}

// GetUserIdBookmark converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserIdBookmark(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id Long

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUserIdBookmark(ctx, id)
	return err
}

// PostUserIdBookmark converts echo context to params.
func (w *ServerInterfaceWrapper) PostUserIdBookmark(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id Long

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostUserIdBookmark(ctx, id)
	return err
}

// DeleteUserIdBookmarkCommunityId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUserIdBookmarkCommunityId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id Long

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "community_id" -------------
	var communityId Long

	err = runtime.BindStyledParameterWithLocation("simple", false, "community_id", runtime.ParamLocationPath, ctx.Param("community_id"), &communityId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter community_id: %s", err))
	}

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteUserIdBookmarkCommunityId(ctx, id, communityId)
	return err
}

// ListUserCommunities converts echo context to params.
func (w *ServerInterfaceWrapper) ListUserCommunities(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id Long

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(TokenScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListUserCommunitiesParams
	// ------------- Optional query parameter "after" -------------

	err = runtime.BindQueryParameter("form", true, false, "after", ctx.QueryParams(), &params.After)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter after: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListUserCommunities(ctx, id, params)
	return err
}

// DeleteUserIdCommunitiesCommunityId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUserIdCommunitiesCommunityId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id Long

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// ------------- Path parameter "community_id" -------------
	var communityId Long

	err = runtime.BindStyledParameterWithLocation("simple", false, "community_id", runtime.ParamLocationPath, ctx.Param("community_id"), &communityId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter community_id: %s", err))
	}

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteUserIdCommunitiesCommunityId(ctx, id, communityId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/community", wrapper.NewCommunity)
	router.GET(baseURL+"/community/search", wrapper.SearchCommunities)
	router.GET(baseURL+"/community/:id", wrapper.GetCommunityById)
	router.GET(baseURL+"/community/:id/restaurants", wrapper.ListCommunityRestaurants)
	router.POST(baseURL+"/community/:id/restaurants", wrapper.AddRestaurantToCommunity)
	router.DELETE(baseURL+"/community/:id/restaurants/:restaurant_id", wrapper.RemoveRestaurantFromCommunity)
	router.GET(baseURL+"/community/:id/restaurants/:restaurant_id/comments", wrapper.GetRestaurantComment)
	router.PUT(baseURL+"/community/:id/restaurants/:restaurant_id/comments", wrapper.UpdateRestaurantComment)
	router.GET(baseURL+"/community/:id/token", wrapper.GetCommunityIdToken)
	router.GET(baseURL+"/community/:id/users", wrapper.ListUsersOfCommunity)
	router.GET(baseURL+"/restaurant/search", wrapper.SearchRestaurants)
	router.GET(baseURL+"/restaurant/:id", wrapper.GetRestaurantId)
	router.POST(baseURL+"/user", wrapper.NewUser)
	router.GET(baseURL+"/user/me", wrapper.GetMyProfile)
	router.POST(baseURL+"/user/me/communities", wrapper.PostUserMeCommunities)
	router.POST(baseURL+"/user/profile", wrapper.UploadProfileImage)
	router.GET(baseURL+"/user/:id/bookmark", wrapper.GetUserIdBookmark)
	router.POST(baseURL+"/user/:id/bookmark", wrapper.PostUserIdBookmark)
	router.DELETE(baseURL+"/user/:id/bookmark/:community_id", wrapper.DeleteUserIdBookmarkCommunityId)
	router.GET(baseURL+"/user/:id/communities", wrapper.ListUserCommunities)
	router.DELETE(baseURL+"/user/:id/communities/:community_id", wrapper.DeleteUserIdCommunitiesCommunityId)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
