// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMcpsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListMcpsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListMcpsResponseBodyItems) *ListMcpsResponseBody
	GetItems() []*ListMcpsResponseBodyItems
	SetMaxResults(v int32) *ListMcpsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListMcpsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListMcpsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMcpsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMcpsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListMcpsResponseBody
	GetTotalCount() *int64
}

type ListMcpsResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list data.
	Items []*ListMcpsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of entries to return per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// Request processed successfully
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// next-page-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMcpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMcpsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMcpsResponseBody) GetItems() []*ListMcpsResponseBodyItems {
	return s.Items
}

func (s *ListMcpsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMcpsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcpsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcpsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMcpsResponseBody) SetCode(v string) *ListMcpsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMcpsResponseBody) SetHttpStatusCode(v int32) *ListMcpsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMcpsResponseBody) SetItems(v []*ListMcpsResponseBodyItems) *ListMcpsResponseBody {
	s.Items = v
	return s
}

func (s *ListMcpsResponseBody) SetMaxResults(v int32) *ListMcpsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMcpsResponseBody) SetMessage(v string) *ListMcpsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMcpsResponseBody) SetNextToken(v string) *ListMcpsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMcpsResponseBody) SetRequestId(v string) *ListMcpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcpsResponseBody) SetSuccess(v bool) *ListMcpsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMcpsResponseBody) SetTotalCount(v int64) *ListMcpsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMcpsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMcpsResponseBodyItems struct {
	// The list of MCP service addresses.
	Addresses []*string `json:"addresses,omitempty" xml:"addresses,omitempty" type:"Repeated"`
	// The backend authentication configuration. enabled indicates whether authentication is enabled. directProxy specifies the custom authentication header for direct proxy. httpToMcp specifies the list of OpenAPI credentials for HTTP_TO_MCP.
	Auth *ListMcpsResponseBodyItemsAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// A sample description that explains the purpose of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
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
}

func (s ListMcpsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItems) GetAddresses() []*string {
	return s.Addresses
}

func (s *ListMcpsResponseBodyItems) GetAuth() *ListMcpsResponseBodyItemsAuth {
	return s.Auth
}

func (s *ListMcpsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListMcpsResponseBodyItems) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *ListMcpsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListMcpsResponseBodyItems) GetProtocol() *string {
	return s.Protocol
}

func (s *ListMcpsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListMcpsResponseBodyItems) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListMcpsResponseBodyItems) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *ListMcpsResponseBodyItems) GetType() *string {
	return s.Type
}

func (s *ListMcpsResponseBodyItems) SetAddresses(v []*string) *ListMcpsResponseBodyItems {
	s.Addresses = v
	return s
}

func (s *ListMcpsResponseBodyItems) SetAuth(v *ListMcpsResponseBodyItemsAuth) *ListMcpsResponseBodyItems {
	s.Auth = v
	return s
}

func (s *ListMcpsResponseBodyItems) SetDescription(v string) *ListMcpsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetMcpServerId(v string) *ListMcpsResponseBodyItems {
	s.McpServerId = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetName(v string) *ListMcpsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetProtocol(v string) *ListMcpsResponseBodyItems {
	s.Protocol = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetStatus(v string) *ListMcpsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetStatusReason(v string) *ListMcpsResponseBodyItems {
	s.StatusReason = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetSwaggerConfig(v string) *ListMcpsResponseBodyItems {
	s.SwaggerConfig = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetType(v string) *ListMcpsResponseBodyItems {
	s.Type = &v
	return s
}

func (s *ListMcpsResponseBodyItems) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMcpsResponseBodyItemsAuth struct {
	// The direct proxy authentication configuration.
	DirectProxy *ListMcpsResponseBodyItemsAuthDirectProxy `json:"directProxy,omitempty" xml:"directProxy,omitempty" type:"Struct"`
	// Indicates whether authentication is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of HTTP_TO_MCP authentication configurations.
	HttpToMcp []*ListMcpsResponseBodyItemsAuthHttpToMcp `json:"httpToMcp,omitempty" xml:"httpToMcp,omitempty" type:"Repeated"`
}

func (s ListMcpsResponseBodyItemsAuth) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsAuth) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsAuth) GetDirectProxy() *ListMcpsResponseBodyItemsAuthDirectProxy {
	return s.DirectProxy
}

func (s *ListMcpsResponseBodyItemsAuth) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListMcpsResponseBodyItemsAuth) GetHttpToMcp() []*ListMcpsResponseBodyItemsAuthHttpToMcp {
	return s.HttpToMcp
}

func (s *ListMcpsResponseBodyItemsAuth) SetDirectProxy(v *ListMcpsResponseBodyItemsAuthDirectProxy) *ListMcpsResponseBodyItemsAuth {
	s.DirectProxy = v
	return s
}

func (s *ListMcpsResponseBodyItemsAuth) SetEnabled(v bool) *ListMcpsResponseBodyItemsAuth {
	s.Enabled = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuth) SetHttpToMcp(v []*ListMcpsResponseBodyItemsAuthHttpToMcp) *ListMcpsResponseBodyItemsAuth {
	s.HttpToMcp = v
	return s
}

func (s *ListMcpsResponseBodyItemsAuth) Validate() error {
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

type ListMcpsResponseBodyItemsAuthDirectProxy struct {
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

func (s ListMcpsResponseBodyItemsAuthDirectProxy) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsAuthDirectProxy) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsAuthDirectProxy) GetName() *string {
	return s.Name
}

func (s *ListMcpsResponseBodyItemsAuthDirectProxy) GetValue() *string {
	return s.Value
}

func (s *ListMcpsResponseBodyItemsAuthDirectProxy) SetName(v string) *ListMcpsResponseBodyItemsAuthDirectProxy {
	s.Name = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuthDirectProxy) SetValue(v string) *ListMcpsResponseBodyItemsAuthDirectProxy {
	s.Value = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuthDirectProxy) Validate() error {
	return dara.Validate(s)
}

type ListMcpsResponseBodyItemsAuthHttpToMcp struct {
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

func (s ListMcpsResponseBodyItemsAuthHttpToMcp) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItemsAuthHttpToMcp) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) GetCredential() *string {
	return s.Credential
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) GetId() *string {
	return s.Id
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) GetName() *string {
	return s.Name
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) GetPosition() *string {
	return s.Position
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) GetType() *string {
	return s.Type
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) SetCredential(v string) *ListMcpsResponseBodyItemsAuthHttpToMcp {
	s.Credential = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) SetId(v string) *ListMcpsResponseBodyItemsAuthHttpToMcp {
	s.Id = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) SetName(v string) *ListMcpsResponseBodyItemsAuthHttpToMcp {
	s.Name = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) SetPosition(v string) *ListMcpsResponseBodyItemsAuthHttpToMcp {
	s.Position = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) SetType(v string) *ListMcpsResponseBodyItemsAuthHttpToMcp {
	s.Type = &v
	return s
}

func (s *ListMcpsResponseBodyItemsAuthHttpToMcp) Validate() error {
	return dara.Validate(s)
}
