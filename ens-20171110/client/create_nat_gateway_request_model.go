// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *CreateNatGatewayRequest
	GetEnsRegionId() *string
	SetInstanceBillingCycle(v string) *CreateNatGatewayRequest
	GetInstanceBillingCycle() *string
	SetInstanceType(v string) *CreateNatGatewayRequest
	GetInstanceType() *string
	SetName(v string) *CreateNatGatewayRequest
	GetName() *string
	SetNetworkId(v string) *CreateNatGatewayRequest
	GetNetworkId() *string
	SetTag(v []*CreateNatGatewayRequestTag) *CreateNatGatewayRequest
	GetTag() []*CreateNatGatewayRequestTag
	SetVSwitchId(v string) *CreateNatGatewayRequest
	GetVSwitchId() *string
}

type CreateNatGatewayRequest struct {
	// The ID of the ENS edge zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-suzhou-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The billing cycle of the instance. Valid values: Hour and Month.
	//
	// example:
	//
	// Hour
	InstanceBillingCycle *string `json:"InstanceBillingCycle,omitempty" xml:"InstanceBillingCycle,omitempty"`
	// The NAT gateway specification. Valid value: **enat.default**.
	//
	// example:
	//
	// enat.default
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The name of the NAT gateway. The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// n-5qj7ykuxmjn7k96l090sp****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The resource tags.
	Tag []*CreateNatGatewayRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-5savh5ngxh8sbj14bu7n****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateNatGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateNatGatewayRequest) GetInstanceBillingCycle() *string {
	return s.InstanceBillingCycle
}

func (s *CreateNatGatewayRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateNatGatewayRequest) GetName() *string {
	return s.Name
}

func (s *CreateNatGatewayRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *CreateNatGatewayRequest) GetTag() []*CreateNatGatewayRequestTag {
	return s.Tag
}

func (s *CreateNatGatewayRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateNatGatewayRequest) SetEnsRegionId(v string) *CreateNatGatewayRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateNatGatewayRequest) SetInstanceBillingCycle(v string) *CreateNatGatewayRequest {
	s.InstanceBillingCycle = &v
	return s
}

func (s *CreateNatGatewayRequest) SetInstanceType(v string) *CreateNatGatewayRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateNatGatewayRequest) SetName(v string) *CreateNatGatewayRequest {
	s.Name = &v
	return s
}

func (s *CreateNatGatewayRequest) SetNetworkId(v string) *CreateNatGatewayRequest {
	s.NetworkId = &v
	return s
}

func (s *CreateNatGatewayRequest) SetTag(v []*CreateNatGatewayRequestTag) *CreateNatGatewayRequest {
	s.Tag = v
	return s
}

func (s *CreateNatGatewayRequest) SetVSwitchId(v string) *CreateNatGatewayRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateNatGatewayRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNatGatewayRequestTag struct {
	// The tag key of the instance. Valid values of N: **1*	- to **20**.
	//
	// - Cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// - Can be up to 64 characters in length.
	//
	// - Cannot be an empty string.
	//
	// example:
	//
	// team
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length, cannot start with acs:, and cannot contain http:// or https://.
	//
	// example:
	//
	// Deep
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNatGatewayRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayRequestTag) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateNatGatewayRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateNatGatewayRequestTag) SetKey(v string) *CreateNatGatewayRequestTag {
	s.Key = &v
	return s
}

func (s *CreateNatGatewayRequestTag) SetValue(v string) *CreateNatGatewayRequestTag {
	s.Value = &v
	return s
}

func (s *CreateNatGatewayRequestTag) Validate() error {
	return dara.Validate(s)
}
