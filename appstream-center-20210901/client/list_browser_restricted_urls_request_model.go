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
	// The ID of the browser delivery group. You must specify either this parameter or BrowserInstanceGroupSetId, but not both. If you specify this parameter, the restricted URLs configured for the delivery group are queried.
	//
	// example:
	//
	// big-3jm9d0abc00example
	BrowserInstanceGroupId *string `json:"BrowserInstanceGroupId,omitempty" xml:"BrowserInstanceGroupId,omitempty"`
	// The ID of the browser delivery group set. You must specify either BrowserInstanceGroupSetId or BrowserInstanceGroupId, but not both.
	//
	// example:
	//
	// set-3jm9d0abc00example
	BrowserInstanceGroupSetId *string `json:"BrowserInstanceGroupSetId,omitempty" xml:"BrowserInstanceGroupSetId,omitempty"`
	// The keyword used to filter access control URLs by using an include match. The match is case-insensitive for English letters. The maximum length is 1,024 characters after leading and trailing spaces are removed. If this parameter is not specified or contains only spaces, no filtering is performed. Special characters are matched literally.
	//
	// example:
	//
	// example.com
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number during a paging query. Default value: 0. If the value is less than 1, such as 0, the first page is returned. If the value is less than 0, a parameter error is returned. Specify 1 or a greater value.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records to return on each page during a paging query. Valid values: 1 to 200. Default value: 200. If the value is less than or equal to 0 or greater than 200, a parameter error is returned.
	//
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
