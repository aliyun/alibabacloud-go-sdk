// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceRabbitMQParameters interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SourceRabbitMQParameters
	GetInstanceId() *string
	SetQueueName(v string) *SourceRabbitMQParameters
	GetQueueName() *string
	SetRegionId(v string) *SourceRabbitMQParameters
	GetRegionId() *string
	SetVirtualHostName(v string) *SourceRabbitMQParameters
	GetVirtualHostName() *string
}

type SourceRabbitMQParameters struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// amqp-cn-nif22u74****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The queue name of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// demo
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The region in which the ApsaraMQ for RabbitMQ instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vhost name of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// eb-connect
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
}

func (s SourceRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s SourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *SourceRabbitMQParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SourceRabbitMQParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *SourceRabbitMQParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *SourceRabbitMQParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *SourceRabbitMQParameters) SetInstanceId(v string) *SourceRabbitMQParameters {
	s.InstanceId = &v
	return s
}

func (s *SourceRabbitMQParameters) SetQueueName(v string) *SourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *SourceRabbitMQParameters) SetRegionId(v string) *SourceRabbitMQParameters {
	s.RegionId = &v
	return s
}

func (s *SourceRabbitMQParameters) SetVirtualHostName(v string) *SourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

func (s *SourceRabbitMQParameters) Validate() error {
	return dara.Validate(s)
}
