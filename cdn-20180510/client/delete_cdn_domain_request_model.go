// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteCdnDomainRequest
	GetDomainName() *string
	SetOwnerAccount(v string) *DeleteCdnDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCdnDomainRequest
	GetOwnerId() *int64
}

type DeleteCdnDomainRequest struct {
	// The accelerated domain name that you want to remove. You can specify only one domain name in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteCdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCdnDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteCdnDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteCdnDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCdnDomainRequest) SetDomainName(v string) *DeleteCdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteCdnDomainRequest) SetOwnerAccount(v string) *DeleteCdnDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCdnDomainRequest) SetOwnerId(v int64) *DeleteCdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCdnDomainRequest) Validate() error {
	return dara.Validate(s)
}
