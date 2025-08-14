// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDNADBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBId(v string) *DeleteDNADBRequest
	GetDBId() *string
	SetOwnerAccount(v string) *DeleteDNADBRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDNADBRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteDNADBRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDNADBRequest
	GetResourceOwnerId() *int64
}

type DeleteDNADBRequest struct {
	// The ID of the media fingerprint library that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// fb712a6890464059b1b2ea7c8647****
	DBId                 *string `json:"DBId,omitempty" xml:"DBId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDNADBRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDNADBRequest) GoString() string {
	return s.String()
}

func (s *DeleteDNADBRequest) GetDBId() *string {
	return s.DBId
}

func (s *DeleteDNADBRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDNADBRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDNADBRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDNADBRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDNADBRequest) SetDBId(v string) *DeleteDNADBRequest {
	s.DBId = &v
	return s
}

func (s *DeleteDNADBRequest) SetOwnerAccount(v string) *DeleteDNADBRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDNADBRequest) SetOwnerId(v int64) *DeleteDNADBRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDNADBRequest) SetResourceOwnerAccount(v string) *DeleteDNADBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDNADBRequest) SetResourceOwnerId(v int64) *DeleteDNADBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDNADBRequest) Validate() error {
	return dara.Validate(s)
}
