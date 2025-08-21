// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBeginSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsrMaxEndSilence(v int32) *BeginSessionResponseBody
	GetAsrMaxEndSilence() *int32
	SetInterruptible(v bool) *BeginSessionResponseBody
	GetInterruptible() *bool
	SetRequestId(v string) *BeginSessionResponseBody
	GetRequestId() *string
	SetSilenceReplyTimeout(v int32) *BeginSessionResponseBody
	GetSilenceReplyTimeout() *int32
	SetWelcomeMessage(v string) *BeginSessionResponseBody
	GetWelcomeMessage() *string
}

type BeginSessionResponseBody struct {
	AsrMaxEndSilence *int32 `json:"AsrMaxEndSilence,omitempty" xml:"AsrMaxEndSilence,omitempty"`
	Interruptible    *bool  `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
	// example:
	//
	// 149C7528-C104-1B50-A4F9-0C5907A8AD9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 静默超时时间
	//
	// example:
	//
	// 5
	SilenceReplyTimeout *int32 `json:"SilenceReplyTimeout,omitempty" xml:"SilenceReplyTimeout,omitempty"`
	// example:
	//
	// 智能对话机器人为您服务，请问有什么可以帮您？
	WelcomeMessage *string `json:"WelcomeMessage,omitempty" xml:"WelcomeMessage,omitempty"`
}

func (s BeginSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BeginSessionResponseBody) GoString() string {
	return s.String()
}

func (s *BeginSessionResponseBody) GetAsrMaxEndSilence() *int32 {
	return s.AsrMaxEndSilence
}

func (s *BeginSessionResponseBody) GetInterruptible() *bool {
	return s.Interruptible
}

func (s *BeginSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BeginSessionResponseBody) GetSilenceReplyTimeout() *int32 {
	return s.SilenceReplyTimeout
}

func (s *BeginSessionResponseBody) GetWelcomeMessage() *string {
	return s.WelcomeMessage
}

func (s *BeginSessionResponseBody) SetAsrMaxEndSilence(v int32) *BeginSessionResponseBody {
	s.AsrMaxEndSilence = &v
	return s
}

func (s *BeginSessionResponseBody) SetInterruptible(v bool) *BeginSessionResponseBody {
	s.Interruptible = &v
	return s
}

func (s *BeginSessionResponseBody) SetRequestId(v string) *BeginSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *BeginSessionResponseBody) SetSilenceReplyTimeout(v int32) *BeginSessionResponseBody {
	s.SilenceReplyTimeout = &v
	return s
}

func (s *BeginSessionResponseBody) SetWelcomeMessage(v string) *BeginSessionResponseBody {
	s.WelcomeMessage = &v
	return s
}

func (s *BeginSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
