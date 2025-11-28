// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDeviceGroupResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableDeviceGroupResponseBody
  GetCode() *string 
  SetHttpStatusCode(v int32) *EnableDeviceGroupResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *EnableDeviceGroupResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableDeviceGroupResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableDeviceGroupResponseBody
  GetSuccess() *bool 
}

type EnableDeviceGroupResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableDeviceGroupResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableDeviceGroupResponseBody) GoString() string {
  return s.String()
}

func (s *EnableDeviceGroupResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableDeviceGroupResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EnableDeviceGroupResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableDeviceGroupResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableDeviceGroupResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableDeviceGroupResponseBody) SetCode(v string) *EnableDeviceGroupResponseBody {
  s.Code = &v
  return s
}

func (s *EnableDeviceGroupResponseBody) SetHttpStatusCode(v int32) *EnableDeviceGroupResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableDeviceGroupResponseBody) SetMessage(v string) *EnableDeviceGroupResponseBody {
  s.Message = &v
  return s
}

func (s *EnableDeviceGroupResponseBody) SetRequestId(v string) *EnableDeviceGroupResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableDeviceGroupResponseBody) SetSuccess(v bool) *EnableDeviceGroupResponseBody {
  s.Success = &v
  return s
}

func (s *EnableDeviceGroupResponseBody) Validate() error {
  return dara.Validate(s)
}

