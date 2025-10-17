// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The ID of the backup set that you want to use to restore data.
	//
	// >  You can call the [DescribeBackups](https://help.aliyun.com/document_detail/612318.html) operation to query the backup sets of the cluster.
	//
	// example:
	//
	// 1880808684
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The region ID of the source cluster.
	//
	// >  This parameter must be specified for cloning clusters across regions.
	//
	// example:
	//
	// cn-beijing
	CloneSourceRegionId *string `json:"CloneSourceRegionId,omitempty" xml:"CloneSourceRegionId,omitempty"`
	// The amount of reserved computing resources. Valid values: 0ACU to 4096ACU. The value must be in increments of 16ACU. Each ACU is approximately equal to 1 core and 4 GB memory.
	//
	// >  This parameter must be specified with a unit.
	//
	// example:
	//
	// 16ACU
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The description of the cluster.
	//
	// 	- The description cannot start with `http://` or `https://`.
	//
	// 	- The description must be 2 to 256 characters in length
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The network type of the cluster. Set the value to **VPC**.
	//
	// example:
	//
	// VPC
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The version of the cluster. Set the value to **5.0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.0
	DBClusterVersion *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	// Specifies whether to enable disk encryption.
	//
	// example:
	//
	// false
	DiskEncryption *bool `json:"DiskEncryption,omitempty" xml:"DiskEncryption,omitempty"`
	// Specifies whether to allocate all reserved computing resources to the user_default resource group. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableDefaultResourcePool *bool `json:"EnableDefaultResourcePool,omitempty" xml:"EnableDefaultResourcePool,omitempty"`
	EnableSSL                 *bool `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// The ID of the key that is used to encrypt disk data.
	//
	// >  This parameter must be specified only when disk encryption is enabled.
	//
	// example:
	//
	// e1935511-cf88-1123-a0f8-1be8d251****
	KmsId *string `json:"KmsId,omitempty" xml:"KmsId,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription type of the subscription cluster. Valid values:
	//
	// 	- **Year**: subscription on a yearly basis.
	//
	// 	- **Month**: subscription on a monthly basis.
	//
	// >  This parameter must be specified when PayType is set to Prepaid.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The product form of the cluster. Valid values:
	//
	// 	- **IntegrationForm**: integrated.
	//
	// 	- **LegacyForm**: Data Lakehouse Edition.
	//
	// example:
	//
	// LegacyForm
	ProductForm *string `json:"ProductForm,omitempty" xml:"ProductForm,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// 	- **BasicVersion**: Basic Edition.
	//
	// 	- **EnterpriseVersion**: Enterprise Edition.
	//
	// >  This parameter must be specified only when ProductForm is set to IntegrationForm.
	//
	// example:
	//
	// BasicVersion
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
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
	// The number of reserved resource nodes.
	//
	// 	- For Enterprise Edition, the default value is 3 and the step size is 3.
	//
	// 	- For Basic Edition, the default value is 1.
	//
	// >  This parameter must be specified only when ProductForm is set to IntegrationForm.
	//
	// example:
	//
	// 3
	ReservedNodeCount *int32 `json:"ReservedNodeCount,omitempty" xml:"ReservedNodeCount,omitempty"`
	// The specifications of reserved resource nodes. Unit: ACUs.
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
	// The method that you want to use to restore data. Valid values:
	//
	// 	- **backup**: restores data from a backup set. You must also specify the **BackupSetId*	- and **SourceDBClusterId*	- parameters.
	//
	// 	- **timepoint**: restores data to a point in time. You must also specify the **RestoreToTime*	- and **SourceDBClusterId*	- parameters.
	//
	// example:
	//
	// backup
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The ID of the secondary vSwitch.
	//
	// >  You cannot set this parameter to a value that is the same as that of the VSwitchId parameter.
	//
	// example:
	//
	// vsw-bp1aadw9k19x451gx****
	SecondaryVSwitchId *string `json:"SecondaryVSwitchId,omitempty" xml:"SecondaryVSwitchId,omitempty"`
	// The ID of the secondary zone.
	//
	// >  You cannot set this parameter to a value that is the same as that of the ZoneId parameter.
	//
	// example:
	//
	// cn-beijing-h
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	// The ID of the source AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// example:
	//
	// amv-bp1r053byu48p****
	SourceDbClusterId *string `json:"SourceDbClusterId,omitempty" xml:"SourceDbClusterId,omitempty"`
	// The amount of reserved storage resources. Valid values: 0ACU to 2064ACU. The value must be in increments of 24ACU. Each ACU is approximately equal to 1 core and 4 GB memory.
	//
	// >  This parameter must be specified with a unit.
	//
	// example:
	//
	// 24ACU
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1at5ze0t5u3xtqn****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1aadw9k19x6cis9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent zone list.
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
	// The key of tag N to add to the cluster. You can use tags to filter clusters. Valid values of N: 1 to 20. The values that you specify for N must be unique and consecutive integers that start from 1. Each value of `Tag.N.Key` is paired with a value of `Tag.N.Value`.
	//
	// >  The tag key can be up to 64 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// example:
	//
	// testkey1
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
