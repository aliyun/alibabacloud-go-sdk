// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainConfigsRequest
	GetDomainName() *string
	SetFunctionNames(v string) *DescribeVodDomainConfigsRequest
	GetFunctionNames() *string
	SetOwnerId(v int64) *DescribeVodDomainConfigsRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeVodDomainConfigsRequest
	GetSecurityToken() *string
}

type DescribeVodDomainConfigsRequest struct {
	// The domain name for CDN.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The feature name. Separate multiple names with commas (,). For more information, see **Feature description**.
	//
	// This parameter is required.
	//
	// example:
	//
	// filetype_based_ttl_set,set_req_host_header
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeVodDomainConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainConfigsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainConfigsRequest) GetFunctionNames() *string {
	return s.FunctionNames
}

func (s *DescribeVodDomainConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainConfigsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVodDomainConfigsRequest) SetDomainName(v string) *DescribeVodDomainConfigsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainConfigsRequest) SetFunctionNames(v string) *DescribeVodDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *DescribeVodDomainConfigsRequest) SetOwnerId(v int64) *DescribeVodDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainConfigsRequest) SetSecurityToken(v string) *DescribeVodDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVodDomainConfigsRequest) Validate() error {
	return dara.Validate(s)
}
