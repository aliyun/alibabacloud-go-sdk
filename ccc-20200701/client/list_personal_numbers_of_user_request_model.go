// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPersonalNumbersOfUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListPersonalNumbersOfUserRequest
	GetInstanceId() *string
	SetIsMember(v bool) *ListPersonalNumbersOfUserRequest
	GetIsMember() *bool
	SetPageNumber(v int32) *ListPersonalNumbersOfUserRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPersonalNumbersOfUserRequest
	GetPageSize() *int32
	SetSearchPattern(v string) *ListPersonalNumbersOfUserRequest
	GetSearchPattern() *string
	SetUserId(v string) *ListPersonalNumbersOfUserRequest
	GetUserId() *string
}

type ListPersonalNumbersOfUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	IsMember *bool `json:"IsMember,omitempty" xml:"IsMember,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0833
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user-test@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListPersonalNumbersOfUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalNumbersOfUserRequest) GoString() string {
	return s.String()
}

func (s *ListPersonalNumbersOfUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPersonalNumbersOfUserRequest) GetIsMember() *bool {
	return s.IsMember
}

func (s *ListPersonalNumbersOfUserRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPersonalNumbersOfUserRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPersonalNumbersOfUserRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListPersonalNumbersOfUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListPersonalNumbersOfUserRequest) SetInstanceId(v string) *ListPersonalNumbersOfUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListPersonalNumbersOfUserRequest) SetIsMember(v bool) *ListPersonalNumbersOfUserRequest {
	s.IsMember = &v
	return s
}

func (s *ListPersonalNumbersOfUserRequest) SetPageNumber(v int32) *ListPersonalNumbersOfUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPersonalNumbersOfUserRequest) SetPageSize(v int32) *ListPersonalNumbersOfUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListPersonalNumbersOfUserRequest) SetSearchPattern(v string) *ListPersonalNumbersOfUserRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListPersonalNumbersOfUserRequest) SetUserId(v string) *ListPersonalNumbersOfUserRequest {
	s.UserId = &v
	return s
}

func (s *ListPersonalNumbersOfUserRequest) Validate() error {
	return dara.Validate(s)
}
