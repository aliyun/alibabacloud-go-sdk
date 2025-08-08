// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMCPToolMeta interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *MCPToolMeta
	GetDescription() *string
	SetInputSchema(v map[string]interface{}) *MCPToolMeta
	GetInputSchema() map[string]interface{}
	SetName(v string) *MCPToolMeta
	GetName() *string
}

type MCPToolMeta struct {
	// example:
	//
	// Here is an example
	Description *string                `json:"description,omitempty" xml:"description,omitempty"`
	InputSchema map[string]interface{} `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	// example:
	//
	// demo-tool
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s MCPToolMeta) String() string {
	return dara.Prettify(s)
}

func (s MCPToolMeta) GoString() string {
	return s.String()
}

func (s *MCPToolMeta) GetDescription() *string {
	return s.Description
}

func (s *MCPToolMeta) GetInputSchema() map[string]interface{} {
	return s.InputSchema
}

func (s *MCPToolMeta) GetName() *string {
	return s.Name
}

func (s *MCPToolMeta) SetDescription(v string) *MCPToolMeta {
	s.Description = &v
	return s
}

func (s *MCPToolMeta) SetInputSchema(v map[string]interface{}) *MCPToolMeta {
	s.InputSchema = v
	return s
}

func (s *MCPToolMeta) SetName(v string) *MCPToolMeta {
	s.Name = &v
	return s
}

func (s *MCPToolMeta) Validate() error {
	return dara.Validate(s)
}
