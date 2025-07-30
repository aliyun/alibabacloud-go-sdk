// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationApiInvokeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *EnableApplicationApiInvokeRequest
  GetApplicationId() *string 
  SetInstanceId(v string) *EnableApplicationApiInvokeRequest
  GetInstanceId() *string 
}

type EnableApplicationApiInvokeRequest struct {
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

func (s EnableApplicationApiInvokeRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationApiInvokeRequest) GoString() string {
  return s.String()
}

func (s *EnableApplicationApiInvokeRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnableApplicationApiInvokeRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableApplicationApiInvokeRequest) SetApplicationId(v string) *EnableApplicationApiInvokeRequest {
  s.ApplicationId = &v
  return s
}

func (s *EnableApplicationApiInvokeRequest) SetInstanceId(v string) *EnableApplicationApiInvokeRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableApplicationApiInvokeRequest) Validate() error {
  return dara.Validate(s)
}

