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
	Insecure *bool `json:"insecure,omitempty" xml:"insecure,omitempty"`
	// example:
	//
	// cm9vdF9jYV9jZXJ0
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
