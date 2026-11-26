// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCompletionMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessages(v []*RunCompletionMessageRequestMessages) *RunCompletionMessageRequest
	GetMessages() []*RunCompletionMessageRequestMessages
	SetModelCode(v string) *RunCompletionMessageRequest
	GetModelCode() *string
	SetStream(v bool) *RunCompletionMessageRequest
	GetStream() *bool
}

type RunCompletionMessageRequest struct {
	// The list of messages in the model request.
	Messages []*RunCompletionMessageRequestMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The model specification. Valid values: TYXM_PLUS and TYXM_TURBO.
	//
	// example:
	//
	// TYXM_PLUS
	ModelCode *string `json:"ModelCode,omitempty" xml:"ModelCode,omitempty"`
	// Specifies whether to enable Server-Sent Events (SSE) responses. Default value: false.
	//
	// example:
	//
	// true
	Stream *bool `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s RunCompletionMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionMessageRequest) GoString() string {
	return s.String()
}

func (s *RunCompletionMessageRequest) GetMessages() []*RunCompletionMessageRequestMessages {
	return s.Messages
}

func (s *RunCompletionMessageRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *RunCompletionMessageRequest) GetStream() *bool {
	return s.Stream
}

func (s *RunCompletionMessageRequest) SetMessages(v []*RunCompletionMessageRequestMessages) *RunCompletionMessageRequest {
	s.Messages = v
	return s
}

func (s *RunCompletionMessageRequest) SetModelCode(v string) *RunCompletionMessageRequest {
	s.ModelCode = &v
	return s
}

func (s *RunCompletionMessageRequest) SetStream(v bool) *RunCompletionMessageRequest {
	s.Stream = &v
	return s
}

func (s *RunCompletionMessageRequest) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunCompletionMessageRequestMessages struct {
	// The prompt content.
	//
	// example:
	//
	// Summarize the conversation in 50 words
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The conversation role. Valid values: user (customer), assistant (customer service representative), system (system message), and function (function).
	//
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s RunCompletionMessageRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionMessageRequestMessages) GoString() string {
	return s.String()
}

func (s *RunCompletionMessageRequestMessages) GetContent() *string {
	return s.Content
}

func (s *RunCompletionMessageRequestMessages) GetRole() *string {
	return s.Role
}

func (s *RunCompletionMessageRequestMessages) SetContent(v string) *RunCompletionMessageRequestMessages {
	s.Content = &v
	return s
}

func (s *RunCompletionMessageRequestMessages) SetRole(v string) *RunCompletionMessageRequestMessages {
	s.Role = &v
	return s
}

func (s *RunCompletionMessageRequestMessages) Validate() error {
	return dara.Validate(s)
}
