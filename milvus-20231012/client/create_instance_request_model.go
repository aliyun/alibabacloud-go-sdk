// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateInstanceRequest
	GetRegionId() *string
	SetAiFunction(v bool) *CreateInstanceRequest
	GetAiFunction() *bool
	SetAutoBackup(v bool) *CreateInstanceRequest
	GetAutoBackup() *bool
	SetAutoPay(v bool) *CreateInstanceRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateInstanceRequest
	GetAutoRenew() *bool
	SetBackupRestoreInfo(v *CreateInstanceRequestBackupRestoreInfo) *CreateInstanceRequest
	GetBackupRestoreInfo() *CreateInstanceRequestBackupRestoreInfo
	SetComponents(v []*CreateInstanceRequestComponents) *CreateInstanceRequest
	GetComponents() []*CreateInstanceRequestComponents
	SetConfiguration(v string) *CreateInstanceRequest
	GetConfiguration() *string
	SetDbAdminPassword(v string) *CreateInstanceRequest
	GetDbAdminPassword() *string
	SetDbVersion(v string) *CreateInstanceRequest
	GetDbVersion() *string
	SetEncrypted(v bool) *CreateInstanceRequest
	GetEncrypted() *bool
	SetHa(v bool) *CreateInstanceRequest
	GetHa() *bool
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetIsMultiAzStorage(v bool) *CreateInstanceRequest
	GetIsMultiAzStorage() *bool
	SetKmsKeyId(v string) *CreateInstanceRequest
	GetKmsKeyId() *string
	SetLoadReplicas(v int32) *CreateInstanceRequest
	GetLoadReplicas() *int32
	SetMultiZoneMode(v string) *CreateInstanceRequest
	GetMultiZoneMode() *string
	SetPaymentDuration(v int32) *CreateInstanceRequest
	GetPaymentDuration() *int32
	SetPaymentDurationUnit(v string) *CreateInstanceRequest
	GetPaymentDurationUnit() *string
	SetPaymentType(v string) *CreateInstanceRequest
	GetPaymentType() *string
	SetPromotionNo(v string) *CreateInstanceRequest
	GetPromotionNo() *string
	SetResourceGroupId(v string) *CreateInstanceRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest
	GetTags() []*CreateInstanceRequestTags
	SetVSwitchIds(v []*CreateInstanceRequestVSwitchIds) *CreateInstanceRequest
	GetVSwitchIds() []*CreateInstanceRequestVSwitchIds
	SetVpcId(v string) *CreateInstanceRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateInstanceRequest
	GetZoneId() *string
	SetClientToken(v string) *CreateInstanceRequest
	GetClientToken() *string
}

type CreateInstanceRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable AI function.
	AiFunction *bool `json:"aiFunction,omitempty" xml:"aiFunction,omitempty"`
	// Specifies whether to enable automatic backup.
	//
	// example:
	//
	// true
	AutoBackup *bool `json:"autoBackup,omitempty" xml:"autoBackup,omitempty"`
	// Specifies whether to enable automatic payment. Default value: true. Valid values:
	//
	// example:
	//
	// true
	AutoPay *bool `json:"autoPay,omitempty" xml:"autoPay,omitempty"`
	// Specifies whether to enable auto-renewal. This parameter takes effect only when the payment type is set to Subscription.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The backup and restoration information.
	BackupRestoreInfo *CreateInstanceRequestBackupRestoreInfo `json:"backupRestoreInfo,omitempty" xml:"backupRestoreInfo,omitempty" type:"Struct"`
	// The component information.
	Components []*CreateInstanceRequestComponents `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	// The configuration items.
	//
	// example:
	//
	// rootCoord:
	//
	//     maxDatabaseNum: 64 # Maximum number of database
	//
	//     maxPartitionNum: 4096
	Configuration *string `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// The database administrator password.
	//
	// example:
	//
	// test12
	DbAdminPassword *string `json:"dbAdminPassword,omitempty" xml:"dbAdminPassword,omitempty"`
	// The Milvus version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2.4
	DbVersion *string `json:"dbVersion,omitempty" xml:"dbVersion,omitempty"`
	// Specifies whether to enable OSS encryption.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// Specifies whether to enable high availability.
	//
	// example:
	//
	// true
	Ha *bool `json:"ha,omitempty" xml:"ha,omitempty"`
	// The instance name.
	//
	// example:
	//
	// milvus-test
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// Specifies whether to enable multi-zone storage.
	IsMultiAzStorage *bool `json:"isMultiAzStorage,omitempty" xml:"isMultiAzStorage,omitempty"`
	// The ID of the KMS key used for encryption.
	//
	// example:
	//
	// key-xxx
	KmsKeyId *string `json:"kmsKeyId,omitempty" xml:"kmsKeyId,omitempty"`
	// The number of load replicas.
	//
	// example:
	//
	// 2
	LoadReplicas *int32 `json:"loadReplicas,omitempty" xml:"loadReplicas,omitempty"`
	// The zone configuration.
	//
	// example:
	//
	// Single
	MultiZoneMode *string `json:"multiZoneMode,omitempty" xml:"multiZoneMode,omitempty"`
	// The payment duration.
	//
	// example:
	//
	// 1
	PaymentDuration *int32 `json:"paymentDuration,omitempty" xml:"paymentDuration,omitempty"`
	// The payment duration unit.
	//
	// example:
	//
	// month
	PaymentDurationUnit *string `json:"paymentDurationUnit,omitempty" xml:"paymentDurationUnit,omitempty"`
	// The payment type.
	//
	// This parameter is required.
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// 22120151****
	PromotionNo *string `json:"promotionNo,omitempty" xml:"promotionNo,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-xxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The instance tags.
	Tags []*CreateInstanceRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The vSwitches.
	//
	// example:
	//
	// ["vsw-123xxx"]
	VSwitchIds []*CreateInstanceRequestVSwitchIds `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-123xxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The primary zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-g
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// xxx
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInstanceRequest) GetAiFunction() *bool {
	return s.AiFunction
}

func (s *CreateInstanceRequest) GetAutoBackup() *bool {
	return s.AutoBackup
}

func (s *CreateInstanceRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateInstanceRequest) GetBackupRestoreInfo() *CreateInstanceRequestBackupRestoreInfo {
	return s.BackupRestoreInfo
}

func (s *CreateInstanceRequest) GetComponents() []*CreateInstanceRequestComponents {
	return s.Components
}

func (s *CreateInstanceRequest) GetConfiguration() *string {
	return s.Configuration
}

func (s *CreateInstanceRequest) GetDbAdminPassword() *string {
	return s.DbAdminPassword
}

func (s *CreateInstanceRequest) GetDbVersion() *string {
	return s.DbVersion
}

func (s *CreateInstanceRequest) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateInstanceRequest) GetHa() *bool {
	return s.Ha
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetIsMultiAzStorage() *bool {
	return s.IsMultiAzStorage
}

func (s *CreateInstanceRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *CreateInstanceRequest) GetLoadReplicas() *int32 {
	return s.LoadReplicas
}

func (s *CreateInstanceRequest) GetMultiZoneMode() *string {
	return s.MultiZoneMode
}

func (s *CreateInstanceRequest) GetPaymentDuration() *int32 {
	return s.PaymentDuration
}

func (s *CreateInstanceRequest) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *CreateInstanceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateInstanceRequest) GetPromotionNo() *string {
	return s.PromotionNo
}

func (s *CreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceRequest) GetTags() []*CreateInstanceRequestTags {
	return s.Tags
}

func (s *CreateInstanceRequest) GetVSwitchIds() []*CreateInstanceRequestVSwitchIds {
	return s.VSwitchIds
}

func (s *CreateInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateInstanceRequest) SetRegionId(v string) *CreateInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceRequest) SetAiFunction(v bool) *CreateInstanceRequest {
	s.AiFunction = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoBackup(v bool) *CreateInstanceRequest {
	s.AutoBackup = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoPay(v bool) *CreateInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetBackupRestoreInfo(v *CreateInstanceRequestBackupRestoreInfo) *CreateInstanceRequest {
	s.BackupRestoreInfo = v
	return s
}

func (s *CreateInstanceRequest) SetComponents(v []*CreateInstanceRequestComponents) *CreateInstanceRequest {
	s.Components = v
	return s
}

func (s *CreateInstanceRequest) SetConfiguration(v string) *CreateInstanceRequest {
	s.Configuration = &v
	return s
}

func (s *CreateInstanceRequest) SetDbAdminPassword(v string) *CreateInstanceRequest {
	s.DbAdminPassword = &v
	return s
}

func (s *CreateInstanceRequest) SetDbVersion(v string) *CreateInstanceRequest {
	s.DbVersion = &v
	return s
}

func (s *CreateInstanceRequest) SetEncrypted(v bool) *CreateInstanceRequest {
	s.Encrypted = &v
	return s
}

func (s *CreateInstanceRequest) SetHa(v bool) *CreateInstanceRequest {
	s.Ha = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetIsMultiAzStorage(v bool) *CreateInstanceRequest {
	s.IsMultiAzStorage = &v
	return s
}

func (s *CreateInstanceRequest) SetKmsKeyId(v string) *CreateInstanceRequest {
	s.KmsKeyId = &v
	return s
}

func (s *CreateInstanceRequest) SetLoadReplicas(v int32) *CreateInstanceRequest {
	s.LoadReplicas = &v
	return s
}

func (s *CreateInstanceRequest) SetMultiZoneMode(v string) *CreateInstanceRequest {
	s.MultiZoneMode = &v
	return s
}

func (s *CreateInstanceRequest) SetPaymentDuration(v int32) *CreateInstanceRequest {
	s.PaymentDuration = &v
	return s
}

func (s *CreateInstanceRequest) SetPaymentDurationUnit(v string) *CreateInstanceRequest {
	s.PaymentDurationUnit = &v
	return s
}

func (s *CreateInstanceRequest) SetPaymentType(v string) *CreateInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInstanceRequest) SetPromotionNo(v string) *CreateInstanceRequest {
	s.PromotionNo = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest {
	s.Tags = v
	return s
}

func (s *CreateInstanceRequest) SetVSwitchIds(v []*CreateInstanceRequestVSwitchIds) *CreateInstanceRequest {
	s.VSwitchIds = v
	return s
}

func (s *CreateInstanceRequest) SetVpcId(v string) *CreateInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceRequest) SetZoneId(v string) *CreateInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	if s.BackupRestoreInfo != nil {
		if err := s.BackupRestoreInfo.Validate(); err != nil {
			return err
		}
	}
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

type CreateInstanceRequestBackupRestoreInfo struct {
	// The backup ID.
	//
	// example:
	//
	// bt-xxxxx
	BackupId *string `json:"backupId,omitempty" xml:"backupId,omitempty"`
	// The backup name.
	//
	// example:
	//
	// Backup1
	BackupName *string `json:"backupName,omitempty" xml:"backupName,omitempty"`
	// The ID of the source backup cluster.
	//
	// example:
	//
	// c-xxxxxxx
	SourceClusterId *string `json:"sourceClusterId,omitempty" xml:"sourceClusterId,omitempty"`
}

func (s CreateInstanceRequestBackupRestoreInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestBackupRestoreInfo) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestBackupRestoreInfo) GetBackupId() *string {
	return s.BackupId
}

func (s *CreateInstanceRequestBackupRestoreInfo) GetBackupName() *string {
	return s.BackupName
}

func (s *CreateInstanceRequestBackupRestoreInfo) GetSourceClusterId() *string {
	return s.SourceClusterId
}

func (s *CreateInstanceRequestBackupRestoreInfo) SetBackupId(v string) *CreateInstanceRequestBackupRestoreInfo {
	s.BackupId = &v
	return s
}

func (s *CreateInstanceRequestBackupRestoreInfo) SetBackupName(v string) *CreateInstanceRequestBackupRestoreInfo {
	s.BackupName = &v
	return s
}

func (s *CreateInstanceRequestBackupRestoreInfo) SetSourceClusterId(v string) *CreateInstanceRequestBackupRestoreInfo {
	s.SourceClusterId = &v
	return s
}

func (s *CreateInstanceRequestBackupRestoreInfo) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestComponents struct {
	// The number of compute units (CUs).
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	CuNum *int32 `json:"cuNum,omitempty" xml:"cuNum,omitempty"`
	// The CU type.
	//
	// example:
	//
	// general
	CuType   *string                                  `json:"cuType,omitempty" xml:"cuType,omitempty"`
	DataDisk *CreateInstanceRequestComponentsDataDisk `json:"dataDisk,omitempty" xml:"dataDisk,omitempty" type:"Struct"`
	// The disk size type for Query Node. Set to Large for storage-optimized, and Normal for compute-optimized or other configurations.
	//
	// example:
	//
	// Normal
	DiskSizeType *string `json:"diskSizeType,omitempty" xml:"diskSizeType,omitempty"`
	// The number of replicas.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
	// The component type.
	//
	// This parameter is required.
	//
	// example:
	//
	// standalone
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateInstanceRequestComponents) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestComponents) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestComponents) GetCuNum() *int32 {
	return s.CuNum
}

func (s *CreateInstanceRequestComponents) GetCuType() *string {
	return s.CuType
}

func (s *CreateInstanceRequestComponents) GetDataDisk() *CreateInstanceRequestComponentsDataDisk {
	return s.DataDisk
}

func (s *CreateInstanceRequestComponents) GetDiskSizeType() *string {
	return s.DiskSizeType
}

func (s *CreateInstanceRequestComponents) GetReplica() *int32 {
	return s.Replica
}

func (s *CreateInstanceRequestComponents) GetType() *string {
	return s.Type
}

func (s *CreateInstanceRequestComponents) SetCuNum(v int32) *CreateInstanceRequestComponents {
	s.CuNum = &v
	return s
}

func (s *CreateInstanceRequestComponents) SetCuType(v string) *CreateInstanceRequestComponents {
	s.CuType = &v
	return s
}

func (s *CreateInstanceRequestComponents) SetDataDisk(v *CreateInstanceRequestComponentsDataDisk) *CreateInstanceRequestComponents {
	s.DataDisk = v
	return s
}

func (s *CreateInstanceRequestComponents) SetDiskSizeType(v string) *CreateInstanceRequestComponents {
	s.DiskSizeType = &v
	return s
}

func (s *CreateInstanceRequestComponents) SetReplica(v int32) *CreateInstanceRequestComponents {
	s.Replica = &v
	return s
}

func (s *CreateInstanceRequestComponents) SetType(v string) *CreateInstanceRequestComponents {
	s.Type = &v
	return s
}

func (s *CreateInstanceRequestComponents) Validate() error {
	if s.DataDisk != nil {
		if err := s.DataDisk.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInstanceRequestComponentsDataDisk struct {
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
	// example:
	//
	// 100
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// alicloud-disk-essd-pl1
	StorageClass *string `json:"storageClass,omitempty" xml:"storageClass,omitempty"`
}

func (s CreateInstanceRequestComponentsDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestComponentsDataDisk) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestComponentsDataDisk) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateInstanceRequestComponentsDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateInstanceRequestComponentsDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateInstanceRequestComponentsDataDisk) GetStorageClass() *string {
	return s.StorageClass
}

func (s *CreateInstanceRequestComponentsDataDisk) SetEnabled(v bool) *CreateInstanceRequestComponentsDataDisk {
	s.Enabled = &v
	return s
}

func (s *CreateInstanceRequestComponentsDataDisk) SetPerformanceLevel(v string) *CreateInstanceRequestComponentsDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateInstanceRequestComponentsDataDisk) SetSize(v int32) *CreateInstanceRequestComponentsDataDisk {
	s.Size = &v
	return s
}

func (s *CreateInstanceRequestComponentsDataDisk) SetStorageClass(v string) *CreateInstanceRequestComponentsDataDisk {
	s.StorageClass = &v
	return s
}

func (s *CreateInstanceRequestComponentsDataDisk) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestTags struct {
	// The key of the resource tag.
	//
	// example:
	//
	// key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the resource tag.
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateInstanceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestTags) SetKey(v string) *CreateInstanceRequestTags {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTags) SetValue(v string) *CreateInstanceRequestTags {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestVSwitchIds struct {
	// The vSwitch ID configuration in the zone.
	//
	// example:
	//
	// vsw-xxx
	VswId *string `json:"vswId,omitempty" xml:"vswId,omitempty"`
	// The zone.
	//
	// example:
	//
	// cn-shanghai-a
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s CreateInstanceRequestVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestVSwitchIds) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestVSwitchIds) GetVswId() *string {
	return s.VswId
}

func (s *CreateInstanceRequestVSwitchIds) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceRequestVSwitchIds) SetVswId(v string) *CreateInstanceRequestVSwitchIds {
	s.VswId = &v
	return s
}

func (s *CreateInstanceRequestVSwitchIds) SetZoneId(v string) *CreateInstanceRequestVSwitchIds {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceRequestVSwitchIds) Validate() error {
	return dara.Validate(s)
}
