// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappPhoneNumberSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetChatappPhoneNumberSettingRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *GetChatappPhoneNumberSettingRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *GetChatappPhoneNumberSettingRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *GetChatappPhoneNumberSettingRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetChatappPhoneNumberSettingRequest
	GetResourceOwnerId() *int64
}

type GetChatappPhoneNumberSettingRequest struct {
	// example:
	//
	// C29393993****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 86138***
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetChatappPhoneNumberSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatappPhoneNumberSettingRequest) GoString() string {
	return s.String()
}

func (s *GetChatappPhoneNumberSettingRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetChatappPhoneNumberSettingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetChatappPhoneNumberSettingRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetChatappPhoneNumberSettingRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetChatappPhoneNumberSettingRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetChatappPhoneNumberSettingRequest) SetCustSpaceId(v string) *GetChatappPhoneNumberSettingRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetChatappPhoneNumberSettingRequest) SetOwnerId(v int64) *GetChatappPhoneNumberSettingRequest {
	s.OwnerId = &v
	return s
}

func (s *GetChatappPhoneNumberSettingRequest) SetPhoneNumber(v string) *GetChatappPhoneNumberSettingRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetChatappPhoneNumberSettingRequest) SetResourceOwnerAccount(v string) *GetChatappPhoneNumberSettingRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetChatappPhoneNumberSettingRequest) SetResourceOwnerId(v int64) *GetChatappPhoneNumberSettingRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetChatappPhoneNumberSettingRequest) Validate() error {
	return dara.Validate(s)
}
