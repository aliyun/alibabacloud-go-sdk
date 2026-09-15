// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicUpdateTemplateEnvdInjectAction interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *PublicUpdateTemplateEnvdInjectAction
	GetEnabled() *bool
}

type PublicUpdateTemplateEnvdInjectAction struct {
	// Specifies whether envd injection is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s PublicUpdateTemplateEnvdInjectAction) String() string {
	return dara.Prettify(s)
}

func (s PublicUpdateTemplateEnvdInjectAction) GoString() string {
	return s.String()
}

func (s *PublicUpdateTemplateEnvdInjectAction) GetEnabled() *bool {
	return s.Enabled
}

func (s *PublicUpdateTemplateEnvdInjectAction) SetEnabled(v bool) *PublicUpdateTemplateEnvdInjectAction {
	s.Enabled = &v
	return s
}

func (s *PublicUpdateTemplateEnvdInjectAction) Validate() error {
	return dara.Validate(s)
}
