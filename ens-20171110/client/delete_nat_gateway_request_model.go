// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceDelete(v bool) *DeleteNatGatewayRequest
	GetForceDelete() *bool
	SetNatGatewayId(v string) *DeleteNatGatewayRequest
	GetNatGatewayId() *string
}

type DeleteNatGatewayRequest struct {
	// Specifies whether to forcefully delete the NAT instance.
	//
	// example:
	//
	// true
	ForceDelete *bool `json:"ForceDelete,omitempty" xml:"ForceDelete,omitempty"`
	// The ID of the NAT gateway that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// nat-5t7nh1cfm6kxiszlttr38****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
}

func (s DeleteNatGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteNatGatewayRequest) GetForceDelete() *bool {
	return s.ForceDelete
}

func (s *DeleteNatGatewayRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DeleteNatGatewayRequest) SetForceDelete(v bool) *DeleteNatGatewayRequest {
	s.ForceDelete = &v
	return s
}

func (s *DeleteNatGatewayRequest) SetNatGatewayId(v string) *DeleteNatGatewayRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DeleteNatGatewayRequest) Validate() error {
	return dara.Validate(s)
}
