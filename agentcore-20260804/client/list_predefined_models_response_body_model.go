// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPredefinedModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPredefinedModelsResponseBody
	GetCode() *string
	SetData(v []*ListPredefinedModelsResponseBodyData) *ListPredefinedModelsResponseBody
	GetData() []*ListPredefinedModelsResponseBodyData
	SetHttpStatusCode(v int32) *ListPredefinedModelsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListPredefinedModelsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPredefinedModelsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPredefinedModelsResponseBody
	GetSuccess() *bool
}

type ListPredefinedModelsResponseBody struct {
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of predefined models.
	Data []*ListPredefinedModelsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s ListPredefinedModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPredefinedModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPredefinedModelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPredefinedModelsResponseBody) GetData() []*ListPredefinedModelsResponseBodyData {
	return s.Data
}

func (s *ListPredefinedModelsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPredefinedModelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPredefinedModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPredefinedModelsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPredefinedModelsResponseBody) SetCode(v string) *ListPredefinedModelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPredefinedModelsResponseBody) SetData(v []*ListPredefinedModelsResponseBodyData) *ListPredefinedModelsResponseBody {
	s.Data = v
	return s
}

func (s *ListPredefinedModelsResponseBody) SetHttpStatusCode(v int32) *ListPredefinedModelsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPredefinedModelsResponseBody) SetMessage(v string) *ListPredefinedModelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPredefinedModelsResponseBody) SetRequestId(v string) *ListPredefinedModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPredefinedModelsResponseBody) SetSuccess(v bool) *ListPredefinedModelsResponseBody {
	s.Success = &v
	return s
}

func (s *ListPredefinedModelsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPredefinedModelsResponseBodyData struct {
	// The model capability configuration.
	Capabilities *ListPredefinedModelsResponseBodyDataCapabilities `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Struct"`
	// The model context window size in tokens. The value must be a positive integer.
	//
	// example:
	//
	// 128000
	ContextSize *int64 `json:"contextSize,omitempty" xml:"contextSize,omitempty"`
	// The maximum number of output tokens supported by the model in a single generation.
	//
	// example:
	//
	// 131072
	MaxTokens *int64 `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	// The upstream model name.
	//
	// example:
	//
	// qwen-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The model release date in the format of YYYY-MM-DD.
	//
	// example:
	//
	// 2025-02-01
	ReleaseDate *string `json:"releaseDate,omitempty" xml:"releaseDate,omitempty"`
}

func (s ListPredefinedModelsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPredefinedModelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPredefinedModelsResponseBodyData) GetCapabilities() *ListPredefinedModelsResponseBodyDataCapabilities {
	return s.Capabilities
}

func (s *ListPredefinedModelsResponseBodyData) GetContextSize() *int64 {
	return s.ContextSize
}

func (s *ListPredefinedModelsResponseBodyData) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *ListPredefinedModelsResponseBodyData) GetModelName() *string {
	return s.ModelName
}

func (s *ListPredefinedModelsResponseBodyData) GetReleaseDate() *string {
	return s.ReleaseDate
}

func (s *ListPredefinedModelsResponseBodyData) SetCapabilities(v *ListPredefinedModelsResponseBodyDataCapabilities) *ListPredefinedModelsResponseBodyData {
	s.Capabilities = v
	return s
}

func (s *ListPredefinedModelsResponseBodyData) SetContextSize(v int64) *ListPredefinedModelsResponseBodyData {
	s.ContextSize = &v
	return s
}

func (s *ListPredefinedModelsResponseBodyData) SetMaxTokens(v int64) *ListPredefinedModelsResponseBodyData {
	s.MaxTokens = &v
	return s
}

func (s *ListPredefinedModelsResponseBodyData) SetModelName(v string) *ListPredefinedModelsResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *ListPredefinedModelsResponseBodyData) SetReleaseDate(v string) *ListPredefinedModelsResponseBodyData {
	s.ReleaseDate = &v
	return s
}

func (s *ListPredefinedModelsResponseBodyData) Validate() error {
	if s.Capabilities != nil {
		if err := s.Capabilities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPredefinedModelsResponseBodyDataCapabilities struct {
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

func (s ListPredefinedModelsResponseBodyDataCapabilities) String() string {
	return dara.Prettify(s)
}

func (s ListPredefinedModelsResponseBodyDataCapabilities) GoString() string {
	return s.String()
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) GetAudio() *bool {
	return s.Audio
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) GetDocument() *bool {
	return s.Document
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) GetMultiToolCall() *bool {
	return s.MultiToolCall
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) GetReasoning() *bool {
	return s.Reasoning
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) GetStreamToolCall() *bool {
	return s.StreamToolCall
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) GetToolCall() *bool {
	return s.ToolCall
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) GetVideo() *bool {
	return s.Video
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) GetVision() *bool {
	return s.Vision
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) SetAudio(v bool) *ListPredefinedModelsResponseBodyDataCapabilities {
	s.Audio = &v
	return s
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) SetDocument(v bool) *ListPredefinedModelsResponseBodyDataCapabilities {
	s.Document = &v
	return s
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) SetMultiToolCall(v bool) *ListPredefinedModelsResponseBodyDataCapabilities {
	s.MultiToolCall = &v
	return s
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) SetReasoning(v bool) *ListPredefinedModelsResponseBodyDataCapabilities {
	s.Reasoning = &v
	return s
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) SetStreamToolCall(v bool) *ListPredefinedModelsResponseBodyDataCapabilities {
	s.StreamToolCall = &v
	return s
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) SetToolCall(v bool) *ListPredefinedModelsResponseBodyDataCapabilities {
	s.ToolCall = &v
	return s
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) SetVideo(v bool) *ListPredefinedModelsResponseBodyDataCapabilities {
	s.Video = &v
	return s
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) SetVision(v bool) *ListPredefinedModelsResponseBodyDataCapabilities {
	s.Vision = &v
	return s
}

func (s *ListPredefinedModelsResponseBodyDataCapabilities) Validate() error {
	return dara.Validate(s)
}
