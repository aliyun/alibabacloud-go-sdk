// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProtocolSettings interface {
	dara.Model
	String() string
	GoString() string
	SetA2AAgentCard(v string) *ProtocolSettings
	GetA2AAgentCard() *string
	SetHeaders(v string) *ProtocolSettings
	GetHeaders() *string
	SetInputBodyJsonSchema(v string) *ProtocolSettings
	GetInputBodyJsonSchema() *string
	SetMethod(v string) *ProtocolSettings
	GetMethod() *string
	SetName(v string) *ProtocolSettings
	GetName() *string
	SetOutputBodyJsonSchema(v string) *ProtocolSettings
	GetOutputBodyJsonSchema() *string
	SetPath(v string) *ProtocolSettings
	GetPath() *string
	SetPathPrefix(v string) *ProtocolSettings
	GetPathPrefix() *string
	SetRequestContentType(v string) *ProtocolSettings
	GetRequestContentType() *string
	SetResponseContentType(v string) *ProtocolSettings
	GetResponseContentType() *string
}

type ProtocolSettings struct {
	// A2A Agent Card
	A2AAgentCard *string `json:"A2AAgentCard,omitempty" xml:"A2AAgentCard,omitempty"`
	// 请求头
	Headers *string `json:"headers,omitempty" xml:"headers,omitempty"`
	// 请求体JSON模式
	InputBodyJsonSchema *string `json:"inputBodyJsonSchema,omitempty" xml:"inputBodyJsonSchema,omitempty"`
	// HTTP方法
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// 协议名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 响应体JSON模式
	OutputBodyJsonSchema *string `json:"outputBodyJsonSchema,omitempty" xml:"outputBodyJsonSchema,omitempty"`
	// 协议路径
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 协议路径前缀
	PathPrefix *string `json:"pathPrefix,omitempty" xml:"pathPrefix,omitempty"`
	// 请求内容类型
	RequestContentType *string `json:"requestContentType,omitempty" xml:"requestContentType,omitempty"`
	// 响应内容类型
	ResponseContentType *string `json:"responseContentType,omitempty" xml:"responseContentType,omitempty"`
}

func (s ProtocolSettings) String() string {
	return dara.Prettify(s)
}

func (s ProtocolSettings) GoString() string {
	return s.String()
}

func (s *ProtocolSettings) GetA2AAgentCard() *string {
	return s.A2AAgentCard
}

func (s *ProtocolSettings) GetHeaders() *string {
	return s.Headers
}

func (s *ProtocolSettings) GetInputBodyJsonSchema() *string {
	return s.InputBodyJsonSchema
}

func (s *ProtocolSettings) GetMethod() *string {
	return s.Method
}

func (s *ProtocolSettings) GetName() *string {
	return s.Name
}

func (s *ProtocolSettings) GetOutputBodyJsonSchema() *string {
	return s.OutputBodyJsonSchema
}

func (s *ProtocolSettings) GetPath() *string {
	return s.Path
}

func (s *ProtocolSettings) GetPathPrefix() *string {
	return s.PathPrefix
}

func (s *ProtocolSettings) GetRequestContentType() *string {
	return s.RequestContentType
}

func (s *ProtocolSettings) GetResponseContentType() *string {
	return s.ResponseContentType
}

func (s *ProtocolSettings) SetA2AAgentCard(v string) *ProtocolSettings {
	s.A2AAgentCard = &v
	return s
}

func (s *ProtocolSettings) SetHeaders(v string) *ProtocolSettings {
	s.Headers = &v
	return s
}

func (s *ProtocolSettings) SetInputBodyJsonSchema(v string) *ProtocolSettings {
	s.InputBodyJsonSchema = &v
	return s
}

func (s *ProtocolSettings) SetMethod(v string) *ProtocolSettings {
	s.Method = &v
	return s
}

func (s *ProtocolSettings) SetName(v string) *ProtocolSettings {
	s.Name = &v
	return s
}

func (s *ProtocolSettings) SetOutputBodyJsonSchema(v string) *ProtocolSettings {
	s.OutputBodyJsonSchema = &v
	return s
}

func (s *ProtocolSettings) SetPath(v string) *ProtocolSettings {
	s.Path = &v
	return s
}

func (s *ProtocolSettings) SetPathPrefix(v string) *ProtocolSettings {
	s.PathPrefix = &v
	return s
}

func (s *ProtocolSettings) SetRequestContentType(v string) *ProtocolSettings {
	s.RequestContentType = &v
	return s
}

func (s *ProtocolSettings) SetResponseContentType(v string) *ProtocolSettings {
	s.ResponseContentType = &v
	return s
}

func (s *ProtocolSettings) Validate() error {
	return dara.Validate(s)
}
