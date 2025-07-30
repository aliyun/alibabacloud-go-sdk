// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersForGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListUsersForGroupRequest
	GetGroupId() *string
	SetInstanceId(v string) *ListUsersForGroupRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *ListUsersForGroupRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListUsersForGroupRequest
	GetPageSize() *int64
	SetUserIds(v []*string) *ListUsersForGroupRequest
	GetUserIds() []*string
}

type ListUsersForGroupRequest struct {
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The account IDs. A maximum of 100 accounts can be queried.
	//
	// example:
	//
	// [ou_001]
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s ListUsersForGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForGroupRequest) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListUsersForGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUsersForGroupRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListUsersForGroupRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUsersForGroupRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *ListUsersForGroupRequest) SetGroupId(v string) *ListUsersForGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ListUsersForGroupRequest) SetInstanceId(v string) *ListUsersForGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUsersForGroupRequest) SetPageNumber(v int64) *ListUsersForGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersForGroupRequest) SetPageSize(v int64) *ListUsersForGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersForGroupRequest) SetUserIds(v []*string) *ListUsersForGroupRequest {
	s.UserIds = v
	return s
}

func (s *ListUsersForGroupRequest) Validate() error {
	return dara.Validate(s)
}
