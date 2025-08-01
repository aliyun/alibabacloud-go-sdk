// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSSLCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateSSLCertRequest
	GetAcceptLanguage() *string
	SetCertIdentifier(v string) *UpdateSSLCertRequest
	GetCertIdentifier() *string
	SetDomainId(v int64) *UpdateSSLCertRequest
	GetDomainId() *int64
	SetGatewayUniqueId(v string) *UpdateSSLCertRequest
	GetGatewayUniqueId() *string
}

type UpdateSSLCertRequest struct {
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
	// 5951436-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The ID of the domain name.
	//
	// example:
	//
	// 210
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-2a99625886d54722be94d92e9a69****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s UpdateSSLCertRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSSLCertRequest) GoString() string {
	return s.String()
}

func (s *UpdateSSLCertRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateSSLCertRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *UpdateSSLCertRequest) GetDomainId() *int64 {
	return s.DomainId
}

func (s *UpdateSSLCertRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateSSLCertRequest) SetAcceptLanguage(v string) *UpdateSSLCertRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateSSLCertRequest) SetCertIdentifier(v string) *UpdateSSLCertRequest {
	s.CertIdentifier = &v
	return s
}

func (s *UpdateSSLCertRequest) SetDomainId(v int64) *UpdateSSLCertRequest {
	s.DomainId = &v
	return s
}

func (s *UpdateSSLCertRequest) SetGatewayUniqueId(v string) *UpdateSSLCertRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateSSLCertRequest) Validate() error {
	return dara.Validate(s)
}
