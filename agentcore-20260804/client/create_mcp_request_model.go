// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateMcpRequestBody) *CreateMcpRequest
	GetBody() *CreateMcpRequestBody
	SetClientToken(v string) *CreateMcpRequest
	GetClientToken() *string
}

type CreateMcpRequest struct {
	// The request body.
	Body *CreateMcpRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The client idempotency token.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426614174000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateMcpRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequest) GoString() string {
	return s.String()
}

func (s *CreateMcpRequest) GetBody() *CreateMcpRequestBody {
	return s.Body
}

func (s *CreateMcpRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateMcpRequest) SetBody(v *CreateMcpRequestBody) *CreateMcpRequest {
	s.Body = v
	return s
}

func (s *CreateMcpRequest) SetClientToken(v string) *CreateMcpRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMcpRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcpRequestBody struct {
	// The list of MCP service addresses.
	//
	// This parameter is required.
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// The backend authentication configuration. When enabled is set to true: for DIRECT_PROXY, specify directProxy (name/value). For HTTP_TO_MCP, specify the httpToMcp array (each item contains id/type/credential. For apiKey, position/name are also required). Multiple authentication objects are supported, and the first one is used as the default upstream credential. HTTP_TO_MCP credentials are merged into the securitySchemes of the Swagger specification.
	Auth *CreateMcpRequestBodyAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// A sample description that explains the purpose of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The MCP name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-mcp-server
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The MCP protocol.
	//
	// example:
	//
	// SSE
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The Swagger configuration. Specify this field if Type is set to HTTP_TO_MCP.
	//
	// example:
	//
	// {"type":"object"}
	SwaggerConfig *string `json:"swaggerConfig,omitempty" xml:"swaggerConfig,omitempty"`
	// The type.
	//
	// This parameter is required.
	//
	// example:
	//
	// DIRECT_PROXY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateMcpRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBody) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBody) GetAddresses() []*string {
	return s.Addresses
}

func (s *CreateMcpRequestBody) GetAuth() *CreateMcpRequestBodyAuth {
	return s.Auth
}

func (s *CreateMcpRequestBody) GetDescription() *string {
	return s.Description
}

func (s *CreateMcpRequestBody) GetName() *string {
	return s.Name
}

func (s *CreateMcpRequestBody) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateMcpRequestBody) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *CreateMcpRequestBody) GetType() *string {
	return s.Type
}

func (s *CreateMcpRequestBody) SetAddresses(v []*string) *CreateMcpRequestBody {
	s.Addresses = v
	return s
}

func (s *CreateMcpRequestBody) SetAuth(v *CreateMcpRequestBodyAuth) *CreateMcpRequestBody {
	s.Auth = v
	return s
}

func (s *CreateMcpRequestBody) SetDescription(v string) *CreateMcpRequestBody {
	s.Description = &v
	return s
}

func (s *CreateMcpRequestBody) SetName(v string) *CreateMcpRequestBody {
	s.Name = &v
	return s
}

func (s *CreateMcpRequestBody) SetProtocol(v string) *CreateMcpRequestBody {
	s.Protocol = &v
	return s
}

func (s *CreateMcpRequestBody) SetSwaggerConfig(v string) *CreateMcpRequestBody {
	s.SwaggerConfig = &v
	return s
}

func (s *CreateMcpRequestBody) SetType(v string) *CreateMcpRequestBody {
	s.Type = &v
	return s
}

func (s *CreateMcpRequestBody) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcpRequestBodyAuth struct {
	// The authentication configuration for direct proxy.
	DirectProxy *CreateMcpRequestBodyAuthDirectProxy `json:"directProxy,omitempty" xml:"directProxy,omitempty" type:"Struct"`
	// Specifies whether to enable authentication.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of HTTP_TO_MCP authentication configurations.
	HttpToMcp []*CreateMcpRequestBodyAuthHttpToMcp `json:"httpToMcp,omitempty" xml:"httpToMcp,omitempty" type:"Repeated"`
}

func (s CreateMcpRequestBodyAuth) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyAuth) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyAuth) GetDirectProxy() *CreateMcpRequestBodyAuthDirectProxy {
	return s.DirectProxy
}

func (s *CreateMcpRequestBodyAuth) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMcpRequestBodyAuth) GetHttpToMcp() []*CreateMcpRequestBodyAuthHttpToMcp {
	return s.HttpToMcp
}

func (s *CreateMcpRequestBodyAuth) SetDirectProxy(v *CreateMcpRequestBodyAuthDirectProxy) *CreateMcpRequestBodyAuth {
	s.DirectProxy = v
	return s
}

func (s *CreateMcpRequestBodyAuth) SetEnabled(v bool) *CreateMcpRequestBodyAuth {
	s.Enabled = &v
	return s
}

func (s *CreateMcpRequestBodyAuth) SetHttpToMcp(v []*CreateMcpRequestBodyAuthHttpToMcp) *CreateMcpRequestBodyAuth {
	s.HttpToMcp = v
	return s
}

func (s *CreateMcpRequestBodyAuth) Validate() error {
	if s.DirectProxy != nil {
		if err := s.DirectProxy.Validate(); err != nil {
			return err
		}
	}
	if s.HttpToMcp != nil {
		for _, item := range s.HttpToMcp {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateMcpRequestBodyAuthDirectProxy struct {
	// The name.
	//
	// example:
	//
	// mcp-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The authentication parameter value.
	//
	// example:
	//
	// example-credential
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateMcpRequestBodyAuthDirectProxy) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyAuthDirectProxy) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyAuthDirectProxy) GetName() *string {
	return s.Name
}

func (s *CreateMcpRequestBodyAuthDirectProxy) GetValue() *string {
	return s.Value
}

func (s *CreateMcpRequestBodyAuthDirectProxy) SetName(v string) *CreateMcpRequestBodyAuthDirectProxy {
	s.Name = &v
	return s
}

func (s *CreateMcpRequestBodyAuthDirectProxy) SetValue(v string) *CreateMcpRequestBodyAuthDirectProxy {
	s.Value = &v
	return s
}

func (s *CreateMcpRequestBodyAuthDirectProxy) Validate() error {
	return dara.Validate(s)
}

type CreateMcpRequestBodyAuthHttpToMcp struct {
	// The authentication credential.
	//
	// example:
	//
	// example-credential
	Credential *string `json:"credential,omitempty" xml:"credential,omitempty"`
	// The authentication scheme ID.
	//
	// example:
	//
	// mcp-1234567890abcdef
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name.
	//
	// example:
	//
	// mcp-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The position of the credential.
	//
	// example:
	//
	// header
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
	// The type.
	//
	// example:
	//
	// basic
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateMcpRequestBodyAuthHttpToMcp) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpRequestBodyAuthHttpToMcp) GoString() string {
	return s.String()
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) GetCredential() *string {
	return s.Credential
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) GetId() *string {
	return s.Id
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) GetName() *string {
	return s.Name
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) GetPosition() *string {
	return s.Position
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) GetType() *string {
	return s.Type
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) SetCredential(v string) *CreateMcpRequestBodyAuthHttpToMcp {
	s.Credential = &v
	return s
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) SetId(v string) *CreateMcpRequestBodyAuthHttpToMcp {
	s.Id = &v
	return s
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) SetName(v string) *CreateMcpRequestBodyAuthHttpToMcp {
	s.Name = &v
	return s
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) SetPosition(v string) *CreateMcpRequestBodyAuthHttpToMcp {
	s.Position = &v
	return s
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) SetType(v string) *CreateMcpRequestBodyAuthHttpToMcp {
	s.Type = &v
	return s
}

func (s *CreateMcpRequestBodyAuthHttpToMcp) Validate() error {
	return dara.Validate(s)
}
