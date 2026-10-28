// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *ModifyUserRequest
	GetBusinessChannel() *string
	SetEmail(v string) *ModifyUserRequest
	GetEmail() *string
	SetEndUserId(v string) *ModifyUserRequest
	GetEndUserId() *string
	SetPhone(v string) *ModifyUserRequest
	GetPhone() *string
}

type ModifyUserRequest struct {
	// The business channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The email address of the convenience account.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The user ID of the convenience account, which corresponds to the username field displayed in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The mobile phone number of the convenience account.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s ModifyUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *ModifyUserRequest) GetEmail() *string {
	return s.Email
}

func (s *ModifyUserRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ModifyUserRequest) GetPhone() *string {
	return s.Phone
}

func (s *ModifyUserRequest) SetBusinessChannel(v string) *ModifyUserRequest {
	s.BusinessChannel = &v
	return s
}

func (s *ModifyUserRequest) SetEmail(v string) *ModifyUserRequest {
	s.Email = &v
	return s
}

func (s *ModifyUserRequest) SetEndUserId(v string) *ModifyUserRequest {
	s.EndUserId = &v
	return s
}

func (s *ModifyUserRequest) SetPhone(v string) *ModifyUserRequest {
	s.Phone = &v
	return s
}

func (s *ModifyUserRequest) Validate() error {
	return dara.Validate(s)
}
