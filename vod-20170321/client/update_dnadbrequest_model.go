// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDNADBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBDescription(v string) *UpdateDNADBRequest
	GetDBDescription() *string
	SetDBId(v string) *UpdateDNADBRequest
	GetDBId() *string
	SetDBName(v string) *UpdateDNADBRequest
	GetDBName() *string
	SetOwnerAccount(v string) *UpdateDNADBRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *UpdateDNADBRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *UpdateDNADBRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *UpdateDNADBRequest
	GetResourceOwnerId() *string
}

type UpdateDNADBRequest struct {
	DBDescription *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	// This parameter is required.
	DBId                 *string `json:"DBId,omitempty" xml:"DBId,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateDNADBRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDNADBRequest) GoString() string {
	return s.String()
}

func (s *UpdateDNADBRequest) GetDBDescription() *string {
	return s.DBDescription
}

func (s *UpdateDNADBRequest) GetDBId() *string {
	return s.DBId
}

func (s *UpdateDNADBRequest) GetDBName() *string {
	return s.DBName
}

func (s *UpdateDNADBRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateDNADBRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *UpdateDNADBRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateDNADBRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *UpdateDNADBRequest) SetDBDescription(v string) *UpdateDNADBRequest {
	s.DBDescription = &v
	return s
}

func (s *UpdateDNADBRequest) SetDBId(v string) *UpdateDNADBRequest {
	s.DBId = &v
	return s
}

func (s *UpdateDNADBRequest) SetDBName(v string) *UpdateDNADBRequest {
	s.DBName = &v
	return s
}

func (s *UpdateDNADBRequest) SetOwnerAccount(v string) *UpdateDNADBRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateDNADBRequest) SetOwnerId(v string) *UpdateDNADBRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateDNADBRequest) SetResourceOwnerAccount(v string) *UpdateDNADBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateDNADBRequest) SetResourceOwnerId(v string) *UpdateDNADBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateDNADBRequest) Validate() error {
	return dara.Validate(s)
}
