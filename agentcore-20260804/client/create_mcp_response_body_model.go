// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMcpResponseBody
	GetCode() *string
	SetData(v *CreateMcpResponseBodyData) *CreateMcpResponseBody
	GetData() *CreateMcpResponseBodyData
	SetHttpStatusCode(v int32) *CreateMcpResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateMcpResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMcpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMcpResponseBody
	GetSuccess() *bool
}

type CreateMcpResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *CreateMcpResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// Request processed successfully
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateMcpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMcpResponseBody) GetData() *CreateMcpResponseBodyData {
	return s.Data
}

func (s *CreateMcpResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateMcpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMcpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcpResponseBody) SetCode(v string) *CreateMcpResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMcpResponseBody) SetData(v *CreateMcpResponseBodyData) *CreateMcpResponseBody {
	s.Data = v
	return s
}

func (s *CreateMcpResponseBody) SetHttpStatusCode(v int32) *CreateMcpResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateMcpResponseBody) SetMessage(v string) *CreateMcpResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMcpResponseBody) SetRequestId(v string) *CreateMcpResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcpResponseBody) SetSuccess(v bool) *CreateMcpResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMcpResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcpResponseBodyData struct {
	// The list of MCP service addresses.
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// The backend authentication configuration. enabled indicates whether authentication is enabled. directProxy specifies custom authentication headers for direct proxy. httpToMcp specifies the OpenAPI credential list for HTTP_TO_MCP.
	Auth *CreateMcpResponseBodyDataAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// A sample description that explains the purpose of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The MCP server configuration.
	//
	// example:
	//
	// {"type":"object"}
	McpServerConfig *string `json:"mcpServerConfig,omitempty" xml:"mcpServerConfig,omitempty"`
	// The MCP server ID.
	//
	// example:
	//
	// mcp-1234567890abcdef
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// The name.
	//
	// example:
	//
	// mcp-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The MCP protocol.
	//
	// example:
	//
	// SSE
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The status.
	//
	// example:
	//
	// CREATING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The status reason.
	//
	// example:
	//
	// Resource processing completed
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// The Swagger configuration.
	//
	// example:
	//
	// {"type":"object"}
	SwaggerConfig *string `json:"swaggerConfig,omitempty" xml:"swaggerConfig,omitempty"`
	// The type.
	//
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateMcpResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyData) GetAddresses() []*string {
	return s.Addresses
}

func (s *CreateMcpResponseBodyData) GetAuth() *CreateMcpResponseBodyDataAuth {
	return s.Auth
}

func (s *CreateMcpResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateMcpResponseBodyData) GetMcpServerConfig() *string {
	return s.McpServerConfig
}

func (s *CreateMcpResponseBodyData) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *CreateMcpResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateMcpResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateMcpResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateMcpResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *CreateMcpResponseBodyData) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *CreateMcpResponseBodyData) GetType() *string {
	return s.Type
}

func (s *CreateMcpResponseBodyData) SetAddresses(v []*string) *CreateMcpResponseBodyData {
	s.Addresses = v
	return s
}

func (s *CreateMcpResponseBodyData) SetAuth(v *CreateMcpResponseBodyDataAuth) *CreateMcpResponseBodyData {
	s.Auth = v
	return s
}

func (s *CreateMcpResponseBodyData) SetDescription(v string) *CreateMcpResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetMcpServerConfig(v string) *CreateMcpResponseBodyData {
	s.McpServerConfig = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetMcpServerId(v string) *CreateMcpResponseBodyData {
	s.McpServerId = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetName(v string) *CreateMcpResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetProtocol(v string) *CreateMcpResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetStatus(v string) *CreateMcpResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetStatusReason(v string) *CreateMcpResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetSwaggerConfig(v string) *CreateMcpResponseBodyData {
	s.SwaggerConfig = &v
	return s
}

func (s *CreateMcpResponseBodyData) SetType(v string) *CreateMcpResponseBodyData {
	s.Type = &v
	return s
}

func (s *CreateMcpResponseBodyData) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcpResponseBodyDataAuth struct {
	// The authentication configuration for direct proxy.
	DirectProxy *CreateMcpResponseBodyDataAuthDirectProxy `json:"directProxy,omitempty" xml:"directProxy,omitempty" type:"Struct"`
	// Specifies whether to enable authentication.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of HTTP_TO_MCP authentication configurations.
	HttpToMcp []*CreateMcpResponseBodyDataAuthHttpToMcp `json:"httpToMcp,omitempty" xml:"httpToMcp,omitempty" type:"Repeated"`
}

func (s CreateMcpResponseBodyDataAuth) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataAuth) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataAuth) GetDirectProxy() *CreateMcpResponseBodyDataAuthDirectProxy {
	return s.DirectProxy
}

func (s *CreateMcpResponseBodyDataAuth) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateMcpResponseBodyDataAuth) GetHttpToMcp() []*CreateMcpResponseBodyDataAuthHttpToMcp {
	return s.HttpToMcp
}

func (s *CreateMcpResponseBodyDataAuth) SetDirectProxy(v *CreateMcpResponseBodyDataAuthDirectProxy) *CreateMcpResponseBodyDataAuth {
	s.DirectProxy = v
	return s
}

func (s *CreateMcpResponseBodyDataAuth) SetEnabled(v bool) *CreateMcpResponseBodyDataAuth {
	s.Enabled = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuth) SetHttpToMcp(v []*CreateMcpResponseBodyDataAuthHttpToMcp) *CreateMcpResponseBodyDataAuth {
	s.HttpToMcp = v
	return s
}

func (s *CreateMcpResponseBodyDataAuth) Validate() error {
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

type CreateMcpResponseBodyDataAuthDirectProxy struct {
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

func (s CreateMcpResponseBodyDataAuthDirectProxy) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataAuthDirectProxy) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataAuthDirectProxy) GetName() *string {
	return s.Name
}

func (s *CreateMcpResponseBodyDataAuthDirectProxy) GetValue() *string {
	return s.Value
}

func (s *CreateMcpResponseBodyDataAuthDirectProxy) SetName(v string) *CreateMcpResponseBodyDataAuthDirectProxy {
	s.Name = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuthDirectProxy) SetValue(v string) *CreateMcpResponseBodyDataAuthDirectProxy {
	s.Value = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuthDirectProxy) Validate() error {
	return dara.Validate(s)
}

type CreateMcpResponseBodyDataAuthHttpToMcp struct {
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

func (s CreateMcpResponseBodyDataAuthHttpToMcp) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpResponseBodyDataAuthHttpToMcp) GoString() string {
	return s.String()
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) GetCredential() *string {
	return s.Credential
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) GetId() *string {
	return s.Id
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) GetName() *string {
	return s.Name
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) GetPosition() *string {
	return s.Position
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) GetType() *string {
	return s.Type
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) SetCredential(v string) *CreateMcpResponseBodyDataAuthHttpToMcp {
	s.Credential = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) SetId(v string) *CreateMcpResponseBodyDataAuthHttpToMcp {
	s.Id = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) SetName(v string) *CreateMcpResponseBodyDataAuthHttpToMcp {
	s.Name = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) SetPosition(v string) *CreateMcpResponseBodyDataAuthHttpToMcp {
	s.Position = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) SetType(v string) *CreateMcpResponseBodyDataAuthHttpToMcp {
	s.Type = &v
	return s
}

func (s *CreateMcpResponseBodyDataAuthHttpToMcp) Validate() error {
	return dara.Validate(s)
}
