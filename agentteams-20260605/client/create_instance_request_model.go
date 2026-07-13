// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateInstanceRequest
	GetClientToken() *string
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetInstanceSpec(v string) *CreateInstanceRequest
	GetInstanceSpec() *string
	SetNetworkType(v string) *CreateInstanceRequest
	GetNetworkType() *string
	SetPaymentType(v string) *CreateInstanceRequest
	GetPaymentType() *string
	SetVpcId(v string) *CreateInstanceRequest
	GetVpcId() *string
	SetZones(v []*CreateInstanceRequestZones) *CreateInstanceRequest
	GetZones() []*CreateInstanceRequestZones
}

type CreateInstanceRequest struct {
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
	VpcId *string                       `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Zones []*CreateInstanceRequestZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *CreateInstanceRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateInstanceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateInstanceRequest) GetZones() []*CreateInstanceRequestZones {
	return s.Zones
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceSpec(v string) *CreateInstanceRequest {
	s.InstanceSpec = &v
	return s
}

func (s *CreateInstanceRequest) SetNetworkType(v string) *CreateInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateInstanceRequest) SetPaymentType(v string) *CreateInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInstanceRequest) SetVpcId(v string) *CreateInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceRequest) SetZones(v []*CreateInstanceRequestZones) *CreateInstanceRequest {
	s.Zones = v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	if s.Zones != nil {
		for _, item := range s.Zones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateInstanceRequestZones struct {
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateInstanceRequestZones) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestZones) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestZones) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateInstanceRequestZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceRequestZones) SetVSwitchId(v string) *CreateInstanceRequestZones {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequestZones) SetZoneId(v string) *CreateInstanceRequestZones {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceRequestZones) Validate() error {
	return dara.Validate(s)
}
