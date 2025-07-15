// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteLiveDomainRequest
	GetDomainName() *string
	SetOwnerAccount(v string) *DeleteLiveDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteLiveDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DeleteLiveDomainRequest
	GetSecurityToken() *string
}

type DeleteLiveDomainRequest struct {
	// The ingest domain or streaming domain that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteLiveDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteLiveDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteLiveDomainRequest) SetDomainName(v string) *DeleteLiveDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveDomainRequest) SetOwnerAccount(v string) *DeleteLiveDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteLiveDomainRequest) SetOwnerId(v int64) *DeleteLiveDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveDomainRequest) SetSecurityToken(v string) *DeleteLiveDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLiveDomainRequest) Validate() error {
	return dara.Validate(s)
}
