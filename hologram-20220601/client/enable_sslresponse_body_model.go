// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSSLResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v bool) *EnableSSLResponseBody
  GetData() *bool 
  SetErrorCode(v string) *EnableSSLResponseBody
  GetErrorCode() *string 
  SetErrorMessage(v string) *EnableSSLResponseBody
  GetErrorMessage() *string 
  SetHttpStatusCode(v string) *EnableSSLResponseBody
  GetHttpStatusCode() *string 
  SetSuccess(v bool) *EnableSSLResponseBody
  GetSuccess() *bool 
  SetRequestId(v string) *EnableSSLResponseBody
  GetRequestId() *string 
}

type EnableSSLResponseBody struct {
  // example:
  // 
  // true
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // null
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // example:
  // 
  // null
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
  // example:
  // 
  // 819A7F0F-2951-540F-BD94-6A41ECF0281F
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s EnableSSLResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSSLResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSSLResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnableSSLResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EnableSSLResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *EnableSSLResponseBody) GetHttpStatusCode() *string  {
  return s.HttpStatusCode
}

func (s *EnableSSLResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableSSLResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSSLResponseBody) SetData(v bool) *EnableSSLResponseBody {
  s.Data = &v
  return s
}

func (s *EnableSSLResponseBody) SetErrorCode(v string) *EnableSSLResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EnableSSLResponseBody) SetErrorMessage(v string) *EnableSSLResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *EnableSSLResponseBody) SetHttpStatusCode(v string) *EnableSSLResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableSSLResponseBody) SetSuccess(v bool) *EnableSSLResponseBody {
  s.Success = &v
  return s
}

func (s *EnableSSLResponseBody) SetRequestId(v string) *EnableSSLResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSSLResponseBody) Validate() error {
  return dara.Validate(s)
}

