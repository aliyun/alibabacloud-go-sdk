// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportOssCheckStatRequest interface {
  dara.Model
  String() string
  GoString() string
  SetByMonth(v bool) *ExportOssCheckStatRequest
  GetByMonth() *bool 
  SetEndDate(v string) *ExportOssCheckStatRequest
  GetEndDate() *string 
  SetParentTaskId(v string) *ExportOssCheckStatRequest
  GetParentTaskId() *string 
  SetRegionId(v string) *ExportOssCheckStatRequest
  GetRegionId() *string 
  SetStartDate(v string) *ExportOssCheckStatRequest
  GetStartDate() *string 
}

type ExportOssCheckStatRequest struct {
  // Whether to support monthly indexing. Values: -true: supported. -false: not supported.
  // 
  // example:
  // 
  // true
  ByMonth *bool `json:"ByMonth,omitempty" xml:"ByMonth,omitempty"`
  // End time of the query, in the format yyyy-MM-dd HH:mm:ss.
  // 
  // example:
  // 
  // 2024-03-11 10:00:00
  EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
  // OSS detection task ID.
  // 
  // example:
  // 
  // P_UX0K5X
  ParentTaskId *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
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

func (s ExportOssCheckStatRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportOssCheckStatRequest) GoString() string {
  return s.String()
}

func (s *ExportOssCheckStatRequest) GetByMonth() *bool  {
  return s.ByMonth
}

func (s *ExportOssCheckStatRequest) GetEndDate() *string  {
  return s.EndDate
}

func (s *ExportOssCheckStatRequest) GetParentTaskId() *string  {
  return s.ParentTaskId
}

func (s *ExportOssCheckStatRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportOssCheckStatRequest) GetStartDate() *string  {
  return s.StartDate
}

func (s *ExportOssCheckStatRequest) SetByMonth(v bool) *ExportOssCheckStatRequest {
  s.ByMonth = &v
  return s
}

func (s *ExportOssCheckStatRequest) SetEndDate(v string) *ExportOssCheckStatRequest {
  s.EndDate = &v
  return s
}

func (s *ExportOssCheckStatRequest) SetParentTaskId(v string) *ExportOssCheckStatRequest {
  s.ParentTaskId = &v
  return s
}

func (s *ExportOssCheckStatRequest) SetRegionId(v string) *ExportOssCheckStatRequest {
  s.RegionId = &v
  return s
}

func (s *ExportOssCheckStatRequest) SetStartDate(v string) *ExportOssCheckStatRequest {
  s.StartDate = &v
  return s
}

func (s *ExportOssCheckStatRequest) Validate() error {
  return dara.Validate(s)
}

