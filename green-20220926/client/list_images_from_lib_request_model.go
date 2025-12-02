// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesFromLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListImagesFromLibRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *ListImagesFromLibRequest
	GetEndDate() *string
	SetImgId(v string) *ListImagesFromLibRequest
	GetImgId() *string
	SetLibId(v string) *ListImagesFromLibRequest
	GetLibId() *string
	SetPageSize(v int32) *ListImagesFromLibRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListImagesFromLibRequest
	GetRegionId() *string
	SetSort(v map[string]*string) *ListImagesFromLibRequest
	GetSort() map[string]*string
	SetStartDate(v string) *ListImagesFromLibRequest
	GetStartDate() *string
}

type ListImagesFromLibRequest struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// End date.
	//
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Image ID.
	//
	// example:
	//
	// 112
	ImgId *string `json:"ImgId,omitempty" xml:"ImgId,omitempty"`
	// Gallery ID.
	//
	// example:
	//
	// custom_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Sort field.
	Sort map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Start date.
	//
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ListImagesFromLibRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImagesFromLibRequest) GoString() string {
	return s.String()
}

func (s *ListImagesFromLibRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListImagesFromLibRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *ListImagesFromLibRequest) GetImgId() *string {
	return s.ImgId
}

func (s *ListImagesFromLibRequest) GetLibId() *string {
	return s.LibId
}

func (s *ListImagesFromLibRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImagesFromLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListImagesFromLibRequest) GetSort() map[string]*string {
	return s.Sort
}

func (s *ListImagesFromLibRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *ListImagesFromLibRequest) SetCurrentPage(v int32) *ListImagesFromLibRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListImagesFromLibRequest) SetEndDate(v string) *ListImagesFromLibRequest {
	s.EndDate = &v
	return s
}

func (s *ListImagesFromLibRequest) SetImgId(v string) *ListImagesFromLibRequest {
	s.ImgId = &v
	return s
}

func (s *ListImagesFromLibRequest) SetLibId(v string) *ListImagesFromLibRequest {
	s.LibId = &v
	return s
}

func (s *ListImagesFromLibRequest) SetPageSize(v int32) *ListImagesFromLibRequest {
	s.PageSize = &v
	return s
}

func (s *ListImagesFromLibRequest) SetRegionId(v string) *ListImagesFromLibRequest {
	s.RegionId = &v
	return s
}

func (s *ListImagesFromLibRequest) SetSort(v map[string]*string) *ListImagesFromLibRequest {
	s.Sort = v
	return s
}

func (s *ListImagesFromLibRequest) SetStartDate(v string) *ListImagesFromLibRequest {
	s.StartDate = &v
	return s
}

func (s *ListImagesFromLibRequest) Validate() error {
	return dara.Validate(s)
}
