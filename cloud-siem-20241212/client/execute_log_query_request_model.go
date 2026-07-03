// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteLogQueryRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndTime(v string) *ExecuteLogQueryRequest
  GetEndTime() *string 
  SetExtendContentPacked(v string) *ExecuteLogQueryRequest
  GetExtendContentPacked() *string 
  SetLang(v string) *ExecuteLogQueryRequest
  GetLang() *string 
  SetLogCondition(v string) *ExecuteLogQueryRequest
  GetLogCondition() *string 
  SetLogProjectName(v string) *ExecuteLogQueryRequest
  GetLogProjectName() *string 
  SetLogQuery(v string) *ExecuteLogQueryRequest
  GetLogQuery() *string 
  SetLogRegionId(v string) *ExecuteLogQueryRequest
  GetLogRegionId() *string 
  SetLogStoreName(v string) *ExecuteLogQueryRequest
  GetLogStoreName() *string 
  SetLogUserId(v int64) *ExecuteLogQueryRequest
  GetLogUserId() *int64 
  SetNormalizationSchemaId(v string) *ExecuteLogQueryRequest
  GetNormalizationSchemaId() *string 
  SetRegionId(v string) *ExecuteLogQueryRequest
  GetRegionId() *string 
  SetRoleFor(v int64) *ExecuteLogQueryRequest
  GetRoleFor() *int64 
  SetStartTime(v string) *ExecuteLogQueryRequest
  GetStartTime() *string 
}

type ExecuteLogQueryRequest struct {
  // The end time.
  // 
  // example:
  // 
  // 1733269771123
  EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
  // Specifies whether to pack non-standard fields into the extension field extend_content. Valid values:
  // 
  // - enabled: Enabled.
  // 
  // - disabled: Disabled.
  // 
  // example:
  // 
  // enabled
  ExtendContentPacked *string `json:"ExtendContentPacked,omitempty" xml:"ExtendContentPacked,omitempty"`
  // The language of the response. Valid values:
  // 
  // - **zh*	- (default): Chinese.
  // 
  // - **en**: English.
  // 
  // example:
  // 
  // zh
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  LogCondition *string `json:"LogCondition,omitempty" xml:"LogCondition,omitempty"`
  // The Simple Log Service project name.
  // 
  // example:
  // 
  // slsaudit-center-173326*******-cn-hangzhou
  LogProjectName *string `json:"LogProjectName,omitempty" xml:"LogProjectName,omitempty"`
  // The Simple Log Service query statement.
  // 
  // example:
  // 
  // *
  LogQuery *string `json:"LogQuery,omitempty" xml:"LogQuery,omitempty"`
  // The log storage region ID.
  // 
  // example:
  // 
  // cn-hangzhou
  LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
  // The Simple Log Service project name.
  // 
  // example:
  // 
  // huawei-cn-cfw
  LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
  // The user ID for data access.
  // 
  // example:
  // 
  // 173326*******
  LogUserId *int64 `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
  // If packing is enabled, you must specify NormalizationSchemaId.
  // 
  // example:
  // 
  // WAF_ALERT_ACTIVITY
  NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
  // The region where the threat analysis data management center is located. Specify the management center based on the region of your assets. Valid values:
  // 
  // - cn-hangzhou: the asset is in the Chinese mainland.
  // 
  // - ap-southeast-1: the asset is outside the Chinese mainland.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The user ID of the member to which the administrator switches the view.
  // 
  // example:
  // 
  // 173326*******
  RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
  // The start time.
  // 
  // example:
  // 
  // 1733269771123
  StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ExecuteLogQueryRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteLogQueryRequest) GoString() string {
  return s.String()
}

func (s *ExecuteLogQueryRequest) GetEndTime() *string  {
  return s.EndTime
}

func (s *ExecuteLogQueryRequest) GetExtendContentPacked() *string  {
  return s.ExtendContentPacked
}

func (s *ExecuteLogQueryRequest) GetLang() *string  {
  return s.Lang
}

func (s *ExecuteLogQueryRequest) GetLogCondition() *string  {
  return s.LogCondition
}

func (s *ExecuteLogQueryRequest) GetLogProjectName() *string  {
  return s.LogProjectName
}

func (s *ExecuteLogQueryRequest) GetLogQuery() *string  {
  return s.LogQuery
}

func (s *ExecuteLogQueryRequest) GetLogRegionId() *string  {
  return s.LogRegionId
}

func (s *ExecuteLogQueryRequest) GetLogStoreName() *string  {
  return s.LogStoreName
}

func (s *ExecuteLogQueryRequest) GetLogUserId() *int64  {
  return s.LogUserId
}

func (s *ExecuteLogQueryRequest) GetNormalizationSchemaId() *string  {
  return s.NormalizationSchemaId
}

func (s *ExecuteLogQueryRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExecuteLogQueryRequest) GetRoleFor() *int64  {
  return s.RoleFor
}

func (s *ExecuteLogQueryRequest) GetStartTime() *string  {
  return s.StartTime
}

func (s *ExecuteLogQueryRequest) SetEndTime(v string) *ExecuteLogQueryRequest {
  s.EndTime = &v
  return s
}

func (s *ExecuteLogQueryRequest) SetExtendContentPacked(v string) *ExecuteLogQueryRequest {
  s.ExtendContentPacked = &v
  return s
}

func (s *ExecuteLogQueryRequest) SetLang(v string) *ExecuteLogQueryRequest {
  s.Lang = &v
  return s
}

func (s *ExecuteLogQueryRequest) SetLogCondition(v string) *ExecuteLogQueryRequest {
  s.LogCondition = &v
  return s
}

func (s *ExecuteLogQueryRequest) SetLogProjectName(v string) *ExecuteLogQueryRequest {
  s.LogProjectName = &v
  return s
}

func (s *ExecuteLogQueryRequest) SetLogQuery(v string) *ExecuteLogQueryRequest {
  s.LogQuery = &v
  return s
}

func (s *ExecuteLogQueryRequest) SetLogRegionId(v string) *ExecuteLogQueryRequest {
  s.LogRegionId = &v
  return s
}

func (s *ExecuteLogQueryRequest) SetLogStoreName(v string) *ExecuteLogQueryRequest {
  s.LogStoreName = &v
  return s
}

func (s *ExecuteLogQueryRequest) SetLogUserId(v int64) *ExecuteLogQueryRequest {
  s.LogUserId = &v
  return s
}

func (s *ExecuteLogQueryRequest) SetNormalizationSchemaId(v string) *ExecuteLogQueryRequest {
  s.NormalizationSchemaId = &v
  return s
}

func (s *ExecuteLogQueryRequest) SetRegionId(v string) *ExecuteLogQueryRequest {
  s.RegionId = &v
  return s
}

func (s *ExecuteLogQueryRequest) SetRoleFor(v int64) *ExecuteLogQueryRequest {
  s.RoleFor = &v
  return s
}

func (s *ExecuteLogQueryRequest) SetStartTime(v string) *ExecuteLogQueryRequest {
  s.StartTime = &v
  return s
}

func (s *ExecuteLogQueryRequest) Validate() error {
  return dara.Validate(s)
}

