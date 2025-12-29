// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v int64) *DeleteContactsRequest
	GetContactId() *int64
	SetOwnerId(v int64) *DeleteContactsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteContactsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteContactsRequest
	GetResourceOwnerId() *int64
}

type DeleteContactsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5454735262
	ContactId            *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactsRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactsRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *DeleteContactsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteContactsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteContactsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteContactsRequest) SetContactId(v int64) *DeleteContactsRequest {
	s.ContactId = &v
	return s
}

func (s *DeleteContactsRequest) SetOwnerId(v int64) *DeleteContactsRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteContactsRequest) SetResourceOwnerAccount(v string) *DeleteContactsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteContactsRequest) SetResourceOwnerId(v int64) *DeleteContactsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteContactsRequest) Validate() error {
	return dara.Validate(s)
}
