// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssCheckFreezeResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetOssCheckFreezeResultRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *GetOssCheckFreezeResultRequest
	GetEndDate() *string
	SetFinishNum(v int64) *GetOssCheckFreezeResultRequest
	GetFinishNum() *int64
	SetPageSize(v int32) *GetOssCheckFreezeResultRequest
	GetPageSize() *int32
	SetQuery(v string) *GetOssCheckFreezeResultRequest
	GetQuery() *string
	SetRegionId(v string) *GetOssCheckFreezeResultRequest
	GetRegionId() *string
	SetSort(v map[string]*string) *GetOssCheckFreezeResultRequest
	GetSort() map[string]*string
	SetStartDate(v string) *GetOssCheckFreezeResultRequest
	GetStartDate() *string
	SetStatus(v int32) *GetOssCheckFreezeResultRequest
	GetStatus() *int32
}

type GetOssCheckFreezeResultRequest struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// End time.
	//
	// example:
	//
	// 2025-05-19 10:05:11
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Number of completed tasks.
	//
	// example:
	//
	// 0
	FinishNum *int64 `json:"FinishNum,omitempty" xml:"FinishNum,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Query condition.
	//
	// example:
	//
	// {\\"TaskId\\":\\"P_O3SI0I\\"}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Sort field.
	Sort map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// Start time.
	//
	// example:
	//
	// 2025-01-09 10:28:54
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Task status.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetOssCheckFreezeResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckFreezeResultRequest) GoString() string {
	return s.String()
}

func (s *GetOssCheckFreezeResultRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetOssCheckFreezeResultRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetOssCheckFreezeResultRequest) GetFinishNum() *int64 {
	return s.FinishNum
}

func (s *GetOssCheckFreezeResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOssCheckFreezeResultRequest) GetQuery() *string {
	return s.Query
}

func (s *GetOssCheckFreezeResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetOssCheckFreezeResultRequest) GetSort() map[string]*string {
	return s.Sort
}

func (s *GetOssCheckFreezeResultRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetOssCheckFreezeResultRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GetOssCheckFreezeResultRequest) SetCurrentPage(v int32) *GetOssCheckFreezeResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetOssCheckFreezeResultRequest) SetEndDate(v string) *GetOssCheckFreezeResultRequest {
	s.EndDate = &v
	return s
}

func (s *GetOssCheckFreezeResultRequest) SetFinishNum(v int64) *GetOssCheckFreezeResultRequest {
	s.FinishNum = &v
	return s
}

func (s *GetOssCheckFreezeResultRequest) SetPageSize(v int32) *GetOssCheckFreezeResultRequest {
	s.PageSize = &v
	return s
}

func (s *GetOssCheckFreezeResultRequest) SetQuery(v string) *GetOssCheckFreezeResultRequest {
	s.Query = &v
	return s
}

func (s *GetOssCheckFreezeResultRequest) SetRegionId(v string) *GetOssCheckFreezeResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetOssCheckFreezeResultRequest) SetSort(v map[string]*string) *GetOssCheckFreezeResultRequest {
	s.Sort = v
	return s
}

func (s *GetOssCheckFreezeResultRequest) SetStartDate(v string) *GetOssCheckFreezeResultRequest {
	s.StartDate = &v
	return s
}

func (s *GetOssCheckFreezeResultRequest) SetStatus(v int32) *GetOssCheckFreezeResultRequest {
	s.Status = &v
	return s
}

func (s *GetOssCheckFreezeResultRequest) Validate() error {
	return dara.Validate(s)
}
