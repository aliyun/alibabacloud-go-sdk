// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBoundTool interface {
	dara.Model
	String() string
	GoString() string
	SetMethod(v string) *BoundTool
	GetMethod() *string
	SetPath(v string) *BoundTool
	GetPath() *string
	SetToolName(v string) *BoundTool
	GetToolName() *string
}

type BoundTool struct {
	// 绑定的 HTTP 请求方法，支持：GET、PUT、POST、PATCH、DELETE、OPTIONS
	//
	// example:
	//
	// POST
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// 绑定的 URL 路径，用于路由匹配
	//
	// example:
	//
	// /api/v1/tools
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 要绑定的工具名称
	//
	// example:
	//
	// my-tool
	ToolName *string `json:"toolName,omitempty" xml:"toolName,omitempty"`
}

func (s BoundTool) String() string {
	return dara.Prettify(s)
}

func (s BoundTool) GoString() string {
	return s.String()
}

func (s *BoundTool) GetMethod() *string {
	return s.Method
}

func (s *BoundTool) GetPath() *string {
	return s.Path
}

func (s *BoundTool) GetToolName() *string {
	return s.ToolName
}

func (s *BoundTool) SetMethod(v string) *BoundTool {
	s.Method = &v
	return s
}

func (s *BoundTool) SetPath(v string) *BoundTool {
	s.Path = &v
	return s
}

func (s *BoundTool) SetToolName(v string) *BoundTool {
	s.ToolName = &v
	return s
}

func (s *BoundTool) Validate() error {
	return dara.Validate(s)
}
