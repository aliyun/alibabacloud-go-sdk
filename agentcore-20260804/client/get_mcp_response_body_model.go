// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMcpResponseBody
	GetCode() *string
	SetData(v *GetMcpResponseBodyData) *GetMcpResponseBody
	GetData() *GetMcpResponseBodyData
	SetHttpStatusCode(v int32) *GetMcpResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMcpResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMcpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMcpResponseBody
	GetSuccess() *bool
}

type GetMcpResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetMcpResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetMcpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMcpResponseBody) GetData() *GetMcpResponseBodyData {
	return s.Data
}

func (s *GetMcpResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMcpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMcpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcpResponseBody) SetCode(v string) *GetMcpResponseBody {
	s.Code = &v
	return s
}

func (s *GetMcpResponseBody) SetData(v *GetMcpResponseBodyData) *GetMcpResponseBody {
	s.Data = v
	return s
}

func (s *GetMcpResponseBody) SetHttpStatusCode(v int32) *GetMcpResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMcpResponseBody) SetMessage(v string) *GetMcpResponseBody {
	s.Message = &v
	return s
}

func (s *GetMcpResponseBody) SetRequestId(v string) *GetMcpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcpResponseBody) SetSuccess(v bool) *GetMcpResponseBody {
	s.Success = &v
	return s
}

func (s *GetMcpResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpResponseBodyData struct {
	// The list of MCP service addresses.
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// The backend authentication configuration. enabled indicates whether authentication is enabled. directProxy specifies the custom authentication header for direct proxy. httpToMcp specifies the list of OpenAPI credentials for HTTP_TO_MCP.
	Auth *GetMcpResponseBodyDataAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
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
	// The MCP service ID.
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
	// The MCP service access URL.
	//
	// example:
	//
	// https://example.com/artifacts/example.zip
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetMcpResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyData) GetAddresses() []*string {
	return s.Addresses
}

func (s *GetMcpResponseBodyData) GetAuth() *GetMcpResponseBodyDataAuth {
	return s.Auth
}

func (s *GetMcpResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetMcpResponseBodyData) GetMcpServerConfig() *string {
	return s.McpServerConfig
}

func (s *GetMcpResponseBodyData) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *GetMcpResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetMcpResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *GetMcpResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetMcpResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetMcpResponseBodyData) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *GetMcpResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetMcpResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetMcpResponseBodyData) SetAddresses(v []*string) *GetMcpResponseBodyData {
	s.Addresses = v
	return s
}

func (s *GetMcpResponseBodyData) SetAuth(v *GetMcpResponseBodyDataAuth) *GetMcpResponseBodyData {
	s.Auth = v
	return s
}

func (s *GetMcpResponseBodyData) SetDescription(v string) *GetMcpResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetMcpResponseBodyData) SetMcpServerConfig(v string) *GetMcpResponseBodyData {
	s.McpServerConfig = &v
	return s
}

func (s *GetMcpResponseBodyData) SetMcpServerId(v string) *GetMcpResponseBodyData {
	s.McpServerId = &v
	return s
}

func (s *GetMcpResponseBodyData) SetName(v string) *GetMcpResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMcpResponseBodyData) SetProtocol(v string) *GetMcpResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetMcpResponseBodyData) SetStatus(v string) *GetMcpResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMcpResponseBodyData) SetStatusReason(v string) *GetMcpResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *GetMcpResponseBodyData) SetSwaggerConfig(v string) *GetMcpResponseBodyData {
	s.SwaggerConfig = &v
	return s
}

func (s *GetMcpResponseBodyData) SetType(v string) *GetMcpResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetMcpResponseBodyData) SetUrl(v string) *GetMcpResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetMcpResponseBodyData) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpResponseBodyDataAuth struct {
	// The direct proxy authentication configuration.
	DirectProxy *GetMcpResponseBodyDataAuthDirectProxy `json:"directProxy,omitempty" xml:"directProxy,omitempty" type:"Struct"`
	// Indicates whether authentication is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of HTTP_TO_MCP authentication configurations.
	HttpToMcp []*GetMcpResponseBodyDataAuthHttpToMcp `json:"httpToMcp,omitempty" xml:"httpToMcp,omitempty" type:"Repeated"`
}

func (s GetMcpResponseBodyDataAuth) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataAuth) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataAuth) GetDirectProxy() *GetMcpResponseBodyDataAuthDirectProxy {
	return s.DirectProxy
}

func (s *GetMcpResponseBodyDataAuth) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetMcpResponseBodyDataAuth) GetHttpToMcp() []*GetMcpResponseBodyDataAuthHttpToMcp {
	return s.HttpToMcp
}

func (s *GetMcpResponseBodyDataAuth) SetDirectProxy(v *GetMcpResponseBodyDataAuthDirectProxy) *GetMcpResponseBodyDataAuth {
	s.DirectProxy = v
	return s
}

func (s *GetMcpResponseBodyDataAuth) SetEnabled(v bool) *GetMcpResponseBodyDataAuth {
	s.Enabled = &v
	return s
}

func (s *GetMcpResponseBodyDataAuth) SetHttpToMcp(v []*GetMcpResponseBodyDataAuthHttpToMcp) *GetMcpResponseBodyDataAuth {
	s.HttpToMcp = v
	return s
}

func (s *GetMcpResponseBodyDataAuth) Validate() error {
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

type GetMcpResponseBodyDataAuthDirectProxy struct {
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

func (s GetMcpResponseBodyDataAuthDirectProxy) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataAuthDirectProxy) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataAuthDirectProxy) GetName() *string {
	return s.Name
}

func (s *GetMcpResponseBodyDataAuthDirectProxy) GetValue() *string {
	return s.Value
}

func (s *GetMcpResponseBodyDataAuthDirectProxy) SetName(v string) *GetMcpResponseBodyDataAuthDirectProxy {
	s.Name = &v
	return s
}

func (s *GetMcpResponseBodyDataAuthDirectProxy) SetValue(v string) *GetMcpResponseBodyDataAuthDirectProxy {
	s.Value = &v
	return s
}

func (s *GetMcpResponseBodyDataAuthDirectProxy) Validate() error {
	return dara.Validate(s)
}

type GetMcpResponseBodyDataAuthHttpToMcp struct {
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

func (s GetMcpResponseBodyDataAuthHttpToMcp) String() string {
	return dara.Prettify(s)
}

func (s GetMcpResponseBodyDataAuthHttpToMcp) GoString() string {
	return s.String()
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) GetCredential() *string {
	return s.Credential
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) GetId() *string {
	return s.Id
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) GetName() *string {
	return s.Name
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) GetPosition() *string {
	return s.Position
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) GetType() *string {
	return s.Type
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) SetCredential(v string) *GetMcpResponseBodyDataAuthHttpToMcp {
	s.Credential = &v
	return s
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) SetId(v string) *GetMcpResponseBodyDataAuthHttpToMcp {
	s.Id = &v
	return s
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) SetName(v string) *GetMcpResponseBodyDataAuthHttpToMcp {
	s.Name = &v
	return s
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) SetPosition(v string) *GetMcpResponseBodyDataAuthHttpToMcp {
	s.Position = &v
	return s
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) SetType(v string) *GetMcpResponseBodyDataAuthHttpToMcp {
	s.Type = &v
	return s
}

func (s *GetMcpResponseBodyDataAuthHttpToMcp) Validate() error {
	return dara.Validate(s)
}
