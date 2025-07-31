// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOssCheckResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListOssCheckResultRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *ListOssCheckResultRequest
	GetEndDate() *string
	SetFinishNum(v int64) *ListOssCheckResultRequest
	GetFinishNum() *int64
	SetPageSize(v int32) *ListOssCheckResultRequest
	GetPageSize() *int32
	SetQuery(v string) *ListOssCheckResultRequest
	GetQuery() *string
	SetRegionId(v string) *ListOssCheckResultRequest
	GetRegionId() *string
	SetSort(v map[string]*string) *ListOssCheckResultRequest
	GetSort() map[string]*string
	SetStartDate(v string) *ListOssCheckResultRequest
	GetStartDate() *string
	SetStatus(v int32) *ListOssCheckResultRequest
	GetStatus() *int32
}

type ListOssCheckResultRequest struct {
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
	// 2
	FinishNum *int64 `json:"FinishNum,omitempty" xml:"FinishNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// {"TaskId":"P_11TL5T"}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sort     map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListOssCheckResultRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOssCheckResultRequest) GoString() string {
	return s.String()
}

func (s *ListOssCheckResultRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListOssCheckResultRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *ListOssCheckResultRequest) GetFinishNum() *int64 {
	return s.FinishNum
}

func (s *ListOssCheckResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOssCheckResultRequest) GetQuery() *string {
	return s.Query
}

func (s *ListOssCheckResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOssCheckResultRequest) GetSort() map[string]*string {
	return s.Sort
}

func (s *ListOssCheckResultRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *ListOssCheckResultRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListOssCheckResultRequest) SetCurrentPage(v int32) *ListOssCheckResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOssCheckResultRequest) SetEndDate(v string) *ListOssCheckResultRequest {
	s.EndDate = &v
	return s
}

func (s *ListOssCheckResultRequest) SetFinishNum(v int64) *ListOssCheckResultRequest {
	s.FinishNum = &v
	return s
}

func (s *ListOssCheckResultRequest) SetPageSize(v int32) *ListOssCheckResultRequest {
	s.PageSize = &v
	return s
}

func (s *ListOssCheckResultRequest) SetQuery(v string) *ListOssCheckResultRequest {
	s.Query = &v
	return s
}

func (s *ListOssCheckResultRequest) SetRegionId(v string) *ListOssCheckResultRequest {
	s.RegionId = &v
	return s
}

func (s *ListOssCheckResultRequest) SetSort(v map[string]*string) *ListOssCheckResultRequest {
	s.Sort = v
	return s
}

func (s *ListOssCheckResultRequest) SetStartDate(v string) *ListOssCheckResultRequest {
	s.StartDate = &v
	return s
}

func (s *ListOssCheckResultRequest) SetStatus(v int32) *ListOssCheckResultRequest {
	s.Status = &v
	return s
}

func (s *ListOssCheckResultRequest) Validate() error {
	return dara.Validate(s)
}
