// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageAICLoginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionName(v string) *ManageAICLoginRequest
	GetActionName() *string
	SetInstanceId(v string) *ManageAICLoginRequest
	GetInstanceId() *string
	SetKeyGroup(v string) *ManageAICLoginRequest
	GetKeyGroup() *string
	SetKeyName(v string) *ManageAICLoginRequest
	GetKeyName() *string
}

type ManageAICLoginRequest struct {
	// Manage actions
	//
	// Valid value:
	//
	// 	- open
	//
	// 	- close
	//
	// This parameter is required.
	//
	// example:
	//
	// open
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aic-xxx-1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Public Key Grouping
	//
	// example:
	//
	// g-test
	KeyGroup *string `json:"KeyGroup,omitempty" xml:"KeyGroup,omitempty"`
	// Public Key Name
	//
	// example:
	//
	// mykey
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
}

func (s ManageAICLoginRequest) String() string {
	return dara.Prettify(s)
}

func (s ManageAICLoginRequest) GoString() string {
	return s.String()
}

func (s *ManageAICLoginRequest) GetActionName() *string {
	return s.ActionName
}

func (s *ManageAICLoginRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ManageAICLoginRequest) GetKeyGroup() *string {
	return s.KeyGroup
}

func (s *ManageAICLoginRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *ManageAICLoginRequest) SetActionName(v string) *ManageAICLoginRequest {
	s.ActionName = &v
	return s
}

func (s *ManageAICLoginRequest) SetInstanceId(v string) *ManageAICLoginRequest {
	s.InstanceId = &v
	return s
}

func (s *ManageAICLoginRequest) SetKeyGroup(v string) *ManageAICLoginRequest {
	s.KeyGroup = &v
	return s
}

func (s *ManageAICLoginRequest) SetKeyName(v string) *ManageAICLoginRequest {
	s.KeyName = &v
	return s
}

func (s *ManageAICLoginRequest) Validate() error {
	return dara.Validate(s)
}
