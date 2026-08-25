// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateModelRequestBody) *CreateModelRequest
	GetBody() *CreateModelRequestBody
	SetClientToken(v string) *CreateModelRequest
	GetClientToken() *string
}

type CreateModelRequest struct {
	// The request body.
	Body *CreateModelRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The client token for idempotence. Not currently supported.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateModelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequest) GoString() string {
	return s.String()
}

func (s *CreateModelRequest) GetBody() *CreateModelRequestBody {
	return s.Body
}

func (s *CreateModelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateModelRequest) SetBody(v *CreateModelRequestBody) *CreateModelRequest {
	s.Body = v
	return s
}

func (s *CreateModelRequest) SetClientToken(v string) *CreateModelRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModelRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateModelRequestBody struct {
	// The model capability configuration.
	Capabilities *CreateModelRequestBodyCapabilities `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Struct"`
	// The model connection ID.
	//
	// This parameter is required.
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
	// The upstream model name.
	//
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
}

func (s CreateModelRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequestBody) GoString() string {
	return s.String()
}

func (s *CreateModelRequestBody) GetCapabilities() *CreateModelRequestBodyCapabilities {
	return s.Capabilities
}

func (s *CreateModelRequestBody) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *CreateModelRequestBody) GetContextSize() *int64 {
	return s.ContextSize
}

func (s *CreateModelRequestBody) GetDescription() *string {
	return s.Description
}

func (s *CreateModelRequestBody) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *CreateModelRequestBody) GetModelName() *string {
	return s.ModelName
}

func (s *CreateModelRequestBody) SetCapabilities(v *CreateModelRequestBodyCapabilities) *CreateModelRequestBody {
	s.Capabilities = v
	return s
}

func (s *CreateModelRequestBody) SetConnectionId(v string) *CreateModelRequestBody {
	s.ConnectionId = &v
	return s
}

func (s *CreateModelRequestBody) SetContextSize(v int64) *CreateModelRequestBody {
	s.ContextSize = &v
	return s
}

func (s *CreateModelRequestBody) SetDescription(v string) *CreateModelRequestBody {
	s.Description = &v
	return s
}

func (s *CreateModelRequestBody) SetMaxTokens(v int64) *CreateModelRequestBody {
	s.MaxTokens = &v
	return s
}

func (s *CreateModelRequestBody) SetModelName(v string) *CreateModelRequestBody {
	s.ModelName = &v
	return s
}

func (s *CreateModelRequestBody) Validate() error {
	if s.Capabilities != nil {
		if err := s.Capabilities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateModelRequestBodyCapabilities struct {
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

func (s CreateModelRequestBodyCapabilities) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequestBodyCapabilities) GoString() string {
	return s.String()
}

func (s *CreateModelRequestBodyCapabilities) GetAudio() *bool {
	return s.Audio
}

func (s *CreateModelRequestBodyCapabilities) GetDocument() *bool {
	return s.Document
}

func (s *CreateModelRequestBodyCapabilities) GetMultiToolCall() *bool {
	return s.MultiToolCall
}

func (s *CreateModelRequestBodyCapabilities) GetReasoning() *bool {
	return s.Reasoning
}

func (s *CreateModelRequestBodyCapabilities) GetStreamToolCall() *bool {
	return s.StreamToolCall
}

func (s *CreateModelRequestBodyCapabilities) GetToolCall() *bool {
	return s.ToolCall
}

func (s *CreateModelRequestBodyCapabilities) GetVideo() *bool {
	return s.Video
}

func (s *CreateModelRequestBodyCapabilities) GetVision() *bool {
	return s.Vision
}

func (s *CreateModelRequestBodyCapabilities) SetAudio(v bool) *CreateModelRequestBodyCapabilities {
	s.Audio = &v
	return s
}

func (s *CreateModelRequestBodyCapabilities) SetDocument(v bool) *CreateModelRequestBodyCapabilities {
	s.Document = &v
	return s
}

func (s *CreateModelRequestBodyCapabilities) SetMultiToolCall(v bool) *CreateModelRequestBodyCapabilities {
	s.MultiToolCall = &v
	return s
}

func (s *CreateModelRequestBodyCapabilities) SetReasoning(v bool) *CreateModelRequestBodyCapabilities {
	s.Reasoning = &v
	return s
}

func (s *CreateModelRequestBodyCapabilities) SetStreamToolCall(v bool) *CreateModelRequestBodyCapabilities {
	s.StreamToolCall = &v
	return s
}

func (s *CreateModelRequestBodyCapabilities) SetToolCall(v bool) *CreateModelRequestBodyCapabilities {
	s.ToolCall = &v
	return s
}

func (s *CreateModelRequestBodyCapabilities) SetVideo(v bool) *CreateModelRequestBodyCapabilities {
	s.Video = &v
	return s
}

func (s *CreateModelRequestBodyCapabilities) SetVision(v bool) *CreateModelRequestBodyCapabilities {
	s.Vision = &v
	return s
}

func (s *CreateModelRequestBodyCapabilities) Validate() error {
	return dara.Validate(s)
}
