// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunLegalAdviceConsultationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RunLegalAdviceConsultationShrinkRequest
	GetAppId() *string
	SetAssistantShrink(v string) *RunLegalAdviceConsultationShrinkRequest
	GetAssistantShrink() *string
	SetExtraShrink(v string) *RunLegalAdviceConsultationShrinkRequest
	GetExtraShrink() *string
	SetStream(v bool) *RunLegalAdviceConsultationShrinkRequest
	GetStream() *bool
	SetThreadShrink(v string) *RunLegalAdviceConsultationShrinkRequest
	GetThreadShrink() *string
}

type RunLegalAdviceConsultationShrinkRequest struct {
	// example:
	//
	// farui
	AppId           *string `json:"appId,omitempty" xml:"appId,omitempty"`
	AssistantShrink *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	ExtraShrink     *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// example:
	//
	// true
	Stream       *bool   `json:"stream,omitempty" xml:"stream,omitempty"`
	ThreadShrink *string `json:"thread,omitempty" xml:"thread,omitempty"`
}

func (s RunLegalAdviceConsultationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunLegalAdviceConsultationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *RunLegalAdviceConsultationShrinkRequest) GetAssistantShrink() *string {
	return s.AssistantShrink
}

func (s *RunLegalAdviceConsultationShrinkRequest) GetExtraShrink() *string {
	return s.ExtraShrink
}

func (s *RunLegalAdviceConsultationShrinkRequest) GetStream() *bool {
	return s.Stream
}

func (s *RunLegalAdviceConsultationShrinkRequest) GetThreadShrink() *string {
	return s.ThreadShrink
}

func (s *RunLegalAdviceConsultationShrinkRequest) SetAppId(v string) *RunLegalAdviceConsultationShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RunLegalAdviceConsultationShrinkRequest) SetAssistantShrink(v string) *RunLegalAdviceConsultationShrinkRequest {
	s.AssistantShrink = &v
	return s
}

func (s *RunLegalAdviceConsultationShrinkRequest) SetExtraShrink(v string) *RunLegalAdviceConsultationShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *RunLegalAdviceConsultationShrinkRequest) SetStream(v bool) *RunLegalAdviceConsultationShrinkRequest {
	s.Stream = &v
	return s
}

func (s *RunLegalAdviceConsultationShrinkRequest) SetThreadShrink(v string) *RunLegalAdviceConsultationShrinkRequest {
	s.ThreadShrink = &v
	return s
}

func (s *RunLegalAdviceConsultationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
