// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSSLConnectionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v bool) *EnableSSLConnectionResponseBody
  GetData() *bool 
  SetErrCode(v string) *EnableSSLConnectionResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *EnableSSLConnectionResponseBody
  GetErrMessage() *string 
  SetHttpStatusCode(v int32) *EnableSSLConnectionResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *EnableSSLConnectionResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableSSLConnectionResponseBody
  GetSuccess() *bool 
}

type EnableSSLConnectionResponseBody struct {
  // example:
  // 
  // 24151320976****
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // InvalidParams
  ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
  // example:
  // 
  // Invalid params: [instance not exists].
  ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // 32A44F0D-BFF6-5664-999A-218BBDE74XXX
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // True
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableSSLConnectionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSSLConnectionResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSSLConnectionResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnableSSLConnectionResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *EnableSSLConnectionResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *EnableSSLConnectionResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EnableSSLConnectionResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSSLConnectionResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableSSLConnectionResponseBody) SetData(v bool) *EnableSSLConnectionResponseBody {
  s.Data = &v
  return s
}

func (s *EnableSSLConnectionResponseBody) SetErrCode(v string) *EnableSSLConnectionResponseBody {
  s.ErrCode = &v
  return s
}

func (s *EnableSSLConnectionResponseBody) SetErrMessage(v string) *EnableSSLConnectionResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *EnableSSLConnectionResponseBody) SetHttpStatusCode(v int32) *EnableSSLConnectionResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableSSLConnectionResponseBody) SetRequestId(v string) *EnableSSLConnectionResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSSLConnectionResponseBody) SetSuccess(v bool) *EnableSSLConnectionResponseBody {
  s.Success = &v
  return s
}

func (s *EnableSSLConnectionResponseBody) Validate() error {
  return dara.Validate(s)
}

