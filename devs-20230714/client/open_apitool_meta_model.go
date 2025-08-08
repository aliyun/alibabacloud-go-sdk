// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenAPIToolMeta interface {
	dara.Model
	String() string
	GoString() string
	SetMethod(v string) *OpenAPIToolMeta
	GetMethod() *string
	SetPath(v string) *OpenAPIToolMeta
	GetPath() *string
	SetToolId(v string) *OpenAPIToolMeta
	GetToolId() *string
	SetToolName(v string) *OpenAPIToolMeta
	GetToolName() *string
}

type OpenAPIToolMeta struct {
	Method   *string `json:"method,omitempty" xml:"method,omitempty"`
	Path     *string `json:"path,omitempty" xml:"path,omitempty"`
	ToolId   *string `json:"toolId,omitempty" xml:"toolId,omitempty"`
	ToolName *string `json:"toolName,omitempty" xml:"toolName,omitempty"`
}

func (s OpenAPIToolMeta) String() string {
	return dara.Prettify(s)
}

func (s OpenAPIToolMeta) GoString() string {
	return s.String()
}

func (s *OpenAPIToolMeta) GetMethod() *string {
	return s.Method
}

func (s *OpenAPIToolMeta) GetPath() *string {
	return s.Path
}

func (s *OpenAPIToolMeta) GetToolId() *string {
	return s.ToolId
}

func (s *OpenAPIToolMeta) GetToolName() *string {
	return s.ToolName
}

func (s *OpenAPIToolMeta) SetMethod(v string) *OpenAPIToolMeta {
	s.Method = &v
	return s
}

func (s *OpenAPIToolMeta) SetPath(v string) *OpenAPIToolMeta {
	s.Path = &v
	return s
}

func (s *OpenAPIToolMeta) SetToolId(v string) *OpenAPIToolMeta {
	s.ToolId = &v
	return s
}

func (s *OpenAPIToolMeta) SetToolName(v string) *OpenAPIToolMeta {
	s.ToolName = &v
	return s
}

func (s *OpenAPIToolMeta) Validate() error {
	return dara.Validate(s)
}
