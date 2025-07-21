// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPhoneBusinessProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *QueryPhoneBusinessProfileRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *QueryPhoneBusinessProfileRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *QueryPhoneBusinessProfileRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *QueryPhoneBusinessProfileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryPhoneBusinessProfileRequest
	GetResourceOwnerId() *int64
}

type QueryPhoneBusinessProfileRequest struct {
	// The space ID of the user within the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2934839388494***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryPhoneBusinessProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPhoneBusinessProfileRequest) GoString() string {
	return s.String()
}

func (s *QueryPhoneBusinessProfileRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *QueryPhoneBusinessProfileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryPhoneBusinessProfileRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *QueryPhoneBusinessProfileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryPhoneBusinessProfileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryPhoneBusinessProfileRequest) SetCustSpaceId(v string) *QueryPhoneBusinessProfileRequest {
	s.CustSpaceId = &v
	return s
}

func (s *QueryPhoneBusinessProfileRequest) SetOwnerId(v int64) *QueryPhoneBusinessProfileRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPhoneBusinessProfileRequest) SetPhoneNumber(v string) *QueryPhoneBusinessProfileRequest {
	s.PhoneNumber = &v
	return s
}

func (s *QueryPhoneBusinessProfileRequest) SetResourceOwnerAccount(v string) *QueryPhoneBusinessProfileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPhoneBusinessProfileRequest) SetResourceOwnerId(v int64) *QueryPhoneBusinessProfileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPhoneBusinessProfileRequest) Validate() error {
	return dara.Validate(s)
}
