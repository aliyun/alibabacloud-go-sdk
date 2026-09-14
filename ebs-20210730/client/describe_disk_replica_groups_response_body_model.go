// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskReplicaGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeDiskReplicaGroupsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeDiskReplicaGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDiskReplicaGroupsResponseBody
	GetPageSize() *int32
	SetReplicaGroups(v []*DescribeDiskReplicaGroupsResponseBodyReplicaGroups) *DescribeDiskReplicaGroupsResponseBody
	GetReplicaGroups() []*DescribeDiskReplicaGroupsResponseBodyReplicaGroups
	SetRequestId(v string) *DescribeDiskReplicaGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDiskReplicaGroupsResponseBody
	GetTotalCount() *int64
}

type DescribeDiskReplicaGroupsResponseBody struct {
	// The query token returned in this call.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about the replication pair-consistent groups.
	ReplicaGroups []*DescribeDiskReplicaGroupsResponseBodyReplicaGroups `json:"ReplicaGroups,omitempty" xml:"ReplicaGroups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// AAA478A0-BEE6-1D42-BEB6-A9CFEAD6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 60
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiskReplicaGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskReplicaGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiskReplicaGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDiskReplicaGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDiskReplicaGroupsResponseBody) GetReplicaGroups() []*DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	return s.ReplicaGroups
}

func (s *DescribeDiskReplicaGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiskReplicaGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDiskReplicaGroupsResponseBody) SetNextToken(v string) *DescribeDiskReplicaGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBody) SetPageNumber(v int32) *DescribeDiskReplicaGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBody) SetPageSize(v int32) *DescribeDiskReplicaGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBody) SetReplicaGroups(v []*DescribeDiskReplicaGroupsResponseBodyReplicaGroups) *DescribeDiskReplicaGroupsResponseBody {
	s.ReplicaGroups = v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBody) SetRequestId(v string) *DescribeDiskReplicaGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBody) SetTotalCount(v int64) *DescribeDiskReplicaGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBody) Validate() error {
	if s.ReplicaGroups != nil {
		for _, item := range s.ReplicaGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiskReplicaGroupsResponseBodyReplicaGroups struct {
	// The bandwidth. Unit: Kbit/s. This parameter is not yet available. The return value is preset by the system.
	//
	// example:
	//
	// 0
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The description of the replication pair-consistent group.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID of the disaster recovery site.
	//
	// example:
	//
	// cn-shanghai
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// The zone ID of the disaster recovery site.
	//
	// example:
	//
	// cn-shanghai-e
	DestinationZoneId *string `json:"DestinationZoneId,omitempty" xml:"DestinationZoneId,omitempty"`
	// Specifies whether to enable replication time control (RTC). Valid values:
	//
	// - false: Disables RTC.
	//
	// - true: Enables RTC.
	//
	// > If you set this parameter to true, RTC is enabled for the replication pair-consistent group and all asynchronous replication pairs that are added to the group.
	//
	// example:
	//
	// true
	EnableRtc *bool `json:"EnableRtc,omitempty" xml:"EnableRtc,omitempty"`
	// The name of the replication pair-consistent group.
	//
	// example:
	//
	// myreplicagrouptest
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The time when the last asynchronous replication was completed for the replication pair-consistent group. This parameter is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1637835114
	LastRecoverPoint *int64 `json:"LastRecoverPoint,omitempty" xml:"LastRecoverPoint,omitempty"`
	// The list of replication pair IDs in the replication pair-consistent group.
	PairIds [][]byte `json:"PairIds,omitempty" xml:"PairIds,omitempty" type:"Repeated"`
	// The number of replication pairs in the replication pair-consistent group.
	//
	// example:
	//
	// 2
	PairNumber *int64 `json:"PairNumber,omitempty" xml:"PairNumber,omitempty"`
	// The initial source region of the replication group.
	//
	// example:
	//
	// cn-beijing
	PrimaryRegion *string `json:"PrimaryRegion,omitempty" xml:"PrimaryRegion,omitempty"`
	// The initial source zone of the replication group.
	//
	// example:
	//
	// cn-beijing-h
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// The recovery point objective (RPO) of the replication pair-consistent group. Unit: seconds.
	//
	// example:
	//
	// 180
	RPO *int64 `json:"RPO,omitempty" xml:"RPO,omitempty"`
	// The ID of the replication pair-consistent group.
	//
	// example:
	//
	// pg-myreplica****
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	// The ID of the resource group to which the replication group belongs.
	//
	// example:
	//
	// rg-aek2a*******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The site of the replication pair and the replication pair-consistent group. Valid values:
	//
	// - production: The production site.
	//
	// - backup: The disaster recovery site.
	//
	// example:
	//
	// production
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// The region ID of the production site.
	//
	// example:
	//
	// cn-beijing
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// The zone ID of the production site.
	//
	// example:
	//
	// cn-beijing-f
	SourceZoneId *string `json:"SourceZoneId,omitempty" xml:"SourceZoneId,omitempty"`
	// The initial destination region of the replication group.
	//
	// example:
	//
	// cn-shanghai
	StandbyRegion *string `json:"StandbyRegion,omitempty" xml:"StandbyRegion,omitempty"`
	// The initial destination zone of the replication group.
	//
	// example:
	//
	// cn-shanghai-e
	StandbyZone *string `json:"StandbyZone,omitempty" xml:"StandbyZone,omitempty"`
	// The status of the replication pair-consistent group. Valid values:
	//
	// - invalid: The replication pair-consistent group is invalid. This status indicates that a replication pair in the group is abnormal.
	//
	// - creating: The replication pair-consistent group is being created.
	//
	// - created: The replication pair-consistent group is created.
	//
	// - create_failed: The replication pair-consistent group failed to be created.
	//
	// - manual_syncing: The replication pair-consistent group is performing a one-time synchronization. The group is also in this state during the first one-time synchronization.
	//
	// - syncing: The replication pair-consistent group is synchronizing data. The group is in this state when data is asynchronously replicated from the primary disk to the secondary disk for a subsequent time.
	//
	// - normal: Normal. When data replication is complete in the current asynchronous replication cycle, the group is in this state.
	//
	// - stopping: The replication pair-consistent group is being stopped.
	//
	// - stopped: The replication pair-consistent group is stopped.
	//
	// - stop_failed: The replication pair-consistent group failed to be stopped.
	//
	// - failovering: A failover is being performed.
	//
	// - failovered: The failover is complete.
	//
	// - failover_failed: The failover failed.
	//
	// - reprotecting: A reverse replication is being performed.
	//
	// - reprotect_failed: The reverse replication failed.
	//
	// - deleting: The replication pair-consistent group is being deleted.
	//
	// - delete_failed: The replication pair-consistent group failed to be deleted.
	//
	// - deleted: The replication pair-consistent group is deleted.
	//
	// example:
	//
	// created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the replication group.
	Tags []*DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeDiskReplicaGroupsResponseBodyReplicaGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetDescription() *string {
	return s.Description
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetDestinationRegionId() *string {
	return s.DestinationRegionId
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetDestinationZoneId() *string {
	return s.DestinationZoneId
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetEnableRtc() *bool {
	return s.EnableRtc
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetLastRecoverPoint() *int64 {
	return s.LastRecoverPoint
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetPairIds() [][]byte {
	return s.PairIds
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetPairNumber() *int64 {
	return s.PairNumber
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetPrimaryRegion() *string {
	return s.PrimaryRegion
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetPrimaryZone() *string {
	return s.PrimaryZone
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetRPO() *int64 {
	return s.RPO
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetReplicaGroupId() *string {
	return s.ReplicaGroupId
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetSite() *string {
	return s.Site
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetSourceRegionId() *string {
	return s.SourceRegionId
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetSourceZoneId() *string {
	return s.SourceZoneId
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetStandbyRegion() *string {
	return s.StandbyRegion
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetStandbyZone() *string {
	return s.StandbyZone
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetStatus() *string {
	return s.Status
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GetTags() []*DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags {
	return s.Tags
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetBandwidth(v int64) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.Bandwidth = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetDescription(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.Description = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetDestinationRegionId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.DestinationRegionId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetDestinationZoneId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.DestinationZoneId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetEnableRtc(v bool) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.EnableRtc = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetGroupName(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetLastRecoverPoint(v int64) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.LastRecoverPoint = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetPairIds(v [][]byte) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.PairIds = v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetPairNumber(v int64) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.PairNumber = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetPrimaryRegion(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.PrimaryRegion = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetPrimaryZone(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.PrimaryZone = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetRPO(v int64) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.RPO = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetReplicaGroupId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.ReplicaGroupId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetResourceGroupId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetSite(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.Site = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetSourceRegionId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.SourceRegionId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetSourceZoneId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.SourceZoneId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetStandbyRegion(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.StandbyRegion = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetStandbyZone(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.StandbyZone = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetStatus(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.Status = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetTags(v []*DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.Tags = v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags struct {
	// The key of the tag of the replication group.
	//
	// example:
	//
	// testKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag of the replication group.
	//
	// example:
	//
	// testValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags) SetTagKey(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags {
	s.TagKey = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags) SetTagValue(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags {
	s.TagValue = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroupsTags) Validate() error {
	return dara.Validate(s)
}
