// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserLevelsOfSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListUserLevelsOfSkillGroupRequest
	GetInstanceId() *string
	SetIsMember(v bool) *ListUserLevelsOfSkillGroupRequest
	GetIsMember() *bool
	SetPageNumber(v int32) *ListUserLevelsOfSkillGroupRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserLevelsOfSkillGroupRequest
	GetPageSize() *int32
	SetSearchPattern(v string) *ListUserLevelsOfSkillGroupRequest
	GetSearchPattern() *string
	SetSkillGroupId(v string) *ListUserLevelsOfSkillGroupRequest
	GetSkillGroupId() *string
}

type ListUserLevelsOfSkillGroupRequest struct {
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
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s ListUserLevelsOfSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserLevelsOfSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *ListUserLevelsOfSkillGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserLevelsOfSkillGroupRequest) GetIsMember() *bool {
	return s.IsMember
}

func (s *ListUserLevelsOfSkillGroupRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserLevelsOfSkillGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserLevelsOfSkillGroupRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListUserLevelsOfSkillGroupRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListUserLevelsOfSkillGroupRequest) SetInstanceId(v string) *ListUserLevelsOfSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupRequest) SetIsMember(v bool) *ListUserLevelsOfSkillGroupRequest {
	s.IsMember = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupRequest) SetPageNumber(v int32) *ListUserLevelsOfSkillGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupRequest) SetPageSize(v int32) *ListUserLevelsOfSkillGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupRequest) SetSearchPattern(v string) *ListUserLevelsOfSkillGroupRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupRequest) SetSkillGroupId(v string) *ListUserLevelsOfSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
