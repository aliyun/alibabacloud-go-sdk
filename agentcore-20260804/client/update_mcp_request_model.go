// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateMcpRequestBody) *UpdateMcpRequest
	GetBody() *UpdateMcpRequestBody
	SetClientToken(v string) *UpdateMcpRequest
	GetClientToken() *string
}

type UpdateMcpRequest struct {
	// The request body.
	Body *UpdateMcpRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The client idempotency token.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426614174000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateMcpRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequest) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequest) GetBody() *UpdateMcpRequestBody {
	return s.Body
}

func (s *UpdateMcpRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateMcpRequest) SetBody(v *UpdateMcpRequestBody) *UpdateMcpRequest {
	s.Body = v
	return s
}

func (s *UpdateMcpRequest) SetClientToken(v string) *UpdateMcpRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateMcpRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMcpRequestBody struct {
	// The list of MCP service addresses.
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// The backend authentication configuration. When enabled=true: for DIRECT_PROXY, specify directProxy (name/value). For HTTP_TO_MCP, specify the httpToMcp array (each item contains id/type/credential; apiKey also requires position/name). Multiple authentication objects are supported, and the first one is the default upstream credential. HTTP_TO_MCP credentials are merged into the securitySchemes of the Swagger specification.
	Auth *UpdateMcpRequestBodyAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// A sample description that explains the purpose of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The Swagger configuration. Specify this field if Type is set to HTTP_TO_MCP.
	//
	// example:
	//
	// {"type":"object"}
	SwaggerConfig *string `json:"swaggerConfig,omitempty" xml:"swaggerConfig,omitempty"`
}

func (s UpdateMcpRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBody) GetAddresses() []*string {
	return s.Addresses
}

func (s *UpdateMcpRequestBody) GetAuth() *UpdateMcpRequestBodyAuth {
	return s.Auth
}

func (s *UpdateMcpRequestBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateMcpRequestBody) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *UpdateMcpRequestBody) SetAddresses(v []*string) *UpdateMcpRequestBody {
	s.Addresses = v
	return s
}

func (s *UpdateMcpRequestBody) SetAuth(v *UpdateMcpRequestBodyAuth) *UpdateMcpRequestBody {
	s.Auth = v
	return s
}

func (s *UpdateMcpRequestBody) SetDescription(v string) *UpdateMcpRequestBody {
	s.Description = &v
	return s
}

func (s *UpdateMcpRequestBody) SetSwaggerConfig(v string) *UpdateMcpRequestBody {
	s.SwaggerConfig = &v
	return s
}

func (s *UpdateMcpRequestBody) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMcpRequestBodyAuth struct {
	// The direct proxy authentication configuration.
	DirectProxy *UpdateMcpRequestBodyAuthDirectProxy `json:"directProxy,omitempty" xml:"directProxy,omitempty" type:"Struct"`
	// Specifies whether to enable authentication.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of HTTP_TO_MCP authentication configurations.
	HttpToMcp []*UpdateMcpRequestBodyAuthHttpToMcp `json:"httpToMcp,omitempty" xml:"httpToMcp,omitempty" type:"Repeated"`
}

func (s UpdateMcpRequestBodyAuth) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyAuth) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyAuth) GetDirectProxy() *UpdateMcpRequestBodyAuthDirectProxy {
	return s.DirectProxy
}

func (s *UpdateMcpRequestBodyAuth) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMcpRequestBodyAuth) GetHttpToMcp() []*UpdateMcpRequestBodyAuthHttpToMcp {
	return s.HttpToMcp
}

func (s *UpdateMcpRequestBodyAuth) SetDirectProxy(v *UpdateMcpRequestBodyAuthDirectProxy) *UpdateMcpRequestBodyAuth {
	s.DirectProxy = v
	return s
}

func (s *UpdateMcpRequestBodyAuth) SetEnabled(v bool) *UpdateMcpRequestBodyAuth {
	s.Enabled = &v
	return s
}

func (s *UpdateMcpRequestBodyAuth) SetHttpToMcp(v []*UpdateMcpRequestBodyAuthHttpToMcp) *UpdateMcpRequestBodyAuth {
	s.HttpToMcp = v
	return s
}

func (s *UpdateMcpRequestBodyAuth) Validate() error {
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

type UpdateMcpRequestBodyAuthDirectProxy struct {
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

func (s UpdateMcpRequestBodyAuthDirectProxy) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyAuthDirectProxy) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyAuthDirectProxy) GetName() *string {
	return s.Name
}

func (s *UpdateMcpRequestBodyAuthDirectProxy) GetValue() *string {
	return s.Value
}

func (s *UpdateMcpRequestBodyAuthDirectProxy) SetName(v string) *UpdateMcpRequestBodyAuthDirectProxy {
	s.Name = &v
	return s
}

func (s *UpdateMcpRequestBodyAuthDirectProxy) SetValue(v string) *UpdateMcpRequestBodyAuthDirectProxy {
	s.Value = &v
	return s
}

func (s *UpdateMcpRequestBodyAuthDirectProxy) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpRequestBodyAuthHttpToMcp struct {
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

func (s UpdateMcpRequestBodyAuthHttpToMcp) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpRequestBodyAuthHttpToMcp) GoString() string {
	return s.String()
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) GetCredential() *string {
	return s.Credential
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) GetId() *string {
	return s.Id
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) GetName() *string {
	return s.Name
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) GetPosition() *string {
	return s.Position
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) GetType() *string {
	return s.Type
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) SetCredential(v string) *UpdateMcpRequestBodyAuthHttpToMcp {
	s.Credential = &v
	return s
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) SetId(v string) *UpdateMcpRequestBodyAuthHttpToMcp {
	s.Id = &v
	return s
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) SetName(v string) *UpdateMcpRequestBodyAuthHttpToMcp {
	s.Name = &v
	return s
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) SetPosition(v string) *UpdateMcpRequestBodyAuthHttpToMcp {
	s.Position = &v
	return s
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) SetType(v string) *UpdateMcpRequestBodyAuthHttpToMcp {
	s.Type = &v
	return s
}

func (s *UpdateMcpRequestBodyAuthHttpToMcp) Validate() error {
	return dara.Validate(s)
}
