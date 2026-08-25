// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListModelsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListModelsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListModelsResponseBodyItems) *ListModelsResponseBody
	GetItems() []*ListModelsResponseBodyItems
	SetMaxResults(v int32) *ListModelsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListModelsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListModelsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListModelsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListModelsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListModelsResponseBody
	GetTotalCount() *int64
}

type ListModelsResponseBody struct {
	// The business status code. A value of SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of models.
	Items []*ListModelsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The number of results per page. Valid values: 0 to 100. If this parameter is not set or set to 0, the default value 10 is used.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The message of the request processing result.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token. Pass the token returned from the previous query. An empty response indicates that no more pages are available.
	//
	// example:
	//
	// bW9kZWwtbWFuYWdlbWVudC1vZmZzZXQ6bW9kZWw6MTA
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of resources that match the query conditions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListModelsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListModelsResponseBody) GetItems() []*ListModelsResponseBodyItems {
	return s.Items
}

func (s *ListModelsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListModelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListModelsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListModelsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListModelsResponseBody) SetCode(v string) *ListModelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListModelsResponseBody) SetHttpStatusCode(v int32) *ListModelsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListModelsResponseBody) SetItems(v []*ListModelsResponseBodyItems) *ListModelsResponseBody {
	s.Items = v
	return s
}

func (s *ListModelsResponseBody) SetMaxResults(v int32) *ListModelsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListModelsResponseBody) SetMessage(v string) *ListModelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListModelsResponseBody) SetNextToken(v string) *ListModelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListModelsResponseBody) SetRequestId(v string) *ListModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelsResponseBody) SetSuccess(v bool) *ListModelsResponseBody {
	s.Success = &v
	return s
}

func (s *ListModelsResponseBody) SetTotalCount(v int64) *ListModelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelsResponseBody) Validate() error {
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

type ListModelsResponseBodyItems struct {
	// The model capability configuration.
	Capabilities *ListModelsResponseBodyItemsCapabilities `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Struct"`
	// The model connection ID.
	//
	// example:
	//
	// mc-1
	ConnectionId *string `json:"connectionId,omitempty" xml:"connectionId,omitempty"`
	// The model context window size, in tokens. The value must be a positive integer.
	//
	// example:
	//
	// 128000
	ContextSize *int64 `json:"contextSize,omitempty" xml:"contextSize,omitempty"`
	// The time when the resource was created, in RFC 3339 UTC format.
	//
	// example:
	//
	// 2026-08-09T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The model description. Maximum length: 255 characters.
	//
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The maximum number of output tokens supported by the model in a single generation.
	//
	// example:
	//
	// 131072
	MaxTokens *int64 `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	// The model ID.
	//
	// example:
	//
	// model-1
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// The upstream model name.
	//
	// example:
	//
	// qwen-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
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

func (s ListModelsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyItems) GetCapabilities() *ListModelsResponseBodyItemsCapabilities {
	return s.Capabilities
}

func (s *ListModelsResponseBodyItems) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *ListModelsResponseBodyItems) GetContextSize() *int64 {
	return s.ContextSize
}

func (s *ListModelsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListModelsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListModelsResponseBodyItems) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *ListModelsResponseBodyItems) GetModelId() *string {
	return s.ModelId
}

func (s *ListModelsResponseBodyItems) GetModelName() *string {
	return s.ModelName
}

func (s *ListModelsResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListModelsResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListModelsResponseBodyItems) SetCapabilities(v *ListModelsResponseBodyItemsCapabilities) *ListModelsResponseBodyItems {
	s.Capabilities = v
	return s
}

func (s *ListModelsResponseBodyItems) SetConnectionId(v string) *ListModelsResponseBodyItems {
	s.ConnectionId = &v
	return s
}

func (s *ListModelsResponseBodyItems) SetContextSize(v int64) *ListModelsResponseBodyItems {
	s.ContextSize = &v
	return s
}

func (s *ListModelsResponseBodyItems) SetCreatedAt(v string) *ListModelsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListModelsResponseBodyItems) SetDescription(v string) *ListModelsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListModelsResponseBodyItems) SetMaxTokens(v int64) *ListModelsResponseBodyItems {
	s.MaxTokens = &v
	return s
}

func (s *ListModelsResponseBodyItems) SetModelId(v string) *ListModelsResponseBodyItems {
	s.ModelId = &v
	return s
}

func (s *ListModelsResponseBodyItems) SetModelName(v string) *ListModelsResponseBodyItems {
	s.ModelName = &v
	return s
}

func (s *ListModelsResponseBodyItems) SetUpdatedAt(v string) *ListModelsResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListModelsResponseBodyItems) SetWorkspaceId(v string) *ListModelsResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListModelsResponseBodyItems) Validate() error {
	if s.Capabilities != nil {
		if err := s.Capabilities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListModelsResponseBodyItemsCapabilities struct {
	// Indicates whether the model supports audio input or output.
	Audio *bool `json:"audio,omitempty" xml:"audio,omitempty"`
	// Indicates whether the model supports document input.
	Document *bool `json:"document,omitempty" xml:"document,omitempty"`
	// Indicates whether the model supports invoking multiple tools in a single response.
	MultiToolCall *bool `json:"multiToolCall,omitempty" xml:"multiToolCall,omitempty"`
	// Indicates whether the model supports reasoning capabilities.
	Reasoning *bool `json:"reasoning,omitempty" xml:"reasoning,omitempty"`
	// Indicates whether the model supports streaming tool calling.
	StreamToolCall *bool `json:"streamToolCall,omitempty" xml:"streamToolCall,omitempty"`
	// Indicates whether the model supports tool calling.
	ToolCall *bool `json:"toolCall,omitempty" xml:"toolCall,omitempty"`
	// Indicates whether the model supports video input.
	Video *bool `json:"video,omitempty" xml:"video,omitempty"`
	// Indicates whether the model supports image input.
	Vision *bool `json:"vision,omitempty" xml:"vision,omitempty"`
}

func (s ListModelsResponseBodyItemsCapabilities) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyItemsCapabilities) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyItemsCapabilities) GetAudio() *bool {
	return s.Audio
}

func (s *ListModelsResponseBodyItemsCapabilities) GetDocument() *bool {
	return s.Document
}

func (s *ListModelsResponseBodyItemsCapabilities) GetMultiToolCall() *bool {
	return s.MultiToolCall
}

func (s *ListModelsResponseBodyItemsCapabilities) GetReasoning() *bool {
	return s.Reasoning
}

func (s *ListModelsResponseBodyItemsCapabilities) GetStreamToolCall() *bool {
	return s.StreamToolCall
}

func (s *ListModelsResponseBodyItemsCapabilities) GetToolCall() *bool {
	return s.ToolCall
}

func (s *ListModelsResponseBodyItemsCapabilities) GetVideo() *bool {
	return s.Video
}

func (s *ListModelsResponseBodyItemsCapabilities) GetVision() *bool {
	return s.Vision
}

func (s *ListModelsResponseBodyItemsCapabilities) SetAudio(v bool) *ListModelsResponseBodyItemsCapabilities {
	s.Audio = &v
	return s
}

func (s *ListModelsResponseBodyItemsCapabilities) SetDocument(v bool) *ListModelsResponseBodyItemsCapabilities {
	s.Document = &v
	return s
}

func (s *ListModelsResponseBodyItemsCapabilities) SetMultiToolCall(v bool) *ListModelsResponseBodyItemsCapabilities {
	s.MultiToolCall = &v
	return s
}

func (s *ListModelsResponseBodyItemsCapabilities) SetReasoning(v bool) *ListModelsResponseBodyItemsCapabilities {
	s.Reasoning = &v
	return s
}

func (s *ListModelsResponseBodyItemsCapabilities) SetStreamToolCall(v bool) *ListModelsResponseBodyItemsCapabilities {
	s.StreamToolCall = &v
	return s
}

func (s *ListModelsResponseBodyItemsCapabilities) SetToolCall(v bool) *ListModelsResponseBodyItemsCapabilities {
	s.ToolCall = &v
	return s
}

func (s *ListModelsResponseBodyItemsCapabilities) SetVideo(v bool) *ListModelsResponseBodyItemsCapabilities {
	s.Video = &v
	return s
}

func (s *ListModelsResponseBodyItemsCapabilities) SetVision(v bool) *ListModelsResponseBodyItemsCapabilities {
	s.Vision = &v
	return s
}

func (s *ListModelsResponseBodyItemsCapabilities) Validate() error {
	return dara.Validate(s)
}
