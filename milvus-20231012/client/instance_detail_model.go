// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceDetail interface {
	dara.Model
	String() string
	GoString() string
	SetAutoBackup(v bool) *InstanceDetail
	GetAutoBackup() *bool
	SetComponents(v []*InstanceDetailComponents) *InstanceDetail
	GetComponents() []*InstanceDetailComponents
	SetConfiguration(v string) *InstanceDetail
	GetConfiguration() *string
	SetCreateTime(v string) *InstanceDetail
	GetCreateTime() *string
	SetDbVersion(v string) *InstanceDetail
	GetDbVersion() *string
	SetEncrypted(v bool) *InstanceDetail
	GetEncrypted() *bool
	SetExpireTime(v string) *InstanceDetail
	GetExpireTime() *string
	SetHa(v bool) *InstanceDetail
	GetHa() *bool
	SetInstanceId(v string) *InstanceDetail
	GetInstanceId() *string
	SetInstanceName(v string) *InstanceDetail
	GetInstanceName() *string
	SetKmsKeyId(v string) *InstanceDetail
	GetKmsKeyId() *string
	SetMultiZoneMode(v string) *InstanceDetail
	GetMultiZoneMode() *string
	SetOrderId(v string) *InstanceDetail
	GetOrderId() *string
	SetPaymentType(v string) *InstanceDetail
	GetPaymentType() *string
	SetRegionId(v string) *InstanceDetail
	GetRegionId() *string
	SetResourceGroupId(v string) *InstanceDetail
	GetResourceGroupId() *string
	SetRunningTime(v int64) *InstanceDetail
	GetRunningTime() *int64
	SetSecurityGroupIds(v []*string) *InstanceDetail
	GetSecurityGroupIds() []*string
	SetStatus(v string) *InstanceDetail
	GetStatus() *string
	SetTags(v []*InstanceDetailTags) *InstanceDetail
	GetTags() []*InstanceDetailTags
	SetVSwitchIds(v []*InstanceDetailVSwitchIds) *InstanceDetail
	GetVSwitchIds() []*InstanceDetailVSwitchIds
	SetVpcId(v string) *InstanceDetail
	GetVpcId() *string
	SetZoneId(v string) *InstanceDetail
	GetZoneId() *string
}

type InstanceDetail struct {
	// The automatic backup configuration.
	//
	// example:
	//
	// true
	AutoBackup *bool `json:"autoBackup,omitempty" xml:"autoBackup,omitempty"`
	// The component information.
	Components []*InstanceDetailComponents `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	// The configuration.
	//
	// example:
	//
	// rootCoord:
	//
	//     maxDatabaseNum: 64 # Maximum number of database
	//
	//     maxPartitionNum: 4096
	Configuration *string `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// The creation time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-04-27T02:04:25Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The instance version.
	//
	// example:
	//
	// 2.5
	DbVersion *string `json:"dbVersion,omitempty" xml:"dbVersion,omitempty"`
	// Indicates whether data encryption is enabled.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// The expiration time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-04-27T02:04:25Z
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// Indicates whether high availability is enabled.
	//
	// example:
	//
	// true
	Ha *bool `json:"ha,omitempty" xml:"ha,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// c-xxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// milvus-test
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// kms key Id。
	//
	// example:
	//
	// key-xxx
	KmsKeyId *string `json:"kmsKeyId,omitempty" xml:"kmsKeyId,omitempty"`
	// The multi-zone deployment mode.
	//
	// example:
	//
	// Single
	MultiZoneMode *string `json:"multiZoneMode,omitempty" xml:"multiZoneMode,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 4751
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// The billing method. Valid values: PayAsYouGo: pay-as-you-go billing method. Subscription: subscription.
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek3dcgyq7pnqwa
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The running time.
	//
	// example:
	//
	// 1
	RunningTime *int64 `json:"runningTime,omitempty" xml:"runningTime,omitempty"`
	// The security group IDs.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty" type:"Repeated"`
	// The instance status. Valid values:
	//
	// - creating: Being created.
	//
	// - running: Running.
	//
	// - updating: Being upgraded. This includes specification changes, configuration changes, and public network access changes.
	//
	// - disable: Unavailable. The cluster has expired and requires renewal to reactivate.
	//
	// - deleting: Being deleted.
	//
	// - deleted: Deleted.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tags.
	Tags []*InstanceDetailTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The vSwitch IDs.
	VSwitchIds []*InstanceDetailVSwitchIds `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-xxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s InstanceDetail) String() string {
	return dara.Prettify(s)
}

func (s InstanceDetail) GoString() string {
	return s.String()
}

func (s *InstanceDetail) GetAutoBackup() *bool {
	return s.AutoBackup
}

func (s *InstanceDetail) GetComponents() []*InstanceDetailComponents {
	return s.Components
}

func (s *InstanceDetail) GetConfiguration() *string {
	return s.Configuration
}

func (s *InstanceDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *InstanceDetail) GetDbVersion() *string {
	return s.DbVersion
}

func (s *InstanceDetail) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *InstanceDetail) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *InstanceDetail) GetHa() *bool {
	return s.Ha
}

func (s *InstanceDetail) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InstanceDetail) GetInstanceName() *string {
	return s.InstanceName
}

func (s *InstanceDetail) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *InstanceDetail) GetMultiZoneMode() *string {
	return s.MultiZoneMode
}

func (s *InstanceDetail) GetOrderId() *string {
	return s.OrderId
}

func (s *InstanceDetail) GetPaymentType() *string {
	return s.PaymentType
}

func (s *InstanceDetail) GetRegionId() *string {
	return s.RegionId
}

func (s *InstanceDetail) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *InstanceDetail) GetRunningTime() *int64 {
	return s.RunningTime
}

func (s *InstanceDetail) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *InstanceDetail) GetStatus() *string {
	return s.Status
}

func (s *InstanceDetail) GetTags() []*InstanceDetailTags {
	return s.Tags
}

func (s *InstanceDetail) GetVSwitchIds() []*InstanceDetailVSwitchIds {
	return s.VSwitchIds
}

func (s *InstanceDetail) GetVpcId() *string {
	return s.VpcId
}

func (s *InstanceDetail) GetZoneId() *string {
	return s.ZoneId
}

func (s *InstanceDetail) SetAutoBackup(v bool) *InstanceDetail {
	s.AutoBackup = &v
	return s
}

func (s *InstanceDetail) SetComponents(v []*InstanceDetailComponents) *InstanceDetail {
	s.Components = v
	return s
}

func (s *InstanceDetail) SetConfiguration(v string) *InstanceDetail {
	s.Configuration = &v
	return s
}

func (s *InstanceDetail) SetCreateTime(v string) *InstanceDetail {
	s.CreateTime = &v
	return s
}

func (s *InstanceDetail) SetDbVersion(v string) *InstanceDetail {
	s.DbVersion = &v
	return s
}

func (s *InstanceDetail) SetEncrypted(v bool) *InstanceDetail {
	s.Encrypted = &v
	return s
}

func (s *InstanceDetail) SetExpireTime(v string) *InstanceDetail {
	s.ExpireTime = &v
	return s
}

func (s *InstanceDetail) SetHa(v bool) *InstanceDetail {
	s.Ha = &v
	return s
}

func (s *InstanceDetail) SetInstanceId(v string) *InstanceDetail {
	s.InstanceId = &v
	return s
}

func (s *InstanceDetail) SetInstanceName(v string) *InstanceDetail {
	s.InstanceName = &v
	return s
}

func (s *InstanceDetail) SetKmsKeyId(v string) *InstanceDetail {
	s.KmsKeyId = &v
	return s
}

func (s *InstanceDetail) SetMultiZoneMode(v string) *InstanceDetail {
	s.MultiZoneMode = &v
	return s
}

func (s *InstanceDetail) SetOrderId(v string) *InstanceDetail {
	s.OrderId = &v
	return s
}

func (s *InstanceDetail) SetPaymentType(v string) *InstanceDetail {
	s.PaymentType = &v
	return s
}

func (s *InstanceDetail) SetRegionId(v string) *InstanceDetail {
	s.RegionId = &v
	return s
}

func (s *InstanceDetail) SetResourceGroupId(v string) *InstanceDetail {
	s.ResourceGroupId = &v
	return s
}

func (s *InstanceDetail) SetRunningTime(v int64) *InstanceDetail {
	s.RunningTime = &v
	return s
}

func (s *InstanceDetail) SetSecurityGroupIds(v []*string) *InstanceDetail {
	s.SecurityGroupIds = v
	return s
}

func (s *InstanceDetail) SetStatus(v string) *InstanceDetail {
	s.Status = &v
	return s
}

func (s *InstanceDetail) SetTags(v []*InstanceDetailTags) *InstanceDetail {
	s.Tags = v
	return s
}

func (s *InstanceDetail) SetVSwitchIds(v []*InstanceDetailVSwitchIds) *InstanceDetail {
	s.VSwitchIds = v
	return s
}

func (s *InstanceDetail) SetVpcId(v string) *InstanceDetail {
	s.VpcId = &v
	return s
}

func (s *InstanceDetail) SetZoneId(v string) *InstanceDetail {
	s.ZoneId = &v
	return s
}

func (s *InstanceDetail) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
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
	if s.VSwitchIds != nil {
		for _, item := range s.VSwitchIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InstanceDetailComponents struct {
	// The number of CUs.
	//
	// example:
	//
	// 4
	CuNum *int32 `json:"cuNum,omitempty" xml:"cuNum,omitempty"`
	// The CU type.
	//
	// example:
	//
	// general
	CuType   *string                           `json:"cuType,omitempty" xml:"cuType,omitempty"`
	DataDisk *InstanceDetailComponentsDataDisk `json:"dataDisk,omitempty" xml:"dataDisk,omitempty" type:"Struct"`
	// The disk size type for the Query Node. Set this parameter to Large for storage-optimized configurations, and to Normal for other configurations.
	//
	// example:
	//
	// Normal
	DiskSizeType *string                             `json:"diskSizeType,omitempty" xml:"diskSizeType,omitempty"`
	PayType      *string                             `json:"payType,omitempty" xml:"payType,omitempty"`
	PodsList     []*InstanceDetailComponentsPodsList `json:"podsList,omitempty" xml:"podsList,omitempty" type:"Repeated"`
	// The number of replicas.
	//
	// example:
	//
	// 1
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
	// The component type.
	//
	// example:
	//
	// data
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InstanceDetailComponents) String() string {
	return dara.Prettify(s)
}

func (s InstanceDetailComponents) GoString() string {
	return s.String()
}

func (s *InstanceDetailComponents) GetCuNum() *int32 {
	return s.CuNum
}

func (s *InstanceDetailComponents) GetCuType() *string {
	return s.CuType
}

func (s *InstanceDetailComponents) GetDataDisk() *InstanceDetailComponentsDataDisk {
	return s.DataDisk
}

func (s *InstanceDetailComponents) GetDiskSizeType() *string {
	return s.DiskSizeType
}

func (s *InstanceDetailComponents) GetPayType() *string {
	return s.PayType
}

func (s *InstanceDetailComponents) GetPodsList() []*InstanceDetailComponentsPodsList {
	return s.PodsList
}

func (s *InstanceDetailComponents) GetReplica() *int32 {
	return s.Replica
}

func (s *InstanceDetailComponents) GetType() *string {
	return s.Type
}

func (s *InstanceDetailComponents) SetCuNum(v int32) *InstanceDetailComponents {
	s.CuNum = &v
	return s
}

func (s *InstanceDetailComponents) SetCuType(v string) *InstanceDetailComponents {
	s.CuType = &v
	return s
}

func (s *InstanceDetailComponents) SetDataDisk(v *InstanceDetailComponentsDataDisk) *InstanceDetailComponents {
	s.DataDisk = v
	return s
}

func (s *InstanceDetailComponents) SetDiskSizeType(v string) *InstanceDetailComponents {
	s.DiskSizeType = &v
	return s
}

func (s *InstanceDetailComponents) SetPayType(v string) *InstanceDetailComponents {
	s.PayType = &v
	return s
}

func (s *InstanceDetailComponents) SetPodsList(v []*InstanceDetailComponentsPodsList) *InstanceDetailComponents {
	s.PodsList = v
	return s
}

func (s *InstanceDetailComponents) SetReplica(v int32) *InstanceDetailComponents {
	s.Replica = &v
	return s
}

func (s *InstanceDetailComponents) SetType(v string) *InstanceDetailComponents {
	s.Type = &v
	return s
}

func (s *InstanceDetailComponents) Validate() error {
	if s.DataDisk != nil {
		if err := s.DataDisk.Validate(); err != nil {
			return err
		}
	}
	if s.PodsList != nil {
		for _, item := range s.PodsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InstanceDetailComponentsDataDisk struct {
	Enabled          *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	PerformanceLevel *string `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
	Size             *int32  `json:"size,omitempty" xml:"size,omitempty"`
	StorageClass     *string `json:"storageClass,omitempty" xml:"storageClass,omitempty"`
}

func (s InstanceDetailComponentsDataDisk) String() string {
	return dara.Prettify(s)
}

func (s InstanceDetailComponentsDataDisk) GoString() string {
	return s.String()
}

func (s *InstanceDetailComponentsDataDisk) GetEnabled() *bool {
	return s.Enabled
}

func (s *InstanceDetailComponentsDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *InstanceDetailComponentsDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *InstanceDetailComponentsDataDisk) GetStorageClass() *string {
	return s.StorageClass
}

func (s *InstanceDetailComponentsDataDisk) SetEnabled(v bool) *InstanceDetailComponentsDataDisk {
	s.Enabled = &v
	return s
}

func (s *InstanceDetailComponentsDataDisk) SetPerformanceLevel(v string) *InstanceDetailComponentsDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *InstanceDetailComponentsDataDisk) SetSize(v int32) *InstanceDetailComponentsDataDisk {
	s.Size = &v
	return s
}

func (s *InstanceDetailComponentsDataDisk) SetStorageClass(v string) *InstanceDetailComponentsDataDisk {
	s.StorageClass = &v
	return s
}

func (s *InstanceDetailComponentsDataDisk) Validate() error {
	return dara.Validate(s)
}

type InstanceDetailComponentsPodsList struct {
	PodId   *string `json:"podId,omitempty" xml:"podId,omitempty"`
	PodName *string `json:"podName,omitempty" xml:"podName,omitempty"`
}

func (s InstanceDetailComponentsPodsList) String() string {
	return dara.Prettify(s)
}

func (s InstanceDetailComponentsPodsList) GoString() string {
	return s.String()
}

func (s *InstanceDetailComponentsPodsList) GetPodId() *string {
	return s.PodId
}

func (s *InstanceDetailComponentsPodsList) GetPodName() *string {
	return s.PodName
}

func (s *InstanceDetailComponentsPodsList) SetPodId(v string) *InstanceDetailComponentsPodsList {
	s.PodId = &v
	return s
}

func (s *InstanceDetailComponentsPodsList) SetPodName(v string) *InstanceDetailComponentsPodsList {
	s.PodName = &v
	return s
}

func (s *InstanceDetailComponentsPodsList) Validate() error {
	return dara.Validate(s)
}

type InstanceDetailTags struct {
	// The tag key.
	//
	// example:
	//
	// k1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s InstanceDetailTags) String() string {
	return dara.Prettify(s)
}

func (s InstanceDetailTags) GoString() string {
	return s.String()
}

func (s *InstanceDetailTags) GetKey() *string {
	return s.Key
}

func (s *InstanceDetailTags) GetValue() *string {
	return s.Value
}

func (s *InstanceDetailTags) SetKey(v string) *InstanceDetailTags {
	s.Key = &v
	return s
}

func (s *InstanceDetailTags) SetValue(v string) *InstanceDetailTags {
	s.Value = &v
	return s
}

func (s *InstanceDetailTags) Validate() error {
	return dara.Validate(s)
}

type InstanceDetailVSwitchIds struct {
	// The vSwitch IDs.
	//
	// example:
	//
	// vsw-xxx
	VswId *string `json:"vswId,omitempty" xml:"vswId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s InstanceDetailVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s InstanceDetailVSwitchIds) GoString() string {
	return s.String()
}

func (s *InstanceDetailVSwitchIds) GetVswId() *string {
	return s.VswId
}

func (s *InstanceDetailVSwitchIds) GetZoneId() *string {
	return s.ZoneId
}

func (s *InstanceDetailVSwitchIds) SetVswId(v string) *InstanceDetailVSwitchIds {
	s.VswId = &v
	return s
}

func (s *InstanceDetailVSwitchIds) SetZoneId(v string) *InstanceDetailVSwitchIds {
	s.ZoneId = &v
	return s
}

func (s *InstanceDetailVSwitchIds) Validate() error {
	return dara.Validate(s)
}
