// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAgentRequest
	GetAppId() *string
	SetChannelId(v string) *GetAgentRequest
	GetChannelId() *string
	SetTaskId(v string) *GetAgentRequest
	GetTaskId() *string
}

type GetAgentRequest struct {
	// example:
	//
	// aec****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 1234
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRequest) GoString() string {
	return s.String()
}

func (s *GetAgentRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetAgentRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *GetAgentRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAgentRequest) SetAppId(v string) *GetAgentRequest {
	s.AppId = &v
	return s
}

func (s *GetAgentRequest) SetChannelId(v string) *GetAgentRequest {
	s.ChannelId = &v
	return s
}

func (s *GetAgentRequest) SetTaskId(v string) *GetAgentRequest {
	s.TaskId = &v
	return s
}

func (s *GetAgentRequest) Validate() error {
	return dara.Validate(s)
}
