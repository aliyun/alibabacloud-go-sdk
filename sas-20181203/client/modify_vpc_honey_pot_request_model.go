// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcHoneyPotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHoneyPotAction(v string) *ModifyVpcHoneyPotRequest
	GetHoneyPotAction() *string
	SetVpcId(v string) *ModifyVpcHoneyPotRequest
	GetVpcId() *string
}

type ModifyVpcHoneyPotRequest struct {
	// Specifies whether to enable or disable the honeypot. Valid values:
	//
	// 	- **disable**
	//
	// 	- **enable**
	//
	// This parameter is required.
	//
	// example:
	//
	// disable
	HoneyPotAction *string `json:"HoneyPotAction,omitempty" xml:"HoneyPotAction,omitempty"`
	// The ID of the virtual private cloud (VPC) on which the honeypot is deployed.
	//
	// >  You can call the [DescribeVpcHoneyPotList](~~DescribeVpcHoneyPotList~~) operation to query the IDs of VPCs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-d7o009q63fqy21r8u****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ModifyVpcHoneyPotRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcHoneyPotRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcHoneyPotRequest) GetHoneyPotAction() *string {
	return s.HoneyPotAction
}

func (s *ModifyVpcHoneyPotRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyVpcHoneyPotRequest) SetHoneyPotAction(v string) *ModifyVpcHoneyPotRequest {
	s.HoneyPotAction = &v
	return s
}

func (s *ModifyVpcHoneyPotRequest) SetVpcId(v string) *ModifyVpcHoneyPotRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyVpcHoneyPotRequest) Validate() error {
	return dara.Validate(s)
}
