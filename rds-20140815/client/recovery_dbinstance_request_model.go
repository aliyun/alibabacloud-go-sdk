// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoveryDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *RecoveryDBInstanceRequest
	GetBackupId() *string
	SetDBInstanceClass(v string) *RecoveryDBInstanceRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *RecoveryDBInstanceRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int32) *RecoveryDBInstanceRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *RecoveryDBInstanceRequest
	GetDBInstanceStorageType() *string
	SetDbNames(v string) *RecoveryDBInstanceRequest
	GetDbNames() *string
	SetInstanceNetworkType(v string) *RecoveryDBInstanceRequest
	GetInstanceNetworkType() *string
	SetPayType(v string) *RecoveryDBInstanceRequest
	GetPayType() *string
	SetPeriod(v string) *RecoveryDBInstanceRequest
	GetPeriod() *string
	SetPrivateIpAddress(v string) *RecoveryDBInstanceRequest
	GetPrivateIpAddress() *string
	SetResourceOwnerId(v int64) *RecoveryDBInstanceRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *RecoveryDBInstanceRequest
	GetRestoreTime() *string
	SetTargetDBInstanceId(v string) *RecoveryDBInstanceRequest
	GetTargetDBInstanceId() *string
	SetUsedTime(v string) *RecoveryDBInstanceRequest
	GetUsedTime() *string
	SetVPCId(v string) *RecoveryDBInstanceRequest
	GetVPCId() *string
	SetVSwitchId(v string) *RecoveryDBInstanceRequest
	GetVSwitchId() *string
}

type RecoveryDBInstanceRequest struct {
	// The backup set ID. You can call the DescribeBackups operation to query backup sets.
	//
	// If you specify this parameter, the **DBInstanceId*	- parameter is optional.
	//
	// > You must specify at least one of **BackupId*	- and **RestoreTime**.
	//
	// example:
	//
	// 29304****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The instance type of the new instance. For more information, see [Instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// example:
	//
	// mssql.x4.medium.s1
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance ID of the original instance.
	//
	// > 	- If you want to recover data by backup set (by specifying the BackupId parameter), this parameter is optional.
	//
	// > 	- If you want to recover data to a point in time (by specifying the RestoreTime parameter), this parameter is required.
	//
	// example:
	//
	// rm-bp18****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The instance storage capacity of the new instance. Unit: GB. For details, see [Instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// > The disk space of the new instance cannot be smaller than that of the original instance.
	//
	// example:
	//
	// 40
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The instance storage type of the new instance. Valid values:
	//
	// 	- **local_ssd/ephemeral_ssd**: local SSD.
	//
	// 	- **cloud_ssd**: standard SSD cloud disk.
	//
	// 	- **cloud_essd**: Enterprise SSD (ESSD) cloud disk.
	//
	// example:
	//
	// cloud_essd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The database name. To restore data to a new instance, use the following format: `Original database name 1,New database name 2`.
	//
	// > To restore data to an existing instance, see [CopyDatabaseBetweenInstances](https://help.aliyun.com/document_detail/2628854.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test1,test2
	DbNames *string `json:"DbNames,omitempty" xml:"DbNames,omitempty"`
	// The network type of the new instance. Valid values:
	//
	// 	- **Classic**: classic network.
	//
	// 	- **VPC**: virtual private cloud (VPC).
	//
	// Default value: the network type of the original instance.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The billing method of the new instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration of the new instance. Valid values:
	//
	// 	- **Year**: year.
	//
	// 	- **Month**: month.
	//
	// > This parameter is required if **PayType*	- is set to **Prepaid**.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The internal IP address of the new instance. The IP address must be within the IP address range of the specified vSwitch. By default, the system automatically assigns an IP address based on the values of **VPCId*	- and **VSwitchId**.
	//
	// example:
	//
	// 172.XX.XX.69
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	ResourceOwnerId  *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Any point in time within the backup retention period. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// If you specify this parameter, the **DBInstanceId*	- parameter is required.
	//
	// > You must specify at least one of **BackupId*	- and **RestoreTime**.
	//
	// example:
	//
	// 2011-06-11T16:00:00Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The instance ID of the target instance.
	//
	// example:
	//
	// rm-bp17****
	TargetDBInstanceId *string `json:"TargetDBInstanceId,omitempty" xml:"TargetDBInstanceId,omitempty"`
	// The subscription duration of the new instance. Valid values:
	//
	// 	- If **Period*	- is set to **Year**, the value of **UsedTime*	- ranges from **1 to 3**.
	//
	// 	- If **Period*	- is set to **Month**, the value of **UsedTime*	- ranges from **1 to 9**.
	//
	// > This parameter is required if **PayType*	- is set to **Prepaid**.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The VPC ID of the new instance.
	//
	// example:
	//
	// vpc-****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the new instance. Separate multiple values with commas (,).
	//
	// example:
	//
	// vsw-****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s RecoveryDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecoveryDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RecoveryDBInstanceRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *RecoveryDBInstanceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *RecoveryDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RecoveryDBInstanceRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *RecoveryDBInstanceRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *RecoveryDBInstanceRequest) GetDbNames() *string {
	return s.DbNames
}

func (s *RecoveryDBInstanceRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *RecoveryDBInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *RecoveryDBInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *RecoveryDBInstanceRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *RecoveryDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RecoveryDBInstanceRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *RecoveryDBInstanceRequest) GetTargetDBInstanceId() *string {
	return s.TargetDBInstanceId
}

func (s *RecoveryDBInstanceRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *RecoveryDBInstanceRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *RecoveryDBInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *RecoveryDBInstanceRequest) SetBackupId(v string) *RecoveryDBInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetDBInstanceClass(v string) *RecoveryDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetDBInstanceId(v string) *RecoveryDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetDBInstanceStorage(v int32) *RecoveryDBInstanceRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetDBInstanceStorageType(v string) *RecoveryDBInstanceRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetDbNames(v string) *RecoveryDBInstanceRequest {
	s.DbNames = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetInstanceNetworkType(v string) *RecoveryDBInstanceRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetPayType(v string) *RecoveryDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetPeriod(v string) *RecoveryDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetPrivateIpAddress(v string) *RecoveryDBInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetResourceOwnerId(v int64) *RecoveryDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetRestoreTime(v string) *RecoveryDBInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetTargetDBInstanceId(v string) *RecoveryDBInstanceRequest {
	s.TargetDBInstanceId = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetUsedTime(v string) *RecoveryDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetVPCId(v string) *RecoveryDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *RecoveryDBInstanceRequest) SetVSwitchId(v string) *RecoveryDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *RecoveryDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
