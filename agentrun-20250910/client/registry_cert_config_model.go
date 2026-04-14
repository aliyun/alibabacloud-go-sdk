// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegistryCertConfig interface {
	dara.Model
	String() string
	GoString() string
	SetInsecure(v bool) *RegistryCertConfig
	GetInsecure() *bool
	SetRootCaCertBase64(v string) *RegistryCertConfig
	GetRootCaCertBase64() *string
}

type RegistryCertConfig struct {
	// 是否跳过TLS证书验证，设置为true时将不验证镜像仓库的证书
	//
	// example:
	//
	// false
	Insecure *bool `json:"insecure,omitempty" xml:"insecure,omitempty"`
	// 镜像仓库的根CA证书Base64编码，用于自签名证书的验证
	//
	// example:
	//
	// cm9vdF9jYV9jZXJ0X2Jhc2U2NA==
	RootCaCertBase64 *string `json:"rootCaCertBase64,omitempty" xml:"rootCaCertBase64,omitempty"`
}

func (s RegistryCertConfig) String() string {
	return dara.Prettify(s)
}

func (s RegistryCertConfig) GoString() string {
	return s.String()
}

func (s *RegistryCertConfig) GetInsecure() *bool {
	return s.Insecure
}

func (s *RegistryCertConfig) GetRootCaCertBase64() *string {
	return s.RootCaCertBase64
}

func (s *RegistryCertConfig) SetInsecure(v bool) *RegistryCertConfig {
	s.Insecure = &v
	return s
}

func (s *RegistryCertConfig) SetRootCaCertBase64(v string) *RegistryCertConfig {
	s.RootCaCertBase64 = &v
	return s
}

func (s *RegistryCertConfig) Validate() error {
	return dara.Validate(s)
}
