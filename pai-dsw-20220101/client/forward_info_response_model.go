// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForwardInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v []*string) *ForwardInfoResponse
	GetAccessType() []*string
	SetConnectInfo(v *ForwardInfoResponseConnectInfo) *ForwardInfoResponse
	GetConnectInfo() *ForwardInfoResponseConnectInfo
	SetContainerName(v string) *ForwardInfoResponse
	GetContainerName() *string
	SetEipAllocationId(v string) *ForwardInfoResponse
	GetEipAllocationId() *string
	SetEnable(v bool) *ForwardInfoResponse
	GetEnable() *bool
	SetExternalPort(v string) *ForwardInfoResponse
	GetExternalPort() *string
	SetForwardPort(v string) *ForwardInfoResponse
	GetForwardPort() *string
	SetName(v string) *ForwardInfoResponse
	GetName() *string
	SetNatGatewayId(v string) *ForwardInfoResponse
	GetNatGatewayId() *string
	SetSSHPublicKey(v string) *ForwardInfoResponse
	GetSSHPublicKey() *string
}

type ForwardInfoResponse struct {
	AccessType  []*string                       `json:"AccessType,omitempty" xml:"AccessType,omitempty" type:"Repeated"`
	ConnectInfo *ForwardInfoResponseConnectInfo `json:"ConnectInfo,omitempty" xml:"ConnectInfo,omitempty" type:"Struct"`
	// example:
	//
	// dsw-notebook
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// example:
	//
	// eip-25877c70gddh****
	EipAllocationId *string `json:"EipAllocationId,omitempty" xml:"EipAllocationId,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 1024
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// example:
	//
	// 22
	ForwardPort *string `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	// example:
	//
	// ssh
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	SSHPublicKey *string `json:"SSHPublicKey,omitempty" xml:"SSHPublicKey,omitempty"`
}

func (s ForwardInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ForwardInfoResponse) GoString() string {
	return s.String()
}

func (s *ForwardInfoResponse) GetAccessType() []*string {
	return s.AccessType
}

func (s *ForwardInfoResponse) GetConnectInfo() *ForwardInfoResponseConnectInfo {
	return s.ConnectInfo
}

func (s *ForwardInfoResponse) GetContainerName() *string {
	return s.ContainerName
}

func (s *ForwardInfoResponse) GetEipAllocationId() *string {
	return s.EipAllocationId
}

func (s *ForwardInfoResponse) GetEnable() *bool {
	return s.Enable
}

func (s *ForwardInfoResponse) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *ForwardInfoResponse) GetForwardPort() *string {
	return s.ForwardPort
}

func (s *ForwardInfoResponse) GetName() *string {
	return s.Name
}

func (s *ForwardInfoResponse) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ForwardInfoResponse) GetSSHPublicKey() *string {
	return s.SSHPublicKey
}

func (s *ForwardInfoResponse) SetAccessType(v []*string) *ForwardInfoResponse {
	s.AccessType = v
	return s
}

func (s *ForwardInfoResponse) SetConnectInfo(v *ForwardInfoResponseConnectInfo) *ForwardInfoResponse {
	s.ConnectInfo = v
	return s
}

func (s *ForwardInfoResponse) SetContainerName(v string) *ForwardInfoResponse {
	s.ContainerName = &v
	return s
}

func (s *ForwardInfoResponse) SetEipAllocationId(v string) *ForwardInfoResponse {
	s.EipAllocationId = &v
	return s
}

func (s *ForwardInfoResponse) SetEnable(v bool) *ForwardInfoResponse {
	s.Enable = &v
	return s
}

func (s *ForwardInfoResponse) SetExternalPort(v string) *ForwardInfoResponse {
	s.ExternalPort = &v
	return s
}

func (s *ForwardInfoResponse) SetForwardPort(v string) *ForwardInfoResponse {
	s.ForwardPort = &v
	return s
}

func (s *ForwardInfoResponse) SetName(v string) *ForwardInfoResponse {
	s.Name = &v
	return s
}

func (s *ForwardInfoResponse) SetNatGatewayId(v string) *ForwardInfoResponse {
	s.NatGatewayId = &v
	return s
}

func (s *ForwardInfoResponse) SetSSHPublicKey(v string) *ForwardInfoResponse {
	s.SSHPublicKey = &v
	return s
}

func (s *ForwardInfoResponse) Validate() error {
	if s.ConnectInfo != nil {
		if err := s.ConnectInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ForwardInfoResponseConnectInfo struct {
	Internet *ForwardInfoResponseConnectInfoInternet `json:"Internet,omitempty" xml:"Internet,omitempty" type:"Struct"`
	Intranet *ForwardInfoResponseConnectInfoIntranet `json:"Intranet,omitempty" xml:"Intranet,omitempty" type:"Struct"`
	// example:
	//
	// DNAT and privateZone are both ready.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// Ready
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
}

func (s ForwardInfoResponseConnectInfo) String() string {
	return dara.Prettify(s)
}

func (s ForwardInfoResponseConnectInfo) GoString() string {
	return s.String()
}

func (s *ForwardInfoResponseConnectInfo) GetInternet() *ForwardInfoResponseConnectInfoInternet {
	return s.Internet
}

func (s *ForwardInfoResponseConnectInfo) GetIntranet() *ForwardInfoResponseConnectInfoIntranet {
	return s.Intranet
}

func (s *ForwardInfoResponseConnectInfo) GetMessage() *string {
	return s.Message
}

func (s *ForwardInfoResponseConnectInfo) GetPhase() *string {
	return s.Phase
}

func (s *ForwardInfoResponseConnectInfo) SetInternet(v *ForwardInfoResponseConnectInfoInternet) *ForwardInfoResponseConnectInfo {
	s.Internet = v
	return s
}

func (s *ForwardInfoResponseConnectInfo) SetIntranet(v *ForwardInfoResponseConnectInfoIntranet) *ForwardInfoResponseConnectInfo {
	s.Intranet = v
	return s
}

func (s *ForwardInfoResponseConnectInfo) SetMessage(v string) *ForwardInfoResponseConnectInfo {
	s.Message = &v
	return s
}

func (s *ForwardInfoResponseConnectInfo) SetPhase(v string) *ForwardInfoResponseConnectInfo {
	s.Phase = &v
	return s
}

func (s *ForwardInfoResponseConnectInfo) Validate() error {
	if s.Internet != nil {
		if err := s.Internet.Validate(); err != nil {
			return err
		}
	}
	if s.Intranet != nil {
		if err := s.Intranet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ForwardInfoResponseConnectInfoInternet struct {
	// example:
	//
	// 47.111.119.114
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// 22
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ForwardInfoResponseConnectInfoInternet) String() string {
	return dara.Prettify(s)
}

func (s ForwardInfoResponseConnectInfoInternet) GoString() string {
	return s.String()
}

func (s *ForwardInfoResponseConnectInfoInternet) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ForwardInfoResponseConnectInfoInternet) GetPort() *string {
	return s.Port
}

func (s *ForwardInfoResponseConnectInfoInternet) SetEndpoint(v string) *ForwardInfoResponseConnectInfoInternet {
	s.Endpoint = &v
	return s
}

func (s *ForwardInfoResponseConnectInfoInternet) SetPort(v string) *ForwardInfoResponseConnectInfoInternet {
	s.Port = &v
	return s
}

func (s *ForwardInfoResponseConnectInfoInternet) Validate() error {
	return dara.Validate(s)
}

type ForwardInfoResponseConnectInfoIntranet struct {
	// example:
	//
	// dsw-notebook-22-urz3u6cnu0uts7ej9r.dsw-5cc6083084818f60.dsw.pai.alibaba.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// 22
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ForwardInfoResponseConnectInfoIntranet) String() string {
	return dara.Prettify(s)
}

func (s ForwardInfoResponseConnectInfoIntranet) GoString() string {
	return s.String()
}

func (s *ForwardInfoResponseConnectInfoIntranet) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ForwardInfoResponseConnectInfoIntranet) GetPort() *string {
	return s.Port
}

func (s *ForwardInfoResponseConnectInfoIntranet) SetEndpoint(v string) *ForwardInfoResponseConnectInfoIntranet {
	s.Endpoint = &v
	return s
}

func (s *ForwardInfoResponseConnectInfoIntranet) SetPort(v string) *ForwardInfoResponseConnectInfoIntranet {
	s.Port = &v
	return s
}

func (s *ForwardInfoResponseConnectInfoIntranet) Validate() error {
	return dara.Validate(s)
}
