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
	Description *string   `json:"description,omitempty"`
	Id          Long      `json:"id"`
	Location    *Location `json:"location,omitempty"`
	Name        string    `json:"name"`
}

// CommunityDetail defines model for community_detail.
type CommunityDetail struct {
	// Embedded struct due to allOf(#/components/schemas/community)
	Community `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	NumRestaurant *int `json:"num_restaurant,omitempty"`
	UserCount     int  `json:"user_count"`
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

// JoinCommunityRequest defines model for joinCommunityRequest.
type JoinCommunityRequest struct {
	UserId Long `json:"user_id"`
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
	After   *PageQuery `json:"after,omitempty"`
	Keyword string     `json:"keyword"`
	Center  *Location  `json:"center,omitempty"`
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
	After   *PageQuery `json:"after,omitempty"`
	Keyword string     `json:"keyword"`
	Center  *Location  `json:"center,omitempty"`
}

// NewUserJSONBody defines parameters for NewUser.
type NewUserJSONBody CreateUserRequest

// PostUserIdBookmarkJSONBody defines parameters for PostUserIdBookmark.
type PostUserIdBookmarkJSONBody struct {
	CommunityId Long `json:"community_id"`
}

// ListUserCommunitiesParams defines parameters for ListUserCommunities.
type ListUserCommunitiesParams struct {
	After *PageQuery `json:"after,omitempty"`
}

// PostUserIdCommunitiesJSONBody defines parameters for PostUserIdCommunities.
type PostUserIdCommunitiesJSONBody JoinCommunityRequest

// NewCommunityJSONRequestBody defines body for NewCommunity for application/json ContentType.
type NewCommunityJSONRequestBody NewCommunityJSONBody

// AddRestaurantToCommunityJSONRequestBody defines body for AddRestaurantToCommunity for application/json ContentType.
type AddRestaurantToCommunityJSONRequestBody AddRestaurantToCommunityJSONBody

// UpdateRestaurantCommentJSONRequestBody defines body for UpdateRestaurantComment for application/json ContentType.
type UpdateRestaurantCommentJSONRequestBody UpdateRestaurantCommentJSONBody

// NewUserJSONRequestBody defines body for NewUser for application/json ContentType.
type NewUserJSONRequestBody NewUserJSONBody

// PostUserIdBookmarkJSONRequestBody defines body for PostUserIdBookmark for application/json ContentType.
type PostUserIdBookmarkJSONRequestBody PostUserIdBookmarkJSONBody

// PostUserIdCommunitiesJSONRequestBody defines body for PostUserIdCommunities for application/json ContentType.
type PostUserIdCommunitiesJSONRequestBody PostUserIdCommunitiesJSONBody

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
	// Join a community
	// (POST /user/{id}/communities)
	PostUserIdCommunities(ctx echo.Context, id Long) error
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

	// ------------- Optional query parameter "center" -------------

	err = runtime.BindQueryParameter("form", true, false, "center", ctx.QueryParams(), &params.Center)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter center: %s", err))
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

	// ------------- Optional query parameter "center" -------------

	err = runtime.BindQueryParameter("form", true, false, "center", ctx.QueryParams(), &params.Center)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter center: %s", err))
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

// PostUserIdCommunities converts echo context to params.
func (w *ServerInterfaceWrapper) PostUserIdCommunities(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id Long

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(TokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostUserIdCommunities(ctx, id)
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
	router.GET(baseURL+"/community/:id/users", wrapper.ListUsersOfCommunity)
	router.GET(baseURL+"/restaurant/search", wrapper.SearchRestaurants)
	router.GET(baseURL+"/restaurant/:id", wrapper.GetRestaurantId)
	router.POST(baseURL+"/user", wrapper.NewUser)
	router.GET(baseURL+"/user/me", wrapper.GetMyProfile)
	router.POST(baseURL+"/user/profile", wrapper.UploadProfileImage)
	router.GET(baseURL+"/user/:id/bookmark", wrapper.GetUserIdBookmark)
	router.POST(baseURL+"/user/:id/bookmark", wrapper.PostUserIdBookmark)
	router.DELETE(baseURL+"/user/:id/bookmark/:community_id", wrapper.DeleteUserIdBookmarkCommunityId)
	router.GET(baseURL+"/user/:id/communities", wrapper.ListUserCommunities)
	router.POST(baseURL+"/user/:id/communities", wrapper.PostUserIdCommunities)
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
