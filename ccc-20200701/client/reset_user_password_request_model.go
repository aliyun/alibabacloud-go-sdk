// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetUserPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ResetUserPasswordRequest
	GetInstanceId() *string
	SetPassword(v string) *ResetUserPasswordRequest
	GetPassword() *string
	SetUserId(v string) *ResetUserPasswordRequest
	GetUserId() *string
}

type ResetUserPasswordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ResetUserPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetUserPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *ResetUserPasswordRequest) GetUserId() *string {
	return s.UserId
}

func (s *ResetUserPasswordRequest) SetInstanceId(v string) *ResetUserPasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetUserPasswordRequest) SetPassword(v string) *ResetUserPasswordRequest {
	s.Password = &v
	return s
}

func (s *ResetUserPasswordRequest) SetUserId(v string) *ResetUserPasswordRequest {
	s.UserId = &v
	return s
}

func (s *ResetUserPasswordRequest) Validate() error {
	return dara.Validate(s)
}
