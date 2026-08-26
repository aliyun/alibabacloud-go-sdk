// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSaseUserTagsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListSaseUserTagsShrinkRequest
	GetCurrentPage() *int64
	SetName(v string) *ListSaseUserTagsShrinkRequest
	GetName() *string
	SetPageSize(v int64) *ListSaseUserTagsShrinkRequest
	GetPageSize() *int64
	SetTagIdsShrink(v string) *ListSaseUserTagsShrinkRequest
	GetTagIdsShrink() *string
}

type ListSaseUserTagsShrinkRequest struct {
	// The page number of the current page in a paging query. Valid values: 1 to 10000.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the user label. The name must be 1 to 128 characters in length.
	//
	// example:
	//
	// boss
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries per page. Settings: 1 to 1000.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The collection of user label IDs.
	TagIdsShrink *string `json:"TagIds,omitempty" xml:"TagIds,omitempty"`
}

func (s ListSaseUserTagsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSaseUserTagsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSaseUserTagsShrinkRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListSaseUserTagsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListSaseUserTagsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSaseUserTagsShrinkRequest) GetTagIdsShrink() *string {
	return s.TagIdsShrink
}

func (s *ListSaseUserTagsShrinkRequest) SetCurrentPage(v int64) *ListSaseUserTagsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListSaseUserTagsShrinkRequest) SetName(v string) *ListSaseUserTagsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListSaseUserTagsShrinkRequest) SetPageSize(v int64) *ListSaseUserTagsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListSaseUserTagsShrinkRequest) SetTagIdsShrink(v string) *ListSaseUserTagsShrinkRequest {
	s.TagIdsShrink = &v
	return s
}

func (s *ListSaseUserTagsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
