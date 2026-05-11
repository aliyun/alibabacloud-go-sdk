// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportStatisticalDataRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBeginTimeLeftRange(v int64) *ExportStatisticalDataRequest
  GetBeginTimeLeftRange() *int64 
  SetBeginTimeRightRange(v int64) *ExportStatisticalDataRequest
  GetBeginTimeRightRange() *int64 
  SetExportType(v string) *ExportStatisticalDataRequest
  GetExportType() *string 
  SetInstanceId(v string) *ExportStatisticalDataRequest
  GetInstanceId() *string 
  SetTimeUnit(v string) *ExportStatisticalDataRequest
  GetTimeUnit() *string 
}

type ExportStatisticalDataRequest struct {
  // example:
  // 
  // 1582266750353
  BeginTimeLeftRange *int64 `json:"BeginTimeLeftRange,omitempty" xml:"BeginTimeLeftRange,omitempty"`
  // example:
  // 
  // 1582266750353
  BeginTimeRightRange *int64 `json:"BeginTimeRightRange,omitempty" xml:"BeginTimeRightRange,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // statistical
  ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 29b52d70-d9fe-4fe0-8476-8aaacbcfdc84
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // Hour
  TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s ExportStatisticalDataRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportStatisticalDataRequest) GoString() string {
  return s.String()
}

func (s *ExportStatisticalDataRequest) GetBeginTimeLeftRange() *int64  {
  return s.BeginTimeLeftRange
}

func (s *ExportStatisticalDataRequest) GetBeginTimeRightRange() *int64  {
  return s.BeginTimeRightRange
}

func (s *ExportStatisticalDataRequest) GetExportType() *string  {
  return s.ExportType
}

func (s *ExportStatisticalDataRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExportStatisticalDataRequest) GetTimeUnit() *string  {
  return s.TimeUnit
}

func (s *ExportStatisticalDataRequest) SetBeginTimeLeftRange(v int64) *ExportStatisticalDataRequest {
  s.BeginTimeLeftRange = &v
  return s
}

func (s *ExportStatisticalDataRequest) SetBeginTimeRightRange(v int64) *ExportStatisticalDataRequest {
  s.BeginTimeRightRange = &v
  return s
}

func (s *ExportStatisticalDataRequest) SetExportType(v string) *ExportStatisticalDataRequest {
  s.ExportType = &v
  return s
}

func (s *ExportStatisticalDataRequest) SetInstanceId(v string) *ExportStatisticalDataRequest {
  s.InstanceId = &v
  return s
}

func (s *ExportStatisticalDataRequest) SetTimeUnit(v string) *ExportStatisticalDataRequest {
  s.TimeUnit = &v
  return s
}

func (s *ExportStatisticalDataRequest) Validate() error {
  return dara.Validate(s)
}

