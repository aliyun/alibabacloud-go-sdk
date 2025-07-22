// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNatGatewayId(v string) *DeleteNatGatewayRequest
	GetNatGatewayId() *string
	SetRegionId(v string) *DeleteNatGatewayRequest
	GetRegionId() *string
}

type DeleteNatGatewayRequest struct {
	// This parameter is required.
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteNatGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteNatGatewayRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DeleteNatGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNatGatewayRequest) SetNatGatewayId(v string) *DeleteNatGatewayRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DeleteNatGatewayRequest) SetRegionId(v string) *DeleteNatGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNatGatewayRequest) Validate() error {
	return dara.Validate(s)
}
