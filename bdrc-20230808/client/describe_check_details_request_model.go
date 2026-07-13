// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeCheckDetailsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCheckDetailsRequest
	GetNextToken() *string
	SetResourceArn(v string) *DescribeCheckDetailsRequest
	GetResourceArn() *string
	SetRuleId(v string) *DescribeCheckDetailsRequest
	GetRuleId() *string
}

type DescribeCheckDetailsRequest struct {
	// The maximum number of entries to return on each page. The default value is 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. Set this parameter to the value of NextToken that is returned from the last API call. For more information about how to set this parameter, see the API description.
	//
	// example:
	//
	// cae**********699
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The unique identifier of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ecs:123***890:cn-shanghai:instance/i-001***90
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// The unique ID of the data protection rule.
	//
	// example:
	//
	// rule-000***dav
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeCheckDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCheckDetailsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCheckDetailsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCheckDetailsRequest) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *DescribeCheckDetailsRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeCheckDetailsRequest) SetMaxResults(v int32) *DescribeCheckDetailsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCheckDetailsRequest) SetNextToken(v string) *DescribeCheckDetailsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCheckDetailsRequest) SetResourceArn(v string) *DescribeCheckDetailsRequest {
	s.ResourceArn = &v
	return s
}

func (s *DescribeCheckDetailsRequest) SetRuleId(v string) *DescribeCheckDetailsRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeCheckDetailsRequest) Validate() error {
	return dara.Validate(s)
}
