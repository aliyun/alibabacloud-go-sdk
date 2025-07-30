// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationProvisioningRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *EnableApplicationProvisioningRequest
  GetApplicationId() *string 
  SetInstanceId(v string) *EnableApplicationProvisioningRequest
  GetInstanceId() *string 
}

type EnableApplicationProvisioningRequest struct {
  // The ID of the application.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // app_mkv7rgt4d7i4u7zqtzev2mxxxx
  ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
  // The ID of the instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableApplicationProvisioningRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationProvisioningRequest) GoString() string {
  return s.String()
}

func (s *EnableApplicationProvisioningRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnableApplicationProvisioningRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableApplicationProvisioningRequest) SetApplicationId(v string) *EnableApplicationProvisioningRequest {
  s.ApplicationId = &v
  return s
}

func (s *EnableApplicationProvisioningRequest) SetInstanceId(v string) *EnableApplicationProvisioningRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableApplicationProvisioningRequest) Validate() error {
  return dara.Validate(s)
}

