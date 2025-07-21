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
	// The name of the vhost that you want to create. Valid values:
	//
	// 	- The name can contain letters, digits, hyphens (-), underscores (_), periods (.), number signs (#), forward slash (/), and at signs (@).
	//
	// 	- The name must be 1 to 255 characters in length.
	//
	// 	- After the vhost is created, you cannot change its name. If you want to change the name of a vhost, delete the vhost and create another vhost.
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
