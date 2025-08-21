// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForwardEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalIp(v string) *CreateForwardEntryRequest
	GetExternalIp() *string
	SetExternalPort(v string) *CreateForwardEntryRequest
	GetExternalPort() *string
	SetForwardEntryName(v string) *CreateForwardEntryRequest
	GetForwardEntryName() *string
	SetHealthCheckPort(v int32) *CreateForwardEntryRequest
	GetHealthCheckPort() *int32
	SetInternalIp(v string) *CreateForwardEntryRequest
	GetInternalIp() *string
	SetInternalPort(v string) *CreateForwardEntryRequest
	GetInternalPort() *string
	SetIpProtocol(v string) *CreateForwardEntryRequest
	GetIpProtocol() *string
	SetNatGatewayId(v string) *CreateForwardEntryRequest
	GetNatGatewayId() *string
	SetStandbyExternalIp(v string) *CreateForwardEntryRequest
	GetStandbyExternalIp() *string
}

type CreateForwardEntryRequest struct {
	// The elastic IP address (EIP) that is used to access the Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// 121.11.36.28
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// The external port or port range that is used for port forwarding.
	//
	// 	- Valid values: 1 to 65535.
	//
	// 	- To specify a port range, separate the first port and the last port with a forward slash (/), such as 10/20.
	//
	// 	- If you set ExternalPort to a port range, you must also set InternalPort to a port range. The number of ports in the port ranges must be the same. For example, if you set ExternalPort to 10/20, you can set InternalPort to 80/90.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The name of the DNAT entry. The name must be 2 to 128 characters in length. The name cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test0
	ForwardEntryName *string `json:"ForwardEntryName,omitempty" xml:"ForwardEntryName,omitempty"`
	// The probe port. The port must be within the internal port range. By default, this parameter is left empty.
	//
	// example:
	//
	// 80
	HealthCheckPort *int32 `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	// The private IP address of the instance that uses the DNAT entry for Internet communication.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.13
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// The internal port or port range that is used for port forwarding.
	//
	// 	- Valid values: 1 to 65535.
	//
	// 	- To specify a port range, separate the first port and the last port with a forward slash (/), such as 10/20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	// The protocol. Valid values:
	//
	// 	- **TCP**: forwards TCP packets.
	//
	// 	- **UDP**: forwards UDP packets.
	//
	// 	- **Any*	- (default): forwards all packets.
	//
	// example:
	//
	// Any
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The ID of the Network Address Translation (NAT) gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// nat-5t7nh1cfm6kxiszlttr383xpo
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The secondary EIP that is used to access the Internet. You need to select a secondary EIP that is bound to NAT. After the DNAT entry is created, the secondary EIP takes effect.
	//
	// example:
	//
	// 101.XXX.XXX.4
	StandbyExternalIp *string `json:"StandbyExternalIp,omitempty" xml:"StandbyExternalIp,omitempty"`
}

func (s CreateForwardEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateForwardEntryRequest) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *CreateForwardEntryRequest) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *CreateForwardEntryRequest) GetForwardEntryName() *string {
	return s.ForwardEntryName
}

func (s *CreateForwardEntryRequest) GetHealthCheckPort() *int32 {
	return s.HealthCheckPort
}

func (s *CreateForwardEntryRequest) GetInternalIp() *string {
	return s.InternalIp
}

func (s *CreateForwardEntryRequest) GetInternalPort() *string {
	return s.InternalPort
}

func (s *CreateForwardEntryRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *CreateForwardEntryRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CreateForwardEntryRequest) GetStandbyExternalIp() *string {
	return s.StandbyExternalIp
}

func (s *CreateForwardEntryRequest) SetExternalIp(v string) *CreateForwardEntryRequest {
	s.ExternalIp = &v
	return s
}

func (s *CreateForwardEntryRequest) SetExternalPort(v string) *CreateForwardEntryRequest {
	s.ExternalPort = &v
	return s
}

func (s *CreateForwardEntryRequest) SetForwardEntryName(v string) *CreateForwardEntryRequest {
	s.ForwardEntryName = &v
	return s
}

func (s *CreateForwardEntryRequest) SetHealthCheckPort(v int32) *CreateForwardEntryRequest {
	s.HealthCheckPort = &v
	return s
}

func (s *CreateForwardEntryRequest) SetInternalIp(v string) *CreateForwardEntryRequest {
	s.InternalIp = &v
	return s
}

func (s *CreateForwardEntryRequest) SetInternalPort(v string) *CreateForwardEntryRequest {
	s.InternalPort = &v
	return s
}

func (s *CreateForwardEntryRequest) SetIpProtocol(v string) *CreateForwardEntryRequest {
	s.IpProtocol = &v
	return s
}

func (s *CreateForwardEntryRequest) SetNatGatewayId(v string) *CreateForwardEntryRequest {
	s.NatGatewayId = &v
	return s
}

func (s *CreateForwardEntryRequest) SetStandbyExternalIp(v string) *CreateForwardEntryRequest {
	s.StandbyExternalIp = &v
	return s
}

func (s *CreateForwardEntryRequest) Validate() error {
	return dara.Validate(s)
}
