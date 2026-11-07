// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrowserBookmarksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrowserInstanceGroupId(v string) *ListBrowserBookmarksRequest
	GetBrowserInstanceGroupId() *string
	SetBrowserInstanceGroupSetId(v string) *ListBrowserBookmarksRequest
	GetBrowserInstanceGroupSetId() *string
	SetKeyword(v string) *ListBrowserBookmarksRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListBrowserBookmarksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBrowserBookmarksRequest
	GetPageSize() *int32
}

type ListBrowserBookmarksRequest struct {
	// example:
	//
	// big-0c7loey7fzjq****
	BrowserInstanceGroupId *string `json:"BrowserInstanceGroupId,omitempty" xml:"BrowserInstanceGroupId,omitempty"`
	// example:
	//
	// set-3jm9d0abc00example
	BrowserInstanceGroupSetId *string `json:"BrowserInstanceGroupSetId,omitempty" xml:"BrowserInstanceGroupSetId,omitempty"`
	// example:
	//
	// example.com
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListBrowserBookmarksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserBookmarksRequest) GoString() string {
	return s.String()
}

func (s *ListBrowserBookmarksRequest) GetBrowserInstanceGroupId() *string {
	return s.BrowserInstanceGroupId
}

func (s *ListBrowserBookmarksRequest) GetBrowserInstanceGroupSetId() *string {
	return s.BrowserInstanceGroupSetId
}

func (s *ListBrowserBookmarksRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListBrowserBookmarksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBrowserBookmarksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBrowserBookmarksRequest) SetBrowserInstanceGroupId(v string) *ListBrowserBookmarksRequest {
	s.BrowserInstanceGroupId = &v
	return s
}

func (s *ListBrowserBookmarksRequest) SetBrowserInstanceGroupSetId(v string) *ListBrowserBookmarksRequest {
	s.BrowserInstanceGroupSetId = &v
	return s
}

func (s *ListBrowserBookmarksRequest) SetKeyword(v string) *ListBrowserBookmarksRequest {
	s.Keyword = &v
	return s
}

func (s *ListBrowserBookmarksRequest) SetPageNumber(v int32) *ListBrowserBookmarksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBrowserBookmarksRequest) SetPageSize(v int32) *ListBrowserBookmarksRequest {
	s.PageSize = &v
	return s
}

func (s *ListBrowserBookmarksRequest) Validate() error {
	return dara.Validate(s)
}
