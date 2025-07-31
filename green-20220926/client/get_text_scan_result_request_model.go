// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextScanResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetTextScanResultRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *GetTextScanResultRequest
	GetEndDate() *string
	SetPageSize(v int32) *GetTextScanResultRequest
	GetPageSize() *int32
	SetQuery(v map[string]*string) *GetTextScanResultRequest
	GetQuery() map[string]*string
	SetRegionId(v string) *GetTextScanResultRequest
	GetRegionId() *string
	SetSort(v map[string]*string) *GetTextScanResultRequest
	GetSort() map[string]*string
	SetStartDate(v string) *GetTextScanResultRequest
	GetStartDate() *string
}

type GetTextScanResultRequest struct {
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
	PageSize *int32             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query    map[string]*string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sort     map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetTextScanResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTextScanResultRequest) GoString() string {
	return s.String()
}

func (s *GetTextScanResultRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetTextScanResultRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetTextScanResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetTextScanResultRequest) GetQuery() map[string]*string {
	return s.Query
}

func (s *GetTextScanResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTextScanResultRequest) GetSort() map[string]*string {
	return s.Sort
}

func (s *GetTextScanResultRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetTextScanResultRequest) SetCurrentPage(v int32) *GetTextScanResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetTextScanResultRequest) SetEndDate(v string) *GetTextScanResultRequest {
	s.EndDate = &v
	return s
}

func (s *GetTextScanResultRequest) SetPageSize(v int32) *GetTextScanResultRequest {
	s.PageSize = &v
	return s
}

func (s *GetTextScanResultRequest) SetQuery(v map[string]*string) *GetTextScanResultRequest {
	s.Query = v
	return s
}

func (s *GetTextScanResultRequest) SetRegionId(v string) *GetTextScanResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetTextScanResultRequest) SetSort(v map[string]*string) *GetTextScanResultRequest {
	s.Sort = v
	return s
}

func (s *GetTextScanResultRequest) SetStartDate(v string) *GetTextScanResultRequest {
	s.StartDate = &v
	return s
}

func (s *GetTextScanResultRequest) Validate() error {
	return dara.Validate(s)
}
