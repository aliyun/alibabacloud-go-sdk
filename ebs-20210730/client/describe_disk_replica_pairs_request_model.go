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
	// The maximum number of entries to return on each page. Use this parameter with NextToken.
	//
	// Valid values: 1 to 500.
	//
	// Default value: 10.
	//
	// example:
	//
	// 1
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the replication pair. Fuzzy matching is supported.
	//
	// example:
	//
	// name***
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The query token. Set this parameter to the NextToken value returned from the previous call to this operation. You do not need to set this parameter for the first call. If you set NextToken, the PageSize and PageNumber parameters are ignored, and the TotalCount value in the response is invalid.
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
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IDs of replication pairs. Specify one or more replication pair IDs. The IDs must be in the `pair-cn-dsa****,pair-cn-asd****` format.
	//
	// If you leave this parameter empty, all replication pairs in the current region are queried. You can specify up to 100 replication pair IDs.
	//
	// example:
	//
	// pair-cn-dsa****
	PairIds *string `json:"PairIds,omitempty" xml:"PairIds,omitempty"`
	// The ID of the region where the primary or secondary disk of the replication pair resides. Call the [DescribeRegions](https://help.aliyun.com/document_detail/354276.html) operation to query the regions that support asynchronous replication.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair-consistent group. Specify the ID of a replication pair-consistent group to query the replication pairs in the group. The ID must be in the `pg-****` format.
	//
	// If you leave this parameter empty, all replication pairs in the current region are queried.
	//
	// > If you set this parameter to `-`, replication pairs that are not in any replication pair-consistent group are returned.
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
	// The site from which to query data. Query data from the production site or the disaster recovery site. Valid values:
	//
	// - production: the production site.
	//
	// - backup: the disaster recovery site.
	//
	// Default value: production.
	//
	// example:
	//
	// production
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// The tags. You can specify up to 20 tags.
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
