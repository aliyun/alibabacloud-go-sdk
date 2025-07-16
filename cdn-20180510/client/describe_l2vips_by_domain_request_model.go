// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeL2VipsByDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeL2VipsByDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeL2VipsByDomainRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeL2VipsByDomainRequest
	GetSecurityToken() *string
}

type DescribeL2VipsByDomainRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
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

func (s DescribeL2VipsByDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeL2VipsByDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeL2VipsByDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeL2VipsByDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeL2VipsByDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeL2VipsByDomainRequest) SetDomainName(v string) *DescribeL2VipsByDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeL2VipsByDomainRequest) SetOwnerId(v int64) *DescribeL2VipsByDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeL2VipsByDomainRequest) SetSecurityToken(v string) *DescribeL2VipsByDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeL2VipsByDomainRequest) Validate() error {
	return dara.Validate(s)
}
