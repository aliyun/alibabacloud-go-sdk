// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyProtocolConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SetPolicyProtocolConfigShrinkRequest
	GetInstanceId() *string
	SetPolicyId(v string) *SetPolicyProtocolConfigShrinkRequest
	GetPolicyId() *string
	SetProtocolConfigShrink(v string) *SetPolicyProtocolConfigShrinkRequest
	GetProtocolConfigShrink() *string
	SetRegionId(v string) *SetPolicyProtocolConfigShrinkRequest
	GetRegionId() *string
}

type SetPolicyProtocolConfigShrinkRequest struct {
	// The bastion host ID.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy that you want to modify.
	//
	// > You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to query the control policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 61
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The protocol control settings.
	//
	// This parameter is required.
	ProtocolConfigShrink *string `json:"ProtocolConfig,omitempty" xml:"ProtocolConfig,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetPolicyProtocolConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyProtocolConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetPolicyProtocolConfigShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPolicyProtocolConfigShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *SetPolicyProtocolConfigShrinkRequest) GetProtocolConfigShrink() *string {
	return s.ProtocolConfigShrink
}

func (s *SetPolicyProtocolConfigShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetPolicyProtocolConfigShrinkRequest) SetInstanceId(v string) *SetPolicyProtocolConfigShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPolicyProtocolConfigShrinkRequest) SetPolicyId(v string) *SetPolicyProtocolConfigShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *SetPolicyProtocolConfigShrinkRequest) SetProtocolConfigShrink(v string) *SetPolicyProtocolConfigShrinkRequest {
	s.ProtocolConfigShrink = &v
	return s
}

func (s *SetPolicyProtocolConfigShrinkRequest) SetRegionId(v string) *SetPolicyProtocolConfigShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *SetPolicyProtocolConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
