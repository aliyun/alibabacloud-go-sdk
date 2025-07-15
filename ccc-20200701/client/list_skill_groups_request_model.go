// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListSkillGroupsRequest
	GetInstanceId() *string
	SetMediaType(v string) *ListSkillGroupsRequest
	GetMediaType() *string
	SetPageNumber(v int32) *ListSkillGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSkillGroupsRequest
	GetPageSize() *int32
	SetSearchPattern(v string) *ListSkillGroupsRequest
	GetSearchPattern() *string
}

type ListSkillGroupsRequest struct {
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
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
}

func (s ListSkillGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListSkillGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSkillGroupsRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ListSkillGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSkillGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillGroupsRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListSkillGroupsRequest) SetInstanceId(v string) *ListSkillGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSkillGroupsRequest) SetMediaType(v string) *ListSkillGroupsRequest {
	s.MediaType = &v
	return s
}

func (s *ListSkillGroupsRequest) SetPageNumber(v int32) *ListSkillGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSkillGroupsRequest) SetPageSize(v int32) *ListSkillGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSkillGroupsRequest) SetSearchPattern(v string) *ListSkillGroupsRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListSkillGroupsRequest) Validate() error {
	return dara.Validate(s)
}
