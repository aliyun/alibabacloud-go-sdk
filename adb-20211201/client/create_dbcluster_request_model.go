// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAINodeNumber(v int32) *CreateDBClusterRequest
	GetAINodeNumber() *int32
	SetAINodeSpec(v string) *CreateDBClusterRequest
	GetAINodeSpec() *string
	SetBackupSetId(v string) *CreateDBClusterRequest
	GetBackupSetId() *string
	SetCloneSourceRegionId(v string) *CreateDBClusterRequest
	GetCloneSourceRegionId() *string
	SetComputeResource(v string) *CreateDBClusterRequest
	GetComputeResource() *string
	SetDBClusterDescription(v string) *CreateDBClusterRequest
	GetDBClusterDescription() *string
	SetDBClusterNetworkType(v string) *CreateDBClusterRequest
	GetDBClusterNetworkType() *string
	SetDBClusterVersion(v string) *CreateDBClusterRequest
	GetDBClusterVersion() *string
	SetDiskEncryption(v bool) *CreateDBClusterRequest
	GetDiskEncryption() *bool
	SetEnableDefaultResourcePool(v bool) *CreateDBClusterRequest
	GetEnableDefaultResourcePool() *bool
	SetEnableSSL(v bool) *CreateDBClusterRequest
	GetEnableSSL() *bool
	SetKmsId(v string) *CreateDBClusterRequest
	GetKmsId() *string
	SetPayType(v string) *CreateDBClusterRequest
	GetPayType() *string
	SetPeriod(v string) *CreateDBClusterRequest
	GetPeriod() *string
	SetProductForm(v string) *CreateDBClusterRequest
	GetProductForm() *string
	SetProductVersion(v string) *CreateDBClusterRequest
	GetProductVersion() *string
	SetRegionId(v string) *CreateDBClusterRequest
	GetRegionId() *string
	SetReservedNodeCount(v int32) *CreateDBClusterRequest
	GetReservedNodeCount() *int32
	SetReservedNodeSize(v string) *CreateDBClusterRequest
	GetReservedNodeSize() *string
	SetResourceGroupId(v string) *CreateDBClusterRequest
	GetResourceGroupId() *string
	SetRestoreToTime(v string) *CreateDBClusterRequest
	GetRestoreToTime() *string
	SetRestoreType(v string) *CreateDBClusterRequest
	GetRestoreType() *string
	SetSecondaryVSwitchId(v string) *CreateDBClusterRequest
	GetSecondaryVSwitchId() *string
	SetSecondaryZoneId(v string) *CreateDBClusterRequest
	GetSecondaryZoneId() *string
	SetSourceDbClusterId(v string) *CreateDBClusterRequest
	GetSourceDbClusterId() *string
	SetStorageResource(v string) *CreateDBClusterRequest
	GetStorageResource() *string
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
	// example:
	//
	// 1
	AINodeNumber *int32 `json:"AINodeNumber,omitempty" xml:"AINodeNumber,omitempty"`
	// example:
	//
	// ADB.MLPlus.4
	AINodeSpec *string `json:"AINodeSpec,omitempty" xml:"AINodeSpec,omitempty"`
	// The ID of the backup set used for restoration from a backup set.
	//
	// > You can call the [DescribeBackups](https://help.aliyun.com/document_detail/612318.html) operation to query the backup list of the cluster.
	//
	// example:
	//
	// 1880808684
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The region of the source cluster.
	//
	// > This parameter is required for cross-region cloning.
	//
	// example:
	//
	// cn-beijing
	CloneSourceRegionId *string `json:"CloneSourceRegionId,omitempty" xml:"CloneSourceRegionId,omitempty"`
	// The compute reserved resources. Valid values: 0 ACU to 4096 ACU, in increments of 16. 1 ACU is approximately equivalent to 1 core and 4 GB of memory.
	//
	// > Include the unit when specifying this parameter.
	//
	// example:
	//
	// 16ACU
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The description of the cluster.
	//
	// - The description cannot start with `http://` or `https://`.
	//
	// - The description must be 2 to 256 characters in length.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The network type of the cluster. Only **VPC*	- (Virtual Private Cloud) is supported.
	//
	// example:
	//
	// VPC
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The version of the Data Lakehouse Edition cluster. Valid values: **5.0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.0
	DBClusterVersion *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	// Specifies whether to enable cloud disk encryption.
	//
	// example:
	//
	// false
	DiskEncryption *bool `json:"DiskEncryption,omitempty" xml:"DiskEncryption,omitempty"`
	// Specifies whether to allocate all compute reserved resources to the default resource group (user_default). Valid values:
	//
	// - **true*	- (default): All compute reserved resources are allocated to the default resource group.
	//
	// - **false**: Not all compute reserved resources are allocated to the default resource group.
	//
	// example:
	//
	// true
	EnableDefaultResourcePool *bool `json:"EnableDefaultResourcePool,omitempty" xml:"EnableDefaultResourcePool,omitempty"`
	// Specifies whether to enable SSL encryption. Valid values:
	//
	// - **true**: SSL encryption is enabled.
	//
	// - **false**: SSL encryption is disabled.
	//
	// example:
	//
	// false
	EnableSSL *bool `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// The ID of the key used to encrypt cloud disk data.
	//
	// > This parameter is used only when cloud disk encryption is enabled for the AnalyticDB for MySQL cluster.
	//
	// example:
	//
	// e1935511-cf88-1123-a0f8-1be8d251****
	KmsId *string `json:"KmsId,omitempty" xml:"KmsId,omitempty"`
	// The billing method. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription type of the subscription cluster. Valid values:
	//
	// - **Year**: subscription on a yearly basis.
	//
	// - **Month**: subscription on a monthly basis.
	//
	// > This parameter is required when PayType is set to Prepaid.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The product form. Valid values:
	//
	// - **IntegrationForm**: integrated form.
	//
	// - **LegacyForm**: Data Lakehouse Edition.
	//
	// example:
	//
	// LegacyForm
	ProductForm *string `json:"ProductForm,omitempty" xml:"ProductForm,omitempty"`
	// The product version. Valid values:
	//
	// - **BasicVersion**: Basic Edition.
	//
	// - **EnterpriseVersion**: Enterprise Edition.
	//
	// > This parameter is required only when ProductForm is set to IntegrationForm.
	//
	// example:
	//
	// BasicVersion
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the region ID of a specific Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of reserved nodes.
	//
	// - Enterprise Edition uses 3 nodes by default, in increments of 3.
	//
	// - Basic Edition uses 1 node by default.
	//
	// > This parameter is required only when ProductForm is set to IntegrationForm.
	//
	// example:
	//
	// 3
	ReservedNodeCount *int32 `json:"ReservedNodeCount,omitempty" xml:"ReservedNodeCount,omitempty"`
	// The node specifications of reserved nodes, in ACUs.
	//
	// example:
	//
	// 8ACU
	ReservedNodeSize *string `json:"ReservedNodeSize,omitempty" xml:"ReservedNodeSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The point in time to which you want to restore data from the backup set.
	//
	// example:
	//
	// 2023-09-20T03:13:56Z
	RestoreToTime *string `json:"RestoreToTime,omitempty" xml:"RestoreToTime,omitempty"`
	// The restoration method. Valid values:
	//
	// 	- **backup**: restores data from a backup set. You must also specify the **BackupSetId*	- and **SourceDBClusterId*	- parameters.
	//
	// 	- **timepoint**: restores data to a specific point in time. You must also specify the **RestoreToTime*	- and **SourceDBClusterId*	- parameters.
	//
	// example:
	//
	// backup
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The secondary vSwitch ID.
	//
	// > The value of this parameter must be different from the value of the VSwitchId parameter.
	//
	// example:
	//
	// vsw-bp1aadw9k19x451gx****
	SecondaryVSwitchId *string `json:"SecondaryVSwitchId,omitempty" xml:"SecondaryVSwitchId,omitempty"`
	// The secondary zone ID.
	//
	// > The value of this parameter must be different from the value of the ZoneId parameter.
	//
	// example:
	//
	// cn-beijing-h
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	// The instance ID of the source AnalyticDB for MySQL Data Warehouse Edition cluster. If this parameter is specified, the Data Lakehouse Edition cluster is used to recover from the Data Warehouse Edition cluster.
	//
	// example:
	//
	// amv-bp1r053byu48p****
	SourceDbClusterId *string `json:"SourceDbClusterId,omitempty" xml:"SourceDbClusterId,omitempty"`
	// The storage reserved resources. Valid values: 0 ACU to 2064 ACU, in increments of 24. 1 ACU is approximately equivalent to 1 core and 4 GB of memory.
	//
	// > Include the unit when specifying this parameter.
	//
	// example:
	//
	// 24ACU
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	// The list of tags.
	Tag []*CreateDBClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The subscription duration of the subscription cluster. Valid values:
	//
	// - When **Period*	- is set to Year, the value of UsedTime ranges from 1 to 3 (integer).
	//
	// - When **Period*	- is set to Month, the value of UsedTime ranges from 1 to 9 (integer).
	//
	// > This parameter is required when PayType is set to **Prepaid**.
	//
	// example:
	//
	// 3
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1at5ze0t5u3xtqn****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1aadw9k19x6cis9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the zone ID of a specific Data Lakehouse Edition cluster.
	//
	// This parameter is required.
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

func (s *CreateDBClusterRequest) GetAINodeNumber() *int32 {
	return s.AINodeNumber
}

func (s *CreateDBClusterRequest) GetAINodeSpec() *string {
	return s.AINodeSpec
}

func (s *CreateDBClusterRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *CreateDBClusterRequest) GetCloneSourceRegionId() *string {
	return s.CloneSourceRegionId
}

func (s *CreateDBClusterRequest) GetComputeResource() *string {
	return s.ComputeResource
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

func (s *CreateDBClusterRequest) GetDiskEncryption() *bool {
	return s.DiskEncryption
}

func (s *CreateDBClusterRequest) GetEnableDefaultResourcePool() *bool {
	return s.EnableDefaultResourcePool
}

func (s *CreateDBClusterRequest) GetEnableSSL() *bool {
	return s.EnableSSL
}

func (s *CreateDBClusterRequest) GetKmsId() *string {
	return s.KmsId
}

func (s *CreateDBClusterRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDBClusterRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDBClusterRequest) GetProductForm() *string {
	return s.ProductForm
}

func (s *CreateDBClusterRequest) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *CreateDBClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBClusterRequest) GetReservedNodeCount() *int32 {
	return s.ReservedNodeCount
}

func (s *CreateDBClusterRequest) GetReservedNodeSize() *string {
	return s.ReservedNodeSize
}

func (s *CreateDBClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBClusterRequest) GetRestoreToTime() *string {
	return s.RestoreToTime
}

func (s *CreateDBClusterRequest) GetRestoreType() *string {
	return s.RestoreType
}

func (s *CreateDBClusterRequest) GetSecondaryVSwitchId() *string {
	return s.SecondaryVSwitchId
}

func (s *CreateDBClusterRequest) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *CreateDBClusterRequest) GetSourceDbClusterId() *string {
	return s.SourceDbClusterId
}

func (s *CreateDBClusterRequest) GetStorageResource() *string {
	return s.StorageResource
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

func (s *CreateDBClusterRequest) SetAINodeNumber(v int32) *CreateDBClusterRequest {
	s.AINodeNumber = &v
	return s
}

func (s *CreateDBClusterRequest) SetAINodeSpec(v string) *CreateDBClusterRequest {
	s.AINodeSpec = &v
	return s
}

func (s *CreateDBClusterRequest) SetBackupSetId(v string) *CreateDBClusterRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateDBClusterRequest) SetCloneSourceRegionId(v string) *CreateDBClusterRequest {
	s.CloneSourceRegionId = &v
	return s
}

func (s *CreateDBClusterRequest) SetComputeResource(v string) *CreateDBClusterRequest {
	s.ComputeResource = &v
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

func (s *CreateDBClusterRequest) SetDiskEncryption(v bool) *CreateDBClusterRequest {
	s.DiskEncryption = &v
	return s
}

func (s *CreateDBClusterRequest) SetEnableDefaultResourcePool(v bool) *CreateDBClusterRequest {
	s.EnableDefaultResourcePool = &v
	return s
}

func (s *CreateDBClusterRequest) SetEnableSSL(v bool) *CreateDBClusterRequest {
	s.EnableSSL = &v
	return s
}

func (s *CreateDBClusterRequest) SetKmsId(v string) *CreateDBClusterRequest {
	s.KmsId = &v
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

func (s *CreateDBClusterRequest) SetProductForm(v string) *CreateDBClusterRequest {
	s.ProductForm = &v
	return s
}

func (s *CreateDBClusterRequest) SetProductVersion(v string) *CreateDBClusterRequest {
	s.ProductVersion = &v
	return s
}

func (s *CreateDBClusterRequest) SetRegionId(v string) *CreateDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBClusterRequest) SetReservedNodeCount(v int32) *CreateDBClusterRequest {
	s.ReservedNodeCount = &v
	return s
}

func (s *CreateDBClusterRequest) SetReservedNodeSize(v string) *CreateDBClusterRequest {
	s.ReservedNodeSize = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceGroupId(v string) *CreateDBClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBClusterRequest) SetRestoreToTime(v string) *CreateDBClusterRequest {
	s.RestoreToTime = &v
	return s
}

func (s *CreateDBClusterRequest) SetRestoreType(v string) *CreateDBClusterRequest {
	s.RestoreType = &v
	return s
}

func (s *CreateDBClusterRequest) SetSecondaryVSwitchId(v string) *CreateDBClusterRequest {
	s.SecondaryVSwitchId = &v
	return s
}

func (s *CreateDBClusterRequest) SetSecondaryZoneId(v string) *CreateDBClusterRequest {
	s.SecondaryZoneId = &v
	return s
}

func (s *CreateDBClusterRequest) SetSourceDbClusterId(v string) *CreateDBClusterRequest {
	s.SourceDbClusterId = &v
	return s
}

func (s *CreateDBClusterRequest) SetStorageResource(v string) *CreateDBClusterRequest {
	s.StorageResource = &v
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
	// The tag key. You can use tags to filter the cluster list. You can specify up to 20 tag pairs. The value of N for each tag pair must be unique and must be a consecutive integer that starts from 1. The value of `Tag.N.Key` corresponds to the value of `Tag.N.Value`.
	//
	// > The tag key can be up to 64 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can use tags to filter the cluster list. You can specify up to 20 tag pairs. The value of N for each tag pair must be unique and must be a consecutive integer that starts from 1. The value of `Tag.N.Key` corresponds to the value of `Tag.N.Value`.
	//
	// > The tag value can be up to 64 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
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
