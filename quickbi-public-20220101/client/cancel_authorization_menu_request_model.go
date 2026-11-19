// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAuthorizationMenuRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataPortalId(v string) *CancelAuthorizationMenuRequest
	GetDataPortalId() *string
	SetMenuIds(v string) *CancelAuthorizationMenuRequest
	GetMenuIds() *string
	SetUserGroupIds(v string) *CancelAuthorizationMenuRequest
	GetUserGroupIds() *string
	SetUserIds(v string) *CancelAuthorizationMenuRequest
	GetUserIds() *string
}

type CancelAuthorizationMenuRequest struct {
	// The ID of the data portal.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0d173abb53e84c8ca7495429163b****
	DataPortalId *string `json:"DataPortalId,omitempty" xml:"DataPortalId,omitempty"`
	// The IDs of the leaf node menus in the data portal.
	//
	// - You can specify multiple IDs. Separate multiple IDs with commas (,). A maximum of 100 IDs can be specified at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54kqgoa****,pg1n135****
	MenuIds *string `json:"MenuIds,omitempty" xml:"MenuIds,omitempty"`
	// The IDs of the user groups.
	//
	// - You can specify multiple IDs. Separate multiple IDs with commas (,). A maximum of 200 IDs can be specified at a time.
	//
	// - UserGroupIds and UserIds cannot be empty at the same time.
	//
	// example:
	//
	// 34fd141d-4598-4093-8c33-8e066dcb****,3d2c23d4-2b41-4af8-a1f5-f6390f32****
	UserGroupIds *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	// The IDs of the users. These are Quick BI user IDs, not Alibaba Cloud UIDs.
	//
	// - You can specify multiple IDs. Separate multiple IDs with commas (,). A maximum of 200 IDs can be specified at a time.
	//
	// example:
	//
	// 204627493484****,121344444790****
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s CancelAuthorizationMenuRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAuthorizationMenuRequest) GoString() string {
	return s.String()
}

func (s *CancelAuthorizationMenuRequest) GetDataPortalId() *string {
	return s.DataPortalId
}

func (s *CancelAuthorizationMenuRequest) GetMenuIds() *string {
	return s.MenuIds
}

func (s *CancelAuthorizationMenuRequest) GetUserGroupIds() *string {
	return s.UserGroupIds
}

func (s *CancelAuthorizationMenuRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *CancelAuthorizationMenuRequest) SetDataPortalId(v string) *CancelAuthorizationMenuRequest {
	s.DataPortalId = &v
	return s
}

func (s *CancelAuthorizationMenuRequest) SetMenuIds(v string) *CancelAuthorizationMenuRequest {
	s.MenuIds = &v
	return s
}

func (s *CancelAuthorizationMenuRequest) SetUserGroupIds(v string) *CancelAuthorizationMenuRequest {
	s.UserGroupIds = &v
	return s
}

func (s *CancelAuthorizationMenuRequest) SetUserIds(v string) *CancelAuthorizationMenuRequest {
	s.UserIds = &v
	return s
}

func (s *CancelAuthorizationMenuRequest) Validate() error {
	return dara.Validate(s)
}
