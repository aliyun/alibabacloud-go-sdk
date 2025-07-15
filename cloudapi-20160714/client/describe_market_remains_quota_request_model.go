// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMarketRemainsQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeMarketRemainsQuotaRequest
	GetDomainName() *string
	SetSecurityToken(v string) *DescribeMarketRemainsQuotaRequest
	GetSecurityToken() *string
}

type DescribeMarketRemainsQuotaRequest struct {
	// The custom domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// *.demo.com
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeMarketRemainsQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMarketRemainsQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeMarketRemainsQuotaRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeMarketRemainsQuotaRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeMarketRemainsQuotaRequest) SetDomainName(v string) *DescribeMarketRemainsQuotaRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeMarketRemainsQuotaRequest) SetSecurityToken(v string) *DescribeMarketRemainsQuotaRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeMarketRemainsQuotaRequest) Validate() error {
	return dara.Validate(s)
}
