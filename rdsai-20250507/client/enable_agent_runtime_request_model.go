// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAgentRuntimeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableAgentRuntimeRequest
  GetClientToken() *string 
  SetInstanceName(v string) *EnableAgentRuntimeRequest
  GetInstanceName() *string 
  SetRegionId(v string) *EnableAgentRuntimeRequest
  GetRegionId() *string 
  SetSecurityGroupId(v string) *EnableAgentRuntimeRequest
  GetSecurityGroupId() *string 
  SetVSwitchId(v string) *EnableAgentRuntimeRequest
  GetVSwitchId() *string 
}

type EnableAgentRuntimeRequest struct {
  // example:
  // 
  // 0c593ea1-3bea-11e9-b96b-88e9fe637760
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // ra-supabase-8moov5lxba****
  InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
  // example:
  // 
  // cn-beijing
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // example:
  // 
  // sg-bp179qkbvlj8ym*****
  SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
  // example:
  // 
  // vsw-9dp2hkpm22gxscfgy****
  VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s EnableAgentRuntimeRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableAgentRuntimeRequest) GoString() string {
  return s.String()
}

func (s *EnableAgentRuntimeRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableAgentRuntimeRequest) GetInstanceName() *string  {
  return s.InstanceName
}

func (s *EnableAgentRuntimeRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableAgentRuntimeRequest) GetSecurityGroupId() *string  {
  return s.SecurityGroupId
}

func (s *EnableAgentRuntimeRequest) GetVSwitchId() *string  {
  return s.VSwitchId
}

func (s *EnableAgentRuntimeRequest) SetClientToken(v string) *EnableAgentRuntimeRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableAgentRuntimeRequest) SetInstanceName(v string) *EnableAgentRuntimeRequest {
  s.InstanceName = &v
  return s
}

func (s *EnableAgentRuntimeRequest) SetRegionId(v string) *EnableAgentRuntimeRequest {
  s.RegionId = &v
  return s
}

func (s *EnableAgentRuntimeRequest) SetSecurityGroupId(v string) *EnableAgentRuntimeRequest {
  s.SecurityGroupId = &v
  return s
}

func (s *EnableAgentRuntimeRequest) SetVSwitchId(v string) *EnableAgentRuntimeRequest {
  s.VSwitchId = &v
  return s
}

func (s *EnableAgentRuntimeRequest) Validate() error {
  return dara.Validate(s)
}

