// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteWorkflowResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ExecuteWorkflowResponseBody
  GetCode() *int32 
  SetData(v *ExecuteWorkflowResponseBodyData) *ExecuteWorkflowResponseBody
  GetData() *ExecuteWorkflowResponseBodyData 
  SetMessage(v string) *ExecuteWorkflowResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteWorkflowResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteWorkflowResponseBody
  GetSuccess() *bool 
}

type ExecuteWorkflowResponseBody struct {
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // If the request is successful, the ID of the workflow instance is returned.
  Data *ExecuteWorkflowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // The error message that is returned only if the corresponding error occurs.
  // 
  // example:
  // 
  // Cannot find product according to your domain.
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 4F68ABED-AC31-4412-9297-D9A8F0401108
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteWorkflowResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteWorkflowResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteWorkflowResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *ExecuteWorkflowResponseBody) GetData() *ExecuteWorkflowResponseBodyData  {
  return s.Data
}

func (s *ExecuteWorkflowResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteWorkflowResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteWorkflowResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteWorkflowResponseBody) SetCode(v int32) *ExecuteWorkflowResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteWorkflowResponseBody) SetData(v *ExecuteWorkflowResponseBodyData) *ExecuteWorkflowResponseBody {
  s.Data = v
  return s
}

func (s *ExecuteWorkflowResponseBody) SetMessage(v string) *ExecuteWorkflowResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteWorkflowResponseBody) SetRequestId(v string) *ExecuteWorkflowResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteWorkflowResponseBody) SetSuccess(v bool) *ExecuteWorkflowResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteWorkflowResponseBody) Validate() error {
  return dara.Validate(s)
}

type ExecuteWorkflowResponseBodyData struct {
  // The workflow instance ID.
  // 
  // example:
  // 
  // 111111
  WfInstanceId *int64 `json:"WfInstanceId,omitempty" xml:"WfInstanceId,omitempty"`
}

func (s ExecuteWorkflowResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecuteWorkflowResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecuteWorkflowResponseBodyData) GetWfInstanceId() *int64  {
  return s.WfInstanceId
}

func (s *ExecuteWorkflowResponseBodyData) SetWfInstanceId(v int64) *ExecuteWorkflowResponseBodyData {
  s.WfInstanceId = &v
  return s
}

func (s *ExecuteWorkflowResponseBodyData) Validate() error {
  return dara.Validate(s)
}

