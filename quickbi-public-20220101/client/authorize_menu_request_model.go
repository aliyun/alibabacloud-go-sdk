// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeMenuRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthPointsValue(v int32) *AuthorizeMenuRequest
	GetAuthPointsValue() *int32
	SetDataPortalId(v string) *AuthorizeMenuRequest
	GetDataPortalId() *string
	SetMenuIds(v string) *AuthorizeMenuRequest
	GetMenuIds() *string
	SetUserGroupIds(v string) *AuthorizeMenuRequest
	GetUserGroupIds() *string
	SetUserIds(v string) *AuthorizeMenuRequest
	GetUserIds() *string
}

type AuthorizeMenuRequest struct {
	// The permissions to grant for the menus. Valid values:
	//
	// - 1: View.
	//
	// - 3: View and export. This is the default value.
	//
	// example:
	//
	// 3
	AuthPointsValue *int32 `json:"AuthPointsValue,omitempty" xml:"AuthPointsValue,omitempty"`
	// The ID of the data portal.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0d173abb53e84c8ca7495429163b****
	DataPortalId *string `json:"DataPortalId,omitempty" xml:"DataPortalId,omitempty"`
	// The IDs of the leaf node menus in the data portal. You can specify multiple IDs. Separate multiple IDs with commas (,). You can specify up to 100 IDs at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54kqgoa****,pg1n135****
	MenuIds *string `json:"MenuIds,omitempty" xml:"MenuIds,omitempty"`
	// The IDs of the user groups. You can specify multiple IDs. Separate multiple IDs with commas (,). You can specify up to 200 IDs at a time. You must specify at least one of UserGroupIds and UserIds.
	//
	// example:
	//
	// 34fd141d-4598-4093-8c33-8e066dcb****,3d2c23d4-2b41-4af8-a1f5-f6390f32****
	UserGroupIds *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	// The IDs of the users. These are Quick BI user IDs, not Alibaba Cloud UIDs. You can specify multiple IDs. Separate multiple IDs with commas (,). You can specify up to 200 IDs at a time.
	//
	// example:
	//
	// 204627493484****,121344444790****
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s AuthorizeMenuRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeMenuRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeMenuRequest) GetAuthPointsValue() *int32 {
	return s.AuthPointsValue
}

func (s *AuthorizeMenuRequest) GetDataPortalId() *string {
	return s.DataPortalId
}

func (s *AuthorizeMenuRequest) GetMenuIds() *string {
	return s.MenuIds
}

func (s *AuthorizeMenuRequest) GetUserGroupIds() *string {
	return s.UserGroupIds
}

func (s *AuthorizeMenuRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *AuthorizeMenuRequest) SetAuthPointsValue(v int32) *AuthorizeMenuRequest {
	s.AuthPointsValue = &v
	return s
}

func (s *AuthorizeMenuRequest) SetDataPortalId(v string) *AuthorizeMenuRequest {
	s.DataPortalId = &v
	return s
}

func (s *AuthorizeMenuRequest) SetMenuIds(v string) *AuthorizeMenuRequest {
	s.MenuIds = &v
	return s
}

func (s *AuthorizeMenuRequest) SetUserGroupIds(v string) *AuthorizeMenuRequest {
	s.UserGroupIds = &v
	return s
}

func (s *AuthorizeMenuRequest) SetUserIds(v string) *AuthorizeMenuRequest {
	s.UserIds = &v
	return s
}

func (s *AuthorizeMenuRequest) Validate() error {
	return dara.Validate(s)
}
