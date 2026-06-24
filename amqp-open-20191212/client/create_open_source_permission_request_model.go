// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSourcePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateOpenSourcePermissionRequest
	GetClientToken() *string
	SetConfigure(v string) *CreateOpenSourcePermissionRequest
	GetConfigure() *string
	SetInstanceId(v string) *CreateOpenSourcePermissionRequest
	GetInstanceId() *string
	SetRead(v string) *CreateOpenSourcePermissionRequest
	GetRead() *string
	SetUserName(v string) *CreateOpenSourcePermissionRequest
	GetUserName() *string
	SetVhost(v string) *CreateOpenSourcePermissionRequest
	GetVhost() *string
	SetWrite(v string) *CreateOpenSourcePermissionRequest
	GetWrite() *string
}

type CreateOpenSourcePermissionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see How to ensure idempotence.
	//
	// example:
	//
	// f6af6bb123988d191bfe01c9a9b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The regular expression for the configure permission.
	//
	// example:
	//
	// ^$
	Configure *string `json:"Configure,omitempty" xml:"Configure,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9nu***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The regular expression for the read permission.
	//
	// example:
	//
	// ^$
	Read *string `json:"Read,omitempty" xml:"Read,omitempty"`
	// The username.
	//
	// example:
	//
	// myName
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The vhost name.
	//
	// example:
	//
	// vhostName
	Vhost *string `json:"Vhost,omitempty" xml:"Vhost,omitempty"`
	// The regular expression for the write permission.
	//
	// example:
	//
	// order_exchange
	Write *string `json:"Write,omitempty" xml:"Write,omitempty"`
}

func (s CreateOpenSourcePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSourcePermissionRequest) GoString() string {
	return s.String()
}

func (s *CreateOpenSourcePermissionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateOpenSourcePermissionRequest) GetConfigure() *string {
	return s.Configure
}

func (s *CreateOpenSourcePermissionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateOpenSourcePermissionRequest) GetRead() *string {
	return s.Read
}

func (s *CreateOpenSourcePermissionRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateOpenSourcePermissionRequest) GetVhost() *string {
	return s.Vhost
}

func (s *CreateOpenSourcePermissionRequest) GetWrite() *string {
	return s.Write
}

func (s *CreateOpenSourcePermissionRequest) SetClientToken(v string) *CreateOpenSourcePermissionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateOpenSourcePermissionRequest) SetConfigure(v string) *CreateOpenSourcePermissionRequest {
	s.Configure = &v
	return s
}

func (s *CreateOpenSourcePermissionRequest) SetInstanceId(v string) *CreateOpenSourcePermissionRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateOpenSourcePermissionRequest) SetRead(v string) *CreateOpenSourcePermissionRequest {
	s.Read = &v
	return s
}

func (s *CreateOpenSourcePermissionRequest) SetUserName(v string) *CreateOpenSourcePermissionRequest {
	s.UserName = &v
	return s
}

func (s *CreateOpenSourcePermissionRequest) SetVhost(v string) *CreateOpenSourcePermissionRequest {
	s.Vhost = &v
	return s
}

func (s *CreateOpenSourcePermissionRequest) SetWrite(v string) *CreateOpenSourcePermissionRequest {
	s.Write = &v
	return s
}

func (s *CreateOpenSourcePermissionRequest) Validate() error {
	return dara.Validate(s)
}
