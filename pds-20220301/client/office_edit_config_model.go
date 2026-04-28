// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfficeEditConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *OfficeEditConfig
	GetEnabled() *bool
}

type OfficeEditConfig struct {
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s OfficeEditConfig) String() string {
	return dara.Prettify(s)
}

func (s OfficeEditConfig) GoString() string {
	return s.String()
}

func (s *OfficeEditConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *OfficeEditConfig) SetEnabled(v bool) *OfficeEditConfig {
	s.Enabled = &v
	return s
}

func (s *OfficeEditConfig) Validate() error {
	return dara.Validate(s)
}
