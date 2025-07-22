// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *NotifyAgentRequest
	GetAppId() *string
	SetChannelId(v string) *NotifyAgentRequest
	GetChannelId() *string
	SetCustomAttribute(v string) *NotifyAgentRequest
	GetCustomAttribute() *string
	SetInterruptable(v bool) *NotifyAgentRequest
	GetInterruptable() *bool
	SetMessage(v string) *NotifyAgentRequest
	GetMessage() *string
	SetPriority(v int32) *NotifyAgentRequest
	GetPriority() *int32
	SetTaskId(v string) *NotifyAgentRequest
	GetTaskId() *string
}

type NotifyAgentRequest struct {
	// example:
	//
	// aec****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// {\\"color\\":\\"blue\\"}
	CustomAttribute *string `json:"CustomAttribute,omitempty" xml:"CustomAttribute,omitempty"`
	// example:
	//
	// true
	Interruptable *bool   `json:"Interruptable,omitempty" xml:"Interruptable,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s NotifyAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s NotifyAgentRequest) GoString() string {
	return s.String()
}

func (s *NotifyAgentRequest) GetAppId() *string {
	return s.AppId
}

func (s *NotifyAgentRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *NotifyAgentRequest) GetCustomAttribute() *string {
	return s.CustomAttribute
}

func (s *NotifyAgentRequest) GetInterruptable() *bool {
	return s.Interruptable
}

func (s *NotifyAgentRequest) GetMessage() *string {
	return s.Message
}

func (s *NotifyAgentRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *NotifyAgentRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *NotifyAgentRequest) SetAppId(v string) *NotifyAgentRequest {
	s.AppId = &v
	return s
}

func (s *NotifyAgentRequest) SetChannelId(v string) *NotifyAgentRequest {
	s.ChannelId = &v
	return s
}

func (s *NotifyAgentRequest) SetCustomAttribute(v string) *NotifyAgentRequest {
	s.CustomAttribute = &v
	return s
}

func (s *NotifyAgentRequest) SetInterruptable(v bool) *NotifyAgentRequest {
	s.Interruptable = &v
	return s
}

func (s *NotifyAgentRequest) SetMessage(v string) *NotifyAgentRequest {
	s.Message = &v
	return s
}

func (s *NotifyAgentRequest) SetPriority(v int32) *NotifyAgentRequest {
	s.Priority = &v
	return s
}

func (s *NotifyAgentRequest) SetTaskId(v string) *NotifyAgentRequest {
	s.TaskId = &v
	return s
}

func (s *NotifyAgentRequest) Validate() error {
	return dara.Validate(s)
}
