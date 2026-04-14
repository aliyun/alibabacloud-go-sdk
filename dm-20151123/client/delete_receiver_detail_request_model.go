// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReceiverDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *DeleteReceiverDetailRequest
	GetEmail() *string
	SetOwnerId(v int64) *DeleteReceiverDetailRequest
	GetOwnerId() *int64
	SetReceiverId(v string) *DeleteReceiverDetailRequest
	GetReceiverId() *string
	SetResourceOwnerAccount(v string) *DeleteReceiverDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteReceiverDetailRequest
	GetResourceOwnerId() *int64
}

type DeleteReceiverDetailRequest struct {
	// The single recipient to be deleted from the recipient list
	//
	// example:
	//
	// test@example.com
	Email   *string `json:"Email,omitempty" xml:"Email,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Recipient list ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 53228b7d80c36257927ecd029ccd3c9a
	ReceiverId           *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteReceiverDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteReceiverDetailRequest) GoString() string {
	return s.String()
}

func (s *DeleteReceiverDetailRequest) GetEmail() *string {
	return s.Email
}

func (s *DeleteReceiverDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteReceiverDetailRequest) GetReceiverId() *string {
	return s.ReceiverId
}

func (s *DeleteReceiverDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteReceiverDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteReceiverDetailRequest) SetEmail(v string) *DeleteReceiverDetailRequest {
	s.Email = &v
	return s
}

func (s *DeleteReceiverDetailRequest) SetOwnerId(v int64) *DeleteReceiverDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteReceiverDetailRequest) SetReceiverId(v string) *DeleteReceiverDetailRequest {
	s.ReceiverId = &v
	return s
}

func (s *DeleteReceiverDetailRequest) SetResourceOwnerAccount(v string) *DeleteReceiverDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteReceiverDetailRequest) SetResourceOwnerId(v int64) *DeleteReceiverDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteReceiverDetailRequest) Validate() error {
	return dara.Validate(s)
}
