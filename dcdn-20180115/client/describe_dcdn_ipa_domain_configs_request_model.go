// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpaDomainConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnIpaDomainConfigsRequest
	GetDomainName() *string
	SetFunctionNames(v string) *DescribeDcdnIpaDomainConfigsRequest
	GetFunctionNames() *string
	SetOwnerId(v int64) *DescribeDcdnIpaDomainConfigsRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnIpaDomainConfigsRequest
	GetSecurityToken() *string
}

type DescribeDcdnIpaDomainConfigsRequest struct {
	// The accelerated domain name. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the feature. Set the value to protogw, which indicates IP Application Accelerator (IPA).
	//
	// This parameter is required.
	//
	// example:
	//
	// protogw
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnIpaDomainConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainConfigsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnIpaDomainConfigsRequest) GetFunctionNames() *string {
	return s.FunctionNames
}

func (s *DescribeDcdnIpaDomainConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnIpaDomainConfigsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnIpaDomainConfigsRequest) SetDomainName(v string) *DescribeDcdnIpaDomainConfigsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsRequest) SetFunctionNames(v string) *DescribeDcdnIpaDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsRequest) SetOwnerId(v int64) *DescribeDcdnIpaDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsRequest) SetSecurityToken(v string) *DescribeDcdnIpaDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsRequest) Validate() error {
	return dara.Validate(s)
}
