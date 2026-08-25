// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateModelResponseBody
	GetCode() *string
	SetData(v *CreateModelResponseBodyData) *CreateModelResponseBody
	GetData() *CreateModelResponseBodyData
	SetHttpStatusCode(v int32) *CreateModelResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateModelResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateModelResponseBody
	GetSuccess() *bool
}

type CreateModelResponseBody struct {
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The model information after creation.
	Data *CreateModelResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s CreateModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateModelResponseBody) GetData() *CreateModelResponseBodyData {
	return s.Data
}

func (s *CreateModelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateModelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateModelResponseBody) SetCode(v string) *CreateModelResponseBody {
	s.Code = &v
	return s
}

func (s *CreateModelResponseBody) SetData(v *CreateModelResponseBodyData) *CreateModelResponseBody {
	s.Data = v
	return s
}

func (s *CreateModelResponseBody) SetHttpStatusCode(v int32) *CreateModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateModelResponseBody) SetMessage(v string) *CreateModelResponseBody {
	s.Message = &v
	return s
}

func (s *CreateModelResponseBody) SetRequestId(v string) *CreateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelResponseBody) SetSuccess(v bool) *CreateModelResponseBody {
	s.Success = &v
	return s
}

func (s *CreateModelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateModelResponseBodyData struct {
	// The model capability configuration.
	Capabilities *CreateModelResponseBodyDataCapabilities `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Struct"`
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
	// The resource creation time in RFC 3339 UTC format.
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
	// The resource last update time in RFC 3339 UTC format.
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

func (s CreateModelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateModelResponseBodyData) GetCapabilities() *CreateModelResponseBodyDataCapabilities {
	return s.Capabilities
}

func (s *CreateModelResponseBodyData) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *CreateModelResponseBodyData) GetContextSize() *int64 {
	return s.ContextSize
}

func (s *CreateModelResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateModelResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateModelResponseBodyData) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *CreateModelResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *CreateModelResponseBodyData) GetModelName() *string {
	return s.ModelName
}

func (s *CreateModelResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateModelResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateModelResponseBodyData) SetCapabilities(v *CreateModelResponseBodyDataCapabilities) *CreateModelResponseBodyData {
	s.Capabilities = v
	return s
}

func (s *CreateModelResponseBodyData) SetConnectionId(v string) *CreateModelResponseBodyData {
	s.ConnectionId = &v
	return s
}

func (s *CreateModelResponseBodyData) SetContextSize(v int64) *CreateModelResponseBodyData {
	s.ContextSize = &v
	return s
}

func (s *CreateModelResponseBodyData) SetCreatedAt(v string) *CreateModelResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateModelResponseBodyData) SetDescription(v string) *CreateModelResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateModelResponseBodyData) SetMaxTokens(v int64) *CreateModelResponseBodyData {
	s.MaxTokens = &v
	return s
}

func (s *CreateModelResponseBodyData) SetModelId(v string) *CreateModelResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *CreateModelResponseBodyData) SetModelName(v string) *CreateModelResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *CreateModelResponseBodyData) SetUpdatedAt(v string) *CreateModelResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *CreateModelResponseBodyData) SetWorkspaceId(v string) *CreateModelResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateModelResponseBodyData) Validate() error {
	if s.Capabilities != nil {
		if err := s.Capabilities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateModelResponseBodyDataCapabilities struct {
	// Specifies whether the model supports audio input or output.
	Audio *bool `json:"audio,omitempty" xml:"audio,omitempty"`
	// Specifies whether the model supports document input.
	Document *bool `json:"document,omitempty" xml:"document,omitempty"`
	// Specifies whether the model is able to invoke multiple tools in a single response.
	MultiToolCall *bool `json:"multiToolCall,omitempty" xml:"multiToolCall,omitempty"`
	// Specifies whether the model supports reasoning capabilities.
	Reasoning *bool `json:"reasoning,omitempty" xml:"reasoning,omitempty"`
	// Specifies whether the model supports streaming tool calling.
	StreamToolCall *bool `json:"streamToolCall,omitempty" xml:"streamToolCall,omitempty"`
	// Specifies whether the model supports tool calling.
	ToolCall *bool `json:"toolCall,omitempty" xml:"toolCall,omitempty"`
	// Specifies whether the model supports video input.
	Video *bool `json:"video,omitempty" xml:"video,omitempty"`
	// Specifies whether the model supports image input.
	Vision *bool `json:"vision,omitempty" xml:"vision,omitempty"`
}

func (s CreateModelResponseBodyDataCapabilities) String() string {
	return dara.Prettify(s)
}

func (s CreateModelResponseBodyDataCapabilities) GoString() string {
	return s.String()
}

func (s *CreateModelResponseBodyDataCapabilities) GetAudio() *bool {
	return s.Audio
}

func (s *CreateModelResponseBodyDataCapabilities) GetDocument() *bool {
	return s.Document
}

func (s *CreateModelResponseBodyDataCapabilities) GetMultiToolCall() *bool {
	return s.MultiToolCall
}

func (s *CreateModelResponseBodyDataCapabilities) GetReasoning() *bool {
	return s.Reasoning
}

func (s *CreateModelResponseBodyDataCapabilities) GetStreamToolCall() *bool {
	return s.StreamToolCall
}

func (s *CreateModelResponseBodyDataCapabilities) GetToolCall() *bool {
	return s.ToolCall
}

func (s *CreateModelResponseBodyDataCapabilities) GetVideo() *bool {
	return s.Video
}

func (s *CreateModelResponseBodyDataCapabilities) GetVision() *bool {
	return s.Vision
}

func (s *CreateModelResponseBodyDataCapabilities) SetAudio(v bool) *CreateModelResponseBodyDataCapabilities {
	s.Audio = &v
	return s
}

func (s *CreateModelResponseBodyDataCapabilities) SetDocument(v bool) *CreateModelResponseBodyDataCapabilities {
	s.Document = &v
	return s
}

func (s *CreateModelResponseBodyDataCapabilities) SetMultiToolCall(v bool) *CreateModelResponseBodyDataCapabilities {
	s.MultiToolCall = &v
	return s
}

func (s *CreateModelResponseBodyDataCapabilities) SetReasoning(v bool) *CreateModelResponseBodyDataCapabilities {
	s.Reasoning = &v
	return s
}

func (s *CreateModelResponseBodyDataCapabilities) SetStreamToolCall(v bool) *CreateModelResponseBodyDataCapabilities {
	s.StreamToolCall = &v
	return s
}

func (s *CreateModelResponseBodyDataCapabilities) SetToolCall(v bool) *CreateModelResponseBodyDataCapabilities {
	s.ToolCall = &v
	return s
}

func (s *CreateModelResponseBodyDataCapabilities) SetVideo(v bool) *CreateModelResponseBodyDataCapabilities {
	s.Video = &v
	return s
}

func (s *CreateModelResponseBodyDataCapabilities) SetVision(v bool) *CreateModelResponseBodyDataCapabilities {
	s.Vision = &v
	return s
}

func (s *CreateModelResponseBodyDataCapabilities) Validate() error {
	return dara.Validate(s)
}
