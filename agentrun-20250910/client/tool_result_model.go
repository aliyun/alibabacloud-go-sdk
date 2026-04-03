// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToolResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ToolResult
	GetCode() *string
	SetData(v *Tool) *ToolResult
	GetData() *Tool
	SetRequestId(v string) *ToolResult
	GetRequestId() *string
}

type ToolResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 工具的详细信息
	//
	// example:
	//
	// {}
	Data *Tool `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ToolResult) String() string {
	return dara.Prettify(s)
}

func (s ToolResult) GoString() string {
	return s.String()
}

func (s *ToolResult) GetCode() *string {
	return s.Code
}

func (s *ToolResult) GetData() *Tool {
	return s.Data
}

func (s *ToolResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ToolResult) SetCode(v string) *ToolResult {
	s.Code = &v
	return s
}

func (s *ToolResult) SetData(v *Tool) *ToolResult {
	s.Data = v
	return s
}

func (s *ToolResult) SetRequestId(v string) *ToolResult {
	s.RequestId = &v
	return s
}

func (s *ToolResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
