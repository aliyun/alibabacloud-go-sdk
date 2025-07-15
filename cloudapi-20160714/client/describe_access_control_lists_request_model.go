// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessControlListsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclName(v string) *DescribeAccessControlListsRequest
	GetAclName() *string
	SetAddressIPVersion(v string) *DescribeAccessControlListsRequest
	GetAddressIPVersion() *string
	SetPageNumber(v int32) *DescribeAccessControlListsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccessControlListsRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeAccessControlListsRequest
	GetSecurityToken() *string
}

type DescribeAccessControlListsRequest struct {
	// The name of the access control policy.
	//
	// example:
	//
	// testAcl
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// IP版本，可以设置为**ipv4**或者**ipv6**。
	//
	// example:
	//
	// ipv6
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// The page number of the current page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeAccessControlListsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsRequest) GetAclName() *string {
	return s.AclName
}

func (s *DescribeAccessControlListsRequest) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *DescribeAccessControlListsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccessControlListsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessControlListsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeAccessControlListsRequest) SetAclName(v string) *DescribeAccessControlListsRequest {
	s.AclName = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetAddressIPVersion(v string) *DescribeAccessControlListsRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetPageNumber(v int32) *DescribeAccessControlListsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetPageSize(v int32) *DescribeAccessControlListsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessControlListsRequest) SetSecurityToken(v string) *DescribeAccessControlListsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAccessControlListsRequest) Validate() error {
	return dara.Validate(s)
}
