// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMailAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMailAddressId(v int32) *ModifyMailAddressRequest
	GetMailAddressId() *int32
	SetOwnerId(v int64) *ModifyMailAddressRequest
	GetOwnerId() *int64
	SetPassword(v string) *ModifyMailAddressRequest
	GetPassword() *string
	SetReplyAddress(v string) *ModifyMailAddressRequest
	GetReplyAddress() *string
	SetResourceOwnerAccount(v string) *ModifyMailAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyMailAddressRequest
	GetResourceOwnerId() *int64
}

type ModifyMailAddressRequest struct {
	// Sending address ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1344565
	MailAddressId *int32 `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// - Length should be 10 to 20 characters, and must include numbers, uppercase letters, and lowercase letters.
	//
	// - Must contain at least 2 digits, 2 uppercase letters, and 2 lowercase letters, and neither the digits nor the letters can consist of a single character repeated.
	//
	// - Cannot be the same as the last set password.
	//
	// example:
	//
	// DM1mail1234
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Reply address
	//
	// example:
	//
	// a***@example.net
	ReplyAddress         *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyMailAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMailAddressRequest) GoString() string {
	return s.String()
}

func (s *ModifyMailAddressRequest) GetMailAddressId() *int32 {
	return s.MailAddressId
}

func (s *ModifyMailAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyMailAddressRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyMailAddressRequest) GetReplyAddress() *string {
	return s.ReplyAddress
}

func (s *ModifyMailAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyMailAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyMailAddressRequest) SetMailAddressId(v int32) *ModifyMailAddressRequest {
	s.MailAddressId = &v
	return s
}

func (s *ModifyMailAddressRequest) SetOwnerId(v int64) *ModifyMailAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyMailAddressRequest) SetPassword(v string) *ModifyMailAddressRequest {
	s.Password = &v
	return s
}

func (s *ModifyMailAddressRequest) SetReplyAddress(v string) *ModifyMailAddressRequest {
	s.ReplyAddress = &v
	return s
}

func (s *ModifyMailAddressRequest) SetResourceOwnerAccount(v string) *ModifyMailAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyMailAddressRequest) SetResourceOwnerId(v int64) *ModifyMailAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyMailAddressRequest) Validate() error {
	return dara.Validate(s)
}
