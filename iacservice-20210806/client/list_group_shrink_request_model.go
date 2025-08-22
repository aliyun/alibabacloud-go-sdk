// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListGroupShrinkRequest
	GetKeyword() *string
	SetPageNumber(v string) *ListGroupShrinkRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListGroupShrinkRequest
	GetPageSize() *string
	SetProjectId(v string) *ListGroupShrinkRequest
	GetProjectId() *string
	SetTagShrink(v string) *ListGroupShrinkRequest
	GetTagShrink() *string
}

type ListGroupShrinkRequest struct {
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 200
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// p-14e80de4866bf7ffed0c4072ed9b37
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGroupShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListGroupShrinkRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListGroupShrinkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListGroupShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListGroupShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListGroupShrinkRequest) SetKeyword(v string) *ListGroupShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListGroupShrinkRequest) SetPageNumber(v string) *ListGroupShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGroupShrinkRequest) SetPageSize(v string) *ListGroupShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupShrinkRequest) SetProjectId(v string) *ListGroupShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListGroupShrinkRequest) SetTagShrink(v string) *ListGroupShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
