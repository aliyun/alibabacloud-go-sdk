// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInternalSlbResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v bool) *EnableInternalSlbResponseBody
  GetData() *bool 
  SetErrCode(v string) *EnableInternalSlbResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *EnableInternalSlbResponseBody
  GetErrMessage() *string 
  SetHttpStatusCode(v int32) *EnableInternalSlbResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *EnableInternalSlbResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableInternalSlbResponseBody
  GetSuccess() *bool 
}

type EnableInternalSlbResponseBody struct {
  // example:
  // 
  // true
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
  // false
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableInternalSlbResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableInternalSlbResponseBody) GoString() string {
  return s.String()
}

func (s *EnableInternalSlbResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnableInternalSlbResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *EnableInternalSlbResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *EnableInternalSlbResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EnableInternalSlbResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableInternalSlbResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableInternalSlbResponseBody) SetData(v bool) *EnableInternalSlbResponseBody {
  s.Data = &v
  return s
}

func (s *EnableInternalSlbResponseBody) SetErrCode(v string) *EnableInternalSlbResponseBody {
  s.ErrCode = &v
  return s
}

func (s *EnableInternalSlbResponseBody) SetErrMessage(v string) *EnableInternalSlbResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *EnableInternalSlbResponseBody) SetHttpStatusCode(v int32) *EnableInternalSlbResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableInternalSlbResponseBody) SetRequestId(v string) *EnableInternalSlbResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableInternalSlbResponseBody) SetSuccess(v bool) *EnableInternalSlbResponseBody {
  s.Success = &v
  return s
}

func (s *EnableInternalSlbResponseBody) Validate() error {
  return dara.Validate(s)
}

