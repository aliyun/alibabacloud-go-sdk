// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSSLCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddSSLCertRequest
	GetAcceptLanguage() *string
	SetCertIdentifier(v string) *AddSSLCertRequest
	GetCertIdentifier() *string
	SetDomainId(v int64) *AddSSLCertRequest
	GetDomainId() *int64
	SetGatewayUniqueId(v string) *AddSSLCertRequest
	GetGatewayUniqueId() *string
}

type AddSSLCertRequest struct {
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
	// The ID of the certificate.
	//
	// example:
	//
	// 5213641-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// 0
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-c70622ff52fe49beb29bea9a6f52****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s AddSSLCertRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSSLCertRequest) GoString() string {
	return s.String()
}

func (s *AddSSLCertRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddSSLCertRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *AddSSLCertRequest) GetDomainId() *int64 {
	return s.DomainId
}

func (s *AddSSLCertRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddSSLCertRequest) SetAcceptLanguage(v string) *AddSSLCertRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddSSLCertRequest) SetCertIdentifier(v string) *AddSSLCertRequest {
	s.CertIdentifier = &v
	return s
}

func (s *AddSSLCertRequest) SetDomainId(v int64) *AddSSLCertRequest {
	s.DomainId = &v
	return s
}

func (s *AddSSLCertRequest) SetGatewayUniqueId(v string) *AddSSLCertRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddSSLCertRequest) Validate() error {
	return dara.Validate(s)
}
