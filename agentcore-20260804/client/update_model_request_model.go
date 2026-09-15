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
	// The model update request body. At least one non-null parameter must be provided among description, contextSize, maxTokens, and capabilities.
	Body *UpdateModelRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The client token for idempotency. Not currently supported.
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
	// The model capability configuration. When an object is provided, it replaces the existing capability configuration as a whole. Capability fields not included in the object are treated as false. Providing an empty object {} sets all capabilities to false. If this parameter is not provided or set to null, the original configuration is retained.
	Capabilities *UpdateModelRequestBodyCapabilities `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Struct"`
	// The context token limit of the model. The minimum value is 1000. The updated value must not be less than maxTokens. If maxTokens is not provided in this request, the existing value is used for validation. If this parameter is not provided or set to null, the original value is retained.
	//
	// example:
	//
	// 131072
	ContextSize *int64 `json:"contextSize,omitempty" xml:"contextSize,omitempty"`
	// The model description. The maximum length is 255 characters after leading and trailing whitespace is removed. Providing an empty string clears the description. If this parameter is not provided or set to null, the original value is retained. Modifying only the description does not refresh the model configuration of associated Agents.
	//
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The maximum number of output tokens per generation. The value must be a positive integer. If contextSize is configured, the updated maxTokens must not exceed contextSize. If contextSize is not provided in this request, the existing value is used for validation. If this parameter is not provided or set to null, the original value is retained.
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
	// Specifies whether the model supports audio input or output. A value of true indicates that it is supported. A value of false indicates that it is not supported.
	Audio *bool `json:"audio,omitempty" xml:"audio,omitempty"`
	// Specifies whether the model supports document input. A value of true indicates that it is supported. A value of false indicates that it is not supported.
	Document *bool `json:"document,omitempty" xml:"document,omitempty"`
	// Specifies whether the model supports invoking multiple tools in a single response. A value of true indicates that it is supported. A value of false indicates that it is not supported.
	MultiToolCall *bool `json:"multiToolCall,omitempty" xml:"multiToolCall,omitempty"`
	// Specifies whether the model supports reasoning. A value of true indicates that it is supported. A value of false indicates that it is not supported. This field is a capability marker and is not used to set reasoning intensity or reasoning token budget.
	Reasoning *bool `json:"reasoning,omitempty" xml:"reasoning,omitempty"`
	// Specifies whether the model supports streaming tool invocation. A value of true indicates that it is supported. A value of false indicates that it is not supported.
	StreamToolCall *bool `json:"streamToolCall,omitempty" xml:"streamToolCall,omitempty"`
	// Specifies whether the model supports tool invocation. A value of true indicates that it is supported. A value of false indicates that it is not supported.
	ToolCall *bool `json:"toolCall,omitempty" xml:"toolCall,omitempty"`
	// Specifies whether the model supports video input. A value of true indicates that it is supported. A value of false indicates that it is not supported.
	Video *bool `json:"video,omitempty" xml:"video,omitempty"`
	// Specifies whether the model supports image input. A value of true indicates that it is supported. A value of false indicates that it is not supported.
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
