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
	// Sender\\"s email address
	//
	// This parameter is required.
	//
	// example:
	//
	// test1@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Reply-to address
	//
	// example:
	//
	// test2@example.com
	ReplyAddress         *string `json:"ReplyAddress,omitempty" xml:"ReplyAddress,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Type of sending. Values:
	//
	// - batch: Bulk emails
	//
	// - trigger: Triggered emails
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
