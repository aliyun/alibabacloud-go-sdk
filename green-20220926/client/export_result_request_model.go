// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportResultRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCurrentPage(v int32) *ExportResultRequest
  GetCurrentPage() *int32 
  SetEndDate(v string) *ExportResultRequest
  GetEndDate() *string 
  SetPageSize(v int32) *ExportResultRequest
  GetPageSize() *int32 
  SetQuery(v string) *ExportResultRequest
  GetQuery() *string 
  SetRegionId(v string) *ExportResultRequest
  GetRegionId() *string 
  SetSort(v map[string]*string) *ExportResultRequest
  GetSort() map[string]*string 
  SetSource(v string) *ExportResultRequest
  GetSource() *string 
  SetStartDate(v string) *ExportResultRequest
  GetStartDate() *string 
}

type ExportResultRequest struct {
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
  Sort map[string]*string `json:"Sort,omitempty" xml:"Sort,omitempty"`
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

func (s ExportResultRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportResultRequest) GoString() string {
  return s.String()
}

func (s *ExportResultRequest) GetCurrentPage() *int32  {
  return s.CurrentPage
}

func (s *ExportResultRequest) GetEndDate() *string  {
  return s.EndDate
}

func (s *ExportResultRequest) GetPageSize() *int32  {
  return s.PageSize
}

func (s *ExportResultRequest) GetQuery() *string  {
  return s.Query
}

func (s *ExportResultRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportResultRequest) GetSort() map[string]*string  {
  return s.Sort
}

func (s *ExportResultRequest) GetSource() *string  {
  return s.Source
}

func (s *ExportResultRequest) GetStartDate() *string  {
  return s.StartDate
}

func (s *ExportResultRequest) SetCurrentPage(v int32) *ExportResultRequest {
  s.CurrentPage = &v
  return s
}

func (s *ExportResultRequest) SetEndDate(v string) *ExportResultRequest {
  s.EndDate = &v
  return s
}

func (s *ExportResultRequest) SetPageSize(v int32) *ExportResultRequest {
  s.PageSize = &v
  return s
}

func (s *ExportResultRequest) SetQuery(v string) *ExportResultRequest {
  s.Query = &v
  return s
}

func (s *ExportResultRequest) SetRegionId(v string) *ExportResultRequest {
  s.RegionId = &v
  return s
}

func (s *ExportResultRequest) SetSort(v map[string]*string) *ExportResultRequest {
  s.Sort = v
  return s
}

func (s *ExportResultRequest) SetSource(v string) *ExportResultRequest {
  s.Source = &v
  return s
}

func (s *ExportResultRequest) SetStartDate(v string) *ExportResultRequest {
  s.StartDate = &v
  return s
}

func (s *ExportResultRequest) Validate() error {
  return dara.Validate(s)
}

