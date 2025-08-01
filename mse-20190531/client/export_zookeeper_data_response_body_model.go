// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportZookeeperDataResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExportZookeeperDataResponseBodyData) *ExportZookeeperDataResponseBody
  GetData() *ExportZookeeperDataResponseBodyData 
  SetDynamicMessage(v string) *ExportZookeeperDataResponseBody
  GetDynamicMessage() *string 
  SetErrorCode(v string) *ExportZookeeperDataResponseBody
  GetErrorCode() *string 
  SetHttpStatusCode(v string) *ExportZookeeperDataResponseBody
  GetHttpStatusCode() *string 
  SetMessage(v string) *ExportZookeeperDataResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExportZookeeperDataResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExportZookeeperDataResponseBody
  GetSuccess() *bool 
}

type ExportZookeeperDataResponseBody struct {
  // The details of the data.
  Data *ExportZookeeperDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
  // 
  // > If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
  // 
  // example:
  // 
  // The specified parameter is invalid.
  DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
  // The error code returned if the request failed.
  // 
  // example:
  // 
  // mse-100-000
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // The HTTP status code returned.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // The message returned.
  // 
  // 	- If the request is successful, a success message is returned.
  // 
  // 	- If the request fails, an error message is returned.
  // 
  // example:
  // 
  // The request was successfully processed.
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // 25EA0A83-9007-4E87-808C-637BE1A****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- `true`: The request was successful.
  // 
  // 	- `false`: The request failed.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportZookeeperDataResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportZookeeperDataResponseBody) GoString() string {
  return s.String()
}

func (s *ExportZookeeperDataResponseBody) GetData() *ExportZookeeperDataResponseBodyData  {
  return s.Data
}

func (s *ExportZookeeperDataResponseBody) GetDynamicMessage() *string  {
  return s.DynamicMessage
}

func (s *ExportZookeeperDataResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *ExportZookeeperDataResponseBody) GetHttpStatusCode() *string  {
  return s.HttpStatusCode
}

func (s *ExportZookeeperDataResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExportZookeeperDataResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportZookeeperDataResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExportZookeeperDataResponseBody) SetData(v *ExportZookeeperDataResponseBodyData) *ExportZookeeperDataResponseBody {
  s.Data = v
  return s
}

func (s *ExportZookeeperDataResponseBody) SetDynamicMessage(v string) *ExportZookeeperDataResponseBody {
  s.DynamicMessage = &v
  return s
}

func (s *ExportZookeeperDataResponseBody) SetErrorCode(v string) *ExportZookeeperDataResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *ExportZookeeperDataResponseBody) SetHttpStatusCode(v string) *ExportZookeeperDataResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExportZookeeperDataResponseBody) SetMessage(v string) *ExportZookeeperDataResponseBody {
  s.Message = &v
  return s
}

func (s *ExportZookeeperDataResponseBody) SetRequestId(v string) *ExportZookeeperDataResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportZookeeperDataResponseBody) SetSuccess(v bool) *ExportZookeeperDataResponseBody {
  s.Success = &v
  return s
}

func (s *ExportZookeeperDataResponseBody) Validate() error {
  return dara.Validate(s)
}

type ExportZookeeperDataResponseBodyData struct {
  // The content of a task.
  ContentMap map[string]interface{} `json:"ContentMap,omitempty" xml:"ContentMap,omitempty"`
  // The creation time.
  // 
  // example:
  // 
  // 1631001140913
  CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
  // The type of the object that is exported. Valid values:
  // 
  // 	- transactionLog: transaction logs
  // 
  // 	- snapshot: snapshots
  // 
  // example:
  // 
  // snapshot
  ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
  // The extended information.
  // 
  // example:
  // 
  // {}
  Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
  // The ID of the task.
  // 
  // example:
  // 
  // 1
  Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
  // The ID of the instance
  // 
  // example:
  // 
  // mse-cn-st21ri2****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The ID of the associated task at the underlying layer. This parameter is used only to troubleshoot failures.
  // 
  // example:
  // 
  // 10
  KubeoneTaskIds *string `json:"KubeoneTaskIds,omitempty" xml:"KubeoneTaskIds,omitempty"`
  // The status of the task. Valid values:
  // 
  // 	- CREATE: The object is being created.
  // 
  // 	- RUNNING: The task is running.
  // 
  // 	- FINISH: The task is completed.
  // 
  // 	- FAILED: The task fails.
  // 
  // 	- EXPIRE: The task has expired.
  // 
  // example:
  // 
  // FINISH
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
  // The last update time.
  // 
  // example:
  // 
  // 1632979237663
  UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ExportZookeeperDataResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExportZookeeperDataResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExportZookeeperDataResponseBodyData) GetContentMap() map[string]interface{}  {
  return s.ContentMap
}

func (s *ExportZookeeperDataResponseBodyData) GetCreateTime() *int64  {
  return s.CreateTime
}

func (s *ExportZookeeperDataResponseBodyData) GetExportType() *string  {
  return s.ExportType
}

func (s *ExportZookeeperDataResponseBodyData) GetExtend() *string  {
  return s.Extend
}

func (s *ExportZookeeperDataResponseBodyData) GetId() *int32  {
  return s.Id
}

func (s *ExportZookeeperDataResponseBodyData) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExportZookeeperDataResponseBodyData) GetKubeoneTaskIds() *string  {
  return s.KubeoneTaskIds
}

func (s *ExportZookeeperDataResponseBodyData) GetStatus() *string  {
  return s.Status
}

func (s *ExportZookeeperDataResponseBodyData) GetUpdateTime() *int64  {
  return s.UpdateTime
}

func (s *ExportZookeeperDataResponseBodyData) SetContentMap(v map[string]interface{}) *ExportZookeeperDataResponseBodyData {
  s.ContentMap = v
  return s
}

func (s *ExportZookeeperDataResponseBodyData) SetCreateTime(v int64) *ExportZookeeperDataResponseBodyData {
  s.CreateTime = &v
  return s
}

func (s *ExportZookeeperDataResponseBodyData) SetExportType(v string) *ExportZookeeperDataResponseBodyData {
  s.ExportType = &v
  return s
}

func (s *ExportZookeeperDataResponseBodyData) SetExtend(v string) *ExportZookeeperDataResponseBodyData {
  s.Extend = &v
  return s
}

func (s *ExportZookeeperDataResponseBodyData) SetId(v int32) *ExportZookeeperDataResponseBodyData {
  s.Id = &v
  return s
}

func (s *ExportZookeeperDataResponseBodyData) SetInstanceId(v string) *ExportZookeeperDataResponseBodyData {
  s.InstanceId = &v
  return s
}

func (s *ExportZookeeperDataResponseBodyData) SetKubeoneTaskIds(v string) *ExportZookeeperDataResponseBodyData {
  s.KubeoneTaskIds = &v
  return s
}

func (s *ExportZookeeperDataResponseBodyData) SetStatus(v string) *ExportZookeeperDataResponseBodyData {
  s.Status = &v
  return s
}

func (s *ExportZookeeperDataResponseBodyData) SetUpdateTime(v int64) *ExportZookeeperDataResponseBodyData {
  s.UpdateTime = &v
  return s
}

func (s *ExportZookeeperDataResponseBodyData) Validate() error {
  return dara.Validate(s)
}

