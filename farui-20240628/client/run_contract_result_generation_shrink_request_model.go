// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunContractResultGenerationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RunContractResultGenerationShrinkRequest
	GetAppId() *string
	SetAssistantShrink(v string) *RunContractResultGenerationShrinkRequest
	GetAssistantShrink() *string
	SetStream(v bool) *RunContractResultGenerationShrinkRequest
	GetStream() *bool
}

type RunContractResultGenerationShrinkRequest struct {
	// example:
	//
	// farui
	AppId           *string `json:"appId,omitempty" xml:"appId,omitempty"`
	AssistantShrink *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s RunContractResultGenerationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunContractResultGenerationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *RunContractResultGenerationShrinkRequest) GetAssistantShrink() *string {
	return s.AssistantShrink
}

func (s *RunContractResultGenerationShrinkRequest) GetStream() *bool {
	return s.Stream
}

func (s *RunContractResultGenerationShrinkRequest) SetAppId(v string) *RunContractResultGenerationShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RunContractResultGenerationShrinkRequest) SetAssistantShrink(v string) *RunContractResultGenerationShrinkRequest {
	s.AssistantShrink = &v
	return s
}

func (s *RunContractResultGenerationShrinkRequest) SetStream(v bool) *RunContractResultGenerationShrinkRequest {
	s.Stream = &v
	return s
}

func (s *RunContractResultGenerationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
