// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBriefSkillGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListBriefSkillGroupsRequest
	GetInstanceId() *string
	SetMediaType(v string) *ListBriefSkillGroupsRequest
	GetMediaType() *string
	SetPageNumber(v int32) *ListBriefSkillGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBriefSkillGroupsRequest
	GetPageSize() *int32
	SetSearchPattern(v string) *ListBriefSkillGroupsRequest
	GetSearchPattern() *string
}

type ListBriefSkillGroupsRequest struct {
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
	// skillgroup
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
}

func (s ListBriefSkillGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBriefSkillGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListBriefSkillGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBriefSkillGroupsRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ListBriefSkillGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBriefSkillGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBriefSkillGroupsRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListBriefSkillGroupsRequest) SetInstanceId(v string) *ListBriefSkillGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBriefSkillGroupsRequest) SetMediaType(v string) *ListBriefSkillGroupsRequest {
	s.MediaType = &v
	return s
}

func (s *ListBriefSkillGroupsRequest) SetPageNumber(v int32) *ListBriefSkillGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBriefSkillGroupsRequest) SetPageSize(v int32) *ListBriefSkillGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListBriefSkillGroupsRequest) SetSearchPattern(v string) *ListBriefSkillGroupsRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListBriefSkillGroupsRequest) Validate() error {
	return dara.Validate(s)
}
