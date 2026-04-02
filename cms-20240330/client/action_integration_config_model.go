// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActionIntegrationConfig interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v []*string) *ActionIntegrationConfig
	GetActions() []*string
	SetEnabled(v bool) *ActionIntegrationConfig
	GetEnabled() *bool
}

type ActionIntegrationConfig struct {
	// 行动集成 ID 列表
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// 是否启用行动集成
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s ActionIntegrationConfig) String() string {
	return dara.Prettify(s)
}

func (s ActionIntegrationConfig) GoString() string {
	return s.String()
}

func (s *ActionIntegrationConfig) GetActions() []*string {
	return s.Actions
}

func (s *ActionIntegrationConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *ActionIntegrationConfig) SetActions(v []*string) *ActionIntegrationConfig {
	s.Actions = v
	return s
}

func (s *ActionIntegrationConfig) SetEnabled(v bool) *ActionIntegrationConfig {
	s.Enabled = &v
	return s
}

func (s *ActionIntegrationConfig) Validate() error {
	return dara.Validate(s)
}
