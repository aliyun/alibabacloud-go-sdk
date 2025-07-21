// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteVirtualHostRequest
	GetInstanceId() *string
	SetVirtualHost(v string) *DeleteVirtualHostRequest
	GetVirtualHost() *string
}

type DeleteVirtualHostRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance to which the vhost you want to delete belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9nu***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the vhost that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s DeleteVirtualHostRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualHostRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualHostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteVirtualHostRequest) GetVirtualHost() *string {
	return s.VirtualHost
}

func (s *DeleteVirtualHostRequest) SetInstanceId(v string) *DeleteVirtualHostRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteVirtualHostRequest) SetVirtualHost(v string) *DeleteVirtualHostRequest {
	s.VirtualHost = &v
	return s
}

func (s *DeleteVirtualHostRequest) Validate() error {
	return dara.Validate(s)
}
