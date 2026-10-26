// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsForProjectAuthorizationPrincipalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListProjectsForProjectAuthorizationPrincipalRequest
	GetInstanceId() *string
	SetKeyword(v string) *ListProjectsForProjectAuthorizationPrincipalRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListProjectsForProjectAuthorizationPrincipalRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProjectsForProjectAuthorizationPrincipalRequest
	GetNextToken() *string
	SetRegionId(v string) *ListProjectsForProjectAuthorizationPrincipalRequest
	GetRegionId() *string
	SetSkip(v int32) *ListProjectsForProjectAuthorizationPrincipalRequest
	GetSkip() *int32
}

type ListProjectsForProjectAuthorizationPrincipalRequest struct {
	// The instance ID of the Bastionhost instance.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the parameter value.
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
	// name
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The maximum number of projects to return per request. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The paging token used to start a paged query. For the first request, do not pass this parameter or pass an empty string of characters.
	//
	// example:
	//
	// 4ieSWJCwxvW3dk3wF.BqkrZmP72nWu5zJ5NWydMqyEs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of items to skip.
	//
	// > This parameter takes effect only when NextToken is left empty.
	//
	// example:
	//
	// 0
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
}

func (s ListProjectsForProjectAuthorizationPrincipalRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsForProjectAuthorizationPrincipalRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsForProjectAuthorizationPrincipalRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListProjectsForProjectAuthorizationPrincipalRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListProjectsForProjectAuthorizationPrincipalRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProjectsForProjectAuthorizationPrincipalRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsForProjectAuthorizationPrincipalRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListProjectsForProjectAuthorizationPrincipalRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListProjectsForProjectAuthorizationPrincipalRequest) SetInstanceId(v string) *ListProjectsForProjectAuthorizationPrincipalRequest {
	s.InstanceId = &v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalRequest) SetKeyword(v string) *ListProjectsForProjectAuthorizationPrincipalRequest {
	s.Keyword = &v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalRequest) SetMaxResults(v int32) *ListProjectsForProjectAuthorizationPrincipalRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalRequest) SetNextToken(v string) *ListProjectsForProjectAuthorizationPrincipalRequest {
	s.NextToken = &v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalRequest) SetRegionId(v string) *ListProjectsForProjectAuthorizationPrincipalRequest {
	s.RegionId = &v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalRequest) SetSkip(v int32) *ListProjectsForProjectAuthorizationPrincipalRequest {
	s.Skip = &v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalRequest) Validate() error {
	return dara.Validate(s)
}
