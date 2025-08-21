// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyForwardEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalIp(v string) *ModifyForwardEntryRequest
	GetExternalIp() *string
	SetExternalPort(v string) *ModifyForwardEntryRequest
	GetExternalPort() *string
	SetForwardEntryId(v string) *ModifyForwardEntryRequest
	GetForwardEntryId() *string
	SetForwardEntryName(v string) *ModifyForwardEntryRequest
	GetForwardEntryName() *string
	SetHealthCheckPort(v int32) *ModifyForwardEntryRequest
	GetHealthCheckPort() *int32
	SetInternalIp(v string) *ModifyForwardEntryRequest
	GetInternalIp() *string
	SetInternalPort(v string) *ModifyForwardEntryRequest
	GetInternalPort() *string
	SetIpProtocol(v string) *ModifyForwardEntryRequest
	GetIpProtocol() *string
}

type ModifyForwardEntryRequest struct {
	// The EIP in the DNAT entry. The public IP address is used to access the Internet.
	//
	// example:
	//
	// 121.XXX.XXX.28
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// The external port or port range that is used for port forwarding.
	//
	// 	- Valid values: 1 to 65535.
	//
	// 	- To specify a port range, separate the first port and the last port with a forward slash (/), such as 10/20. The first port and the last port are included.
	//
	// 	- If you set ExternalPort to a port range, you must also set InternalPort to a port range. The number of ports in the port ranges must be the same. For example, if you set ExternalPort to 10/20, you can set InternalPort to 80/90.
	//
	// 	- The maximum port range is 1000.
	//
	// example:
	//
	// 22
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The ID of the DNAT entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// dnat-5tfjp3537mi6iokl59g5c****
	ForwardEntryId *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
	// The name of the DNAT entry. The name must be 2 to 128 characters in length. It cannot start with `http://` or `https://`.
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
	// example:
	//
	// 10.XXX.XXX.50
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// The private port or port range that is used in port forwarding.
	//
	// 	- Valid values: 1 to 65535.
	//
	// 	- To specify a port range, separate the first port and the last port with a forward slash (/), such as 10/20. The first port and the last port are included.
	//
	// 	- If you set InternalPort to a port range, you must also set ExternalPort to a port range. The number of ports in the port ranges must be the same. For example, if you set ExternalPort to 10/20, you can set InternalPort to 80/90.
	//
	// 	- The maximum port range is 1000.
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
}

func (s ModifyForwardEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardEntryRequest) GoString() string {
	return s.String()
}

func (s *ModifyForwardEntryRequest) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *ModifyForwardEntryRequest) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *ModifyForwardEntryRequest) GetForwardEntryId() *string {
	return s.ForwardEntryId
}

func (s *ModifyForwardEntryRequest) GetForwardEntryName() *string {
	return s.ForwardEntryName
}

func (s *ModifyForwardEntryRequest) GetHealthCheckPort() *int32 {
	return s.HealthCheckPort
}

func (s *ModifyForwardEntryRequest) GetInternalIp() *string {
	return s.InternalIp
}

func (s *ModifyForwardEntryRequest) GetInternalPort() *string {
	return s.InternalPort
}

func (s *ModifyForwardEntryRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifyForwardEntryRequest) SetExternalIp(v string) *ModifyForwardEntryRequest {
	s.ExternalIp = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetExternalPort(v string) *ModifyForwardEntryRequest {
	s.ExternalPort = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetForwardEntryId(v string) *ModifyForwardEntryRequest {
	s.ForwardEntryId = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetForwardEntryName(v string) *ModifyForwardEntryRequest {
	s.ForwardEntryName = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetHealthCheckPort(v int32) *ModifyForwardEntryRequest {
	s.HealthCheckPort = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetInternalIp(v string) *ModifyForwardEntryRequest {
	s.InternalIp = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetInternalPort(v string) *ModifyForwardEntryRequest {
	s.InternalPort = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetIpProtocol(v string) *ModifyForwardEntryRequest {
	s.IpProtocol = &v
	return s
}

func (s *ModifyForwardEntryRequest) Validate() error {
	return dara.Validate(s)
}
