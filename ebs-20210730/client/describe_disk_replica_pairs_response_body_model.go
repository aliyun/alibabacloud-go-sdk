// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskReplicaPairsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeDiskReplicaPairsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeDiskReplicaPairsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDiskReplicaPairsResponseBody
	GetPageSize() *int32
	SetReplicaPairs(v []*DescribeDiskReplicaPairsResponseBodyReplicaPairs) *DescribeDiskReplicaPairsResponseBody
	GetReplicaPairs() []*DescribeDiskReplicaPairsResponseBodyReplicaPairs
	SetRequestId(v string) *DescribeDiskReplicaPairsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDiskReplicaPairsResponseBody
	GetTotalCount() *int64
}

type DescribeDiskReplicaPairsResponseBody struct {
	// The query token returned from this call.
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
	// The replication pairs.
	ReplicaPairs []*DescribeDiskReplicaPairsResponseBodyReplicaPairs `json:"ReplicaPairs,omitempty" xml:"ReplicaPairs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// AAA478A0-BEE6-1D42-BEB6-A9CFEAD6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 60
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiskReplicaPairsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskReplicaPairsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiskReplicaPairsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDiskReplicaPairsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDiskReplicaPairsResponseBody) GetReplicaPairs() []*DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	return s.ReplicaPairs
}

func (s *DescribeDiskReplicaPairsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiskReplicaPairsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDiskReplicaPairsResponseBody) SetNextToken(v string) *DescribeDiskReplicaPairsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetPageNumber(v int32) *DescribeDiskReplicaPairsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetPageSize(v int32) *DescribeDiskReplicaPairsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetReplicaPairs(v []*DescribeDiskReplicaPairsResponseBodyReplicaPairs) *DescribeDiskReplicaPairsResponseBody {
	s.ReplicaPairs = v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetRequestId(v string) *DescribeDiskReplicaPairsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetTotalCount(v int64) *DescribeDiskReplicaPairsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) Validate() error {
	if s.ReplicaPairs != nil {
		for _, item := range s.ReplicaPairs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiskReplicaPairsResponseBodyReplicaPairs struct {
	// The bandwidth used for asynchronous replication. Unit: Kbit/s.
	//
	// example:
	//
	// 10240
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The billing method of the replication pair.
	//
	// Valid values:
	//
	// - PREPAY: subscription.
	//
	// - POSTPAY: pay-as-you-go.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The creation time. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1649750977
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the replication pair.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the secondary disk.
	//
	// example:
	//
	// d-asdfjl2342kj2l3k4****
	DestinationDiskId *string `json:"DestinationDiskId,omitempty" xml:"DestinationDiskId,omitempty"`
	// The region of the secondary disk.
	//
	// example:
	//
	// cn-shanghai
	DestinationRegion *string `json:"DestinationRegion,omitempty" xml:"DestinationRegion,omitempty"`
	// The zone of the secondary disk.
	//
	// example:
	//
	// cn-shanghai-b
	DestinationZoneId *string `json:"DestinationZoneId,omitempty" xml:"DestinationZoneId,omitempty"`
	// Specifies whether real-time control (RTC) is enabled. Valid values:
	//
	// - false: Disabled.
	//
	// - true: Enabled.
	//
	// > If the replication pair is in a replication pair-consistent group, the value of this parameter is the same as that of the group.
	EnableRtc *bool `json:"EnableRtc,omitempty" xml:"EnableRtc,omitempty"`
	// The expiration time of the replication pair. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1649750977
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The time when the last asynchronous replication was completed. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1649751977
	LastRecoverPoint *int64 `json:"LastRecoverPoint,omitempty" xml:"LastRecoverPoint,omitempty"`
	// The name of the replication pair.
	//
	// example:
	//
	// TestReplicaPair
	PairName *string `json:"PairName,omitempty" xml:"PairName,omitempty"`
	// The initial source region of the replication pair.
	//
	// example:
	//
	// cn-beijing
	PrimaryRegion *string `json:"PrimaryRegion,omitempty" xml:"PrimaryRegion,omitempty"`
	// The initial source zone of the replication pair.
	//
	// example:
	//
	// cn-beijing-a
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// The recovery point objective (RPO) of the replication pair. Unit: seconds.
	//
	// example:
	//
	// 900
	RPO *int64 `json:"RPO,omitempty" xml:"RPO,omitempty"`
	// The ID of the replication pair-consistent group to which the replication pair belongs.
	//
	// example:
	//
	// pg-xxxx****
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	// The name of the replication pair-consistent group to which the replication pair belongs.
	//
	// example:
	//
	// pg-name****
	ReplicaGroupName *string `json:"ReplicaGroupName,omitempty" xml:"ReplicaGroupName,omitempty"`
	// The ID of the replication pair.
	//
	// example:
	//
	// pair-cn-dsa****
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
	// The ID of the resource group to which the replication pair belongs.
	//
	// example:
	//
	// rg-acfmvs*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The site type of the replication pair or replication pair-consistent group. Valid values:
	//
	// - production: the production site.
	//
	// - backup: the disaster recovery site.
	//
	// example:
	//
	// production
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// The ID of the primary disk.
	//
	// example:
	//
	// d-bp131n0q38u3a4zi****
	SourceDiskId *string `json:"SourceDiskId,omitempty" xml:"SourceDiskId,omitempty"`
	// The region of the primary disk.
	//
	// example:
	//
	// cn-beijing
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// The zone of the primary disk.
	//
	// example:
	//
	// cn-beijing-a
	SourceZoneId *string `json:"SourceZoneId,omitempty" xml:"SourceZoneId,omitempty"`
	// The initial destination region of the replication pair.
	//
	// example:
	//
	// cn-shanghai
	StandbyRegion *string `json:"StandbyRegion,omitempty" xml:"StandbyRegion,omitempty"`
	// The initial destination zone of the replication pair.
	//
	// example:
	//
	// cn-shanghai-b
	StandbyZone *string `json:"StandbyZone,omitempty" xml:"StandbyZone,omitempty"`
	// The status of the replication pair. Valid values:
	//
	// - invalid: The replication pair is invalid. This status indicates that the replication pair is not working correctly.
	//
	// - creating: The replication pair is being created.
	//
	// - created: The replication pair is created.
	//
	// - create_failed: The replication pair failed to be created.
	//
	// - initial_syncing: The replication pair is in the initial synchronization state. After a replication pair is created and started, it enters this state during the first asynchronous replication of data from the primary disk to the secondary disk.
	//
	// - manual_syncing: The replication pair is being manually synchronized. After the manual synchronization is complete, the replication pair returns to the stopped state. If it is the first one-time synchronization, the status is also manual_syncing.
	//
	// - syncing: The replication pair is synchronizing data. The replication pair is in this state when data is asynchronously replicated from the primary disk to the secondary disk for a second or subsequent time.
	//
	// - normal: The replication pair is in the normal state. The replication pair enters this state when data replication is complete in the current replication cycle.
	//
	// - stopping: The replication pair is being stopped.
	//
	// - stopped: The replication pair is stopped.
	//
	// - stop_failed: The replication pair failed to be stopped.
	//
	// - failovering: A failover is in progress.
	//
	// - failovered: The failover is complete.
	//
	// - failover_failed: The failover failed.
	//
	// - reprotecting: A reverse replication is in progress.
	//
	// - reprotect_failed: The reverse replication failed.
	//
	// - deleting: The replication pair is being deleted.
	//
	// - delete_failed: The replication pair failed to be deleted.
	//
	// - deleted: The replication pair is deleted.
	//
	// example:
	//
	// created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status message of the replication pair. This parameter is returned when the Status is `invalid` or `create_failed`. Valid values:
	//
	// - PrePayOrderExpired: The subscription replication pair has expired.
	//
	// - PostPayOrderCeaseService: The service for the pay-as-you-go replication pair is suspended, usually due to an overdue payment.
	//
	// - DeviceRemoved: The primary or secondary disk is deleted.
	//
	// - DeviceKeyChanged: The `DeviceKey` mapping of the primary or secondary disk has changed.
	//
	// - DeviceSizeChanged: The `DeviceSize` of the primary or secondary disk has changed.
	//
	// - OperationDenied.QuotaExceed: The number of created replication pairs exceeds the quota.
	//
	// example:
	//
	// PrePayOrderExpired
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The tags of the replication pair.
	Tags []*DescribeDiskReplicaPairsResponseBodyReplicaPairsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeDiskReplicaPairsResponseBodyReplicaPairs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskReplicaPairsResponseBodyReplicaPairs) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetDescription() *string {
	return s.Description
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetDestinationDiskId() *string {
	return s.DestinationDiskId
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetDestinationRegion() *string {
	return s.DestinationRegion
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetDestinationZoneId() *string {
	return s.DestinationZoneId
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetEnableRtc() *bool {
	return s.EnableRtc
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetExpiredTime() *int64 {
	return s.ExpiredTime
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetLastRecoverPoint() *int64 {
	return s.LastRecoverPoint
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetPairName() *string {
	return s.PairName
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetPrimaryRegion() *string {
	return s.PrimaryRegion
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetPrimaryZone() *string {
	return s.PrimaryZone
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetRPO() *int64 {
	return s.RPO
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetReplicaGroupId() *string {
	return s.ReplicaGroupId
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetReplicaGroupName() *string {
	return s.ReplicaGroupName
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetReplicaPairId() *string {
	return s.ReplicaPairId
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetSite() *string {
	return s.Site
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetSourceDiskId() *string {
	return s.SourceDiskId
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetSourceRegion() *string {
	return s.SourceRegion
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetSourceZoneId() *string {
	return s.SourceZoneId
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetStandbyRegion() *string {
	return s.StandbyRegion
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetStandbyZone() *string {
	return s.StandbyZone
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetStatus() *string {
	return s.Status
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) GetTags() []*DescribeDiskReplicaPairsResponseBodyReplicaPairsTags {
	return s.Tags
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetBandwidth(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Bandwidth = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetChargeType(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ChargeType = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetCreateTime(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.CreateTime = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetDescription(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Description = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetDestinationDiskId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.DestinationDiskId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetDestinationRegion(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.DestinationRegion = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetDestinationZoneId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.DestinationZoneId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetEnableRtc(v bool) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.EnableRtc = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetExpiredTime(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetLastRecoverPoint(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.LastRecoverPoint = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetPairName(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.PairName = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetPrimaryRegion(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.PrimaryRegion = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetPrimaryZone(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.PrimaryZone = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetRPO(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.RPO = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetReplicaGroupId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ReplicaGroupId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetReplicaGroupName(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ReplicaGroupName = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetReplicaPairId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ReplicaPairId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetResourceGroupId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetSite(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Site = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetSourceDiskId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.SourceDiskId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetSourceRegion(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.SourceRegion = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetSourceZoneId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.SourceZoneId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetStandbyRegion(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.StandbyRegion = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetStandbyZone(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.StandbyZone = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetStatus(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Status = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetStatusMessage(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.StatusMessage = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetTags(v []*DescribeDiskReplicaPairsResponseBodyReplicaPairsTags) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Tags = v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) Validate() error {
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

type DescribeDiskReplicaPairsResponseBodyReplicaPairsTags struct {
	// The key of the tag.
	//
	// example:
	//
	// testKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// testValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDiskReplicaPairsResponseBodyReplicaPairsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskReplicaPairsResponseBodyReplicaPairsTags) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairsTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairsTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairsTags) SetTagKey(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairsTags {
	s.TagKey = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairsTags) SetTagValue(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairsTags {
	s.TagValue = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairsTags) Validate() error {
	return dara.Validate(s)
}
