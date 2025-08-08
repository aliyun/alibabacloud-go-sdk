// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTool interface {
	dara.Model
	String() string
	GoString() string
	SetMethod(v string) *Tool
	GetMethod() *string
	SetPath(v string) *Tool
	GetPath() *string
	SetToolId(v string) *Tool
	GetToolId() *string
	SetToolName(v string) *Tool
	GetToolName() *string
}

type Tool struct {
	Method   *string `json:"method,omitempty" xml:"method,omitempty"`
	Path     *string `json:"path,omitempty" xml:"path,omitempty"`
	ToolId   *string `json:"toolId,omitempty" xml:"toolId,omitempty"`
	ToolName *string `json:"toolName,omitempty" xml:"toolName,omitempty"`
}

func (s Tool) String() string {
	return dara.Prettify(s)
}

func (s Tool) GoString() string {
	return s.String()
}

func (s *Tool) GetMethod() *string {
	return s.Method
}

func (s *Tool) GetPath() *string {
	return s.Path
}

func (s *Tool) GetToolId() *string {
	return s.ToolId
}

func (s *Tool) GetToolName() *string {
	return s.ToolName
}

func (s *Tool) SetMethod(v string) *Tool {
	s.Method = &v
	return s
}

func (s *Tool) SetPath(v string) *Tool {
	s.Path = &v
	return s
}

func (s *Tool) SetToolId(v string) *Tool {
	s.ToolId = &v
	return s
}

func (s *Tool) SetToolName(v string) *Tool {
	s.ToolName = &v
	return s
}

func (s *Tool) Validate() error {
	return dara.Validate(s)
}
