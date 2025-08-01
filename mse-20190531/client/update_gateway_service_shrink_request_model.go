// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayServiceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayServiceShrinkRequest
	GetAcceptLanguage() *string
	SetDnsServerListShrink(v string) *UpdateGatewayServiceShrinkRequest
	GetDnsServerListShrink() *string
	SetGatewayId(v int64) *UpdateGatewayServiceShrinkRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayServiceShrinkRequest
	GetGatewayUniqueId() *string
	SetId(v string) *UpdateGatewayServiceShrinkRequest
	GetId() *string
	SetIpListShrink(v string) *UpdateGatewayServiceShrinkRequest
	GetIpListShrink() *string
	SetName(v string) *UpdateGatewayServiceShrinkRequest
	GetName() *string
	SetServicePort(v string) *UpdateGatewayServiceShrinkRequest
	GetServicePort() *string
	SetServiceProtocol(v string) *UpdateGatewayServiceShrinkRequest
	GetServiceProtocol() *string
	SetTlsSetting(v string) *UpdateGatewayServiceShrinkRequest
	GetTlsSetting() *string
}

type UpdateGatewayServiceShrinkRequest struct {
	// example:
	//
	// zh
	AcceptLanguage      *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	DnsServerListShrink *string `json:"DnsServerList,omitempty" xml:"DnsServerList,omitempty"`
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
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IpListShrink *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
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

func (s UpdateGatewayServiceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayServiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayServiceShrinkRequest) GetDnsServerListShrink() *string {
	return s.DnsServerListShrink
}

func (s *UpdateGatewayServiceShrinkRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayServiceShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayServiceShrinkRequest) GetId() *string {
	return s.Id
}

func (s *UpdateGatewayServiceShrinkRequest) GetIpListShrink() *string {
	return s.IpListShrink
}

func (s *UpdateGatewayServiceShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateGatewayServiceShrinkRequest) GetServicePort() *string {
	return s.ServicePort
}

func (s *UpdateGatewayServiceShrinkRequest) GetServiceProtocol() *string {
	return s.ServiceProtocol
}

func (s *UpdateGatewayServiceShrinkRequest) GetTlsSetting() *string {
	return s.TlsSetting
}

func (s *UpdateGatewayServiceShrinkRequest) SetAcceptLanguage(v string) *UpdateGatewayServiceShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayServiceShrinkRequest) SetDnsServerListShrink(v string) *UpdateGatewayServiceShrinkRequest {
	s.DnsServerListShrink = &v
	return s
}

func (s *UpdateGatewayServiceShrinkRequest) SetGatewayId(v int64) *UpdateGatewayServiceShrinkRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayServiceShrinkRequest) SetGatewayUniqueId(v string) *UpdateGatewayServiceShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayServiceShrinkRequest) SetId(v string) *UpdateGatewayServiceShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayServiceShrinkRequest) SetIpListShrink(v string) *UpdateGatewayServiceShrinkRequest {
	s.IpListShrink = &v
	return s
}

func (s *UpdateGatewayServiceShrinkRequest) SetName(v string) *UpdateGatewayServiceShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateGatewayServiceShrinkRequest) SetServicePort(v string) *UpdateGatewayServiceShrinkRequest {
	s.ServicePort = &v
	return s
}

func (s *UpdateGatewayServiceShrinkRequest) SetServiceProtocol(v string) *UpdateGatewayServiceShrinkRequest {
	s.ServiceProtocol = &v
	return s
}

func (s *UpdateGatewayServiceShrinkRequest) SetTlsSetting(v string) *UpdateGatewayServiceShrinkRequest {
	s.TlsSetting = &v
	return s
}

func (s *UpdateGatewayServiceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
