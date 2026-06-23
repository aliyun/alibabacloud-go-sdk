// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateVirtualHostRequest
	GetInstanceId() *string
	SetVirtualHost(v string) *CreateVirtualHostRequest
	GetVirtualHost() *string
}

type CreateVirtualHostRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9n***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the vhost to create. The name must meet the following requirements:
	//
	// - The name can contain only letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slashes (/), and at signs (@).
	//
	// - The name must be 1 to 255 characters in length.
	//
	// - The vhost name cannot be changed after creation. To use a different name, delete the vhost and create a new one.
	//
	// This parameter is required.
	//
	// example:
	//
	// Demo
	VirtualHost *string `json:"VirtualHost,omitempty" xml:"VirtualHost,omitempty"`
}

func (s CreateVirtualHostRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualHostRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualHostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateVirtualHostRequest) GetVirtualHost() *string {
	return s.VirtualHost
}

func (s *CreateVirtualHostRequest) SetInstanceId(v string) *CreateVirtualHostRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateVirtualHostRequest) SetVirtualHost(v string) *CreateVirtualHostRequest {
	s.VirtualHost = &v
	return s
}

func (s *CreateVirtualHostRequest) Validate() error {
	return dara.Validate(s)
}
