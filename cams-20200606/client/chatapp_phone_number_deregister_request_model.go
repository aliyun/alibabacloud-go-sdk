// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappPhoneNumberDeregisterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *ChatappPhoneNumberDeregisterRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *ChatappPhoneNumberDeregisterRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *ChatappPhoneNumberDeregisterRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *ChatappPhoneNumberDeregisterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ChatappPhoneNumberDeregisterRequest
	GetResourceOwnerId() *int64
}

type ChatappPhoneNumberDeregisterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
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

func (s ChatappPhoneNumberDeregisterRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatappPhoneNumberDeregisterRequest) GoString() string {
	return s.String()
}

func (s *ChatappPhoneNumberDeregisterRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ChatappPhoneNumberDeregisterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChatappPhoneNumberDeregisterRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ChatappPhoneNumberDeregisterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ChatappPhoneNumberDeregisterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ChatappPhoneNumberDeregisterRequest) SetCustSpaceId(v string) *ChatappPhoneNumberDeregisterRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ChatappPhoneNumberDeregisterRequest) SetOwnerId(v int64) *ChatappPhoneNumberDeregisterRequest {
	s.OwnerId = &v
	return s
}

func (s *ChatappPhoneNumberDeregisterRequest) SetPhoneNumber(v string) *ChatappPhoneNumberDeregisterRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ChatappPhoneNumberDeregisterRequest) SetResourceOwnerAccount(v string) *ChatappPhoneNumberDeregisterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ChatappPhoneNumberDeregisterRequest) SetResourceOwnerId(v int64) *ChatappPhoneNumberDeregisterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ChatappPhoneNumberDeregisterRequest) Validate() error {
	return dara.Validate(s)
}
