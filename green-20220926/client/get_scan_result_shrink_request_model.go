// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScanResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetScanResultShrinkRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *GetScanResultShrinkRequest
	GetEndDate() *string
	SetPageSize(v int32) *GetScanResultShrinkRequest
	GetPageSize() *int32
	SetQueryShrink(v string) *GetScanResultShrinkRequest
	GetQueryShrink() *string
	SetRegionId(v string) *GetScanResultShrinkRequest
	GetRegionId() *string
	SetResourceType(v string) *GetScanResultShrinkRequest
	GetResourceType() *string
	SetSortShrink(v string) *GetScanResultShrinkRequest
	GetSortShrink() *string
	SetStartDate(v string) *GetScanResultShrinkRequest
	GetStartDate() *string
}

type GetScanResultShrinkRequest struct {
	// Current page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// End time.
	//
	// example:
	//
	// 2023-08-24 10:01:55
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Search criteria.
	QueryShrink *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource type.
	//
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Sort fields.
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Start time.
	//
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetScanResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScanResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetScanResultShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetScanResultShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetScanResultShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetScanResultShrinkRequest) GetQueryShrink() *string {
	return s.QueryShrink
}

func (s *GetScanResultShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetScanResultShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetScanResultShrinkRequest) GetSortShrink() *string {
	return s.SortShrink
}

func (s *GetScanResultShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetScanResultShrinkRequest) SetCurrentPage(v int32) *GetScanResultShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetScanResultShrinkRequest) SetEndDate(v string) *GetScanResultShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetScanResultShrinkRequest) SetPageSize(v int32) *GetScanResultShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetScanResultShrinkRequest) SetQueryShrink(v string) *GetScanResultShrinkRequest {
	s.QueryShrink = &v
	return s
}

func (s *GetScanResultShrinkRequest) SetRegionId(v string) *GetScanResultShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetScanResultShrinkRequest) SetResourceType(v string) *GetScanResultShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *GetScanResultShrinkRequest) SetSortShrink(v string) *GetScanResultShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *GetScanResultShrinkRequest) SetStartDate(v string) *GetScanResultShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetScanResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
