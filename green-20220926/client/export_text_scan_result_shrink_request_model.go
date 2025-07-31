// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportTextScanResultShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndDate(v string) *ExportTextScanResultShrinkRequest
  GetEndDate() *string 
  SetQueryShrink(v string) *ExportTextScanResultShrinkRequest
  GetQueryShrink() *string 
  SetRegionId(v string) *ExportTextScanResultShrinkRequest
  GetRegionId() *string 
  SetStartDate(v string) *ExportTextScanResultShrinkRequest
  GetStartDate() *string 
}

type ExportTextScanResultShrinkRequest struct {
  // example:
  // 
  // 2024-03-11 10:00:00
  EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
  QueryShrink *string `json:"Query,omitempty" xml:"Query,omitempty"`
  // example:
  // 
  // cn-shanghai
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // example:
  // 
  // 2024-03-10 10:00:00
  StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportTextScanResultShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportTextScanResultShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExportTextScanResultShrinkRequest) GetEndDate() *string  {
  return s.EndDate
}

func (s *ExportTextScanResultShrinkRequest) GetQueryShrink() *string  {
  return s.QueryShrink
}

func (s *ExportTextScanResultShrinkRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportTextScanResultShrinkRequest) GetStartDate() *string  {
  return s.StartDate
}

func (s *ExportTextScanResultShrinkRequest) SetEndDate(v string) *ExportTextScanResultShrinkRequest {
  s.EndDate = &v
  return s
}

func (s *ExportTextScanResultShrinkRequest) SetQueryShrink(v string) *ExportTextScanResultShrinkRequest {
  s.QueryShrink = &v
  return s
}

func (s *ExportTextScanResultShrinkRequest) SetRegionId(v string) *ExportTextScanResultShrinkRequest {
  s.RegionId = &v
  return s
}

func (s *ExportTextScanResultShrinkRequest) SetStartDate(v string) *ExportTextScanResultShrinkRequest {
  s.StartDate = &v
  return s
}

func (s *ExportTextScanResultShrinkRequest) Validate() error {
  return dara.Validate(s)
}

