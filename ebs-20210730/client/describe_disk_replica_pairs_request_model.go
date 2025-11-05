// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskReplicaPairsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *DescribeDiskReplicaPairsRequest
	GetMaxResults() *int64
	SetName(v string) *DescribeDiskReplicaPairsRequest
	GetName() *string
	SetNextToken(v string) *DescribeDiskReplicaPairsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeDiskReplicaPairsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDiskReplicaPairsRequest
	GetPageSize() *int32
	SetPairIds(v string) *DescribeDiskReplicaPairsRequest
	GetPairIds() *string
	SetRegionId(v string) *DescribeDiskReplicaPairsRequest
	GetRegionId() *string
	SetReplicaGroupId(v string) *DescribeDiskReplicaPairsRequest
	GetReplicaGroupId() *string
	SetResourceGroupId(v string) *DescribeDiskReplicaPairsRequest
	GetResourceGroupId() *string
	SetSite(v string) *DescribeDiskReplicaPairsRequest
	GetSite() *string
	SetTag(v []*DescribeDiskReplicaPairsRequestTag) *DescribeDiskReplicaPairsRequest
	GetTag() []*DescribeDiskReplicaPairsRequestTag
}

type DescribeDiskReplicaPairsRequest struct {
	// The maximum number of entries per page. You can use this parameter together with NextToken.
	//
	// Valid values: 1 to 500.
	//
	// Default value: 10.
	//
	// example:
	//
	// 1
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the replication pair. Fuzzy search is supported.
	//
	// example:
	//
	// name***
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken. If you specify NextToken, the PageSize and PageNumber request parameters do not take effect, and the TotalCount response parameter is invalid.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 5
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IDs of replication pairs. You can specify the IDs of one or more replication pairs and separate the IDs with commas (,). Example: `pair-cn-dsa****,pair-cn-asd****`.
	//
	// This parameter is empty by default, which indicates that all replication pairs in the specified region are queried. You can specify a maximum of 100 replication pair IDs.
	//
	// example:
	//
	// pair-cn-dsa****
	PairIds *string `json:"PairIds,omitempty" xml:"PairIds,omitempty"`
	// The region ID of the primary or secondary disk in the replication pair. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/354276.html) operation to query the most recent list of regions in which async replication is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair-consistent group. You can specify the ID of a replication pair-consistent group to query the replication pairs in the group. Example: `pg-****`.
	//
	// This parameter is empty by default, which indicates that all replication pairs in the specified region are queried.
	//
	// >  If this parameter is set to`-`, replication pairs that are not added to any replication pair-consistent groups are returned.
	//
	// example:
	//
	// pg-****
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	// The ID of the resource group to which the replication pair belongs.
	//
	// example:
	//
	// rg-acfmvs******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the site from which the information of replication pairs is retrieved. Valid value:
	//
	// 	- production: primary site
	//
	// 	- backup: secondary site
	//
	// Default value: production.
	//
	// example:
	//
	// production
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// The tags. Up to 20 tags are supported.
	Tag []*DescribeDiskReplicaPairsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDiskReplicaPairsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskReplicaPairsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeDiskReplicaPairsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeDiskReplicaPairsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiskReplicaPairsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDiskReplicaPairsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDiskReplicaPairsRequest) GetPairIds() *string {
	return s.PairIds
}

func (s *DescribeDiskReplicaPairsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiskReplicaPairsRequest) GetReplicaGroupId() *string {
	return s.ReplicaGroupId
}

func (s *DescribeDiskReplicaPairsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDiskReplicaPairsRequest) GetSite() *string {
	return s.Site
}

func (s *DescribeDiskReplicaPairsRequest) GetTag() []*DescribeDiskReplicaPairsRequestTag {
	return s.Tag
}

func (s *DescribeDiskReplicaPairsRequest) SetMaxResults(v int64) *DescribeDiskReplicaPairsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetName(v string) *DescribeDiskReplicaPairsRequest {
	s.Name = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetNextToken(v string) *DescribeDiskReplicaPairsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetPageNumber(v int32) *DescribeDiskReplicaPairsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetPageSize(v int32) *DescribeDiskReplicaPairsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetPairIds(v string) *DescribeDiskReplicaPairsRequest {
	s.PairIds = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetRegionId(v string) *DescribeDiskReplicaPairsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetReplicaGroupId(v string) *DescribeDiskReplicaPairsRequest {
	s.ReplicaGroupId = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetResourceGroupId(v string) *DescribeDiskReplicaPairsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetSite(v string) *DescribeDiskReplicaPairsRequest {
	s.Site = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetTag(v []*DescribeDiskReplicaPairsRequestTag) *DescribeDiskReplicaPairsRequest {
	s.Tag = v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) Validate() error {
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

type DescribeDiskReplicaPairsRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDiskReplicaPairsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskReplicaPairsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDiskReplicaPairsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDiskReplicaPairsRequestTag) SetKey(v string) *DescribeDiskReplicaPairsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequestTag) SetValue(v string) *DescribeDiskReplicaPairsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequestTag) Validate() error {
	return dara.Validate(s)
}
