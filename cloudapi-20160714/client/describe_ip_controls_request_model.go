// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpControlsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpControlId(v string) *DescribeIpControlsRequest
	GetIpControlId() *string
	SetIpControlName(v string) *DescribeIpControlsRequest
	GetIpControlName() *string
	SetIpControlType(v string) *DescribeIpControlsRequest
	GetIpControlType() *string
	SetPageNumber(v int32) *DescribeIpControlsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIpControlsRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeIpControlsRequest
	GetSecurityToken() *string
}

type DescribeIpControlsRequest struct {
	// The ID of the ACL. The ID is unique.
	//
	// example:
	//
	// 7ea91319a34d48a09b5c9c871d9768b1
	IpControlId *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	// The name of the ACL.
	//
	// example:
	//
	// ACL test
	IpControlName *string `json:"IpControlName,omitempty" xml:"IpControlName,omitempty"`
	// The type of the ACL. Valid values:
	//
	// 	- **ALLOW**: a whitelist
	//
	// 	- **REFUSE**: a blacklist
	//
	// example:
	//
	// ALLOW
	IpControlType *string `json:"IpControlType,omitempty" xml:"IpControlType,omitempty"`
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

func (s DescribeIpControlsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpControlsRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpControlsRequest) GetIpControlId() *string {
	return s.IpControlId
}

func (s *DescribeIpControlsRequest) GetIpControlName() *string {
	return s.IpControlName
}

func (s *DescribeIpControlsRequest) GetIpControlType() *string {
	return s.IpControlType
}

func (s *DescribeIpControlsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIpControlsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIpControlsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeIpControlsRequest) SetIpControlId(v string) *DescribeIpControlsRequest {
	s.IpControlId = &v
	return s
}

func (s *DescribeIpControlsRequest) SetIpControlName(v string) *DescribeIpControlsRequest {
	s.IpControlName = &v
	return s
}

func (s *DescribeIpControlsRequest) SetIpControlType(v string) *DescribeIpControlsRequest {
	s.IpControlType = &v
	return s
}

func (s *DescribeIpControlsRequest) SetPageNumber(v int32) *DescribeIpControlsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeIpControlsRequest) SetPageSize(v int32) *DescribeIpControlsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIpControlsRequest) SetSecurityToken(v string) *DescribeIpControlsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeIpControlsRequest) Validate() error {
	return dara.Validate(s)
}
