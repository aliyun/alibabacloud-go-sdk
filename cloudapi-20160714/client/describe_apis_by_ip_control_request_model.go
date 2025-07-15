// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisByIpControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpControlId(v string) *DescribeApisByIpControlRequest
	GetIpControlId() *string
	SetPageNumber(v int32) *DescribeApisByIpControlRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApisByIpControlRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeApisByIpControlRequest
	GetSecurityToken() *string
}

type DescribeApisByIpControlRequest struct {
	// The ID of the ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	IpControlId *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApisByIpControlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByIpControlRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisByIpControlRequest) GetIpControlId() *string {
	return s.IpControlId
}

func (s *DescribeApisByIpControlRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApisByIpControlRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApisByIpControlRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApisByIpControlRequest) SetIpControlId(v string) *DescribeApisByIpControlRequest {
	s.IpControlId = &v
	return s
}

func (s *DescribeApisByIpControlRequest) SetPageNumber(v int32) *DescribeApisByIpControlRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByIpControlRequest) SetPageSize(v int32) *DescribeApisByIpControlRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByIpControlRequest) SetSecurityToken(v string) *DescribeApisByIpControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApisByIpControlRequest) Validate() error {
	return dara.Validate(s)
}
