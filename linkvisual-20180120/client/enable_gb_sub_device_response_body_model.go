// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableGbSubDeviceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableGbSubDeviceResponseBody
  GetCode() *string 
  SetData(v *EnableGbSubDeviceResponseBodyData) *EnableGbSubDeviceResponseBody
  GetData() *EnableGbSubDeviceResponseBodyData 
  SetErrorMessage(v string) *EnableGbSubDeviceResponseBody
  GetErrorMessage() *string 
  SetRequestId(v string) *EnableGbSubDeviceResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableGbSubDeviceResponseBody
  GetSuccess() *bool 
}

type EnableGbSubDeviceResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *EnableGbSubDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableGbSubDeviceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableGbSubDeviceResponseBody) GoString() string {
  return s.String()
}

func (s *EnableGbSubDeviceResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableGbSubDeviceResponseBody) GetData() *EnableGbSubDeviceResponseBodyData  {
  return s.Data
}

func (s *EnableGbSubDeviceResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *EnableGbSubDeviceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableGbSubDeviceResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableGbSubDeviceResponseBody) SetCode(v string) *EnableGbSubDeviceResponseBody {
  s.Code = &v
  return s
}

func (s *EnableGbSubDeviceResponseBody) SetData(v *EnableGbSubDeviceResponseBodyData) *EnableGbSubDeviceResponseBody {
  s.Data = v
  return s
}

func (s *EnableGbSubDeviceResponseBody) SetErrorMessage(v string) *EnableGbSubDeviceResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *EnableGbSubDeviceResponseBody) SetRequestId(v string) *EnableGbSubDeviceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableGbSubDeviceResponseBody) SetSuccess(v bool) *EnableGbSubDeviceResponseBody {
  s.Success = &v
  return s
}

func (s *EnableGbSubDeviceResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnableGbSubDeviceResponseBodyData struct {
  DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
  ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s EnableGbSubDeviceResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnableGbSubDeviceResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnableGbSubDeviceResponseBodyData) GetDeviceName() *string  {
  return s.DeviceName
}

func (s *EnableGbSubDeviceResponseBodyData) GetProductKey() *string  {
  return s.ProductKey
}

func (s *EnableGbSubDeviceResponseBodyData) SetDeviceName(v string) *EnableGbSubDeviceResponseBodyData {
  s.DeviceName = &v
  return s
}

func (s *EnableGbSubDeviceResponseBodyData) SetProductKey(v string) *EnableGbSubDeviceResponseBodyData {
  s.ProductKey = &v
  return s
}

func (s *EnableGbSubDeviceResponseBodyData) Validate() error {
  return dara.Validate(s)
}

