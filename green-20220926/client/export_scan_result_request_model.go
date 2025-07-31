// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportScanResultRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCurrentPage(v int32) *ExportScanResultRequest
  GetCurrentPage() *int32 
  SetEndDate(v string) *ExportScanResultRequest
  GetEndDate() *string 
  SetPageSize(v int32) *ExportScanResultRequest
  GetPageSize() *int32 
  SetQuery(v map[string]*string) *ExportScanResultRequest
  GetQuery() map[string]*string 
  SetRegionId(v string) *ExportScanResultRequest
  GetRegionId() *string 
  SetResourceType(v string) *ExportScanResultRequest
  GetResourceType() *string 
  SetSort(v map[string]*string) *ExportScanResultRequest
  GetSort() map[string]*string 
  SetStartDate(v string) *ExportScanResultRequest
  GetStartDate() *string 
}

type ExportScanResultRequest struct {
  // example:
  // 
  // 1
  CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
  // example:
  // 
  // 2024-03-11 10:00:00
  EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
  // example:
  // 
  // 20
  PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
  Query map[string]*string `json:"Query,omitempty" xml:"Query,omitempty"`
  // example:
  // 
  // cn-shanghai
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // example:
  // 
  // text
  ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
  Sort map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
  // example:
  // 
  // 2024-03-10 10:00:00
  StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportScanResultRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportScanResultRequest) GoString() string {
  return s.String()
}

func (s *ExportScanResultRequest) GetCurrentPage() *int32  {
  return s.CurrentPage
}

func (s *ExportScanResultRequest) GetEndDate() *string  {
  return s.EndDate
}

func (s *ExportScanResultRequest) GetPageSize() *int32  {
  return s.PageSize
}

func (s *ExportScanResultRequest) GetQuery() map[string]*string  {
  return s.Query
}

func (s *ExportScanResultRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportScanResultRequest) GetResourceType() *string  {
  return s.ResourceType
}

func (s *ExportScanResultRequest) GetSort() map[string]*string  {
  return s.Sort
}

func (s *ExportScanResultRequest) GetStartDate() *string  {
  return s.StartDate
}

func (s *ExportScanResultRequest) SetCurrentPage(v int32) *ExportScanResultRequest {
  s.CurrentPage = &v
  return s
}

func (s *ExportScanResultRequest) SetEndDate(v string) *ExportScanResultRequest {
  s.EndDate = &v
  return s
}

func (s *ExportScanResultRequest) SetPageSize(v int32) *ExportScanResultRequest {
  s.PageSize = &v
  return s
}

func (s *ExportScanResultRequest) SetQuery(v map[string]*string) *ExportScanResultRequest {
  s.Query = v
  return s
}

func (s *ExportScanResultRequest) SetRegionId(v string) *ExportScanResultRequest {
  s.RegionId = &v
  return s
}

func (s *ExportScanResultRequest) SetResourceType(v string) *ExportScanResultRequest {
  s.ResourceType = &v
  return s
}

func (s *ExportScanResultRequest) SetSort(v map[string]*string) *ExportScanResultRequest {
  s.Sort = v
  return s
}

func (s *ExportScanResultRequest) SetStartDate(v string) *ExportScanResultRequest {
  s.StartDate = &v
  return s
}

func (s *ExportScanResultRequest) Validate() error {
  return dara.Validate(s)
}

