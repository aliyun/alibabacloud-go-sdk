// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskReplicaGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupIds(v string) *DescribeDiskReplicaGroupsRequest
	GetGroupIds() *string
	SetMaxResults(v int64) *DescribeDiskReplicaGroupsRequest
	GetMaxResults() *int64
	SetName(v string) *DescribeDiskReplicaGroupsRequest
	GetName() *string
	SetNextToken(v string) *DescribeDiskReplicaGroupsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeDiskReplicaGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDiskReplicaGroupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDiskReplicaGroupsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDiskReplicaGroupsRequest
	GetResourceGroupId() *string
	SetSite(v string) *DescribeDiskReplicaGroupsRequest
	GetSite() *string
	SetTag(v []*DescribeDiskReplicaGroupsRequestTag) *DescribeDiskReplicaGroupsRequest
	GetTag() []*DescribeDiskReplicaGroupsRequestTag
}

type DescribeDiskReplicaGroupsRequest struct {
	// The IDs of the replication pair-consistent groups. You can specify the IDs of one or more replication pair-consistent groups. Separate the IDs with commas (,).
	//
	// This parameter is empty by default, which indicates that all replication pair-consistent groups in the specified region are queried. You can specify up to the IDs of 100 replication pair-consistent groups.
	//
	// example:
	//
	// AAAAAdDWBF2****
	GroupIds *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// The maximum number of entries per page. You can use this parameter together with NextToken.
	//
	// Valid values: 1 to 500.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the replication pair-consistent group. You can perform a fuzzy search.
	//
	// example:
	//
	// pg-name***
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken. If you specify NextToken, the PageSize and PageNumber request parameters do not take effect, and the TotalCount response parameter is invalid.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 5
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region to which the replication pair-consistent group belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the replication pair-consistent group belongs.
	//
	// example:
	//
	// rg-aekz*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the site from which the information of replication pair-consistent groups is retrieved. This parameter is used for scenarios where data is replicated across zones in replication pairs.
	//
	// 	- If this parameter is not specified, information such as the status of replication pair-consistent groups at the primary site is queried and returned.
	//
	// 	- Otherwise, information such as the state of replication pairs at the site specified by the Site parameter is queried and returned. Valid values:
	//
	//     	- production: primary site
	//
	//     	- backup: secondary site
	//
	// example:
	//
	// production
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// The tags to add to the replication pair-consistent group. You can specify up to 20 tags.
	Tag []*DescribeDiskReplicaGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDiskReplicaGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskReplicaGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaGroupsRequest) GetGroupIds() *string {
	return s.GroupIds
}

func (s *DescribeDiskReplicaGroupsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeDiskReplicaGroupsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeDiskReplicaGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiskReplicaGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDiskReplicaGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDiskReplicaGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiskReplicaGroupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDiskReplicaGroupsRequest) GetSite() *string {
	return s.Site
}

func (s *DescribeDiskReplicaGroupsRequest) GetTag() []*DescribeDiskReplicaGroupsRequestTag {
	return s.Tag
}

func (s *DescribeDiskReplicaGroupsRequest) SetGroupIds(v string) *DescribeDiskReplicaGroupsRequest {
	s.GroupIds = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetMaxResults(v int64) *DescribeDiskReplicaGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetName(v string) *DescribeDiskReplicaGroupsRequest {
	s.Name = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetNextToken(v string) *DescribeDiskReplicaGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetPageNumber(v int32) *DescribeDiskReplicaGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetPageSize(v int32) *DescribeDiskReplicaGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetRegionId(v string) *DescribeDiskReplicaGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetResourceGroupId(v string) *DescribeDiskReplicaGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetSite(v string) *DescribeDiskReplicaGroupsRequest {
	s.Site = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetTag(v []*DescribeDiskReplicaGroupsRequestTag) *DescribeDiskReplicaGroupsRequest {
	s.Tag = v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiskReplicaGroupsRequestTag struct {
	// The key of tag N of the replication pair-consistent group.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the replication pair-consistent group.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDiskReplicaGroupsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskReplicaGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaGroupsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDiskReplicaGroupsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDiskReplicaGroupsRequestTag) SetKey(v string) *DescribeDiskReplicaGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequestTag) SetValue(v string) *DescribeDiskReplicaGroupsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequestTag) Validate() error {
	return dara.Validate(s)
}
