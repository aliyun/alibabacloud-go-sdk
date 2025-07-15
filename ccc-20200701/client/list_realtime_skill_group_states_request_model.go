// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeSkillGroupStatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListRealtimeSkillGroupStatesRequest
	GetInstanceId() *string
	SetMediaType(v string) *ListRealtimeSkillGroupStatesRequest
	GetMediaType() *string
	SetPageNumber(v int32) *ListRealtimeSkillGroupStatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRealtimeSkillGroupStatesRequest
	GetPageSize() *int32
	SetSkillGroupIdList(v string) *ListRealtimeSkillGroupStatesRequest
	GetSkillGroupIdList() *string
}

type ListRealtimeSkillGroupStatesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MediaType  *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
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
	// ["skillgroup1@ccc-test", "skillgroup2@ccc-test"]
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
}

func (s ListRealtimeSkillGroupStatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeSkillGroupStatesRequest) GoString() string {
	return s.String()
}

func (s *ListRealtimeSkillGroupStatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRealtimeSkillGroupStatesRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ListRealtimeSkillGroupStatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRealtimeSkillGroupStatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRealtimeSkillGroupStatesRequest) GetSkillGroupIdList() *string {
	return s.SkillGroupIdList
}

func (s *ListRealtimeSkillGroupStatesRequest) SetInstanceId(v string) *ListRealtimeSkillGroupStatesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesRequest) SetMediaType(v string) *ListRealtimeSkillGroupStatesRequest {
	s.MediaType = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesRequest) SetPageNumber(v int32) *ListRealtimeSkillGroupStatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesRequest) SetPageSize(v int32) *ListRealtimeSkillGroupStatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesRequest) SetSkillGroupIdList(v string) *ListRealtimeSkillGroupStatesRequest {
	s.SkillGroupIdList = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesRequest) Validate() error {
	return dara.Validate(s)
}
