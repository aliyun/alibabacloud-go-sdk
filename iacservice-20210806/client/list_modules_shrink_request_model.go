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
	SetModuleName(v string) *ListModulesShrinkRequest
	GetModuleName() *string
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
	// The group ID.
	//
	// example:
	//
	// g-kw1a50tj8rk7cki2q8bbat
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The search keyword. Supports fuzzy match on template names.
	//
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The template name.
	//
	// example:
	//
	// ModuleName
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of results per page. Default value: 20. Minimum value: 1. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The project ID.
	//
	// example:
	//
	// p-al1c58tb2lu9oej36kclvf
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The list of template tags.
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

func (s *ListModulesShrinkRequest) GetModuleName() *string {
	return s.ModuleName
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

func (s *ListModulesShrinkRequest) SetModuleName(v string) *ListModulesShrinkRequest {
	s.ModuleName = &v
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
