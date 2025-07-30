// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *EnableApplicationRequest
  GetApplicationId() *string 
  SetInstanceId(v string) *EnableApplicationRequest
  GetInstanceId() *string 
}

type EnableApplicationRequest struct {
  // The ID of the application that you want to enable.
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

func (s EnableApplicationRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationRequest) GoString() string {
  return s.String()
}

func (s *EnableApplicationRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnableApplicationRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableApplicationRequest) SetApplicationId(v string) *EnableApplicationRequest {
  s.ApplicationId = &v
  return s
}

func (s *EnableApplicationRequest) SetInstanceId(v string) *EnableApplicationRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableApplicationRequest) Validate() error {
  return dara.Validate(s)
}

