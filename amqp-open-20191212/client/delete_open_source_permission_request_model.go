// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpenSourcePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteOpenSourcePermissionRequest
	GetInstanceId() *string
	SetUserName(v string) *DeleteOpenSourcePermissionRequest
	GetUserName() *string
	SetVhost(v string) *DeleteOpenSourcePermissionRequest
	GetVhost() *string
}

type DeleteOpenSourcePermissionRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9nu***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// myName
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The vhost name.
	//
	// This parameter is required.
	//
	// example:
	//
	// production
	Vhost *string `json:"Vhost,omitempty" xml:"Vhost,omitempty"`
}

func (s DeleteOpenSourcePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSourcePermissionRequest) GoString() string {
	return s.String()
}

func (s *DeleteOpenSourcePermissionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteOpenSourcePermissionRequest) GetUserName() *string {
	return s.UserName
}

func (s *DeleteOpenSourcePermissionRequest) GetVhost() *string {
	return s.Vhost
}

func (s *DeleteOpenSourcePermissionRequest) SetInstanceId(v string) *DeleteOpenSourcePermissionRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteOpenSourcePermissionRequest) SetUserName(v string) *DeleteOpenSourcePermissionRequest {
	s.UserName = &v
	return s
}

func (s *DeleteOpenSourcePermissionRequest) SetVhost(v string) *DeleteOpenSourcePermissionRequest {
	s.Vhost = &v
	return s
}

func (s *DeleteOpenSourcePermissionRequest) Validate() error {
	return dara.Validate(s)
}
