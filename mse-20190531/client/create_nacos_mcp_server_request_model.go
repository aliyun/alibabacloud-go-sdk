// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNacosMcpServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateNacosMcpServerRequest
	GetAcceptLanguage() *string
	SetEncryptToolSpec(v bool) *CreateNacosMcpServerRequest
	GetEncryptToolSpec() *bool
	SetEndpointSpecification(v string) *CreateNacosMcpServerRequest
	GetEndpointSpecification() *string
	SetInstanceId(v string) *CreateNacosMcpServerRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *CreateNacosMcpServerRequest
	GetNamespaceId() *string
	SetServerName(v string) *CreateNacosMcpServerRequest
	GetServerName() *string
	SetServerSpecification(v string) *CreateNacosMcpServerRequest
	GetServerSpecification() *string
	SetToolSpecification(v string) *CreateNacosMcpServerRequest
	GetToolSpecification() *string
	SetYamlConfig(v string) *CreateNacosMcpServerRequest
	GetYamlConfig() *string
}

type CreateNacosMcpServerRequest struct {
	// example:
	//
	// zh
	AcceptLanguage  *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	EncryptToolSpec *bool   `json:"EncryptToolSpec,omitempty" xml:"EncryptToolSpec,omitempty"`
	// example:
	//
	// {"type":"REF","data":{"namespaceId":"public","groupName":"mcp-endpoints","serviceName":"mcp-demo"}}
	EndpointSpecification *string `json:"EndpointSpecification,omitempty" xml:"EndpointSpecification,omitempty"`
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// public
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// mcp-demo
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
	// example:
	//
	// {"description":"a mcp demo service","capabilities":[],"allVersions":[{"version":"1.0.0","release_date":"2025-08-13T01:58:42Z","is_latest":true}],"remoteServerConfig":{"exportPath":"/sse"},"enabled":true,"versionDetail":{"version":"1.0.1"},"backendEndpoints":[{"path":"/sse","address":"127.0.0.1","port":8080}],"protocol":"mcp-sse","name":"mcp-demo","frontProtocol":"mcp-sse"}
	ServerSpecification *string `json:"ServerSpecification,omitempty" xml:"ServerSpecification,omitempty"`
	// example:
	//
	// {"tools":[{"name":"demo-tool","description":"a demo tool","inputSchema":{"type":"object","properties":{"name":{"type":"string","description":"name"}}}}],"toolsMeta":{"demo-tool":{"enabled":true}}}
	ToolSpecification *string `json:"ToolSpecification,omitempty" xml:"ToolSpecification,omitempty"`
	// example:
	//
	// allowTools:
	//
	// - demo-tool
	//
	// securityScheme: {}
	//
	// server:
	//
	//   name: mcp-demo
	//
	// tools:
	//
	// - args:
	//
	//   - name: name
	//
	//     description: name
	//
	//     type: string
	//
	//   description: a demo tool
	//
	//   name: demo-tool
	YamlConfig *string `json:"YamlConfig,omitempty" xml:"YamlConfig,omitempty"`
}

func (s CreateNacosMcpServerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNacosMcpServerRequest) GoString() string {
	return s.String()
}

func (s *CreateNacosMcpServerRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateNacosMcpServerRequest) GetEncryptToolSpec() *bool {
	return s.EncryptToolSpec
}

func (s *CreateNacosMcpServerRequest) GetEndpointSpecification() *string {
	return s.EndpointSpecification
}

func (s *CreateNacosMcpServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateNacosMcpServerRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateNacosMcpServerRequest) GetServerName() *string {
	return s.ServerName
}

func (s *CreateNacosMcpServerRequest) GetServerSpecification() *string {
	return s.ServerSpecification
}

func (s *CreateNacosMcpServerRequest) GetToolSpecification() *string {
	return s.ToolSpecification
}

func (s *CreateNacosMcpServerRequest) GetYamlConfig() *string {
	return s.YamlConfig
}

func (s *CreateNacosMcpServerRequest) SetAcceptLanguage(v string) *CreateNacosMcpServerRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateNacosMcpServerRequest) SetEncryptToolSpec(v bool) *CreateNacosMcpServerRequest {
	s.EncryptToolSpec = &v
	return s
}

func (s *CreateNacosMcpServerRequest) SetEndpointSpecification(v string) *CreateNacosMcpServerRequest {
	s.EndpointSpecification = &v
	return s
}

func (s *CreateNacosMcpServerRequest) SetInstanceId(v string) *CreateNacosMcpServerRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNacosMcpServerRequest) SetNamespaceId(v string) *CreateNacosMcpServerRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateNacosMcpServerRequest) SetServerName(v string) *CreateNacosMcpServerRequest {
	s.ServerName = &v
	return s
}

func (s *CreateNacosMcpServerRequest) SetServerSpecification(v string) *CreateNacosMcpServerRequest {
	s.ServerSpecification = &v
	return s
}

func (s *CreateNacosMcpServerRequest) SetToolSpecification(v string) *CreateNacosMcpServerRequest {
	s.ToolSpecification = &v
	return s
}

func (s *CreateNacosMcpServerRequest) SetYamlConfig(v string) *CreateNacosMcpServerRequest {
	s.YamlConfig = &v
	return s
}

func (s *CreateNacosMcpServerRequest) Validate() error {
	return dara.Validate(s)
}
