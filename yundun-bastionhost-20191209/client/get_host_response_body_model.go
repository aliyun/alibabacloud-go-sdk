// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHost(v *GetHostResponseBodyHost) *GetHostResponseBody
	GetHost() *GetHostResponseBodyHost
	SetRequestId(v string) *GetHostResponseBody
	GetRequestId() *string
}

type GetHostResponseBody struct {
	// The returned information about the host.
	Host *GetHostResponseBodyHost `json:"Host,omitempty" xml:"Host,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHostResponseBody) GoString() string {
	return s.String()
}

func (s *GetHostResponseBody) GetHost() *GetHostResponseBodyHost {
	return s.Host
}

func (s *GetHostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHostResponseBody) SetHost(v *GetHostResponseBodyHost) *GetHostResponseBody {
	s.Host = v
	return s
}

func (s *GetHostResponseBody) SetRequestId(v string) *GetHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHostResponseBody) Validate() error {
	if s.Host != nil {
		if err := s.Host.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHostResponseBodyHost struct {
	// The address type of the host. Valid values:
	//
	// 	- **Public**: a public address
	//
	// 	- **Private**: a private address
	//
	// example:
	//
	// Public
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The description of the host.
	//
	// example:
	//
	// host
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the host.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The hostname.
	//
	// example:
	//
	// host
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The internal endpoint of the host. The value is a domain name or an IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	// The public address of the host. The value is a domain name or an IP address.
	//
	// example:
	//
	// 1.1.XX.XX
	HostPublicAddress *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	// The ID of the network domain to which the host belongs.
	//
	// example:
	//
	// 1
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The operating system of the host. Valid values:
	//
	// 	- **Linux**
	//
	// 	- **Windows**
	//
	// example:
	//
	// Linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The preferred key exchange algorithm of the host. The value of this parameter is returned if OSType is set to Linux. Valid values:
	//
	// 	- **default**
	//
	// 	- **diffie-hellman-group1-sha1**
	//
	// 	- **diffie-hellman-group14-sha1**
	//
	// 	- **diffie-hellman-group-exchange-sha1**
	//
	// example:
	//
	// default
	PrefKex *string `json:"PrefKex,omitempty" xml:"PrefKex,omitempty"`
	// The protocol information about the host.
	Protocols []*GetHostResponseBodyHostProtocols `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The source of the host. Valid values:
	//
	// 	- **Local**: a host in a data center
	//
	// 	- **Ecs**: an Elastic Compute Service (ECS) instance
	//
	// 	- **Rds**: a host in an ApsaraDB MyBase dedicated cluster
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the ECS instance or the host in an ApsaraDB MyBase dedicated cluster.
	//
	// >  If **Local*	- is returned for the **Source*	- parameter, no value is returned for this parameter.
	//
	// example:
	//
	// i-bp19ienyt0yax748****
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The status of the host. Valid values:
	//
	// 	- **Normal**: normal
	//
	// 	- **Release**: released
	//
	// example:
	//
	// Normal
	SourceInstanceState *string `json:"SourceInstanceState,omitempty" xml:"SourceInstanceState,omitempty"`
}

func (s GetHostResponseBodyHost) String() string {
	return dara.Prettify(s)
}

func (s GetHostResponseBodyHost) GoString() string {
	return s.String()
}

func (s *GetHostResponseBodyHost) GetActiveAddressType() *string {
	return s.ActiveAddressType
}

func (s *GetHostResponseBodyHost) GetComment() *string {
	return s.Comment
}

func (s *GetHostResponseBodyHost) GetHostId() *string {
	return s.HostId
}

func (s *GetHostResponseBodyHost) GetHostName() *string {
	return s.HostName
}

func (s *GetHostResponseBodyHost) GetHostPrivateAddress() *string {
	return s.HostPrivateAddress
}

func (s *GetHostResponseBodyHost) GetHostPublicAddress() *string {
	return s.HostPublicAddress
}

func (s *GetHostResponseBodyHost) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *GetHostResponseBodyHost) GetOSType() *string {
	return s.OSType
}

func (s *GetHostResponseBodyHost) GetPrefKex() *string {
	return s.PrefKex
}

func (s *GetHostResponseBodyHost) GetProtocols() []*GetHostResponseBodyHostProtocols {
	return s.Protocols
}

func (s *GetHostResponseBodyHost) GetSource() *string {
	return s.Source
}

func (s *GetHostResponseBodyHost) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *GetHostResponseBodyHost) GetSourceInstanceState() *string {
	return s.SourceInstanceState
}

func (s *GetHostResponseBodyHost) SetActiveAddressType(v string) *GetHostResponseBodyHost {
	s.ActiveAddressType = &v
	return s
}

func (s *GetHostResponseBodyHost) SetComment(v string) *GetHostResponseBodyHost {
	s.Comment = &v
	return s
}

func (s *GetHostResponseBodyHost) SetHostId(v string) *GetHostResponseBodyHost {
	s.HostId = &v
	return s
}

func (s *GetHostResponseBodyHost) SetHostName(v string) *GetHostResponseBodyHost {
	s.HostName = &v
	return s
}

func (s *GetHostResponseBodyHost) SetHostPrivateAddress(v string) *GetHostResponseBodyHost {
	s.HostPrivateAddress = &v
	return s
}

func (s *GetHostResponseBodyHost) SetHostPublicAddress(v string) *GetHostResponseBodyHost {
	s.HostPublicAddress = &v
	return s
}

func (s *GetHostResponseBodyHost) SetNetworkDomainId(v string) *GetHostResponseBodyHost {
	s.NetworkDomainId = &v
	return s
}

func (s *GetHostResponseBodyHost) SetOSType(v string) *GetHostResponseBodyHost {
	s.OSType = &v
	return s
}

func (s *GetHostResponseBodyHost) SetPrefKex(v string) *GetHostResponseBodyHost {
	s.PrefKex = &v
	return s
}

func (s *GetHostResponseBodyHost) SetProtocols(v []*GetHostResponseBodyHostProtocols) *GetHostResponseBodyHost {
	s.Protocols = v
	return s
}

func (s *GetHostResponseBodyHost) SetSource(v string) *GetHostResponseBodyHost {
	s.Source = &v
	return s
}

func (s *GetHostResponseBodyHost) SetSourceInstanceId(v string) *GetHostResponseBodyHost {
	s.SourceInstanceId = &v
	return s
}

func (s *GetHostResponseBodyHost) SetSourceInstanceState(v string) *GetHostResponseBodyHost {
	s.SourceInstanceState = &v
	return s
}

func (s *GetHostResponseBodyHost) Validate() error {
	if s.Protocols != nil {
		for _, item := range s.Protocols {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHostResponseBodyHostProtocols struct {
	// The fingerprint of the host. This parameter uniquely identifies a host. A value is returned for this parameter only if you have performed O\\&M operations on the host by using the bastion host. Otherwise, no value is returned.
	//
	// example:
	//
	// ssh-ed25519|3e:46:5a:e1:1f:0d:39:7e:61:35:d5:fa:7b:2b:**:**
	HostFingerPrint *string `json:"HostFingerPrint,omitempty" xml:"HostFingerPrint,omitempty"`
	// The service port of the host.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol that is used to connect to the host. Valid values:
	//
	// 	- **SSH**
	//
	// 	- **RDP**
	//
	// example:
	//
	// SSH
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
}

func (s GetHostResponseBodyHostProtocols) String() string {
	return dara.Prettify(s)
}

func (s GetHostResponseBodyHostProtocols) GoString() string {
	return s.String()
}

func (s *GetHostResponseBodyHostProtocols) GetHostFingerPrint() *string {
	return s.HostFingerPrint
}

func (s *GetHostResponseBodyHostProtocols) GetPort() *int32 {
	return s.Port
}

func (s *GetHostResponseBodyHostProtocols) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *GetHostResponseBodyHostProtocols) SetHostFingerPrint(v string) *GetHostResponseBodyHostProtocols {
	s.HostFingerPrint = &v
	return s
}

func (s *GetHostResponseBodyHostProtocols) SetPort(v int32) *GetHostResponseBodyHostProtocols {
	s.Port = &v
	return s
}

func (s *GetHostResponseBodyHostProtocols) SetProtocolName(v string) *GetHostResponseBodyHostProtocols {
	s.ProtocolName = &v
	return s
}

func (s *GetHostResponseBodyHostProtocols) Validate() error {
	return dara.Validate(s)
}
