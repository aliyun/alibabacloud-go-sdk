// This file is auto-generated, don't edit it. Thanks.
package endToEndRealTimeDialog

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
}

type Error struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *Error) SetCode(v string) *Error {
  s.Code = &v
  return s
}

func (s *Error) SetMessage(v string) *Error {
  s.Message = &v
  return s
}

func (s *Error) Validate() error {
  return dara.Validate(s)
}

