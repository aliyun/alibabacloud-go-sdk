// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSSLCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListSSLCertRequest
	GetAcceptLanguage() *string
	SetCertName(v string) *ListSSLCertRequest
	GetCertName() *string
	SetGatewayUniqueId(v string) *ListSSLCertRequest
	GetGatewayUniqueId() *string
}

type ListSSLCertRequest struct {
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
	// example:
	//
	// certabc
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-3f97e2989c344f35ab3fd62b19f1****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s ListSSLCertRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSSLCertRequest) GoString() string {
	return s.String()
}

func (s *ListSSLCertRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListSSLCertRequest) GetCertName() *string {
	return s.CertName
}

func (s *ListSSLCertRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListSSLCertRequest) SetAcceptLanguage(v string) *ListSSLCertRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListSSLCertRequest) SetCertName(v string) *ListSSLCertRequest {
	s.CertName = &v
	return s
}

func (s *ListSSLCertRequest) SetGatewayUniqueId(v string) *ListSSLCertRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListSSLCertRequest) Validate() error {
	return dara.Validate(s)
}
