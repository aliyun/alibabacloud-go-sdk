// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCdnDomainExistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *CheckCdnDomainExistRequest
	GetDomainName() *string
	SetOwnerAccount(v string) *CheckCdnDomainExistRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckCdnDomainExistRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *CheckCdnDomainExistRequest
	GetSecurityToken() *string
}

type CheckCdnDomainExistRequest struct {
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

func (s CheckCdnDomainExistRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCdnDomainExistRequest) GoString() string {
	return s.String()
}

func (s *CheckCdnDomainExistRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CheckCdnDomainExistRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckCdnDomainExistRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckCdnDomainExistRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CheckCdnDomainExistRequest) SetDomainName(v string) *CheckCdnDomainExistRequest {
	s.DomainName = &v
	return s
}

func (s *CheckCdnDomainExistRequest) SetOwnerAccount(v string) *CheckCdnDomainExistRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckCdnDomainExistRequest) SetOwnerId(v int64) *CheckCdnDomainExistRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckCdnDomainExistRequest) SetSecurityToken(v string) *CheckCdnDomainExistRequest {
	s.SecurityToken = &v
	return s
}

func (s *CheckCdnDomainExistRequest) Validate() error {
	return dara.Validate(s)
}
