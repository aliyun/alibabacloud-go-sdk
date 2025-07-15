// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillGroupStatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListSkillGroupStatesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListSkillGroupStatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSkillGroupStatesRequest
	GetPageSize() *int32
	SetSkillGroupIds(v string) *ListSkillGroupStatesRequest
	GetSkillGroupIds() *string
}

type ListSkillGroupStatesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ["skillgroup1@ccc-test","skillgroup2@ccc-test"]
	SkillGroupIds *string `json:"SkillGroupIds,omitempty" xml:"SkillGroupIds,omitempty"`
}

func (s ListSkillGroupStatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupStatesRequest) GoString() string {
	return s.String()
}

func (s *ListSkillGroupStatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSkillGroupStatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSkillGroupStatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillGroupStatesRequest) GetSkillGroupIds() *string {
	return s.SkillGroupIds
}

func (s *ListSkillGroupStatesRequest) SetInstanceId(v string) *ListSkillGroupStatesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSkillGroupStatesRequest) SetPageNumber(v int32) *ListSkillGroupStatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSkillGroupStatesRequest) SetPageSize(v int32) *ListSkillGroupStatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSkillGroupStatesRequest) SetSkillGroupIds(v string) *ListSkillGroupStatesRequest {
	s.SkillGroupIds = &v
	return s
}

func (s *ListSkillGroupStatesRequest) Validate() error {
	return dara.Validate(s)
}
