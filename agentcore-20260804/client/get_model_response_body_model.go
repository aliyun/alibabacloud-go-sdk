// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetModelResponseBody
	GetCode() *string
	SetData(v *GetModelResponseBodyData) *GetModelResponseBody
	GetData() *GetModelResponseBodyData
	SetHttpStatusCode(v int32) *GetModelResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetModelResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetModelResponseBody
	GetSuccess() *bool
}

type GetModelResponseBody struct {
	// The business status code. A value of SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The model details.
	Data *GetModelResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. A value of 200 indicates success.
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

func (s GetModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetModelResponseBody) GetData() *GetModelResponseBodyData {
	return s.Data
}

func (s *GetModelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetModelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetModelResponseBody) SetCode(v string) *GetModelResponseBody {
	s.Code = &v
	return s
}

func (s *GetModelResponseBody) SetData(v *GetModelResponseBodyData) *GetModelResponseBody {
	s.Data = v
	return s
}

func (s *GetModelResponseBody) SetHttpStatusCode(v int32) *GetModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetModelResponseBody) SetMessage(v string) *GetModelResponseBody {
	s.Message = &v
	return s
}

func (s *GetModelResponseBody) SetRequestId(v string) *GetModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelResponseBody) SetSuccess(v bool) *GetModelResponseBody {
	s.Success = &v
	return s
}

func (s *GetModelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModelResponseBodyData struct {
	// The model capability configuration.
	Capabilities *GetModelResponseBodyDataCapabilities `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Struct"`
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
	// The model description, up to 255 characters.
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
	// The region ID to which the resource belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
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

func (s GetModelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetModelResponseBodyData) GetCapabilities() *GetModelResponseBodyDataCapabilities {
	return s.Capabilities
}

func (s *GetModelResponseBodyData) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *GetModelResponseBodyData) GetContextSize() *int64 {
	return s.ContextSize
}

func (s *GetModelResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetModelResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetModelResponseBodyData) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *GetModelResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *GetModelResponseBodyData) GetModelName() *string {
	return s.ModelName
}

func (s *GetModelResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetModelResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetModelResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetModelResponseBodyData) SetCapabilities(v *GetModelResponseBodyDataCapabilities) *GetModelResponseBodyData {
	s.Capabilities = v
	return s
}

func (s *GetModelResponseBodyData) SetConnectionId(v string) *GetModelResponseBodyData {
	s.ConnectionId = &v
	return s
}

func (s *GetModelResponseBodyData) SetContextSize(v int64) *GetModelResponseBodyData {
	s.ContextSize = &v
	return s
}

func (s *GetModelResponseBodyData) SetCreatedAt(v string) *GetModelResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetModelResponseBodyData) SetDescription(v string) *GetModelResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetModelResponseBodyData) SetMaxTokens(v int64) *GetModelResponseBodyData {
	s.MaxTokens = &v
	return s
}

func (s *GetModelResponseBodyData) SetModelId(v string) *GetModelResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *GetModelResponseBodyData) SetModelName(v string) *GetModelResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *GetModelResponseBodyData) SetRegionId(v string) *GetModelResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetModelResponseBodyData) SetUpdatedAt(v string) *GetModelResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetModelResponseBodyData) SetWorkspaceId(v string) *GetModelResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetModelResponseBodyData) Validate() error {
	if s.Capabilities != nil {
		if err := s.Capabilities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModelResponseBodyDataCapabilities struct {
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

func (s GetModelResponseBodyDataCapabilities) String() string {
	return dara.Prettify(s)
}

func (s GetModelResponseBodyDataCapabilities) GoString() string {
	return s.String()
}

func (s *GetModelResponseBodyDataCapabilities) GetAudio() *bool {
	return s.Audio
}

func (s *GetModelResponseBodyDataCapabilities) GetDocument() *bool {
	return s.Document
}

func (s *GetModelResponseBodyDataCapabilities) GetMultiToolCall() *bool {
	return s.MultiToolCall
}

func (s *GetModelResponseBodyDataCapabilities) GetReasoning() *bool {
	return s.Reasoning
}

func (s *GetModelResponseBodyDataCapabilities) GetStreamToolCall() *bool {
	return s.StreamToolCall
}

func (s *GetModelResponseBodyDataCapabilities) GetToolCall() *bool {
	return s.ToolCall
}

func (s *GetModelResponseBodyDataCapabilities) GetVideo() *bool {
	return s.Video
}

func (s *GetModelResponseBodyDataCapabilities) GetVision() *bool {
	return s.Vision
}

func (s *GetModelResponseBodyDataCapabilities) SetAudio(v bool) *GetModelResponseBodyDataCapabilities {
	s.Audio = &v
	return s
}

func (s *GetModelResponseBodyDataCapabilities) SetDocument(v bool) *GetModelResponseBodyDataCapabilities {
	s.Document = &v
	return s
}

func (s *GetModelResponseBodyDataCapabilities) SetMultiToolCall(v bool) *GetModelResponseBodyDataCapabilities {
	s.MultiToolCall = &v
	return s
}

func (s *GetModelResponseBodyDataCapabilities) SetReasoning(v bool) *GetModelResponseBodyDataCapabilities {
	s.Reasoning = &v
	return s
}

func (s *GetModelResponseBodyDataCapabilities) SetStreamToolCall(v bool) *GetModelResponseBodyDataCapabilities {
	s.StreamToolCall = &v
	return s
}

func (s *GetModelResponseBodyDataCapabilities) SetToolCall(v bool) *GetModelResponseBodyDataCapabilities {
	s.ToolCall = &v
	return s
}

func (s *GetModelResponseBodyDataCapabilities) SetVideo(v bool) *GetModelResponseBodyDataCapabilities {
	s.Video = &v
	return s
}

func (s *GetModelResponseBodyDataCapabilities) SetVision(v bool) *GetModelResponseBodyDataCapabilities {
	s.Vision = &v
	return s
}

func (s *GetModelResponseBodyDataCapabilities) Validate() error {
	return dara.Validate(s)
}
