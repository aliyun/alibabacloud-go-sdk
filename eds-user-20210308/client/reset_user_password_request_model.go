// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetUserPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNotifyType(v int32) *ResetUserPasswordRequest
	GetNotifyType() *int32
	SetUsers(v []*string) *ResetUserPasswordRequest
	GetUsers() []*string
}

type ResetUserPasswordRequest struct {
	// The method to notify the user after the password is reset.
	//
	// > Alibaba Cloud accounts of the international site do not support sending notification through text messages.
	//
	// example:
	//
	// 1
	NotifyType *int32 `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// The names of the convenience users whose passwords you want to reset.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ResetUserPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordRequest) GetNotifyType() *int32 {
	return s.NotifyType
}

func (s *ResetUserPasswordRequest) GetUsers() []*string {
	return s.Users
}

func (s *ResetUserPasswordRequest) SetNotifyType(v int32) *ResetUserPasswordRequest {
	s.NotifyType = &v
	return s
}

func (s *ResetUserPasswordRequest) SetUsers(v []*string) *ResetUserPasswordRequest {
	s.Users = v
	return s
}

func (s *ResetUserPasswordRequest) Validate() error {
	return dara.Validate(s)
}
