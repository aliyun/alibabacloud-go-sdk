// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsDomainConfigsRequest
	GetDomainName() *string
	SetFunctionNames(v string) *DescribeVsDomainConfigsRequest
	GetFunctionNames() *string
	SetOwnerId(v int64) *DescribeVsDomainConfigsRequest
	GetOwnerId() *int64
}

type DescribeVsDomainConfigsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// path_based_ttl_set,oss_auth
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVsDomainConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainConfigsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainConfigsRequest) GetFunctionNames() *string {
	return s.FunctionNames
}

func (s *DescribeVsDomainConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsDomainConfigsRequest) SetDomainName(v string) *DescribeVsDomainConfigsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainConfigsRequest) SetFunctionNames(v string) *DescribeVsDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *DescribeVsDomainConfigsRequest) SetOwnerId(v int64) *DescribeVsDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainConfigsRequest) Validate() error {
	return dara.Validate(s)
}
