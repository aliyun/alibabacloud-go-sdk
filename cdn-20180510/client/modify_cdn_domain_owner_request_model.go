// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdnDomainOwnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *ModifyCdnDomainOwnerRequest
	GetDomainName() *string
	SetOwnerAccount(v string) *ModifyCdnDomainOwnerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCdnDomainOwnerRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *ModifyCdnDomainOwnerRequest
	GetSecurityToken() *string
}

type ModifyCdnDomainOwnerRequest struct {
	// The accelerated domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyCdnDomainOwnerRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdnDomainOwnerRequest) GoString() string {
	return s.String()
}

func (s *ModifyCdnDomainOwnerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ModifyCdnDomainOwnerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCdnDomainOwnerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCdnDomainOwnerRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyCdnDomainOwnerRequest) SetDomainName(v string) *ModifyCdnDomainOwnerRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyCdnDomainOwnerRequest) SetOwnerAccount(v string) *ModifyCdnDomainOwnerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCdnDomainOwnerRequest) SetOwnerId(v int64) *ModifyCdnDomainOwnerRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCdnDomainOwnerRequest) SetSecurityToken(v string) *ModifyCdnDomainOwnerRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyCdnDomainOwnerRequest) Validate() error {
	return dara.Validate(s)
}
