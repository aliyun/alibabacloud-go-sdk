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
	SetResponseFormatType(v string) *RunCompletionMessageRequest
	GetResponseFormatType() *string
}

type RunCompletionMessageRequest struct {
	// This parameter is required.
	Messages []*RunCompletionMessageRequestMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// example:
	//
	// ccai-14b
	ModelCode *string `json:"ModelCode,omitempty" xml:"ModelCode,omitempty"`
	// example:
	//
	// false
	Stream             *bool   `json:"Stream,omitempty" xml:"Stream,omitempty"`
	ResponseFormatType *string `json:"responseFormatType,omitempty" xml:"responseFormatType,omitempty"`
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

func (s *RunCompletionMessageRequest) GetResponseFormatType() *string {
	return s.ResponseFormatType
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

func (s *RunCompletionMessageRequest) SetResponseFormatType(v string) *RunCompletionMessageRequest {
	s.ResponseFormatType = &v
	return s
}

func (s *RunCompletionMessageRequest) Validate() error {
	return dara.Validate(s)
}

type RunCompletionMessageRequestMessages struct {
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
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
