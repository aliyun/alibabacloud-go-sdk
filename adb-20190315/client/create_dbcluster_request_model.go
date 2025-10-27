// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetID(v string) *CreateDBClusterRequest
	GetBackupSetID() *string
	SetClientToken(v string) *CreateDBClusterRequest
	GetClientToken() *string
	SetComputeResource(v string) *CreateDBClusterRequest
	GetComputeResource() *string
	SetDBClusterCategory(v string) *CreateDBClusterRequest
	GetDBClusterCategory() *string
	SetDBClusterClass(v string) *CreateDBClusterRequest
	GetDBClusterClass() *string
	SetDBClusterDescription(v string) *CreateDBClusterRequest
	GetDBClusterDescription() *string
	SetDBClusterNetworkType(v string) *CreateDBClusterRequest
	GetDBClusterNetworkType() *string
	SetDBClusterVersion(v string) *CreateDBClusterRequest
	GetDBClusterVersion() *string
	SetDBNodeGroupCount(v string) *CreateDBClusterRequest
	GetDBNodeGroupCount() *string
	SetDBNodeStorage(v string) *CreateDBClusterRequest
	GetDBNodeStorage() *string
	SetDiskEncryption(v bool) *CreateDBClusterRequest
	GetDiskEncryption() *bool
	SetElasticIOResource(v string) *CreateDBClusterRequest
	GetElasticIOResource() *string
	SetEnableSSL(v bool) *CreateDBClusterRequest
	GetEnableSSL() *bool
	SetExecutorCount(v string) *CreateDBClusterRequest
	GetExecutorCount() *string
	SetKmsId(v string) *CreateDBClusterRequest
	GetKmsId() *string
	SetMode(v string) *CreateDBClusterRequest
	GetMode() *string
	SetOwnerAccount(v string) *CreateDBClusterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDBClusterRequest
	GetOwnerId() *int64
	SetPayType(v string) *CreateDBClusterRequest
	GetPayType() *string
	SetPeriod(v string) *CreateDBClusterRequest
	GetPeriod() *string
	SetRegionId(v string) *CreateDBClusterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBClusterRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDBClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBClusterRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *CreateDBClusterRequest
	GetRestoreTime() *string
	SetRestoreType(v string) *CreateDBClusterRequest
	GetRestoreType() *string
	SetSourceDBInstanceName(v string) *CreateDBClusterRequest
	GetSourceDBInstanceName() *string
	SetStorageResource(v string) *CreateDBClusterRequest
	GetStorageResource() *string
	SetStorageType(v string) *CreateDBClusterRequest
	GetStorageType() *string
	SetTag(v []*CreateDBClusterRequestTag) *CreateDBClusterRequest
	GetTag() []*CreateDBClusterRequestTag
	SetUsedTime(v string) *CreateDBClusterRequest
	GetUsedTime() *string
	SetVPCId(v string) *CreateDBClusterRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateDBClusterRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateDBClusterRequest
	GetZoneId() *string
}

type CreateDBClusterRequest struct {
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	BackupSetID *string `json:"BackupSetID,omitempty" xml:"BackupSetID,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The value is case-sensitive and can contain a maximum of 64 ASCII characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-t7241****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The computing resources of the cluster. You can use computing resources to compute data. The increase in the computing resources can accelerate data queries. The computing resources are available for Cluster Edition and Basic Edition.
	//
	// 	- Computing resources for Cluster Edition include 16 cores and 64 GB memory, 24 cores and 96 GB memory, and 32 cores or more. Cluster Edition supports resource isolation, scheduled scaling, and tiered storage of hot and cold data.
	//
	// 	- Computing resources for Basic Edition include 8 cores and 32 GB memory and 16 cores and 64 GB memory. Alibaba Cloud does not provide an SLA guarantee for Basic Edition, and 4 to 8 hours are required for a failover. We recommend that you do not use Basic Edition in production environments.
	//
	// >
	//
	// 	- You can call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/190632.html) operation to query the available computing resources in a region.
	//
	// 	- This parameter must be specified when Mode is set to **Flexible**.
	//
	// example:
	//
	// 32Core128GB
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// 	- **Cluster**: reserved mode for Cluster Edition
	//
	// <!---->
	//
	// 	- **MixedStorage**: elastic mode for Cluster Edition
	//
	// >  If the DBClusterCategory parameter is set to Cluster, you must set the Mode parameter to Reserver. If the DBClusterCategory parameter is set to MixedStorage, you must set the Mode parameter to Flexible. Otherwise, the cluster fails to be created.
	//
	// This parameter is required.
	//
	// example:
	//
	// Cluster
	DBClusterCategory *string `json:"DBClusterCategory,omitempty" xml:"DBClusterCategory,omitempty"`
	// The specification of the cluster. Valid values:
	//
	// 	- **C8**
	//
	// 	- **C32**
	//
	// >  This parameter is required if the Mode parameter is set to Reserver.
	//
	// example:
	//
	// C8
	DBClusterClass *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	// The description of the cluster.
	//
	// 	- The description cannot start with `http://` or `https`.
	//
	// 	- The description must be 2 to 256 characters in length.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The network type of the cluster. Set the value to **VPC**.
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The version of the cluster. Set the value to **3.0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3.0
	DBClusterVersion *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	// The number of node groups. Valid values: 1 to 200 (integer).
	//
	// >  This parameter is required if the Mode parameter is set to Reserver.
	//
	// example:
	//
	// 2
	DBNodeGroupCount *string `json:"DBNodeGroupCount,omitempty" xml:"DBNodeGroupCount,omitempty"`
	// The storage capacity of the cluster. Unit: GB.
	//
	// 	- Valid values when DBClusterClass is set to C8: 100 to 1000
	//
	// 	- Valid values when DBClusterClass is set to C32: 100 to 8000
	//
	// > 	- This parameter is required if the Mode parameter is set to Reserver.
	//
	// > 	- 1000 The storage capacity less than 1,000 GB increases in 100 GB increments. The storage capacity greater than 1,000 GB increases in 1,000 GB increments.
	//
	// example:
	//
	// 200
	DBNodeStorage *string `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	// Indicates whether disk encryption is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DiskEncryption *bool `json:"DiskEncryption,omitempty" xml:"DiskEncryption,omitempty"`
	// The number of elastic I/O units (EIUs). For more information, see [Elasticity of the storage layer](https://help.aliyun.com/document_detail/189505.html).
	//
	// example:
	//
	// 0
	ElasticIOResource *string `json:"ElasticIOResource,omitempty" xml:"ElasticIOResource,omitempty"`
	// Specifies whether to enable SSL encryption. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableSSL *bool `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	ExecutorCount *string `json:"ExecutorCount,omitempty" xml:"ExecutorCount,omitempty"`
	// The Key Management Service (KMS) ID that is used for disk encryption. This parameter takes effect only when DiskEncryption is set to true.
	//
	// example:
	//
	// xxxxxxxx-xxxx-xxxx-xxxx-xxxx
	KmsId *string `json:"KmsId,omitempty" xml:"KmsId,omitempty"`
	// The mode of the cluster. Valid values:
	//
	// 	- **Reserver**: the reserved mode.
	//
	// 	- **Flexible**: the elastic mode.
	//
	// example:
	//
	// Reserver
	Mode         *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription type of the subscription cluster. Valid values:
	//
	// 	- **Year**: subscription on a yearly basis
	//
	// 	- **Month**: subscription on a monthly basis
	//
	// >  This parameter is required if the PayType parameter is set to Prepaid.
	//
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the cluster belongs.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	SourceDBInstanceName *string `json:"SourceDBInstanceName,omitempty" xml:"SourceDBInstanceName,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The tags to add to the cluster.
	Tag []*CreateDBClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The subscription period of the subscription cluster.
	//
	// 	- Valid values when Period is set to Year: 1, 2, and 3 (integer)
	//
	// 	- Valid values when Period is set to Month: 1 to 9 (integer)
	//
	// > 	- This parameter is required if the PayType parameter is set to Prepaid.
	//
	// > 	- Longer subscription periods offer more savings. Purchasing a cluster for one year is more cost-effective than purchasing the cluster for 10 or 11 months.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The virtual private cloud (VPC) ID of the cluster.
	//
	// example:
	//
	// vpc-bp1at5ze0t5u3xtqn****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the cluster.
	//
	// example:
	//
	// vsw-bp1aadw9k19x6cis9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateDBClusterRequest) GetBackupSetID() *string {
	return s.BackupSetID
}

func (s *CreateDBClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBClusterRequest) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *CreateDBClusterRequest) GetDBClusterCategory() *string {
	return s.DBClusterCategory
}

func (s *CreateDBClusterRequest) GetDBClusterClass() *string {
	return s.DBClusterClass
}

func (s *CreateDBClusterRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *CreateDBClusterRequest) GetDBClusterNetworkType() *string {
	return s.DBClusterNetworkType
}

func (s *CreateDBClusterRequest) GetDBClusterVersion() *string {
	return s.DBClusterVersion
}

func (s *CreateDBClusterRequest) GetDBNodeGroupCount() *string {
	return s.DBNodeGroupCount
}

func (s *CreateDBClusterRequest) GetDBNodeStorage() *string {
	return s.DBNodeStorage
}

func (s *CreateDBClusterRequest) GetDiskEncryption() *bool {
	return s.DiskEncryption
}

func (s *CreateDBClusterRequest) GetElasticIOResource() *string {
	return s.ElasticIOResource
}

func (s *CreateDBClusterRequest) GetEnableSSL() *bool {
	return s.EnableSSL
}

func (s *CreateDBClusterRequest) GetExecutorCount() *string {
	return s.ExecutorCount
}

func (s *CreateDBClusterRequest) GetKmsId() *string {
	return s.KmsId
}

func (s *CreateDBClusterRequest) GetMode() *string {
	return s.Mode
}

func (s *CreateDBClusterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBClusterRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDBClusterRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDBClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBClusterRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *CreateDBClusterRequest) GetRestoreType() *string {
	return s.RestoreType
}

func (s *CreateDBClusterRequest) GetSourceDBInstanceName() *string {
	return s.SourceDBInstanceName
}

func (s *CreateDBClusterRequest) GetStorageResource() *string {
	return s.StorageResource
}

func (s *CreateDBClusterRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateDBClusterRequest) GetTag() []*CreateDBClusterRequestTag {
	return s.Tag
}

func (s *CreateDBClusterRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateDBClusterRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDBClusterRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBClusterRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBClusterRequest) SetBackupSetID(v string) *CreateDBClusterRequest {
	s.BackupSetID = &v
	return s
}

func (s *CreateDBClusterRequest) SetClientToken(v string) *CreateDBClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBClusterRequest) SetComputeResource(v string) *CreateDBClusterRequest {
	s.ComputeResource = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterCategory(v string) *CreateDBClusterRequest {
	s.DBClusterCategory = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterClass(v string) *CreateDBClusterRequest {
	s.DBClusterClass = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterDescription(v string) *CreateDBClusterRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterNetworkType(v string) *CreateDBClusterRequest {
	s.DBClusterNetworkType = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterVersion(v string) *CreateDBClusterRequest {
	s.DBClusterVersion = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBNodeGroupCount(v string) *CreateDBClusterRequest {
	s.DBNodeGroupCount = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBNodeStorage(v string) *CreateDBClusterRequest {
	s.DBNodeStorage = &v
	return s
}

func (s *CreateDBClusterRequest) SetDiskEncryption(v bool) *CreateDBClusterRequest {
	s.DiskEncryption = &v
	return s
}

func (s *CreateDBClusterRequest) SetElasticIOResource(v string) *CreateDBClusterRequest {
	s.ElasticIOResource = &v
	return s
}

func (s *CreateDBClusterRequest) SetEnableSSL(v bool) *CreateDBClusterRequest {
	s.EnableSSL = &v
	return s
}

func (s *CreateDBClusterRequest) SetExecutorCount(v string) *CreateDBClusterRequest {
	s.ExecutorCount = &v
	return s
}

func (s *CreateDBClusterRequest) SetKmsId(v string) *CreateDBClusterRequest {
	s.KmsId = &v
	return s
}

func (s *CreateDBClusterRequest) SetMode(v string) *CreateDBClusterRequest {
	s.Mode = &v
	return s
}

func (s *CreateDBClusterRequest) SetOwnerAccount(v string) *CreateDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBClusterRequest) SetOwnerId(v int64) *CreateDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBClusterRequest) SetPayType(v string) *CreateDBClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBClusterRequest) SetPeriod(v string) *CreateDBClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateDBClusterRequest) SetRegionId(v string) *CreateDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceGroupId(v string) *CreateDBClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceOwnerAccount(v string) *CreateDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceOwnerId(v int64) *CreateDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBClusterRequest) SetRestoreTime(v string) *CreateDBClusterRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateDBClusterRequest) SetRestoreType(v string) *CreateDBClusterRequest {
	s.RestoreType = &v
	return s
}

func (s *CreateDBClusterRequest) SetSourceDBInstanceName(v string) *CreateDBClusterRequest {
	s.SourceDBInstanceName = &v
	return s
}

func (s *CreateDBClusterRequest) SetStorageResource(v string) *CreateDBClusterRequest {
	s.StorageResource = &v
	return s
}

func (s *CreateDBClusterRequest) SetStorageType(v string) *CreateDBClusterRequest {
	s.StorageType = &v
	return s
}

func (s *CreateDBClusterRequest) SetTag(v []*CreateDBClusterRequestTag) *CreateDBClusterRequest {
	s.Tag = v
	return s
}

func (s *CreateDBClusterRequest) SetUsedTime(v string) *CreateDBClusterRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBClusterRequest) SetVPCId(v string) *CreateDBClusterRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBClusterRequest) SetVSwitchId(v string) *CreateDBClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBClusterRequest) SetZoneId(v string) *CreateDBClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBClusterRequest) Validate() error {
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

type CreateDBClusterRequestTag struct {
	// The key of tag N to add to the cluster. You can use tags to filter clusters. Valid values of N: 1 to 20. The values that you specify for N must be unique and consecutive integers that start from 1. Each value of `Tag.N.Key` is paired with a value of `Tag.N.Value`.
	//
	// >  The tag key can be up to 64 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the cluster. You can use tags to filter clusters. Valid values of N: 1 to 20. The values that you specify for N must be unique and consecutive integers that start from 1. Each value of `Tag.N.Key` is paired with a value of `Tag.N.Value`.
	//
	// >  The tag value can be up to 64 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// example:
	//
	// test1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDBClusterRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDBClusterRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDBClusterRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDBClusterRequestTag) SetKey(v string) *CreateDBClusterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDBClusterRequestTag) SetValue(v string) *CreateDBClusterRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDBClusterRequestTag) Validate() error {
	return dara.Validate(s)
}
