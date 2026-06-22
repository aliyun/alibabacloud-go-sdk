// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryChatappPhoneNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *QueryChatappPhoneNumbersRequest
	GetCustSpaceId() *string
	SetIsvCode(v string) *QueryChatappPhoneNumbersRequest
	GetIsvCode() *string
	SetOwnerId(v int64) *QueryChatappPhoneNumbersRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryChatappPhoneNumbersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryChatappPhoneNumbersRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *QueryChatappPhoneNumbersRequest
	GetStatus() *string
}

type QueryChatappPhoneNumbersRequest struct {
	// The Space ID of the ISV sub-customer.
	//
	// example:
	//
	// 29348393884****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The verification code of the Independent Software Vendor (ISV). This code is used to verify if the user is authorized by the ISV.
	//
	// example:
	//
	// aksik93kdkkxmwol9****
	IsvCode              *string `json:"IsvCode,omitempty" xml:"IsvCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the phone number.
	//
	// example:
	//
	// VERIFIED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryChatappPhoneNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryChatappPhoneNumbersRequest) GoString() string {
	return s.String()
}

func (s *QueryChatappPhoneNumbersRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *QueryChatappPhoneNumbersRequest) GetIsvCode() *string {
	return s.IsvCode
}

func (s *QueryChatappPhoneNumbersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryChatappPhoneNumbersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryChatappPhoneNumbersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryChatappPhoneNumbersRequest) GetStatus() *string {
	return s.Status
}

func (s *QueryChatappPhoneNumbersRequest) SetCustSpaceId(v string) *QueryChatappPhoneNumbersRequest {
	s.CustSpaceId = &v
	return s
}

func (s *QueryChatappPhoneNumbersRequest) SetIsvCode(v string) *QueryChatappPhoneNumbersRequest {
	s.IsvCode = &v
	return s
}

func (s *QueryChatappPhoneNumbersRequest) SetOwnerId(v int64) *QueryChatappPhoneNumbersRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryChatappPhoneNumbersRequest) SetResourceOwnerAccount(v string) *QueryChatappPhoneNumbersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryChatappPhoneNumbersRequest) SetResourceOwnerId(v int64) *QueryChatappPhoneNumbersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryChatappPhoneNumbersRequest) SetStatus(v string) *QueryChatappPhoneNumbersRequest {
	s.Status = &v
	return s
}

func (s *QueryChatappPhoneNumbersRequest) Validate() error {
	return dara.Validate(s)
}
