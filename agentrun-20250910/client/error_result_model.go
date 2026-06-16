// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iErrorResult interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ErrorResult
  GetCode() *string 
  SetMessage(v string) *ErrorResult
  GetMessage() *string 
  SetRequestId(v string) *ErrorResult
  GetRequestId() *string 
}

type ErrorResult struct {
  // The error code. SUCCESS indicates the request was successful. Otherwise, a specific error code is returned, such as ERR_BAD_REQUEST, ERR_VALIDATION_FAILED, or ERR_INTERNAL_SERVER_ERROR.
  // 
  // example:
  // 
  // SUCCESS
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  // A human-readable description of the error.
  // 
  // example:
  // 
  // Conflict Transaction, process failed
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // The unique request identifier, used for troubleshooting.
  // 
  // example:
  // 
  // 55D4BE40-2811-5CFB-8482-E0E98D575B1E
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ErrorResult) String() string {
  return dara.Prettify(s)
}

func (s ErrorResult) GoString() string {
  return s.String()
}

func (s *ErrorResult) GetCode() *string  {
  return s.Code
}

func (s *ErrorResult) GetMessage() *string  {
  return s.Message
}

func (s *ErrorResult) GetRequestId() *string  {
  return s.RequestId
}

func (s *ErrorResult) SetCode(v string) *ErrorResult {
  s.Code = &v
  return s
}

func (s *ErrorResult) SetMessage(v string) *ErrorResult {
  s.Message = &v
  return s
}

func (s *ErrorResult) SetRequestId(v string) *ErrorResult {
  s.RequestId = &v
  return s
}

func (s *ErrorResult) Validate() error {
  return dara.Validate(s)
}

