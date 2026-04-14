// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReceiverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesc(v string) *CreateReceiverRequest
	GetDesc() *string
	SetOwnerId(v int64) *CreateReceiverRequest
	GetOwnerId() *int64
	SetReceiversAlias(v string) *CreateReceiverRequest
	GetReceiversAlias() *string
	SetReceiversName(v string) *CreateReceiverRequest
	GetReceiversName() *string
	SetResourceOwnerAccount(v string) *CreateReceiverRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateReceiverRequest
	GetResourceOwnerId() *int64
}

type CreateReceiverRequest struct {
	// List description.
	//
	// example:
	//
	// the description
	Desc    *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// List alias, an email address less than 30 characters long.
	//
	// This parameter is required.
	//
	// example:
	//
	// a***@example.net
	ReceiversAlias *string `json:"ReceiversAlias,omitempty" xml:"ReceiversAlias,omitempty"`
	// List name, must be unique, with a length of 1-30 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ReceiversName        *string `json:"ReceiversName,omitempty" xml:"ReceiversName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateReceiverRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReceiverRequest) GoString() string {
	return s.String()
}

func (s *CreateReceiverRequest) GetDesc() *string {
	return s.Desc
}

func (s *CreateReceiverRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateReceiverRequest) GetReceiversAlias() *string {
	return s.ReceiversAlias
}

func (s *CreateReceiverRequest) GetReceiversName() *string {
	return s.ReceiversName
}

func (s *CreateReceiverRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateReceiverRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateReceiverRequest) SetDesc(v string) *CreateReceiverRequest {
	s.Desc = &v
	return s
}

func (s *CreateReceiverRequest) SetOwnerId(v int64) *CreateReceiverRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateReceiverRequest) SetReceiversAlias(v string) *CreateReceiverRequest {
	s.ReceiversAlias = &v
	return s
}

func (s *CreateReceiverRequest) SetReceiversName(v string) *CreateReceiverRequest {
	s.ReceiversName = &v
	return s
}

func (s *CreateReceiverRequest) SetResourceOwnerAccount(v string) *CreateReceiverRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateReceiverRequest) SetResourceOwnerId(v int64) *CreateReceiverRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateReceiverRequest) Validate() error {
	return dara.Validate(s)
}
