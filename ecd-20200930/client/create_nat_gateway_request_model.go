// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateNatGatewayRequest
	GetRegionId() *string
	SetVSwitchId(v string) *CreateNatGatewayRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateNatGatewayRequest
	GetVpcId() *string
}

type CreateNatGatewayRequest struct {
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateNatGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNatGatewayRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateNatGatewayRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateNatGatewayRequest) SetRegionId(v string) *CreateNatGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNatGatewayRequest) SetVSwitchId(v string) *CreateNatGatewayRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateNatGatewayRequest) SetVpcId(v string) *CreateNatGatewayRequest {
	s.VpcId = &v
	return s
}

func (s *CreateNatGatewayRequest) Validate() error {
	return dara.Validate(s)
}
