// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ImportServicesRequest
	GetAcceptLanguage() *string
	SetFcAlias(v string) *ImportServicesRequest
	GetFcAlias() *string
	SetFcServiceName(v string) *ImportServicesRequest
	GetFcServiceName() *string
	SetFcVersion(v string) *ImportServicesRequest
	GetFcVersion() *string
	SetGatewayUniqueId(v string) *ImportServicesRequest
	GetGatewayUniqueId() *string
	SetServiceList(v []*ImportServicesRequestServiceList) *ImportServicesRequest
	GetServiceList() []*ImportServicesRequestServiceList
	SetSourceId(v int64) *ImportServicesRequest
	GetSourceId() *int64
	SetSourceType(v string) *ImportServicesRequest
	GetSourceType() *string
	SetTlsSetting(v string) *ImportServicesRequest
	GetTlsSetting() *string
}

type ImportServicesRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	FcAlias        *string `json:"FcAlias,omitempty" xml:"FcAlias,omitempty"`
	FcServiceName  *string `json:"FcServiceName,omitempty" xml:"FcServiceName,omitempty"`
	FcVersion      *string `json:"FcVersion,omitempty" xml:"FcVersion,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-77e1153db6e14c0a8b1fae20bcb89ca5
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The information about services.
	//
	// example:
	//
	// DNS
	ServiceList []*ImportServicesRequestServiceList `json:"ServiceList,omitempty" xml:"ServiceList,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	SourceId *int64 `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The service source. Valid values:
	//
	// 	- MSE: MSE Nacos instance
	//
	// 	- K8s: ACK cluster
	//
	// 	- VIP: fixed address
	//
	// 	- DNS: DNS domain
	//
	// example:
	//
	// DNS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The Transport Layer Security (TLS) settings. Valid values:
	//
	// 	- mode: TLS mode
	//
	// 	- certId: certificate ID
	//
	// 	- caCertId: CA certificate ID
	//
	// 	- caCertContent: CA certificate public key
	//
	// 	- sni: service name identification
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

func (s ImportServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportServicesRequest) GoString() string {
	return s.String()
}

func (s *ImportServicesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ImportServicesRequest) GetFcAlias() *string {
	return s.FcAlias
}

func (s *ImportServicesRequest) GetFcServiceName() *string {
	return s.FcServiceName
}

func (s *ImportServicesRequest) GetFcVersion() *string {
	return s.FcVersion
}

func (s *ImportServicesRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ImportServicesRequest) GetServiceList() []*ImportServicesRequestServiceList {
	return s.ServiceList
}

func (s *ImportServicesRequest) GetSourceId() *int64 {
	return s.SourceId
}

func (s *ImportServicesRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ImportServicesRequest) GetTlsSetting() *string {
	return s.TlsSetting
}

func (s *ImportServicesRequest) SetAcceptLanguage(v string) *ImportServicesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ImportServicesRequest) SetFcAlias(v string) *ImportServicesRequest {
	s.FcAlias = &v
	return s
}

func (s *ImportServicesRequest) SetFcServiceName(v string) *ImportServicesRequest {
	s.FcServiceName = &v
	return s
}

func (s *ImportServicesRequest) SetFcVersion(v string) *ImportServicesRequest {
	s.FcVersion = &v
	return s
}

func (s *ImportServicesRequest) SetGatewayUniqueId(v string) *ImportServicesRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *ImportServicesRequest) SetServiceList(v []*ImportServicesRequestServiceList) *ImportServicesRequest {
	s.ServiceList = v
	return s
}

func (s *ImportServicesRequest) SetSourceId(v int64) *ImportServicesRequest {
	s.SourceId = &v
	return s
}

func (s *ImportServicesRequest) SetSourceType(v string) *ImportServicesRequest {
	s.SourceType = &v
	return s
}

func (s *ImportServicesRequest) SetTlsSetting(v string) *ImportServicesRequest {
	s.TlsSetting = &v
	return s
}

func (s *ImportServicesRequest) Validate() error {
	if s.ServiceList != nil {
		for _, item := range s.ServiceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImportServicesRequestServiceList struct {
	DnsServerList []*string `json:"DnsServerList,omitempty" xml:"DnsServerList,omitempty" type:"Repeated"`
	// The group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The IP addresses of the service.
	Ips []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	// The name of the service.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// public
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	SaeAppId  *string `json:"SaeAppId,omitempty" xml:"SaeAppId,omitempty"`
	// The port of the service.
	//
	// example:
	//
	// 8080
	ServicePort *int64 `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// The protocol of the service.
	//
	// example:
	//
	// GRPC, HTTP
	ServiceProtocol *string `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
}

func (s ImportServicesRequestServiceList) String() string {
	return dara.Prettify(s)
}

func (s ImportServicesRequestServiceList) GoString() string {
	return s.String()
}

func (s *ImportServicesRequestServiceList) GetDnsServerList() []*string {
	return s.DnsServerList
}

func (s *ImportServicesRequestServiceList) GetGroupName() *string {
	return s.GroupName
}

func (s *ImportServicesRequestServiceList) GetIps() []*string {
	return s.Ips
}

func (s *ImportServicesRequestServiceList) GetName() *string {
	return s.Name
}

func (s *ImportServicesRequestServiceList) GetNamespace() *string {
	return s.Namespace
}

func (s *ImportServicesRequestServiceList) GetSaeAppId() *string {
	return s.SaeAppId
}

func (s *ImportServicesRequestServiceList) GetServicePort() *int64 {
	return s.ServicePort
}

func (s *ImportServicesRequestServiceList) GetServiceProtocol() *string {
	return s.ServiceProtocol
}

func (s *ImportServicesRequestServiceList) SetDnsServerList(v []*string) *ImportServicesRequestServiceList {
	s.DnsServerList = v
	return s
}

func (s *ImportServicesRequestServiceList) SetGroupName(v string) *ImportServicesRequestServiceList {
	s.GroupName = &v
	return s
}

func (s *ImportServicesRequestServiceList) SetIps(v []*string) *ImportServicesRequestServiceList {
	s.Ips = v
	return s
}

func (s *ImportServicesRequestServiceList) SetName(v string) *ImportServicesRequestServiceList {
	s.Name = &v
	return s
}

func (s *ImportServicesRequestServiceList) SetNamespace(v string) *ImportServicesRequestServiceList {
	s.Namespace = &v
	return s
}

func (s *ImportServicesRequestServiceList) SetSaeAppId(v string) *ImportServicesRequestServiceList {
	s.SaeAppId = &v
	return s
}

func (s *ImportServicesRequestServiceList) SetServicePort(v int64) *ImportServicesRequestServiceList {
	s.ServicePort = &v
	return s
}

func (s *ImportServicesRequestServiceList) SetServiceProtocol(v string) *ImportServicesRequestServiceList {
	s.ServiceProtocol = &v
	return s
}

func (s *ImportServicesRequestServiceList) Validate() error {
	return dara.Validate(s)
}
