// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportWarningRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDealed(v string) *ExportWarningRequest
  GetDealed() *string 
  SetExportType(v string) *ExportWarningRequest
  GetExportType() *string 
  SetGroupId(v int64) *ExportWarningRequest
  GetGroupId() *int64 
  SetIsCleartextPwd(v int32) *ExportWarningRequest
  GetIsCleartextPwd() *int32 
  SetIsSummaryExport(v int32) *ExportWarningRequest
  GetIsSummaryExport() *int32 
  SetLang(v string) *ExportWarningRequest
  GetLang() *string 
  SetRiskIds(v string) *ExportWarningRequest
  GetRiskIds() *string 
  SetRiskLevels(v string) *ExportWarningRequest
  GetRiskLevels() *string 
  SetRiskName(v string) *ExportWarningRequest
  GetRiskName() *string 
  SetSourceIp(v string) *ExportWarningRequest
  GetSourceIp() *string 
  SetStatusList(v string) *ExportWarningRequest
  GetStatusList() *string 
  SetStrategyId(v int64) *ExportWarningRequest
  GetStrategyId() *int64 
  SetSubTypeNames(v string) *ExportWarningRequest
  GetSubTypeNames() *string 
  SetTypeName(v string) *ExportWarningRequest
  GetTypeName() *string 
  SetTypeNames(v string) *ExportWarningRequest
  GetTypeNames() *string 
  SetUuids(v string) *ExportWarningRequest
  GetUuids() *string 
}

type ExportWarningRequest struct {
  // Specifies whether the baseline issue is handled. Valid values:
  // 
  // - **Y**: The issue is handled.
  // 
  // - **N**: The issue is not handled.
  // 
  // example:
  // 
  // N
  Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
  // The type of the export task. Set the value to **hc_check_warning**, which indicates exporting baseline check results.
  // 
  // example:
  // 
  // hc_check_warning
  ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
  // The ID of the server group.
  // 
  // > You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to obtain the ID of the server group.
  // 
  // example:
  // 
  // 13007754
  GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
  // Specifies the export method for weak password baseline check results. Valid values:
  // 
  // - **0**: Weak passwords are masked before export.
  // 
  // - **1**: Weak passwords are exported in plaintext.
  // 
  // example:
  // 
  // 0
  IsCleartextPwd *int32 `json:"IsCleartextPwd,omitempty" xml:"IsCleartextPwd,omitempty"`
  // Specifies whether to export the baseline check results in summary mode. Valid values:
  // 
  // - **0**: The results are not exported in summary mode.
  // 
  // - **1**: The results are exported in summary mode.
  // 
  // example:
  // 
  // 1
  IsSummaryExport *int32 `json:"IsSummaryExport,omitempty" xml:"IsSummaryExport,omitempty"`
  // Specifies the language type for the request and response messages. Default value: **zh**. Valid values:
  // 
  // - **zh**: Chinese
  // 
  // - **en**: English
  // 
  // example:
  // 
  // zh
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // The IDs of the risk items in the baseline check results. Separate multiple risk item IDs with commas (,).
  // 
  // example:
  // 
  // 123,124
  RiskIds *string `json:"RiskIds,omitempty" xml:"RiskIds,omitempty"`
  // The risk levels of the baselines. Separate multiple risk levels with commas (,). Valid values:
  // 
  // - **high**: High
  // 
  // - **medium**: Medium
  // 
  // - **low**: Low
  // 
  // example:
  // 
  // high,medium
  RiskLevels *string `json:"RiskLevels,omitempty" xml:"RiskLevels,omitempty"`
  // The name of the baseline.
  // 
  // example:
  // 
  // Alibaba Cloud Standard - Windows 2016/2019  Security Baseline
  RiskName *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
  // The source IP address.
  // 
  // example:
  // 
  // 192.0.XX.XX
  SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
  // The statuses of the risk items in the baseline check results. Separate multiple status values with commas (,). Valid values:
  // 
  // - **3**: Passed
  // 
  // - **1**: Failed
  // 
  // example:
  // 
  // 1,3
  StatusList *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
  // The ID of the baseline check policy.
  // 
  // example:
  // 
  // 12
  StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
  // Specifies the baseline subtypes to which the baseline check results to be exported belong. Separate multiple subtypes with commas (,).
  // 
  // > The value is the **TypeName*	- field under the **SubTypes*	- field returned by the [DescribeRiskType](~~DescribeRiskType~~) operation.
  // 
  // example:
  // 
  // hc_middleware_ack_master
  SubTypeNames *string `json:"SubTypeNames,omitempty" xml:"SubTypeNames,omitempty"`
  // Specifies the baseline type to which the baseline check results to be exported belong.
  // 
  // > The value is the **TypeName*	- field returned by the [DescribeRiskType](~~DescribeRiskType~~) operation. If both **TypeName*	- and **TypeNames*	- are specified, only **TypeName*	- takes effect.
  // 
  // example:
  // 
  // hc_container
  TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
  // Specifies the baseline types to which the baseline check results to be exported belong. Separate multiple types with commas (,).
  // 
  // > The value is the **TypeName*	- field returned by the [DescribeRiskType](~~DescribeRiskType~~) operation. If both **TypeName*	- and **TypeNames*	- are specified, only **TypeName*	- takes effect.
  // 
  // example:
  // 
  // hc_container,cis
  TypeNames *string `json:"TypeNames,omitempty" xml:"TypeNames,omitempty"`
  // The UUIDs of the server assets for which you want to export the baseline check results. Separate multiple UUIDs with commas (,).
  // 
  // example:
  // 
  // inet-7c676676-06fa-442e-90fb-b802e****,inet-7c676676-06fa-442e-90fb-b****
  Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s ExportWarningRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportWarningRequest) GoString() string {
  return s.String()
}

func (s *ExportWarningRequest) GetDealed() *string  {
  return s.Dealed
}

func (s *ExportWarningRequest) GetExportType() *string  {
  return s.ExportType
}

func (s *ExportWarningRequest) GetGroupId() *int64  {
  return s.GroupId
}

func (s *ExportWarningRequest) GetIsCleartextPwd() *int32  {
  return s.IsCleartextPwd
}

func (s *ExportWarningRequest) GetIsSummaryExport() *int32  {
  return s.IsSummaryExport
}

func (s *ExportWarningRequest) GetLang() *string  {
  return s.Lang
}

func (s *ExportWarningRequest) GetRiskIds() *string  {
  return s.RiskIds
}

func (s *ExportWarningRequest) GetRiskLevels() *string  {
  return s.RiskLevels
}

func (s *ExportWarningRequest) GetRiskName() *string  {
  return s.RiskName
}

func (s *ExportWarningRequest) GetSourceIp() *string  {
  return s.SourceIp
}

func (s *ExportWarningRequest) GetStatusList() *string  {
  return s.StatusList
}

func (s *ExportWarningRequest) GetStrategyId() *int64  {
  return s.StrategyId
}

func (s *ExportWarningRequest) GetSubTypeNames() *string  {
  return s.SubTypeNames
}

func (s *ExportWarningRequest) GetTypeName() *string  {
  return s.TypeName
}

func (s *ExportWarningRequest) GetTypeNames() *string  {
  return s.TypeNames
}

func (s *ExportWarningRequest) GetUuids() *string  {
  return s.Uuids
}

func (s *ExportWarningRequest) SetDealed(v string) *ExportWarningRequest {
  s.Dealed = &v
  return s
}

func (s *ExportWarningRequest) SetExportType(v string) *ExportWarningRequest {
  s.ExportType = &v
  return s
}

func (s *ExportWarningRequest) SetGroupId(v int64) *ExportWarningRequest {
  s.GroupId = &v
  return s
}

func (s *ExportWarningRequest) SetIsCleartextPwd(v int32) *ExportWarningRequest {
  s.IsCleartextPwd = &v
  return s
}

func (s *ExportWarningRequest) SetIsSummaryExport(v int32) *ExportWarningRequest {
  s.IsSummaryExport = &v
  return s
}

func (s *ExportWarningRequest) SetLang(v string) *ExportWarningRequest {
  s.Lang = &v
  return s
}

func (s *ExportWarningRequest) SetRiskIds(v string) *ExportWarningRequest {
  s.RiskIds = &v
  return s
}

func (s *ExportWarningRequest) SetRiskLevels(v string) *ExportWarningRequest {
  s.RiskLevels = &v
  return s
}

func (s *ExportWarningRequest) SetRiskName(v string) *ExportWarningRequest {
  s.RiskName = &v
  return s
}

func (s *ExportWarningRequest) SetSourceIp(v string) *ExportWarningRequest {
  s.SourceIp = &v
  return s
}

func (s *ExportWarningRequest) SetStatusList(v string) *ExportWarningRequest {
  s.StatusList = &v
  return s
}

func (s *ExportWarningRequest) SetStrategyId(v int64) *ExportWarningRequest {
  s.StrategyId = &v
  return s
}

func (s *ExportWarningRequest) SetSubTypeNames(v string) *ExportWarningRequest {
  s.SubTypeNames = &v
  return s
}

func (s *ExportWarningRequest) SetTypeName(v string) *ExportWarningRequest {
  s.TypeName = &v
  return s
}

func (s *ExportWarningRequest) SetTypeNames(v string) *ExportWarningRequest {
  s.TypeNames = &v
  return s
}

func (s *ExportWarningRequest) SetUuids(v string) *ExportWarningRequest {
  s.Uuids = &v
  return s
}

func (s *ExportWarningRequest) Validate() error {
  return dara.Validate(s)
}

