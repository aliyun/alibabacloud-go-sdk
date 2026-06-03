// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorization(v string) *ChatRequest
	GetAuthorization() *string
	SetExternalUserId(v string) *ChatRequest
	GetExternalUserId() *string
	SetInput(v []*ChatRequestInput) *ChatRequest
	GetInput() []*ChatRequestInput
	SetRoutingKey(v string) *ChatRequest
	GetRoutingKey() *string
	SetSessionId(v string) *ChatRequest
	GetSessionId() *string
	SetSettings(v *ChatRequestSettings) *ChatRequest
	GetSettings() *ChatRequestSettings
	SetStreamOptions(v *ChatRequestStreamOptions) *ChatRequest
	GetStreamOptions() *ChatRequestStreamOptions
	SetTemplateId(v string) *ChatRequest
	GetTemplateId() *string
}

type ChatRequest struct {
	// example:
	//
	// Bearer%20eyJhb****...****k
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
	// example:
	//
	// test-user
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	// example:
	//
	// [{"Role":"user","Content":[{"Type":"text","Text":"你好"}]}]
	Input []*ChatRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	RoutingKey *string `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty"`
	// example:
	//
	// test-session-001
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// {"OutputFileMode": "url"}
	Settings *ChatRequestSettings `json:"Settings,omitempty" xml:"Settings,omitempty" type:"Struct"`
	// example:
	//
	// {"IncludeReasoning": false, "IncludeToolCalls": false}
	StreamOptions *ChatRequestStreamOptions `json:"StreamOptions,omitempty" xml:"StreamOptions,omitempty" type:"Struct"`
	// example:
	//
	// template-abc123
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ChatRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatRequest) GoString() string {
	return s.String()
}

func (s *ChatRequest) GetAuthorization() *string {
	return s.Authorization
}

func (s *ChatRequest) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *ChatRequest) GetInput() []*ChatRequestInput {
	return s.Input
}

func (s *ChatRequest) GetRoutingKey() *string {
	return s.RoutingKey
}

func (s *ChatRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatRequest) GetSettings() *ChatRequestSettings {
	return s.Settings
}

func (s *ChatRequest) GetStreamOptions() *ChatRequestStreamOptions {
	return s.StreamOptions
}

func (s *ChatRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ChatRequest) SetAuthorization(v string) *ChatRequest {
	s.Authorization = &v
	return s
}

func (s *ChatRequest) SetExternalUserId(v string) *ChatRequest {
	s.ExternalUserId = &v
	return s
}

func (s *ChatRequest) SetInput(v []*ChatRequestInput) *ChatRequest {
	s.Input = v
	return s
}

func (s *ChatRequest) SetRoutingKey(v string) *ChatRequest {
	s.RoutingKey = &v
	return s
}

func (s *ChatRequest) SetSessionId(v string) *ChatRequest {
	s.SessionId = &v
	return s
}

func (s *ChatRequest) SetSettings(v *ChatRequestSettings) *ChatRequest {
	s.Settings = v
	return s
}

func (s *ChatRequest) SetStreamOptions(v *ChatRequestStreamOptions) *ChatRequest {
	s.StreamOptions = v
	return s
}

func (s *ChatRequest) SetTemplateId(v string) *ChatRequest {
	s.TemplateId = &v
	return s
}

func (s *ChatRequest) Validate() error {
	if s.Input != nil {
		for _, item := range s.Input {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Settings != nil {
		if err := s.Settings.Validate(); err != nil {
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

type ChatRequestInput struct {
	Content []*ChatRequestInputContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ChatRequestInput) String() string {
	return dara.Prettify(s)
}

func (s ChatRequestInput) GoString() string {
	return s.String()
}

func (s *ChatRequestInput) GetContent() []*ChatRequestInputContent {
	return s.Content
}

func (s *ChatRequestInput) GetRole() *string {
	return s.Role
}

func (s *ChatRequestInput) SetContent(v []*ChatRequestInputContent) *ChatRequestInput {
	s.Content = v
	return s
}

func (s *ChatRequestInput) SetRole(v string) *ChatRequestInput {
	s.Role = &v
	return s
}

func (s *ChatRequestInput) Validate() error {
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

type ChatRequestInputContent struct {
	// example:
	//
	// report.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// /workspace/report.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// https://example.com/img.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// 帮我分析这张图片
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ChatRequestInputContent) String() string {
	return dara.Prettify(s)
}

func (s ChatRequestInputContent) GoString() string {
	return s.String()
}

func (s *ChatRequestInputContent) GetFileName() *string {
	return s.FileName
}

func (s *ChatRequestInputContent) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ChatRequestInputContent) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ChatRequestInputContent) GetText() *string {
	return s.Text
}

func (s *ChatRequestInputContent) GetType() *string {
	return s.Type
}

func (s *ChatRequestInputContent) SetFileName(v string) *ChatRequestInputContent {
	s.FileName = &v
	return s
}

func (s *ChatRequestInputContent) SetFileUrl(v string) *ChatRequestInputContent {
	s.FileUrl = &v
	return s
}

func (s *ChatRequestInputContent) SetImageUrl(v string) *ChatRequestInputContent {
	s.ImageUrl = &v
	return s
}

func (s *ChatRequestInputContent) SetText(v string) *ChatRequestInputContent {
	s.Text = &v
	return s
}

func (s *ChatRequestInputContent) SetType(v string) *ChatRequestInputContent {
	s.Type = &v
	return s
}

func (s *ChatRequestInputContent) Validate() error {
	return dara.Validate(s)
}

type ChatRequestSettings struct {
	// example:
	//
	// base64
	OutputFileMode *string `json:"OutputFileMode,omitempty" xml:"OutputFileMode,omitempty"`
}

func (s ChatRequestSettings) String() string {
	return dara.Prettify(s)
}

func (s ChatRequestSettings) GoString() string {
	return s.String()
}

func (s *ChatRequestSettings) GetOutputFileMode() *string {
	return s.OutputFileMode
}

func (s *ChatRequestSettings) SetOutputFileMode(v string) *ChatRequestSettings {
	s.OutputFileMode = &v
	return s
}

func (s *ChatRequestSettings) Validate() error {
	return dara.Validate(s)
}

type ChatRequestStreamOptions struct {
	// example:
	//
	// true
	IncludeReasoning *bool `json:"IncludeReasoning,omitempty" xml:"IncludeReasoning,omitempty"`
	// example:
	//
	// true
	IncludeToolCalls *bool `json:"IncludeToolCalls,omitempty" xml:"IncludeToolCalls,omitempty"`
}

func (s ChatRequestStreamOptions) String() string {
	return dara.Prettify(s)
}

func (s ChatRequestStreamOptions) GoString() string {
	return s.String()
}

func (s *ChatRequestStreamOptions) GetIncludeReasoning() *bool {
	return s.IncludeReasoning
}

func (s *ChatRequestStreamOptions) GetIncludeToolCalls() *bool {
	return s.IncludeToolCalls
}

func (s *ChatRequestStreamOptions) SetIncludeReasoning(v bool) *ChatRequestStreamOptions {
	s.IncludeReasoning = &v
	return s
}

func (s *ChatRequestStreamOptions) SetIncludeToolCalls(v bool) *ChatRequestStreamOptions {
	s.IncludeToolCalls = &v
	return s
}

func (s *ChatRequestStreamOptions) Validate() error {
	return dara.Validate(s)
}
