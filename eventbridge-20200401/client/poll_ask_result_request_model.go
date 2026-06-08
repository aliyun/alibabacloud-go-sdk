// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPollAskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *PollAskResultRequest
	GetAgentName() *string
	SetMessageId(v string) *PollAskResultRequest
	GetMessageId() *string
}

type PollAskResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// demo-luma-agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// msg_xxx
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s PollAskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s PollAskResultRequest) GoString() string {
	return s.String()
}

func (s *PollAskResultRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *PollAskResultRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *PollAskResultRequest) SetAgentName(v string) *PollAskResultRequest {
	s.AgentName = &v
	return s
}

func (s *PollAskResultRequest) SetMessageId(v string) *PollAskResultRequest {
	s.MessageId = &v
	return s
}

func (s *PollAskResultRequest) Validate() error {
	return dara.Validate(s)
}
