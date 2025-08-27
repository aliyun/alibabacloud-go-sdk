// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalUserAddResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExternalUserAddResponseBody
  GetCode() *string 
  SetMessage(v string) *ExternalUserAddResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExternalUserAddResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExternalUserAddResponseBody
  GetSuccess() *bool 
  SetTraceId(v string) *ExternalUserAddResponseBody
  GetTraceId() *string 
}

type ExternalUserAddResponseBody struct {
  // example:
  // 
  // success
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // example:
  // 
  // B72B39C8-32DE-558D-AD1C-D53F11F6ADFE
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  // example:
  // 
  // 210f079416784321627628333de4ab
  TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ExternalUserAddResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserAddResponseBody) GoString() string {
  return s.String()
}

func (s *ExternalUserAddResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExternalUserAddResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExternalUserAddResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExternalUserAddResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExternalUserAddResponseBody) GetTraceId() *string  {
  return s.TraceId
}

func (s *ExternalUserAddResponseBody) SetCode(v string) *ExternalUserAddResponseBody {
  s.Code = &v
  return s
}

func (s *ExternalUserAddResponseBody) SetMessage(v string) *ExternalUserAddResponseBody {
  s.Message = &v
  return s
}

func (s *ExternalUserAddResponseBody) SetRequestId(v string) *ExternalUserAddResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExternalUserAddResponseBody) SetSuccess(v bool) *ExternalUserAddResponseBody {
  s.Success = &v
  return s
}

func (s *ExternalUserAddResponseBody) SetTraceId(v string) *ExternalUserAddResponseBody {
  s.TraceId = &v
  return s
}

func (s *ExternalUserAddResponseBody) Validate() error {
  return dara.Validate(s)
}

