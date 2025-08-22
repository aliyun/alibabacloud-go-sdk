// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnIpaDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteDcdnIpaDomainRequest
	GetDomainName() *string
	SetOwnerAccount(v string) *DeleteDcdnIpaDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDcdnIpaDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteDcdnIpaDomainRequest
	GetSecurityToken() *string
}

type DeleteDcdnIpaDomainRequest struct {
	// The accelerated domain name that you want to delete. You can specify only one accelerated domain name in each request.
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

func (s DeleteDcdnIpaDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnIpaDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnIpaDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteDcdnIpaDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDcdnIpaDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDcdnIpaDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteDcdnIpaDomainRequest) SetDomainName(v string) *DeleteDcdnIpaDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteDcdnIpaDomainRequest) SetOwnerAccount(v string) *DeleteDcdnIpaDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDcdnIpaDomainRequest) SetOwnerId(v int64) *DeleteDcdnIpaDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDcdnIpaDomainRequest) SetSecurityToken(v string) *DeleteDcdnIpaDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteDcdnIpaDomainRequest) Validate() error {
	return dara.Validate(s)
}
