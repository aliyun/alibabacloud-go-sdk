// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrowserRestrictedURLsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListBrowserRestrictedURLsResponseBodyData) *ListBrowserRestrictedURLsResponseBody
	GetData() []*ListBrowserRestrictedURLsResponseBodyData
	SetPageNumber(v int32) *ListBrowserRestrictedURLsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBrowserRestrictedURLsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListBrowserRestrictedURLsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListBrowserRestrictedURLsResponseBody
	GetTotalCount() *int32
}

type ListBrowserRestrictedURLsResponseBody struct {
	Data []*ListBrowserRestrictedURLsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// D4A13D91-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBrowserRestrictedURLsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserRestrictedURLsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBrowserRestrictedURLsResponseBody) GetData() []*ListBrowserRestrictedURLsResponseBodyData {
	return s.Data
}

func (s *ListBrowserRestrictedURLsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBrowserRestrictedURLsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBrowserRestrictedURLsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBrowserRestrictedURLsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBrowserRestrictedURLsResponseBody) SetData(v []*ListBrowserRestrictedURLsResponseBodyData) *ListBrowserRestrictedURLsResponseBody {
	s.Data = v
	return s
}

func (s *ListBrowserRestrictedURLsResponseBody) SetPageNumber(v int32) *ListBrowserRestrictedURLsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBrowserRestrictedURLsResponseBody) SetPageSize(v int32) *ListBrowserRestrictedURLsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBrowserRestrictedURLsResponseBody) SetRequestId(v string) *ListBrowserRestrictedURLsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBrowserRestrictedURLsResponseBody) SetTotalCount(v int32) *ListBrowserRestrictedURLsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBrowserRestrictedURLsResponseBody) Validate() error {
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

type ListBrowserRestrictedURLsResponseBodyData struct {
	// example:
	//
	// ru-3jm9d0abc00example
	RestrictedURLId *string `json:"RestrictedURLId,omitempty" xml:"RestrictedURLId,omitempty"`
	// example:
	//
	// https://www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListBrowserRestrictedURLsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserRestrictedURLsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBrowserRestrictedURLsResponseBodyData) GetRestrictedURLId() *string {
	return s.RestrictedURLId
}

func (s *ListBrowserRestrictedURLsResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *ListBrowserRestrictedURLsResponseBodyData) SetRestrictedURLId(v string) *ListBrowserRestrictedURLsResponseBodyData {
	s.RestrictedURLId = &v
	return s
}

func (s *ListBrowserRestrictedURLsResponseBodyData) SetUrl(v string) *ListBrowserRestrictedURLsResponseBodyData {
	s.Url = &v
	return s
}

func (s *ListBrowserRestrictedURLsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
