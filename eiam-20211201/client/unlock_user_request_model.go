// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UnlockUserRequest
	GetInstanceId() *string
	SetUserId(v string) *UnlockUserRequest
	GetUserId() *string
}

type UnlockUserRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UnlockUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UnlockUserRequest) GoString() string {
	return s.String()
}

func (s *UnlockUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnlockUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *UnlockUserRequest) SetInstanceId(v string) *UnlockUserRequest {
	s.InstanceId = &v
	return s
}

func (s *UnlockUserRequest) SetUserId(v string) *UnlockUserRequest {
	s.UserId = &v
	return s
}

func (s *UnlockUserRequest) Validate() error {
	return dara.Validate(s)
}
