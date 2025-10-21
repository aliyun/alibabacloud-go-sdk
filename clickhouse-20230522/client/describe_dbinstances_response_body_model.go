// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDBInstancesResponseBodyData) *DescribeDBInstancesResponseBody
	GetData() *DescribeDBInstancesResponseBodyData
	SetRequestId(v string) *DescribeDBInstancesResponseBody
	GetRequestId() *string
}

type DescribeDBInstancesResponseBody struct {
	// The returned result.
	Data *DescribeDBInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBody) GetData() *DescribeDBInstancesResponseBodyData {
	return s.Data
}

func (s *DescribeDBInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstancesResponseBody) SetData(v *DescribeDBInstancesResponseBodyData) *DescribeDBInstancesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetRequestId(v string) *DescribeDBInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstancesResponseBodyData struct {
	// The clusters.
	DBInstances []*DescribeDBInstancesResponseBodyDataDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
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
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyData) GetDBInstances() []*DescribeDBInstancesResponseBodyDataDBInstances {
	return s.DBInstances
}

func (s *DescribeDBInstancesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstancesResponseBodyData) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeDBInstancesResponseBodyData) SetDBInstances(v []*DescribeDBInstancesResponseBodyDataDBInstances) *DescribeDBInstancesResponseBodyData {
	s.DBInstances = v
	return s
}

func (s *DescribeDBInstancesResponseBodyData) SetPageNumber(v int32) *DescribeDBInstancesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyData) SetPageSize(v int32) *DescribeDBInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyData) SetTotalCount(v string) *DescribeDBInstancesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyData) Validate() error {
	if s.DBInstances != nil {
		for _, item := range s.DBInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancesResponseBodyDataDBInstances struct {
	// The user ID.
	//
	// example:
	//
	// 1294****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The channel ID.
	//
	// example:
	//
	// 186681****
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PrePaid: subscription
	//
	// 	- PostPaid: pay-as-you-go
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2022-12-04 21:16:15
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Indicates whether the release protection feature is enabled for the cluster.
	//
	// example:
	//
	// False
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// test_desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The engine type.
	//
	// example:
	//
	// clickhouse
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine version.
	//
	// example:
	//
	// 22.8
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the cluster expires.
	//
	// example:
	//
	// 2024-02-16 11:51:06
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The lock mode.
	//
	// example:
	//
	// 0
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the cluster was locked.
	//
	// example:
	//
	// null
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The end time of the maintenance window.
	//
	// example:
	//
	// 04:00:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window.
	//
	// example:
	//
	// 00:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The maximum capacity for elastic scaling.
	//
	// example:
	//
	// 13
	ScaleMax *int32 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum capacity for elastic scaling.
	//
	// example:
	//
	// 1
	ScaleMin *int32 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The cluster status.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// oss
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The tags.
	Tags []*DescribeDBInstancesResponseBodyDataDBInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-8vb5mw****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-uf6kg****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDataDBInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDataDBInstances) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetBid() *string {
	return s.Bid
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetDeletionProtection() *string {
	return s.DeletionProtection
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetDescription() *string {
	return s.Description
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetScaleMax() *int32 {
	return s.ScaleMax
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetScaleMin() *int32 {
	return s.ScaleMin
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetTags() []*DescribeDBInstancesResponseBodyDataDBInstancesTags {
	return s.Tags
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetAliUid(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.AliUid = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetBid(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.Bid = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetChargeType(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetCreateTime(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetDBInstanceId(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetDeletionProtection(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetDescription(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.Description = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetEngine(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetEngineVersion(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetExpireTime(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetLockMode(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetLockReason(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetMaintainEndTime(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetMaintainStartTime(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetRegionId(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetResourceGroupId(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetScaleMax(v int32) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.ScaleMax = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetScaleMin(v int32) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.ScaleMin = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetStatus(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.Status = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetStorageType(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetTags(v []*DescribeDBInstancesResponseBodyDataDBInstancesTags) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.Tags = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetVSwitchId(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetVpcId(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) SetZoneId(v string) *DescribeDBInstancesResponseBodyDataDBInstances {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstances) Validate() error {
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

type DescribeDBInstancesResponseBodyDataDBInstancesTags struct {
	// The tag key.
	//
	// example:
	//
	// tag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDataDBInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDataDBInstancesTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDataDBInstancesTags) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstancesResponseBodyDataDBInstancesTags) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstancesResponseBodyDataDBInstancesTags) SetKey(v string) *DescribeDBInstancesResponseBodyDataDBInstancesTags {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstancesTags) SetValue(v string) *DescribeDBInstancesResponseBodyDataDBInstancesTags {
	s.Value = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDataDBInstancesTags) Validate() error {
	return dara.Validate(s)
}
