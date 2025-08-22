// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *DescribeDcdnDomainConfigsRequest
	GetConfigId() *string
	SetDomainName(v string) *DescribeDcdnDomainConfigsRequest
	GetDomainName() *string
	SetFunctionNames(v string) *DescribeDcdnDomainConfigsRequest
	GetFunctionNames() *string
	SetOwnerId(v int64) *DescribeDcdnDomainConfigsRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnDomainConfigsRequest
	GetSecurityToken() *string
}

type DescribeDcdnDomainConfigsRequest struct {
	// The ID of the configuration.
	//
	// example:
	//
	// 5003576
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The names of the features to query. Separate features with commas (,).
	//
	// example:
	//
	// filetype_based_ttl_set,set_req_host_header
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnDomainConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainConfigsRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeDcdnDomainConfigsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainConfigsRequest) GetFunctionNames() *string {
	return s.FunctionNames
}

func (s *DescribeDcdnDomainConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnDomainConfigsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnDomainConfigsRequest) SetConfigId(v string) *DescribeDcdnDomainConfigsRequest {
	s.ConfigId = &v
	return s
}

func (s *DescribeDcdnDomainConfigsRequest) SetDomainName(v string) *DescribeDcdnDomainConfigsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainConfigsRequest) SetFunctionNames(v string) *DescribeDcdnDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *DescribeDcdnDomainConfigsRequest) SetOwnerId(v int64) *DescribeDcdnDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDomainConfigsRequest) SetSecurityToken(v string) *DescribeDcdnDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnDomainConfigsRequest) Validate() error {
	return dara.Validate(s)
}
