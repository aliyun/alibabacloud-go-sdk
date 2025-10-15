// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationTokenRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *EnableApplicationTokenRequest
  GetApplicationId() *string 
  SetApplicationTokenId(v string) *EnableApplicationTokenRequest
  GetApplicationTokenId() *string 
  SetInstanceId(v string) *EnableApplicationTokenRequest
  GetInstanceId() *string 
}

type EnableApplicationTokenRequest struct {
  // IDaaS的应用资源ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // app_mkv7rgt4d7i4u7zqtzev2mxxxx
  ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
  // IDaaS的应用资源TokenID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // token_sfrwerxxxxxxxxxxxxxx
  ApplicationTokenId *string `json:"ApplicationTokenId,omitempty" xml:"ApplicationTokenId,omitempty"`
  // IDaaS EIAM实例的ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableApplicationTokenRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationTokenRequest) GoString() string {
  return s.String()
}

func (s *EnableApplicationTokenRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnableApplicationTokenRequest) GetApplicationTokenId() *string  {
  return s.ApplicationTokenId
}

func (s *EnableApplicationTokenRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableApplicationTokenRequest) SetApplicationId(v string) *EnableApplicationTokenRequest {
  s.ApplicationId = &v
  return s
}

func (s *EnableApplicationTokenRequest) SetApplicationTokenId(v string) *EnableApplicationTokenRequest {
  s.ApplicationTokenId = &v
  return s
}

func (s *EnableApplicationTokenRequest) SetInstanceId(v string) *EnableApplicationTokenRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableApplicationTokenRequest) Validate() error {
  return dara.Validate(s)
}

