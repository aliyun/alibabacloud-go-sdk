// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListModulesRequest
	GetGroupId() *string
	SetKeyword(v string) *ListModulesRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListModulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModulesRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListModulesRequest
	GetProjectId() *string
	SetTag(v []*ListModulesRequestTag) *ListModulesRequest
	GetTag() []*ListModulesRequestTag
}

type ListModulesRequest struct {
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize  *int32                   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProjectId *string                  `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Tag       []*ListModulesRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s ListModulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModulesRequest) GoString() string {
	return s.String()
}

func (s *ListModulesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListModulesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListModulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModulesRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListModulesRequest) GetTag() []*ListModulesRequestTag {
	return s.Tag
}

func (s *ListModulesRequest) SetGroupId(v string) *ListModulesRequest {
	s.GroupId = &v
	return s
}

func (s *ListModulesRequest) SetKeyword(v string) *ListModulesRequest {
	s.Keyword = &v
	return s
}

func (s *ListModulesRequest) SetPageNumber(v int32) *ListModulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModulesRequest) SetPageSize(v int32) *ListModulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListModulesRequest) SetProjectId(v string) *ListModulesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListModulesRequest) SetTag(v []*ListModulesRequestTag) *ListModulesRequest {
	s.Tag = v
	return s
}

func (s *ListModulesRequest) Validate() error {
	return dara.Validate(s)
}

type ListModulesRequestTag struct {
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListModulesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListModulesRequestTag) GoString() string {
	return s.String()
}

func (s *ListModulesRequestTag) GetTagKey() *string {
	return s.TagKey
}

func (s *ListModulesRequestTag) GetTagValue() *string {
	return s.TagValue
}

func (s *ListModulesRequestTag) SetTagKey(v string) *ListModulesRequestTag {
	s.TagKey = &v
	return s
}

func (s *ListModulesRequestTag) SetTagValue(v string) *ListModulesRequestTag {
	s.TagValue = &v
	return s
}

func (s *ListModulesRequestTag) Validate() error {
	return dara.Validate(s)
}
