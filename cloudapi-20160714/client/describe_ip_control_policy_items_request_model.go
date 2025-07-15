// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpControlPolicyItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpControlId(v string) *DescribeIpControlPolicyItemsRequest
	GetIpControlId() *string
	SetPageNumber(v int32) *DescribeIpControlPolicyItemsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIpControlPolicyItemsRequest
	GetPageSize() *int32
	SetPolicyItemId(v string) *DescribeIpControlPolicyItemsRequest
	GetPolicyItemId() *string
	SetSecurityToken(v string) *DescribeIpControlPolicyItemsRequest
	GetSecurityToken() *string
}

type DescribeIpControlPolicyItemsRequest struct {
	// The ID of the ACL. The ID is unique.
	//
	// example:
	//
	// 7ea91319a34d48a09b5c9c871d9768b1
	IpControlId *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// P151617000829241
	PolicyItemId  *string `json:"PolicyItemId,omitempty" xml:"PolicyItemId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeIpControlPolicyItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpControlPolicyItemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpControlPolicyItemsRequest) GetIpControlId() *string {
	return s.IpControlId
}

func (s *DescribeIpControlPolicyItemsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIpControlPolicyItemsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIpControlPolicyItemsRequest) GetPolicyItemId() *string {
	return s.PolicyItemId
}

func (s *DescribeIpControlPolicyItemsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeIpControlPolicyItemsRequest) SetIpControlId(v string) *DescribeIpControlPolicyItemsRequest {
	s.IpControlId = &v
	return s
}

func (s *DescribeIpControlPolicyItemsRequest) SetPageNumber(v int32) *DescribeIpControlPolicyItemsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeIpControlPolicyItemsRequest) SetPageSize(v int32) *DescribeIpControlPolicyItemsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIpControlPolicyItemsRequest) SetPolicyItemId(v string) *DescribeIpControlPolicyItemsRequest {
	s.PolicyItemId = &v
	return s
}

func (s *DescribeIpControlPolicyItemsRequest) SetSecurityToken(v string) *DescribeIpControlPolicyItemsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeIpControlPolicyItemsRequest) Validate() error {
	return dara.Validate(s)
}
