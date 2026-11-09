// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAuthPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentity(v string) *DeleteCustomAuthPermissionRequest
	GetIdentity() *string
	SetIdentityType(v string) *DeleteCustomAuthPermissionRequest
	GetIdentityType() *string
	SetInstanceId(v string) *DeleteCustomAuthPermissionRequest
	GetInstanceId() *string
	SetTopic(v string) *DeleteCustomAuthPermissionRequest
	GetTopic() *string
}

type DeleteCustomAuthPermissionRequest struct {
	// The username or client ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Identity *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	// The identity type. Valid values:
	//
	// - USER: username only.
	//
	// - CLIENT: specific client.
	//
	// This parameter is required.
	//
	// example:
	//
	// USER
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// The ID of the Micro MSMQ MQTT instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-cn-0pp1ldu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The authorized topic. Multi-level MQTT topics and wildcards are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// test/t1
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s DeleteCustomAuthPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAuthPermissionRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomAuthPermissionRequest) GetIdentity() *string {
	return s.Identity
}

func (s *DeleteCustomAuthPermissionRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *DeleteCustomAuthPermissionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCustomAuthPermissionRequest) GetTopic() *string {
	return s.Topic
}

func (s *DeleteCustomAuthPermissionRequest) SetIdentity(v string) *DeleteCustomAuthPermissionRequest {
	s.Identity = &v
	return s
}

func (s *DeleteCustomAuthPermissionRequest) SetIdentityType(v string) *DeleteCustomAuthPermissionRequest {
	s.IdentityType = &v
	return s
}

func (s *DeleteCustomAuthPermissionRequest) SetInstanceId(v string) *DeleteCustomAuthPermissionRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCustomAuthPermissionRequest) SetTopic(v string) *DeleteCustomAuthPermissionRequest {
	s.Topic = &v
	return s
}

func (s *DeleteCustomAuthPermissionRequest) Validate() error {
	return dara.Validate(s)
}
