// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyPhoneWithTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *VerifyPhoneWithTokenRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *VerifyPhoneWithTokenRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *VerifyPhoneWithTokenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *VerifyPhoneWithTokenRequest
	GetResourceOwnerId() *int64
	SetSpToken(v string) *VerifyPhoneWithTokenRequest
	GetSpToken() *string
}

type VerifyPhoneWithTokenRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1380000****
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The token for phone number verification that is obtained by the JavaScript SDK.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dfafdafad542****
	SpToken *string `json:"SpToken,omitempty" xml:"SpToken,omitempty"`
}

func (s VerifyPhoneWithTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyPhoneWithTokenRequest) GoString() string {
	return s.String()
}

func (s *VerifyPhoneWithTokenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *VerifyPhoneWithTokenRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *VerifyPhoneWithTokenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *VerifyPhoneWithTokenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *VerifyPhoneWithTokenRequest) GetSpToken() *string {
	return s.SpToken
}

func (s *VerifyPhoneWithTokenRequest) SetOwnerId(v int64) *VerifyPhoneWithTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyPhoneWithTokenRequest) SetPhoneNumber(v string) *VerifyPhoneWithTokenRequest {
	s.PhoneNumber = &v
	return s
}

func (s *VerifyPhoneWithTokenRequest) SetResourceOwnerAccount(v string) *VerifyPhoneWithTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VerifyPhoneWithTokenRequest) SetResourceOwnerId(v int64) *VerifyPhoneWithTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *VerifyPhoneWithTokenRequest) SetSpToken(v string) *VerifyPhoneWithTokenRequest {
	s.SpToken = &v
	return s
}

func (s *VerifyPhoneWithTokenRequest) Validate() error {
	return dara.Validate(s)
}
