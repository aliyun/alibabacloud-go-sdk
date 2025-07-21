// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveReplyMailAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ApproveReplyMailAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ApproveReplyMailAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ApproveReplyMailAddressRequest
	GetResourceOwnerId() *int64
	SetTicket(v string) *ApproveReplyMailAddressRequest
	GetTicket() *string
}

type ApproveReplyMailAddressRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Email address Ticket credential, part of the string in the verification email\\"s URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// a724068dac9a45d19574375adeca0d7d
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s ApproveReplyMailAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ApproveReplyMailAddressRequest) GoString() string {
	return s.String()
}

func (s *ApproveReplyMailAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ApproveReplyMailAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ApproveReplyMailAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ApproveReplyMailAddressRequest) GetTicket() *string {
	return s.Ticket
}

func (s *ApproveReplyMailAddressRequest) SetOwnerId(v int64) *ApproveReplyMailAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ApproveReplyMailAddressRequest) SetResourceOwnerAccount(v string) *ApproveReplyMailAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ApproveReplyMailAddressRequest) SetResourceOwnerId(v int64) *ApproveReplyMailAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ApproveReplyMailAddressRequest) SetTicket(v string) *ApproveReplyMailAddressRequest {
	s.Ticket = &v
	return s
}

func (s *ApproveReplyMailAddressRequest) Validate() error {
	return dara.Validate(s)
}
