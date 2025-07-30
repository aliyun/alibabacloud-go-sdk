// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationSsoRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *EnableApplicationSsoRequest
  GetApplicationId() *string 
  SetInstanceId(v string) *EnableApplicationSsoRequest
  GetInstanceId() *string 
}

type EnableApplicationSsoRequest struct {
  // The application ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // app_mkv7rgt4d7i4u7zqtzev2mxxxx
  ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
  // The instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableApplicationSsoRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationSsoRequest) GoString() string {
  return s.String()
}

func (s *EnableApplicationSsoRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnableApplicationSsoRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableApplicationSsoRequest) SetApplicationId(v string) *EnableApplicationSsoRequest {
  s.ApplicationId = &v
  return s
}

func (s *EnableApplicationSsoRequest) SetInstanceId(v string) *EnableApplicationSsoRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableApplicationSsoRequest) Validate() error {
  return dara.Validate(s)
}

