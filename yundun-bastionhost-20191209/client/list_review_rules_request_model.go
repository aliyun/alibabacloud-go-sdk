// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReviewRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListReviewRulesRequest
	GetInstanceId() *string
	SetKeyword(v string) *ListReviewRulesRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListReviewRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListReviewRulesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListReviewRulesRequest
	GetRegionId() *string
	SetSkip(v int32) *ListReviewRulesRequest
	GetSkip() *int32
}

type ListReviewRulesRequest struct {
	// The ID of the Bastionhost instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the audit policy to query. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The maximum number of audit policy entries to return in a single request. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start a paged query. Do not specify this parameter or specify an empty character string for the first request.
	//
	// example:
	//
	// 4ieSWJCwxvW3dk3wF.BqkrZmP72nWu5zJ5NWydMqyEs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the Bastionhost instance. For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of entries to skip. This parameter takes effect only when NextToken is left empty.
	//
	// example:
	//
	// 0
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
}

func (s ListReviewRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListReviewRulesRequest) GoString() string {
	return s.String()
}

func (s *ListReviewRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListReviewRulesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListReviewRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListReviewRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListReviewRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListReviewRulesRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListReviewRulesRequest) SetInstanceId(v string) *ListReviewRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListReviewRulesRequest) SetKeyword(v string) *ListReviewRulesRequest {
	s.Keyword = &v
	return s
}

func (s *ListReviewRulesRequest) SetMaxResults(v int32) *ListReviewRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListReviewRulesRequest) SetNextToken(v string) *ListReviewRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListReviewRulesRequest) SetRegionId(v string) *ListReviewRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListReviewRulesRequest) SetSkip(v int32) *ListReviewRulesRequest {
	s.Skip = &v
	return s
}

func (s *ListReviewRulesRequest) Validate() error {
	return dara.Validate(s)
}
