// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSmartAccessGatewayDnsForwardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *AddSmartAccessGatewayDnsForwardRequest
	GetDomain() *string
	SetMasterIp(v string) *AddSmartAccessGatewayDnsForwardRequest
	GetMasterIp() *string
	SetMode(v string) *AddSmartAccessGatewayDnsForwardRequest
	GetMode() *string
	SetOutboundPortIndex(v string) *AddSmartAccessGatewayDnsForwardRequest
	GetOutboundPortIndex() *string
	SetOutboundPortName(v string) *AddSmartAccessGatewayDnsForwardRequest
	GetOutboundPortName() *string
	SetOutboundPortType(v string) *AddSmartAccessGatewayDnsForwardRequest
	GetOutboundPortType() *string
	SetRegionId(v string) *AddSmartAccessGatewayDnsForwardRequest
	GetRegionId() *string
	SetSagInsId(v string) *AddSmartAccessGatewayDnsForwardRequest
	GetSagInsId() *string
	SetSagSn(v string) *AddSmartAccessGatewayDnsForwardRequest
	GetSagSn() *string
	SetSlaveIp(v string) *AddSmartAccessGatewayDnsForwardRequest
	GetSlaveIp() *string
}

type AddSmartAccessGatewayDnsForwardRequest struct {
	// The domain name.
	//
	// >
	//
	// 	- Wildcard domain names are supported. You can use the wildcard character asterisk (\\*) to specify a wildcard domain name.
	//
	// For example, you can enter \\*.baidu.com to specify the domain name baidu.com.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The primary DNS server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 114.114.114.114
	MasterIp *string `json:"MasterIp,omitempty" xml:"MasterIp,omitempty"`
	// The forwarding mode.
	//
	// >
	//
	// 	- This parameter is not in use. Ignore this parameter.
	//
	// example:
	//
	// first
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The number of the egress port.
	//
	// >
	//
	// 	- This parameter is optional if OutboundPortType is set to PhysicalPort. Ignore this parameter if OutboundPortType is set to Tunnel.
	//
	// 	- The value of OutboundPortIndex is the unique index of the port name specified by poOutboundPortName. For example, 0 is the index for the port named eth0, and 2 is the index for the port named lte.
	//
	// example:
	//
	// 0
	OutboundPortIndex *string `json:"OutboundPortIndex,omitempty" xml:"OutboundPortIndex,omitempty"`
	// The egress port.
	//
	// >
	//
	// 	- This parameter is optional if OutboundPortType is set to PhysicalPort. Ignore this parameter if OutboundPortType is set to Tunnel.
	//
	// 	- The value of OutboundPortIndex is the unique index of the port name specified by poOutboundPortName. For example, 0 is the index for the port named eth0, and 2 is the index for the port named lte.
	//
	// example:
	//
	// eth0
	OutboundPortName *string `json:"OutboundPortName,omitempty" xml:"OutboundPortName,omitempty"`
	// The type of the egress port.
	//
	// >
	//
	// 	- A value of Tunnel specifies a tunnel, and a value of PhysicalPort specifies a physical port.
	//
	// example:
	//
	// Tunnel
	OutboundPortType *string `json:"OutboundPortType,omitempty" xml:"OutboundPortType,omitempty"`
	// The ID of the region where the instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-v9un1ccz22owd76lf8
	SagInsId *string `json:"SagInsId,omitempty" xml:"SagInsId,omitempty"`
	// The serial number (SN) of the device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sagf4dkqh78
	SagSn *string `json:"SagSn,omitempty" xml:"SagSn,omitempty"`
	// The secondary DNS server.
	//
	// example:
	//
	// 172.16.0.14
	SlaveIp *string `json:"SlaveIp,omitempty" xml:"SlaveIp,omitempty"`
}

func (s AddSmartAccessGatewayDnsForwardRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSmartAccessGatewayDnsForwardRequest) GoString() string {
	return s.String()
}

func (s *AddSmartAccessGatewayDnsForwardRequest) GetDomain() *string {
	return s.Domain
}

func (s *AddSmartAccessGatewayDnsForwardRequest) GetMasterIp() *string {
	return s.MasterIp
}

func (s *AddSmartAccessGatewayDnsForwardRequest) GetMode() *string {
	return s.Mode
}

func (s *AddSmartAccessGatewayDnsForwardRequest) GetOutboundPortIndex() *string {
	return s.OutboundPortIndex
}

func (s *AddSmartAccessGatewayDnsForwardRequest) GetOutboundPortName() *string {
	return s.OutboundPortName
}

func (s *AddSmartAccessGatewayDnsForwardRequest) GetOutboundPortType() *string {
	return s.OutboundPortType
}

func (s *AddSmartAccessGatewayDnsForwardRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddSmartAccessGatewayDnsForwardRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *AddSmartAccessGatewayDnsForwardRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *AddSmartAccessGatewayDnsForwardRequest) GetSlaveIp() *string {
	return s.SlaveIp
}

func (s *AddSmartAccessGatewayDnsForwardRequest) SetDomain(v string) *AddSmartAccessGatewayDnsForwardRequest {
	s.Domain = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardRequest) SetMasterIp(v string) *AddSmartAccessGatewayDnsForwardRequest {
	s.MasterIp = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardRequest) SetMode(v string) *AddSmartAccessGatewayDnsForwardRequest {
	s.Mode = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardRequest) SetOutboundPortIndex(v string) *AddSmartAccessGatewayDnsForwardRequest {
	s.OutboundPortIndex = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardRequest) SetOutboundPortName(v string) *AddSmartAccessGatewayDnsForwardRequest {
	s.OutboundPortName = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardRequest) SetOutboundPortType(v string) *AddSmartAccessGatewayDnsForwardRequest {
	s.OutboundPortType = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardRequest) SetRegionId(v string) *AddSmartAccessGatewayDnsForwardRequest {
	s.RegionId = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardRequest) SetSagInsId(v string) *AddSmartAccessGatewayDnsForwardRequest {
	s.SagInsId = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardRequest) SetSagSn(v string) *AddSmartAccessGatewayDnsForwardRequest {
	s.SagSn = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardRequest) SetSlaveIp(v string) *AddSmartAccessGatewayDnsForwardRequest {
	s.SlaveIp = &v
	return s
}

func (s *AddSmartAccessGatewayDnsForwardRequest) Validate() error {
	return dara.Validate(s)
}
