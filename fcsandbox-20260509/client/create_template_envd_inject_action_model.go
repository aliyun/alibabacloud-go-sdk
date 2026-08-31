// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateEnvdInjectAction interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *CreateTemplateEnvdInjectAction
	GetEnabled() *bool
}

type CreateTemplateEnvdInjectAction struct {
	// Specifies whether to enable envd injection.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreateTemplateEnvdInjectAction) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateEnvdInjectAction) GoString() string {
	return s.String()
}

func (s *CreateTemplateEnvdInjectAction) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateTemplateEnvdInjectAction) SetEnabled(v bool) *CreateTemplateEnvdInjectAction {
	s.Enabled = &v
	return s
}

func (s *CreateTemplateEnvdInjectAction) Validate() error {
	return dara.Validate(s)
}
