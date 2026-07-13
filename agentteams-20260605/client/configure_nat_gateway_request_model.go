// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureNatGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ConfigureNatGatewayRequest
	GetClientToken() *string
	SetDescription(v string) *ConfigureNatGatewayRequest
	GetDescription() *string
	SetEipAllocationId(v string) *ConfigureNatGatewayRequest
	GetEipAllocationId() *string
	SetEipBandwidth(v int32) *ConfigureNatGatewayRequest
	GetEipBandwidth() *int32
	SetInstanceId(v string) *ConfigureNatGatewayRequest
	GetInstanceId() *string
	SetNatGatewayInstanceId(v string) *ConfigureNatGatewayRequest
	GetNatGatewayInstanceId() *string
}

type ConfigureNatGatewayRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// create public nat gateway for agentteams instance
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EipAllocationId *string `json:"EipAllocationId,omitempty" xml:"EipAllocationId,omitempty"`
	// example:
	//
	// 5
	EipBandwidth *int32 `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// <instanceId>
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NatGatewayInstanceId *string `json:"NatGatewayInstanceId,omitempty" xml:"NatGatewayInstanceId,omitempty"`
}

func (s ConfigureNatGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigureNatGatewayRequest) GoString() string {
	return s.String()
}

func (s *ConfigureNatGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ConfigureNatGatewayRequest) GetDescription() *string {
	return s.Description
}

func (s *ConfigureNatGatewayRequest) GetEipAllocationId() *string {
	return s.EipAllocationId
}

func (s *ConfigureNatGatewayRequest) GetEipBandwidth() *int32 {
	return s.EipBandwidth
}

func (s *ConfigureNatGatewayRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigureNatGatewayRequest) GetNatGatewayInstanceId() *string {
	return s.NatGatewayInstanceId
}

func (s *ConfigureNatGatewayRequest) SetClientToken(v string) *ConfigureNatGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *ConfigureNatGatewayRequest) SetDescription(v string) *ConfigureNatGatewayRequest {
	s.Description = &v
	return s
}

func (s *ConfigureNatGatewayRequest) SetEipAllocationId(v string) *ConfigureNatGatewayRequest {
	s.EipAllocationId = &v
	return s
}

func (s *ConfigureNatGatewayRequest) SetEipBandwidth(v int32) *ConfigureNatGatewayRequest {
	s.EipBandwidth = &v
	return s
}

func (s *ConfigureNatGatewayRequest) SetInstanceId(v string) *ConfigureNatGatewayRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigureNatGatewayRequest) SetNatGatewayInstanceId(v string) *ConfigureNatGatewayRequest {
	s.NatGatewayInstanceId = &v
	return s
}

func (s *ConfigureNatGatewayRequest) Validate() error {
	return dara.Validate(s)
}
