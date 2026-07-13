// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateInstanceShrinkRequest
	GetClientToken() *string
	SetInstanceName(v string) *CreateInstanceShrinkRequest
	GetInstanceName() *string
	SetInstanceSpec(v string) *CreateInstanceShrinkRequest
	GetInstanceSpec() *string
	SetNetworkType(v string) *CreateInstanceShrinkRequest
	GetNetworkType() *string
	SetPaymentType(v string) *CreateInstanceShrinkRequest
	GetPaymentType() *string
	SetVpcId(v string) *CreateInstanceShrinkRequest
	GetVpcId() *string
	SetZonesShrink(v string) *CreateInstanceShrinkRequest
	GetZonesShrink() *string
}

type CreateInstanceShrinkRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agentteams-demo
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SMALL_X1
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PRIVATE_NET
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1xxxx
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZonesShrink *string `json:"Zones,omitempty" xml:"Zones,omitempty"`
}

func (s CreateInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateInstanceShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceShrinkRequest) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *CreateInstanceShrinkRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateInstanceShrinkRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateInstanceShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateInstanceShrinkRequest) GetZonesShrink() *string {
	return s.ZonesShrink
}

func (s *CreateInstanceShrinkRequest) SetClientToken(v string) *CreateInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetInstanceName(v string) *CreateInstanceShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetInstanceSpec(v string) *CreateInstanceShrinkRequest {
	s.InstanceSpec = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetNetworkType(v string) *CreateInstanceShrinkRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetPaymentType(v string) *CreateInstanceShrinkRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetVpcId(v string) *CreateInstanceShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetZonesShrink(v string) *CreateInstanceShrinkRequest {
	s.ZonesShrink = &v
	return s
}

func (s *CreateInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
