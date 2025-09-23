// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteMigrationOperationResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExecuteMigrationOperationResponseBody
  GetCode() *string 
  SetData(v *ExecuteMigrationOperationResponseBodyData) *ExecuteMigrationOperationResponseBody
  GetData() *ExecuteMigrationOperationResponseBodyData 
  SetDynamicCode(v string) *ExecuteMigrationOperationResponseBody
  GetDynamicCode() *string 
  SetDynamicMessage(v string) *ExecuteMigrationOperationResponseBody
  GetDynamicMessage() *string 
  SetHttpStatusCode(v int32) *ExecuteMigrationOperationResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExecuteMigrationOperationResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteMigrationOperationResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteMigrationOperationResponseBody
  GetSuccess() *bool 
}

type ExecuteMigrationOperationResponseBody struct {
  // example:
  // 
  // Topic.NotFound
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  Data *ExecuteMigrationOperationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // example:
  // 
  // InstanceId
  DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
  // example:
  // 
  // instanceId
  DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  // example:
  // 
  // The topic already exists.
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // example:
  // 
  // 814BCD66-2D76-5080-AAD2-E50E5BDB0995
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteMigrationOperationResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteMigrationOperationResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteMigrationOperationResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExecuteMigrationOperationResponseBody) GetData() *ExecuteMigrationOperationResponseBodyData  {
  return s.Data
}

func (s *ExecuteMigrationOperationResponseBody) GetDynamicCode() *string  {
  return s.DynamicCode
}

func (s *ExecuteMigrationOperationResponseBody) GetDynamicMessage() *string  {
  return s.DynamicMessage
}

func (s *ExecuteMigrationOperationResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecuteMigrationOperationResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteMigrationOperationResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteMigrationOperationResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteMigrationOperationResponseBody) SetCode(v string) *ExecuteMigrationOperationResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteMigrationOperationResponseBody) SetData(v *ExecuteMigrationOperationResponseBodyData) *ExecuteMigrationOperationResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteMigrationOperationResponseBody) SetDynamicCode(v string) *ExecuteMigrationOperationResponseBody {
  s.DynamicCode = &v
  return s
}

func (s *ExecuteMigrationOperationResponseBody) SetDynamicMessage(v string) *ExecuteMigrationOperationResponseBody {
  s.DynamicMessage = &v
  return s
}

func (s *ExecuteMigrationOperationResponseBody) SetHttpStatusCode(v int32) *ExecuteMigrationOperationResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteMigrationOperationResponseBody) SetMessage(v string) *ExecuteMigrationOperationResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteMigrationOperationResponseBody) SetRequestId(v string) *ExecuteMigrationOperationResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteMigrationOperationResponseBody) SetSuccess(v bool) *ExecuteMigrationOperationResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteMigrationOperationResponseBody) Validate() error {
  return dara.Validate(s)
}

type ExecuteMigrationOperationResponseBodyData struct {
  // example:
  // 
  // MIGRATE_WAIT
  BusinessStatus *string `json:"businessStatus,omitempty" xml:"businessStatus,omitempty"`
  // example:
  // 
  // 2022-08-01 20:05:50
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
  // example:
  // 
  // 6
  MigrationId *int64 `json:"migrationId,omitempty" xml:"migrationId,omitempty"`
  // example:
  // 
  // 110
  OperationId *int64 `json:"operationId,omitempty" xml:"operationId,omitempty"`
  // example:
  // 
  // group01
  OperationKey *string `json:"operationKey,omitempty" xml:"operationKey,omitempty"`
  OperationParam *ExecuteMigrationOperationResponseBodyDataOperationParam `json:"operationParam,omitempty" xml:"operationParam,omitempty" type:"Struct"`
  OperationResult *ExecuteMigrationOperationResponseBodyDataOperationResult `json:"operationResult,omitempty" xml:"operationResult,omitempty" type:"Struct"`
  // example:
  // 
  // DONE
  OperationStatus *string `json:"operationStatus,omitempty" xml:"operationStatus,omitempty"`
  // example:
  // 
  // CONNECT_NETWORK
  OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty"`
  // example:
  // 
  // MIGRATE_METADATA
  StageType *string `json:"stageType,omitempty" xml:"stageType,omitempty"`
  // example:
  // 
  // 2022-08-01 20:05:50
  UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ExecuteMigrationOperationResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteMigrationOperationResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteMigrationOperationResponseBodyData) GetBusinessStatus() *string  {
  return s.BusinessStatus
}

func (s *ExecuteMigrationOperationResponseBodyData) GetCreateTime() *string  {
  return s.CreateTime
}

func (s *ExecuteMigrationOperationResponseBodyData) GetMigrationId() *int64  {
  return s.MigrationId
}

func (s *ExecuteMigrationOperationResponseBodyData) GetOperationId() *int64  {
  return s.OperationId
}

func (s *ExecuteMigrationOperationResponseBodyData) GetOperationKey() *string  {
  return s.OperationKey
}

func (s *ExecuteMigrationOperationResponseBodyData) GetOperationParam() *ExecuteMigrationOperationResponseBodyDataOperationParam  {
  return s.OperationParam
}

func (s *ExecuteMigrationOperationResponseBodyData) GetOperationResult() *ExecuteMigrationOperationResponseBodyDataOperationResult  {
  return s.OperationResult
}

func (s *ExecuteMigrationOperationResponseBodyData) GetOperationStatus() *string  {
  return s.OperationStatus
}

func (s *ExecuteMigrationOperationResponseBodyData) GetOperationType() *string  {
  return s.OperationType
}

func (s *ExecuteMigrationOperationResponseBodyData) GetStageType() *string  {
  return s.StageType
}

func (s *ExecuteMigrationOperationResponseBodyData) GetUpdateTime() *string  {
  return s.UpdateTime
}

func (s *ExecuteMigrationOperationResponseBodyData) SetBusinessStatus(v string) *ExecuteMigrationOperationResponseBodyData {
  s.BusinessStatus = &v
  return s
}

func (s *ExecuteMigrationOperationResponseBodyData) SetCreateTime(v string) *ExecuteMigrationOperationResponseBodyData {
  s.CreateTime = &v
  return s
}

func (s *ExecuteMigrationOperationResponseBodyData) SetMigrationId(v int64) *ExecuteMigrationOperationResponseBodyData {
  s.MigrationId = &v
  return s
}

func (s *ExecuteMigrationOperationResponseBodyData) SetOperationId(v int64) *ExecuteMigrationOperationResponseBodyData {
  s.OperationId = &v
  return s
}

func (s *ExecuteMigrationOperationResponseBodyData) SetOperationKey(v string) *ExecuteMigrationOperationResponseBodyData {
  s.OperationKey = &v
  return s
}

func (s *ExecuteMigrationOperationResponseBodyData) SetOperationParam(v *ExecuteMigrationOperationResponseBodyDataOperationParam) *ExecuteMigrationOperationResponseBodyData {
  s.OperationParam = v
  return s
}

func (s *ExecuteMigrationOperationResponseBodyData) SetOperationResult(v *ExecuteMigrationOperationResponseBodyDataOperationResult) *ExecuteMigrationOperationResponseBodyData {
  s.OperationResult = v
  return s
}

func (s *ExecuteMigrationOperationResponseBodyData) SetOperationStatus(v string) *ExecuteMigrationOperationResponseBodyData {
  s.OperationStatus = &v
  return s
}

func (s *ExecuteMigrationOperationResponseBodyData) SetOperationType(v string) *ExecuteMigrationOperationResponseBodyData {
  s.OperationType = &v
  return s
}

func (s *ExecuteMigrationOperationResponseBodyData) SetStageType(v string) *ExecuteMigrationOperationResponseBodyData {
  s.StageType = &v
  return s
}

func (s *ExecuteMigrationOperationResponseBodyData) SetUpdateTime(v string) *ExecuteMigrationOperationResponseBodyData {
  s.UpdateTime = &v
  return s
}

func (s *ExecuteMigrationOperationResponseBodyData) Validate() error {
  return dara.Validate(s)
}

type ExecuteMigrationOperationResponseBodyDataOperationParam struct {
  // example:
  // 
  // {}
  ParamData interface{} `json:"paramData,omitempty" xml:"paramData,omitempty"`
}

func (s ExecuteMigrationOperationResponseBodyDataOperationParam) String() string {
  return dara.Prettify(s)
}

func (s ExecuteMigrationOperationResponseBodyDataOperationParam) GoString() string {
  return s.String()
}

func (s *ExecuteMigrationOperationResponseBodyDataOperationParam) GetParamData() interface{}  {
  return s.ParamData
}

func (s *ExecuteMigrationOperationResponseBodyDataOperationParam) SetParamData(v interface{}) *ExecuteMigrationOperationResponseBodyDataOperationParam {
  s.ParamData = v
  return s
}

func (s *ExecuteMigrationOperationResponseBodyDataOperationParam) Validate() error {
  return dara.Validate(s)
}

type ExecuteMigrationOperationResponseBodyDataOperationResult struct {
  // example:
  // 
  // {}
  ResultData interface{} `json:"resultData,omitempty" xml:"resultData,omitempty"`
}

func (s ExecuteMigrationOperationResponseBodyDataOperationResult) String() string {
  return dara.Prettify(s)
}

func (s ExecuteMigrationOperationResponseBodyDataOperationResult) GoString() string {
  return s.String()
}

func (s *ExecuteMigrationOperationResponseBodyDataOperationResult) GetResultData() interface{}  {
  return s.ResultData
}

func (s *ExecuteMigrationOperationResponseBodyDataOperationResult) SetResultData(v interface{}) *ExecuteMigrationOperationResponseBodyDataOperationResult {
  s.ResultData = v
  return s
}

func (s *ExecuteMigrationOperationResponseBodyDataOperationResult) Validate() error {
  return dara.Validate(s)
}

