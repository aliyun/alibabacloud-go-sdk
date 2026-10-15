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
  SetQuery(v string) *ExportCipStatsRequest
  GetQuery() *string 
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
  // Specifies whether monthly indexing is supported. Valid values:
  // 
  // - **true**: Supported.
  // 
  // - **false**: Not supported.
  // 
  // example:
  // 
  // true
  ByMonth *bool `json:"ByMonth,omitempty" xml:"ByMonth,omitempty"`
  // The end time of the query. Format: yyyy-MM-dd HH:mm:ss.
  // 
  // example:
  // 
  // 2024-04-16 09:00:00
  EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
  // The export type. Valid values:
  // 
  // - **level**: Export by risk level.
  // 
  // - **label**: Export by label.
  // 
  // example:
  // 
  // label
  ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
  // The task label to export.
  // 
  // example:
  // 
  // xx
  Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
  // The query condition.
  // 
  // example:
  // 
  // {"Classify":"guard-scene"}
  Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
  // The region ID.
  // 
  // example:
  // 
  // cn-shanghai
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The resource type.
  // 
  // example:
  // 
  // text
  ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
  // The service code.
  // 
  // example:
  // 
  // baselineCheck
  ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
  // The start time of the query. Format: yyyy-MM-dd HH:mm:ss.
  // 
  // example:
  // 
  // 2024-04-15 09:00:00
  StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
  // The UID of the RAM user.
  // 
  // example:
  // 
  // 268220485413130979
  SubUid *string `json:"SubUid,omitempty" xml:"SubUid,omitempty"`
  // The type. Valid values:
  // 
  // - **cip**: Content Moderation invoke volume statistics.
  // 
  // - **risk_level**: Content Moderation security risk level statistics.
  // 
  // - **content_moderation**: AI Safety Guardrail content compliance risk level and tag statistics.
  // 
  // - **sensitive_data**: AI Safety Guardrail sensitive data risk level and tag statistics.
  // 
  // - **prompt_attack**: AI Safety Guardrail prompt risk level and tag statistics.
  // 
  // example:
  // 
  // risk_level
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

func (s *ExportCipStatsRequest) GetQuery() *string  {
  return s.Query
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

func (s *ExportCipStatsRequest) SetQuery(v string) *ExportCipStatsRequest {
  s.Query = &v
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

