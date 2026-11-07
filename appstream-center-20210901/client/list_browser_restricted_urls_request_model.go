// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrowserRestrictedURLsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrowserInstanceGroupId(v string) *ListBrowserRestrictedURLsRequest
	GetBrowserInstanceGroupId() *string
	SetBrowserInstanceGroupSetId(v string) *ListBrowserRestrictedURLsRequest
	GetBrowserInstanceGroupSetId() *string
	SetKeyword(v string) *ListBrowserRestrictedURLsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListBrowserRestrictedURLsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBrowserRestrictedURLsRequest
	GetPageSize() *int32
}

type ListBrowserRestrictedURLsRequest struct {
	// example:
	//
	// big-3jm9d0abc00example
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListBrowserRestrictedURLsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserRestrictedURLsRequest) GoString() string {
	return s.String()
}

func (s *ListBrowserRestrictedURLsRequest) GetBrowserInstanceGroupId() *string {
	return s.BrowserInstanceGroupId
}

func (s *ListBrowserRestrictedURLsRequest) GetBrowserInstanceGroupSetId() *string {
	return s.BrowserInstanceGroupSetId
}

func (s *ListBrowserRestrictedURLsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListBrowserRestrictedURLsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBrowserRestrictedURLsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBrowserRestrictedURLsRequest) SetBrowserInstanceGroupId(v string) *ListBrowserRestrictedURLsRequest {
	s.BrowserInstanceGroupId = &v
	return s
}

func (s *ListBrowserRestrictedURLsRequest) SetBrowserInstanceGroupSetId(v string) *ListBrowserRestrictedURLsRequest {
	s.BrowserInstanceGroupSetId = &v
	return s
}

func (s *ListBrowserRestrictedURLsRequest) SetKeyword(v string) *ListBrowserRestrictedURLsRequest {
	s.Keyword = &v
	return s
}

func (s *ListBrowserRestrictedURLsRequest) SetPageNumber(v int32) *ListBrowserRestrictedURLsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBrowserRestrictedURLsRequest) SetPageSize(v int32) *ListBrowserRestrictedURLsRequest {
	s.PageSize = &v
	return s
}

func (s *ListBrowserRestrictedURLsRequest) Validate() error {
	return dara.Validate(s)
}
