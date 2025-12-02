// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOssCheckResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListOssCheckResultShrinkRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *ListOssCheckResultShrinkRequest
	GetEndDate() *string
	SetFinishNum(v int64) *ListOssCheckResultShrinkRequest
	GetFinishNum() *int64
	SetPageSize(v int32) *ListOssCheckResultShrinkRequest
	GetPageSize() *int32
	SetQuery(v string) *ListOssCheckResultShrinkRequest
	GetQuery() *string
	SetRegionId(v string) *ListOssCheckResultShrinkRequest
	GetRegionId() *string
	SetSortShrink(v string) *ListOssCheckResultShrinkRequest
	GetSortShrink() *string
	SetStartDate(v string) *ListOssCheckResultShrinkRequest
	GetStartDate() *string
	SetStatus(v int32) *ListOssCheckResultShrinkRequest
	GetStatus() *int32
}

type ListOssCheckResultShrinkRequest struct {
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
	// Number of completed tasks.
	//
	// example:
	//
	// 2
	FinishNum *int64 `json:"FinishNum,omitempty" xml:"FinishNum,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Search condition.
	//
	// example:
	//
	// {"TaskId":"P_11TL5T"}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Sort field.
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Start date.
	//
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Task status.
	//
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListOssCheckResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOssCheckResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListOssCheckResultShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListOssCheckResultShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *ListOssCheckResultShrinkRequest) GetFinishNum() *int64 {
	return s.FinishNum
}

func (s *ListOssCheckResultShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOssCheckResultShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *ListOssCheckResultShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOssCheckResultShrinkRequest) GetSortShrink() *string {
	return s.SortShrink
}

func (s *ListOssCheckResultShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *ListOssCheckResultShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListOssCheckResultShrinkRequest) SetCurrentPage(v int32) *ListOssCheckResultShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) SetEndDate(v string) *ListOssCheckResultShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) SetFinishNum(v int64) *ListOssCheckResultShrinkRequest {
	s.FinishNum = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) SetPageSize(v int32) *ListOssCheckResultShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) SetQuery(v string) *ListOssCheckResultShrinkRequest {
	s.Query = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) SetRegionId(v string) *ListOssCheckResultShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) SetSortShrink(v string) *ListOssCheckResultShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) SetStartDate(v string) *ListOssCheckResultShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) SetStatus(v int32) *ListOssCheckResultShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListOssCheckResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
