// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAuthPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffect(v string) *UpdateCustomAuthPermissionRequest
	GetEffect() *string
	SetIdentity(v string) *UpdateCustomAuthPermissionRequest
	GetIdentity() *string
	SetIdentityType(v string) *UpdateCustomAuthPermissionRequest
	GetIdentityType() *string
	SetInstanceId(v string) *UpdateCustomAuthPermissionRequest
	GetInstanceId() *string
	SetPermitAction(v string) *UpdateCustomAuthPermissionRequest
	GetPermitAction() *string
	SetTopic(v string) *UpdateCustomAuthPermissionRequest
	GetTopic() *string
}

type UpdateCustomAuthPermissionRequest struct {
	// Specifies whether to allow or deny access.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALLOW
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// Username or Client ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Identity *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	// The identity type. Valid values:
	//
	// 	- USER
	//
	// 	- CLIENT
	//
	// This parameter is required.
	//
	// example:
	//
	// USER
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// ID of the Cloud Message Queue MQTT version instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-0pp12gl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The permissions that you want to grant.
	//
	// This parameter is required.
	//
	// example:
	//
	// PUB_SUB
	PermitAction *string `json:"PermitAction,omitempty" xml:"PermitAction,omitempty"`
	// Authorized Topic, supporting multi-level MQTT topics and wildcards.
	//
	// This parameter is required.
	//
	// example:
	//
	// test/t1
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s UpdateCustomAuthPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAuthPermissionRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomAuthPermissionRequest) GetEffect() *string {
	return s.Effect
}

func (s *UpdateCustomAuthPermissionRequest) GetIdentity() *string {
	return s.Identity
}

func (s *UpdateCustomAuthPermissionRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *UpdateCustomAuthPermissionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCustomAuthPermissionRequest) GetPermitAction() *string {
	return s.PermitAction
}

func (s *UpdateCustomAuthPermissionRequest) GetTopic() *string {
	return s.Topic
}

func (s *UpdateCustomAuthPermissionRequest) SetEffect(v string) *UpdateCustomAuthPermissionRequest {
	s.Effect = &v
	return s
}

func (s *UpdateCustomAuthPermissionRequest) SetIdentity(v string) *UpdateCustomAuthPermissionRequest {
	s.Identity = &v
	return s
}

func (s *UpdateCustomAuthPermissionRequest) SetIdentityType(v string) *UpdateCustomAuthPermissionRequest {
	s.IdentityType = &v
	return s
}

func (s *UpdateCustomAuthPermissionRequest) SetInstanceId(v string) *UpdateCustomAuthPermissionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCustomAuthPermissionRequest) SetPermitAction(v string) *UpdateCustomAuthPermissionRequest {
	s.PermitAction = &v
	return s
}

func (s *UpdateCustomAuthPermissionRequest) SetTopic(v string) *UpdateCustomAuthPermissionRequest {
	s.Topic = &v
	return s
}

func (s *UpdateCustomAuthPermissionRequest) Validate() error {
	return dara.Validate(s)
}
