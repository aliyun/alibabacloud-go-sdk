// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEipAffinity(v bool) *CreateSnatEntryRequest
	GetEipAffinity() *bool
	SetIdleTimeout(v int32) *CreateSnatEntryRequest
	GetIdleTimeout() *int32
	SetIspAffinity(v bool) *CreateSnatEntryRequest
	GetIspAffinity() *bool
	SetNatGatewayId(v string) *CreateSnatEntryRequest
	GetNatGatewayId() *string
	SetSnatEntryName(v string) *CreateSnatEntryRequest
	GetSnatEntryName() *string
	SetSnatIp(v string) *CreateSnatEntryRequest
	GetSnatIp() *string
	SetSourceCIDR(v string) *CreateSnatEntryRequest
	GetSourceCIDR() *string
	SetSourceNetworkId(v string) *CreateSnatEntryRequest
	GetSourceNetworkId() *string
	SetSourceVSwitchId(v string) *CreateSnatEntryRequest
	GetSourceVSwitchId() *string
	SetStandbySnatIp(v string) *CreateSnatEntryRequest
	GetStandbySnatIp() *string
}

type CreateSnatEntryRequest struct {
	// Specifies whether to enable IP affinity. If you do not specify this parameter, IP affinity is enabled by default. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// >  After you enable IP affinity, if multiple EIPs are associated with an SNAT entry, one client uses the same EIP to for communication. If IP affinity is disabled, the client uses a random EIP for communication.
	//
	// example:
	//
	// false
	EipAffinity *bool `json:"EipAffinity,omitempty" xml:"EipAffinity,omitempty"`
	// The timeout period for idle connections. Valid values: **1*	- to **86400**. Unit: seconds.
	//
	// example:
	//
	// 15
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// Whether to enable operator affinity. Value taking:
	//
	// - false:Do not open.
	//
	// - true:Open.
	//
	// example:
	//
	// true
	IspAffinity *bool `json:"IspAffinity,omitempty" xml:"IspAffinity,omitempty"`
	// The ID of the Network Address Translation (NAT) gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// nat-5tawjw5j7sgd2deujxuk0****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The name of the SNAT entry. The name must be 1 to 128 characters in length. The name cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test0
	SnatEntryName *string `json:"SnatEntryName,omitempty" xml:"SnatEntryName,omitempty"`
	// The elastic IP address (EIP) in the SNAT entry. Separate multiple EIPs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 120.XXX.XXX.71
	SnatIp *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
	// The CIDR block. You can specify the CIDR block of a network, a vSwitch, or an instance. You can also specify a custom CIDR block. All instances within the CIDR block can access the Internet or external networks by using SNAT.
	//
	// >  If you specify **SourceVSwitchId*	- and **SourceCIDR**, **SourceVSwitchId*	- does not take effect. The value that you specified for **SourceCIDR*	- takes precedence.
	//
	// example:
	//
	// 10.0.0.0/24
	SourceCIDR *string `json:"SourceCIDR,omitempty" xml:"SourceCIDR,omitempty"`
	// The ID of the network. This parameter specifies that all ENS instances in the network can use the SNAT entry to access the Internet.
	//
	// >  If you specify **SourceNetworkId*	- and **SourceVSwitchId*	- or **SourceCIDR**, **SourceNetworkId*	- does not take effect. The value that you specified for **SourceCIDR*	- takes precedence. Priority: **SourceCIDR*	- > **SourceVSwitchId*	- > **SourceNetworkId**.
	//
	// example:
	//
	// n-2zeuphj08tt7q3brd****
	SourceNetworkId *string `json:"SourceNetworkId,omitempty" xml:"SourceNetworkId,omitempty"`
	// The ID of the vSwitch that you need to access over the Internet. This parameter specifies that Edge Node Service (ENS) instances in the vSwitch can use the SNAT entry to access the Internet.
	//
	// >  If you specify **SourceVSwitchId*	- and **SourceCIDR**, **SourceVSwitchId*	- does not take effect. The value that you specified for **SourceCIDR*	- takes precedence.
	//
	// example:
	//
	// vsw-bp1hwx7gi495q260p****
	SourceVSwitchId *string `json:"SourceVSwitchId,omitempty" xml:"SourceVSwitchId,omitempty"`
	// The secondary EIP in the SNAT entry. Separate multiple secondary EIPs with commas (,).
	//
	// example:
	//
	// 101.XXX.XXX.7
	StandbySnatIp *string `json:"StandbySnatIp,omitempty" xml:"StandbySnatIp,omitempty"`
}

func (s CreateSnatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSnatEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateSnatEntryRequest) GetEipAffinity() *bool {
	return s.EipAffinity
}

func (s *CreateSnatEntryRequest) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *CreateSnatEntryRequest) GetIspAffinity() *bool {
	return s.IspAffinity
}

func (s *CreateSnatEntryRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CreateSnatEntryRequest) GetSnatEntryName() *string {
	return s.SnatEntryName
}

func (s *CreateSnatEntryRequest) GetSnatIp() *string {
	return s.SnatIp
}

func (s *CreateSnatEntryRequest) GetSourceCIDR() *string {
	return s.SourceCIDR
}

func (s *CreateSnatEntryRequest) GetSourceNetworkId() *string {
	return s.SourceNetworkId
}

func (s *CreateSnatEntryRequest) GetSourceVSwitchId() *string {
	return s.SourceVSwitchId
}

func (s *CreateSnatEntryRequest) GetStandbySnatIp() *string {
	return s.StandbySnatIp
}

func (s *CreateSnatEntryRequest) SetEipAffinity(v bool) *CreateSnatEntryRequest {
	s.EipAffinity = &v
	return s
}

func (s *CreateSnatEntryRequest) SetIdleTimeout(v int32) *CreateSnatEntryRequest {
	s.IdleTimeout = &v
	return s
}

func (s *CreateSnatEntryRequest) SetIspAffinity(v bool) *CreateSnatEntryRequest {
	s.IspAffinity = &v
	return s
}

func (s *CreateSnatEntryRequest) SetNatGatewayId(v string) *CreateSnatEntryRequest {
	s.NatGatewayId = &v
	return s
}

func (s *CreateSnatEntryRequest) SetSnatEntryName(v string) *CreateSnatEntryRequest {
	s.SnatEntryName = &v
	return s
}

func (s *CreateSnatEntryRequest) SetSnatIp(v string) *CreateSnatEntryRequest {
	s.SnatIp = &v
	return s
}

func (s *CreateSnatEntryRequest) SetSourceCIDR(v string) *CreateSnatEntryRequest {
	s.SourceCIDR = &v
	return s
}

func (s *CreateSnatEntryRequest) SetSourceNetworkId(v string) *CreateSnatEntryRequest {
	s.SourceNetworkId = &v
	return s
}

func (s *CreateSnatEntryRequest) SetSourceVSwitchId(v string) *CreateSnatEntryRequest {
	s.SourceVSwitchId = &v
	return s
}

func (s *CreateSnatEntryRequest) SetStandbySnatIp(v string) *CreateSnatEntryRequest {
	s.StandbySnatIp = &v
	return s
}

func (s *CreateSnatEntryRequest) Validate() error {
	return dara.Validate(s)
}
