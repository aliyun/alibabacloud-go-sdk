// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyAccessTimeRangeConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTimeRangeConfigShrink(v string) *SetPolicyAccessTimeRangeConfigShrinkRequest
	GetAccessTimeRangeConfigShrink() *string
	SetInstanceId(v string) *SetPolicyAccessTimeRangeConfigShrinkRequest
	GetInstanceId() *string
	SetPolicyId(v string) *SetPolicyAccessTimeRangeConfigShrinkRequest
	GetPolicyId() *string
	SetRegionId(v string) *SetPolicyAccessTimeRangeConfigShrinkRequest
	GetRegionId() *string
}

type SetPolicyAccessTimeRangeConfigShrinkRequest struct {
	// The logon period limits.
	//
	// This parameter is required.
	AccessTimeRangeConfigShrink *string `json:"AccessTimeRangeConfig,omitempty" xml:"AccessTimeRangeConfig,omitempty"`
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
	// The control policy ID.
	//
	// >  You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to query the control policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetPolicyAccessTimeRangeConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyAccessTimeRangeConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetPolicyAccessTimeRangeConfigShrinkRequest) GetAccessTimeRangeConfigShrink() *string {
	return s.AccessTimeRangeConfigShrink
}

func (s *SetPolicyAccessTimeRangeConfigShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPolicyAccessTimeRangeConfigShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *SetPolicyAccessTimeRangeConfigShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetPolicyAccessTimeRangeConfigShrinkRequest) SetAccessTimeRangeConfigShrink(v string) *SetPolicyAccessTimeRangeConfigShrinkRequest {
	s.AccessTimeRangeConfigShrink = &v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigShrinkRequest) SetInstanceId(v string) *SetPolicyAccessTimeRangeConfigShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigShrinkRequest) SetPolicyId(v string) *SetPolicyAccessTimeRangeConfigShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigShrinkRequest) SetRegionId(v string) *SetPolicyAccessTimeRangeConfigShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *SetPolicyAccessTimeRangeConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
