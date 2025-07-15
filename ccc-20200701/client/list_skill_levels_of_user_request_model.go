// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillLevelsOfUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListSkillLevelsOfUserRequest
	GetInstanceId() *string
	SetIsMember(v bool) *ListSkillLevelsOfUserRequest
	GetIsMember() *bool
	SetPageNumber(v int32) *ListSkillLevelsOfUserRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSkillLevelsOfUserRequest
	GetPageSize() *int32
	SetSearchPattern(v string) *ListSkillLevelsOfUserRequest
	GetSearchPattern() *string
	SetUserId(v string) *ListSkillLevelsOfUserRequest
	GetUserId() *string
}

type ListSkillLevelsOfUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// skillgroup
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListSkillLevelsOfUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillLevelsOfUserRequest) GoString() string {
	return s.String()
}

func (s *ListSkillLevelsOfUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSkillLevelsOfUserRequest) GetIsMember() *bool {
	return s.IsMember
}

func (s *ListSkillLevelsOfUserRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSkillLevelsOfUserRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillLevelsOfUserRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListSkillLevelsOfUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListSkillLevelsOfUserRequest) SetInstanceId(v string) *ListSkillLevelsOfUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSkillLevelsOfUserRequest) SetIsMember(v bool) *ListSkillLevelsOfUserRequest {
	s.IsMember = &v
	return s
}

func (s *ListSkillLevelsOfUserRequest) SetPageNumber(v int32) *ListSkillLevelsOfUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSkillLevelsOfUserRequest) SetPageSize(v int32) *ListSkillLevelsOfUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListSkillLevelsOfUserRequest) SetSearchPattern(v string) *ListSkillLevelsOfUserRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListSkillLevelsOfUserRequest) SetUserId(v string) *ListSkillLevelsOfUserRequest {
	s.UserId = &v
	return s
}

func (s *ListSkillLevelsOfUserRequest) Validate() error {
	return dara.Validate(s)
}
