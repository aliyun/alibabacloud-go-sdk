// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnDomainDetailRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeDcdnDomainDetailRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnDomainDetailRequest
	GetSecurityToken() *string
}

type DescribeDcdnDomainDetailRequest struct {
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

func (s DescribeDcdnDomainDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainDetailRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnDomainDetailRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnDomainDetailRequest) SetDomainName(v string) *DescribeDcdnDomainDetailRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainDetailRequest) SetOwnerId(v int64) *DescribeDcdnDomainDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainDetailRequest) SetSecurityToken(v string) *DescribeDcdnDomainDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnDomainDetailRequest) Validate() error {
	return dara.Validate(s)
}
