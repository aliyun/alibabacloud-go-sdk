// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeNodeNumber(v int32) *AddGatewayRequest
	GetFeNodeNumber() *int32
	SetGatewayName(v string) *AddGatewayRequest
	GetGatewayName() *string
	SetInstanceId(v string) *AddGatewayRequest
	GetInstanceId() *string
}

type AddGatewayRequest struct {
	// example:
	//
	// 3
	FeNodeNumber *int32 `json:"FeNodeNumber,omitempty" xml:"FeNodeNumber,omitempty"`
	// example:
	//
	// eas_automation_test
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AddGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRequest) GoString() string {
	return s.String()
}

func (s *AddGatewayRequest) GetFeNodeNumber() *int32 {
	return s.FeNodeNumber
}

func (s *AddGatewayRequest) GetGatewayName() *string {
	return s.GatewayName
}

func (s *AddGatewayRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddGatewayRequest) SetFeNodeNumber(v int32) *AddGatewayRequest {
	s.FeNodeNumber = &v
	return s
}

func (s *AddGatewayRequest) SetGatewayName(v string) *AddGatewayRequest {
	s.GatewayName = &v
	return s
}

func (s *AddGatewayRequest) SetInstanceId(v string) *AddGatewayRequest {
	s.InstanceId = &v
	return s
}

func (s *AddGatewayRequest) Validate() error {
	return dara.Validate(s)
}
