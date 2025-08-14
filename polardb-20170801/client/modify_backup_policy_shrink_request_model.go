// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedDataPoliciesShrink(v string) *ModifyBackupPolicyShrinkRequest
	GetAdvancedDataPoliciesShrink() *string
	SetBackupFrequency(v string) *ModifyBackupPolicyShrinkRequest
	GetBackupFrequency() *string
	SetBackupPolicyLevel(v string) *ModifyBackupPolicyShrinkRequest
	GetBackupPolicyLevel() *string
	SetBackupRetentionPolicyOnClusterDeletion(v string) *ModifyBackupPolicyShrinkRequest
	GetBackupRetentionPolicyOnClusterDeletion() *string
	SetDBClusterId(v string) *ModifyBackupPolicyShrinkRequest
	GetDBClusterId() *string
	SetDataLevel1BackupFrequency(v string) *ModifyBackupPolicyShrinkRequest
	GetDataLevel1BackupFrequency() *string
	SetDataLevel1BackupPeriod(v string) *ModifyBackupPolicyShrinkRequest
	GetDataLevel1BackupPeriod() *string
	SetDataLevel1BackupRetentionPeriod(v string) *ModifyBackupPolicyShrinkRequest
	GetDataLevel1BackupRetentionPeriod() *string
	SetDataLevel1BackupTime(v string) *ModifyBackupPolicyShrinkRequest
	GetDataLevel1BackupTime() *string
	SetDataLevel2BackupAnotherRegionRegion(v string) *ModifyBackupPolicyShrinkRequest
	GetDataLevel2BackupAnotherRegionRegion() *string
	SetDataLevel2BackupAnotherRegionRetentionPeriod(v string) *ModifyBackupPolicyShrinkRequest
	GetDataLevel2BackupAnotherRegionRetentionPeriod() *string
	SetDataLevel2BackupPeriod(v string) *ModifyBackupPolicyShrinkRequest
	GetDataLevel2BackupPeriod() *string
	SetDataLevel2BackupRetentionPeriod(v string) *ModifyBackupPolicyShrinkRequest
	GetDataLevel2BackupRetentionPeriod() *string
	SetOwnerAccount(v string) *ModifyBackupPolicyShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyBackupPolicyShrinkRequest
	GetOwnerId() *int64
	SetPreferredBackupPeriod(v string) *ModifyBackupPolicyShrinkRequest
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *ModifyBackupPolicyShrinkRequest
	GetPreferredBackupTime() *string
	SetResourceOwnerAccount(v string) *ModifyBackupPolicyShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyBackupPolicyShrinkRequest
	GetResourceOwnerId() *int64
}

type ModifyBackupPolicyShrinkRequest struct {
	AdvancedDataPoliciesShrink *string `json:"AdvancedDataPolicies,omitempty" xml:"AdvancedDataPolicies,omitempty"`
	// The backup frequency. Default value: Normal. Valid values:
	//
	// 	- **Normal**: standard backup. The system backs up data once a day.
	//
	// 	- **2/24H**: enhanced backup. The system backs up data every 2 hours.
	//
	// 	- **3/24H**: enhanced backup. The system backs up data every 3 hours.
	//
	// 	- **4/24H**: enhanced backup. The system backs up data every 4 hours.
	//
	// >- If you enable enhanced backup, all backups are retained for 24 hours. For backup files that are created earlier than the previous 24 hours, the system permanently retains only the first backup that is created after 00:00 every day and deletes the rest.
	//
	// >- If you enable enhanced backup, **PreferredBackupPeriod*	- is automatically set to all days in a week (from Monday to Sunday).
	//
	// >- This parameter is invalid if the region where your PolarDB for MySQL cluster is deployed supports the cross-region backup feature. For information about the regions that support the cross-region backup feature, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// Normal
	BackupFrequency   *string `json:"BackupFrequency,omitempty" xml:"BackupFrequency,omitempty"`
	BackupPolicyLevel *string `json:"BackupPolicyLevel,omitempty" xml:"BackupPolicyLevel,omitempty"`
	// Specifies whether to retain backups when a cluster is deleted. Valid values:
	//
	// 	- **ALL**: permanently retains all backups.
	//
	// 	- **LATEST**: permanently retains the most recent backup.
	//
	// 	- **NONE**: does not retain backups.
	//
	// >  The default value of the parameter is NONE.
	//
	// example:
	//
	// NONE
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// The ID of the cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query information about all clusters that are deployed in a specified region, such as the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp13wz9586voc****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The frequency of level-1 backups. Default value: Normal. Valid values:
	//
	// 	- **Normal**: standard backup. The system backs up data once a day.
	//
	// 	- **2/24H**: enhanced backup. The system backs up data every 2 hours.
	//
	// 	- **3/24H**: enhanced backup. The system backs up data every 3 hours.
	//
	// 	- **4/24H**: enhanced backup. The system backs up data every 4 hours.
	//
	// >- This parameter is invalid for PolarDB for Oracle clusters or PolarDB for PostgreSQL clusters.
	//
	// >- This parameter is invalid if the region where your PolarDB for MySQL cluster is deployed does not support the cross-region backup feature. For information about the regions that support the cross-region backup feature, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// Normal
	DataLevel1BackupFrequency *string `json:"DataLevel1BackupFrequency,omitempty" xml:"DataLevel1BackupFrequency,omitempty"`
	// The backup cycle of level-1 backups. Valid values:
	//
	// 	- **Monday**
	//
	// 	- **Tuesday**
	//
	// 	- **Wednesday**
	//
	// 	- **Thursday**
	//
	// 	- **Friday**
	//
	// 	- **Saturday**
	//
	// 	- **Sunday**
	//
	// >- You need to specify at least two values. Separate multiple values with commas (,).
	//
	// >- This parameter is invalid for PolarDB for Oracle clusters or PolarDB for PostgreSQL clusters.
	//
	// >- This parameter is invalid if the region where your PolarDB for MySQL cluster is deployed does not support the cross-region backup feature. For information about the regions that support the cross-region backup feature, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// Monday,Tuesday
	DataLevel1BackupPeriod *string `json:"DataLevel1BackupPeriod,omitempty" xml:"DataLevel1BackupPeriod,omitempty"`
	// The retention period of level-1 backups. Valid values: 3 to 14. Unit: days.
	//
	// example:
	//
	// 3
	DataLevel1BackupRetentionPeriod *string `json:"DataLevel1BackupRetentionPeriod,omitempty" xml:"DataLevel1BackupRetentionPeriod,omitempty"`
	// The time period during which automatic backup for level-1 backup is performed. The time period is in the `hh:mmZ-hh:mmZ` format and is displayed in UTC. The start time and end time are on the hour and have an interval of 1 hour. Example: `14:00Z-15:00Z`.
	//
	// >- This parameter is invalid for PolarDB for Oracle clusters or PolarDB for PostgreSQL clusters.
	//
	// >- This parameter is invalid if the region where your PolarDB for MySQL cluster is deployed does not support the cross-region backup feature. For information about the regions that support the cross-region backup feature, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// 15:00Z-16:00Z
	DataLevel1BackupTime *string `json:"DataLevel1BackupTime,omitempty" xml:"DataLevel1BackupTime,omitempty"`
	// The region where the cross-region level-2 backup is stored. For information about regions that support the cross-region backup feature, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// cn-hangzhou
	DataLevel2BackupAnotherRegionRegion *string `json:"DataLevel2BackupAnotherRegionRegion,omitempty" xml:"DataLevel2BackupAnotherRegionRegion,omitempty"`
	// The retention period of cross-region level-2 backups. Valid values:
	//
	// 	- **0**: The cross-region level-2 backup feature is disabled.
	//
	// 	- **30 to 7300**: Cross-region level-2 backups are retained for 30 to 7,300 days.
	//
	// 	- **1**: Cross-region level-2 backups are permanently retained.
	//
	// >  The default value of the parameter is **0**.
	//
	// example:
	//
	// 30
	DataLevel2BackupAnotherRegionRetentionPeriod *string `json:"DataLevel2BackupAnotherRegionRetentionPeriod,omitempty" xml:"DataLevel2BackupAnotherRegionRetentionPeriod,omitempty"`
	// The backup cycle of level-2 backups. Valid values:
	//
	// 	- **Monday**
	//
	// 	- **Tuesday**
	//
	// 	- **Wednesday**
	//
	// 	- **Thursday**
	//
	// 	- **Friday**
	//
	// 	- **Saturday**
	//
	// 	- **Sunday**
	//
	// >- You need to specify at least two values. Separate multiple values with commas (,).
	//
	// >- This parameter is invalid for PolarDB for Oracle clusters or PolarDB for PostgreSQL clusters.
	//
	// >- This parameter is invalid if the region where your PolarDB for MySQL cluster is deployed does not support the cross-region backup feature. For information about the regions that support the cross-region backup feature, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// Monday,Tuesday
	DataLevel2BackupPeriod *string `json:"DataLevel2BackupPeriod,omitempty" xml:"DataLevel2BackupPeriod,omitempty"`
	// The retention period of level-2 backups. Valid values:
	//
	// 	- **0**: The level-2 backup feature is disabled.
	//
	// 	- **30 to 7300**: Level-2 backups are retained for 30 to 7,300 days.
	//
	// 	- **1**: Level-2 backups are permanently retained.
	//
	// >  The default value of this parameter is **0**.
	//
	// example:
	//
	// 0
	DataLevel2BackupRetentionPeriod *string `json:"DataLevel2BackupRetentionPeriod,omitempty" xml:"DataLevel2BackupRetentionPeriod,omitempty"`
	OwnerAccount                    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The backup cycle. Valid values:
	//
	// 	- **Monday**
	//
	// 	- **Tuesday**
	//
	// 	- **Wednesday**
	//
	// 	- **Thursday**
	//
	// 	- **Friday**
	//
	// 	- **Saturday**
	//
	// 	- **Sunday**
	//
	// >- You need to specify at least two values. Separate multiple values with commas (,).
	//
	// >- This parameter is invalid if the region where your PolarDB for MySQL cluster is deployed supports the cross-region backup feature. For information about the regions that support the cross-region backup feature, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// Monday,Tuesday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The time period during which automatic backup for level-1 backup is performed. The format is `hh:mmZ-hh:mmZ` format. The time is displayed in UTC. The start time and end time are on the hour and with an interval of one hour. Example: `14:00Z-15:00Z`.
	//
	// example:
	//
	// 15:00Z-16:00Z
	PreferredBackupTime  *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyBackupPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyShrinkRequest) GetAdvancedDataPoliciesShrink() *string {
	return s.AdvancedDataPoliciesShrink
}

func (s *ModifyBackupPolicyShrinkRequest) GetBackupFrequency() *string {
	return s.BackupFrequency
}

func (s *ModifyBackupPolicyShrinkRequest) GetBackupPolicyLevel() *string {
	return s.BackupPolicyLevel
}

func (s *ModifyBackupPolicyShrinkRequest) GetBackupRetentionPolicyOnClusterDeletion() *string {
	return s.BackupRetentionPolicyOnClusterDeletion
}

func (s *ModifyBackupPolicyShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyBackupPolicyShrinkRequest) GetDataLevel1BackupFrequency() *string {
	return s.DataLevel1BackupFrequency
}

func (s *ModifyBackupPolicyShrinkRequest) GetDataLevel1BackupPeriod() *string {
	return s.DataLevel1BackupPeriod
}

func (s *ModifyBackupPolicyShrinkRequest) GetDataLevel1BackupRetentionPeriod() *string {
	return s.DataLevel1BackupRetentionPeriod
}

func (s *ModifyBackupPolicyShrinkRequest) GetDataLevel1BackupTime() *string {
	return s.DataLevel1BackupTime
}

func (s *ModifyBackupPolicyShrinkRequest) GetDataLevel2BackupAnotherRegionRegion() *string {
	return s.DataLevel2BackupAnotherRegionRegion
}

func (s *ModifyBackupPolicyShrinkRequest) GetDataLevel2BackupAnotherRegionRetentionPeriod() *string {
	return s.DataLevel2BackupAnotherRegionRetentionPeriod
}

func (s *ModifyBackupPolicyShrinkRequest) GetDataLevel2BackupPeriod() *string {
	return s.DataLevel2BackupPeriod
}

func (s *ModifyBackupPolicyShrinkRequest) GetDataLevel2BackupRetentionPeriod() *string {
	return s.DataLevel2BackupRetentionPeriod
}

func (s *ModifyBackupPolicyShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyBackupPolicyShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyBackupPolicyShrinkRequest) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *ModifyBackupPolicyShrinkRequest) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *ModifyBackupPolicyShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyBackupPolicyShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyBackupPolicyShrinkRequest) SetAdvancedDataPoliciesShrink(v string) *ModifyBackupPolicyShrinkRequest {
	s.AdvancedDataPoliciesShrink = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetBackupFrequency(v string) *ModifyBackupPolicyShrinkRequest {
	s.BackupFrequency = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetBackupPolicyLevel(v string) *ModifyBackupPolicyShrinkRequest {
	s.BackupPolicyLevel = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetBackupRetentionPolicyOnClusterDeletion(v string) *ModifyBackupPolicyShrinkRequest {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetDBClusterId(v string) *ModifyBackupPolicyShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetDataLevel1BackupFrequency(v string) *ModifyBackupPolicyShrinkRequest {
	s.DataLevel1BackupFrequency = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetDataLevel1BackupPeriod(v string) *ModifyBackupPolicyShrinkRequest {
	s.DataLevel1BackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetDataLevel1BackupRetentionPeriod(v string) *ModifyBackupPolicyShrinkRequest {
	s.DataLevel1BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetDataLevel1BackupTime(v string) *ModifyBackupPolicyShrinkRequest {
	s.DataLevel1BackupTime = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetDataLevel2BackupAnotherRegionRegion(v string) *ModifyBackupPolicyShrinkRequest {
	s.DataLevel2BackupAnotherRegionRegion = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetDataLevel2BackupAnotherRegionRetentionPeriod(v string) *ModifyBackupPolicyShrinkRequest {
	s.DataLevel2BackupAnotherRegionRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetDataLevel2BackupPeriod(v string) *ModifyBackupPolicyShrinkRequest {
	s.DataLevel2BackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetDataLevel2BackupRetentionPeriod(v string) *ModifyBackupPolicyShrinkRequest {
	s.DataLevel2BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetOwnerAccount(v string) *ModifyBackupPolicyShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetOwnerId(v int64) *ModifyBackupPolicyShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyShrinkRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyShrinkRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetResourceOwnerAccount(v string) *ModifyBackupPolicyShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetResourceOwnerId(v int64) *ModifyBackupPolicyShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
