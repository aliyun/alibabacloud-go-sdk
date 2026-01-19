// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationM2MClientRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *EnableApplicationM2MClientRequest
  GetApplicationId() *string 
  SetInstanceId(v string) *EnableApplicationM2MClientRequest
  GetInstanceId() *string 
}

type EnableApplicationM2MClientRequest struct {
  // IDaaS的应用资源ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // app_mkv7rgt4d7i4u7zqtzev2mxxxx
  ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
  // IDaaS EIAM实例的ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableApplicationM2MClientRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationM2MClientRequest) GoString() string {
  return s.String()
}

func (s *EnableApplicationM2MClientRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnableApplicationM2MClientRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableApplicationM2MClientRequest) SetApplicationId(v string) *EnableApplicationM2MClientRequest {
  s.ApplicationId = &v
  return s
}

func (s *EnableApplicationM2MClientRequest) SetInstanceId(v string) *EnableApplicationM2MClientRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableApplicationM2MClientRequest) Validate() error {
  return dara.Validate(s)
}

