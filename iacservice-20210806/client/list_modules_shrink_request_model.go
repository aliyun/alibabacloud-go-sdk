// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListModulesShrinkRequest
	GetGroupId() *string
	SetKeyword(v string) *ListModulesShrinkRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListModulesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModulesShrinkRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListModulesShrinkRequest
	GetProjectId() *string
	SetTagShrink(v string) *ListModulesShrinkRequest
	GetTagShrink() *string
}

type ListModulesShrinkRequest struct {
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
	PageSize  *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListModulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListModulesShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListModulesShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListModulesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModulesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModulesShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListModulesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListModulesShrinkRequest) SetGroupId(v string) *ListModulesShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *ListModulesShrinkRequest) SetKeyword(v string) *ListModulesShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListModulesShrinkRequest) SetPageNumber(v int32) *ListModulesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModulesShrinkRequest) SetPageSize(v int32) *ListModulesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListModulesShrinkRequest) SetProjectId(v string) *ListModulesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListModulesShrinkRequest) SetTagShrink(v string) *ListModulesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListModulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
