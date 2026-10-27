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
	// The conversation message array. Creative Agent uses the last user message as the current objective and executes it in combination with the session history.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"role":"user","content":"Generate an ink wash landscape painting"}]
	Messages []*AigcChatCompletionStreamRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// The extended input for this image creation request. Use the parameters object to set generation parameters. Do not use this field to pass model addresses or keys.
	//
	// example:
	//
	// {"parameters":{"size":"1024*1024","n":1}}
	Metadata *AigcChatCompletionStreamRequestMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// A compatibility placeholder field. Can be omitted or set to an empty string. You cannot use this field to switch the model used by the instance. The model field in the response is an empty string.
	//
	// example:
	//
	// " "
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// Specifies whether to use SSE streaming responses. This parameter must be set to true for this operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// The HTTP JSON key is stream_options. Set include_usage to retrieve the token usage for this call.
	//
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
	// The message content. Can be text. A multimodal array can contain text, image_url, or video_url. The Creative Agent Lite edition performs image-to-image when a reference image is provided and text-to-image when no reference image is provided. It does not accept video.
	//
	// example:
	//
	// Generate a futuristic city nightscape after rain
	Content []*AigcChatCompletionStreamRequestMessagesContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// The message role. Place the current generation objective in a user message. History can include user, assistant, and tool.
	//
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
	// The reference image when type is set to image_url. Supports HTTPS URLs or Base64 Data URLs.
	ImageUrl *AigcChatCompletionStreamRequestMessagesContentImageUrl `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Struct"`
	// The generation or editing objective when type is set to text.
	//
	// example:
	//
	// Change the background of the reference image to a seaside sunset
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// The content part type: text, image_url, or video_url.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The video input when type is set to video_url. Only accepted by the Creative Agent Dedicated edition Planner. The Creative Agent Lite edition rejects this type.
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
	// The publicly accessible HTTPS URL of the reference image, or a complete Base64 image Data URL (such as data:image/png;base64,...). Signed URLs must remain valid during request execution. Re-upload or update the link for expired or unauthorized images.
	//
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
	// The video URL.
	//
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
	// The image parameters for this request. Only size, n, seed, negative_prompt, num_inference_steps, and guidance_scale are supported. Omitted fields use the default configurations. response_format is not supported.
	//
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
	// The guidance scale. Must be a finite number between 0 and 20.
	//
	// example:
	//
	// 0.0
	GuidanceScale *float64 `json:"guidanceScale,omitempty" xml:"guidanceScale,omitempty"`
	// The number of images to generate for this request. Must be an integer in the range of 1 to 6.
	//
	// example:
	//
	// 1
	N *int64 `json:"n,omitempty" xml:"n,omitempty"`
	// The negative prompt that describes content you do not want to appear in the result.
	//
	// example:
	//
	// low quality, blurry, text, watermark
	NegativePrompt *string `json:"negativePrompt,omitempty" xml:"negativePrompt,omitempty"`
	// The number of model inference iterations. Must be an integer in the range of 1 to 200.
	//
	// example:
	//
	// 8
	NumInferenceSteps *int64 `json:"numInferenceSteps,omitempty" xml:"numInferenceSteps,omitempty"`
	// The random seed. Must be an integer in the range of 0 to 2147483647.
	//
	// example:
	//
	// 42
	Seed *int64 `json:"seed,omitempty" xml:"seed,omitempty"`
	// The output pixel dimensions, such as `1024*1024` or `1280*720`. The format `1024x1024` is also accepted. When specified, this value overrides the default canvas dimensions for this request. Use width and height values supported by the model.
	//
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
	// The HTTP JSON key is include_usage. When set to true, an independent usage chunk is returned before the end marker. The choices field in this chunk is an empty array. Token usage does not represent the number of images or the final bill.
	//
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
