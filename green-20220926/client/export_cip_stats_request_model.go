// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportCipStatsRequest interface {
  dara.Model
  String() string
  GoString() string
  SetByMonth(v bool) *ExportCipStatsRequest
  GetByMonth() *bool 
  SetEndDate(v string) *ExportCipStatsRequest
  GetEndDate() *string 
  SetExportType(v string) *ExportCipStatsRequest
  GetExportType() *string 
  SetLabel(v string) *ExportCipStatsRequest
  GetLabel() *string 
  SetRegionId(v string) *ExportCipStatsRequest
  GetRegionId() *string 
  SetResourceType(v string) *ExportCipStatsRequest
  GetResourceType() *string 
  SetServiceCode(v string) *ExportCipStatsRequest
  GetServiceCode() *string 
  SetStartDate(v string) *ExportCipStatsRequest
  GetStartDate() *string 
  SetSubUid(v string) *ExportCipStatsRequest
  GetSubUid() *string 
  SetType(v string) *ExportCipStatsRequest
  GetType() *string 
}

type ExportCipStatsRequest struct {
  // example:
  // 
  // true
  ByMonth *bool `json:"ByMonth,omitempty" xml:"ByMonth,omitempty"`
  // example:
  // 
  // 2024-04-16 09:00:00
  EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
  ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
  // example:
  // 
  // xx
  Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
  // example:
  // 
  // cn-shanghai
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // example:
  // 
  // text
  ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
  ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
  // example:
  // 
  // 2024-04-15 09:00:00
  StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
  // example:
  // 
  // 268220485413130979
  SubUid *string `json:"SubUid,omitempty" xml:"SubUid,omitempty"`
  Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ExportCipStatsRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportCipStatsRequest) GoString() string {
  return s.String()
}

func (s *ExportCipStatsRequest) GetByMonth() *bool  {
  return s.ByMonth
}

func (s *ExportCipStatsRequest) GetEndDate() *string  {
  return s.EndDate
}

func (s *ExportCipStatsRequest) GetExportType() *string  {
  return s.ExportType
}

func (s *ExportCipStatsRequest) GetLabel() *string  {
  return s.Label
}

func (s *ExportCipStatsRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportCipStatsRequest) GetResourceType() *string  {
  return s.ResourceType
}

func (s *ExportCipStatsRequest) GetServiceCode() *string  {
  return s.ServiceCode
}

func (s *ExportCipStatsRequest) GetStartDate() *string  {
  return s.StartDate
}

func (s *ExportCipStatsRequest) GetSubUid() *string  {
  return s.SubUid
}

func (s *ExportCipStatsRequest) GetType() *string  {
  return s.Type
}

func (s *ExportCipStatsRequest) SetByMonth(v bool) *ExportCipStatsRequest {
  s.ByMonth = &v
  return s
}

func (s *ExportCipStatsRequest) SetEndDate(v string) *ExportCipStatsRequest {
  s.EndDate = &v
  return s
}

func (s *ExportCipStatsRequest) SetExportType(v string) *ExportCipStatsRequest {
  s.ExportType = &v
  return s
}

func (s *ExportCipStatsRequest) SetLabel(v string) *ExportCipStatsRequest {
  s.Label = &v
  return s
}

func (s *ExportCipStatsRequest) SetRegionId(v string) *ExportCipStatsRequest {
  s.RegionId = &v
  return s
}

func (s *ExportCipStatsRequest) SetResourceType(v string) *ExportCipStatsRequest {
  s.ResourceType = &v
  return s
}

func (s *ExportCipStatsRequest) SetServiceCode(v string) *ExportCipStatsRequest {
  s.ServiceCode = &v
  return s
}

func (s *ExportCipStatsRequest) SetStartDate(v string) *ExportCipStatsRequest {
  s.StartDate = &v
  return s
}

func (s *ExportCipStatsRequest) SetSubUid(v string) *ExportCipStatsRequest {
  s.SubUid = &v
  return s
}

func (s *ExportCipStatsRequest) SetType(v string) *ExportCipStatsRequest {
  s.Type = &v
  return s
}

func (s *ExportCipStatsRequest) Validate() error {
  return dara.Validate(s)
}

