// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPhoneNumberOnlineTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *QueryPhoneNumberOnlineTimeRequest
	GetAuthCode() *string
	SetInputNumber(v string) *QueryPhoneNumberOnlineTimeRequest
	GetInputNumber() *string
	SetMask(v string) *QueryPhoneNumberOnlineTimeRequest
	GetMask() *string
	SetOwnerId(v int64) *QueryPhoneNumberOnlineTimeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryPhoneNumberOnlineTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryPhoneNumberOnlineTimeRequest
	GetResourceOwnerId() *int64
}

type QueryPhoneNumberOnlineTimeRequest struct {
	// example:
	//
	// 示例值示例值示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// example:
	//
	// 示例值
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// example:
	//
	// 示例值示例值
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryPhoneNumberOnlineTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPhoneNumberOnlineTimeRequest) GoString() string {
	return s.String()
}

func (s *QueryPhoneNumberOnlineTimeRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *QueryPhoneNumberOnlineTimeRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *QueryPhoneNumberOnlineTimeRequest) GetMask() *string {
	return s.Mask
}

func (s *QueryPhoneNumberOnlineTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryPhoneNumberOnlineTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryPhoneNumberOnlineTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryPhoneNumberOnlineTimeRequest) SetAuthCode(v string) *QueryPhoneNumberOnlineTimeRequest {
	s.AuthCode = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeRequest) SetInputNumber(v string) *QueryPhoneNumberOnlineTimeRequest {
	s.InputNumber = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeRequest) SetMask(v string) *QueryPhoneNumberOnlineTimeRequest {
	s.Mask = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeRequest) SetOwnerId(v int64) *QueryPhoneNumberOnlineTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeRequest) SetResourceOwnerAccount(v string) *QueryPhoneNumberOnlineTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeRequest) SetResourceOwnerId(v int64) *QueryPhoneNumberOnlineTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPhoneNumberOnlineTimeRequest) Validate() error {
	return dara.Validate(s)
}
