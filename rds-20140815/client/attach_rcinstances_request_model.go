// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachRCInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *AttachRCInstancesRequest
	GetInstanceIds() []*string
	SetKeyPair(v string) *AttachRCInstancesRequest
	GetKeyPair() *string
	SetPassword(v string) *AttachRCInstancesRequest
	GetPassword() *string
	SetRegionId(v string) *AttachRCInstancesRequest
	GetRegionId() *string
	SetVpcId(v string) *AttachRCInstancesRequest
	GetVpcId() *string
}

type AttachRCInstancesRequest struct {
	// The node IDs.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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

func (s AttachRCInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachRCInstancesRequest) GoString() string {
	return s.String()
}

func (s *AttachRCInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *AttachRCInstancesRequest) GetKeyPair() *string {
	return s.KeyPair
}

func (s *AttachRCInstancesRequest) GetPassword() *string {
	return s.Password
}

func (s *AttachRCInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachRCInstancesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *AttachRCInstancesRequest) SetInstanceIds(v []*string) *AttachRCInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *AttachRCInstancesRequest) SetKeyPair(v string) *AttachRCInstancesRequest {
	s.KeyPair = &v
	return s
}

func (s *AttachRCInstancesRequest) SetPassword(v string) *AttachRCInstancesRequest {
	s.Password = &v
	return s
}

func (s *AttachRCInstancesRequest) SetRegionId(v string) *AttachRCInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *AttachRCInstancesRequest) SetVpcId(v string) *AttachRCInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *AttachRCInstancesRequest) Validate() error {
	return dara.Validate(s)
}
