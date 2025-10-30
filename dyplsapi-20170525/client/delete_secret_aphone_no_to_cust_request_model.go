// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretAPhoneNoToCustRequest interface {
	dara.Model
	String() string
	GoString() string
	SetANoWhiteGroupId(v int64) *DeleteSecretAPhoneNoToCustRequest
	GetANoWhiteGroupId() *int64
	SetOwnerId(v int64) *DeleteSecretAPhoneNoToCustRequest
	GetOwnerId() *int64
	SetPhoneNoA(v string) *DeleteSecretAPhoneNoToCustRequest
	GetPhoneNoA() *string
	SetResourceOwnerAccount(v string) *DeleteSecretAPhoneNoToCustRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSecretAPhoneNoToCustRequest
	GetResourceOwnerId() *int64
}

type DeleteSecretAPhoneNoToCustRequest struct {
	// A号码组ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 51
	ANoWhiteGroupId *int64 `json:"ANoWhiteGroupId,omitempty" xml:"ANoWhiteGroupId,omitempty"`
	OwnerId         *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// A号码
	//
	// This parameter is required.
	//
	// example:
	//
	// 130*****1234
	PhoneNoA             *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteSecretAPhoneNoToCustRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretAPhoneNoToCustRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretAPhoneNoToCustRequest) GetANoWhiteGroupId() *int64 {
	return s.ANoWhiteGroupId
}

func (s *DeleteSecretAPhoneNoToCustRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSecretAPhoneNoToCustRequest) GetPhoneNoA() *string {
	return s.PhoneNoA
}

func (s *DeleteSecretAPhoneNoToCustRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSecretAPhoneNoToCustRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSecretAPhoneNoToCustRequest) SetANoWhiteGroupId(v int64) *DeleteSecretAPhoneNoToCustRequest {
	s.ANoWhiteGroupId = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustRequest) SetOwnerId(v int64) *DeleteSecretAPhoneNoToCustRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustRequest) SetPhoneNoA(v string) *DeleteSecretAPhoneNoToCustRequest {
	s.PhoneNoA = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustRequest) SetResourceOwnerAccount(v string) *DeleteSecretAPhoneNoToCustRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustRequest) SetResourceOwnerId(v int64) *DeleteSecretAPhoneNoToCustRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSecretAPhoneNoToCustRequest) Validate() error {
	return dara.Validate(s)
}
