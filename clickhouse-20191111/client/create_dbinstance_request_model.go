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
	SetTags(v []*CreateDBInstanceRequestTags) *CreateDBInstanceRequest
	GetTags() []*CreateDBInstanceRequestTags
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
	// >This parameter takes effect only when PayType is set to Prepaid.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The backup set ID. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/360339.html) operation to query the backup set ID.
	//
	// >This parameter is required when you restore data for an ApsaraDB for ClickHouse cluster.
	//
	// example:
	//
	// b-12af23adsf
	BackupSetID *string `json:"BackupSetID,omitempty" xml:"BackupSetID,omitempty"`
	// The client token that is used to ensure the idempotence of the request. The value is a string that contains up to 64 ASCII characters.
	//
	// example:
	//
	// AB
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The replica configuration. Valid values:
	//
	// - **Basic**: single-replica edition.
	//
	// - **HighAvailability**: double-replica edition.
	//
	// This parameter is required.
	//
	// example:
	//
	// Basic
	DBClusterCategory *string `json:"DBClusterCategory,omitempty" xml:"DBClusterCategory,omitempty"`
	// The cluster specifications.
	//
	// <props="china">
	//
	// - Single-replica edition. Valid values:
	//
	//     - **LS20**: large storage, 20 cores, 88 GB.
	//
	//     - **LS40**: large storage, 40 cores, 176 GB.
	//
	//     - **LS80**: large storage, 80 cores, 352 GB.
	//
	//     - **S8**: standard, 8 cores, 32 GB.
	//
	//     - **S16**: standard, 16 cores, 64 GB.
	//
	//     - **S32**: standard, 32 cores, 128 GB.
	//
	//     - **S64**: standard, 64 cores, 256 GB.
	//
	//     - **S80**: standard, 80 cores, 384 GB.
	//
	//     - **S104**: standard, 104 cores, 384 GB.
	//
	// - Double-replica edition. Valid values:
	//
	//     - **LC20**: large storage, 20 cores, 88 GB.
	//
	//     - **LC40**: large storage, 40 cores, 176 GB.
	//
	//     - **LC80**: large storage, 80 cores, 352 GB.
	//
	//     - **C8**: standard, 8 cores, 32 GB.
	//
	//     - **C16**: standard, 16 cores, 64 GB.
	//
	//     - **C32**: standard, 32 cores, 128 GB.
	//
	//     - **C64**: standard, 64 cores, 256 GB.
	//
	//     - **C80**: standard, 80 cores, 384 GB.
	//
	//     - **C104**: standard, 104 cores, 384 GB.
	//
	//
	// <props="intl">
	//
	// - Single-replica edition. Valid values:
	//
	//   - **S8**: 8 cores, 32 GB.
	//
	//   - **S16**: 16 cores, 64 GB.
	//
	//   - **S32**: 32 cores, 128 GB.
	//
	//   - **S64**: 64 cores, 256 GB.
	//
	//   - **S104**: 104 cores, 384 GB.
	//
	// - Double-replica edition. Valid values:
	//
	//   - **C8**: 8 cores, 32 GB.
	//
	//   - **C16**: 16 cores, 64 GB.
	//
	//   - **C32**: 32 cores, 128 GB.
	//
	//   - **C64**: 64 cores, 256 GB.
	//
	//   - **C104**: 104 cores, 384 GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// S8
	DBClusterClass *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The network type. Only VPC is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The kernel version. Valid values:
	//
	// - **21.8.10.19**
	//
	// - **22.8.5.29**
	//
	// This parameter is required.
	//
	// example:
	//
	// 22.8.5.29
	DBClusterVersion *string `json:"DBClusterVersion,omitempty" xml:"DBClusterVersion,omitempty"`
	// The number of nodes.
	//
	// - Single-replica edition: valid values: 1 to 48.
	//
	// - Double-replica edition: valid values: 1 to 24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DBNodeGroupCount *string `json:"DBNodeGroupCount,omitempty" xml:"DBNodeGroupCount,omitempty"`
	// The storage capacity per node. Valid values: 100 to 32000. Unit: GB.
	//
	// >The step size is 100 GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	DBNodeStorage *string `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	// The storage type. Valid values:
	//
	// <props="china">- **CloudESSD_PL0**: PL0 ESSD.
	//
	// - **CloudESSD**: PL1 ESSD.
	//
	// - **CloudESSD_PL2**: PL2 ESSD.
	//
	// - **CloudESSD_PL3**: PL3 ESSD.
	//
	// - **CloudEfficiency**: ultra cloud disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudESSD_PL2
	DbNodeStorageType *string `json:"DbNodeStorageType,omitempty" xml:"DbNodeStorageType,omitempty"`
	// The ID of the key used for cloud disk encryption. This parameter is required when EncryptionType is set to CloudDisk.
	//
	// You can view the key ID in the Key Management Service (KMS) console or create a key.
	//
	// >If EncryptionType is not specified, you do not need to specify this parameter.
	//
	// example:
	//
	// 0d2470df-da7b-4786-b981-9a164dae****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The encryption type. Only cloud disk encryption is supported. Set the value to **CloudDisk**.
	//
	// >If this parameter is not specified, data is not encrypted.
	//
	// example:
	//
	// CloudDisk
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// 	Notice: This parameter takes effect and is required only when PayType is set to Prepaid.
	//
	// - **Year**: subscription on a yearly basis.
	//
	// - **Month**: subscription on a monthly basis.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the region ID.
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
	// The source cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query the cluster ID.
	//
	// >This parameter is required when you restore data for an ApsaraDB for ClickHouse cluster.
	//
	// example:
	//
	// cc-bp1lxbo89u950****
	SourceDBClusterId *string `json:"SourceDBClusterId,omitempty" xml:"SourceDBClusterId,omitempty"`
	// The tag information.
	Tags []*CreateDBInstanceRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The subscription duration of the subscription cluster.
	//
	// 	Notice: This parameter takes effect and is required only when PayType is set to Prepaid.
	//
	// - If Period is set to Year, valid values: 1 to 3 (integer).
	//
	// - If Period is set to Month, valid values: 1 to 9 (integer).
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// VPC ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The secondary vSwitch.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchBak *string `json:"VSwitchBak,omitempty" xml:"VSwitchBak,omitempty"`
	// The secondary vSwitch 2.
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
	// The secondary zone 2.
	//
	// example:
	//
	// cn-hangzhou-j
	ZondIdBak2 *string `json:"ZondIdBak2,omitempty" xml:"ZondIdBak2,omitempty"`
	// The zone ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the zone ID.
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

func (s *CreateDBInstanceRequest) GetTags() []*CreateDBInstanceRequestTags {
	return s.Tags
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

func (s *CreateDBInstanceRequest) SetTags(v []*CreateDBInstanceRequestTags) *CreateDBInstanceRequest {
	s.Tags = v
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

type CreateDBInstanceRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// user123
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Example string
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateDBInstanceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateDBInstanceRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateDBInstanceRequestTags) SetKey(v string) *CreateDBInstanceRequestTags {
	s.Key = &v
	return s
}

func (s *CreateDBInstanceRequestTags) SetValue(v string) *CreateDBInstanceRequestTags {
	s.Value = &v
	return s
}

func (s *CreateDBInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}
