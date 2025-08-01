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
	ServiceListShrink *string `json:"ServiceList,omitempty" xml:"ServiceList,omitempty"`
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
