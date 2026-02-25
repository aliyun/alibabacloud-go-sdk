// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeNodeNumber(v int32) *UpdateGatewayRequest
	GetFeNodeNumber() *int32
	SetGatewayId(v string) *UpdateGatewayRequest
	GetGatewayId() *string
	SetGatewayName(v string) *UpdateGatewayRequest
	GetGatewayName() *string
	SetInstanceId(v string) *UpdateGatewayRequest
	GetInstanceId() *string
}

type UpdateGatewayRequest struct {
	// example:
	//
	// 3
	FeNodeNumber *int32 `json:"FeNodeNumber,omitempty" xml:"FeNodeNumber,omitempty"`
	// example:
	//
	// dg-65u7d65p5960fjq7
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// polar-byd-glm-47
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRequest) GetFeNodeNumber() *int32 {
	return s.FeNodeNumber
}

func (s *UpdateGatewayRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *UpdateGatewayRequest) GetGatewayName() *string {
	return s.GatewayName
}

func (s *UpdateGatewayRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateGatewayRequest) SetFeNodeNumber(v int32) *UpdateGatewayRequest {
	s.FeNodeNumber = &v
	return s
}

func (s *UpdateGatewayRequest) SetGatewayId(v string) *UpdateGatewayRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayRequest) SetGatewayName(v string) *UpdateGatewayRequest {
	s.GatewayName = &v
	return s
}

func (s *UpdateGatewayRequest) SetInstanceId(v string) *UpdateGatewayRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateGatewayRequest) Validate() error {
	return dara.Validate(s)
}
