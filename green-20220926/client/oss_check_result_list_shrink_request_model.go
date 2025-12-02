// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOssCheckResultListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *OssCheckResultListShrinkRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *OssCheckResultListShrinkRequest
	GetEndDate() *string
	SetFinishNum(v int64) *OssCheckResultListShrinkRequest
	GetFinishNum() *int64
	SetPageSize(v int32) *OssCheckResultListShrinkRequest
	GetPageSize() *int32
	SetQuery(v string) *OssCheckResultListShrinkRequest
	GetQuery() *string
	SetRegionId(v string) *OssCheckResultListShrinkRequest
	GetRegionId() *string
	SetSortShrink(v string) *OssCheckResultListShrinkRequest
	GetSortShrink() *string
	SetStartDate(v string) *OssCheckResultListShrinkRequest
	GetStartDate() *string
	SetStatus(v int32) *OssCheckResultListShrinkRequest
	GetStatus() *int32
}

type OssCheckResultListShrinkRequest struct {
	// Page size.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Start date.
	//
	// example:
	//
	// 2023-10-21 16:08:38
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Region ID.
	//
	// example:
	//
	// 55
	FinishNum *int64 `json:"FinishNum,omitempty" xml:"FinishNum,omitempty"`
	// Query condition.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// End date.
	//
	// example:
	//
	// {}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// Sort field.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Current page number.
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// System-defined parameter. Value: **OssCheckResultList**.
	//
	// example:
	//
	// 2023-08-21 16:08:38
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Number of completed items.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s OssCheckResultListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OssCheckResultListShrinkRequest) GoString() string {
	return s.String()
}

func (s *OssCheckResultListShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *OssCheckResultListShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *OssCheckResultListShrinkRequest) GetFinishNum() *int64 {
	return s.FinishNum
}

func (s *OssCheckResultListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *OssCheckResultListShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *OssCheckResultListShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OssCheckResultListShrinkRequest) GetSortShrink() *string {
	return s.SortShrink
}

func (s *OssCheckResultListShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *OssCheckResultListShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *OssCheckResultListShrinkRequest) SetCurrentPage(v int32) *OssCheckResultListShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *OssCheckResultListShrinkRequest) SetEndDate(v string) *OssCheckResultListShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *OssCheckResultListShrinkRequest) SetFinishNum(v int64) *OssCheckResultListShrinkRequest {
	s.FinishNum = &v
	return s
}

func (s *OssCheckResultListShrinkRequest) SetPageSize(v int32) *OssCheckResultListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *OssCheckResultListShrinkRequest) SetQuery(v string) *OssCheckResultListShrinkRequest {
	s.Query = &v
	return s
}

func (s *OssCheckResultListShrinkRequest) SetRegionId(v string) *OssCheckResultListShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *OssCheckResultListShrinkRequest) SetSortShrink(v string) *OssCheckResultListShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *OssCheckResultListShrinkRequest) SetStartDate(v string) *OssCheckResultListShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *OssCheckResultListShrinkRequest) SetStatus(v int32) *OssCheckResultListShrinkRequest {
	s.Status = &v
	return s
}

func (s *OssCheckResultListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
