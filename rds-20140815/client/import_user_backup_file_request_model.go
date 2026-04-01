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
	// A JSON array that consists of the information about the full backup file stored as an object in an OSS bucket. Example: `{"Bucket":"test", "Object":"test/test_db_employees.xb","Location":"ap-southeast-1"}`
	//
	// The JSON array contains the following fields:
	//
	// 	- **Bucket**: The name of the OSS bucket in which the full backup file is stored as an object. You can call the [GetBucket](https://help.aliyun.com/document_detail/31965.html) operation to query the name of the bucket.
	//
	// 	- **Object**: The path of the full backup file that is stored as an object in the OSS bucket. You can call the [GetObject](https://help.aliyun.com/document_detail/31980.html) operation to query the path of the object.
	//
	// 	- **Location**: The ID of the region in which the OSS bucket is located. You can call the [GetBucketLocation](https://help.aliyun.com/document_detail/31967.html) operation to query the region of the bucket.
	//
	// example:
	//
	// {"Bucket":"test", "Object":"test/test_db_employees.xb","Location":"ap-southeast-1"}
	BackupFile *string `json:"BackupFile,omitempty" xml:"BackupFile,omitempty"`
	// The region ID of the OSS bucket where the full backup file of the self-managed MySQL database is located. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	BucketRegion     *string `json:"BucketRegion,omitempty" xml:"BucketRegion,omitempty"`
	BuildReplication *bool   `json:"BuildReplication,omitempty" xml:"BuildReplication,omitempty"`
	// The description of the full backup file.
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
	// The version of the database engine that is run on the self-managed MySQL database and ApsaraDB RDS for MySQL instance. Set the value to **5.7**.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	MasterInfo    *string `json:"MasterInfo,omitempty" xml:"MasterInfo,omitempty"`
	// example:
	//
	// oss
	Mode    *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the DescribeRegions operation to query the most recent region list.
	//
	// > 	- The value of this parameter is the ID of the region in which you want to create the instance.
	//
	// > 	- The value of this parameter must be consistent with the value of **BucketRegion**.
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
	// The amount of storage that is required to restore the data of the full backup file. Unit: GB.
	//
	// > 	- The default value of this parameter is 5 times the size of the full backup file.
	//
	// > 	- The minimum value of this parameter is 20.
	//
	// example:
	//
	// 20
	RestoreSize *int32 `json:"RestoreSize,omitempty" xml:"RestoreSize,omitempty"`
	// The retention period of the full backup file. Unit: days. Valid values: any **non-zero*	- positive integer.
	//
	// example:
	//
	// 30
	Retention  *int32  `json:"Retention,omitempty" xml:"Retention,omitempty"`
	SourceInfo *string `json:"SourceInfo,omitempty" xml:"SourceInfo,omitempty"`
	// The zone ID. You can call the DescribeRegions operation to query the zone ID.
	//
	// > 	- If you specify this parameter, the system creates a snapshot in single-digit seconds, which greatly reduces the time that is required to import the full backup file.
	//
	// > 	- When you call the CreateDBInstance operation to create an instance by using the full backup file, the instance is created in the zone that you specify for this parameter.
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
