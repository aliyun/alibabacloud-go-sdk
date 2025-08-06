// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDNADBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBId(v string) *GetDNADBRequest
	GetDBId() *string
	SetOwnerAccount(v string) *GetDNADBRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetDNADBRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetDNADBRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetDNADBRequest
	GetResourceOwnerId() *string
}

type GetDNADBRequest struct {
	// This parameter is required.
	DBId                 *string `json:"DBId,omitempty" xml:"DBId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetDNADBRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDNADBRequest) GoString() string {
	return s.String()
}

func (s *GetDNADBRequest) GetDBId() *string {
	return s.DBId
}

func (s *GetDNADBRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetDNADBRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetDNADBRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetDNADBRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetDNADBRequest) SetDBId(v string) *GetDNADBRequest {
	s.DBId = &v
	return s
}

func (s *GetDNADBRequest) SetOwnerAccount(v string) *GetDNADBRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetDNADBRequest) SetOwnerId(v string) *GetDNADBRequest {
	s.OwnerId = &v
	return s
}

func (s *GetDNADBRequest) SetResourceOwnerAccount(v string) *GetDNADBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetDNADBRequest) SetResourceOwnerId(v string) *GetDNADBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetDNADBRequest) Validate() error {
	return dara.Validate(s)
}
