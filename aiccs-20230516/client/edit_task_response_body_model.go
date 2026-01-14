// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditTaskResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAccessDeniedDetail(v string) *EditTaskResponseBody
  GetAccessDeniedDetail() *string 
  SetCode(v int64) *EditTaskResponseBody
  GetCode() *int64 
  SetMessage(v string) *EditTaskResponseBody
  GetMessage() *string 
  SetModel(v *EditTaskResponseBodyModel) *EditTaskResponseBody
  GetModel() *EditTaskResponseBodyModel 
  SetRequestId(v string) *EditTaskResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EditTaskResponseBody
  GetSuccess() *bool 
  SetTimestamp(v int64) *EditTaskResponseBody
  GetTimestamp() *int64 
}

type EditTaskResponseBody struct {
  AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
  // example:
  // 
  // 0
  Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // 示例值示例值
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  Model *EditTaskResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
  // example:
  // 
  // 5453cc9b-02bc-4dbb-9527-f28a9100b912
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
  // example:
  // 
  // 1686225227338
  Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s EditTaskResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditTaskResponseBody) GoString() string {
  return s.String()
}

func (s *EditTaskResponseBody) GetAccessDeniedDetail() *string  {
  return s.AccessDeniedDetail
}

func (s *EditTaskResponseBody) GetCode() *int64  {
  return s.Code
}

func (s *EditTaskResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EditTaskResponseBody) GetModel() *EditTaskResponseBodyModel  {
  return s.Model
}

func (s *EditTaskResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditTaskResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EditTaskResponseBody) GetTimestamp() *int64  {
  return s.Timestamp
}

func (s *EditTaskResponseBody) SetAccessDeniedDetail(v string) *EditTaskResponseBody {
  s.AccessDeniedDetail = &v
  return s
}

func (s *EditTaskResponseBody) SetCode(v int64) *EditTaskResponseBody {
  s.Code = &v
  return s
}

func (s *EditTaskResponseBody) SetMessage(v string) *EditTaskResponseBody {
  s.Message = &v
  return s
}

func (s *EditTaskResponseBody) SetModel(v *EditTaskResponseBodyModel) *EditTaskResponseBody {
  s.Model = v
  return s
}

func (s *EditTaskResponseBody) SetRequestId(v string) *EditTaskResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditTaskResponseBody) SetSuccess(v bool) *EditTaskResponseBody {
  s.Success = &v
  return s
}

func (s *EditTaskResponseBody) SetTimestamp(v int64) *EditTaskResponseBody {
  s.Timestamp = &v
  return s
}

func (s *EditTaskResponseBody) Validate() error {
  if s.Model != nil {
    if err := s.Model.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EditTaskResponseBodyModel struct {
  // 任务ID
  // 
  // example:
  // 
  // 92
  TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s EditTaskResponseBodyModel) String() string {
  return dara.Prettify(s)
}

func (s EditTaskResponseBodyModel) GoString() string {
  return s.String()
}

func (s *EditTaskResponseBodyModel) GetTaskId() *int64  {
  return s.TaskId
}

func (s *EditTaskResponseBodyModel) SetTaskId(v int64) *EditTaskResponseBodyModel {
  s.TaskId = &v
  return s
}

func (s *EditTaskResponseBodyModel) Validate() error {
  return dara.Validate(s)
}

