// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappVerifyAndRegisterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *ChatappVerifyAndRegisterRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *ChatappVerifyAndRegisterRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *ChatappVerifyAndRegisterRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *ChatappVerifyAndRegisterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ChatappVerifyAndRegisterRequest
	GetResourceOwnerId() *int64
	SetVerifyCode(v string) *ChatappVerifyAndRegisterRequest
	GetVerifyCode() *string
}

type ChatappVerifyAndRegisterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 29389299388383
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 86138000000
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123466
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s ChatappVerifyAndRegisterRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatappVerifyAndRegisterRequest) GoString() string {
	return s.String()
}

func (s *ChatappVerifyAndRegisterRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ChatappVerifyAndRegisterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChatappVerifyAndRegisterRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ChatappVerifyAndRegisterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ChatappVerifyAndRegisterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ChatappVerifyAndRegisterRequest) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *ChatappVerifyAndRegisterRequest) SetCustSpaceId(v string) *ChatappVerifyAndRegisterRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ChatappVerifyAndRegisterRequest) SetOwnerId(v int64) *ChatappVerifyAndRegisterRequest {
	s.OwnerId = &v
	return s
}

func (s *ChatappVerifyAndRegisterRequest) SetPhoneNumber(v string) *ChatappVerifyAndRegisterRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ChatappVerifyAndRegisterRequest) SetResourceOwnerAccount(v string) *ChatappVerifyAndRegisterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ChatappVerifyAndRegisterRequest) SetResourceOwnerId(v int64) *ChatappVerifyAndRegisterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ChatappVerifyAndRegisterRequest) SetVerifyCode(v string) *ChatappVerifyAndRegisterRequest {
	s.VerifyCode = &v
	return s
}

func (s *ChatappVerifyAndRegisterRequest) Validate() error {
	return dara.Validate(s)
}
