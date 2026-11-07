// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportServicesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ImportServicesShrinkRequest
	GetAcceptLanguage() *string
	SetFcAlias(v string) *ImportServicesShrinkRequest
	GetFcAlias() *string
	SetFcServiceName(v string) *ImportServicesShrinkRequest
	GetFcServiceName() *string
	SetFcVersion(v string) *ImportServicesShrinkRequest
	GetFcVersion() *string
	SetGatewayUniqueId(v string) *ImportServicesShrinkRequest
	GetGatewayUniqueId() *string
	SetServiceListShrink(v string) *ImportServicesShrinkRequest
	GetServiceListShrink() *string
	SetSourceId(v int64) *ImportServicesShrinkRequest
	GetSourceId() *int64
	SetSourceType(v string) *ImportServicesShrinkRequest
	GetSourceType() *string
	SetTlsSetting(v string) *ImportServicesShrinkRequest
	GetTlsSetting() *string
}

type ImportServicesShrinkRequest struct {
	// The language of the response. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The Function Compute service alias. This parameter is required when the service source is Function Compute.
	//
	// example:
	//
	// test-alias
	FcAlias *string `json:"FcAlias,omitempty" xml:"FcAlias,omitempty"`
	// The Function Compute service name. This parameter is required when the service source is Function Compute.
	//
	// example:
	//
	// guide
	FcServiceName *string `json:"FcServiceName,omitempty" xml:"FcServiceName,omitempty"`
	// The Function Compute service version. This parameter is required when the service source is Function Compute.
	//
	// example:
	//
	// LATEST
	FcVersion *string `json:"FcVersion,omitempty" xml:"FcVersion,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-77e1153db6e14c0a8b1fae20bcb89ca5
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The service list.
	//
	// example:
	//
	// DNS
	ServiceListShrink *string `json:"ServiceList,omitempty" xml:"ServiceList,omitempty"`
	// The service source ID. This parameter is used when you add a service from a specified source.
	//
	// example:
	//
	// 100
	SourceId *int64 `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The service source. Valid values:
	//
	// - MSE: Microservices Engine (MSE Nacos).
	//
	// - K8s: Container Service for Kubernetes (ACK).
	//
	// - VIP: Fixed address.
	//
	// - DNS: DNS domain name.
	//
	// example:
	//
	// DNS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The TLS settings.
	//
	// - TLS pattern: mode
	//
	// - Certificate ID: certId
	//
	// - CA certificate ID: caCertId
	//
	// - CA certificate public key: caCertContent
	//
	// - Service name identity: sni
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

func (s ImportServicesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportServicesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportServicesShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ImportServicesShrinkRequest) GetFcAlias() *string {
	return s.FcAlias
}

func (s *ImportServicesShrinkRequest) GetFcServiceName() *string {
	return s.FcServiceName
}

func (s *ImportServicesShrinkRequest) GetFcVersion() *string {
	return s.FcVersion
}

func (s *ImportServicesShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ImportServicesShrinkRequest) GetServiceListShrink() *string {
	return s.ServiceListShrink
}

func (s *ImportServicesShrinkRequest) GetSourceId() *int64 {
	return s.SourceId
}

func (s *ImportServicesShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ImportServicesShrinkRequest) GetTlsSetting() *string {
	return s.TlsSetting
}

func (s *ImportServicesShrinkRequest) SetAcceptLanguage(v string) *ImportServicesShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ImportServicesShrinkRequest) SetFcAlias(v string) *ImportServicesShrinkRequest {
	s.FcAlias = &v
	return s
}

func (s *ImportServicesShrinkRequest) SetFcServiceName(v string) *ImportServicesShrinkRequest {
	s.FcServiceName = &v
	return s
}

func (s *ImportServicesShrinkRequest) SetFcVersion(v string) *ImportServicesShrinkRequest {
	s.FcVersion = &v
	return s
}

func (s *ImportServicesShrinkRequest) SetGatewayUniqueId(v string) *ImportServicesShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *ImportServicesShrinkRequest) SetServiceListShrink(v string) *ImportServicesShrinkRequest {
	s.ServiceListShrink = &v
	return s
}

func (s *ImportServicesShrinkRequest) SetSourceId(v int64) *ImportServicesShrinkRequest {
	s.SourceId = &v
	return s
}

func (s *ImportServicesShrinkRequest) SetSourceType(v string) *ImportServicesShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *ImportServicesShrinkRequest) SetTlsSetting(v string) *ImportServicesShrinkRequest {
	s.TlsSetting = &v
	return s
}

func (s *ImportServicesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
