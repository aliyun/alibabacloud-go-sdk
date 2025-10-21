// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iErrorDetails interface {
  dara.Model
  String() string
  GoString() string
  SetColumnNumber(v string) *ErrorDetails
  GetColumnNumber() *string 
  SetEndColumnNumber(v string) *ErrorDetails
  GetEndColumnNumber() *string 
  SetEndLineNumber(v string) *ErrorDetails
  GetEndLineNumber() *string 
  SetInvalidflinkConf(v []*string) *ErrorDetails
  GetInvalidflinkConf() []*string 
  SetLineNumber(v string) *ErrorDetails
  GetLineNumber() *string 
  SetMessage(v string) *ErrorDetails
  GetMessage() *string 
}

type ErrorDetails struct {
  ColumnNumber *string `json:"columnNumber,omitempty" xml:"columnNumber,omitempty"`
  EndColumnNumber *string `json:"endColumnNumber,omitempty" xml:"endColumnNumber,omitempty"`
  EndLineNumber *string `json:"endLineNumber,omitempty" xml:"endLineNumber,omitempty"`
  InvalidflinkConf []*string `json:"invalidflinkConf,omitempty" xml:"invalidflinkConf,omitempty" type:"Repeated"`
  LineNumber *string `json:"lineNumber,omitempty" xml:"lineNumber,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ErrorDetails) String() string {
  return dara.Prettify(s)
}

func (s ErrorDetails) GoString() string {
  return s.String()
}

func (s *ErrorDetails) GetColumnNumber() *string  {
  return s.ColumnNumber
}

func (s *ErrorDetails) GetEndColumnNumber() *string  {
  return s.EndColumnNumber
}

func (s *ErrorDetails) GetEndLineNumber() *string  {
  return s.EndLineNumber
}

func (s *ErrorDetails) GetInvalidflinkConf() []*string  {
  return s.InvalidflinkConf
}

func (s *ErrorDetails) GetLineNumber() *string  {
  return s.LineNumber
}

func (s *ErrorDetails) GetMessage() *string  {
  return s.Message
}

func (s *ErrorDetails) SetColumnNumber(v string) *ErrorDetails {
  s.ColumnNumber = &v
  return s
}

func (s *ErrorDetails) SetEndColumnNumber(v string) *ErrorDetails {
  s.EndColumnNumber = &v
  return s
}

func (s *ErrorDetails) SetEndLineNumber(v string) *ErrorDetails {
  s.EndLineNumber = &v
  return s
}

func (s *ErrorDetails) SetInvalidflinkConf(v []*string) *ErrorDetails {
  s.InvalidflinkConf = v
  return s
}

func (s *ErrorDetails) SetLineNumber(v string) *ErrorDetails {
  s.LineNumber = &v
  return s
}

func (s *ErrorDetails) SetMessage(v string) *ErrorDetails {
  s.Message = &v
  return s
}

func (s *ErrorDetails) Validate() error {
  return dara.Validate(s)
}

