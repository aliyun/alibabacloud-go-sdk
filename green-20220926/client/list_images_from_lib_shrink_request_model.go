// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesFromLibShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListImagesFromLibShrinkRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *ListImagesFromLibShrinkRequest
	GetEndDate() *string
	SetImgId(v string) *ListImagesFromLibShrinkRequest
	GetImgId() *string
	SetLibId(v string) *ListImagesFromLibShrinkRequest
	GetLibId() *string
	SetPageSize(v int32) *ListImagesFromLibShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListImagesFromLibShrinkRequest
	GetRegionId() *string
	SetSortShrink(v string) *ListImagesFromLibShrinkRequest
	GetSortShrink() *string
	SetStartDate(v string) *ListImagesFromLibShrinkRequest
	GetStartDate() *string
}

type ListImagesFromLibShrinkRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 112
	ImgId *string `json:"ImgId,omitempty" xml:"ImgId,omitempty"`
	// example:
	//
	// custom_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ListImagesFromLibShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImagesFromLibShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListImagesFromLibShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListImagesFromLibShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *ListImagesFromLibShrinkRequest) GetImgId() *string {
	return s.ImgId
}

func (s *ListImagesFromLibShrinkRequest) GetLibId() *string {
	return s.LibId
}

func (s *ListImagesFromLibShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImagesFromLibShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListImagesFromLibShrinkRequest) GetSortShrink() *string {
	return s.SortShrink
}

func (s *ListImagesFromLibShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *ListImagesFromLibShrinkRequest) SetCurrentPage(v int32) *ListImagesFromLibShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListImagesFromLibShrinkRequest) SetEndDate(v string) *ListImagesFromLibShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *ListImagesFromLibShrinkRequest) SetImgId(v string) *ListImagesFromLibShrinkRequest {
	s.ImgId = &v
	return s
}

func (s *ListImagesFromLibShrinkRequest) SetLibId(v string) *ListImagesFromLibShrinkRequest {
	s.LibId = &v
	return s
}

func (s *ListImagesFromLibShrinkRequest) SetPageSize(v int32) *ListImagesFromLibShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListImagesFromLibShrinkRequest) SetRegionId(v string) *ListImagesFromLibShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListImagesFromLibShrinkRequest) SetSortShrink(v string) *ListImagesFromLibShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *ListImagesFromLibShrinkRequest) SetStartDate(v string) *ListImagesFromLibShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *ListImagesFromLibShrinkRequest) Validate() error {
	return dara.Validate(s)
}
