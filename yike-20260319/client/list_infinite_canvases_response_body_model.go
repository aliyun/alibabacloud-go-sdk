// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInfiniteCanvasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanvasList(v []*ListInfiniteCanvasesResponseBodyCanvasList) *ListInfiniteCanvasesResponseBody
	GetCanvasList() []*ListInfiniteCanvasesResponseBodyCanvasList
	SetPageNo(v int32) *ListInfiniteCanvasesResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListInfiniteCanvasesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListInfiniteCanvasesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListInfiniteCanvasesResponseBody
	GetTotalCount() *int32
}

type ListInfiniteCanvasesResponseBody struct {
	CanvasList []*ListInfiniteCanvasesResponseBodyCanvasList `json:"CanvasList,omitempty" xml:"CanvasList,omitempty" type:"Repeated"`
	// example:
	//
	// 16
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInfiniteCanvasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInfiniteCanvasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInfiniteCanvasesResponseBody) GetCanvasList() []*ListInfiniteCanvasesResponseBodyCanvasList {
	return s.CanvasList
}

func (s *ListInfiniteCanvasesResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListInfiniteCanvasesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInfiniteCanvasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInfiniteCanvasesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInfiniteCanvasesResponseBody) SetCanvasList(v []*ListInfiniteCanvasesResponseBodyCanvasList) *ListInfiniteCanvasesResponseBody {
	s.CanvasList = v
	return s
}

func (s *ListInfiniteCanvasesResponseBody) SetPageNo(v int32) *ListInfiniteCanvasesResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListInfiniteCanvasesResponseBody) SetPageSize(v int32) *ListInfiniteCanvasesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInfiniteCanvasesResponseBody) SetRequestId(v string) *ListInfiniteCanvasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInfiniteCanvasesResponseBody) SetTotalCount(v int32) *ListInfiniteCanvasesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInfiniteCanvasesResponseBody) Validate() error {
	if s.CanvasList != nil {
		for _, item := range s.CanvasList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInfiniteCanvasesResponseBodyCanvasList struct {
	// example:
	//
	// canvas_xxx
	CanvasId *string `json:"CanvasId,omitempty" xml:"CanvasId,omitempty"`
	// example:
	//
	// http://example.com/cover.png
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// example:
	//
	// 2025-12-26T10:21:17Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2026-03-18T10:03:56Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// http://example.com/thumbnail2.png
	Thumbnail *string `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	// example:
	//
	// example
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListInfiniteCanvasesResponseBodyCanvasList) String() string {
	return dara.Prettify(s)
}

func (s ListInfiniteCanvasesResponseBodyCanvasList) GoString() string {
	return s.String()
}

func (s *ListInfiniteCanvasesResponseBodyCanvasList) GetCanvasId() *string {
	return s.CanvasId
}

func (s *ListInfiniteCanvasesResponseBodyCanvasList) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ListInfiniteCanvasesResponseBodyCanvasList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListInfiniteCanvasesResponseBodyCanvasList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListInfiniteCanvasesResponseBodyCanvasList) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *ListInfiniteCanvasesResponseBodyCanvasList) GetTitle() *string {
	return s.Title
}

func (s *ListInfiniteCanvasesResponseBodyCanvasList) SetCanvasId(v string) *ListInfiniteCanvasesResponseBodyCanvasList {
	s.CanvasId = &v
	return s
}

func (s *ListInfiniteCanvasesResponseBodyCanvasList) SetCoverUrl(v string) *ListInfiniteCanvasesResponseBodyCanvasList {
	s.CoverUrl = &v
	return s
}

func (s *ListInfiniteCanvasesResponseBodyCanvasList) SetGmtCreate(v string) *ListInfiniteCanvasesResponseBodyCanvasList {
	s.GmtCreate = &v
	return s
}

func (s *ListInfiniteCanvasesResponseBodyCanvasList) SetGmtModified(v string) *ListInfiniteCanvasesResponseBodyCanvasList {
	s.GmtModified = &v
	return s
}

func (s *ListInfiniteCanvasesResponseBodyCanvasList) SetThumbnail(v string) *ListInfiniteCanvasesResponseBodyCanvasList {
	s.Thumbnail = &v
	return s
}

func (s *ListInfiniteCanvasesResponseBodyCanvasList) SetTitle(v string) *ListInfiniteCanvasesResponseBodyCanvasList {
	s.Title = &v
	return s
}

func (s *ListInfiniteCanvasesResponseBodyCanvasList) Validate() error {
	return dara.Validate(s)
}
