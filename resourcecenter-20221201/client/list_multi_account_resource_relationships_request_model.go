// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountResourceRelationshipsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMultiAccountResourceRelationshipsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMultiAccountResourceRelationshipsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListMultiAccountResourceRelationshipsRequest
	GetRegionId() *string
	SetRelatedResourceFilter(v []*ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) *ListMultiAccountResourceRelationshipsRequest
	GetRelatedResourceFilter() []*ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter
	SetResourceId(v string) *ListMultiAccountResourceRelationshipsRequest
	GetResourceId() *string
	SetResourceType(v string) *ListMultiAccountResourceRelationshipsRequest
	GetResourceType() *string
	SetScope(v string) *ListMultiAccountResourceRelationshipsRequest
	GetScope() *string
}

type ListMultiAccountResourceRelationshipsRequest struct {
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 500.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The filter conditions for resources associated with the resource.
	RelatedResourceFilter []*ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter `json:"RelatedResourceFilter,omitempty" xml:"RelatedResourceFilter,omitempty" type:"Repeated"`
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-eb3hji****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::ACK::Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The search scope. Valid values:
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched. You can call the [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html) operation to query the ID.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to query the ID.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to query the ID.
	//
	// 	- ID of a member: Resources within the member are searched. You can call the [ListAccounts](https://help.aliyun.com/document_detail/160016.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ListMultiAccountResourceRelationshipsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountResourceRelationshipsRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceRelationshipsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMultiAccountResourceRelationshipsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiAccountResourceRelationshipsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListMultiAccountResourceRelationshipsRequest) GetRelatedResourceFilter() []*ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter {
	return s.RelatedResourceFilter
}

func (s *ListMultiAccountResourceRelationshipsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListMultiAccountResourceRelationshipsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListMultiAccountResourceRelationshipsRequest) GetScope() *string {
	return s.Scope
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetMaxResults(v int32) *ListMultiAccountResourceRelationshipsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetNextToken(v string) *ListMultiAccountResourceRelationshipsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetRegionId(v string) *ListMultiAccountResourceRelationshipsRequest {
	s.RegionId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetRelatedResourceFilter(v []*ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) *ListMultiAccountResourceRelationshipsRequest {
	s.RelatedResourceFilter = v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetResourceId(v string) *ListMultiAccountResourceRelationshipsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetResourceType(v string) *ListMultiAccountResourceRelationshipsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) SetScope(v string) *ListMultiAccountResourceRelationshipsRequest {
	s.Scope = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequest) Validate() error {
	return dara.Validate(s)
}

type ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter struct {
	// The key of the filter condition. For more information, see `Supported filter parameters`.
	//
	// example:
	//
	// RelatedResourceRegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching method.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) GetKey() *string {
	return s.Key
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) GetMatchType() *string {
	return s.MatchType
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) GetValue() []*string {
	return s.Value
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) SetKey(v string) *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter {
	s.Key = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) SetMatchType(v string) *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter {
	s.MatchType = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) SetValue(v []*string) *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter {
	s.Value = v
	return s
}

func (s *ListMultiAccountResourceRelationshipsRequestRelatedResourceFilter) Validate() error {
	return dara.Validate(s)
}
