// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableGbSubDeviceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDeviceName(v string) *EnableGbSubDeviceRequest
  GetDeviceName() *string 
  SetIotId(v string) *EnableGbSubDeviceRequest
  GetIotId() *string 
  SetIotInstanceId(v string) *EnableGbSubDeviceRequest
  GetIotInstanceId() *string 
  SetProductKey(v string) *EnableGbSubDeviceRequest
  GetProductKey() *string 
  SetSubDeviceId(v string) *EnableGbSubDeviceRequest
  GetSubDeviceId() *string 
}

type EnableGbSubDeviceRequest struct {
  DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
  IotId *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
  // This parameter is required.
  IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
  ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
  // This parameter is required.
  SubDeviceId *string `json:"SubDeviceId,omitempty" xml:"SubDeviceId,omitempty"`
}

func (s EnableGbSubDeviceRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableGbSubDeviceRequest) GoString() string {
  return s.String()
}

func (s *EnableGbSubDeviceRequest) GetDeviceName() *string  {
  return s.DeviceName
}

func (s *EnableGbSubDeviceRequest) GetIotId() *string  {
  return s.IotId
}

func (s *EnableGbSubDeviceRequest) GetIotInstanceId() *string  {
  return s.IotInstanceId
}

func (s *EnableGbSubDeviceRequest) GetProductKey() *string  {
  return s.ProductKey
}

func (s *EnableGbSubDeviceRequest) GetSubDeviceId() *string  {
  return s.SubDeviceId
}

func (s *EnableGbSubDeviceRequest) SetDeviceName(v string) *EnableGbSubDeviceRequest {
  s.DeviceName = &v
  return s
}

func (s *EnableGbSubDeviceRequest) SetIotId(v string) *EnableGbSubDeviceRequest {
  s.IotId = &v
  return s
}

func (s *EnableGbSubDeviceRequest) SetIotInstanceId(v string) *EnableGbSubDeviceRequest {
  s.IotInstanceId = &v
  return s
}

func (s *EnableGbSubDeviceRequest) SetProductKey(v string) *EnableGbSubDeviceRequest {
  s.ProductKey = &v
  return s
}

func (s *EnableGbSubDeviceRequest) SetSubDeviceId(v string) *EnableGbSubDeviceRequest {
  s.SubDeviceId = &v
  return s
}

func (s *EnableGbSubDeviceRequest) Validate() error {
  return dara.Validate(s)
}

