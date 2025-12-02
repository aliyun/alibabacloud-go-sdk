// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOssCheckResultListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *OssCheckResultListRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *OssCheckResultListRequest
	GetEndDate() *string
	SetFinishNum(v int64) *OssCheckResultListRequest
	GetFinishNum() *int64
	SetPageSize(v int32) *OssCheckResultListRequest
	GetPageSize() *int32
	SetQuery(v string) *OssCheckResultListRequest
	GetQuery() *string
	SetRegionId(v string) *OssCheckResultListRequest
	GetRegionId() *string
	SetSort(v map[string]*string) *OssCheckResultListRequest
	GetSort() map[string]*string
	SetStartDate(v string) *OssCheckResultListRequest
	GetStartDate() *string
	SetStatus(v int32) *OssCheckResultListRequest
	GetStatus() *int32
}

type OssCheckResultListRequest struct {
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
	Sort map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
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

func (s OssCheckResultListRequest) String() string {
	return dara.Prettify(s)
}

func (s OssCheckResultListRequest) GoString() string {
	return s.String()
}

func (s *OssCheckResultListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *OssCheckResultListRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *OssCheckResultListRequest) GetFinishNum() *int64 {
	return s.FinishNum
}

func (s *OssCheckResultListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *OssCheckResultListRequest) GetQuery() *string {
	return s.Query
}

func (s *OssCheckResultListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OssCheckResultListRequest) GetSort() map[string]*string {
	return s.Sort
}

func (s *OssCheckResultListRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *OssCheckResultListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *OssCheckResultListRequest) SetCurrentPage(v int32) *OssCheckResultListRequest {
	s.CurrentPage = &v
	return s
}

func (s *OssCheckResultListRequest) SetEndDate(v string) *OssCheckResultListRequest {
	s.EndDate = &v
	return s
}

func (s *OssCheckResultListRequest) SetFinishNum(v int64) *OssCheckResultListRequest {
	s.FinishNum = &v
	return s
}

func (s *OssCheckResultListRequest) SetPageSize(v int32) *OssCheckResultListRequest {
	s.PageSize = &v
	return s
}

func (s *OssCheckResultListRequest) SetQuery(v string) *OssCheckResultListRequest {
	s.Query = &v
	return s
}

func (s *OssCheckResultListRequest) SetRegionId(v string) *OssCheckResultListRequest {
	s.RegionId = &v
	return s
}

func (s *OssCheckResultListRequest) SetSort(v map[string]*string) *OssCheckResultListRequest {
	s.Sort = v
	return s
}

func (s *OssCheckResultListRequest) SetStartDate(v string) *OssCheckResultListRequest {
	s.StartDate = &v
	return s
}

func (s *OssCheckResultListRequest) SetStatus(v int32) *OssCheckResultListRequest {
	s.Status = &v
	return s
}

func (s *OssCheckResultListRequest) Validate() error {
	return dara.Validate(s)
}
