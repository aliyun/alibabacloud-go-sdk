// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTriggerNodeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExecuteTriggerNodeResponseBody
  GetCode() *string 
  SetHttpStatusCode(v int32) *ExecuteTriggerNodeResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExecuteTriggerNodeResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExecuteTriggerNodeResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteTriggerNodeResponseBody
  GetSuccess() *bool 
}

type ExecuteTriggerNodeResponseBody struct {
  // example:
  // 
  // OK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // internal error
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteTriggerNodeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTriggerNodeResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteTriggerNodeResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExecuteTriggerNodeResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExecuteTriggerNodeResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExecuteTriggerNodeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteTriggerNodeResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteTriggerNodeResponseBody) SetCode(v string) *ExecuteTriggerNodeResponseBody {
  s.Code = &v
  return s
}

func (s *ExecuteTriggerNodeResponseBody) SetHttpStatusCode(v int32) *ExecuteTriggerNodeResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExecuteTriggerNodeResponseBody) SetMessage(v string) *ExecuteTriggerNodeResponseBody {
  s.Message = &v
  return s
}

func (s *ExecuteTriggerNodeResponseBody) SetRequestId(v string) *ExecuteTriggerNodeResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteTriggerNodeResponseBody) SetSuccess(v bool) *ExecuteTriggerNodeResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteTriggerNodeResponseBody) Validate() error {
  return dara.Validate(s)
}

