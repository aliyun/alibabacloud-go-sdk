// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrowserBookmarksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListBrowserBookmarksResponseBodyData) *ListBrowserBookmarksResponseBody
	GetData() []*ListBrowserBookmarksResponseBodyData
	SetPageNumber(v int32) *ListBrowserBookmarksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBrowserBookmarksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListBrowserBookmarksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListBrowserBookmarksResponseBody
	GetTotalCount() *int32
}

type ListBrowserBookmarksResponseBody struct {
	Data []*ListBrowserBookmarksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 01A0C2ED-95F2-1A37-9FC6-4A395179****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBrowserBookmarksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserBookmarksResponseBody) GoString() string {
	return s.String()
}

func (s *ListBrowserBookmarksResponseBody) GetData() []*ListBrowserBookmarksResponseBodyData {
	return s.Data
}

func (s *ListBrowserBookmarksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBrowserBookmarksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBrowserBookmarksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBrowserBookmarksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBrowserBookmarksResponseBody) SetData(v []*ListBrowserBookmarksResponseBodyData) *ListBrowserBookmarksResponseBody {
	s.Data = v
	return s
}

func (s *ListBrowserBookmarksResponseBody) SetPageNumber(v int32) *ListBrowserBookmarksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBrowserBookmarksResponseBody) SetPageSize(v int32) *ListBrowserBookmarksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBrowserBookmarksResponseBody) SetRequestId(v string) *ListBrowserBookmarksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBrowserBookmarksResponseBody) SetTotalCount(v int32) *ListBrowserBookmarksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBrowserBookmarksResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBrowserBookmarksResponseBodyData struct {
	// example:
	//
	// 工作
	BookmarkFolder *string `json:"BookmarkFolder,omitempty" xml:"BookmarkFolder,omitempty"`
	// example:
	//
	// bm-0c7loey7fzjq****
	BookmarkId *string `json:"BookmarkId,omitempty" xml:"BookmarkId,omitempty"`
	// example:
	//
	// 阿里云
	BookmarkName *string `json:"BookmarkName,omitempty" xml:"BookmarkName,omitempty"`
	// example:
	//
	// https://www.aliyun.com
	BookmarkURL *string `json:"BookmarkURL,omitempty" xml:"BookmarkURL,omitempty"`
}

func (s ListBrowserBookmarksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserBookmarksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBrowserBookmarksResponseBodyData) GetBookmarkFolder() *string {
	return s.BookmarkFolder
}

func (s *ListBrowserBookmarksResponseBodyData) GetBookmarkId() *string {
	return s.BookmarkId
}

func (s *ListBrowserBookmarksResponseBodyData) GetBookmarkName() *string {
	return s.BookmarkName
}

func (s *ListBrowserBookmarksResponseBodyData) GetBookmarkURL() *string {
	return s.BookmarkURL
}

func (s *ListBrowserBookmarksResponseBodyData) SetBookmarkFolder(v string) *ListBrowserBookmarksResponseBodyData {
	s.BookmarkFolder = &v
	return s
}

func (s *ListBrowserBookmarksResponseBodyData) SetBookmarkId(v string) *ListBrowserBookmarksResponseBodyData {
	s.BookmarkId = &v
	return s
}

func (s *ListBrowserBookmarksResponseBodyData) SetBookmarkName(v string) *ListBrowserBookmarksResponseBodyData {
	s.BookmarkName = &v
	return s
}

func (s *ListBrowserBookmarksResponseBodyData) SetBookmarkURL(v string) *ListBrowserBookmarksResponseBodyData {
	s.BookmarkURL = &v
	return s
}

func (s *ListBrowserBookmarksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
