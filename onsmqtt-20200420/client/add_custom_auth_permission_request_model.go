// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomAuthPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffect(v string) *AddCustomAuthPermissionRequest
	GetEffect() *string
	SetIdentity(v string) *AddCustomAuthPermissionRequest
	GetIdentity() *string
	SetIdentityType(v string) *AddCustomAuthPermissionRequest
	GetIdentityType() *string
	SetInstanceId(v string) *AddCustomAuthPermissionRequest
	GetInstanceId() *string
	SetPermitAction(v string) *AddCustomAuthPermissionRequest
	GetPermitAction() *string
	SetTopic(v string) *AddCustomAuthPermissionRequest
	GetTopic() *string
}

type AddCustomAuthPermissionRequest struct {
	// The authorization effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALLOW
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
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
	// The ID of the MSMQ for MQTT instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-cn-0pp12gl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The authorized action.
	//
	// This parameter is required.
	//
	// example:
	//
	// PUB_SUB
	PermitAction *string `json:"PermitAction,omitempty" xml:"PermitAction,omitempty"`
	// The authorized topic. Multi-level MQTT topics and wildcards are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// test/t1
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s AddCustomAuthPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCustomAuthPermissionRequest) GoString() string {
	return s.String()
}

func (s *AddCustomAuthPermissionRequest) GetEffect() *string {
	return s.Effect
}

func (s *AddCustomAuthPermissionRequest) GetIdentity() *string {
	return s.Identity
}

func (s *AddCustomAuthPermissionRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *AddCustomAuthPermissionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddCustomAuthPermissionRequest) GetPermitAction() *string {
	return s.PermitAction
}

func (s *AddCustomAuthPermissionRequest) GetTopic() *string {
	return s.Topic
}

func (s *AddCustomAuthPermissionRequest) SetEffect(v string) *AddCustomAuthPermissionRequest {
	s.Effect = &v
	return s
}

func (s *AddCustomAuthPermissionRequest) SetIdentity(v string) *AddCustomAuthPermissionRequest {
	s.Identity = &v
	return s
}

func (s *AddCustomAuthPermissionRequest) SetIdentityType(v string) *AddCustomAuthPermissionRequest {
	s.IdentityType = &v
	return s
}

func (s *AddCustomAuthPermissionRequest) SetInstanceId(v string) *AddCustomAuthPermissionRequest {
	s.InstanceId = &v
	return s
}

func (s *AddCustomAuthPermissionRequest) SetPermitAction(v string) *AddCustomAuthPermissionRequest {
	s.PermitAction = &v
	return s
}

func (s *AddCustomAuthPermissionRequest) SetTopic(v string) *AddCustomAuthPermissionRequest {
	s.Topic = &v
	return s
}

func (s *AddCustomAuthPermissionRequest) Validate() error {
	return dara.Validate(s)
}
