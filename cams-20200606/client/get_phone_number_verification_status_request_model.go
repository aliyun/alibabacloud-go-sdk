// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneNumberVerificationStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetPhoneNumberVerificationStatusRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *GetPhoneNumberVerificationStatusRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *GetPhoneNumberVerificationStatusRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *GetPhoneNumberVerificationStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetPhoneNumberVerificationStatusRequest
	GetResourceOwnerId() *int64
}

type GetPhoneNumberVerificationStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetPhoneNumberVerificationStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberVerificationStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberVerificationStatusRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetPhoneNumberVerificationStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetPhoneNumberVerificationStatusRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetPhoneNumberVerificationStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetPhoneNumberVerificationStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetPhoneNumberVerificationStatusRequest) SetCustSpaceId(v string) *GetPhoneNumberVerificationStatusRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusRequest) SetOwnerId(v int64) *GetPhoneNumberVerificationStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusRequest) SetPhoneNumber(v string) *GetPhoneNumberVerificationStatusRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusRequest) SetResourceOwnerAccount(v string) *GetPhoneNumberVerificationStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusRequest) SetResourceOwnerId(v int64) *GetPhoneNumberVerificationStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusRequest) Validate() error {
	return dara.Validate(s)
}
