// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportScanResultShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCurrentPage(v int32) *ExportScanResultShrinkRequest
  GetCurrentPage() *int32 
  SetEndDate(v string) *ExportScanResultShrinkRequest
  GetEndDate() *string 
  SetPageSize(v int32) *ExportScanResultShrinkRequest
  GetPageSize() *int32 
  SetQueryShrink(v string) *ExportScanResultShrinkRequest
  GetQueryShrink() *string 
  SetRegionId(v string) *ExportScanResultShrinkRequest
  GetRegionId() *string 
  SetResourceType(v string) *ExportScanResultShrinkRequest
  GetResourceType() *string 
  SetSortShrink(v string) *ExportScanResultShrinkRequest
  GetSortShrink() *string 
  SetStartDate(v string) *ExportScanResultShrinkRequest
  GetStartDate() *string 
}

type ExportScanResultShrinkRequest struct {
  // Current page number.
  // 
  // example:
  // 
  // 1
  CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
  // End time of the query, in the format yyyy-MM-dd HH:mm:ss.
  // 
  // example:
  // 
  // 2024-03-11 10:00:00
  EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
  // Page size.
  // 
  // example:
  // 
  // 20
  PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
  // Query content.
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
  // text
  ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
  // Sort fields.
  SortShrink *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
  // Start time of the query, in the format yyyy-MM-dd HH:mm:ss.
  // 
  // example:
  // 
  // 2024-03-10 10:00:00
  StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportScanResultShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportScanResultShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExportScanResultShrinkRequest) GetCurrentPage() *int32  {
  return s.CurrentPage
}

func (s *ExportScanResultShrinkRequest) GetEndDate() *string  {
  return s.EndDate
}

func (s *ExportScanResultShrinkRequest) GetPageSize() *int32  {
  return s.PageSize
}

func (s *ExportScanResultShrinkRequest) GetQueryShrink() *string  {
  return s.QueryShrink
}

func (s *ExportScanResultShrinkRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportScanResultShrinkRequest) GetResourceType() *string  {
  return s.ResourceType
}

func (s *ExportScanResultShrinkRequest) GetSortShrink() *string  {
  return s.SortShrink
}

func (s *ExportScanResultShrinkRequest) GetStartDate() *string  {
  return s.StartDate
}

func (s *ExportScanResultShrinkRequest) SetCurrentPage(v int32) *ExportScanResultShrinkRequest {
  s.CurrentPage = &v
  return s
}

func (s *ExportScanResultShrinkRequest) SetEndDate(v string) *ExportScanResultShrinkRequest {
  s.EndDate = &v
  return s
}

func (s *ExportScanResultShrinkRequest) SetPageSize(v int32) *ExportScanResultShrinkRequest {
  s.PageSize = &v
  return s
}

func (s *ExportScanResultShrinkRequest) SetQueryShrink(v string) *ExportScanResultShrinkRequest {
  s.QueryShrink = &v
  return s
}

func (s *ExportScanResultShrinkRequest) SetRegionId(v string) *ExportScanResultShrinkRequest {
  s.RegionId = &v
  return s
}

func (s *ExportScanResultShrinkRequest) SetResourceType(v string) *ExportScanResultShrinkRequest {
  s.ResourceType = &v
  return s
}

func (s *ExportScanResultShrinkRequest) SetSortShrink(v string) *ExportScanResultShrinkRequest {
  s.SortShrink = &v
  return s
}

func (s *ExportScanResultShrinkRequest) SetStartDate(v string) *ExportScanResultShrinkRequest {
  s.StartDate = &v
  return s
}

func (s *ExportScanResultShrinkRequest) Validate() error {
  return dara.Validate(s)
}

