// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListEnvironmentsShrinkRequest
	GetKeyword() *string
	SetLabelSelectorShrink(v string) *ListEnvironmentsShrinkRequest
	GetLabelSelectorShrink() *string
	SetPageNumber(v int64) *ListEnvironmentsShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListEnvironmentsShrinkRequest
	GetPageSize() *int64
}

type ListEnvironmentsShrinkRequest struct {
	// example:
	//
	// dev
	Keyword             *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelectorShrink *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListEnvironmentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListEnvironmentsShrinkRequest) GetLabelSelectorShrink() *string {
	return s.LabelSelectorShrink
}

func (s *ListEnvironmentsShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListEnvironmentsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListEnvironmentsShrinkRequest) SetKeyword(v string) *ListEnvironmentsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListEnvironmentsShrinkRequest) SetLabelSelectorShrink(v string) *ListEnvironmentsShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

func (s *ListEnvironmentsShrinkRequest) SetPageNumber(v int64) *ListEnvironmentsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEnvironmentsShrinkRequest) SetPageSize(v int64) *ListEnvironmentsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
