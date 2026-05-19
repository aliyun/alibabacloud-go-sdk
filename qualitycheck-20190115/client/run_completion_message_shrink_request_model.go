// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCompletionMessageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessagesShrink(v string) *RunCompletionMessageShrinkRequest
	GetMessagesShrink() *string
	SetModelCode(v string) *RunCompletionMessageShrinkRequest
	GetModelCode() *string
	SetStream(v bool) *RunCompletionMessageShrinkRequest
	GetStream() *bool
}

type RunCompletionMessageShrinkRequest struct {
	MessagesShrink *string `json:"Messages,omitempty" xml:"Messages,omitempty"`
	// example:
	//
	// TYXM_PLUS
	ModelCode *string `json:"ModelCode,omitempty" xml:"ModelCode,omitempty"`
	// example:
	//
	// true
	Stream *bool `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s RunCompletionMessageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunCompletionMessageShrinkRequest) GetMessagesShrink() *string {
	return s.MessagesShrink
}

func (s *RunCompletionMessageShrinkRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *RunCompletionMessageShrinkRequest) GetStream() *bool {
	return s.Stream
}

func (s *RunCompletionMessageShrinkRequest) SetMessagesShrink(v string) *RunCompletionMessageShrinkRequest {
	s.MessagesShrink = &v
	return s
}

func (s *RunCompletionMessageShrinkRequest) SetModelCode(v string) *RunCompletionMessageShrinkRequest {
	s.ModelCode = &v
	return s
}

func (s *RunCompletionMessageShrinkRequest) SetStream(v bool) *RunCompletionMessageShrinkRequest {
	s.Stream = &v
	return s
}

func (s *RunCompletionMessageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
