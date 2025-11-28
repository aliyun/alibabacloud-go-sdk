// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDeviceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableDeviceResponseBody
  GetCode() *string 
  SetHttpStatusCode(v int32) *EnableDeviceResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *EnableDeviceResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableDeviceResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableDeviceResponseBody
  GetSuccess() *bool 
}

type EnableDeviceResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableDeviceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableDeviceResponseBody) GoString() string {
  return s.String()
}

func (s *EnableDeviceResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableDeviceResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EnableDeviceResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableDeviceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableDeviceResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableDeviceResponseBody) SetCode(v string) *EnableDeviceResponseBody {
  s.Code = &v
  return s
}

func (s *EnableDeviceResponseBody) SetHttpStatusCode(v int32) *EnableDeviceResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableDeviceResponseBody) SetMessage(v string) *EnableDeviceResponseBody {
  s.Message = &v
  return s
}

func (s *EnableDeviceResponseBody) SetRequestId(v string) *EnableDeviceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableDeviceResponseBody) SetSuccess(v bool) *EnableDeviceResponseBody {
  s.Success = &v
  return s
}

func (s *EnableDeviceResponseBody) Validate() error {
  return dara.Validate(s)
}

