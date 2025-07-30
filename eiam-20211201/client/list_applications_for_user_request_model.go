// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *ListApplicationsForUserRequest
	GetApplicationIds() []*string
	SetInstanceId(v string) *ListApplicationsForUserRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *ListApplicationsForUserRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListApplicationsForUserRequest
	GetPageSize() *int64
	SetQueryMode(v string) *ListApplicationsForUserRequest
	GetQueryMode() *string
	SetUserId(v string) *ListApplicationsForUserRequest
	GetUserId() *string
}

type ListApplicationsForUserRequest struct {
	// The IDs of the applications that the EIAM account can access. You can query a maximum of 100 application IDs at a time.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query mode. Default value: **OnlyDirect**. Valid values:
	//
	// 	- OnlyDirect: Only the direct permissions are queried. Direct permissions are the permissions that are directly granted to the account.
	//
	// 	- IncludeInherit: Both the permissions that are directly granted to the account and the inherited permissions are queried. Inherited permissions are the permissions that an account inherits from the parent organization or the group to which the account belongs.
	//
	// example:
	//
	// OnlyDirect
	QueryMode *string `json:"QueryMode,omitempty" xml:"QueryMode,omitempty"`
	// The ID of the EIAM account.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListApplicationsForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsForUserRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *ListApplicationsForUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationsForUserRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListApplicationsForUserRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApplicationsForUserRequest) GetQueryMode() *string {
	return s.QueryMode
}

func (s *ListApplicationsForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListApplicationsForUserRequest) SetApplicationIds(v []*string) *ListApplicationsForUserRequest {
	s.ApplicationIds = v
	return s
}

func (s *ListApplicationsForUserRequest) SetInstanceId(v string) *ListApplicationsForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationsForUserRequest) SetPageNumber(v int64) *ListApplicationsForUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsForUserRequest) SetPageSize(v int64) *ListApplicationsForUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsForUserRequest) SetQueryMode(v string) *ListApplicationsForUserRequest {
	s.QueryMode = &v
	return s
}

func (s *ListApplicationsForUserRequest) SetUserId(v string) *ListApplicationsForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListApplicationsForUserRequest) Validate() error {
	return dara.Validate(s)
}
