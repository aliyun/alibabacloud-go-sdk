// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetModelConnectionResponseBody
	GetCode() *string
	SetData(v *GetModelConnectionResponseBodyData) *GetModelConnectionResponseBody
	GetData() *GetModelConnectionResponseBodyData
	SetHttpStatusCode(v int32) *GetModelConnectionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetModelConnectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetModelConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetModelConnectionResponseBody
	GetSuccess() *bool
}

type GetModelConnectionResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetModelConnectionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetModelConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelConnectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetModelConnectionResponseBody) GetData() *GetModelConnectionResponseBodyData {
	return s.Data
}

func (s *GetModelConnectionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetModelConnectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetModelConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelConnectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetModelConnectionResponseBody) SetCode(v string) *GetModelConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *GetModelConnectionResponseBody) SetData(v *GetModelConnectionResponseBodyData) *GetModelConnectionResponseBody {
	s.Data = v
	return s
}

func (s *GetModelConnectionResponseBody) SetHttpStatusCode(v int32) *GetModelConnectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetModelConnectionResponseBody) SetMessage(v string) *GetModelConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *GetModelConnectionResponseBody) SetRequestId(v string) *GetModelConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelConnectionResponseBody) SetSuccess(v bool) *GetModelConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *GetModelConnectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModelConnectionResponseBodyData struct {
	// example:
	//
	// 1
	ApiKeyCount *int32    `json:"apiKeyCount,omitempty" xml:"apiKeyCount,omitempty"`
	ApiKeys     []*string `json:"apiKeys,omitempty" xml:"apiKeys,omitempty" type:"Repeated"`
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
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
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

func (s GetModelConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetModelConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetModelConnectionResponseBodyData) GetApiKeyCount() *int32 {
	return s.ApiKeyCount
}

func (s *GetModelConnectionResponseBodyData) GetApiKeys() []*string {
	return s.ApiKeys
}

func (s *GetModelConnectionResponseBodyData) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *GetModelConnectionResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetModelConnectionResponseBodyData) GetCredentialConfigured() *bool {
	return s.CredentialConfigured
}

func (s *GetModelConnectionResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetModelConnectionResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetModelConnectionResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetModelConnectionResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *GetModelConnectionResponseBodyData) GetProviderType() *string {
	return s.ProviderType
}

func (s *GetModelConnectionResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetModelConnectionResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetModelConnectionResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetModelConnectionResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetModelConnectionResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetModelConnectionResponseBodyData) SetApiKeyCount(v int32) *GetModelConnectionResponseBodyData {
	s.ApiKeyCount = &v
	return s
}

func (s *GetModelConnectionResponseBodyData) SetApiKeys(v []*string) *GetModelConnectionResponseBodyData {
	s.ApiKeys = v
	return s
}

func (s *GetModelConnectionResponseBodyData) SetConnectionId(v string) *GetModelConnectionResponseBodyData {
	s.ConnectionId = &v
	return s
}

func (s *GetModelConnectionResponseBodyData) SetCreatedAt(v string) *GetModelConnectionResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetModelConnectionResponseBodyData) SetCredentialConfigured(v bool) *GetModelConnectionResponseBodyData {
	s.CredentialConfigured = &v
	return s
}

func (s *GetModelConnectionResponseBodyData) SetDescription(v string) *GetModelConnectionResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetModelConnectionResponseBodyData) SetEndpoint(v string) *GetModelConnectionResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetModelConnectionResponseBodyData) SetName(v string) *GetModelConnectionResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetModelConnectionResponseBodyData) SetProtocol(v string) *GetModelConnectionResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetModelConnectionResponseBodyData) SetProviderType(v string) *GetModelConnectionResponseBodyData {
	s.ProviderType = &v
	return s
}

func (s *GetModelConnectionResponseBodyData) SetRegionId(v string) *GetModelConnectionResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetModelConnectionResponseBodyData) SetStatus(v string) *GetModelConnectionResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetModelConnectionResponseBodyData) SetStatusReason(v string) *GetModelConnectionResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *GetModelConnectionResponseBodyData) SetUpdatedAt(v string) *GetModelConnectionResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetModelConnectionResponseBodyData) SetWorkspaceId(v string) *GetModelConnectionResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetModelConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
