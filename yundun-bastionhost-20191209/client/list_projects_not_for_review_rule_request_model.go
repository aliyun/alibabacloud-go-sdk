// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsNotForReviewRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListProjectsNotForReviewRuleRequest
	GetInstanceId() *string
	SetKeyword(v string) *ListProjectsNotForReviewRuleRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListProjectsNotForReviewRuleRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProjectsNotForReviewRuleRequest
	GetNextToken() *string
	SetRegionId(v string) *ListProjectsNotForReviewRuleRequest
	GetRegionId() *string
	SetReviewRuleId(v int64) *ListProjectsNotForReviewRuleRequest
	GetReviewRuleId() *int64
	SetSkip(v int32) *ListProjectsNotForReviewRuleRequest
	GetSkip() *int32
}

type ListProjectsNotForReviewRuleRequest struct {
	// The ID of the bastion host instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-7mz2ve7h00a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the project to query. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The maximum number of projects to return per request. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for paging. Leave this parameter empty or do not specify this character for the first request.
	//
	// example:
	//
	// 4ieSWJCwxvW3dk3wF.BqkrZmP72nWu5zJ5NWydMqyEs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the audit policy.
	//
	// > You can invoke the ListReviewRules operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReviewRuleId *int64 `json:"ReviewRuleId,omitempty" xml:"ReviewRuleId,omitempty"`
	// The number of entries to skip.
	//
	// > This parameter takes effect only when NextToken is empty.
	//
	// example:
	//
	// 0
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
}

func (s ListProjectsNotForReviewRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsNotForReviewRuleRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsNotForReviewRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListProjectsNotForReviewRuleRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListProjectsNotForReviewRuleRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProjectsNotForReviewRuleRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsNotForReviewRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListProjectsNotForReviewRuleRequest) GetReviewRuleId() *int64 {
	return s.ReviewRuleId
}

func (s *ListProjectsNotForReviewRuleRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListProjectsNotForReviewRuleRequest) SetInstanceId(v string) *ListProjectsNotForReviewRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ListProjectsNotForReviewRuleRequest) SetKeyword(v string) *ListProjectsNotForReviewRuleRequest {
	s.Keyword = &v
	return s
}

func (s *ListProjectsNotForReviewRuleRequest) SetMaxResults(v int32) *ListProjectsNotForReviewRuleRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsNotForReviewRuleRequest) SetNextToken(v string) *ListProjectsNotForReviewRuleRequest {
	s.NextToken = &v
	return s
}

func (s *ListProjectsNotForReviewRuleRequest) SetRegionId(v string) *ListProjectsNotForReviewRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ListProjectsNotForReviewRuleRequest) SetReviewRuleId(v int64) *ListProjectsNotForReviewRuleRequest {
	s.ReviewRuleId = &v
	return s
}

func (s *ListProjectsNotForReviewRuleRequest) SetSkip(v int32) *ListProjectsNotForReviewRuleRequest {
	s.Skip = &v
	return s
}

func (s *ListProjectsNotForReviewRuleRequest) Validate() error {
	return dara.Validate(s)
}
