// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalErrorSAEWeb interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ExternalErrorSAEWeb
  GetCode() *int32 
  SetErrorCode(v string) *ExternalErrorSAEWeb
  GetErrorCode() *string 
  SetMessage(v string) *ExternalErrorSAEWeb
  GetMessage() *string 
  SetRequestId(v string) *ExternalErrorSAEWeb
  GetRequestId() *string 
  SetSuccess(v bool) *ExternalErrorSAEWeb
  GetSuccess() *bool 
}

type ExternalErrorSAEWeb struct {
  Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
  ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExternalErrorSAEWeb) String() string {
  return dara.Prettify(s)
}

func (s ExternalErrorSAEWeb) GoString() string {
  return s.String()
}

func (s *ExternalErrorSAEWeb) GetCode() *int32  {
  return s.Code
}

func (s *ExternalErrorSAEWeb) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *ExternalErrorSAEWeb) GetMessage() *string  {
  return s.Message
}

func (s *ExternalErrorSAEWeb) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExternalErrorSAEWeb) GetSuccess() *bool  {
  return s.Success
}

func (s *ExternalErrorSAEWeb) SetCode(v int32) *ExternalErrorSAEWeb {
  s.Code = &v
  return s
}

func (s *ExternalErrorSAEWeb) SetErrorCode(v string) *ExternalErrorSAEWeb {
  s.ErrorCode = &v
  return s
}

func (s *ExternalErrorSAEWeb) SetMessage(v string) *ExternalErrorSAEWeb {
  s.Message = &v
  return s
}

func (s *ExternalErrorSAEWeb) SetRequestId(v string) *ExternalErrorSAEWeb {
  s.RequestId = &v
  return s
}

func (s *ExternalErrorSAEWeb) SetSuccess(v bool) *ExternalErrorSAEWeb {
  s.Success = &v
  return s
}

func (s *ExternalErrorSAEWeb) Validate() error {
  return dara.Validate(s)
}

