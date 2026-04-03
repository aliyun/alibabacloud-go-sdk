// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMcpProxyConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetHooks(v []*Hook) *McpProxyConfiguration
	GetHooks() []*Hook
}

type McpProxyConfiguration struct {
	// MCP 代理的钩子函数配置列表，每个钩子可以在请求处理的不同阶段执行自定义逻辑
	Hooks []*Hook `json:"hooks" xml:"hooks" type:"Repeated"`
}

func (s McpProxyConfiguration) String() string {
	return dara.Prettify(s)
}

func (s McpProxyConfiguration) GoString() string {
	return s.String()
}

func (s *McpProxyConfiguration) GetHooks() []*Hook {
	return s.Hooks
}

func (s *McpProxyConfiguration) SetHooks(v []*Hook) *McpProxyConfiguration {
	s.Hooks = v
	return s
}

func (s *McpProxyConfiguration) Validate() error {
	if s.Hooks != nil {
		for _, item := range s.Hooks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
