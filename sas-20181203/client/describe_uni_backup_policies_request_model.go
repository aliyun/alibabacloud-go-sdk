// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUniBackupPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeUniBackupPoliciesRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeUniBackupPoliciesRequest
	GetPageSize() *int32
	SetPolicyName(v string) *DescribeUniBackupPoliciesRequest
	GetPolicyName() *string
}

type DescribeUniBackupPoliciesRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the anti-ransomware policy.
	//
	// example:
	//
	// auto_oracle_Hpm
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s DescribeUniBackupPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupPoliciesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeUniBackupPoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUniBackupPoliciesRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribeUniBackupPoliciesRequest) SetCurrentPage(v int32) *DescribeUniBackupPoliciesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeUniBackupPoliciesRequest) SetPageSize(v int32) *DescribeUniBackupPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUniBackupPoliciesRequest) SetPolicyName(v string) *DescribeUniBackupPoliciesRequest {
	s.PolicyName = &v
	return s
}

func (s *DescribeUniBackupPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
