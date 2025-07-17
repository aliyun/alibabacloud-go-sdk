// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListUsersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUsersRequest
	GetPageSize() *int32
	SetRole(v string) *ListUsersRequest
	GetRole() *string
	SetSearchKey(v string) *ListUsersRequest
	GetSearchKey() *string
	SetTid(v int64) *ListUsersRequest
	GetTid() *int64
	SetUserState(v string) *ListUsersRequest
	GetUserState() *string
}

type ListUsersRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// **
	//
	// Valid values: 10, 20, 50, and 100.***	- Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The role that is assigned to the user. Valid values:
	//
	// 	- **USER**: a regular user.
	//
	// 	- **DBA*	- : a database administrator (DBA).
	//
	// 	- **ADMIN**: a Data Management (DMS) administrator.
	//
	// 	- **SECURITY_ADMIN**: a security administrator.
	//
	// 	- **STRUCT_READ_ONLY**: a schema read-only user.
	//
	// >  To check your role, move the pointer over the profile picture in the upper-right corner of the DMS console.
	//
	// example:
	//
	// DBA
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The search keyword. Fuzzy match is supported.
	//
	// example:
	//
	// test
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The status of the user. Valid values:
	//
	// 	- **NORMAL**: The user is normal.
	//
	// 	- **DISABLE**: The user is disabled.
	//
	// 	- **DELETE**: The user is deleted.
	//
	// example:
	//
	// NORMAL
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
}

func (s ListUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUsersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUsersRequest) GetRole() *string {
	return s.Role
}

func (s *ListUsersRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListUsersRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListUsersRequest) GetUserState() *string {
	return s.UserState
}

func (s *ListUsersRequest) SetPageNumber(v int32) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequest) SetRole(v string) *ListUsersRequest {
	s.Role = &v
	return s
}

func (s *ListUsersRequest) SetSearchKey(v string) *ListUsersRequest {
	s.SearchKey = &v
	return s
}

func (s *ListUsersRequest) SetTid(v int64) *ListUsersRequest {
	s.Tid = &v
	return s
}

func (s *ListUsersRequest) SetUserState(v string) *ListUsersRequest {
	s.UserState = &v
	return s
}

func (s *ListUsersRequest) Validate() error {
	return dara.Validate(s)
}
