// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *ListContactsRequest
	GetBizType() *string
	SetContactId(v int64) *ListContactsRequest
	GetContactId() *int64
	SetOwnerId(v int64) *ListContactsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListContactsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListContactsRequest
	GetResourceOwnerId() *int64
}

type ListContactsRequest struct {
	// example:
	//
	// dytns
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 1298647
	ContactId            *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContactsRequest) GoString() string {
	return s.String()
}

func (s *ListContactsRequest) GetBizType() *string {
	return s.BizType
}

func (s *ListContactsRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *ListContactsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListContactsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListContactsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListContactsRequest) SetBizType(v string) *ListContactsRequest {
	s.BizType = &v
	return s
}

func (s *ListContactsRequest) SetContactId(v int64) *ListContactsRequest {
	s.ContactId = &v
	return s
}

func (s *ListContactsRequest) SetOwnerId(v int64) *ListContactsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListContactsRequest) SetResourceOwnerAccount(v string) *ListContactsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListContactsRequest) SetResourceOwnerId(v int64) *ListContactsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListContactsRequest) Validate() error {
	return dara.Validate(s)
}
