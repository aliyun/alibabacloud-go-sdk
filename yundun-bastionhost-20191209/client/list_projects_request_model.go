// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListProjectsRequest
	GetInstanceId() *string
	SetKeyword(v string) *ListProjectsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListProjectsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProjectsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListProjectsRequest
	GetRegionId() *string
	SetSkip(v int32) *ListProjectsRequest
	GetSkip() *int32
}

type ListProjectsRequest struct {
	// The ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to retrieve this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the project to query. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of entries to query. If you do not specify this parameter, 20 entries are queried by default. A maximum of 100 entries can be queried in a single request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token required for the next cursor-based query.
	//
	// > Leave this parameter empty for the first query. For subsequent requests, specify the value of the NextToken field returned in the previous response.
	//
	// example:
	//
	// 42bc833a0a0002dae0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of entries to skip.
	//
	// > This parameter takes effect only when the NextToken parameter is left empty.
	//
	// example:
	//
	// 2
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListProjectsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListProjectsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProjectsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListProjectsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListProjectsRequest) SetInstanceId(v string) *ListProjectsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListProjectsRequest) SetKeyword(v string) *ListProjectsRequest {
	s.Keyword = &v
	return s
}

func (s *ListProjectsRequest) SetMaxResults(v int32) *ListProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsRequest) SetNextToken(v string) *ListProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProjectsRequest) SetRegionId(v string) *ListProjectsRequest {
	s.RegionId = &v
	return s
}

func (s *ListProjectsRequest) SetSkip(v int32) *ListProjectsRequest {
	s.Skip = &v
	return s
}

func (s *ListProjectsRequest) Validate() error {
	return dara.Validate(s)
}
