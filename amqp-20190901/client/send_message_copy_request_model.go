// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageCopyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *SendMessageCopyRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *SendMessageCopyRequest
	GetInstanceId() *string
	SetProcessToken(v string) *SendMessageCopyRequest
	GetProcessToken() *string
	SetQueueName(v string) *SendMessageCopyRequest
	GetQueueName() *string
	SetVhostName(v string) *SendMessageCopyRequest
	GetVhostName() *string
}

type SendMessageCopyRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	ProcessToken *string `json:"ProcessToken,omitempty" xml:"ProcessToken,omitempty"`
	// This parameter is required.
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s SendMessageCopyRequest) String() string {
	return dara.Prettify(s)
}

func (s SendMessageCopyRequest) GoString() string {
	return s.String()
}

func (s *SendMessageCopyRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *SendMessageCopyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SendMessageCopyRequest) GetProcessToken() *string {
	return s.ProcessToken
}

func (s *SendMessageCopyRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *SendMessageCopyRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *SendMessageCopyRequest) SetConsoleSessionId(v string) *SendMessageCopyRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *SendMessageCopyRequest) SetInstanceId(v string) *SendMessageCopyRequest {
	s.InstanceId = &v
	return s
}

func (s *SendMessageCopyRequest) SetProcessToken(v string) *SendMessageCopyRequest {
	s.ProcessToken = &v
	return s
}

func (s *SendMessageCopyRequest) SetQueueName(v string) *SendMessageCopyRequest {
	s.QueueName = &v
	return s
}

func (s *SendMessageCopyRequest) SetVhostName(v string) *SendMessageCopyRequest {
	s.VhostName = &v
	return s
}

func (s *SendMessageCopyRequest) Validate() error {
	return dara.Validate(s)
}
