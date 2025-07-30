// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListGroupsForUserRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *ListGroupsForUserRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListGroupsForUserRequest
	GetPageSize() *int64
	SetUserId(v string) *ListGroupsForUserRequest
	GetUserId() *string
}

type ListGroupsForUserRequest struct {
	// The instance ID.
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
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListGroupsForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGroupsForUserRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListGroupsForUserRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListGroupsForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListGroupsForUserRequest) SetInstanceId(v string) *ListGroupsForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListGroupsForUserRequest) SetPageNumber(v int64) *ListGroupsForUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGroupsForUserRequest) SetPageSize(v int64) *ListGroupsForUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupsForUserRequest) SetUserId(v string) *ListGroupsForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListGroupsForUserRequest) Validate() error {
	return dara.Validate(s)
}
