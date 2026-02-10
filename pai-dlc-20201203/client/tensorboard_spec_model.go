// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTensorboardSpec interface {
	dara.Model
	String() string
	GoString() string
	SetEcsType(v string) *TensorboardSpec
	GetEcsType() *string
	SetSecurityGroupId(v string) *TensorboardSpec
	GetSecurityGroupId() *string
	SetSwitchId(v string) *TensorboardSpec
	GetSwitchId() *string
	SetVpcId(v string) *TensorboardSpec
	GetVpcId() *string
}

type TensorboardSpec struct {
	// The specifications of the ECS instance configured for the pay-as-you-go Tensorboard task.
	//
	// example:
	//
	// ecs.g6.large
	EcsType *string `json:"EcsType,omitempty" xml:"EcsType,omitempty"`
	// The ID of the security group configured for the pay-as-you-go Tensorboard task.
	//
	// example:
	//
	// sg-xxxxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the vSwitch configured for the pay-as-you-go Tensorboard task.
	//
	// example:
	//
	// vsw-xxxx
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) configured for the pay-as-you-go Tensorboard task.
	//
	// example:
	//
	// vpc-xxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s TensorboardSpec) String() string {
	return dara.Prettify(s)
}

func (s TensorboardSpec) GoString() string {
	return s.String()
}

func (s *TensorboardSpec) GetEcsType() *string {
	return s.EcsType
}

func (s *TensorboardSpec) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *TensorboardSpec) GetSwitchId() *string {
	return s.SwitchId
}

func (s *TensorboardSpec) GetVpcId() *string {
	return s.VpcId
}

func (s *TensorboardSpec) SetEcsType(v string) *TensorboardSpec {
	s.EcsType = &v
	return s
}

func (s *TensorboardSpec) SetSecurityGroupId(v string) *TensorboardSpec {
	s.SecurityGroupId = &v
	return s
}

func (s *TensorboardSpec) SetSwitchId(v string) *TensorboardSpec {
	s.SwitchId = &v
	return s
}

func (s *TensorboardSpec) SetVpcId(v string) *TensorboardSpec {
	s.VpcId = &v
	return s
}

func (s *TensorboardSpec) Validate() error {
	return dara.Validate(s)
}
