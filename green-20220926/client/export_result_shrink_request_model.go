// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportResultShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCurrentPage(v int32) *ExportResultShrinkRequest
  GetCurrentPage() *int32 
  SetEndDate(v string) *ExportResultShrinkRequest
  GetEndDate() *string 
  SetPageSize(v int32) *ExportResultShrinkRequest
  GetPageSize() *int32 
  SetQuery(v string) *ExportResultShrinkRequest
  GetQuery() *string 
  SetRegionId(v string) *ExportResultShrinkRequest
  GetRegionId() *string 
  SetSortShrink(v string) *ExportResultShrinkRequest
  GetSortShrink() *string 
  SetSource(v string) *ExportResultShrinkRequest
  GetSource() *string 
  SetStartDate(v string) *ExportResultShrinkRequest
  GetStartDate() *string 
}

type ExportResultShrinkRequest struct {
  // Page number of the query result. Default is 1.
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
  // Number of items per page in the query result.
  // 
  // example:
  // 
  // 20
  PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
  // Query condition.
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
  // Operation source.
  // 
  // example:
  // 
  // disposal
  Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
  // Start date.
  // 
  // example:
  // 
  // 2023-08-11 09:00:19
  StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportResultShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportResultShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExportResultShrinkRequest) GetCurrentPage() *int32  {
  return s.CurrentPage
}

func (s *ExportResultShrinkRequest) GetEndDate() *string  {
  return s.EndDate
}

func (s *ExportResultShrinkRequest) GetPageSize() *int32  {
  return s.PageSize
}

func (s *ExportResultShrinkRequest) GetQuery() *string  {
  return s.Query
}

func (s *ExportResultShrinkRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportResultShrinkRequest) GetSortShrink() *string  {
  return s.SortShrink
}

func (s *ExportResultShrinkRequest) GetSource() *string  {
  return s.Source
}

func (s *ExportResultShrinkRequest) GetStartDate() *string  {
  return s.StartDate
}

func (s *ExportResultShrinkRequest) SetCurrentPage(v int32) *ExportResultShrinkRequest {
  s.CurrentPage = &v
  return s
}

func (s *ExportResultShrinkRequest) SetEndDate(v string) *ExportResultShrinkRequest {
  s.EndDate = &v
  return s
}

func (s *ExportResultShrinkRequest) SetPageSize(v int32) *ExportResultShrinkRequest {
  s.PageSize = &v
  return s
}

func (s *ExportResultShrinkRequest) SetQuery(v string) *ExportResultShrinkRequest {
  s.Query = &v
  return s
}

func (s *ExportResultShrinkRequest) SetRegionId(v string) *ExportResultShrinkRequest {
  s.RegionId = &v
  return s
}

func (s *ExportResultShrinkRequest) SetSortShrink(v string) *ExportResultShrinkRequest {
  s.SortShrink = &v
  return s
}

func (s *ExportResultShrinkRequest) SetSource(v string) *ExportResultShrinkRequest {
  s.Source = &v
  return s
}

func (s *ExportResultShrinkRequest) SetStartDate(v string) *ExportResultShrinkRequest {
  s.StartDate = &v
  return s
}

func (s *ExportResultShrinkRequest) Validate() error {
  return dara.Validate(s)
}

