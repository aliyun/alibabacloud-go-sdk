// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationResourceServerRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *EnableApplicationResourceServerRequest
  GetApplicationId() *string 
  SetInstanceId(v string) *EnableApplicationResourceServerRequest
  GetInstanceId() *string 
}

type EnableApplicationResourceServerRequest struct {
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

func (s EnableApplicationResourceServerRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationResourceServerRequest) GoString() string {
  return s.String()
}

func (s *EnableApplicationResourceServerRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnableApplicationResourceServerRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableApplicationResourceServerRequest) SetApplicationId(v string) *EnableApplicationResourceServerRequest {
  s.ApplicationId = &v
  return s
}

func (s *EnableApplicationResourceServerRequest) SetInstanceId(v string) *EnableApplicationResourceServerRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableApplicationResourceServerRequest) Validate() error {
  return dara.Validate(s)
}

