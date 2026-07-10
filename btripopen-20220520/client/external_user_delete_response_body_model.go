// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalUserDeleteResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExternalUserDeleteResponseBody
  GetCode() *string 
  SetMessage(v string) *ExternalUserDeleteResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExternalUserDeleteResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExternalUserDeleteResponseBody
  GetSuccess() *bool 
  SetTraceId(v string) *ExternalUserDeleteResponseBody
  GetTraceId() *string 
}

type ExternalUserDeleteResponseBody struct {
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ExternalUserDeleteResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserDeleteResponseBody) GoString() string {
  return s.String()
}

func (s *ExternalUserDeleteResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExternalUserDeleteResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExternalUserDeleteResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExternalUserDeleteResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExternalUserDeleteResponseBody) GetTraceId() *string  {
  return s.TraceId
}

func (s *ExternalUserDeleteResponseBody) SetCode(v string) *ExternalUserDeleteResponseBody {
  s.Code = &v
  return s
}

func (s *ExternalUserDeleteResponseBody) SetMessage(v string) *ExternalUserDeleteResponseBody {
  s.Message = &v
  return s
}

func (s *ExternalUserDeleteResponseBody) SetRequestId(v string) *ExternalUserDeleteResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExternalUserDeleteResponseBody) SetSuccess(v bool) *ExternalUserDeleteResponseBody {
  s.Success = &v
  return s
}

func (s *ExternalUserDeleteResponseBody) SetTraceId(v string) *ExternalUserDeleteResponseBody {
  s.TraceId = &v
  return s
}

func (s *ExternalUserDeleteResponseBody) Validate() error {
  return dara.Validate(s)
}

