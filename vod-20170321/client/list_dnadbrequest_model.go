// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDNADBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ListDNADBRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *ListDNADBRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *ListDNADBRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *ListDNADBRequest
	GetResourceOwnerId() *string
}

type ListDNADBRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListDNADBRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDNADBRequest) GoString() string {
	return s.String()
}

func (s *ListDNADBRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListDNADBRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListDNADBRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListDNADBRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *ListDNADBRequest) SetOwnerAccount(v string) *ListDNADBRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListDNADBRequest) SetOwnerId(v string) *ListDNADBRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDNADBRequest) SetResourceOwnerAccount(v string) *ListDNADBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListDNADBRequest) SetResourceOwnerId(v string) *ListDNADBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListDNADBRequest) Validate() error {
	return dara.Validate(s)
}
