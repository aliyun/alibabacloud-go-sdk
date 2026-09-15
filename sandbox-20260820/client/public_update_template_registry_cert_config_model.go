// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicUpdateTemplateRegistryCertConfig interface {
	dara.Model
	String() string
	GoString() string
	SetInsecure(v bool) *PublicUpdateTemplateRegistryCertConfig
	GetInsecure() *bool
}

type PublicUpdateTemplateRegistryCertConfig struct {
	Insecure *bool `json:"insecure,omitempty" xml:"insecure,omitempty"`
}

func (s PublicUpdateTemplateRegistryCertConfig) String() string {
	return dara.Prettify(s)
}

func (s PublicUpdateTemplateRegistryCertConfig) GoString() string {
	return s.String()
}

func (s *PublicUpdateTemplateRegistryCertConfig) GetInsecure() *bool {
	return s.Insecure
}

func (s *PublicUpdateTemplateRegistryCertConfig) SetInsecure(v bool) *PublicUpdateTemplateRegistryCertConfig {
	s.Insecure = &v
	return s
}

func (s *PublicUpdateTemplateRegistryCertConfig) Validate() error {
	return dara.Validate(s)
}
