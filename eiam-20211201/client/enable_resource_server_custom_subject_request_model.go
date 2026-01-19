// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableResourceServerCustomSubjectRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApplicationId(v string) *EnableResourceServerCustomSubjectRequest
  GetApplicationId() *string 
  SetInstanceId(v string) *EnableResourceServerCustomSubjectRequest
  GetInstanceId() *string 
}

type EnableResourceServerCustomSubjectRequest struct {
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

func (s EnableResourceServerCustomSubjectRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableResourceServerCustomSubjectRequest) GoString() string {
  return s.String()
}

func (s *EnableResourceServerCustomSubjectRequest) GetApplicationId() *string  {
  return s.ApplicationId
}

func (s *EnableResourceServerCustomSubjectRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableResourceServerCustomSubjectRequest) SetApplicationId(v string) *EnableResourceServerCustomSubjectRequest {
  s.ApplicationId = &v
  return s
}

func (s *EnableResourceServerCustomSubjectRequest) SetInstanceId(v string) *EnableResourceServerCustomSubjectRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableResourceServerCustomSubjectRequest) Validate() error {
  return dara.Validate(s)
}

