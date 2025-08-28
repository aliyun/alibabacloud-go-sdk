// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *SendVerificationRequest
	GetBizType() *string
	SetOwnerId(v int64) *SendVerificationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SendVerificationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SendVerificationRequest
	GetResourceOwnerId() *int64
	SetTarget(v string) *SendVerificationRequest
	GetTarget() *string
	SetVerifyType(v string) *SendVerificationRequest
	GetVerifyType() *string
}

type SendVerificationRequest struct {
	// The business type. Set the value to **CONTACT**.
	//
	// This parameter is required.
	//
	// example:
	//
	// CONTACT
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The mobile phone number that receives the SMS verification code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 150****0000
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The mode of sending the SMS verification code. Set the value to **SMS**.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s SendVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s SendVerificationRequest) GoString() string {
	return s.String()
}

func (s *SendVerificationRequest) GetBizType() *string {
	return s.BizType
}

func (s *SendVerificationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SendVerificationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SendVerificationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SendVerificationRequest) GetTarget() *string {
	return s.Target
}

func (s *SendVerificationRequest) GetVerifyType() *string {
	return s.VerifyType
}

func (s *SendVerificationRequest) SetBizType(v string) *SendVerificationRequest {
	s.BizType = &v
	return s
}

func (s *SendVerificationRequest) SetOwnerId(v int64) *SendVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *SendVerificationRequest) SetResourceOwnerAccount(v string) *SendVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendVerificationRequest) SetResourceOwnerId(v int64) *SendVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendVerificationRequest) SetTarget(v string) *SendVerificationRequest {
	s.Target = &v
	return s
}

func (s *SendVerificationRequest) SetVerifyType(v string) *SendVerificationRequest {
	s.VerifyType = &v
	return s
}

func (s *SendVerificationRequest) Validate() error {
	return dara.Validate(s)
}
