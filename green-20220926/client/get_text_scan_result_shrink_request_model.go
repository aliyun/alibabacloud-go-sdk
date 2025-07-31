// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextScanResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetTextScanResultShrinkRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *GetTextScanResultShrinkRequest
	GetEndDate() *string
	SetPageSize(v int32) *GetTextScanResultShrinkRequest
	GetPageSize() *int32
	SetQueryShrink(v string) *GetTextScanResultShrinkRequest
	GetQueryShrink() *string
	SetRegionId(v string) *GetTextScanResultShrinkRequest
	GetRegionId() *string
	SetSortShrink(v string) *GetTextScanResultShrinkRequest
	GetSortShrink() *string
	SetStartDate(v string) *GetTextScanResultShrinkRequest
	GetStartDate() *string
}

type GetTextScanResultShrinkRequest struct {
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
	// 10
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryShrink *string `json:"Query,omitempty" xml:"Query,omitempty"`
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

func (s GetTextScanResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTextScanResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetTextScanResultShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetTextScanResultShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetTextScanResultShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetTextScanResultShrinkRequest) GetQueryShrink() *string {
	return s.QueryShrink
}

func (s *GetTextScanResultShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTextScanResultShrinkRequest) GetSortShrink() *string {
	return s.SortShrink
}

func (s *GetTextScanResultShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetTextScanResultShrinkRequest) SetCurrentPage(v int32) *GetTextScanResultShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetTextScanResultShrinkRequest) SetEndDate(v string) *GetTextScanResultShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetTextScanResultShrinkRequest) SetPageSize(v int32) *GetTextScanResultShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetTextScanResultShrinkRequest) SetQueryShrink(v string) *GetTextScanResultShrinkRequest {
	s.QueryShrink = &v
	return s
}

func (s *GetTextScanResultShrinkRequest) SetRegionId(v string) *GetTextScanResultShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetTextScanResultShrinkRequest) SetSortShrink(v string) *GetTextScanResultShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *GetTextScanResultShrinkRequest) SetStartDate(v string) *GetTextScanResultShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetTextScanResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
