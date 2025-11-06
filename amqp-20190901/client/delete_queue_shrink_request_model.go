// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQueueShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollina(v string) *DeleteQueueShrinkRequest
	GetCollina() *string
	SetConsoleSessionId(v string) *DeleteQueueShrinkRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *DeleteQueueShrinkRequest
	GetInstanceId() *string
	SetQueueName(v string) *DeleteQueueShrinkRequest
	GetQueueName() *string
	SetQueueNamesShrink(v string) *DeleteQueueShrinkRequest
	GetQueueNamesShrink() *string
	SetUmidToken(v string) *DeleteQueueShrinkRequest
	GetUmidToken() *string
	SetVhostName(v string) *DeleteQueueShrinkRequest
	GetVhostName() *string
}

type DeleteQueueShrinkRequest struct {
	Collina          *string `json:"Collina,omitempty" xml:"Collina,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueueName        *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	QueueNamesShrink *string `json:"QueueNames,omitempty" xml:"QueueNames,omitempty"`
	UmidToken        *string `json:"UmidToken,omitempty" xml:"UmidToken,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s DeleteQueueShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQueueShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteQueueShrinkRequest) GetCollina() *string {
	return s.Collina
}

func (s *DeleteQueueShrinkRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *DeleteQueueShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteQueueShrinkRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *DeleteQueueShrinkRequest) GetQueueNamesShrink() *string {
	return s.QueueNamesShrink
}

func (s *DeleteQueueShrinkRequest) GetUmidToken() *string {
	return s.UmidToken
}

func (s *DeleteQueueShrinkRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *DeleteQueueShrinkRequest) SetCollina(v string) *DeleteQueueShrinkRequest {
	s.Collina = &v
	return s
}

func (s *DeleteQueueShrinkRequest) SetConsoleSessionId(v string) *DeleteQueueShrinkRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *DeleteQueueShrinkRequest) SetInstanceId(v string) *DeleteQueueShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteQueueShrinkRequest) SetQueueName(v string) *DeleteQueueShrinkRequest {
	s.QueueName = &v
	return s
}

func (s *DeleteQueueShrinkRequest) SetQueueNamesShrink(v string) *DeleteQueueShrinkRequest {
	s.QueueNamesShrink = &v
	return s
}

func (s *DeleteQueueShrinkRequest) SetUmidToken(v string) *DeleteQueueShrinkRequest {
	s.UmidToken = &v
	return s
}

func (s *DeleteQueueShrinkRequest) SetVhostName(v string) *DeleteQueueShrinkRequest {
	s.VhostName = &v
	return s
}

func (s *DeleteQueueShrinkRequest) Validate() error {
	return dara.Validate(s)
}
