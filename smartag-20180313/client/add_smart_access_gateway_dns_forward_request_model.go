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
	// >- Wildcard domain names and full match are supported. You can enter an asterisk (*) for the domain name to match all. Examples:
	//
	// *.baidu.com. The allowed input format is baidu.com.
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
	// >- This field is invalid. You can ignore it.
	//
	// example:
	//
	// first
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The outbound port index.
	//
	// >- This parameter is optional when OutboundPortType is set to PhysicalPort, and is not required when OutboundPortType is set to Tunnel.
	//
	// >- OutboundPortIndex and OutboundPortName have a one-to-one mapping: 0 - eth0; 2 - lte.
	//
	// example:
	//
	// 0
	OutboundPortIndex *string `json:"OutboundPortIndex,omitempty" xml:"OutboundPortIndex,omitempty"`
	// The outbound port.
	//
	// >- This parameter is optional when OutboundPortType is set to PhysicalPort, and is not required when OutboundPortType is set to Tunnel.
	//
	// >- OutboundPortIndex and OutboundPortName have a one-to-one mapping: 0 - eth0; 2 - lte.
	//
	// example:
	//
	// eth0
	OutboundPortName *string `json:"OutboundPortName,omitempty" xml:"OutboundPortName,omitempty"`
	// The outbound port type.
	//
	// >- A value of Tunnel indicates a tunnel port, and a value of PhysicalPort indicates a physical port.
	//
	// example:
	//
	// Tunnel
	OutboundPortType *string `json:"OutboundPortType,omitempty" xml:"OutboundPortType,omitempty"`
	// The region ID of the Smart Access Gateway instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Smart Access Gateway instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-v9un1ccz22owd76lf8
	SagInsId *string `json:"SagInsId,omitempty" xml:"SagInsId,omitempty"`
	// The serial number of the device.
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
