// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScanResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetScanResultRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *GetScanResultRequest
	GetEndDate() *string
	SetPageSize(v int32) *GetScanResultRequest
	GetPageSize() *int32
	SetQuery(v map[string]*string) *GetScanResultRequest
	GetQuery() map[string]*string
	SetRegionId(v string) *GetScanResultRequest
	GetRegionId() *string
	SetResourceType(v string) *GetScanResultRequest
	GetResourceType() *string
	SetSort(v map[string]*string) *GetScanResultRequest
	GetSort() map[string]*string
	SetStartDate(v string) *GetScanResultRequest
	GetStartDate() *string
}

type GetScanResultRequest struct {
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// image
	ResourceType *string            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Sort         map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 2023-08-11 09:00:19
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetScanResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScanResultRequest) GoString() string {
	return s.String()
}

func (s *GetScanResultRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetScanResultRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetScanResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetScanResultRequest) GetQuery() map[string]*string {
	return s.Query
}

func (s *GetScanResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetScanResultRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetScanResultRequest) GetSort() map[string]*string {
	return s.Sort
}

func (s *GetScanResultRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetScanResultRequest) SetCurrentPage(v int32) *GetScanResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetScanResultRequest) SetEndDate(v string) *GetScanResultRequest {
	s.EndDate = &v
	return s
}

func (s *GetScanResultRequest) SetPageSize(v int32) *GetScanResultRequest {
	s.PageSize = &v
	return s
}

func (s *GetScanResultRequest) SetQuery(v map[string]*string) *GetScanResultRequest {
	s.Query = v
	return s
}

func (s *GetScanResultRequest) SetRegionId(v string) *GetScanResultRequest {
	s.RegionId = &v
	return s
}

func (s *GetScanResultRequest) SetResourceType(v string) *GetScanResultRequest {
	s.ResourceType = &v
	return s
}

func (s *GetScanResultRequest) SetSort(v map[string]*string) *GetScanResultRequest {
	s.Sort = v
	return s
}

func (s *GetScanResultRequest) SetStartDate(v string) *GetScanResultRequest {
	s.StartDate = &v
	return s
}

func (s *GetScanResultRequest) Validate() error {
	return dara.Validate(s)
}
