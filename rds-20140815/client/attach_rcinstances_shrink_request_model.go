// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachRCInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *AttachRCInstancesShrinkRequest
	GetInstanceIdsShrink() *string
	SetKeyPair(v string) *AttachRCInstancesShrinkRequest
	GetKeyPair() *string
	SetPassword(v string) *AttachRCInstancesShrinkRequest
	GetPassword() *string
	SetRegionId(v string) *AttachRCInstancesShrinkRequest
	GetRegionId() *string
	SetVpcId(v string) *AttachRCInstancesShrinkRequest
	GetVpcId() *string
}

type AttachRCInstancesShrinkRequest struct {
	// The node IDs.
	//
	// This parameter is required.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The key pair of the node.
	//
	// example:
	//
	// Custom_test
	KeyPair *string `json:"KeyPair,omitempty" xml:"KeyPair,omitempty"`
	// The logon password of the node.
	//
	// example:
	//
	// testPassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// > This is a reserved parameter.
	//
	// example:
	//
	// None
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s AttachRCInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachRCInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachRCInstancesShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *AttachRCInstancesShrinkRequest) GetKeyPair() *string {
	return s.KeyPair
}

func (s *AttachRCInstancesShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *AttachRCInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachRCInstancesShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *AttachRCInstancesShrinkRequest) SetInstanceIdsShrink(v string) *AttachRCInstancesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *AttachRCInstancesShrinkRequest) SetKeyPair(v string) *AttachRCInstancesShrinkRequest {
	s.KeyPair = &v
	return s
}

func (s *AttachRCInstancesShrinkRequest) SetPassword(v string) *AttachRCInstancesShrinkRequest {
	s.Password = &v
	return s
}

func (s *AttachRCInstancesShrinkRequest) SetRegionId(v string) *AttachRCInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *AttachRCInstancesShrinkRequest) SetVpcId(v string) *AttachRCInstancesShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *AttachRCInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
