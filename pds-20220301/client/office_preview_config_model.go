// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfficePreviewConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *OfficePreviewConfig
	GetEnabled() *bool
}

type OfficePreviewConfig struct {
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s OfficePreviewConfig) String() string {
	return dara.Prettify(s)
}

func (s OfficePreviewConfig) GoString() string {
	return s.String()
}

func (s *OfficePreviewConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *OfficePreviewConfig) SetEnabled(v bool) *OfficePreviewConfig {
	s.Enabled = &v
	return s
}

func (s *OfficePreviewConfig) Validate() error {
	return dara.Validate(s)
}
