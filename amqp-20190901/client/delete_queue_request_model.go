// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollina(v string) *DeleteQueueRequest
	GetCollina() *string
	SetConsoleSessionId(v string) *DeleteQueueRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *DeleteQueueRequest
	GetInstanceId() *string
	SetQueueName(v string) *DeleteQueueRequest
	GetQueueName() *string
	SetQueueNames(v map[string]interface{}) *DeleteQueueRequest
	GetQueueNames() map[string]interface{}
	SetUmidToken(v string) *DeleteQueueRequest
	GetUmidToken() *string
	SetVhostName(v string) *DeleteQueueRequest
	GetVhostName() *string
}

type DeleteQueueRequest struct {
	Collina          *string                `json:"Collina,omitempty" xml:"Collina,omitempty"`
	ConsoleSessionId *string                `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	InstanceId       *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueueName        *string                `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	QueueNames       map[string]interface{} `json:"QueueNames,omitempty" xml:"QueueNames,omitempty"`
	UmidToken        *string                `json:"UmidToken,omitempty" xml:"UmidToken,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s DeleteQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQueueRequest) GoString() string {
	return s.String()
}

func (s *DeleteQueueRequest) GetCollina() *string {
	return s.Collina
}

func (s *DeleteQueueRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *DeleteQueueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteQueueRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *DeleteQueueRequest) GetQueueNames() map[string]interface{} {
	return s.QueueNames
}

func (s *DeleteQueueRequest) GetUmidToken() *string {
	return s.UmidToken
}

func (s *DeleteQueueRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *DeleteQueueRequest) SetCollina(v string) *DeleteQueueRequest {
	s.Collina = &v
	return s
}

func (s *DeleteQueueRequest) SetConsoleSessionId(v string) *DeleteQueueRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *DeleteQueueRequest) SetInstanceId(v string) *DeleteQueueRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteQueueRequest) SetQueueName(v string) *DeleteQueueRequest {
	s.QueueName = &v
	return s
}

func (s *DeleteQueueRequest) SetQueueNames(v map[string]interface{}) *DeleteQueueRequest {
	s.QueueNames = v
	return s
}

func (s *DeleteQueueRequest) SetUmidToken(v string) *DeleteQueueRequest {
	s.UmidToken = &v
	return s
}

func (s *DeleteQueueRequest) SetVhostName(v string) *DeleteQueueRequest {
	s.VhostName = &v
	return s
}

func (s *DeleteQueueRequest) Validate() error {
	return dara.Validate(s)
}
