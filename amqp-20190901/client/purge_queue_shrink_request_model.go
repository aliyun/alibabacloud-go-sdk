// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurgeQueueShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollina(v string) *PurgeQueueShrinkRequest
	GetCollina() *string
	SetConsoleSessionId(v string) *PurgeQueueShrinkRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *PurgeQueueShrinkRequest
	GetInstanceId() *string
	SetQueueName(v string) *PurgeQueueShrinkRequest
	GetQueueName() *string
	SetQueueNamesShrink(v string) *PurgeQueueShrinkRequest
	GetQueueNamesShrink() *string
	SetUmidToken(v string) *PurgeQueueShrinkRequest
	GetUmidToken() *string
	SetVhostName(v string) *PurgeQueueShrinkRequest
	GetVhostName() *string
}

type PurgeQueueShrinkRequest struct {
	Collina          *string `json:"Collina,omitempty" xml:"Collina,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueueName        *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	QueueNamesShrink *string `json:"QueueNames,omitempty" xml:"QueueNames,omitempty"`
	UmidToken        *string `json:"UmidToken,omitempty" xml:"UmidToken,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s PurgeQueueShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PurgeQueueShrinkRequest) GoString() string {
	return s.String()
}

func (s *PurgeQueueShrinkRequest) GetCollina() *string {
	return s.Collina
}

func (s *PurgeQueueShrinkRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *PurgeQueueShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PurgeQueueShrinkRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *PurgeQueueShrinkRequest) GetQueueNamesShrink() *string {
	return s.QueueNamesShrink
}

func (s *PurgeQueueShrinkRequest) GetUmidToken() *string {
	return s.UmidToken
}

func (s *PurgeQueueShrinkRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *PurgeQueueShrinkRequest) SetCollina(v string) *PurgeQueueShrinkRequest {
	s.Collina = &v
	return s
}

func (s *PurgeQueueShrinkRequest) SetConsoleSessionId(v string) *PurgeQueueShrinkRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *PurgeQueueShrinkRequest) SetInstanceId(v string) *PurgeQueueShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *PurgeQueueShrinkRequest) SetQueueName(v string) *PurgeQueueShrinkRequest {
	s.QueueName = &v
	return s
}

func (s *PurgeQueueShrinkRequest) SetQueueNamesShrink(v string) *PurgeQueueShrinkRequest {
	s.QueueNamesShrink = &v
	return s
}

func (s *PurgeQueueShrinkRequest) SetUmidToken(v string) *PurgeQueueShrinkRequest {
	s.UmidToken = &v
	return s
}

func (s *PurgeQueueShrinkRequest) SetVhostName(v string) *PurgeQueueShrinkRequest {
	s.VhostName = &v
	return s
}

func (s *PurgeQueueShrinkRequest) Validate() error {
	return dara.Validate(s)
}
