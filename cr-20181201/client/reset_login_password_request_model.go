// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetLoginPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ResetLoginPasswordRequest
	GetInstanceId() *string
	SetPassword(v string) *ResetLoginPasswordRequest
	GetPassword() *string
}

type ResetLoginPasswordRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new password that you specify to log on to the instance. The password must be 8 to 32 bits in length, and must contain at least two of the following character types: letters, special characters, and digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ResetLoginPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetLoginPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetLoginPasswordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetLoginPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *ResetLoginPasswordRequest) SetInstanceId(v string) *ResetLoginPasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetLoginPasswordRequest) SetPassword(v string) *ResetLoginPasswordRequest {
	s.Password = &v
	return s
}

func (s *ResetLoginPasswordRequest) Validate() error {
	return dara.Validate(s)
}
