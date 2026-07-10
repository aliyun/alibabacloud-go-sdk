// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalUserUpdateResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExternalUserUpdateResponseBody
  GetCode() *string 
  SetMessage(v string) *ExternalUserUpdateResponseBody
  GetMessage() *string 
  SetRequestId(v string) *ExternalUserUpdateResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExternalUserUpdateResponseBody
  GetSuccess() *bool 
  SetTraceId(v string) *ExternalUserUpdateResponseBody
  GetTraceId() *string 
}

type ExternalUserUpdateResponseBody struct {
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ExternalUserUpdateResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserUpdateResponseBody) GoString() string {
  return s.String()
}

func (s *ExternalUserUpdateResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExternalUserUpdateResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExternalUserUpdateResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExternalUserUpdateResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExternalUserUpdateResponseBody) GetTraceId() *string  {
  return s.TraceId
}

func (s *ExternalUserUpdateResponseBody) SetCode(v string) *ExternalUserUpdateResponseBody {
  s.Code = &v
  return s
}

func (s *ExternalUserUpdateResponseBody) SetMessage(v string) *ExternalUserUpdateResponseBody {
  s.Message = &v
  return s
}

func (s *ExternalUserUpdateResponseBody) SetRequestId(v string) *ExternalUserUpdateResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExternalUserUpdateResponseBody) SetSuccess(v bool) *ExternalUserUpdateResponseBody {
  s.Success = &v
  return s
}

func (s *ExternalUserUpdateResponseBody) SetTraceId(v string) *ExternalUserUpdateResponseBody {
  s.TraceId = &v
  return s
}

func (s *ExternalUserUpdateResponseBody) Validate() error {
  return dara.Validate(s)
}

