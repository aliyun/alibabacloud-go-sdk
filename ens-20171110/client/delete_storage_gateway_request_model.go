// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStorageGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *DeleteStorageGatewayRequest
	GetGatewayId() *string
}

type DeleteStorageGatewayRequest struct {
	// The ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgw-****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
}

func (s DeleteStorageGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStorageGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteStorageGatewayRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *DeleteStorageGatewayRequest) SetGatewayId(v string) *DeleteStorageGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *DeleteStorageGatewayRequest) Validate() error {
	return dara.Validate(s)
}
