// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayIntranetLinkedVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVSwitchId(v string) *DeleteGatewayIntranetLinkedVpcRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DeleteGatewayIntranetLinkedVpcRequest
	GetVpcId() *string
}

type DeleteGatewayIntranetLinkedVpcRequest struct {
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-8vbqn2at0kljjxxxx****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeleteGatewayIntranetLinkedVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayIntranetLinkedVpcRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayIntranetLinkedVpcRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DeleteGatewayIntranetLinkedVpcRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DeleteGatewayIntranetLinkedVpcRequest) SetVSwitchId(v string) *DeleteGatewayIntranetLinkedVpcRequest {
	s.VSwitchId = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcRequest) SetVpcId(v string) *DeleteGatewayIntranetLinkedVpcRequest {
	s.VpcId = &v
	return s
}

func (s *DeleteGatewayIntranetLinkedVpcRequest) Validate() error {
	return dara.Validate(s)
}
