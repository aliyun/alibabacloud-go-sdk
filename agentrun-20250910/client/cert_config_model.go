// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCertName(v string) *CertConfig
	GetCertName() *string
	SetCertificate(v string) *CertConfig
	GetCertificate() *string
	SetPrivateKey(v string) *CertConfig
	GetPrivateKey() *string
}

type CertConfig struct {
	// 证书的名称。支持字母、数字、下划线（_）和短划线（-），不能以数字和短划线（-）开头。长度范围为 1~128 个字符。
	//
	// example:
	//
	// my-cert
	CertName *string `json:"certName,omitempty" xml:"certName,omitempty"`
	// 证书，如果是证书链，则需要依次填写多个证书。
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\nxxxxx\n-----END CERTIFICATE-----
	Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty"`
	// 私钥。
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY-----\nxxxxx\n-----END RSA PRIVATE KEY-----
	PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty"`
}

func (s CertConfig) String() string {
	return dara.Prettify(s)
}

func (s CertConfig) GoString() string {
	return s.String()
}

func (s *CertConfig) GetCertName() *string {
	return s.CertName
}

func (s *CertConfig) GetCertificate() *string {
	return s.Certificate
}

func (s *CertConfig) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *CertConfig) SetCertName(v string) *CertConfig {
	s.CertName = &v
	return s
}

func (s *CertConfig) SetCertificate(v string) *CertConfig {
	s.Certificate = &v
	return s
}

func (s *CertConfig) SetPrivateKey(v string) *CertConfig {
	s.PrivateKey = &v
	return s
}

func (s *CertConfig) Validate() error {
	return dara.Validate(s)
}
