// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsNotForProjectAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationId(v int64) *ListProjectsNotForProjectAuthorizationRequest
	GetAuthorizationId() *int64
	SetInstanceId(v string) *ListProjectsNotForProjectAuthorizationRequest
	GetInstanceId() *string
	SetKeyword(v string) *ListProjectsNotForProjectAuthorizationRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListProjectsNotForProjectAuthorizationRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProjectsNotForProjectAuthorizationRequest
	GetNextToken() *string
	SetRegionId(v string) *ListProjectsNotForProjectAuthorizationRequest
	GetRegionId() *string
	SetSkip(v int32) *ListProjectsNotForProjectAuthorizationRequest
	GetSkip() *int32
}

type ListProjectsNotForProjectAuthorizationRequest struct {
	// The project authorization ID.
	//
	// > You can call the ListProjectAuthorizations operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuthorizationId *int64 `json:"AuthorizationId,omitempty" xml:"AuthorizationId,omitempty"`
	// The ID of the bastion host instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the bastion host instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost_std_intl-sg-uq833e2dz02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project name to query. Fuzzy match is supported.
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
	// The paging token. Leave this parameter empty or do not specify this parameter for the first request.
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
	// The number of entries to skip.
	//
	// > This parameter takes effect only when NextToken is empty.
	//
	// example:
	//
	// 0
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
}

func (s ListProjectsNotForProjectAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsNotForProjectAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsNotForProjectAuthorizationRequest) GetAuthorizationId() *int64 {
	return s.AuthorizationId
}

func (s *ListProjectsNotForProjectAuthorizationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListProjectsNotForProjectAuthorizationRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListProjectsNotForProjectAuthorizationRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProjectsNotForProjectAuthorizationRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsNotForProjectAuthorizationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListProjectsNotForProjectAuthorizationRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListProjectsNotForProjectAuthorizationRequest) SetAuthorizationId(v int64) *ListProjectsNotForProjectAuthorizationRequest {
	s.AuthorizationId = &v
	return s
}

func (s *ListProjectsNotForProjectAuthorizationRequest) SetInstanceId(v string) *ListProjectsNotForProjectAuthorizationRequest {
	s.InstanceId = &v
	return s
}

func (s *ListProjectsNotForProjectAuthorizationRequest) SetKeyword(v string) *ListProjectsNotForProjectAuthorizationRequest {
	s.Keyword = &v
	return s
}

func (s *ListProjectsNotForProjectAuthorizationRequest) SetMaxResults(v int32) *ListProjectsNotForProjectAuthorizationRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsNotForProjectAuthorizationRequest) SetNextToken(v string) *ListProjectsNotForProjectAuthorizationRequest {
	s.NextToken = &v
	return s
}

func (s *ListProjectsNotForProjectAuthorizationRequest) SetRegionId(v string) *ListProjectsNotForProjectAuthorizationRequest {
	s.RegionId = &v
	return s
}

func (s *ListProjectsNotForProjectAuthorizationRequest) SetSkip(v int32) *ListProjectsNotForProjectAuthorizationRequest {
	s.Skip = &v
	return s
}

func (s *ListProjectsNotForProjectAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
