// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultPlayDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SetDefaultPlayDomainRequest
	GetDomainName() *string
	SetOwnerAccount(v string) *SetDefaultPlayDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SetDefaultPlayDomainRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SetDefaultPlayDomainRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SetDefaultPlayDomainRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *SetDefaultPlayDomainRequest
	GetResourceRealOwnerId() *int64
}

type SetDefaultPlayDomainRequest struct {
	// This parameter is required.
	DomainName           *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
}

func (s SetDefaultPlayDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultPlayDomainRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultPlayDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetDefaultPlayDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetDefaultPlayDomainRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SetDefaultPlayDomainRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetDefaultPlayDomainRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SetDefaultPlayDomainRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *SetDefaultPlayDomainRequest) SetDomainName(v string) *SetDefaultPlayDomainRequest {
	s.DomainName = &v
	return s
}

func (s *SetDefaultPlayDomainRequest) SetOwnerAccount(v string) *SetDefaultPlayDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetDefaultPlayDomainRequest) SetOwnerId(v string) *SetDefaultPlayDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDefaultPlayDomainRequest) SetResourceOwnerAccount(v string) *SetDefaultPlayDomainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetDefaultPlayDomainRequest) SetResourceOwnerId(v string) *SetDefaultPlayDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetDefaultPlayDomainRequest) SetResourceRealOwnerId(v int64) *SetDefaultPlayDomainRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *SetDefaultPlayDomainRequest) Validate() error {
	return dara.Validate(s)
}
