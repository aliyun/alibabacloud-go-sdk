// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateModelConnectionResponseBody
	GetCode() *string
	SetData(v *CreateModelConnectionResponseBodyData) *CreateModelConnectionResponseBody
	GetData() *CreateModelConnectionResponseBodyData
	SetHttpStatusCode(v int32) *CreateModelConnectionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateModelConnectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateModelConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateModelConnectionResponseBody
	GetSuccess() *bool
}

type CreateModelConnectionResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateModelConnectionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// request-1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateModelConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelConnectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateModelConnectionResponseBody) GetData() *CreateModelConnectionResponseBodyData {
	return s.Data
}

func (s *CreateModelConnectionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateModelConnectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateModelConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelConnectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateModelConnectionResponseBody) SetCode(v string) *CreateModelConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateModelConnectionResponseBody) SetData(v *CreateModelConnectionResponseBodyData) *CreateModelConnectionResponseBody {
	s.Data = v
	return s
}

func (s *CreateModelConnectionResponseBody) SetHttpStatusCode(v int32) *CreateModelConnectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateModelConnectionResponseBody) SetMessage(v string) *CreateModelConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateModelConnectionResponseBody) SetRequestId(v string) *CreateModelConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelConnectionResponseBody) SetSuccess(v bool) *CreateModelConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateModelConnectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateModelConnectionResponseBodyData struct {
	// example:
	//
	// 1
	ApiKeyCount *int32 `json:"apiKeyCount,omitempty" xml:"apiKeyCount,omitempty"`
	// example:
	//
	// mc-1
	ConnectionId *string `json:"connectionId,omitempty" xml:"connectionId,omitempty"`
	// example:
	//
	// 2026-08-09T00:00:00Z
	CreatedAt            *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CredentialConfigured *bool   `json:"credentialConfigured,omitempty" xml:"credentialConfigured,omitempty"`
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// https://dashscope.aliyuncs.com/compatible-mode/v1
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// OpenAI/v1
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// qwen
	ProviderType *string `json:"providerType,omitempty" xml:"providerType,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// GatewayOperationException
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// example:
	//
	// 2026-08-09T00:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// ws-1
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateModelConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateModelConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateModelConnectionResponseBodyData) GetApiKeyCount() *int32 {
	return s.ApiKeyCount
}

func (s *CreateModelConnectionResponseBodyData) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *CreateModelConnectionResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateModelConnectionResponseBodyData) GetCredentialConfigured() *bool {
	return s.CredentialConfigured
}

func (s *CreateModelConnectionResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateModelConnectionResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateModelConnectionResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateModelConnectionResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateModelConnectionResponseBodyData) GetProviderType() *string {
	return s.ProviderType
}

func (s *CreateModelConnectionResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateModelConnectionResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *CreateModelConnectionResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateModelConnectionResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateModelConnectionResponseBodyData) SetApiKeyCount(v int32) *CreateModelConnectionResponseBodyData {
	s.ApiKeyCount = &v
	return s
}

func (s *CreateModelConnectionResponseBodyData) SetConnectionId(v string) *CreateModelConnectionResponseBodyData {
	s.ConnectionId = &v
	return s
}

func (s *CreateModelConnectionResponseBodyData) SetCreatedAt(v string) *CreateModelConnectionResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateModelConnectionResponseBodyData) SetCredentialConfigured(v bool) *CreateModelConnectionResponseBodyData {
	s.CredentialConfigured = &v
	return s
}

func (s *CreateModelConnectionResponseBodyData) SetDescription(v string) *CreateModelConnectionResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateModelConnectionResponseBodyData) SetEndpoint(v string) *CreateModelConnectionResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *CreateModelConnectionResponseBodyData) SetName(v string) *CreateModelConnectionResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateModelConnectionResponseBodyData) SetProtocol(v string) *CreateModelConnectionResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *CreateModelConnectionResponseBodyData) SetProviderType(v string) *CreateModelConnectionResponseBodyData {
	s.ProviderType = &v
	return s
}

func (s *CreateModelConnectionResponseBodyData) SetStatus(v string) *CreateModelConnectionResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateModelConnectionResponseBodyData) SetStatusReason(v string) *CreateModelConnectionResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *CreateModelConnectionResponseBodyData) SetUpdatedAt(v string) *CreateModelConnectionResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *CreateModelConnectionResponseBodyData) SetWorkspaceId(v string) *CreateModelConnectionResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateModelConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
