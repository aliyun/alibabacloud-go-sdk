// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProbeMcpOriginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ProbeMcpOriginRequestBody) *ProbeMcpOriginRequest
	GetBody() *ProbeMcpOriginRequestBody
}

type ProbeMcpOriginRequest struct {
	// The scope of the origin server configuration must not exceed that of the creation interface. The control fields for this validation only restrict the validation scope and do not extend the origin server access configuration.
	//
	// This parameter is required.
	Body *ProbeMcpOriginRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s ProbeMcpOriginRequest) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginRequest) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginRequest) GetBody() *ProbeMcpOriginRequestBody {
	return s.Body
}

func (s *ProbeMcpOriginRequest) SetBody(v *ProbeMcpOriginRequestBody) *ProbeMcpOriginRequest {
	s.Body = v
	return s
}

func (s *ProbeMcpOriginRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ProbeMcpOriginRequestBody struct {
	// A single HTTP(S) URL that reuses the addresses of the origin server used during creation. Embedded credentials, batch addresses, or protected targets are not accepted.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://mcp.example.com/mcp
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// Reuses the enabled, directProxy, and httpToMcp origin server authentication fields and validations from the creation interface. When enabled is set to true, you must provide a configuration that matches the access type. codePackage or additional headers are not accepted.
	Auth *ProbeMcpOriginRequestBodyAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// Applicable only to HTTP. Reuses the JSON or YAML string semantics from the creation interface. This field is required when validating the definition. Remote documents or references are not automatically fetched.
	//
	// example:
	//
	// {"openapi":"3.0.3","info":{"title":"Inventory API","version":"1.0.0"},"paths":{"/inventory":{"get":{"operationId":"listInventory","responses":{"200":{"description":"OK"}}}}}}
	SwaggerConfig *string `json:"swaggerConfig,omitempty" xml:"swaggerConfig,omitempty"`
	// Valid values: DIRECT_PROXY or HTTP_TO_MCP. CODE_PACKAGE is not accepted.
	//
	// This parameter is required.
	//
	// example:
	//
	// DIRECT_PROXY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ProbeMcpOriginRequestBody) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginRequestBody) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginRequestBody) GetAddress() *string {
	return s.Address
}

func (s *ProbeMcpOriginRequestBody) GetAuth() *ProbeMcpOriginRequestBodyAuth {
	return s.Auth
}

func (s *ProbeMcpOriginRequestBody) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *ProbeMcpOriginRequestBody) GetType() *string {
	return s.Type
}

func (s *ProbeMcpOriginRequestBody) SetAddress(v string) *ProbeMcpOriginRequestBody {
	s.Address = &v
	return s
}

func (s *ProbeMcpOriginRequestBody) SetAuth(v *ProbeMcpOriginRequestBodyAuth) *ProbeMcpOriginRequestBody {
	s.Auth = v
	return s
}

func (s *ProbeMcpOriginRequestBody) SetSwaggerConfig(v string) *ProbeMcpOriginRequestBody {
	s.SwaggerConfig = &v
	return s
}

func (s *ProbeMcpOriginRequestBody) SetType(v string) *ProbeMcpOriginRequestBody {
	s.Type = &v
	return s
}

func (s *ProbeMcpOriginRequestBody) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ProbeMcpOriginRequestBodyAuth struct {
	// The backend authentication configuration for the direct proxy.
	DirectProxy *ProbeMcpOriginRequestBodyAuthDirectProxy `json:"directProxy,omitempty" xml:"directProxy,omitempty" type:"Struct"`
	// Specifies whether to enable the configuration.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of backend authentication configurations for HTTP-to-MCP.
	HttpToMcp []*ProbeMcpOriginRequestBodyAuthHttpToMcp `json:"httpToMcp,omitempty" xml:"httpToMcp,omitempty" type:"Repeated"`
}

func (s ProbeMcpOriginRequestBodyAuth) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginRequestBodyAuth) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginRequestBodyAuth) GetDirectProxy() *ProbeMcpOriginRequestBodyAuthDirectProxy {
	return s.DirectProxy
}

func (s *ProbeMcpOriginRequestBodyAuth) GetEnabled() *bool {
	return s.Enabled
}

func (s *ProbeMcpOriginRequestBodyAuth) GetHttpToMcp() []*ProbeMcpOriginRequestBodyAuthHttpToMcp {
	return s.HttpToMcp
}

func (s *ProbeMcpOriginRequestBodyAuth) SetDirectProxy(v *ProbeMcpOriginRequestBodyAuthDirectProxy) *ProbeMcpOriginRequestBodyAuth {
	s.DirectProxy = v
	return s
}

func (s *ProbeMcpOriginRequestBodyAuth) SetEnabled(v bool) *ProbeMcpOriginRequestBodyAuth {
	s.Enabled = &v
	return s
}

func (s *ProbeMcpOriginRequestBodyAuth) SetHttpToMcp(v []*ProbeMcpOriginRequestBodyAuthHttpToMcp) *ProbeMcpOriginRequestBodyAuth {
	s.HttpToMcp = v
	return s
}

func (s *ProbeMcpOriginRequestBodyAuth) Validate() error {
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

type ProbeMcpOriginRequestBodyAuthDirectProxy struct {
	// The name of the backend authentication header.
	//
	// example:
	//
	// Authorization
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the backend authentication header.
	//
	// example:
	//
	// example-api-key
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ProbeMcpOriginRequestBodyAuthDirectProxy) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginRequestBodyAuthDirectProxy) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginRequestBodyAuthDirectProxy) GetName() *string {
	return s.Name
}

func (s *ProbeMcpOriginRequestBodyAuthDirectProxy) GetValue() *string {
	return s.Value
}

func (s *ProbeMcpOriginRequestBodyAuthDirectProxy) SetName(v string) *ProbeMcpOriginRequestBodyAuthDirectProxy {
	s.Name = &v
	return s
}

func (s *ProbeMcpOriginRequestBodyAuthDirectProxy) SetValue(v string) *ProbeMcpOriginRequestBodyAuthDirectProxy {
	s.Value = &v
	return s
}

func (s *ProbeMcpOriginRequestBodyAuthDirectProxy) Validate() error {
	return dara.Validate(s)
}

type ProbeMcpOriginRequestBodyAuthHttpToMcp struct {
	// The credential for the backend authentication.
	//
	// example:
	//
	// example-api-key
	Credential *string `json:"credential,omitempty" xml:"credential,omitempty"`
	// The ID of the backend authentication configuration.
	//
	// example:
	//
	// api-key-auth
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the API key parameter.
	//
	// example:
	//
	// X-API-Key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The location where the API key is passed.
	//
	// example:
	//
	// header
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
	// The type of the backend authentication.
	//
	// example:
	//
	// apiKey
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ProbeMcpOriginRequestBodyAuthHttpToMcp) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginRequestBodyAuthHttpToMcp) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginRequestBodyAuthHttpToMcp) GetCredential() *string {
	return s.Credential
}

func (s *ProbeMcpOriginRequestBodyAuthHttpToMcp) GetId() *string {
	return s.Id
}

func (s *ProbeMcpOriginRequestBodyAuthHttpToMcp) GetName() *string {
	return s.Name
}

func (s *ProbeMcpOriginRequestBodyAuthHttpToMcp) GetPosition() *string {
	return s.Position
}

func (s *ProbeMcpOriginRequestBodyAuthHttpToMcp) GetType() *string {
	return s.Type
}

func (s *ProbeMcpOriginRequestBodyAuthHttpToMcp) SetCredential(v string) *ProbeMcpOriginRequestBodyAuthHttpToMcp {
	s.Credential = &v
	return s
}

func (s *ProbeMcpOriginRequestBodyAuthHttpToMcp) SetId(v string) *ProbeMcpOriginRequestBodyAuthHttpToMcp {
	s.Id = &v
	return s
}

func (s *ProbeMcpOriginRequestBodyAuthHttpToMcp) SetName(v string) *ProbeMcpOriginRequestBodyAuthHttpToMcp {
	s.Name = &v
	return s
}

func (s *ProbeMcpOriginRequestBodyAuthHttpToMcp) SetPosition(v string) *ProbeMcpOriginRequestBodyAuthHttpToMcp {
	s.Position = &v
	return s
}

func (s *ProbeMcpOriginRequestBodyAuthHttpToMcp) SetType(v string) *ProbeMcpOriginRequestBodyAuthHttpToMcp {
	s.Type = &v
	return s
}

func (s *ProbeMcpOriginRequestBodyAuthHttpToMcp) Validate() error {
	return dara.Validate(s)
}
