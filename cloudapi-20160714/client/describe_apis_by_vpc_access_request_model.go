// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisByVpcAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeApisByVpcAccessRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApisByVpcAccessRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeApisByVpcAccessRequest
	GetSecurityToken() *string
	SetVpcName(v string) *DescribeApisByVpcAccessRequest
	GetVpcName() *string
}

type DescribeApisByVpcAccessRequest struct {
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The name of the VPC access authorization.
	//
	// example:
	//
	// lynkco-iov-uat
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeApisByVpcAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByVpcAccessRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisByVpcAccessRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApisByVpcAccessRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApisByVpcAccessRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApisByVpcAccessRequest) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeApisByVpcAccessRequest) SetPageNumber(v int32) *DescribeApisByVpcAccessRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByVpcAccessRequest) SetPageSize(v int32) *DescribeApisByVpcAccessRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByVpcAccessRequest) SetSecurityToken(v string) *DescribeApisByVpcAccessRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApisByVpcAccessRequest) SetVpcName(v string) *DescribeApisByVpcAccessRequest {
	s.VpcName = &v
	return s
}

func (s *DescribeApisByVpcAccessRequest) Validate() error {
	return dara.Validate(s)
}
