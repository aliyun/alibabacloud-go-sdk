// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListModelConnectionsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListModelConnectionsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListModelConnectionsResponseBodyItems) *ListModelConnectionsResponseBody
	GetItems() []*ListModelConnectionsResponseBodyItems
	SetMaxResults(v int32) *ListModelConnectionsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListModelConnectionsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListModelConnectionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListModelConnectionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListModelConnectionsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListModelConnectionsResponseBody
	GetTotalCount() *int64
}

type ListModelConnectionsResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                   `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Items          []*ListModelConnectionsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// bW9kZWwtbWFuYWdlbWVudC1vZmZzZXQ6bW9kZWwtY29ubmVjdGlvbjoxMA
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// request-1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListModelConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelConnectionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListModelConnectionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListModelConnectionsResponseBody) GetItems() []*ListModelConnectionsResponseBodyItems {
	return s.Items
}

func (s *ListModelConnectionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListModelConnectionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListModelConnectionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelConnectionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListModelConnectionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListModelConnectionsResponseBody) SetCode(v string) *ListModelConnectionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListModelConnectionsResponseBody) SetHttpStatusCode(v int32) *ListModelConnectionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListModelConnectionsResponseBody) SetItems(v []*ListModelConnectionsResponseBodyItems) *ListModelConnectionsResponseBody {
	s.Items = v
	return s
}

func (s *ListModelConnectionsResponseBody) SetMaxResults(v int32) *ListModelConnectionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListModelConnectionsResponseBody) SetMessage(v string) *ListModelConnectionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListModelConnectionsResponseBody) SetNextToken(v string) *ListModelConnectionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListModelConnectionsResponseBody) SetRequestId(v string) *ListModelConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelConnectionsResponseBody) SetSuccess(v bool) *ListModelConnectionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListModelConnectionsResponseBody) SetTotalCount(v int64) *ListModelConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelConnectionsResponseBody) Validate() error {
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

type ListModelConnectionsResponseBodyItems struct {
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
	Endpoint *string                                        `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Models   []*ListModelConnectionsResponseBodyItemsModels `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
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

func (s ListModelConnectionsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListModelConnectionsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListModelConnectionsResponseBodyItems) GetApiKeyCount() *int32 {
	return s.ApiKeyCount
}

func (s *ListModelConnectionsResponseBodyItems) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *ListModelConnectionsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListModelConnectionsResponseBodyItems) GetCredentialConfigured() *bool {
	return s.CredentialConfigured
}

func (s *ListModelConnectionsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListModelConnectionsResponseBodyItems) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListModelConnectionsResponseBodyItems) GetModels() []*ListModelConnectionsResponseBodyItemsModels {
	return s.Models
}

func (s *ListModelConnectionsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListModelConnectionsResponseBodyItems) GetProtocol() *string {
	return s.Protocol
}

func (s *ListModelConnectionsResponseBodyItems) GetProviderType() *string {
	return s.ProviderType
}

func (s *ListModelConnectionsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListModelConnectionsResponseBodyItems) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListModelConnectionsResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListModelConnectionsResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListModelConnectionsResponseBodyItems) SetApiKeyCount(v int32) *ListModelConnectionsResponseBodyItems {
	s.ApiKeyCount = &v
	return s
}

func (s *ListModelConnectionsResponseBodyItems) SetConnectionId(v string) *ListModelConnectionsResponseBodyItems {
	s.ConnectionId = &v
	return s
}

func (s *ListModelConnectionsResponseBodyItems) SetCreatedAt(v string) *ListModelConnectionsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListModelConnectionsResponseBodyItems) SetCredentialConfigured(v bool) *ListModelConnectionsResponseBodyItems {
	s.CredentialConfigured = &v
	return s
}

func (s *ListModelConnectionsResponseBodyItems) SetDescription(v string) *ListModelConnectionsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListModelConnectionsResponseBodyItems) SetEndpoint(v string) *ListModelConnectionsResponseBodyItems {
	s.Endpoint = &v
	return s
}

func (s *ListModelConnectionsResponseBodyItems) SetModels(v []*ListModelConnectionsResponseBodyItemsModels) *ListModelConnectionsResponseBodyItems {
	s.Models = v
	return s
}

func (s *ListModelConnectionsResponseBodyItems) SetName(v string) *ListModelConnectionsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListModelConnectionsResponseBodyItems) SetProtocol(v string) *ListModelConnectionsResponseBodyItems {
	s.Protocol = &v
	return s
}

func (s *ListModelConnectionsResponseBodyItems) SetProviderType(v string) *ListModelConnectionsResponseBodyItems {
	s.ProviderType = &v
	return s
}

func (s *ListModelConnectionsResponseBodyItems) SetStatus(v string) *ListModelConnectionsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListModelConnectionsResponseBodyItems) SetStatusReason(v string) *ListModelConnectionsResponseBodyItems {
	s.StatusReason = &v
	return s
}

func (s *ListModelConnectionsResponseBodyItems) SetUpdatedAt(v string) *ListModelConnectionsResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListModelConnectionsResponseBodyItems) SetWorkspaceId(v string) *ListModelConnectionsResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListModelConnectionsResponseBodyItems) Validate() error {
	if s.Models != nil {
		for _, item := range s.Models {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelConnectionsResponseBodyItemsModels struct {
	// example:
	//
	// model-1
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// qwen-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
}

func (s ListModelConnectionsResponseBodyItemsModels) String() string {
	return dara.Prettify(s)
}

func (s ListModelConnectionsResponseBodyItemsModels) GoString() string {
	return s.String()
}

func (s *ListModelConnectionsResponseBodyItemsModels) GetModelId() *string {
	return s.ModelId
}

func (s *ListModelConnectionsResponseBodyItemsModels) GetModelName() *string {
	return s.ModelName
}

func (s *ListModelConnectionsResponseBodyItemsModels) SetModelId(v string) *ListModelConnectionsResponseBodyItemsModels {
	s.ModelId = &v
	return s
}

func (s *ListModelConnectionsResponseBodyItemsModels) SetModelName(v string) *ListModelConnectionsResponseBodyItemsModels {
	s.ModelName = &v
	return s
}

func (s *ListModelConnectionsResponseBodyItemsModels) Validate() error {
	return dara.Validate(s)
}
