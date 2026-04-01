// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreDdrTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *RestoreDdrTableRequest
	GetBackupId() *string
	SetClientToken(v string) *RestoreDdrTableRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *RestoreDdrTableRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *RestoreDdrTableRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RestoreDdrTableRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *RestoreDdrTableRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *RestoreDdrTableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestoreDdrTableRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *RestoreDdrTableRequest
	GetRestoreTime() *string
	SetRestoreType(v string) *RestoreDdrTableRequest
	GetRestoreType() *string
	SetSourceDBInstanceName(v string) *RestoreDdrTableRequest
	GetSourceDBInstanceName() *string
	SetSourceRegion(v string) *RestoreDdrTableRequest
	GetSourceRegion() *string
	SetTableMeta(v string) *RestoreDdrTableRequest
	GetTableMeta() *string
}

type RestoreDdrTableRequest struct {
	// The cross-region backup set ID. You can call the DescribeCrossRegionBackups operation to query the IDs of the backup sets that are available to an instance.
	//
	// >  If you set the **RestoreType*	- parameter to **0**, you must also specify the BackupId parameter.
	//
	// example:
	//
	// 279563
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The source instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bpxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the destination instance. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to which you want to restore data. The point in time that you specify must be earlier than the current time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// > If **RestoreType*	- is set to **BackupTime**, you must specify this parameter.
	//
	// example:
	//
	// 2020-04-25T16:00:00Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The method that is used to restore data. Valid values:
	//
	// 	- **0**: restores data from a backup set. If you set this parameter to 0, you must also specify the **BackupSetId*	- parameter.
	//
	// 	- **1**: restores data to a point in time. If you set this parameter to 1, you must also specify the **RestoreTime**, **SourceRegion**, and **SourceDBInstanceName*	- parameters.
	//
	// Default value: **0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The ID of the source instance whose data you want to restore to a point in time.
	//
	// >  If you set the **RestoreType*	- parameter to **1**, you must also specify the SourceDBInstanceName parameter.
	//
	// example:
	//
	// rm-bpxxxxx
	SourceDBInstanceName *string `json:"SourceDBInstanceName,omitempty" xml:"SourceDBInstanceName,omitempty"`
	// The region ID of the source instance if you want to restore data to a point in time.
	//
	// > : If you set **RestoreType*	- to **1**, you must also specify this parameter.
	//
	// example:
	//
	// cn-beijing
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// The names of the databases and tables that you want to restore. The value is in the following format: `[{"type":"db","name":"<The name of Database 1 on the source instance>","newname":"<The name of Database 1 on the destination instance>","tables":[{"type":"table","name":"<The name of Table 1 in Database 1 on the source instance>","newname":"<The name of Table 1 in Database 1 on the destination instance>"},{"type":"table","name":"<The name of Table 2 in Database 1 on the source instance>","newname":"<The name of Table 2 in Database 1 on the destination instance>"}]},{"type":"db","name":"<The name of Database 2 on the source instance>","newname":"<The name of Database 2 on the destination instance>","tables":[{"type":"table","name":"<The name of Table 3 in Database 2 on the source instance>","newname":"<The name of Table 3 in Database 2 on the destination instance>"},{"type":"table","name":"<The name of Table 4 in Database 2 on the source instance>","newname":"<The name of Table 4 in Database 2 on the destination instance>"}]}]`
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"type":"db","name":"testdb1","newname":"testdb1","tables":[{"type":"table","name":"test1","newname":"test1_backup"},{"type":"table","name":"test2","newname":"test2_backup"}]}]
	TableMeta *string `json:"TableMeta,omitempty" xml:"TableMeta,omitempty"`
}

func (s RestoreDdrTableRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreDdrTableRequest) GoString() string {
	return s.String()
}

func (s *RestoreDdrTableRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *RestoreDdrTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestoreDdrTableRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RestoreDdrTableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestoreDdrTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestoreDdrTableRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RestoreDdrTableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestoreDdrTableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestoreDdrTableRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *RestoreDdrTableRequest) GetRestoreType() *string {
	return s.RestoreType
}

func (s *RestoreDdrTableRequest) GetSourceDBInstanceName() *string {
	return s.SourceDBInstanceName
}

func (s *RestoreDdrTableRequest) GetSourceRegion() *string {
	return s.SourceRegion
}

func (s *RestoreDdrTableRequest) GetTableMeta() *string {
	return s.TableMeta
}

func (s *RestoreDdrTableRequest) SetBackupId(v string) *RestoreDdrTableRequest {
	s.BackupId = &v
	return s
}

func (s *RestoreDdrTableRequest) SetClientToken(v string) *RestoreDdrTableRequest {
	s.ClientToken = &v
	return s
}

func (s *RestoreDdrTableRequest) SetDBInstanceId(v string) *RestoreDdrTableRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RestoreDdrTableRequest) SetOwnerId(v int64) *RestoreDdrTableRequest {
	s.OwnerId = &v
	return s
}

func (s *RestoreDdrTableRequest) SetRegionId(v string) *RestoreDdrTableRequest {
	s.RegionId = &v
	return s
}

func (s *RestoreDdrTableRequest) SetResourceGroupId(v string) *RestoreDdrTableRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RestoreDdrTableRequest) SetResourceOwnerAccount(v string) *RestoreDdrTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestoreDdrTableRequest) SetResourceOwnerId(v int64) *RestoreDdrTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestoreDdrTableRequest) SetRestoreTime(v string) *RestoreDdrTableRequest {
	s.RestoreTime = &v
	return s
}

func (s *RestoreDdrTableRequest) SetRestoreType(v string) *RestoreDdrTableRequest {
	s.RestoreType = &v
	return s
}

func (s *RestoreDdrTableRequest) SetSourceDBInstanceName(v string) *RestoreDdrTableRequest {
	s.SourceDBInstanceName = &v
	return s
}

func (s *RestoreDdrTableRequest) SetSourceRegion(v string) *RestoreDdrTableRequest {
	s.SourceRegion = &v
	return s
}

func (s *RestoreDdrTableRequest) SetTableMeta(v string) *RestoreDdrTableRequest {
	s.TableMeta = &v
	return s
}

func (s *RestoreDdrTableRequest) Validate() error {
	return dara.Validate(s)
}
