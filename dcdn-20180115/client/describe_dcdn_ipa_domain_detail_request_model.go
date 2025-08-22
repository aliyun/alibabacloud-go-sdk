// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpaDomainDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnIpaDomainDetailRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeDcdnIpaDomainDetailRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnIpaDomainDetailRequest
	GetSecurityToken() *string
}

type DescribeDcdnIpaDomainDetailRequest struct {
	// The accelerated domain names for which you want to query basic information. You can specify only one domain name in each request.
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

func (s DescribeDcdnIpaDomainDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainDetailRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnIpaDomainDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnIpaDomainDetailRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnIpaDomainDetailRequest) SetDomainName(v string) *DescribeDcdnIpaDomainDetailRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailRequest) SetOwnerId(v int64) *DescribeDcdnIpaDomainDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailRequest) SetSecurityToken(v string) *DescribeDcdnIpaDomainDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailRequest) Validate() error {
	return dara.Validate(s)
}
