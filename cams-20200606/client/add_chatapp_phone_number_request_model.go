// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddChatappPhoneNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCc(v string) *AddChatappPhoneNumberRequest
	GetCc() *string
	SetCustSpaceId(v string) *AddChatappPhoneNumberRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *AddChatappPhoneNumberRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *AddChatappPhoneNumberRequest
	GetPhoneNumber() *string
	SetPreValidateId(v string) *AddChatappPhoneNumberRequest
	GetPreValidateId() *string
	SetResourceOwnerAccount(v string) *AddChatappPhoneNumberRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddChatappPhoneNumberRequest
	GetResourceOwnerId() *int64
	SetVerifiedName(v string) *AddChatappPhoneNumberRequest
	GetVerifiedName() *string
}

type AddChatappPhoneNumberRequest struct {
	// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86
	Cc *string `json:"Cc,omitempty" xml:"Cc,omitempty"`
	// Adds a phone number for a WhatsApp Business account (WABA).
	//
	// This parameter is required.
	//
	// example:
	//
	// 93928389****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// AddChatappPhoneNumber
	//
	// This parameter is required.
	//
	// example:
	//
	// 1380000****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Deprecated
	//
	// cams:ChatappPhoneNumberRegister
	//
	// example:
	//
	// 1020****
	PreValidateId        *string `json:"PreValidateId,omitempty" xml:"PreValidateId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Private
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba
	VerifiedName *string `json:"VerifiedName,omitempty" xml:"VerifiedName,omitempty"`
}

func (s AddChatappPhoneNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s AddChatappPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *AddChatappPhoneNumberRequest) GetCc() *string {
	return s.Cc
}

func (s *AddChatappPhoneNumberRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *AddChatappPhoneNumberRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddChatappPhoneNumberRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *AddChatappPhoneNumberRequest) GetPreValidateId() *string {
	return s.PreValidateId
}

func (s *AddChatappPhoneNumberRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddChatappPhoneNumberRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddChatappPhoneNumberRequest) GetVerifiedName() *string {
	return s.VerifiedName
}

func (s *AddChatappPhoneNumberRequest) SetCc(v string) *AddChatappPhoneNumberRequest {
	s.Cc = &v
	return s
}

func (s *AddChatappPhoneNumberRequest) SetCustSpaceId(v string) *AddChatappPhoneNumberRequest {
	s.CustSpaceId = &v
	return s
}

func (s *AddChatappPhoneNumberRequest) SetOwnerId(v int64) *AddChatappPhoneNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *AddChatappPhoneNumberRequest) SetPhoneNumber(v string) *AddChatappPhoneNumberRequest {
	s.PhoneNumber = &v
	return s
}

func (s *AddChatappPhoneNumberRequest) SetPreValidateId(v string) *AddChatappPhoneNumberRequest {
	s.PreValidateId = &v
	return s
}

func (s *AddChatappPhoneNumberRequest) SetResourceOwnerAccount(v string) *AddChatappPhoneNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddChatappPhoneNumberRequest) SetResourceOwnerId(v int64) *AddChatappPhoneNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddChatappPhoneNumberRequest) SetVerifiedName(v string) *AddChatappPhoneNumberRequest {
	s.VerifiedName = &v
	return s
}

func (s *AddChatappPhoneNumberRequest) Validate() error {
	return dara.Validate(s)
}
