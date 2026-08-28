// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateModelConnectionResponseBody
	GetCode() *string
	SetData(v *UpdateModelConnectionResponseBodyData) *UpdateModelConnectionResponseBody
	GetData() *UpdateModelConnectionResponseBodyData
	SetHttpStatusCode(v int32) *UpdateModelConnectionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateModelConnectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateModelConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateModelConnectionResponseBody
	GetSuccess() *bool
}

type UpdateModelConnectionResponseBody struct {
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The updated model connection information.
	Data *UpdateModelConnectionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. The value 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request processing result message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateModelConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelConnectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateModelConnectionResponseBody) GetData() *UpdateModelConnectionResponseBodyData {
	return s.Data
}

func (s *UpdateModelConnectionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateModelConnectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateModelConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateModelConnectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateModelConnectionResponseBody) SetCode(v string) *UpdateModelConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateModelConnectionResponseBody) SetData(v *UpdateModelConnectionResponseBodyData) *UpdateModelConnectionResponseBody {
	s.Data = v
	return s
}

func (s *UpdateModelConnectionResponseBody) SetHttpStatusCode(v int32) *UpdateModelConnectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateModelConnectionResponseBody) SetMessage(v string) *UpdateModelConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateModelConnectionResponseBody) SetRequestId(v string) *UpdateModelConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModelConnectionResponseBody) SetSuccess(v bool) *UpdateModelConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateModelConnectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateModelConnectionResponseBodyData struct {
	// The number of API keys configured in the model connection.
	//
	// example:
	//
	// 1
	ApiKeyCount *int32 `json:"apiKeyCount,omitempty" xml:"apiKeyCount,omitempty"`
	// The model connection ID.
	//
	// example:
	//
	// mc-1
	ConnectionId *string `json:"connectionId,omitempty" xml:"connectionId,omitempty"`
	// The time when the resource was created, in RFC 3339 UTC format.
	//
	// example:
	//
	// 2026-08-09T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Indicates whether access credentials have been configured for the model connection.
	CredentialConfigured *bool `json:"credentialConfigured,omitempty" xml:"credentialConfigured,omitempty"`
	// The description of the model connection. The description can be up to 255 characters in length.
	//
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The absolute HTTP or HTTPS address of the upstream model service. The address can be up to 1024 characters in length.
	//
	// example:
	//
	// https://dashscope.aliyuncs.com/compatible-mode/v1
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The model connection name. The name must be 1 to 128 non-whitespace characters in length.
	//
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The model invocation protocol. Currently, only OpenAI/v1 is supported. If not specified in Settings when the model connection is created, this default value is used.
	//
	// example:
	//
	// OpenAI/v1
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The model provider type.
	//
	// example:
	//
	// qwen
	ProviderType *string `json:"providerType,omitempty" xml:"providerType,omitempty"`
	// The resource status.
	//
	// example:
	//
	// Active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The failure summary returned when the model connection fails to be published or fails to be deleted but remains in the Deleting state. This value is empty for other states.
	//
	// example:
	//
	// GatewayOperationException
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// The time when the resource was last updated, in RFC 3339 UTC format.
	//
	// example:
	//
	// 2026-08-09T00:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-1
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateModelConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateModelConnectionResponseBodyData) GetApiKeyCount() *int32 {
	return s.ApiKeyCount
}

func (s *UpdateModelConnectionResponseBodyData) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *UpdateModelConnectionResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateModelConnectionResponseBodyData) GetCredentialConfigured() *bool {
	return s.CredentialConfigured
}

func (s *UpdateModelConnectionResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelConnectionResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateModelConnectionResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateModelConnectionResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateModelConnectionResponseBodyData) GetProviderType() *string {
	return s.ProviderType
}

func (s *UpdateModelConnectionResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateModelConnectionResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *UpdateModelConnectionResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateModelConnectionResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateModelConnectionResponseBodyData) SetApiKeyCount(v int32) *UpdateModelConnectionResponseBodyData {
	s.ApiKeyCount = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetConnectionId(v string) *UpdateModelConnectionResponseBodyData {
	s.ConnectionId = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetCreatedAt(v string) *UpdateModelConnectionResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetCredentialConfigured(v bool) *UpdateModelConnectionResponseBodyData {
	s.CredentialConfigured = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetDescription(v string) *UpdateModelConnectionResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetEndpoint(v string) *UpdateModelConnectionResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetName(v string) *UpdateModelConnectionResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetProtocol(v string) *UpdateModelConnectionResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetProviderType(v string) *UpdateModelConnectionResponseBodyData {
	s.ProviderType = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetStatus(v string) *UpdateModelConnectionResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetStatusReason(v string) *UpdateModelConnectionResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetUpdatedAt(v string) *UpdateModelConnectionResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetWorkspaceId(v string) *UpdateModelConnectionResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
