// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeCdnDomainDetailRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeCdnDomainDetailRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeCdnDomainDetailRequest
	GetSecurityToken() *string
}

type DescribeCdnDomainDetailRequest struct {
	// The accelerated domain name. You can specify only one domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnDomainDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainDetailRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnDomainDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCdnDomainDetailRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCdnDomainDetailRequest) SetDomainName(v string) *DescribeCdnDomainDetailRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainDetailRequest) SetOwnerId(v int64) *DescribeCdnDomainDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnDomainDetailRequest) SetSecurityToken(v string) *DescribeCdnDomainDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnDomainDetailRequest) Validate() error {
	return dara.Validate(s)
}
