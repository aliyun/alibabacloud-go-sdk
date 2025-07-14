// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebCertConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCertName(v string) *WebCertConfig
	GetCertName() *string
	SetCertificate(v string) *WebCertConfig
	GetCertificate() *string
	SetPrivateKey(v string) *WebCertConfig
	GetPrivateKey() *string
}

type WebCertConfig struct {
	CertName    *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	PrivateKey  *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
}

func (s WebCertConfig) String() string {
	return dara.Prettify(s)
}

func (s WebCertConfig) GoString() string {
	return s.String()
}

func (s *WebCertConfig) GetCertName() *string {
	return s.CertName
}

func (s *WebCertConfig) GetCertificate() *string {
	return s.Certificate
}

func (s *WebCertConfig) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *WebCertConfig) SetCertName(v string) *WebCertConfig {
	s.CertName = &v
	return s
}

func (s *WebCertConfig) SetCertificate(v string) *WebCertConfig {
	s.Certificate = &v
	return s
}

func (s *WebCertConfig) SetPrivateKey(v string) *WebCertConfig {
	s.PrivateKey = &v
	return s
}

func (s *WebCertConfig) Validate() error {
	return dara.Validate(s)
}
