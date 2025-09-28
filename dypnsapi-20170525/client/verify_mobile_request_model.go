// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyMobileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessCode(v string) *VerifyMobileRequest
	GetAccessCode() *string
	SetOutId(v string) *VerifyMobileRequest
	GetOutId() *string
	SetOwnerId(v int64) *VerifyMobileRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *VerifyMobileRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *VerifyMobileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *VerifyMobileRequest
	GetResourceOwnerId() *int64
}

type VerifyMobileRequest struct {
	// The token obtained by the SDK for your app.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dfafdafad542****
	AccessCode *string `json:"AccessCode,omitempty" xml:"AccessCode,omitempty"`
	// The external ID.
	//
	// example:
	//
	// 123456
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13800****00
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s VerifyMobileRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyMobileRequest) GoString() string {
	return s.String()
}

func (s *VerifyMobileRequest) GetAccessCode() *string {
	return s.AccessCode
}

func (s *VerifyMobileRequest) GetOutId() *string {
	return s.OutId
}

func (s *VerifyMobileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *VerifyMobileRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *VerifyMobileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *VerifyMobileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *VerifyMobileRequest) SetAccessCode(v string) *VerifyMobileRequest {
	s.AccessCode = &v
	return s
}

func (s *VerifyMobileRequest) SetOutId(v string) *VerifyMobileRequest {
	s.OutId = &v
	return s
}

func (s *VerifyMobileRequest) SetOwnerId(v int64) *VerifyMobileRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyMobileRequest) SetPhoneNumber(v string) *VerifyMobileRequest {
	s.PhoneNumber = &v
	return s
}

func (s *VerifyMobileRequest) SetResourceOwnerAccount(v string) *VerifyMobileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VerifyMobileRequest) SetResourceOwnerId(v int64) *VerifyMobileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *VerifyMobileRequest) Validate() error {
	return dara.Validate(s)
}
