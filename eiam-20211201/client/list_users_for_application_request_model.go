// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersForApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListUsersForApplicationRequest
	GetApplicationId() *string
	SetApplicationRoleId(v string) *ListUsersForApplicationRequest
	GetApplicationRoleId() *string
	SetInstanceId(v string) *ListUsersForApplicationRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *ListUsersForApplicationRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListUsersForApplicationRequest
	GetPageSize() *int64
	SetUserIds(v []*string) *ListUsersForApplicationRequest
	GetUserIds() []*string
}

type ListUsersForApplicationRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 应用角色ID。
	//
	// example:
	//
	// app_role_mkv7rgt4ds8d8v0qtzev2mxxxx
	ApplicationRoleId *string `json:"ApplicationRoleId,omitempty" xml:"ApplicationRoleId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of account IDs. A maximum of 100 account IDs are supported in a single request.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s ListUsersForApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListUsersForApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListUsersForApplicationRequest) GetApplicationRoleId() *string {
	return s.ApplicationRoleId
}

func (s *ListUsersForApplicationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUsersForApplicationRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListUsersForApplicationRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUsersForApplicationRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *ListUsersForApplicationRequest) SetApplicationId(v string) *ListUsersForApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListUsersForApplicationRequest) SetApplicationRoleId(v string) *ListUsersForApplicationRequest {
	s.ApplicationRoleId = &v
	return s
}

func (s *ListUsersForApplicationRequest) SetInstanceId(v string) *ListUsersForApplicationRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUsersForApplicationRequest) SetPageNumber(v int64) *ListUsersForApplicationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersForApplicationRequest) SetPageSize(v int64) *ListUsersForApplicationRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersForApplicationRequest) SetUserIds(v []*string) *ListUsersForApplicationRequest {
	s.UserIds = v
	return s
}

func (s *ListUsersForApplicationRequest) Validate() error {
	return dara.Validate(s)
}
