// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOpenSourcePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateOpenSourcePermissionRequest
	GetClientToken() *string
	SetConfigure(v string) *UpdateOpenSourcePermissionRequest
	GetConfigure() *string
	SetInstanceId(v string) *UpdateOpenSourcePermissionRequest
	GetInstanceId() *string
	SetRead(v string) *UpdateOpenSourcePermissionRequest
	GetRead() *string
	SetUserName(v string) *UpdateOpenSourcePermissionRequest
	GetUserName() *string
	SetVhost(v string) *UpdateOpenSourcePermissionRequest
	GetVhost() *string
	SetWrite(v string) *UpdateOpenSourcePermissionRequest
	GetWrite() *string
}

type UpdateOpenSourcePermissionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// f6af6bb123988d191bfe01c9a9b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The regular expression for configure permissions.
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
	// The regular expression for read permissions.
	//
	// example:
	//
	// ^$
	Read *string `json:"Read,omitempty" xml:"Read,omitempty"`
	// The username.
	//
	// example:
	//
	// MyName
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The vhost name.
	//
	// example:
	//
	// production
	Vhost *string `json:"Vhost,omitempty" xml:"Vhost,omitempty"`
	// The regular expression for write permissions.
	//
	// example:
	//
	// order_exchange
	Write *string `json:"Write,omitempty" xml:"Write,omitempty"`
}

func (s UpdateOpenSourcePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpenSourcePermissionRequest) GoString() string {
	return s.String()
}

func (s *UpdateOpenSourcePermissionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateOpenSourcePermissionRequest) GetConfigure() *string {
	return s.Configure
}

func (s *UpdateOpenSourcePermissionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateOpenSourcePermissionRequest) GetRead() *string {
	return s.Read
}

func (s *UpdateOpenSourcePermissionRequest) GetUserName() *string {
	return s.UserName
}

func (s *UpdateOpenSourcePermissionRequest) GetVhost() *string {
	return s.Vhost
}

func (s *UpdateOpenSourcePermissionRequest) GetWrite() *string {
	return s.Write
}

func (s *UpdateOpenSourcePermissionRequest) SetClientToken(v string) *UpdateOpenSourcePermissionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateOpenSourcePermissionRequest) SetConfigure(v string) *UpdateOpenSourcePermissionRequest {
	s.Configure = &v
	return s
}

func (s *UpdateOpenSourcePermissionRequest) SetInstanceId(v string) *UpdateOpenSourcePermissionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateOpenSourcePermissionRequest) SetRead(v string) *UpdateOpenSourcePermissionRequest {
	s.Read = &v
	return s
}

func (s *UpdateOpenSourcePermissionRequest) SetUserName(v string) *UpdateOpenSourcePermissionRequest {
	s.UserName = &v
	return s
}

func (s *UpdateOpenSourcePermissionRequest) SetVhost(v string) *UpdateOpenSourcePermissionRequest {
	s.Vhost = &v
	return s
}

func (s *UpdateOpenSourcePermissionRequest) SetWrite(v string) *UpdateOpenSourcePermissionRequest {
	s.Write = &v
	return s
}

func (s *UpdateOpenSourcePermissionRequest) Validate() error {
	return dara.Validate(s)
}
