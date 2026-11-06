// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCloudAccountPrivilegeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableCloudAccountPrivilegeRequest
  GetClientToken() *string 
  SetCloudAccountId(v string) *EnableCloudAccountPrivilegeRequest
  GetCloudAccountId() *string 
  SetInstanceId(v string) *EnableCloudAccountPrivilegeRequest
  GetInstanceId() *string 
}

type EnableCloudAccountPrivilegeRequest struct {
  // Ensures the idempotence of the request. You can generate a parameter value on your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters and can be up to 64 characters in length.
  // 
  // example:
  // 
  // client-token-example
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // The ID of the cloud account.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ac_mkv7rgt4d7i4u7zqtzev2mxxxx
  CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
  // The ID of the IDaaS EIAM instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableCloudAccountPrivilegeRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableCloudAccountPrivilegeRequest) GoString() string {
  return s.String()
}

func (s *EnableCloudAccountPrivilegeRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableCloudAccountPrivilegeRequest) GetCloudAccountId() *string  {
  return s.CloudAccountId
}

func (s *EnableCloudAccountPrivilegeRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableCloudAccountPrivilegeRequest) SetClientToken(v string) *EnableCloudAccountPrivilegeRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableCloudAccountPrivilegeRequest) SetCloudAccountId(v string) *EnableCloudAccountPrivilegeRequest {
  s.CloudAccountId = &v
  return s
}

func (s *EnableCloudAccountPrivilegeRequest) SetInstanceId(v string) *EnableCloudAccountPrivilegeRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableCloudAccountPrivilegeRequest) Validate() error {
  return dara.Validate(s)
}

