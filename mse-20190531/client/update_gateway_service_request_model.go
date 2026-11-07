// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayServiceRequest
	GetAcceptLanguage() *string
	SetDnsServerList(v []*string) *UpdateGatewayServiceRequest
	GetDnsServerList() []*string
	SetGatewayId(v int64) *UpdateGatewayServiceRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayServiceRequest
	GetGatewayUniqueId() *string
	SetId(v string) *UpdateGatewayServiceRequest
	GetId() *string
	SetIpList(v []*string) *UpdateGatewayServiceRequest
	GetIpList() []*string
	SetName(v string) *UpdateGatewayServiceRequest
	GetName() *string
	SetServicePort(v string) *UpdateGatewayServiceRequest
	GetServicePort() *string
	SetServiceProtocol(v string) *UpdateGatewayServiceRequest
	GetServiceProtocol() *string
	SetTlsSetting(v string) *UpdateGatewayServiceRequest
	GetTlsSetting() *string
}

type UpdateGatewayServiceRequest struct {
	// The language in which you want the results to be displayed. Valid values: zh (default): Chinese. en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The list of Domain Name System (DNS) servers (for DNS domain name services only).
	DnsServerList []*string `json:"DnsServerList,omitempty" xml:"DnsServerList,omitempty" type:"Repeated"`
	// The gateway ID.
	//
	// example:
	//
	// 501
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-c9bc5afd61014165bd58f621b491*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// 322
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The address list of fixed address services or domain name list of Domain Name System (DNS) services.
	IpList []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// The service name. This parameter is deprecated and cannot be modified. The value of this parameter must be the same as the current service name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The service port. Specify this parameter for DNS services.
	//
	// example:
	//
	// 80
	ServicePort *string `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// The service protocol.
	//
	// - HTTP
	//
	// - HTTPS
	//
	// - HTTP2
	//
	// - GRPC
	//
	// - DUBBO
	//
	// example:
	//
	// HTTP
	ServiceProtocol *string `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	// The Transport Layer Security (TLS) settings. Valid values:
	//
	// - mode: TLS mode
	//
	// - certId: certificate ID
	//
	// - caCertId: CA certificate ID
	//
	// - caCertContent: CA certificate public key
	//
	// - sni: service name identification
	//
	// example:
	//
	// {
	//
	//       "mode": "MUTUAL",
	//
	//       "certId": "1*****-cn-hangzhou",
	//
	//       "caCertContent": "123",
	//
	//       "sni": "ceshi"
	//
	// }
	TlsSetting *string `json:"TlsSetting,omitempty" xml:"TlsSetting,omitempty"`
}

func (s UpdateGatewayServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayServiceRequest) GetDnsServerList() []*string {
	return s.DnsServerList
}

func (s *UpdateGatewayServiceRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayServiceRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayServiceRequest) GetId() *string {
	return s.Id
}

func (s *UpdateGatewayServiceRequest) GetIpList() []*string {
	return s.IpList
}

func (s *UpdateGatewayServiceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateGatewayServiceRequest) GetServicePort() *string {
	return s.ServicePort
}

func (s *UpdateGatewayServiceRequest) GetServiceProtocol() *string {
	return s.ServiceProtocol
}

func (s *UpdateGatewayServiceRequest) GetTlsSetting() *string {
	return s.TlsSetting
}

func (s *UpdateGatewayServiceRequest) SetAcceptLanguage(v string) *UpdateGatewayServiceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayServiceRequest) SetDnsServerList(v []*string) *UpdateGatewayServiceRequest {
	s.DnsServerList = v
	return s
}

func (s *UpdateGatewayServiceRequest) SetGatewayId(v int64) *UpdateGatewayServiceRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayServiceRequest) SetGatewayUniqueId(v string) *UpdateGatewayServiceRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayServiceRequest) SetId(v string) *UpdateGatewayServiceRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayServiceRequest) SetIpList(v []*string) *UpdateGatewayServiceRequest {
	s.IpList = v
	return s
}

func (s *UpdateGatewayServiceRequest) SetName(v string) *UpdateGatewayServiceRequest {
	s.Name = &v
	return s
}

func (s *UpdateGatewayServiceRequest) SetServicePort(v string) *UpdateGatewayServiceRequest {
	s.ServicePort = &v
	return s
}

func (s *UpdateGatewayServiceRequest) SetServiceProtocol(v string) *UpdateGatewayServiceRequest {
	s.ServiceProtocol = &v
	return s
}

func (s *UpdateGatewayServiceRequest) SetTlsSetting(v string) *UpdateGatewayServiceRequest {
	s.TlsSetting = &v
	return s
}

func (s *UpdateGatewayServiceRequest) Validate() error {
	return dara.Validate(s)
}
