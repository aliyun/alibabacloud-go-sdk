// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateModelResponseBody
	GetCode() *string
	SetData(v *UpdateModelResponseBodyData) *UpdateModelResponseBody
	GetData() *UpdateModelResponseBodyData
	SetHttpStatusCode(v int32) *UpdateModelResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateModelResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateModelResponseBody
	GetSuccess() *bool
}

type UpdateModelResponseBody struct {
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The updated model information.
	Data *UpdateModelResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. The value 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The message of the request processing result.
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

func (s UpdateModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateModelResponseBody) GetData() *UpdateModelResponseBodyData {
	return s.Data
}

func (s *UpdateModelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateModelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateModelResponseBody) SetCode(v string) *UpdateModelResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateModelResponseBody) SetData(v *UpdateModelResponseBodyData) *UpdateModelResponseBody {
	s.Data = v
	return s
}

func (s *UpdateModelResponseBody) SetHttpStatusCode(v int32) *UpdateModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateModelResponseBody) SetMessage(v string) *UpdateModelResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateModelResponseBody) SetRequestId(v string) *UpdateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModelResponseBody) SetSuccess(v bool) *UpdateModelResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateModelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateModelResponseBodyData struct {
	// The model capability configurations.
	Capabilities *UpdateModelResponseBodyDataCapabilities `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Struct"`
	// The model connection ID.
	//
	// example:
	//
	// mc-1
	ConnectionId *string `json:"connectionId,omitempty" xml:"connectionId,omitempty"`
	// The model context window size, in tokens. Must be a positive integer.
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
	// The maximum number of output tokens supported per model generation.
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

func (s UpdateModelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateModelResponseBodyData) GetCapabilities() *UpdateModelResponseBodyDataCapabilities {
	return s.Capabilities
}

func (s *UpdateModelResponseBodyData) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *UpdateModelResponseBodyData) GetContextSize() *int64 {
	return s.ContextSize
}

func (s *UpdateModelResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateModelResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelResponseBodyData) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *UpdateModelResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *UpdateModelResponseBodyData) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateModelResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateModelResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateModelResponseBodyData) SetCapabilities(v *UpdateModelResponseBodyDataCapabilities) *UpdateModelResponseBodyData {
	s.Capabilities = v
	return s
}

func (s *UpdateModelResponseBodyData) SetConnectionId(v string) *UpdateModelResponseBodyData {
	s.ConnectionId = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetContextSize(v int64) *UpdateModelResponseBodyData {
	s.ContextSize = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetCreatedAt(v string) *UpdateModelResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetDescription(v string) *UpdateModelResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetMaxTokens(v int64) *UpdateModelResponseBodyData {
	s.MaxTokens = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetModelId(v string) *UpdateModelResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetModelName(v string) *UpdateModelResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetUpdatedAt(v string) *UpdateModelResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateModelResponseBodyData) SetWorkspaceId(v string) *UpdateModelResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateModelResponseBodyData) Validate() error {
	if s.Capabilities != nil {
		if err := s.Capabilities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateModelResponseBodyDataCapabilities struct {
	// Indicates whether the model supports audio input or output.
	Audio *bool `json:"audio,omitempty" xml:"audio,omitempty"`
	// Indicates whether the model supports document input.
	Document *bool `json:"document,omitempty" xml:"document,omitempty"`
	// Indicates whether the model is able to invoke multiple tool calling requests in a single response.
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

func (s UpdateModelResponseBodyDataCapabilities) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelResponseBodyDataCapabilities) GoString() string {
	return s.String()
}

func (s *UpdateModelResponseBodyDataCapabilities) GetAudio() *bool {
	return s.Audio
}

func (s *UpdateModelResponseBodyDataCapabilities) GetDocument() *bool {
	return s.Document
}

func (s *UpdateModelResponseBodyDataCapabilities) GetMultiToolCall() *bool {
	return s.MultiToolCall
}

func (s *UpdateModelResponseBodyDataCapabilities) GetReasoning() *bool {
	return s.Reasoning
}

func (s *UpdateModelResponseBodyDataCapabilities) GetStreamToolCall() *bool {
	return s.StreamToolCall
}

func (s *UpdateModelResponseBodyDataCapabilities) GetToolCall() *bool {
	return s.ToolCall
}

func (s *UpdateModelResponseBodyDataCapabilities) GetVideo() *bool {
	return s.Video
}

func (s *UpdateModelResponseBodyDataCapabilities) GetVision() *bool {
	return s.Vision
}

func (s *UpdateModelResponseBodyDataCapabilities) SetAudio(v bool) *UpdateModelResponseBodyDataCapabilities {
	s.Audio = &v
	return s
}

func (s *UpdateModelResponseBodyDataCapabilities) SetDocument(v bool) *UpdateModelResponseBodyDataCapabilities {
	s.Document = &v
	return s
}

func (s *UpdateModelResponseBodyDataCapabilities) SetMultiToolCall(v bool) *UpdateModelResponseBodyDataCapabilities {
	s.MultiToolCall = &v
	return s
}

func (s *UpdateModelResponseBodyDataCapabilities) SetReasoning(v bool) *UpdateModelResponseBodyDataCapabilities {
	s.Reasoning = &v
	return s
}

func (s *UpdateModelResponseBodyDataCapabilities) SetStreamToolCall(v bool) *UpdateModelResponseBodyDataCapabilities {
	s.StreamToolCall = &v
	return s
}

func (s *UpdateModelResponseBodyDataCapabilities) SetToolCall(v bool) *UpdateModelResponseBodyDataCapabilities {
	s.ToolCall = &v
	return s
}

func (s *UpdateModelResponseBodyDataCapabilities) SetVideo(v bool) *UpdateModelResponseBodyDataCapabilities {
	s.Video = &v
	return s
}

func (s *UpdateModelResponseBodyDataCapabilities) SetVision(v bool) *UpdateModelResponseBodyDataCapabilities {
	s.Vision = &v
	return s
}

func (s *UpdateModelResponseBodyDataCapabilities) Validate() error {
	return dara.Validate(s)
}
