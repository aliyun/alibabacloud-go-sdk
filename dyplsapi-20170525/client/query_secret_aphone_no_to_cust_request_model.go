// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecretAPhoneNoToCustRequest interface {
	dara.Model
	String() string
	GoString() string
	SetANoWhiteGroupId(v int64) *QuerySecretAPhoneNoToCustRequest
	GetANoWhiteGroupId() *int64
	SetOwnerId(v int64) *QuerySecretAPhoneNoToCustRequest
	GetOwnerId() *int64
	SetPhoneNoA(v string) *QuerySecretAPhoneNoToCustRequest
	GetPhoneNoA() *string
	SetResourceOwnerAccount(v string) *QuerySecretAPhoneNoToCustRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySecretAPhoneNoToCustRequest
	GetResourceOwnerId() *int64
}

type QuerySecretAPhoneNoToCustRequest struct {
	// 号码组ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 58
	ANoWhiteGroupId *int64 `json:"ANoWhiteGroupId,omitempty" xml:"ANoWhiteGroupId,omitempty"`
	OwnerId         *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// A号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 130*****8888
	PhoneNoA             *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QuerySecretAPhoneNoToCustRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretAPhoneNoToCustRequest) GoString() string {
	return s.String()
}

func (s *QuerySecretAPhoneNoToCustRequest) GetANoWhiteGroupId() *int64 {
	return s.ANoWhiteGroupId
}

func (s *QuerySecretAPhoneNoToCustRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySecretAPhoneNoToCustRequest) GetPhoneNoA() *string {
	return s.PhoneNoA
}

func (s *QuerySecretAPhoneNoToCustRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySecretAPhoneNoToCustRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySecretAPhoneNoToCustRequest) SetANoWhiteGroupId(v int64) *QuerySecretAPhoneNoToCustRequest {
	s.ANoWhiteGroupId = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustRequest) SetOwnerId(v int64) *QuerySecretAPhoneNoToCustRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustRequest) SetPhoneNoA(v string) *QuerySecretAPhoneNoToCustRequest {
	s.PhoneNoA = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustRequest) SetResourceOwnerAccount(v string) *QuerySecretAPhoneNoToCustRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustRequest) SetResourceOwnerId(v int64) *QuerySecretAPhoneNoToCustRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustRequest) Validate() error {
	return dara.Validate(s)
}
