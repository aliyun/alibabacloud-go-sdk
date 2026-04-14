// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckReplyToMailAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CheckReplyToMailAddressRequest
	GetLang() *string
	SetMailAddressId(v int32) *CheckReplyToMailAddressRequest
	GetMailAddressId() *int32
	SetOwnerId(v int64) *CheckReplyToMailAddressRequest
	GetOwnerId() *int64
	SetRegion(v string) *CheckReplyToMailAddressRequest
	GetRegion() *string
	SetResourceOwnerAccount(v string) *CheckReplyToMailAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckReplyToMailAddressRequest
	GetResourceOwnerId() *int64
}

type CheckReplyToMailAddressRequest struct {
	// Language.
	//
	// en is English, and any other value or an empty value defaults to Chinese.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Sender Address ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 126545
	MailAddressId *int32 `json:"MailAddressId,omitempty" xml:"MailAddressId,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Region
	//
	// example:
	//
	// cn-hangzhou
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckReplyToMailAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckReplyToMailAddressRequest) GoString() string {
	return s.String()
}

func (s *CheckReplyToMailAddressRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckReplyToMailAddressRequest) GetMailAddressId() *int32 {
	return s.MailAddressId
}

func (s *CheckReplyToMailAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckReplyToMailAddressRequest) GetRegion() *string {
	return s.Region
}

func (s *CheckReplyToMailAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckReplyToMailAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckReplyToMailAddressRequest) SetLang(v string) *CheckReplyToMailAddressRequest {
	s.Lang = &v
	return s
}

func (s *CheckReplyToMailAddressRequest) SetMailAddressId(v int32) *CheckReplyToMailAddressRequest {
	s.MailAddressId = &v
	return s
}

func (s *CheckReplyToMailAddressRequest) SetOwnerId(v int64) *CheckReplyToMailAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckReplyToMailAddressRequest) SetRegion(v string) *CheckReplyToMailAddressRequest {
	s.Region = &v
	return s
}

func (s *CheckReplyToMailAddressRequest) SetResourceOwnerAccount(v string) *CheckReplyToMailAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckReplyToMailAddressRequest) SetResourceOwnerId(v int64) *CheckReplyToMailAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckReplyToMailAddressRequest) Validate() error {
	return dara.Validate(s)
}
