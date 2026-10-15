// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateRegistryCertConfig interface {
	dara.Model
	String() string
	GoString() string
	SetInsecure(v bool) *CreateTemplateRegistryCertConfig
	GetInsecure() *bool
}

type CreateTemplateRegistryCertConfig struct {
	// Specifies whether to skip HTTPS certificate validation for the image repository. Enable this option only for private repositories that use self-signed certificates.
	Insecure *bool `json:"insecure,omitempty" xml:"insecure,omitempty"`
}

func (s CreateTemplateRegistryCertConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRegistryCertConfig) GoString() string {
	return s.String()
}

func (s *CreateTemplateRegistryCertConfig) GetInsecure() *bool {
	return s.Insecure
}

func (s *CreateTemplateRegistryCertConfig) SetInsecure(v bool) *CreateTemplateRegistryCertConfig {
	s.Insecure = &v
	return s
}

func (s *CreateTemplateRegistryCertConfig) Validate() error {
	return dara.Validate(s)
}
