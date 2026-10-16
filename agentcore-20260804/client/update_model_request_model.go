// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateModelRequestBody) *UpdateModelRequest
	GetBody() *UpdateModelRequestBody
	SetClientToken(v string) *UpdateModelRequest
	GetClientToken() *string
}

type UpdateModelRequest struct {
	// The model update request body. You must specify at least one non-null parameter among description, contextSize, maxTokens, and capabilities.
	Body *UpdateModelRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The client token used for idempotence. This parameter is not supported.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateModelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelRequest) GetBody() *UpdateModelRequestBody {
	return s.Body
}

func (s *UpdateModelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateModelRequest) SetBody(v *UpdateModelRequestBody) *UpdateModelRequest {
	s.Body = v
	return s
}

func (s *UpdateModelRequest) SetClientToken(v string) *UpdateModelRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateModelRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateModelRequestBody struct {
	// The model capability configuration. When an object is passed, the original capability configuration is replaced as a whole. Capability fields that are not specified in the object are treated as false. Passing an empty object {} sets all capabilities to false. If this parameter is not passed or is set to null, the original configuration is retained.
	Capabilities *UpdateModelRequestBodyCapabilities `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Struct"`
	// The context token limit of the model. The minimum value is 1000. The updated value must not be less than maxTokens. If maxTokens is not specified in the current request, the existing value is used for validation. If this parameter is not passed or is set to null, the original value is retained.
	//
	// example:
	//
	// 131072
	ContextSize *int64 `json:"contextSize,omitempty" xml:"contextSize,omitempty"`
	// The model description. The maximum length is 255 characters after leading and trailing whitespace is removed. Passing an empty string clears the description. If this parameter is not passed or is set to null, the original value is retained. Modifying only the description does not refresh the model configurations of associated agents.
	//
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The maximum number of output tokens for a single generation of the model. The value must be a positive integer. If contextSize is configured, the updated maxTokens must not be greater than contextSize. If contextSize is not specified in the current request, the existing value is used for validation. If this parameter is not passed or is set to null, the original value is retained.
	//
	// example:
	//
	// 8192
	MaxTokens *int64 `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
}

func (s UpdateModelRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateModelRequestBody) GetCapabilities() *UpdateModelRequestBodyCapabilities {
	return s.Capabilities
}

func (s *UpdateModelRequestBody) GetContextSize() *int64 {
	return s.ContextSize
}

func (s *UpdateModelRequestBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelRequestBody) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *UpdateModelRequestBody) SetCapabilities(v *UpdateModelRequestBodyCapabilities) *UpdateModelRequestBody {
	s.Capabilities = v
	return s
}

func (s *UpdateModelRequestBody) SetContextSize(v int64) *UpdateModelRequestBody {
	s.ContextSize = &v
	return s
}

func (s *UpdateModelRequestBody) SetDescription(v string) *UpdateModelRequestBody {
	s.Description = &v
	return s
}

func (s *UpdateModelRequestBody) SetMaxTokens(v int64) *UpdateModelRequestBody {
	s.MaxTokens = &v
	return s
}

func (s *UpdateModelRequestBody) Validate() error {
	if s.Capabilities != nil {
		if err := s.Capabilities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateModelRequestBodyCapabilities struct {
	// Specifies whether the model supports audio input or output. Valid values:
	//
	// - true: Supported.
	//
	// - false: Not supported.
	Audio *bool `json:"audio,omitempty" xml:"audio,omitempty"`
	// Specifies whether the model supports document input. Valid values:
	//
	// - true: Supported.
	//
	// - false: Not supported.
	Document *bool `json:"document,omitempty" xml:"document,omitempty"`
	// Specifies whether the model is able to invoke tool calling multiple times in a single response. Valid values:
	//
	// - true: Supported.
	//
	// - false: Not supported.
	MultiToolCall *bool `json:"multiToolCall,omitempty" xml:"multiToolCall,omitempty"`
	// Specifies whether the model supports reasoning capabilities. Valid values:
	//
	// - true: Supported.
	//
	// - false: Not supported.
	//
	// This field is a capability flag and is not used to set the reasoning intensity or reasoning token budget.
	Reasoning *bool `json:"reasoning,omitempty" xml:"reasoning,omitempty"`
	// Specifies whether the model is able to invoke streaming tool calling. Valid values:
	//
	// - true: Supported.
	//
	// - false: Not supported.
	StreamToolCall *bool `json:"streamToolCall,omitempty" xml:"streamToolCall,omitempty"`
	// Specifies whether the model is able to invoke tool calling. Valid values:
	//
	// - true: Supported.
	//
	// - false: Not supported.
	ToolCall *bool `json:"toolCall,omitempty" xml:"toolCall,omitempty"`
	// Specifies whether the model supports video input. Valid values:
	//
	// - true: Supported.
	//
	// - false: Not supported.
	Video *bool `json:"video,omitempty" xml:"video,omitempty"`
	// Specifies whether the model supports image input. Valid values:
	//
	// - true: Supported.
	//
	// - false: Not supported.
	Vision *bool `json:"vision,omitempty" xml:"vision,omitempty"`
}

func (s UpdateModelRequestBodyCapabilities) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelRequestBodyCapabilities) GoString() string {
	return s.String()
}

func (s *UpdateModelRequestBodyCapabilities) GetAudio() *bool {
	return s.Audio
}

func (s *UpdateModelRequestBodyCapabilities) GetDocument() *bool {
	return s.Document
}

func (s *UpdateModelRequestBodyCapabilities) GetMultiToolCall() *bool {
	return s.MultiToolCall
}

func (s *UpdateModelRequestBodyCapabilities) GetReasoning() *bool {
	return s.Reasoning
}

func (s *UpdateModelRequestBodyCapabilities) GetStreamToolCall() *bool {
	return s.StreamToolCall
}

func (s *UpdateModelRequestBodyCapabilities) GetToolCall() *bool {
	return s.ToolCall
}

func (s *UpdateModelRequestBodyCapabilities) GetVideo() *bool {
	return s.Video
}

func (s *UpdateModelRequestBodyCapabilities) GetVision() *bool {
	return s.Vision
}

func (s *UpdateModelRequestBodyCapabilities) SetAudio(v bool) *UpdateModelRequestBodyCapabilities {
	s.Audio = &v
	return s
}

func (s *UpdateModelRequestBodyCapabilities) SetDocument(v bool) *UpdateModelRequestBodyCapabilities {
	s.Document = &v
	return s
}

func (s *UpdateModelRequestBodyCapabilities) SetMultiToolCall(v bool) *UpdateModelRequestBodyCapabilities {
	s.MultiToolCall = &v
	return s
}

func (s *UpdateModelRequestBodyCapabilities) SetReasoning(v bool) *UpdateModelRequestBodyCapabilities {
	s.Reasoning = &v
	return s
}

func (s *UpdateModelRequestBodyCapabilities) SetStreamToolCall(v bool) *UpdateModelRequestBodyCapabilities {
	s.StreamToolCall = &v
	return s
}

func (s *UpdateModelRequestBodyCapabilities) SetToolCall(v bool) *UpdateModelRequestBodyCapabilities {
	s.ToolCall = &v
	return s
}

func (s *UpdateModelRequestBodyCapabilities) SetVideo(v bool) *UpdateModelRequestBodyCapabilities {
	s.Video = &v
	return s
}

func (s *UpdateModelRequestBodyCapabilities) SetVision(v bool) *UpdateModelRequestBodyCapabilities {
	s.Vision = &v
	return s
}

func (s *UpdateModelRequestBodyCapabilities) Validate() error {
	return dara.Validate(s)
}
