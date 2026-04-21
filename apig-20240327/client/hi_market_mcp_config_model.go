// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketMcpConfig interface {
	dara.Model
	String() string
	GoString() string
	SetMcpServerConfig(v *HiMarketMcpConfigMcpServerConfig) *HiMarketMcpConfig
	GetMcpServerConfig() *HiMarketMcpConfigMcpServerConfig
	SetMcpServerName(v string) *HiMarketMcpConfig
	GetMcpServerName() *string
	SetMeta(v *HiMarketMcpConfigMeta) *HiMarketMcpConfig
	GetMeta() *HiMarketMcpConfigMeta
	SetTools(v string) *HiMarketMcpConfig
	GetTools() *string
}

type HiMarketMcpConfig struct {
	McpServerConfig *HiMarketMcpConfigMcpServerConfig `json:"mcpServerConfig,omitempty" xml:"mcpServerConfig,omitempty" type:"Struct"`
	McpServerName   *string                           `json:"mcpServerName,omitempty" xml:"mcpServerName,omitempty"`
	Meta            *HiMarketMcpConfigMeta            `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	Tools           *string                           `json:"tools,omitempty" xml:"tools,omitempty"`
}

func (s HiMarketMcpConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketMcpConfig) GoString() string {
	return s.String()
}

func (s *HiMarketMcpConfig) GetMcpServerConfig() *HiMarketMcpConfigMcpServerConfig {
	return s.McpServerConfig
}

func (s *HiMarketMcpConfig) GetMcpServerName() *string {
	return s.McpServerName
}

func (s *HiMarketMcpConfig) GetMeta() *HiMarketMcpConfigMeta {
	return s.Meta
}

func (s *HiMarketMcpConfig) GetTools() *string {
	return s.Tools
}

func (s *HiMarketMcpConfig) SetMcpServerConfig(v *HiMarketMcpConfigMcpServerConfig) *HiMarketMcpConfig {
	s.McpServerConfig = v
	return s
}

func (s *HiMarketMcpConfig) SetMcpServerName(v string) *HiMarketMcpConfig {
	s.McpServerName = &v
	return s
}

func (s *HiMarketMcpConfig) SetMeta(v *HiMarketMcpConfigMeta) *HiMarketMcpConfig {
	s.Meta = v
	return s
}

func (s *HiMarketMcpConfig) SetTools(v string) *HiMarketMcpConfig {
	s.Tools = &v
	return s
}

func (s *HiMarketMcpConfig) Validate() error {
	if s.McpServerConfig != nil {
		if err := s.McpServerConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Meta != nil {
		if err := s.Meta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HiMarketMcpConfigMcpServerConfig struct {
	Domains []*HiMarketDomain `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
	Path    *string           `json:"path,omitempty" xml:"path,omitempty"`
}

func (s HiMarketMcpConfigMcpServerConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketMcpConfigMcpServerConfig) GoString() string {
	return s.String()
}

func (s *HiMarketMcpConfigMcpServerConfig) GetDomains() []*HiMarketDomain {
	return s.Domains
}

func (s *HiMarketMcpConfigMcpServerConfig) GetPath() *string {
	return s.Path
}

func (s *HiMarketMcpConfigMcpServerConfig) SetDomains(v []*HiMarketDomain) *HiMarketMcpConfigMcpServerConfig {
	s.Domains = v
	return s
}

func (s *HiMarketMcpConfigMcpServerConfig) SetPath(v string) *HiMarketMcpConfigMcpServerConfig {
	s.Path = &v
	return s
}

func (s *HiMarketMcpConfigMcpServerConfig) Validate() error {
	if s.Domains != nil {
		for _, item := range s.Domains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HiMarketMcpConfigMeta struct {
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s HiMarketMcpConfigMeta) String() string {
	return dara.Prettify(s)
}

func (s HiMarketMcpConfigMeta) GoString() string {
	return s.String()
}

func (s *HiMarketMcpConfigMeta) GetProtocol() *string {
	return s.Protocol
}

func (s *HiMarketMcpConfigMeta) SetProtocol(v string) *HiMarketMcpConfigMeta {
	s.Protocol = &v
	return s
}

func (s *HiMarketMcpConfigMeta) Validate() error {
	return dara.Validate(s)
}
