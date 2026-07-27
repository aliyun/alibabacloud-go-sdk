// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAgentRuntimeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBranchName(v string) *EnableAgentRuntimeRequest
  GetBranchName() *string 
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
  BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
  // The idempotency parameter.
  // 
  // example:
  // 
  // 0c593ea1-3bea-11e9-b96b-88e9fe637760
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // The instance ID of the AI application.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ra-supabase-8moov5lxba****
  InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
  // The region ID of the instance.
  // 
  // example:
  // 
  // cn-beijing
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The security group ID used to create an endpoint.
  // 
  // **If not specified**: The system performs automatic creation of a security group named **sg-aliyun-rds-created-supabase-sandbox*	- in the VPC where the instance resides. No manual operation is required.
  // 
  // **If specified**: Make sure that the specified security group allows the CIDR block of the VPC where the Supabase instance resides (both inbound and outbound directions must be allowed). Otherwise, network connectivity issues may occur.
  // 
  // 	Notice: The endpoint is created only once. When the first Supabase instance in a VPC enables the sandbox and Edge Routine function, the system performs automatic creation of the endpoint. When other Supabase instances in the same VPC enable this capability later, the existing endpoint is reused and no new endpoint is created.
  // 
  // example:
  // 
  // sg-bp179qkbvlj8ym*****
  SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
  // The vSwitch ID used to create an endpoint. If this parameter is not specified, the vSwitch of the Supabase instance is used by default.
  // 
  // 	Notice: The endpoint is created only once. When the first Supabase instance in a VPC enables the sandbox and Edge Routine function, the system performs automatic creation of the endpoint. When other Supabase instances in the same VPC enable this capability later, the existing endpoint is reused and no new endpoint is created.
  // 
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

func (s *EnableAgentRuntimeRequest) GetBranchName() *string  {
  return s.BranchName
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

func (s *EnableAgentRuntimeRequest) SetBranchName(v string) *EnableAgentRuntimeRequest {
  s.BranchName = &v
  return s
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

