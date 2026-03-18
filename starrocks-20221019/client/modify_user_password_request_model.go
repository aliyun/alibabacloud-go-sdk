// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyUserPasswordRequest
	GetInstanceId() *string
	SetPassword(v string) *ModifyUserPasswordRequest
	GetPassword() *string
}

type ModifyUserPasswordRequest struct {
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ***
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ModifyUserPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyUserPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyUserPasswordRequest) SetInstanceId(v string) *ModifyUserPasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUserPasswordRequest) SetPassword(v string) *ModifyUserPasswordRequest {
	s.Password = &v
	return s
}

func (s *ModifyUserPasswordRequest) Validate() error {
	return dara.Validate(s)
}
