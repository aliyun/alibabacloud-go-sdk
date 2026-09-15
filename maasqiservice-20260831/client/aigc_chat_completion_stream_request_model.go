// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAigcChatCompletionStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessages(v []*AigcChatCompletionStreamRequestMessages) *AigcChatCompletionStreamRequest
	GetMessages() []*AigcChatCompletionStreamRequestMessages
	SetMetadata(v *AigcChatCompletionStreamRequestMetadata) *AigcChatCompletionStreamRequest
	GetMetadata() *AigcChatCompletionStreamRequestMetadata
	SetModel(v string) *AigcChatCompletionStreamRequest
	GetModel() *string
	SetStream(v bool) *AigcChatCompletionStreamRequest
	GetStream() *bool
	SetStreamOptions(v *AigcChatCompletionStreamRequestStreamOptions) *AigcChatCompletionStreamRequest
	GetStreamOptions() *AigcChatCompletionStreamRequestStreamOptions
}

type AigcChatCompletionStreamRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"role":"user","content":"生成一张水墨山水画"}]
	Messages []*AigcChatCompletionStreamRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// example:
	//
	// {"parameters":{"size":"1024*1024","n":1}}
	Metadata *AigcChatCompletionStreamRequestMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	Model    *string                                  `json:"model,omitempty" xml:"model,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// example:
	//
	// {"include_usage":true}
	StreamOptions *AigcChatCompletionStreamRequestStreamOptions `json:"streamOptions,omitempty" xml:"streamOptions,omitempty" type:"Struct"`
}

func (s AigcChatCompletionStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s AigcChatCompletionStreamRequest) GoString() string {
	return s.String()
}

func (s *AigcChatCompletionStreamRequest) GetMessages() []*AigcChatCompletionStreamRequestMessages {
	return s.Messages
}

func (s *AigcChatCompletionStreamRequest) GetMetadata() *AigcChatCompletionStreamRequestMetadata {
	return s.Metadata
}

func (s *AigcChatCompletionStreamRequest) GetModel() *string {
	return s.Model
}

func (s *AigcChatCompletionStreamRequest) GetStream() *bool {
	return s.Stream
}

func (s *AigcChatCompletionStreamRequest) GetStreamOptions() *AigcChatCompletionStreamRequestStreamOptions {
	return s.StreamOptions
}

func (s *AigcChatCompletionStreamRequest) SetMessages(v []*AigcChatCompletionStreamRequestMessages) *AigcChatCompletionStreamRequest {
	s.Messages = v
	return s
}

func (s *AigcChatCompletionStreamRequest) SetMetadata(v *AigcChatCompletionStreamRequestMetadata) *AigcChatCompletionStreamRequest {
	s.Metadata = v
	return s
}

func (s *AigcChatCompletionStreamRequest) SetModel(v string) *AigcChatCompletionStreamRequest {
	s.Model = &v
	return s
}

func (s *AigcChatCompletionStreamRequest) SetStream(v bool) *AigcChatCompletionStreamRequest {
	s.Stream = &v
	return s
}

func (s *AigcChatCompletionStreamRequest) SetStreamOptions(v *AigcChatCompletionStreamRequestStreamOptions) *AigcChatCompletionStreamRequest {
	s.StreamOptions = v
	return s
}

func (s *AigcChatCompletionStreamRequest) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	if s.StreamOptions != nil {
		if err := s.StreamOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AigcChatCompletionStreamRequestMessages struct {
	// example:
	//
	// 生成一张雨后未来城市夜景
	Content []*AigcChatCompletionStreamRequestMessagesContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s AigcChatCompletionStreamRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s AigcChatCompletionStreamRequestMessages) GoString() string {
	return s.String()
}

func (s *AigcChatCompletionStreamRequestMessages) GetContent() []*AigcChatCompletionStreamRequestMessagesContent {
	return s.Content
}

func (s *AigcChatCompletionStreamRequestMessages) GetRole() *string {
	return s.Role
}

func (s *AigcChatCompletionStreamRequestMessages) SetContent(v []*AigcChatCompletionStreamRequestMessagesContent) *AigcChatCompletionStreamRequestMessages {
	s.Content = v
	return s
}

func (s *AigcChatCompletionStreamRequestMessages) SetRole(v string) *AigcChatCompletionStreamRequestMessages {
	s.Role = &v
	return s
}

func (s *AigcChatCompletionStreamRequestMessages) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AigcChatCompletionStreamRequestMessagesContent struct {
	ImageUrl *AigcChatCompletionStreamRequestMessagesContentImageUrl `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Struct"`
	// example:
	//
	// 把参考图背景改成海边日落
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// text
	Type     *string                                                 `json:"type,omitempty" xml:"type,omitempty"`
	VideoUrl *AigcChatCompletionStreamRequestMessagesContentVideoUrl `json:"videoUrl,omitempty" xml:"videoUrl,omitempty" type:"Struct"`
}

func (s AigcChatCompletionStreamRequestMessagesContent) String() string {
	return dara.Prettify(s)
}

func (s AigcChatCompletionStreamRequestMessagesContent) GoString() string {
	return s.String()
}

func (s *AigcChatCompletionStreamRequestMessagesContent) GetImageUrl() *AigcChatCompletionStreamRequestMessagesContentImageUrl {
	return s.ImageUrl
}

func (s *AigcChatCompletionStreamRequestMessagesContent) GetText() *string {
	return s.Text
}

func (s *AigcChatCompletionStreamRequestMessagesContent) GetType() *string {
	return s.Type
}

func (s *AigcChatCompletionStreamRequestMessagesContent) GetVideoUrl() *AigcChatCompletionStreamRequestMessagesContentVideoUrl {
	return s.VideoUrl
}

func (s *AigcChatCompletionStreamRequestMessagesContent) SetImageUrl(v *AigcChatCompletionStreamRequestMessagesContentImageUrl) *AigcChatCompletionStreamRequestMessagesContent {
	s.ImageUrl = v
	return s
}

func (s *AigcChatCompletionStreamRequestMessagesContent) SetText(v string) *AigcChatCompletionStreamRequestMessagesContent {
	s.Text = &v
	return s
}

func (s *AigcChatCompletionStreamRequestMessagesContent) SetType(v string) *AigcChatCompletionStreamRequestMessagesContent {
	s.Type = &v
	return s
}

func (s *AigcChatCompletionStreamRequestMessagesContent) SetVideoUrl(v *AigcChatCompletionStreamRequestMessagesContentVideoUrl) *AigcChatCompletionStreamRequestMessagesContent {
	s.VideoUrl = v
	return s
}

func (s *AigcChatCompletionStreamRequestMessagesContent) Validate() error {
	if s.ImageUrl != nil {
		if err := s.ImageUrl.Validate(); err != nil {
			return err
		}
	}
	if s.VideoUrl != nil {
		if err := s.VideoUrl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AigcChatCompletionStreamRequestMessagesContentImageUrl struct {
	// example:
	//
	// https://example.com/input.png
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s AigcChatCompletionStreamRequestMessagesContentImageUrl) String() string {
	return dara.Prettify(s)
}

func (s AigcChatCompletionStreamRequestMessagesContentImageUrl) GoString() string {
	return s.String()
}

func (s *AigcChatCompletionStreamRequestMessagesContentImageUrl) GetUrl() *string {
	return s.Url
}

func (s *AigcChatCompletionStreamRequestMessagesContentImageUrl) SetUrl(v string) *AigcChatCompletionStreamRequestMessagesContentImageUrl {
	s.Url = &v
	return s
}

func (s *AigcChatCompletionStreamRequestMessagesContentImageUrl) Validate() error {
	return dara.Validate(s)
}

type AigcChatCompletionStreamRequestMessagesContentVideoUrl struct {
	// example:
	//
	// https://example.com/input.mp4
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s AigcChatCompletionStreamRequestMessagesContentVideoUrl) String() string {
	return dara.Prettify(s)
}

func (s AigcChatCompletionStreamRequestMessagesContentVideoUrl) GoString() string {
	return s.String()
}

func (s *AigcChatCompletionStreamRequestMessagesContentVideoUrl) GetUrl() *string {
	return s.Url
}

func (s *AigcChatCompletionStreamRequestMessagesContentVideoUrl) SetUrl(v string) *AigcChatCompletionStreamRequestMessagesContentVideoUrl {
	s.Url = &v
	return s
}

func (s *AigcChatCompletionStreamRequestMessagesContentVideoUrl) Validate() error {
	return dara.Validate(s)
}

type AigcChatCompletionStreamRequestMetadata struct {
	// example:
	//
	// {"size":"1024*1024","n":1,"num_inference_steps":8}
	Parameters *AigcChatCompletionStreamRequestMetadataParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
}

func (s AigcChatCompletionStreamRequestMetadata) String() string {
	return dara.Prettify(s)
}

func (s AigcChatCompletionStreamRequestMetadata) GoString() string {
	return s.String()
}

func (s *AigcChatCompletionStreamRequestMetadata) GetParameters() *AigcChatCompletionStreamRequestMetadataParameters {
	return s.Parameters
}

func (s *AigcChatCompletionStreamRequestMetadata) SetParameters(v *AigcChatCompletionStreamRequestMetadataParameters) *AigcChatCompletionStreamRequestMetadata {
	s.Parameters = v
	return s
}

func (s *AigcChatCompletionStreamRequestMetadata) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AigcChatCompletionStreamRequestMetadataParameters struct {
	// example:
	//
	// 0.0
	GuidanceScale *float64 `json:"guidanceScale,omitempty" xml:"guidanceScale,omitempty"`
	// example:
	//
	// 1
	N *int64 `json:"n,omitempty" xml:"n,omitempty"`
	// example:
	//
	// 低质量、模糊、文字、水印
	NegativePrompt *string `json:"negativePrompt,omitempty" xml:"negativePrompt,omitempty"`
	// example:
	//
	// 8
	NumInferenceSteps *int64 `json:"numInferenceSteps,omitempty" xml:"numInferenceSteps,omitempty"`
	// example:
	//
	// 42
	Seed *int64 `json:"seed,omitempty" xml:"seed,omitempty"`
	// example:
	//
	// 1024*1024
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s AigcChatCompletionStreamRequestMetadataParameters) String() string {
	return dara.Prettify(s)
}

func (s AigcChatCompletionStreamRequestMetadataParameters) GoString() string {
	return s.String()
}

func (s *AigcChatCompletionStreamRequestMetadataParameters) GetGuidanceScale() *float64 {
	return s.GuidanceScale
}

func (s *AigcChatCompletionStreamRequestMetadataParameters) GetN() *int64 {
	return s.N
}

func (s *AigcChatCompletionStreamRequestMetadataParameters) GetNegativePrompt() *string {
	return s.NegativePrompt
}

func (s *AigcChatCompletionStreamRequestMetadataParameters) GetNumInferenceSteps() *int64 {
	return s.NumInferenceSteps
}

func (s *AigcChatCompletionStreamRequestMetadataParameters) GetSeed() *int64 {
	return s.Seed
}

func (s *AigcChatCompletionStreamRequestMetadataParameters) GetSize() *string {
	return s.Size
}

func (s *AigcChatCompletionStreamRequestMetadataParameters) SetGuidanceScale(v float64) *AigcChatCompletionStreamRequestMetadataParameters {
	s.GuidanceScale = &v
	return s
}

func (s *AigcChatCompletionStreamRequestMetadataParameters) SetN(v int64) *AigcChatCompletionStreamRequestMetadataParameters {
	s.N = &v
	return s
}

func (s *AigcChatCompletionStreamRequestMetadataParameters) SetNegativePrompt(v string) *AigcChatCompletionStreamRequestMetadataParameters {
	s.NegativePrompt = &v
	return s
}

func (s *AigcChatCompletionStreamRequestMetadataParameters) SetNumInferenceSteps(v int64) *AigcChatCompletionStreamRequestMetadataParameters {
	s.NumInferenceSteps = &v
	return s
}

func (s *AigcChatCompletionStreamRequestMetadataParameters) SetSeed(v int64) *AigcChatCompletionStreamRequestMetadataParameters {
	s.Seed = &v
	return s
}

func (s *AigcChatCompletionStreamRequestMetadataParameters) SetSize(v string) *AigcChatCompletionStreamRequestMetadataParameters {
	s.Size = &v
	return s
}

func (s *AigcChatCompletionStreamRequestMetadataParameters) Validate() error {
	return dara.Validate(s)
}

type AigcChatCompletionStreamRequestStreamOptions struct {
	// example:
	//
	// true
	IncludeUsage *bool `json:"includeUsage,omitempty" xml:"includeUsage,omitempty"`
}

func (s AigcChatCompletionStreamRequestStreamOptions) String() string {
	return dara.Prettify(s)
}

func (s AigcChatCompletionStreamRequestStreamOptions) GoString() string {
	return s.String()
}

func (s *AigcChatCompletionStreamRequestStreamOptions) GetIncludeUsage() *bool {
	return s.IncludeUsage
}

func (s *AigcChatCompletionStreamRequestStreamOptions) SetIncludeUsage(v bool) *AigcChatCompletionStreamRequestStreamOptions {
	s.IncludeUsage = &v
	return s
}

func (s *AigcChatCompletionStreamRequestStreamOptions) Validate() error {
	return dara.Validate(s)
}
