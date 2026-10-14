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
	// The Elastic IP Address (EIP) that provides public network access.
	//
	// The EIP must not be occupied by an SNAT rule of the same NAT gateway. You can check the IpLists.UsingStatus field returned by DescribeNatGateways to confirm that the EIP is in the Idle state.
	//
	// This parameter is required.
	//
	// example:
	//
	// 121.XXX.XXX.28
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// The external port or port range for port forwarding.
	//
	// - The port range must be within 1 to 65535.
	//
	// - To forward traffic within a port range, separate the start and end ports with a forward slash (/), such as 10/20.
	//
	// - If ExternalPort is set to a port range, InternalPort must also be set to a port range with the same number of ports. For example, if ExternalPort is set to 10/20, InternalPort must be set to 80/90.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The name of the DNAT rule. The name must be 2 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test0
	ForwardEntryName *string `json:"ForwardEntryName,omitempty" xml:"ForwardEntryName,omitempty"`
	// The health check port for DNAT. The port must be within the internal port range. Default value: empty.
	//
	// example:
	//
	// 80
	HealthCheckPort *int32 `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	// The private IP address of the instance that communicates with the Internet through the DNAT entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.XXX.XXX
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// The internal port or port range for port forwarding.
	//
	// - The port range must be within 1 to 65535.
	//
	// - To forward traffic within a port range, separate the start and end ports with a forward slash (/), such as 10/20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	// The protocol type. Valid values:
	//
	// - **TCP**: forwards TCP packets.
	//
	// - **UDP**: forwards UDP packets.
	//
	// - **Any*	- (default): forwards all packets.
	//
	// example:
	//
	// Any
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The ID of the NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// nat-5t7nh1cfm6kxiszlttr38****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The standby EIP that provides public network access. The standby EIP must be bound to the NAT gateway. After the DNAT entry is created, the standby EIP takes effect.
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
