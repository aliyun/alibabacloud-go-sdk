// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSaseUserTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListSaseUserTagsRequest
	GetCurrentPage() *int64
	SetName(v string) *ListSaseUserTagsRequest
	GetName() *string
	SetPageSize(v int64) *ListSaseUserTagsRequest
	GetPageSize() *int64
	SetTagIds(v []*string) *ListSaseUserTagsRequest
	GetTagIds() []*string
}

type ListSaseUserTagsRequest struct {
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
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s ListSaseUserTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSaseUserTagsRequest) GoString() string {
	return s.String()
}

func (s *ListSaseUserTagsRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListSaseUserTagsRequest) GetName() *string {
	return s.Name
}

func (s *ListSaseUserTagsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSaseUserTagsRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *ListSaseUserTagsRequest) SetCurrentPage(v int64) *ListSaseUserTagsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListSaseUserTagsRequest) SetName(v string) *ListSaseUserTagsRequest {
	s.Name = &v
	return s
}

func (s *ListSaseUserTagsRequest) SetPageSize(v int64) *ListSaseUserTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSaseUserTagsRequest) SetTagIds(v []*string) *ListSaseUserTagsRequest {
	s.TagIds = v
	return s
}

func (s *ListSaseUserTagsRequest) Validate() error {
	return dara.Validate(s)
}
