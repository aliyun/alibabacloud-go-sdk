// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iErrorResponse interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ErrorResponse
  GetCode() *string 
  SetMessage(v string) *ErrorResponse
  GetMessage() *string 
  SetRequestId(v string) *ErrorResponse
  GetRequestId() *string 
}

type ErrorResponse struct {
  // example:
  // 
  // InternalServerError
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  // example:
  // 
  // internal server error
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // example:
  // 
  // 90D6B8F5-FE97-4509-9AAB-367836C51818
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ErrorResponse) String() string {
  return dara.Prettify(s)
}

func (s ErrorResponse) GoString() string {
  return s.String()
}

func (s *ErrorResponse) GetCode() *string  {
  return s.Code
}

func (s *ErrorResponse) GetMessage() *string  {
  return s.Message
}

func (s *ErrorResponse) GetRequestId() *string  {
  return s.RequestId
}

func (s *ErrorResponse) SetCode(v string) *ErrorResponse {
  s.Code = &v
  return s
}

func (s *ErrorResponse) SetMessage(v string) *ErrorResponse {
  s.Message = &v
  return s
}

func (s *ErrorResponse) SetRequestId(v string) *ErrorResponse {
  s.RequestId = &v
  return s
}

func (s *ErrorResponse) Validate() error {
  return dara.Validate(s)
}

