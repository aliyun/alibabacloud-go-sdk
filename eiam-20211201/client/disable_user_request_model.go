// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DisableUserRequest
	GetInstanceId() *string
	SetUserId(v string) *DisableUserRequest
	GetUserId() *string
}

type DisableUserRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DisableUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableUserRequest) GoString() string {
	return s.String()
}

func (s *DisableUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *DisableUserRequest) SetInstanceId(v string) *DisableUserRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableUserRequest) SetUserId(v string) *DisableUserRequest {
	s.UserId = &v
	return s
}

func (s *DisableUserRequest) Validate() error {
	return dara.Validate(s)
}
