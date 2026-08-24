// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMailAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateMailAddressRequest
	GetAccountName() *string
	SetAddressType(v string) *CreateMailAddressRequest
	GetAddressType() *string
	SetOwnerId(v int64) *CreateMailAddressRequest
	GetOwnerId() *int64
	SetReplyAddress(v string) *CreateMailAddressRequest
	GetReplyAddress() *string
	SetResourceOwnerAccount(v string) *CreateMailAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateMailAddressRequest
	GetResourceOwnerId() *int64
	SetSendtype(v string) *CreateMailAddressRequest
	GetSendtype() *string
}

type CreateMailAddressRequest struct {
	// The sender address.
	//
	// This parameter is required.
	//
	// example:
	//
	// Account+@+domain
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The type of the address to create. Valid values:
	//
	// EXTERNAL: The domain name of the address to create has not been created in this system.
	//
	// INTERNAL: The domain name of the address to create has already been created in this system.
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The reply-to address.
	//
	// example:
	//
	// test1***@example.net
	ReplyAddress         *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of email. Valid values:
	//
	// - batch: batch email
	//
	// - trigger: triggered email
	//
	// This parameter is required.
	//
	// example:
	//
	// batch
	Sendtype *string `json:"Sendtype,omitempty" xml:"Sendtype,omitempty"`
}

func (s CreateMailAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMailAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateMailAddressRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateMailAddressRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *CreateMailAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateMailAddressRequest) GetReplyAddress() *string {
	return s.ReplyAddress
}

func (s *CreateMailAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateMailAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateMailAddressRequest) GetSendtype() *string {
	return s.Sendtype
}

func (s *CreateMailAddressRequest) SetAccountName(v string) *CreateMailAddressRequest {
	s.AccountName = &v
	return s
}

func (s *CreateMailAddressRequest) SetAddressType(v string) *CreateMailAddressRequest {
	s.AddressType = &v
	return s
}

func (s *CreateMailAddressRequest) SetOwnerId(v int64) *CreateMailAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMailAddressRequest) SetReplyAddress(v string) *CreateMailAddressRequest {
	s.ReplyAddress = &v
	return s
}

func (s *CreateMailAddressRequest) SetResourceOwnerAccount(v string) *CreateMailAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMailAddressRequest) SetResourceOwnerId(v int64) *CreateMailAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMailAddressRequest) SetSendtype(v string) *CreateMailAddressRequest {
	s.Sendtype = &v
	return s
}

func (s *CreateMailAddressRequest) Validate() error {
	return dara.Validate(s)
}
