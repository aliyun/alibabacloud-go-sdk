// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactsByIdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContacts(v string) *DeleteContactsByIdsRequest
	GetContacts() *string
	SetOwnerId(v int64) *DeleteContactsByIdsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteContactsByIdsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteContactsByIdsRequest
	GetResourceOwnerId() *int64
}

type DeleteContactsByIdsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	Contacts             *string `json:"Contacts,omitempty" xml:"Contacts,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteContactsByIdsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactsByIdsRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactsByIdsRequest) GetContacts() *string {
	return s.Contacts
}

func (s *DeleteContactsByIdsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteContactsByIdsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteContactsByIdsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteContactsByIdsRequest) SetContacts(v string) *DeleteContactsByIdsRequest {
	s.Contacts = &v
	return s
}

func (s *DeleteContactsByIdsRequest) SetOwnerId(v int64) *DeleteContactsByIdsRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteContactsByIdsRequest) SetResourceOwnerAccount(v string) *DeleteContactsByIdsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteContactsByIdsRequest) SetResourceOwnerId(v int64) *DeleteContactsByIdsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteContactsByIdsRequest) Validate() error {
	return dara.Validate(s)
}
