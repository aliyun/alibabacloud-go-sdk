// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListUsersRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListUsersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUsersRequest
	GetPageSize() *int32
	SetSearchPattern(v string) *ListUsersRequest
	GetSearchPattern() *string
	SetSkillGroupId(v string) *ListUsersRequest
	GetSkillGroupId() *string
}

type ListUsersRequest struct {
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
	// agent
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	SkillGroupId  *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s ListUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUsersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUsersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUsersRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListUsersRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListUsersRequest) SetInstanceId(v string) *ListUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v int32) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequest) SetSearchPattern(v string) *ListUsersRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListUsersRequest) SetSkillGroupId(v string) *ListUsersRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListUsersRequest) Validate() error {
	return dara.Validate(s)
}
