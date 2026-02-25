// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *DeleteGatewayRequest
	GetGatewayId() *string
	SetInstanceId(v string) *DeleteGatewayRequest
	GetInstanceId() *string
}

type DeleteGatewayRequest struct {
	// example:
	//
	// dg-2r69r8kpmn56k5s3
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *DeleteGatewayRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteGatewayRequest) SetGatewayId(v string) *DeleteGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayRequest) SetInstanceId(v string) *DeleteGatewayRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteGatewayRequest) Validate() error {
	return dara.Validate(s)
}
