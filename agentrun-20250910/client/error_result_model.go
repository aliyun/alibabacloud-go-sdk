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
  // SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  // 错误信息描述
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // 唯一的请求标识符，用于问题追踪
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

