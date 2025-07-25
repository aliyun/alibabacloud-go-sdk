// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSandboxCacheConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *SandboxCacheConfig
	GetEnabled() *bool
}

type SandboxCacheConfig struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s SandboxCacheConfig) String() string {
	return dara.Prettify(s)
}

func (s SandboxCacheConfig) GoString() string {
	return s.String()
}

func (s *SandboxCacheConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *SandboxCacheConfig) SetEnabled(v bool) *SandboxCacheConfig {
	s.Enabled = &v
	return s
}

func (s *SandboxCacheConfig) Validate() error {
	return dara.Validate(s)
}
