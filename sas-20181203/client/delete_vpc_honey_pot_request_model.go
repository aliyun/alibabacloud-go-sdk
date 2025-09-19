// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcHoneyPotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVpcId(v string) *DeleteVpcHoneyPotRequest
	GetVpcId() *string
}

type DeleteVpcHoneyPotRequest struct {
	// The ID of the virtual private cloud (VPC) on which the honeypot is deployed.
	//
	// > You can call the [DescribeVpcHoneyPotList](~~DescribeVpcHoneyPotList~~) operation to query the IDs of VPCs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-d7o009q63fqy21r8u****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeleteVpcHoneyPotRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcHoneyPotRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcHoneyPotRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DeleteVpcHoneyPotRequest) SetVpcId(v string) *DeleteVpcHoneyPotRequest {
	s.VpcId = &v
	return s
}

func (s *DeleteVpcHoneyPotRequest) Validate() error {
	return dara.Validate(s)
}
