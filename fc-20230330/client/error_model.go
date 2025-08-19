// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iError interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *Error
  GetCode() *string 
  SetMessage(v string) *Error
  GetMessage() *string 
  SetRequestId(v string) *Error
  GetRequestId() *string 
}

type Error struct {
  // example:
  // 
  // FunctionNotFound
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // function not found
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 1-64e70cf1-5cbef92ea8fc8c42899cf5d1
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s Error) String() string {
  return dara.Prettify(s)
}

func (s Error) GoString() string {
  return s.String()
}

func (s *Error) GetCode() *string  {
  return s.Code
}

func (s *Error) GetMessage() *string  {
  return s.Message
}

func (s *Error) GetRequestId() *string  {
  return s.RequestId
}

func (s *Error) SetCode(v string) *Error {
  s.Code = &v
  return s
}

func (s *Error) SetMessage(v string) *Error {
  s.Message = &v
  return s
}

func (s *Error) SetRequestId(v string) *Error {
  s.RequestId = &v
  return s
}

func (s *Error) Validate() error {
  return dara.Validate(s)
}

