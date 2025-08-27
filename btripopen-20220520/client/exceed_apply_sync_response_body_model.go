// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExceedApplySyncResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExceedApplySyncResponseBody
  GetCode() *string 
  SetMessage(v string) *ExceedApplySyncResponseBody
  GetMessage() *string 
  SetModule(v bool) *ExceedApplySyncResponseBody
  GetModule() *bool 
  SetRequestId(v string) *ExceedApplySyncResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExceedApplySyncResponseBody
  GetSuccess() *bool 
  SetTraceId(v string) *ExceedApplySyncResponseBody
  GetTraceId() *string 
}

type ExceedApplySyncResponseBody struct {
  // example:
  // 
  // SUCCESS
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // module。
  // 
  // example:
  // 
  // {\\"list\\": [], \\"pageSize\\": 20, \\"pageNo\\": 1}
  Module *bool `json:"module,omitempty" xml:"module,omitempty"`
  // example:
  // 
  // 407543AF-2BD9-5890-BD92-9D1AB7218B27
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  // example:
  // 
  // 21041ce316577904808056433edbb2
  TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ExceedApplySyncResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExceedApplySyncResponseBody) GoString() string {
  return s.String()
}

func (s *ExceedApplySyncResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExceedApplySyncResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExceedApplySyncResponseBody) GetModule() *bool  {
  return s.Module
}

func (s *ExceedApplySyncResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExceedApplySyncResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExceedApplySyncResponseBody) GetTraceId() *string  {
  return s.TraceId
}

func (s *ExceedApplySyncResponseBody) SetCode(v string) *ExceedApplySyncResponseBody {
  s.Code = &v
  return s
}

func (s *ExceedApplySyncResponseBody) SetMessage(v string) *ExceedApplySyncResponseBody {
  s.Message = &v
  return s
}

func (s *ExceedApplySyncResponseBody) SetModule(v bool) *ExceedApplySyncResponseBody {
  s.Module = &v
  return s
}

func (s *ExceedApplySyncResponseBody) SetRequestId(v string) *ExceedApplySyncResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExceedApplySyncResponseBody) SetSuccess(v bool) *ExceedApplySyncResponseBody {
  s.Success = &v
  return s
}

func (s *ExceedApplySyncResponseBody) SetTraceId(v string) *ExceedApplySyncResponseBody {
  s.TraceId = &v
  return s
}

func (s *ExceedApplySyncResponseBody) Validate() error {
  return dara.Validate(s)
}

