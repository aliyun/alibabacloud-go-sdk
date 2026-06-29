// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteManualNodeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExecuteManualNodeResponseBody
  GetCode() *string 
  SetFlowId(v string) *ExecuteManualNodeResponseBody
  GetFlowId() *string 
  SetHttpStatusCode(v int32) *ExecuteManualNodeResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExecuteManualNodeResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteManualNodeResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteManualNodeResponseBody
  GetSuccess() *bool 
}

type ExecuteManualNodeResponseBody struct {
  // The error code. A value of OK indicates that the request was successful.
  // 
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The workflow ID.
  // 
  // example:
  // 
  // f_1231_121324
  FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
  // The HTTP status code returned by the backend.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // The error message.
  // 
  // example:
  // 
  // successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful.
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteManualNodeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteManualNodeResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteManualNodeResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExecuteManualNodeResponseBody) GetFlowId() *string  {
  return s.FlowId
}

func (s *ExecuteManualNodeResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecuteManualNodeResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteManualNodeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteManualNodeResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteManualNodeResponseBody) SetCode(v string) *ExecuteManualNodeResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteManualNodeResponseBody) SetFlowId(v string) *ExecuteManualNodeResponseBody {
  s.FlowId = &v
  return s
}

func (s *ExecuteManualNodeResponseBody) SetHttpStatusCode(v int32) *ExecuteManualNodeResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteManualNodeResponseBody) SetMessage(v string) *ExecuteManualNodeResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteManualNodeResponseBody) SetRequestId(v string) *ExecuteManualNodeResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteManualNodeResponseBody) SetSuccess(v bool) *ExecuteManualNodeResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteManualNodeResponseBody) Validate() error {
  return dara.Validate(s)
}

