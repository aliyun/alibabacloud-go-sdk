// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *CreateDBInstanceRequest
	GetAutoRenew() *bool
	SetBackupSetID(v string) *CreateDBInstanceRequest
	GetBackupSetID() *string
	SetClientToken(v string) *CreateDBInstanceRequest
	GetClientToken() *string
	SetDBClusterCategory(v string) *CreateDBInstanceRequest
	GetDBClusterCategory() *string
	SetDBClusterClass(v string) *CreateDBInstanceRequest
	GetDBClusterClass() *string
	SetDBClusterDescription(v string) *CreateDBInstanceRequest
	GetDBClusterDescription() *string
	SetDBClusterNetworkType(v string) *CreateDBInstanceRequest
	GetDBClusterNetworkType() *string
	SetDBClusterVersion(v string) *CreateDBInstanceRequest
	GetDBClusterVersion() *string
	SetDBNodeGroupCount(v string) *CreateDBInstanceRequest
	GetDBNodeGroupCount() *string
	SetDBNodeStorage(v string) *CreateDBInstanceRequest
	GetDBNodeStorage() *string
	SetDbNodeStorageType(v string) *CreateDBInstanceRequest
	GetDbNodeStorageType() *string
	SetEncryptionKey(v string) *CreateDBInstanceRequest
	GetEncryptionKey() *string
	SetEncryptionType(v string) *CreateDBInstanceRequest
	GetEncryptionType() *string
	SetOwnerAccount(v string) *CreateDBInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDBInstanceRequest
	GetOwnerId() *int64
	SetPayType(v string) *CreateDBInstanceRequest
	GetPayType() *string
	SetPeriod(v string) *CreateDBInstanceRequest
	GetPeriod() *string
	SetRegionId(v string) *CreateDBInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBInstanceRequest
	GetResourceOwnerId() *int64
	SetSourceDBClusterId(v string) *CreateDBInstanceRequest
	GetSourceDBClusterId() *string
	SetUsedTime(v string) *CreateDBInstanceRequest
	GetUsedTime() *string
	SetVPCId(v string) *CreateDBInstanceRequest
	GetVPCId() *string
	SetVSwitchBak(v string) *CreateDBInstanceRequest
	GetVSwitchBak() *string
	SetVSwitchBak2(v string) *CreateDBInstanceRequest
	GetVSwitchBak2() *string
	SetVSwitchId(v string) *CreateDBInstanceRequest
	GetVSwitchId() *string
	SetZondIdBak2(v string) *CreateDBInstanceRequest
	GetZondIdBak2() *string
	SetZoneId(v string) *CreateDBInstanceRequest
	GetZoneId() *string
	SetZoneIdBak(v string) *CreateDBInstanceRequest
	GetZoneIdBak() *string
}

type CreateDBInstanceRequest struct {
	// Specifies whether to enable auto-renewal.
	//
	// >  This parameter is valid only if the value of PayType is set to Prepaid.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the backup set. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/360339.html) operation to query the backup sets.
	//
	// >  If you want to restore the data of an ApsaraDB for ClickHouse cluster, this parameter is required.
	//
	// example:
	//
	// b-12af23adsf
	BackupSetID *string `json:"BackupSetID,omitempty" xml:"BackupSetID,omitempty"`
	// The client token that is used to ensure the idempotence of the request. The value is a string and can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// AB
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// 	- **Basic**: Single-replica Edition
	//
	// 	- **HighAvailability**: Double-replica Edition
	//
	// This parameter is required.
	//
	// example:
	//
	// Basic
	DBClusterCategory *string `json:"DBClusterCategory,omitempty" xml:"DBClusterCategory,omitempty"`
	// The specifications of the cluster.
	//
	// 	- Valid values for a Single-replica Edition cluster:
	//
	//     	- **S8**: 8 cores and 32 GB of memory
	//
	//     	- **S16**: 16 cores and 64 GB of memory
	//
	//     	- **S32**: 32 cores and 128 GB of memory
	//
	//     	- **S64**: 64 cores and 256 GB of memory
	//
	//     	- **S104**: 104 cores and 384 GB of memory
	//
	// 	- Valid values for a Double-replica Edition cluster:
	//
	//     	- **C8**: 8 cores and 32 GB of memory
	//
	//     	- **C16**: 16 cores and 64 GB of memory
	//
	//     	- **C32**: 32 cores and 128 GB of memory
	//
	//     	- **C64**: 64 cores and 256 GB of memory
	//
	//     	- **C104**: 104 cores and 384 GB of memory
	//
	// This parameter is required.
	//
	// example:
	//
	// S8
	DBClusterClass *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The network type of the cluster. Only Virtual Private Cloud (VPC) is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The kernel version. Valid values:
	//
	// 	- **21.8.10.19**
	//
	// 	- **22.8.5.29**
	//
	// This parameter is required.
	//
	// example:
	//
	// 21.8.10.19
	DBClusterVersion *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	// The number of nodes.
	//
	// 	- Valid values when the cluster is of Single-replica Edition: 1 to 48.
	//
	// 	- Valid values when the cluster is of Double-replica Edition: 1 to 24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DBNodeGroupCount *string `json:"DBNodeGroupCount,omitempty" xml:"DBNodeGroupCount,omitempty"`
	// The storage capacity of a single node. Valid values: 100 to 32000. Unit: GB.
	//
	// >  This value is a multiple of 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	DBNodeStorage *string `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	// The storage type of the cluster. Valid values:
	//
	// 	- **CloudESSD**: The cluster uses an enhanced SSD (ESSD) of performance level 1 (PL1).
	//
	// 	- **CloudESSD_PL2**: The cluster uses an ESSD of PL2.
	//
	// 	- **CloudESSD_PL3**: The cluster uses an ESSD of PL3.
	//
	// 	- **CloudEfficiency**: The cluster uses an ultra disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudESSD_PL2
	DbNodeStorageType *string `json:"DbNodeStorageType,omitempty" xml:"DbNodeStorageType,omitempty"`
	// You must specify this parameter when EncryptionType is set to CloudDisk.
	//
	// The ID of the key that is used to encrypt data on disks. You can obtain the ID of the key from the Key Management Service (KMS) console. You can also create a key.
	//
	// >  If EncryptionType is empty, you do not need to specify this parameter.
	//
	// example:
	//
	// 0d2470df-da7b-4786-b981-9a164dae****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The encryption type. Set the value to **CloudDisk**, which indicates that only disk encryption is supported.
	//
	// >  If this parameter is not specified, data is not encrypted.
	//
	// example:
	//
	// CloudDisk
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// 	- **Postpaid**: The cluster uses the pay-as-you-go billing method.
	//
	// 	- **Prepaid**: The cluster uses the subscription billing method.
	//
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration for the cluster. Valid values:
	//
	// >  This parameter is required only when PayType is set to Prepaid.
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
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
	// The ID of the source cluster. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query backup set IDs.
	//
	// >  If you want to restore the data of an ApsaraDB for ClickHouse cluster, this parameter is required.
	//
	// example:
	//
	// cc-bp1lxbo89u950****
	SourceDBClusterId *string `json:"SourceDBClusterId,omitempty" xml:"SourceDBClusterId,omitempty"`
	// The subscription duration of the subscription cluster.
	//
	// >  This parameter is required only when PayType is set to Prepaid.
	//
	// 	- Valid values when Period is set to Year: 1 to 3 (integer)
	//
	// 	- Valid values when Period is set to Month: 1 to 9 (integer)
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The VPC ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch in the secondary zone for the VPC.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchBak *string `json:"VSwitchBak,omitempty" xml:"VSwitchBak,omitempty"`
	// The vSwitch in secondary zone 2 for the VPC.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchBak2 *string `json:"VSwitchBak2,omitempty" xml:"VSwitchBak2,omitempty"`
	// The vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The secondary zone 2 of the instance.
	//
	// example:
	//
	// cn-hangzhou-j
	ZondIdBak2 *string `json:"ZondIdBak2,omitempty" xml:"ZondIdBak2,omitempty"`
	// The zone ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The secondary zone.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneIdBak *string `json:"ZoneIdBak,omitempty" xml:"ZoneIdBak,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateDBInstanceRequest) GetBackupSetID() *string {
	return s.BackupSetID
}

func (s *CreateDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceRequest) GetDBClusterCategory() *string {
	return s.DBClusterCategory
}

func (s *CreateDBInstanceRequest) GetDBClusterClass() *string {
	return s.DBClusterClass
}

func (s *CreateDBInstanceRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *CreateDBInstanceRequest) GetDBClusterNetworkType() *string {
	return s.DBClusterNetworkType
}

func (s *CreateDBInstanceRequest) GetDBClusterVersion() *string {
	return s.DBClusterVersion
}

func (s *CreateDBInstanceRequest) GetDBNodeGroupCount() *string {
	return s.DBNodeGroupCount
}

func (s *CreateDBInstanceRequest) GetDBNodeStorage() *string {
	return s.DBNodeStorage
}

func (s *CreateDBInstanceRequest) GetDbNodeStorageType() *string {
	return s.DbNodeStorageType
}

func (s *CreateDBInstanceRequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *CreateDBInstanceRequest) GetEncryptionType() *string {
	return s.EncryptionType
}

func (s *CreateDBInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDBInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBInstanceRequest) GetSourceDBClusterId() *string {
	return s.SourceDBClusterId
}

func (s *CreateDBInstanceRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateDBInstanceRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDBInstanceRequest) GetVSwitchBak() *string {
	return s.VSwitchBak
}

func (s *CreateDBInstanceRequest) GetVSwitchBak2() *string {
	return s.VSwitchBak2
}

func (s *CreateDBInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBInstanceRequest) GetZondIdBak2() *string {
	return s.ZondIdBak2
}

func (s *CreateDBInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBInstanceRequest) GetZoneIdBak() *string {
	return s.ZoneIdBak
}

func (s *CreateDBInstanceRequest) SetAutoRenew(v bool) *CreateDBInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDBInstanceRequest) SetBackupSetID(v string) *CreateDBInstanceRequest {
	s.BackupSetID = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBClusterCategory(v string) *CreateDBInstanceRequest {
	s.DBClusterCategory = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBClusterClass(v string) *CreateDBInstanceRequest {
	s.DBClusterClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBClusterDescription(v string) *CreateDBInstanceRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBClusterNetworkType(v string) *CreateDBInstanceRequest {
	s.DBClusterNetworkType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBClusterVersion(v string) *CreateDBInstanceRequest {
	s.DBClusterVersion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBNodeGroupCount(v string) *CreateDBInstanceRequest {
	s.DBNodeGroupCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBNodeStorage(v string) *CreateDBInstanceRequest {
	s.DBNodeStorage = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDbNodeStorageType(v string) *CreateDBInstanceRequest {
	s.DbNodeStorageType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEncryptionKey(v string) *CreateDBInstanceRequest {
	s.EncryptionKey = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEncryptionType(v string) *CreateDBInstanceRequest {
	s.EncryptionType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetOwnerAccount(v string) *CreateDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetOwnerId(v int64) *CreateDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPayType(v string) *CreateDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPeriod(v string) *CreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRegionId(v string) *CreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceGroupId(v string) *CreateDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceOwnerAccount(v string) *CreateDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceOwnerId(v int64) *CreateDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSourceDBClusterId(v string) *CreateDBInstanceRequest {
	s.SourceDBClusterId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetUsedTime(v string) *CreateDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVPCId(v string) *CreateDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchBak(v string) *CreateDBInstanceRequest {
	s.VSwitchBak = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchBak2(v string) *CreateDBInstanceRequest {
	s.VSwitchBak2 = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchId(v string) *CreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZondIdBak2(v string) *CreateDBInstanceRequest {
	s.ZondIdBak2 = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneIdBak(v string) *CreateDBInstanceRequest {
	s.ZoneIdBak = &v
	return s
}

func (s *CreateDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
