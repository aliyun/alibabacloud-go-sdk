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
	// example:
	//
	// zh
	AcceptLanguage *string   `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	DnsServerList  []*string `json:"DnsServerList,omitempty" xml:"DnsServerList,omitempty" type:"Repeated"`
	// example:
	//
	// 501
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// example:
	//
	// gw-c9bc5afd61014165bd58f621b491*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// example:
	//
	// 322
	Id     *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	IpList []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 80
	ServicePort *string `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// example:
	//
	// HTTP
	ServiceProtocol *string `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
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
