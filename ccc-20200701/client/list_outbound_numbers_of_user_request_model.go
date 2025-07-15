// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutboundNumbersOfUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListOutboundNumbersOfUserRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListOutboundNumbersOfUserRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOutboundNumbersOfUserRequest
	GetPageSize() *int32
	SetSkillGroupIdList(v string) *ListOutboundNumbersOfUserRequest
	GetSkillGroupIdList() *string
	SetUserId(v string) *ListOutboundNumbersOfUserRequest
	GetUserId() *string
}

type ListOutboundNumbersOfUserRequest struct {
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
	// ["skillgroup1@ccc-test","skillgroup2@ccc-test"]
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListOutboundNumbersOfUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundNumbersOfUserRequest) GoString() string {
	return s.String()
}

func (s *ListOutboundNumbersOfUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOutboundNumbersOfUserRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOutboundNumbersOfUserRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOutboundNumbersOfUserRequest) GetSkillGroupIdList() *string {
	return s.SkillGroupIdList
}

func (s *ListOutboundNumbersOfUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListOutboundNumbersOfUserRequest) SetInstanceId(v string) *ListOutboundNumbersOfUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOutboundNumbersOfUserRequest) SetPageNumber(v int32) *ListOutboundNumbersOfUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOutboundNumbersOfUserRequest) SetPageSize(v int32) *ListOutboundNumbersOfUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListOutboundNumbersOfUserRequest) SetSkillGroupIdList(v string) *ListOutboundNumbersOfUserRequest {
	s.SkillGroupIdList = &v
	return s
}

func (s *ListOutboundNumbersOfUserRequest) SetUserId(v string) *ListOutboundNumbersOfUserRequest {
	s.UserId = &v
	return s
}

func (s *ListOutboundNumbersOfUserRequest) Validate() error {
	return dara.Validate(s)
}
