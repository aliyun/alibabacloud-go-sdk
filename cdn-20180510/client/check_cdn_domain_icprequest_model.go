// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCdnDomainICPRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *CheckCdnDomainICPRequest
	GetDomainName() *string
	SetOwnerAccount(v string) *CheckCdnDomainICPRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckCdnDomainICPRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *CheckCdnDomainICPRequest
	GetSecurityToken() *string
}

type CheckCdnDomainICPRequest struct {
	// The domain name.
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

func (s CheckCdnDomainICPRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCdnDomainICPRequest) GoString() string {
	return s.String()
}

func (s *CheckCdnDomainICPRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CheckCdnDomainICPRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckCdnDomainICPRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckCdnDomainICPRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CheckCdnDomainICPRequest) SetDomainName(v string) *CheckCdnDomainICPRequest {
	s.DomainName = &v
	return s
}

func (s *CheckCdnDomainICPRequest) SetOwnerAccount(v string) *CheckCdnDomainICPRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckCdnDomainICPRequest) SetOwnerId(v int64) *CheckCdnDomainICPRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckCdnDomainICPRequest) SetSecurityToken(v string) *CheckCdnDomainICPRequest {
	s.SecurityToken = &v
	return s
}

func (s *CheckCdnDomainICPRequest) Validate() error {
	return dara.Validate(s)
}
