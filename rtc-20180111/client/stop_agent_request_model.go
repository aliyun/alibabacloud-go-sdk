// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StopAgentRequest
	GetAppId() *string
	SetChannelId(v string) *StopAgentRequest
	GetChannelId() *string
	SetTaskId(v string) *StopAgentRequest
	GetTaskId() *string
}

type StopAgentRequest struct {
	// The application ID. Only a single ID is supported. You can create and query application IDs in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// aec****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The joined channel. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The task ID. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAgentRequest) GoString() string {
	return s.String()
}

func (s *StopAgentRequest) GetAppId() *string {
	return s.AppId
}

func (s *StopAgentRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StopAgentRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopAgentRequest) SetAppId(v string) *StopAgentRequest {
	s.AppId = &v
	return s
}

func (s *StopAgentRequest) SetChannelId(v string) *StopAgentRequest {
	s.ChannelId = &v
	return s
}

func (s *StopAgentRequest) SetTaskId(v string) *StopAgentRequest {
	s.TaskId = &v
	return s
}

func (s *StopAgentRequest) Validate() error {
	return dara.Validate(s)
}
