// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextModerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *TextModerationRequest
	GetService() *string
	SetServiceParameters(v string) *TextModerationRequest
	GetServiceParameters() *string
}

type TextModerationRequest struct {
	// The type of the moderation service. Valid values: nickname_detection: user nickname chat_detection: chat interactions comment_detection: dynamic comments pgc_detection: professionally-generated content (PGC) teaching materials
	//
	// Valid values:
	//
	// 	- pgc_detection: moderation of PGC teaching materials
	//
	// 	- nickname_detection: user nickname moderation
	//
	// 	- comment_multilingual_pro: multi-language moderation in international business scenarios
	//
	// 	- chat_detection: moderation of interactive content of private chats
	//
	// 	- ad_compliance_detection: advertising law compliance identification
	//
	// 	- comment_detection: moderation of comment content of public chats
	//
	// 	- ai_art_detection: AI-generated text identfication
	//
	// example:
	//
	// nickname_detection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// example:
	//
	// {"content":"Content to be moderated"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s TextModerationRequest) String() string {
	return dara.Prettify(s)
}

func (s TextModerationRequest) GoString() string {
	return s.String()
}

func (s *TextModerationRequest) GetService() *string {
	return s.Service
}

func (s *TextModerationRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *TextModerationRequest) SetService(v string) *TextModerationRequest {
	s.Service = &v
	return s
}

func (s *TextModerationRequest) SetServiceParameters(v string) *TextModerationRequest {
	s.ServiceParameters = &v
	return s
}

func (s *TextModerationRequest) Validate() error {
	return dara.Validate(s)
}
