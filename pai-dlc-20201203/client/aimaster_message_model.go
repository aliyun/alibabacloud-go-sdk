// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAIMasterMessage interface {
	dara.Model
	String() string
	GoString() string
	SetExtended(v string) *AIMasterMessage
	GetExtended() *string
	SetJobRestartCount(v int32) *AIMasterMessage
	GetJobRestartCount() *int32
	SetMessageContent(v string) *AIMasterMessage
	GetMessageContent() *string
	SetMessageEvent(v string) *AIMasterMessage
	GetMessageEvent() *string
	SetMessageVersion(v int32) *AIMasterMessage
	GetMessageVersion() *int32
	SetRestartType(v string) *AIMasterMessage
	GetRestartType() *string
}

type AIMasterMessage struct {
	Extended        *string `json:"Extended,omitempty" xml:"Extended,omitempty"`
	JobRestartCount *int32  `json:"JobRestartCount,omitempty" xml:"JobRestartCount,omitempty"`
	MessageContent  *string `json:"MessageContent,omitempty" xml:"MessageContent,omitempty"`
	MessageEvent    *string `json:"MessageEvent,omitempty" xml:"MessageEvent,omitempty"`
	MessageVersion  *int32  `json:"MessageVersion,omitempty" xml:"MessageVersion,omitempty"`
	RestartType     *string `json:"RestartType,omitempty" xml:"RestartType,omitempty"`
}

func (s AIMasterMessage) String() string {
	return dara.Prettify(s)
}

func (s AIMasterMessage) GoString() string {
	return s.String()
}

func (s *AIMasterMessage) GetExtended() *string {
	return s.Extended
}

func (s *AIMasterMessage) GetJobRestartCount() *int32 {
	return s.JobRestartCount
}

func (s *AIMasterMessage) GetMessageContent() *string {
	return s.MessageContent
}

func (s *AIMasterMessage) GetMessageEvent() *string {
	return s.MessageEvent
}

func (s *AIMasterMessage) GetMessageVersion() *int32 {
	return s.MessageVersion
}

func (s *AIMasterMessage) GetRestartType() *string {
	return s.RestartType
}

func (s *AIMasterMessage) SetExtended(v string) *AIMasterMessage {
	s.Extended = &v
	return s
}

func (s *AIMasterMessage) SetJobRestartCount(v int32) *AIMasterMessage {
	s.JobRestartCount = &v
	return s
}

func (s *AIMasterMessage) SetMessageContent(v string) *AIMasterMessage {
	s.MessageContent = &v
	return s
}

func (s *AIMasterMessage) SetMessageEvent(v string) *AIMasterMessage {
	s.MessageEvent = &v
	return s
}

func (s *AIMasterMessage) SetMessageVersion(v int32) *AIMasterMessage {
	s.MessageVersion = &v
	return s
}

func (s *AIMasterMessage) SetRestartType(v string) *AIMasterMessage {
	s.RestartType = &v
	return s
}

func (s *AIMasterMessage) Validate() error {
	return dara.Validate(s)
}
