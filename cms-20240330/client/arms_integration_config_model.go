// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArmsIntegrationConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *ArmsIntegrationConfig
	GetEnabled() *bool
}

type ArmsIntegrationConfig struct {
	// 是否启用 ARMS 集成
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s ArmsIntegrationConfig) String() string {
	return dara.Prettify(s)
}

func (s ArmsIntegrationConfig) GoString() string {
	return s.String()
}

func (s *ArmsIntegrationConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *ArmsIntegrationConfig) SetEnabled(v bool) *ArmsIntegrationConfig {
	s.Enabled = &v
	return s
}

func (s *ArmsIntegrationConfig) Validate() error {
	return dara.Validate(s)
}
