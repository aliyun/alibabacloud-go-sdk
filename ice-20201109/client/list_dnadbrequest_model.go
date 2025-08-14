// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDNADBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBIds(v string) *ListDNADBRequest
	GetDBIds() *string
	SetOwnerAccount(v string) *ListDNADBRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListDNADBRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListDNADBRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListDNADBRequest
	GetResourceOwnerId() *int64
}

type ListDNADBRequest struct {
	// The IDs of the media fingerprint libraries. We recommend that you query at most 10 libraries at a time. Separate multiple library IDs with commas (,).
	//
	// example:
	//
	// 2288c6ca184c0e47098a5b665e2a12****,78dc866518b843259669df58ed30****
	DBIds                *string `json:"DBIds,omitempty" xml:"DBIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListDNADBRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDNADBRequest) GoString() string {
	return s.String()
}

func (s *ListDNADBRequest) GetDBIds() *string {
	return s.DBIds
}

func (s *ListDNADBRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListDNADBRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListDNADBRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListDNADBRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListDNADBRequest) SetDBIds(v string) *ListDNADBRequest {
	s.DBIds = &v
	return s
}

func (s *ListDNADBRequest) SetOwnerAccount(v string) *ListDNADBRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListDNADBRequest) SetOwnerId(v int64) *ListDNADBRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDNADBRequest) SetResourceOwnerAccount(v string) *ListDNADBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListDNADBRequest) SetResourceOwnerId(v int64) *ListDNADBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListDNADBRequest) Validate() error {
	return dara.Validate(s)
}
