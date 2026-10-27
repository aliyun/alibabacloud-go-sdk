// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCuaChatCompletionStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessages(v []*CuaChatCompletionStreamRequestMessages) *CuaChatCompletionStreamRequest
	GetMessages() []*CuaChatCompletionStreamRequestMessages
	SetMetadata(v *CuaChatCompletionStreamRequestMetadata) *CuaChatCompletionStreamRequest
	GetMetadata() *CuaChatCompletionStreamRequestMetadata
	SetModel(v string) *CuaChatCompletionStreamRequest
	GetModel() *string
	SetStream(v bool) *CuaChatCompletionStreamRequest
	GetStream() *bool
	SetStreamOptions(v *CuaChatCompletionStreamRequestStreamOptions) *CuaChatCompletionStreamRequest
	GetStreamOptions() *CuaChatCompletionStreamRequestStreamOptions
	SetTools(v []map[string]interface{}) *CuaChatCompletionStreamRequest
	GetTools() []map[string]interface{}
}

type CuaChatCompletionStreamRequest struct {
	// This parameter is required.
	Messages []*CuaChatCompletionStreamRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	Metadata *CuaChatCompletionStreamRequestMetadata   `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// example:
	//
	// " "
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Stream        *bool                                        `json:"stream,omitempty" xml:"stream,omitempty"`
	StreamOptions *CuaChatCompletionStreamRequestStreamOptions `json:"streamOptions,omitempty" xml:"streamOptions,omitempty" type:"Struct"`
	// example:
	//
	// []
	Tools []map[string]interface{} `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s CuaChatCompletionStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s CuaChatCompletionStreamRequest) GoString() string {
	return s.String()
}

func (s *CuaChatCompletionStreamRequest) GetMessages() []*CuaChatCompletionStreamRequestMessages {
	return s.Messages
}

func (s *CuaChatCompletionStreamRequest) GetMetadata() *CuaChatCompletionStreamRequestMetadata {
	return s.Metadata
}

func (s *CuaChatCompletionStreamRequest) GetModel() *string {
	return s.Model
}

func (s *CuaChatCompletionStreamRequest) GetStream() *bool {
	return s.Stream
}

func (s *CuaChatCompletionStreamRequest) GetStreamOptions() *CuaChatCompletionStreamRequestStreamOptions {
	return s.StreamOptions
}

func (s *CuaChatCompletionStreamRequest) GetTools() []map[string]interface{} {
	return s.Tools
}

func (s *CuaChatCompletionStreamRequest) SetMessages(v []*CuaChatCompletionStreamRequestMessages) *CuaChatCompletionStreamRequest {
	s.Messages = v
	return s
}

func (s *CuaChatCompletionStreamRequest) SetMetadata(v *CuaChatCompletionStreamRequestMetadata) *CuaChatCompletionStreamRequest {
	s.Metadata = v
	return s
}

func (s *CuaChatCompletionStreamRequest) SetModel(v string) *CuaChatCompletionStreamRequest {
	s.Model = &v
	return s
}

func (s *CuaChatCompletionStreamRequest) SetStream(v bool) *CuaChatCompletionStreamRequest {
	s.Stream = &v
	return s
}

func (s *CuaChatCompletionStreamRequest) SetStreamOptions(v *CuaChatCompletionStreamRequestStreamOptions) *CuaChatCompletionStreamRequest {
	s.StreamOptions = v
	return s
}

func (s *CuaChatCompletionStreamRequest) SetTools(v []map[string]interface{}) *CuaChatCompletionStreamRequest {
	s.Tools = v
	return s
}

func (s *CuaChatCompletionStreamRequest) Validate() error {
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

type CuaChatCompletionStreamRequestMessages struct {
	// This parameter is required.
	Content []*CuaChatCompletionStreamRequestMessagesContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// call_cua_1
	ToolCallId *string `json:"toolCallId,omitempty" xml:"toolCallId,omitempty"`
}

func (s CuaChatCompletionStreamRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s CuaChatCompletionStreamRequestMessages) GoString() string {
	return s.String()
}

func (s *CuaChatCompletionStreamRequestMessages) GetContent() []*CuaChatCompletionStreamRequestMessagesContent {
	return s.Content
}

func (s *CuaChatCompletionStreamRequestMessages) GetRole() *string {
	return s.Role
}

func (s *CuaChatCompletionStreamRequestMessages) GetToolCallId() *string {
	return s.ToolCallId
}

func (s *CuaChatCompletionStreamRequestMessages) SetContent(v []*CuaChatCompletionStreamRequestMessagesContent) *CuaChatCompletionStreamRequestMessages {
	s.Content = v
	return s
}

func (s *CuaChatCompletionStreamRequestMessages) SetRole(v string) *CuaChatCompletionStreamRequestMessages {
	s.Role = &v
	return s
}

func (s *CuaChatCompletionStreamRequestMessages) SetToolCallId(v string) *CuaChatCompletionStreamRequestMessages {
	s.ToolCallId = &v
	return s
}

func (s *CuaChatCompletionStreamRequestMessages) Validate() error {
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

type CuaChatCompletionStreamRequestMessagesContent struct {
	// example:
	//
	// data:image/png;base64,<IMAGE_BASE64>
	ImageData *string                                                `json:"imageData,omitempty" xml:"imageData,omitempty"`
	ImageUrl  *CuaChatCompletionStreamRequestMessagesContentImageUrl `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Struct"`
	// example:
	//
	// 打开浏览器并访问指定页面。
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// image_url
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CuaChatCompletionStreamRequestMessagesContent) String() string {
	return dara.Prettify(s)
}

func (s CuaChatCompletionStreamRequestMessagesContent) GoString() string {
	return s.String()
}

func (s *CuaChatCompletionStreamRequestMessagesContent) GetImageData() *string {
	return s.ImageData
}

func (s *CuaChatCompletionStreamRequestMessagesContent) GetImageUrl() *CuaChatCompletionStreamRequestMessagesContentImageUrl {
	return s.ImageUrl
}

func (s *CuaChatCompletionStreamRequestMessagesContent) GetText() *string {
	return s.Text
}

func (s *CuaChatCompletionStreamRequestMessagesContent) GetType() *string {
	return s.Type
}

func (s *CuaChatCompletionStreamRequestMessagesContent) SetImageData(v string) *CuaChatCompletionStreamRequestMessagesContent {
	s.ImageData = &v
	return s
}

func (s *CuaChatCompletionStreamRequestMessagesContent) SetImageUrl(v *CuaChatCompletionStreamRequestMessagesContentImageUrl) *CuaChatCompletionStreamRequestMessagesContent {
	s.ImageUrl = v
	return s
}

func (s *CuaChatCompletionStreamRequestMessagesContent) SetText(v string) *CuaChatCompletionStreamRequestMessagesContent {
	s.Text = &v
	return s
}

func (s *CuaChatCompletionStreamRequestMessagesContent) SetType(v string) *CuaChatCompletionStreamRequestMessagesContent {
	s.Type = &v
	return s
}

func (s *CuaChatCompletionStreamRequestMessagesContent) Validate() error {
	if s.ImageUrl != nil {
		if err := s.ImageUrl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CuaChatCompletionStreamRequestMessagesContentImageUrl struct {
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/desktop.png
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CuaChatCompletionStreamRequestMessagesContentImageUrl) String() string {
	return dara.Prettify(s)
}

func (s CuaChatCompletionStreamRequestMessagesContentImageUrl) GoString() string {
	return s.String()
}

func (s *CuaChatCompletionStreamRequestMessagesContentImageUrl) GetUrl() *string {
	return s.Url
}

func (s *CuaChatCompletionStreamRequestMessagesContentImageUrl) SetUrl(v string) *CuaChatCompletionStreamRequestMessagesContentImageUrl {
	s.Url = &v
	return s
}

func (s *CuaChatCompletionStreamRequestMessagesContentImageUrl) Validate() error {
	return dara.Validate(s)
}

type CuaChatCompletionStreamRequestMetadata struct {
	// example:
	//
	// 1080
	ScreenHeight *int `json:"screenHeight,omitempty" xml:"screenHeight,omitempty"`
	// example:
	//
	// 1920
	ScreenWidth *int `json:"screenWidth,omitempty" xml:"screenWidth,omitempty"`
}

func (s CuaChatCompletionStreamRequestMetadata) String() string {
	return dara.Prettify(s)
}

func (s CuaChatCompletionStreamRequestMetadata) GoString() string {
	return s.String()
}

func (s *CuaChatCompletionStreamRequestMetadata) GetScreenHeight() *int {
	return s.ScreenHeight
}

func (s *CuaChatCompletionStreamRequestMetadata) GetScreenWidth() *int {
	return s.ScreenWidth
}

func (s *CuaChatCompletionStreamRequestMetadata) SetScreenHeight(v int) *CuaChatCompletionStreamRequestMetadata {
	s.ScreenHeight = &v
	return s
}

func (s *CuaChatCompletionStreamRequestMetadata) SetScreenWidth(v int) *CuaChatCompletionStreamRequestMetadata {
	s.ScreenWidth = &v
	return s
}

func (s *CuaChatCompletionStreamRequestMetadata) Validate() error {
	return dara.Validate(s)
}

type CuaChatCompletionStreamRequestStreamOptions struct {
	// example:
	//
	// true
	IncludeUsage *bool `json:"includeUsage,omitempty" xml:"includeUsage,omitempty"`
}

func (s CuaChatCompletionStreamRequestStreamOptions) String() string {
	return dara.Prettify(s)
}

func (s CuaChatCompletionStreamRequestStreamOptions) GoString() string {
	return s.String()
}

func (s *CuaChatCompletionStreamRequestStreamOptions) GetIncludeUsage() *bool {
	return s.IncludeUsage
}

func (s *CuaChatCompletionStreamRequestStreamOptions) SetIncludeUsage(v bool) *CuaChatCompletionStreamRequestStreamOptions {
	s.IncludeUsage = &v
	return s
}

func (s *CuaChatCompletionStreamRequestStreamOptions) Validate() error {
	return dara.Validate(s)
}
