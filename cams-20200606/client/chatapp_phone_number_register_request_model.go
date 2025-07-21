// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappPhoneNumberRegisterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *ChatappPhoneNumberRegisterRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *ChatappPhoneNumberRegisterRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *ChatappPhoneNumberRegisterRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *ChatappPhoneNumberRegisterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ChatappPhoneNumberRegisterRequest
	GetResourceOwnerId() *int64
}

type ChatappPhoneNumberRegisterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 939283893939
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8613800000000
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ChatappPhoneNumberRegisterRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatappPhoneNumberRegisterRequest) GoString() string {
	return s.String()
}

func (s *ChatappPhoneNumberRegisterRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ChatappPhoneNumberRegisterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChatappPhoneNumberRegisterRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ChatappPhoneNumberRegisterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ChatappPhoneNumberRegisterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ChatappPhoneNumberRegisterRequest) SetCustSpaceId(v string) *ChatappPhoneNumberRegisterRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ChatappPhoneNumberRegisterRequest) SetOwnerId(v int64) *ChatappPhoneNumberRegisterRequest {
	s.OwnerId = &v
	return s
}

func (s *ChatappPhoneNumberRegisterRequest) SetPhoneNumber(v string) *ChatappPhoneNumberRegisterRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ChatappPhoneNumberRegisterRequest) SetResourceOwnerAccount(v string) *ChatappPhoneNumberRegisterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ChatappPhoneNumberRegisterRequest) SetResourceOwnerId(v int64) *ChatappPhoneNumberRegisterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ChatappPhoneNumberRegisterRequest) Validate() error {
	return dara.Validate(s)
}
