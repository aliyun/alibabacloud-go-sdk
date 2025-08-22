// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteDcdnDomainRequest
	GetDomainName() *string
	SetOwnerAccount(v string) *DeleteDcdnDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDcdnDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteDcdnDomainRequest
	GetSecurityToken() *string
}

type DeleteDcdnDomainRequest struct {
	// The accelerated domain name to be deleted. You can specify only one domain name.
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

func (s DeleteDcdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteDcdnDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDcdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDcdnDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteDcdnDomainRequest) SetDomainName(v string) *DeleteDcdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteDcdnDomainRequest) SetOwnerAccount(v string) *DeleteDcdnDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDcdnDomainRequest) SetOwnerId(v int64) *DeleteDcdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDcdnDomainRequest) SetSecurityToken(v string) *DeleteDcdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteDcdnDomainRequest) Validate() error {
	return dara.Validate(s)
}
