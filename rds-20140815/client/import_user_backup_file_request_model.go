// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportUserBackupFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupFile(v string) *ImportUserBackupFileRequest
	GetBackupFile() *string
	SetBucketRegion(v string) *ImportUserBackupFileRequest
	GetBucketRegion() *string
	SetBuildReplication(v bool) *ImportUserBackupFileRequest
	GetBuildReplication() *bool
	SetComment(v string) *ImportUserBackupFileRequest
	GetComment() *string
	SetDBInstanceId(v string) *ImportUserBackupFileRequest
	GetDBInstanceId() *string
	SetEngineVersion(v string) *ImportUserBackupFileRequest
	GetEngineVersion() *string
	SetMasterInfo(v string) *ImportUserBackupFileRequest
	GetMasterInfo() *string
	SetMode(v string) *ImportUserBackupFileRequest
	GetMode() *string
	SetOwnerId(v int64) *ImportUserBackupFileRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ImportUserBackupFileRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ImportUserBackupFileRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ImportUserBackupFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ImportUserBackupFileRequest
	GetResourceOwnerId() *int64
	SetRestoreSize(v int32) *ImportUserBackupFileRequest
	GetRestoreSize() *int32
	SetRetention(v int32) *ImportUserBackupFileRequest
	GetRetention() *int32
	SetSourceInfo(v string) *ImportUserBackupFileRequest
	GetSourceInfo() *string
	SetZoneId(v string) *ImportUserBackupFileRequest
	GetZoneId() *string
}

type ImportUserBackupFileRequest struct {
	// The JSON array that describes the information about the backup file in the OSS bucket. Example:
	//
	// `{"Bucket":"test", "Object":"test/test_db_employees.xb","Location":"ap-southeast-1"}`
	//
	// Parameters in the array:
	//
	// 	- **Bucket**: the name of the OSS bucket where the backup file is stored. You can call the [GetBucket](https://help.aliyun.com/document_detail/31965.html) operation to query the name.
	//
	// 	- **Object**: the detailed path of the directory where the backup file is stored. You can call the [GetObject](https://help.aliyun.com/document_detail/31980.html) operation to query the path.
	//
	// 	- **Location**: the region ID of the OSS bucket. You can call the [GetBucketLocation](https://help.aliyun.com/document_detail/31967.html) operation to query the region ID.
	//
	// This parameter is required when **Mode*	- is set to **oss**. For **Location**, specify the region ID of the OSS bucket. You can call the GetBucketLocation operation to query the region ID.
	//
	// example:
	//
	// {"Bucket":"test", "Object":"test/test_db_employees.xb","Location":"ap-southeast-1"}
	BackupFile *string `json:"BackupFile,omitempty" xml:"BackupFile,omitempty"`
	// The region ID of the OSS bucket where the backup file of the self-managed MySQL 5.7 database is stored. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	BucketRegion *string `json:"BucketRegion,omitempty" xml:"BucketRegion,omitempty"`
	// Specifies whether to automatically set up replication. Valid values:
	//
	// - true: yes. The MasterInfo parameter is required.
	//
	// - false: no.
	//
	// > This parameter takes effect only for native replication instances. You must specify the DBInstanceId parameter when you call this operation.
	//
	// example:
	//
	// true
	BuildReplication *bool `json:"BuildReplication,omitempty" xml:"BuildReplication,omitempty"`
	// The description of the user backup to be imported.
	//
	// example:
	//
	// BackupTest
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The version of the MySQL database. Valid values: **5.7*	- and **8.0**.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The case-sensitive JSON array that provides the master information for setting up MySQL replication. Example:
	//
	// ```
	//
	// {"masterIp":"172.20.xx.xx","masterPort":"3306","masterUser":"replica","masterPassword":"W33uopkehBQ="}
	//
	// ```
	//
	// Parameters in the array:
	//
	// - masterIp: the IP address of the primary database.
	//
	// - masterPort: the port of the primary database.
	//
	// - masterUser: the replication account of the primary database.
	//
	// - masterPassword: the password of the replication account of the primary database. The password must be Base64-encoded.
	//
	// > This parameter takes effect only for native replication instances. You must specify the DBInstanceId parameter when you call this operation.
	//
	// example:
	//
	// {"masterIp":"172.20.xx.xx","masterPort":"3306","masterUser":"replica","masterPassword":"W33uopkehBQ="}
	MasterInfo *string `json:"MasterInfo,omitempty" xml:"MasterInfo,omitempty"`
	// The import mode. Valid values:
	//
	// - oss: downloads the backup from OSS and imports it.
	//
	// - stream: imports the backup over the network.
	//
	// example:
	//
	// oss
	Mode    *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the ApsaraDB RDS instance. You can call the DescribeRegions operation to query the most recent region list.
	//
	// > 	- The value of this parameter is the region ID where you want to create the ApsaraDB RDS instance.
	//
	// > 	- The value must be the same as the value of the **BucketRegion*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to query the resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The storage space required to restore the user backup. Unit: GB.
	//
	// > 	- The default value is five times the size of the backup file.
	//
	// > 	- The minimum value is 20.
	//
	// example:
	//
	// 20
	RestoreSize *int32 `json:"RestoreSize,omitempty" xml:"RestoreSize,omitempty"`
	// The retention period of the user backup file. Unit: days. The value must be an integer greater than **0**.
	//
	// example:
	//
	// 30
	Retention *int32 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The case-sensitive JSON array that provides the source information of the full backup. Example:
	//
	// ```
	//
	// {"sourceIp":"172.20.xx.xx","sourcePort":"9999"}
	//
	// ```
	//
	// Parameters in the array:
	//
	// - sourceIp: the source IP address.
	//
	// - sourcePort: the port listened by Netcat on the source.
	//
	// > This parameter takes effect only for native replication instances. You must specify the DBInstanceId parameter when you call this operation.
	//
	// example:
	//
	// {"sourceIp":"172.20.xx.xx","sourcePort":"9999"}
	SourceInfo *string `json:"SourceInfo,omitempty" xml:"SourceInfo,omitempty"`
	// The zone ID. You can call the DescribeRegions operation to query the most recent zone list.
	//
	// > 	- After you specify a zone, the system creates a snapshot within seconds in the zone, which significantly reduces the time required to import the backup.
	//
	// > 	- When you call the CreateDBInstance operation to create an instance by using the user backup, the instance is created in this zone.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ImportUserBackupFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportUserBackupFileRequest) GoString() string {
	return s.String()
}

func (s *ImportUserBackupFileRequest) GetBackupFile() *string {
	return s.BackupFile
}

func (s *ImportUserBackupFileRequest) GetBucketRegion() *string {
	return s.BucketRegion
}

func (s *ImportUserBackupFileRequest) GetBuildReplication() *bool {
	return s.BuildReplication
}

func (s *ImportUserBackupFileRequest) GetComment() *string {
	return s.Comment
}

func (s *ImportUserBackupFileRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ImportUserBackupFileRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *ImportUserBackupFileRequest) GetMasterInfo() *string {
	return s.MasterInfo
}

func (s *ImportUserBackupFileRequest) GetMode() *string {
	return s.Mode
}

func (s *ImportUserBackupFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ImportUserBackupFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ImportUserBackupFileRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ImportUserBackupFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ImportUserBackupFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ImportUserBackupFileRequest) GetRestoreSize() *int32 {
	return s.RestoreSize
}

func (s *ImportUserBackupFileRequest) GetRetention() *int32 {
	return s.Retention
}

func (s *ImportUserBackupFileRequest) GetSourceInfo() *string {
	return s.SourceInfo
}

func (s *ImportUserBackupFileRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ImportUserBackupFileRequest) SetBackupFile(v string) *ImportUserBackupFileRequest {
	s.BackupFile = &v
	return s
}

func (s *ImportUserBackupFileRequest) SetBucketRegion(v string) *ImportUserBackupFileRequest {
	s.BucketRegion = &v
	return s
}

func (s *ImportUserBackupFileRequest) SetBuildReplication(v bool) *ImportUserBackupFileRequest {
	s.BuildReplication = &v
	return s
}

func (s *ImportUserBackupFileRequest) SetComment(v string) *ImportUserBackupFileRequest {
	s.Comment = &v
	return s
}

func (s *ImportUserBackupFileRequest) SetDBInstanceId(v string) *ImportUserBackupFileRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ImportUserBackupFileRequest) SetEngineVersion(v string) *ImportUserBackupFileRequest {
	s.EngineVersion = &v
	return s
}

func (s *ImportUserBackupFileRequest) SetMasterInfo(v string) *ImportUserBackupFileRequest {
	s.MasterInfo = &v
	return s
}

func (s *ImportUserBackupFileRequest) SetMode(v string) *ImportUserBackupFileRequest {
	s.Mode = &v
	return s
}

func (s *ImportUserBackupFileRequest) SetOwnerId(v int64) *ImportUserBackupFileRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportUserBackupFileRequest) SetRegionId(v string) *ImportUserBackupFileRequest {
	s.RegionId = &v
	return s
}

func (s *ImportUserBackupFileRequest) SetResourceGroupId(v string) *ImportUserBackupFileRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ImportUserBackupFileRequest) SetResourceOwnerAccount(v string) *ImportUserBackupFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportUserBackupFileRequest) SetResourceOwnerId(v int64) *ImportUserBackupFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ImportUserBackupFileRequest) SetRestoreSize(v int32) *ImportUserBackupFileRequest {
	s.RestoreSize = &v
	return s
}

func (s *ImportUserBackupFileRequest) SetRetention(v int32) *ImportUserBackupFileRequest {
	s.Retention = &v
	return s
}

func (s *ImportUserBackupFileRequest) SetSourceInfo(v string) *ImportUserBackupFileRequest {
	s.SourceInfo = &v
	return s
}

func (s *ImportUserBackupFileRequest) SetZoneId(v string) *ImportUserBackupFileRequest {
	s.ZoneId = &v
	return s
}

func (s *ImportUserBackupFileRequest) Validate() error {
	return dara.Validate(s)
}
