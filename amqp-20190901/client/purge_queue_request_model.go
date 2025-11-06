// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurgeQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollina(v string) *PurgeQueueRequest
	GetCollina() *string
	SetConsoleSessionId(v string) *PurgeQueueRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *PurgeQueueRequest
	GetInstanceId() *string
	SetQueueName(v string) *PurgeQueueRequest
	GetQueueName() *string
	SetQueueNames(v map[string]interface{}) *PurgeQueueRequest
	GetQueueNames() map[string]interface{}
	SetUmidToken(v string) *PurgeQueueRequest
	GetUmidToken() *string
	SetVhostName(v string) *PurgeQueueRequest
	GetVhostName() *string
}

type PurgeQueueRequest struct {
	Collina          *string                `json:"Collina,omitempty" xml:"Collina,omitempty"`
	ConsoleSessionId *string                `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	InstanceId       *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueueName        *string                `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	QueueNames       map[string]interface{} `json:"QueueNames,omitempty" xml:"QueueNames,omitempty"`
	UmidToken        *string                `json:"UmidToken,omitempty" xml:"UmidToken,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s PurgeQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s PurgeQueueRequest) GoString() string {
	return s.String()
}

func (s *PurgeQueueRequest) GetCollina() *string {
	return s.Collina
}

func (s *PurgeQueueRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *PurgeQueueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PurgeQueueRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *PurgeQueueRequest) GetQueueNames() map[string]interface{} {
	return s.QueueNames
}

func (s *PurgeQueueRequest) GetUmidToken() *string {
	return s.UmidToken
}

func (s *PurgeQueueRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *PurgeQueueRequest) SetCollina(v string) *PurgeQueueRequest {
	s.Collina = &v
	return s
}

func (s *PurgeQueueRequest) SetConsoleSessionId(v string) *PurgeQueueRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *PurgeQueueRequest) SetInstanceId(v string) *PurgeQueueRequest {
	s.InstanceId = &v
	return s
}

func (s *PurgeQueueRequest) SetQueueName(v string) *PurgeQueueRequest {
	s.QueueName = &v
	return s
}

func (s *PurgeQueueRequest) SetQueueNames(v map[string]interface{}) *PurgeQueueRequest {
	s.QueueNames = v
	return s
}

func (s *PurgeQueueRequest) SetUmidToken(v string) *PurgeQueueRequest {
	s.UmidToken = &v
	return s
}

func (s *PurgeQueueRequest) SetVhostName(v string) *PurgeQueueRequest {
	s.VhostName = &v
	return s
}

func (s *PurgeQueueRequest) Validate() error {
	return dara.Validate(s)
}
