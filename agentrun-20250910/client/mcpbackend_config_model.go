// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMCPBackendConfig interface {
	dara.Model
	String() string
	GoString() string
	SetScene(v string) *MCPBackendConfig
	GetScene() *string
	SetServices(v []*MCPServiceConfig) *MCPBackendConfig
	GetServices() []*MCPServiceConfig
}

type MCPBackendConfig struct {
	Scene    *string             `json:"scene,omitempty" xml:"scene,omitempty"`
	Services []*MCPServiceConfig `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s MCPBackendConfig) String() string {
	return dara.Prettify(s)
}

func (s MCPBackendConfig) GoString() string {
	return s.String()
}

func (s *MCPBackendConfig) GetScene() *string {
	return s.Scene
}

func (s *MCPBackendConfig) GetServices() []*MCPServiceConfig {
	return s.Services
}

func (s *MCPBackendConfig) SetScene(v string) *MCPBackendConfig {
	s.Scene = &v
	return s
}

func (s *MCPBackendConfig) SetServices(v []*MCPServiceConfig) *MCPBackendConfig {
	s.Services = v
	return s
}

func (s *MCPBackendConfig) Validate() error {
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
