// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMembersNotForProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListMembersNotForProjectRequest
	GetInstanceId() *string
	SetKeyword(v string) *ListMembersNotForProjectRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListMembersNotForProjectRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMembersNotForProjectRequest
	GetNextToken() *string
	SetProjectId(v int64) *ListMembersNotForProjectRequest
	GetProjectId() *int64
	SetRegionId(v string) *ListMembersNotForProjectRequest
	GetRegionId() *string
	SetSkip(v int32) *ListMembersNotForProjectRequest
	GetSkip() *int32
}

type ListMembersNotForProjectRequest struct {
	// The instance ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-i7m2d7zrw11
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The UIDs of the member accounts to query. Fuzzy match is supported.
	//
	// example:
	//
	// 162****5213
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The maximum number of entries to return. If you do not specify this parameter, the default value 20 is used. You can query up to 100 entries at a time.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. For the first request, leave this parameter empty or pass an empty string.
	//
	// example:
	//
	// 4ieSWJCwxvW3dk3wF.BqkrZmP72nWu5zJ5NWydMqyEs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The project ID.
	//
	// > You can call the ListProjects operation to obtain this parameter.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
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
	// > This parameter takes effect only when NextToken is empty.
	//
	// example:
	//
	// 0
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
}

func (s ListMembersNotForProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMembersNotForProjectRequest) GoString() string {
	return s.String()
}

func (s *ListMembersNotForProjectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListMembersNotForProjectRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListMembersNotForProjectRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMembersNotForProjectRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMembersNotForProjectRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListMembersNotForProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListMembersNotForProjectRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListMembersNotForProjectRequest) SetInstanceId(v string) *ListMembersNotForProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *ListMembersNotForProjectRequest) SetKeyword(v string) *ListMembersNotForProjectRequest {
	s.Keyword = &v
	return s
}

func (s *ListMembersNotForProjectRequest) SetMaxResults(v int32) *ListMembersNotForProjectRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMembersNotForProjectRequest) SetNextToken(v string) *ListMembersNotForProjectRequest {
	s.NextToken = &v
	return s
}

func (s *ListMembersNotForProjectRequest) SetProjectId(v int64) *ListMembersNotForProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *ListMembersNotForProjectRequest) SetRegionId(v string) *ListMembersNotForProjectRequest {
	s.RegionId = &v
	return s
}

func (s *ListMembersNotForProjectRequest) SetSkip(v int32) *ListMembersNotForProjectRequest {
	s.Skip = &v
	return s
}

func (s *ListMembersNotForProjectRequest) Validate() error {
	return dara.Validate(s)
}
