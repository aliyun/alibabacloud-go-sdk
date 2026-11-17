// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectSupplyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *GetDirectSupplyRequest
	GetGatewayId() *string
}

type GetDirectSupplyRequest struct {
	// The ID of the gateway instance for which you want to query direct-supply models.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-xxxxxxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
}

func (s GetDirectSupplyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDirectSupplyRequest) GoString() string {
	return s.String()
}

func (s *GetDirectSupplyRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetDirectSupplyRequest) SetGatewayId(v string) *GetDirectSupplyRequest {
	s.GatewayId = &v
	return s
}

func (s *GetDirectSupplyRequest) Validate() error {
	return dara.Validate(s)
}
