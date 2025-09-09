// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyApprovalConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalConfig(v *SetPolicyApprovalConfigRequestApprovalConfig) *SetPolicyApprovalConfigRequest
	GetApprovalConfig() *SetPolicyApprovalConfigRequestApprovalConfig
	SetInstanceId(v string) *SetPolicyApprovalConfigRequest
	GetInstanceId() *string
	SetPolicyId(v string) *SetPolicyApprovalConfigRequest
	GetPolicyId() *string
	SetRegionId(v string) *SetPolicyApprovalConfigRequest
	GetRegionId() *string
}

type SetPolicyApprovalConfigRequest struct {
	// The O&M approval setting in the control policy.
	//
	// This parameter is required.
	ApprovalConfig *SetPolicyApprovalConfigRequestApprovalConfig `json:"ApprovalConfig,omitempty" xml:"ApprovalConfig,omitempty" type:"Struct"`
	// The bastion host ID.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy that you want to modify.
	//
	// >  You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to query the control policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetPolicyApprovalConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyApprovalConfigRequest) GoString() string {
	return s.String()
}

func (s *SetPolicyApprovalConfigRequest) GetApprovalConfig() *SetPolicyApprovalConfigRequestApprovalConfig {
	return s.ApprovalConfig
}

func (s *SetPolicyApprovalConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPolicyApprovalConfigRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *SetPolicyApprovalConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetPolicyApprovalConfigRequest) SetApprovalConfig(v *SetPolicyApprovalConfigRequestApprovalConfig) *SetPolicyApprovalConfigRequest {
	s.ApprovalConfig = v
	return s
}

func (s *SetPolicyApprovalConfigRequest) SetInstanceId(v string) *SetPolicyApprovalConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPolicyApprovalConfigRequest) SetPolicyId(v string) *SetPolicyApprovalConfigRequest {
	s.PolicyId = &v
	return s
}

func (s *SetPolicyApprovalConfigRequest) SetRegionId(v string) *SetPolicyApprovalConfigRequest {
	s.RegionId = &v
	return s
}

func (s *SetPolicyApprovalConfigRequest) Validate() error {
	return dara.Validate(s)
}

type SetPolicyApprovalConfigRequestApprovalConfig struct {
	// Specifies whether to enable O&M approval in the control policy. Valid values:
	//
	// 	- **On**: enables O&M approval.
	//
	// 	- **Off**: disables O&M approval.
	//
	// This parameter is required.
	//
	// example:
	//
	// On
	SwitchStatus *string `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
}

func (s SetPolicyApprovalConfigRequestApprovalConfig) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyApprovalConfigRequestApprovalConfig) GoString() string {
	return s.String()
}

func (s *SetPolicyApprovalConfigRequestApprovalConfig) GetSwitchStatus() *string {
	return s.SwitchStatus
}

func (s *SetPolicyApprovalConfigRequestApprovalConfig) SetSwitchStatus(v string) *SetPolicyApprovalConfigRequestApprovalConfig {
	s.SwitchStatus = &v
	return s
}

func (s *SetPolicyApprovalConfigRequestApprovalConfig) Validate() error {
	return dara.Validate(s)
}
