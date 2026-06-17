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
	SetModel(v string) *ChatRequest
	GetModel() *string
	SetResume(v bool) *ChatRequest
	GetResume() *bool
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
	// Bearer + JWT returned by GetAccessToken. URL-encode the entire string and pass it as a query parameter.
	//
	// example:
	//
	// Bearer%20eyJhb****...****k
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
	// The user ID from the external system.
	//
	// example:
	//
	// test-user
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
	// The message list (JSON string), sorted in chronological order.
	//
	// example:
	//
	// [{"Role":"user","Content":[{"Type":"text","Text":"你好"}]}]
	Input  []*ChatRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Repeated"`
	Model  *string             `json:"Model,omitempty" xml:"Model,omitempty"`
	Resume *bool               `json:"Resume,omitempty" xml:"Resume,omitempty"`
	// The routing key that specifies the backend instance to process the request.
	//
	// example:
	//
	// ""
	RoutingKey *string `json:"RoutingKey,omitempty" xml:"RoutingKey,omitempty"`
	// The session ID for multi-turn conversation context persistence.
	//
	// example:
	//
	// test-session-001
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The additional settings. Contains the output file mode control parameter OutputFileMode (string, valid values: url or base64. Defaults to base64 for legacy compatibility. We recommend url).
	//
	// example:
	//
	// {"OutputFileMode": "url"}
	Settings *ChatRequestSettings `json:"Settings,omitempty" xml:"Settings,omitempty" type:"Struct"`
	// The streaming output control options. Contains IncludeReasoning (boolean, default true, specifies whether to include the model thinking process) and IncludeToolCalls (boolean, default true, specifies whether to include tool invocation details). If not specified or set to a null object, the behavior is consistent with the legacy version.
	//
	// example:
	//
	// {"IncludeReasoning": false, "IncludeToolCalls": false}
	StreamOptions *ChatRequestStreamOptions `json:"StreamOptions,omitempty" xml:"StreamOptions,omitempty" type:"Struct"`
	// The agent template ID.
	//
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

func (s *ChatRequest) GetModel() *string {
	return s.Model
}

func (s *ChatRequest) GetResume() *bool {
	return s.Resume
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

func (s *ChatRequest) SetModel(v string) *ChatRequest {
	s.Model = &v
	return s
}

func (s *ChatRequest) SetResume(v bool) *ChatRequest {
	s.Resume = &v
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
	// The content block list.
	Content []*ChatRequestInputContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The message role.
	//
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
	// The file path or URL (Type=file).
	//
	// example:
	//
	// /workspace/report.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The image URL or Base64-encoded string (Type=image).
	//
	// example:
	//
	// https://example.com/img.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The text content (Type=text).
	//
	// example:
	//
	// 帮我分析这张图片
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The content type.
	//
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
	// Controls the file output mode. Valid values: url or base64. If this parameter is not specified, base64 is used by default for legacy compatibility.
	//
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
	// Specifies whether to include the model thinking process. When set to false, the SSE stream does not include messages with Type="reasoning" or their content events.
	//
	// example:
	//
	// true
	IncludeReasoning *bool `json:"IncludeReasoning,omitempty" xml:"IncludeReasoning,omitempty"`
	// Specifies whether to include tool invocation details. When set to false, the SSE stream does not include messages of type plugin_call, plugin_call_output, mcp_call, or mcp_call_output, or their content events.
	//
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
