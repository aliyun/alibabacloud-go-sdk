// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomScenePoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeCustomScenePoliciesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCustomScenePoliciesRequest
	GetPageSize() *int32
	SetPolicyId(v int64) *DescribeCustomScenePoliciesRequest
	GetPolicyId() *int64
}

type DescribeCustomScenePoliciesRequest struct {
	// The number of the page to return. Valid values: **1 to 100000**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**. Valid values: **5**, **10**, or **20**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 1234****
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DescribeCustomScenePoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomScenePoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomScenePoliciesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCustomScenePoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCustomScenePoliciesRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DescribeCustomScenePoliciesRequest) SetPageNumber(v int32) *DescribeCustomScenePoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomScenePoliciesRequest) SetPageSize(v int32) *DescribeCustomScenePoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomScenePoliciesRequest) SetPolicyId(v int64) *DescribeCustomScenePoliciesRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribeCustomScenePoliciesRequest) Validate() error {
	return dara.Validate(s)
}
