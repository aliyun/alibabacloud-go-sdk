// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDBInstanceAttributeResponseBodyData) *DescribeDBInstanceAttributeResponseBody
	GetData() *DescribeDBInstanceAttributeResponseBodyData
	SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceAttributeResponseBody struct {
	// The result returned.
	Data *DescribeDBInstanceAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBody) GetData() *DescribeDBInstanceAttributeResponseBodyData {
	return s.Data
}

func (s *DescribeDBInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceAttributeResponseBody) SetData(v *DescribeDBInstanceAttributeResponseBodyData) *DescribeDBInstanceAttributeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceAttributeResponseBodyData struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 140692647406****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The channel ID.
	//
	// example:
	//
	// PD39050615820269****
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// example:
	//
	// enterprise
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The billing method. Enterprise Edition clusters use the pay-as-you-go billing method.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// activation
	ClickObserveServiceStatus *string `json:"ClickObserveServiceStatus,omitempty" xml:"ClickObserveServiceStatus,omitempty"`
	// The time when the cluster was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2023-09-14T08:14:48Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Indicates whether the release protection feature is enabled for the cluster.
	//
	// example:
	//
	// 0/1
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The deployment mode of the cluster. Valid values: single_az and multi_az.
	//
	// 	- single_az: indicates that the server nodes are deployed in the primary zone. The ID of the primary zone is specified by the ZoneID parameter.
	//
	// 	- multi_az: indicates that the server nodes are deployed in multiple zones. The information about the zones is specified by the MultiZones parameter.
	//
	// The keeper nodes are deployed in multiple zones.
	//
	// example:
	//
	// single_az
	DeploySchema *string `json:"DeploySchema,omitempty" xml:"DeploySchema,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// Used for test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The disabled database ports. Multiple database ports are separated by commas (,).
	//
	// example:
	//
	// 9001,8123
	DisabledPorts *string `json:"DisabledPorts,omitempty" xml:"DisabledPorts,omitempty"`
	// The engine type.
	//
	// example:
	//
	// clickhouse
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The minor engine version of the cluster.
	//
	// example:
	//
	// 23.8.1.41495_6
	EngineMinorVersion *string `json:"EngineMinorVersion,omitempty" xml:"EngineMinorVersion,omitempty"`
	// The engine version.
	//
	// example:
	//
	// 23.8
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the cluster expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// >  Pay-as-you-go clusters never expire. If the cluster is a pay-as-you-go cluster, an empty string is returned for this parameter.
	//
	// example:
	//
	// 2024-04-17T08:14:48Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The latest minor engine version.
	//
	// example:
	//
	// 23.8.1.41495_6
	LatestEngineMinorVersion *string `json:"LatestEngineMinorVersion,omitempty" xml:"LatestEngineMinorVersion,omitempty"`
	// The lock mode of the cluster.
	//
	// example:
	//
	// 0
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the cluster was locked.
	//
	// example:
	//
	// nolock
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The end time of the maintenance window.
	//
	// example:
	//
	// 21:00
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window.
	//
	// example:
	//
	// 12:00
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The information about the zones.
	MultiZones []*DescribeDBInstanceAttributeResponseBodyDataMultiZones `json:"MultiZones,omitempty" xml:"MultiZones,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	NodeCount *string `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// example:
	//
	// 32
	NodeScaleMax *string `json:"NodeScaleMax,omitempty" xml:"NodeScaleMax,omitempty"`
	// example:
	//
	// 4
	NodeScaleMin *string `json:"NodeScaleMin,omitempty" xml:"NodeScaleMin,omitempty"`
	// The nodes.
	Nodes []*DescribeDBInstanceAttributeResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The size of the object storage space.
	//
	// example:
	//
	// 13
	ObjectStoreSize *string `json:"ObjectStoreSize,omitempty" xml:"ObjectStoreSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// rg-acfmzygvt54****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The maximum capacity for elastic scaling.
	//
	// example:
	//
	// 32
	ScaleMax *int32 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum capacity for elastic scaling.
	//
	// example:
	//
	// 8
	ScaleMin *int32 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The cluster status.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 100
	StorageQuota *string `json:"StorageQuota,omitempty" xml:"StorageQuota,omitempty"`
	// The size of the storage space. Unit: GB.
	//
	// example:
	//
	// 12
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The storage type.
	//
	// example:
	//
	// 100
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The details of the tags.
	Tags []*DescribeDBInstanceAttributeResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf67ij56zm9x4uc6hmilg
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-wz9duj8xd6r1gzhsg*****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetBid() *string {
	return s.Bid
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetClickObserveServiceStatus() *string {
	return s.ClickObserveServiceStatus
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetDeploySchema() *string {
	return s.DeploySchema
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetDisabledPorts() *string {
	return s.DisabledPorts
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetEngineMinorVersion() *string {
	return s.EngineMinorVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetLatestEngineMinorVersion() *string {
	return s.LatestEngineMinorVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetMultiZones() []*DescribeDBInstanceAttributeResponseBodyDataMultiZones {
	return s.MultiZones
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetNodeCount() *string {
	return s.NodeCount
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetNodeScaleMax() *string {
	return s.NodeScaleMax
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetNodeScaleMin() *string {
	return s.NodeScaleMin
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetNodes() []*DescribeDBInstanceAttributeResponseBodyDataNodes {
	return s.Nodes
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetObjectStoreSize() *string {
	return s.ObjectStoreSize
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetScaleMax() *int32 {
	return s.ScaleMax
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetScaleMin() *int32 {
	return s.ScaleMin
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetStorageQuota() *string {
	return s.StorageQuota
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetTags() []*DescribeDBInstanceAttributeResponseBodyDataTags {
	return s.Tags
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstanceAttributeResponseBodyData) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetAliUid(v int64) *DescribeDBInstanceAttributeResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetBid(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.Bid = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetCategory(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.Category = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetChargeType(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetClickObserveServiceStatus(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.ClickObserveServiceStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetCreateTime(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetDBInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetDeletionProtection(v bool) *DescribeDBInstanceAttributeResponseBodyData {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetDeploySchema(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.DeploySchema = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetDescription(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetDisabledPorts(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.DisabledPorts = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetEngine(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetEngineMinorVersion(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.EngineMinorVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetEngineVersion(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetExpireTime(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetLatestEngineMinorVersion(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.LatestEngineMinorVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetLockMode(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetLockReason(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetMaintainEndTime(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetMaintainStartTime(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetMultiZones(v []*DescribeDBInstanceAttributeResponseBodyDataMultiZones) *DescribeDBInstanceAttributeResponseBodyData {
	s.MultiZones = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetNodeCount(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.NodeCount = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetNodeScaleMax(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.NodeScaleMax = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetNodeScaleMin(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.NodeScaleMin = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetNodes(v []*DescribeDBInstanceAttributeResponseBodyDataNodes) *DescribeDBInstanceAttributeResponseBodyData {
	s.Nodes = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetObjectStoreSize(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.ObjectStoreSize = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetRegionId(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetResourceGroupId(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetScaleMax(v int32) *DescribeDBInstanceAttributeResponseBodyData {
	s.ScaleMax = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetScaleMin(v int32) *DescribeDBInstanceAttributeResponseBodyData {
	s.ScaleMin = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetStatus(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetStorageQuota(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.StorageQuota = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetStorageSize(v int32) *DescribeDBInstanceAttributeResponseBodyData {
	s.StorageSize = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetStorageType(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetTags(v []*DescribeDBInstanceAttributeResponseBodyDataTags) *DescribeDBInstanceAttributeResponseBodyData {
	s.Tags = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetVpcId(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyData {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyData) Validate() error {
	if s.MultiZones != nil {
		for _, item := range s.MultiZones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type DescribeDBInstanceAttributeResponseBodyDataMultiZones struct {
	// The vSwitch IDs.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDataMultiZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDataMultiZones) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDataMultiZones) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeDBInstanceAttributeResponseBodyDataMultiZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstanceAttributeResponseBodyDataMultiZones) SetVSwitchIds(v []*string) *DescribeDBInstanceAttributeResponseBodyDataMultiZones {
	s.VSwitchIds = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDataMultiZones) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyDataMultiZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDataMultiZones) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDataNodes struct {
	// The node status.
	//
	// example:
	//
	// active
	NodeStatus *string `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDataNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDataNodes) GetNodeStatus() *string {
	return s.NodeStatus
}

func (s *DescribeDBInstanceAttributeResponseBodyDataNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstanceAttributeResponseBodyDataNodes) SetNodeStatus(v string) *DescribeDBInstanceAttributeResponseBodyDataNodes {
	s.NodeStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDataNodes) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyDataNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDataNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDataTags struct {
	// The key of the tag.
	//
	// example:
	//
	// id
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// ck
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstanceAttributeResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstanceAttributeResponseBodyDataTags) SetKey(v string) *DescribeDBInstanceAttributeResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDataTags) SetValue(v string) *DescribeDBInstanceAttributeResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
