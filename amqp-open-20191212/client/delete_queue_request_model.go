// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteQueueRequest
	GetInstanceId() *string
	SetQueueName(v string) *DeleteQueueRequest
	GetQueueName() *string
	SetVirtualHost(v string) *DeleteQueueRequest
	GetVirtualHost() *string
}

type DeleteQueueRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1880770869023***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The queue name.
	//
	// This parameter is required.
	//
	// example:
	//
	// DemoQueue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The vhost name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s DeleteQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQueueRequest) GoString() string {
	return s.String()
}

func (s *DeleteQueueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteQueueRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *DeleteQueueRequest) GetVirtualHost() *string {
	return s.VirtualHost
}

func (s *DeleteQueueRequest) SetInstanceId(v string) *DeleteQueueRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteQueueRequest) SetQueueName(v string) *DeleteQueueRequest {
	s.QueueName = &v
	return s
}

func (s *DeleteQueueRequest) SetVirtualHost(v string) *DeleteQueueRequest {
	s.VirtualHost = &v
	return s
}

func (s *DeleteQueueRequest) Validate() error {
	return dara.Validate(s)
}
