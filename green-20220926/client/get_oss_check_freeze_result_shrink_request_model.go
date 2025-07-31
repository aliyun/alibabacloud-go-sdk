// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssCheckFreezeResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetOssCheckFreezeResultShrinkRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *GetOssCheckFreezeResultShrinkRequest
	GetEndDate() *string
	SetFinishNum(v int64) *GetOssCheckFreezeResultShrinkRequest
	GetFinishNum() *int64
	SetPageSize(v int32) *GetOssCheckFreezeResultShrinkRequest
	GetPageSize() *int32
	SetQuery(v string) *GetOssCheckFreezeResultShrinkRequest
	GetQuery() *string
	SetRegionId(v string) *GetOssCheckFreezeResultShrinkRequest
	GetRegionId() *string
	SetSortShrink(v string) *GetOssCheckFreezeResultShrinkRequest
	GetSortShrink() *string
	SetStartDate(v string) *GetOssCheckFreezeResultShrinkRequest
	GetStartDate() *string
	SetStatus(v int32) *GetOssCheckFreezeResultShrinkRequest
	GetStatus() *int32
}

type GetOssCheckFreezeResultShrinkRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2025-05-19 10:05:11
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 0
	FinishNum *int64 `json:"FinishNum,omitempty" xml:"FinishNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// {\\"TaskId\\":\\"P_O3SI0I\\"}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2025-01-09 10:28:54
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetOssCheckFreezeResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckFreezeResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetOssCheckFreezeResultShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetOssCheckFreezeResultShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetOssCheckFreezeResultShrinkRequest) GetFinishNum() *int64 {
	return s.FinishNum
}

func (s *GetOssCheckFreezeResultShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOssCheckFreezeResultShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *GetOssCheckFreezeResultShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetOssCheckFreezeResultShrinkRequest) GetSortShrink() *string {
	return s.SortShrink
}

func (s *GetOssCheckFreezeResultShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetOssCheckFreezeResultShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GetOssCheckFreezeResultShrinkRequest) SetCurrentPage(v int32) *GetOssCheckFreezeResultShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetOssCheckFreezeResultShrinkRequest) SetEndDate(v string) *GetOssCheckFreezeResultShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *GetOssCheckFreezeResultShrinkRequest) SetFinishNum(v int64) *GetOssCheckFreezeResultShrinkRequest {
	s.FinishNum = &v
	return s
}

func (s *GetOssCheckFreezeResultShrinkRequest) SetPageSize(v int32) *GetOssCheckFreezeResultShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetOssCheckFreezeResultShrinkRequest) SetQuery(v string) *GetOssCheckFreezeResultShrinkRequest {
	s.Query = &v
	return s
}

func (s *GetOssCheckFreezeResultShrinkRequest) SetRegionId(v string) *GetOssCheckFreezeResultShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetOssCheckFreezeResultShrinkRequest) SetSortShrink(v string) *GetOssCheckFreezeResultShrinkRequest {
	s.SortShrink = &v
	return s
}

func (s *GetOssCheckFreezeResultShrinkRequest) SetStartDate(v string) *GetOssCheckFreezeResultShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *GetOssCheckFreezeResultShrinkRequest) SetStatus(v int32) *GetOssCheckFreezeResultShrinkRequest {
	s.Status = &v
	return s
}

func (s *GetOssCheckFreezeResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
