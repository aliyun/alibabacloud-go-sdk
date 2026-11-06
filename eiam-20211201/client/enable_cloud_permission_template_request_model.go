// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCloudPermissionTemplateRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCloudAccountId(v string) *EnableCloudPermissionTemplateRequest
  GetCloudAccountId() *string 
  SetCloudPermissionTemplateId(v string) *EnableCloudPermissionTemplateRequest
  GetCloudPermissionTemplateId() *string 
  SetInstanceId(v string) *EnableCloudPermissionTemplateRequest
  GetInstanceId() *string 
}

type EnableCloudPermissionTemplateRequest struct {
  // The cloud account ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ac_mkv7rgt4d7i4u7zqtzev2mxxxx
  CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
  // The permission template ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cpt_mkv7rgt4d7i4u7zqtzev2mxxxx
  CloudPermissionTemplateId *string `json:"CloudPermissionTemplateId,omitempty" xml:"CloudPermissionTemplateId,omitempty"`
  // The ID of the IDaaS EIAM instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableCloudPermissionTemplateRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableCloudPermissionTemplateRequest) GoString() string {
  return s.String()
}

func (s *EnableCloudPermissionTemplateRequest) GetCloudAccountId() *string  {
  return s.CloudAccountId
}

func (s *EnableCloudPermissionTemplateRequest) GetCloudPermissionTemplateId() *string  {
  return s.CloudPermissionTemplateId
}

func (s *EnableCloudPermissionTemplateRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableCloudPermissionTemplateRequest) SetCloudAccountId(v string) *EnableCloudPermissionTemplateRequest {
  s.CloudAccountId = &v
  return s
}

func (s *EnableCloudPermissionTemplateRequest) SetCloudPermissionTemplateId(v string) *EnableCloudPermissionTemplateRequest {
  s.CloudPermissionTemplateId = &v
  return s
}

func (s *EnableCloudPermissionTemplateRequest) SetInstanceId(v string) *EnableCloudPermissionTemplateRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableCloudPermissionTemplateRequest) Validate() error {
  return dara.Validate(s)
}

