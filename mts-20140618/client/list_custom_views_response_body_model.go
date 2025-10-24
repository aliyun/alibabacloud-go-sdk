// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomViewsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomViews(v *ListCustomViewsResponseBodyCustomViews) *ListCustomViewsResponseBody
	GetCustomViews() *ListCustomViewsResponseBodyCustomViews
	SetPageNumber(v int32) *ListCustomViewsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomViewsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCustomViewsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListCustomViewsResponseBody
	GetTotalCount() *int64
}

type ListCustomViewsResponseBody struct {
	CustomViews *ListCustomViewsResponseBodyCustomViews `json:"CustomViews,omitempty" xml:"CustomViews,omitempty" type:"Struct"`
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
	// 580e8ce3-3b80-44c5-9f3f-36ac3cc5bdd5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomViewsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomViewsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomViewsResponseBody) GetCustomViews() *ListCustomViewsResponseBodyCustomViews {
	return s.CustomViews
}

func (s *ListCustomViewsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomViewsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomViewsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomViewsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCustomViewsResponseBody) SetCustomViews(v *ListCustomViewsResponseBodyCustomViews) *ListCustomViewsResponseBody {
	s.CustomViews = v
	return s
}

func (s *ListCustomViewsResponseBody) SetPageNumber(v int32) *ListCustomViewsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomViewsResponseBody) SetPageSize(v int32) *ListCustomViewsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomViewsResponseBody) SetRequestId(v string) *ListCustomViewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomViewsResponseBody) SetTotalCount(v int64) *ListCustomViewsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomViewsResponseBody) Validate() error {
	if s.CustomViews != nil {
		if err := s.CustomViews.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomViewsResponseBodyCustomViews struct {
	CustomView []*ListCustomViewsResponseBodyCustomViewsCustomView `json:"CustomView,omitempty" xml:"CustomView,omitempty" type:"Repeated"`
}

func (s ListCustomViewsResponseBodyCustomViews) String() string {
	return dara.Prettify(s)
}

func (s ListCustomViewsResponseBodyCustomViews) GoString() string {
	return s.String()
}

func (s *ListCustomViewsResponseBodyCustomViews) GetCustomView() []*ListCustomViewsResponseBodyCustomViewsCustomView {
	return s.CustomView
}

func (s *ListCustomViewsResponseBodyCustomViews) SetCustomView(v []*ListCustomViewsResponseBodyCustomViewsCustomView) *ListCustomViewsResponseBodyCustomViews {
	s.CustomView = v
	return s
}

func (s *ListCustomViewsResponseBodyCustomViews) Validate() error {
	if s.CustomView != nil {
		for _, item := range s.CustomView {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomViewsResponseBodyCustomViewsCustomView struct {
	// example:
	//
	// 1
	CustomViewId *string `json:"CustomViewId,omitempty" xml:"CustomViewId,omitempty"`
	// example:
	//
	// http://``127.66.**.**``/photo.jpeg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s ListCustomViewsResponseBodyCustomViewsCustomView) String() string {
	return dara.Prettify(s)
}

func (s ListCustomViewsResponseBodyCustomViewsCustomView) GoString() string {
	return s.String()
}

func (s *ListCustomViewsResponseBodyCustomViewsCustomView) GetCustomViewId() *string {
	return s.CustomViewId
}

func (s *ListCustomViewsResponseBodyCustomViewsCustomView) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ListCustomViewsResponseBodyCustomViewsCustomView) SetCustomViewId(v string) *ListCustomViewsResponseBodyCustomViewsCustomView {
	s.CustomViewId = &v
	return s
}

func (s *ListCustomViewsResponseBodyCustomViewsCustomView) SetImageUrl(v string) *ListCustomViewsResponseBodyCustomViewsCustomView {
	s.ImageUrl = &v
	return s
}

func (s *ListCustomViewsResponseBodyCustomViewsCustomView) Validate() error {
	return dara.Validate(s)
}
