// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteVodDomainRequest
	GetDomainName() *string
	SetOwnerAccount(v string) *DeleteVodDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVodDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteVodDomainRequest
	GetSecurityToken() *string
}

type DeleteVodDomainRequest struct {
	// The domain name for CDN that you want to delete.
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

func (s DeleteVodDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteVodDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteVodDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVodDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVodDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteVodDomainRequest) SetDomainName(v string) *DeleteVodDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteVodDomainRequest) SetOwnerAccount(v string) *DeleteVodDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVodDomainRequest) SetOwnerId(v int64) *DeleteVodDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVodDomainRequest) SetSecurityToken(v string) *DeleteVodDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteVodDomainRequest) Validate() error {
	return dara.Validate(s)
}
