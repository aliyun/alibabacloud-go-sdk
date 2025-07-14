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
	// Authorizes the permissions of the menu. Valid values:
	//
	// 	- 1: view
	//
	// 	- 3: View + Export (default)
	//
	// example:
	//
	// 3
	AuthPointsValue *int32 `json:"AuthPointsValue,omitempty" xml:"AuthPointsValue,omitempty"`
	// The ID of the BI portal.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0d173abb53e84c8ca7495429163b****
	DataPortalId *string `json:"DataPortalId,omitempty" xml:"DataPortalId,omitempty"`
	// The menu ID of the BI portal leaf node.
	//
	// 	- The directory menu cannot be authorized.
	//
	// 	- You can upload multiple parameters at a time. Separate multiple IDs with commas (,). The maximum number of parameters that can be modified at a time is 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54kqgoa****,pg1n135****
	MenuIds *string `json:"MenuIds,omitempty" xml:"MenuIds,omitempty"`
	// The IDs of the user groups.
	//
	// 	- You can upload multiple parameters at a time. Separate multiple IDs with commas (,). The maximum number of parameters that can be modified at a time is 200.
	//
	// 	- UserGroupIds and UserIds cannot be empty at the same time
	//
	// example:
	//
	// 34fd141d-4598-4093-8c33-8e066dcb****,3d2c23d4-2b41-4af8-a1f5-f6390f32****
	UserGroupIds *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	// The IDs of the end users. The UserID of the Quick BI is used instead of the UID of Alibaba Cloud.
	//
	// 	- You can upload multiple parameters at a time. Separate multiple IDs with commas (,). The maximum number of parameters that can be modified at a time is 200.
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
