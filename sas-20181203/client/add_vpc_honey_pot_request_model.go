// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVpcHoneyPotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVpcId(v string) *AddVpcHoneyPotRequest
	GetVpcId() *string
}

type AddVpcHoneyPotRequest struct {
	// The ID of the virtual private cloud (VPC) in which you want to create a honeypot.
	//
	// > You can call the [DescribeVpcList](~~DescribeVpcList~~) operation to obtain the VPC ID. The VPC ID is the value of the InstanceId parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-p0w5fgkfsl5a6791q****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s AddVpcHoneyPotRequest) String() string {
	return dara.Prettify(s)
}

func (s AddVpcHoneyPotRequest) GoString() string {
	return s.String()
}

func (s *AddVpcHoneyPotRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *AddVpcHoneyPotRequest) SetVpcId(v string) *AddVpcHoneyPotRequest {
	s.VpcId = &v
	return s
}

func (s *AddVpcHoneyPotRequest) Validate() error {
	return dara.Validate(s)
}
