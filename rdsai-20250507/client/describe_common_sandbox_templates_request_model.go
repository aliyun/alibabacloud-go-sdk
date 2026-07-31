// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonSandboxTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeCommonSandboxTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCommonSandboxTemplatesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeCommonSandboxTemplatesRequest
	GetRegionId() *string
}

type DescribeCommonSandboxTemplatesRequest struct {
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// None
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that indicates the position from which the query starts. Set this parameter to empty to start from the beginning.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLYz/NvD85HpgBeRBCusEIeVQ0dHZH9jr+NP3X9Jx0iSoql55b9nd4PIDm252/a0f+U=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCommonSandboxTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonSandboxTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCommonSandboxTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCommonSandboxTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCommonSandboxTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCommonSandboxTemplatesRequest) SetMaxResults(v int32) *DescribeCommonSandboxTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCommonSandboxTemplatesRequest) SetNextToken(v string) *DescribeCommonSandboxTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCommonSandboxTemplatesRequest) SetRegionId(v string) *DescribeCommonSandboxTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCommonSandboxTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
