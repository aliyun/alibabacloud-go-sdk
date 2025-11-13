// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunContractRuleGenerationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RunContractRuleGenerationShrinkRequest
	GetAppId() *string
	SetAssistantShrink(v string) *RunContractRuleGenerationShrinkRequest
	GetAssistantShrink() *string
	SetStream(v bool) *RunContractRuleGenerationShrinkRequest
	GetStream() *bool
}

type RunContractRuleGenerationShrinkRequest struct {
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

func (s RunContractRuleGenerationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunContractRuleGenerationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunContractRuleGenerationShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *RunContractRuleGenerationShrinkRequest) GetAssistantShrink() *string {
	return s.AssistantShrink
}

func (s *RunContractRuleGenerationShrinkRequest) GetStream() *bool {
	return s.Stream
}

func (s *RunContractRuleGenerationShrinkRequest) SetAppId(v string) *RunContractRuleGenerationShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RunContractRuleGenerationShrinkRequest) SetAssistantShrink(v string) *RunContractRuleGenerationShrinkRequest {
	s.AssistantShrink = &v
	return s
}

func (s *RunContractRuleGenerationShrinkRequest) SetStream(v bool) *RunContractRuleGenerationShrinkRequest {
	s.Stream = &v
	return s
}

func (s *RunContractRuleGenerationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
