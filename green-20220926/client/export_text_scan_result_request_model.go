// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportTextScanResultRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndDate(v string) *ExportTextScanResultRequest
  GetEndDate() *string 
  SetQuery(v map[string]*string) *ExportTextScanResultRequest
  GetQuery() map[string]*string 
  SetRegionId(v string) *ExportTextScanResultRequest
  GetRegionId() *string 
  SetStartDate(v string) *ExportTextScanResultRequest
  GetStartDate() *string 
}

type ExportTextScanResultRequest struct {
  // End time of the query, in the format yyyy-MM-dd HH:mm:ss.
  // 
  // example:
  // 
  // 2024-03-11 10:00:00
  EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
  // Query conditions.
  Query map[string]*string `json:"Query,omitempty" xml:"Query,omitempty"`
  // Region ID.
  // 
  // example:
  // 
  // cn-shanghai
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // Start time of the query, in the format yyyy-MM-dd HH:mm:ss.
  // 
  // example:
  // 
  // 2024-03-10 10:00:00
  StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportTextScanResultRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportTextScanResultRequest) GoString() string {
  return s.String()
}

func (s *ExportTextScanResultRequest) GetEndDate() *string  {
  return s.EndDate
}

func (s *ExportTextScanResultRequest) GetQuery() map[string]*string  {
  return s.Query
}

func (s *ExportTextScanResultRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportTextScanResultRequest) GetStartDate() *string  {
  return s.StartDate
}

func (s *ExportTextScanResultRequest) SetEndDate(v string) *ExportTextScanResultRequest {
  s.EndDate = &v
  return s
}

func (s *ExportTextScanResultRequest) SetQuery(v map[string]*string) *ExportTextScanResultRequest {
  s.Query = v
  return s
}

func (s *ExportTextScanResultRequest) SetRegionId(v string) *ExportTextScanResultRequest {
  s.RegionId = &v
  return s
}

func (s *ExportTextScanResultRequest) SetStartDate(v string) *ExportTextScanResultRequest {
  s.StartDate = &v
  return s
}

func (s *ExportTextScanResultRequest) Validate() error {
  return dara.Validate(s)
}

