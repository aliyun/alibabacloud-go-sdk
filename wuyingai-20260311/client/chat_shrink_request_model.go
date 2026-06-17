// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorization(v string) *ChatShrinkRequest
	GetAuthorization() *string
	SetExternalUserId(v string) *ChatShrinkRequest
	GetExternalUserId() *string
	SetInputShrink(v string) *ChatShrinkRequest
	GetInputShrink() *string
	SetModel(v string) *ChatShrinkRequest
	GetModel() *string
	SetResume(v bool) *ChatShrinkRequest
	GetResume() *bool
	SetRoutingKey(v string) *ChatShrinkRequest
	GetRoutingKey() *string
	SetSessionId(v string) *ChatShrinkRequest
	GetSessionId() *string
	SetSettingsShrink(v string) *ChatShrinkRequest
	GetSettingsShrink() *string
	SetStreamOptionsShrink(v string) *ChatShrinkRequest
	GetStreamOptionsShrink() *string
	SetTemplateId(v string) *ChatShrinkRequest
	GetTemplateId() *string
}

type ChatShrinkRequest struct {
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
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	Model       *string `json:"Model,omitempty" xml:"Model,omitempty"`
	Resume      *bool   `json:"Resume,omitempty" xml:"Resume,omitempty"`
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
	SettingsShrink *string `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// The streaming output control options. Contains IncludeReasoning (boolean, default true, specifies whether to include the model thinking process) and IncludeToolCalls (boolean, default true, specifies whether to include tool invocation details). If not specified or set to a null object, the behavior is consistent with the legacy version.
	//
	// example:
	//
	// {"IncludeReasoning": false, "IncludeToolCalls": false}
	StreamOptionsShrink *string `json:"StreamOptions,omitempty" xml:"StreamOptions,omitempty"`
	// The agent template ID.
	//
	// example:
	//
	// template-abc123
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ChatShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChatShrinkRequest) GetAuthorization() *string {
	return s.Authorization
}

func (s *ChatShrinkRequest) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *ChatShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *ChatShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *ChatShrinkRequest) GetResume() *bool {
	return s.Resume
}

func (s *ChatShrinkRequest) GetRoutingKey() *string {
	return s.RoutingKey
}

func (s *ChatShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatShrinkRequest) GetSettingsShrink() *string {
	return s.SettingsShrink
}

func (s *ChatShrinkRequest) GetStreamOptionsShrink() *string {
	return s.StreamOptionsShrink
}

func (s *ChatShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ChatShrinkRequest) SetAuthorization(v string) *ChatShrinkRequest {
	s.Authorization = &v
	return s
}

func (s *ChatShrinkRequest) SetExternalUserId(v string) *ChatShrinkRequest {
	s.ExternalUserId = &v
	return s
}

func (s *ChatShrinkRequest) SetInputShrink(v string) *ChatShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *ChatShrinkRequest) SetModel(v string) *ChatShrinkRequest {
	s.Model = &v
	return s
}

func (s *ChatShrinkRequest) SetResume(v bool) *ChatShrinkRequest {
	s.Resume = &v
	return s
}

func (s *ChatShrinkRequest) SetRoutingKey(v string) *ChatShrinkRequest {
	s.RoutingKey = &v
	return s
}

func (s *ChatShrinkRequest) SetSessionId(v string) *ChatShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *ChatShrinkRequest) SetSettingsShrink(v string) *ChatShrinkRequest {
	s.SettingsShrink = &v
	return s
}

func (s *ChatShrinkRequest) SetStreamOptionsShrink(v string) *ChatShrinkRequest {
	s.StreamOptionsShrink = &v
	return s
}

func (s *ChatShrinkRequest) SetTemplateId(v string) *ChatShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *ChatShrinkRequest) Validate() error {
	return dara.Validate(s)
}
