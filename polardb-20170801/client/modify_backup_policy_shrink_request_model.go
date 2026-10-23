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
	// The advanced backup policy.
	//
	// > 	- 	- This parameter is not supported by PolarDB for PostgreSQL (Compatible with Oracle) and PolarDB for PostgreSQL.
	//
	// > 	- 	- This parameter is supported only by clusters whose BackupPolicyLevel is set to Advanced.
	AdvancedDataPoliciesShrink *string `json:"AdvancedDataPolicies,omitempty" xml:"AdvancedDataPolicies,omitempty"`
	// The backup frequency. Valid values:
	//
	// - **Normal*	- (default): standard backup, performed once a day at a scheduled time.
	//
	// - **2/24H**: high-frequency backup, performed once every 2 hours.
	//
	// - **3/24H**: high-frequency backup, performed once every 3 hours.
	//
	// - **4/24H**: high-frequency backup, performed once every 4 hours.
	//
	// > 	- 	- After high-frequency backup is enabled, all backups completed within 24 hours are retained. For backups that exceed 24 hours, the system retains only the first backup completed after 00:00 each day and deletes the rest.
	//
	// > 	- 	- After high-frequency backup is enabled, all days of the week (Monday to Sunday) are selected by default for the backup epoch **PreferredBackupPeriod*	- parameter.
	//
	// > 	- 	- This parameter is not supported if the region where your PolarDB for MySQL cluster resides supports the cross-region backup feature. For more information about the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// > 	- 	- After advanced backup is enabled, this parameter no longer takes effect. Use the AdvancedDataPolicies parameter instead.
	//
	// example:
	//
	// Normal
	BackupFrequency *string `json:"BackupFrequency,omitempty" xml:"BackupFrequency,omitempty"`
	// The level of the backup policy. Valid values:
	//
	// 	- **Normal**: standard backup
	//
	// 	- **Advanced**: advanced backup
	//
	// > 	- 	- This parameter is not supported by PolarDB for PostgreSQL (Compatible with Oracle) and PolarDB for PostgreSQL.
	//
	// > 	- 	- You can check the AdvancedPolicyOption parameter in the response of the [DescribeBackupPolicy](https://help.aliyun.com/document_detail/2319231.html) operation to determine whether the cluster supports advanced backup. If the cluster supports advanced backup, you can apply for this feature in [Advanced backup settings](~611727~~).
	//
	// > 	- 	- After advanced backup is enabled, you **cannot*	- roll back to standard backup.
	//
	// example:
	//
	// Normal
	BackupPolicyLevel *string `json:"BackupPolicyLevel,omitempty" xml:"BackupPolicyLevel,omitempty"`
	// Specifies whether to retain backups when a cluster is deleted. Valid values:
	//
	// - **ALL**: long-term retention (LTR) of all backups.
	//
	// - **LATEST**: long-term retention (LTR) of the last backup.
	//
	// - **NONE**: does not retain backup sets.
	//
	// > Default value: NONE.
	//
	// example:
	//
	// NONE
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query all clusters in a specific region, including their cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp13wz9586voc****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The backup frequency. Valid values:
	//
	// - **Normal*	- (default): standard backup, performed once a day at a scheduled time
	//
	// - **2/24H**: high-frequency backup, performed once every 2 hours
	//
	// - **3/24H**: high-frequency backup, performed once every 3 hours
	//
	// - **4/24H**: high-frequency backup, performed once every 4 hours
	//
	// > 	- 	- This parameter is not supported by PolarDB for PostgreSQL (Compatible with Oracle) and PolarDB for PostgreSQL.
	//
	// > 	- 	- This parameter is not supported if the region where your PolarDB for MySQL cluster resides does not support cross-region backup. For more information about the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// > 	- 	- After advanced backup is enabled, this parameter no longer takes effect. Use the AdvancedDataPolicies parameter instead.
	//
	// example:
	//
	// Normal
	DataLevel1BackupFrequency *string `json:"DataLevel1BackupFrequency,omitempty" xml:"DataLevel1BackupFrequency,omitempty"`
	// The backup cycle for level-1 backups. Valid values:
	//
	// 	- **Monday**: Monday
	//
	// 	- **Tuesday**: Tuesday
	//
	// 	- **Wednesday**: Wednesday
	//
	// 	- **Thursday**: Thursday
	//
	// 	- **Friday**: Friday
	//
	// 	- **Saturday**: Saturday
	//
	// 	- **Sunday**: Sunday
	//
	// > 	- 	- You must select at least two days. Separate multiple values with commas (,).
	//
	// > 	- 	- This parameter is not supported by PolarDB for PostgreSQL (Compatible with Oracle) and PolarDB for PostgreSQL.
	//
	// > 	- 	- This parameter is not supported if the region where your PolarDB for MySQL cluster resides does not support cross-region backup. For more information about the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// > 	- 	- After advanced backup is enabled, this parameter no longer takes effect. Use the AdvancedDataPolicies parameter instead.
	//
	// example:
	//
	// Monday,Tuesday
	DataLevel1BackupPeriod *string `json:"DataLevel1BackupPeriod,omitempty" xml:"DataLevel1BackupPeriod,omitempty"`
	// The retention period of level-1 backups. Valid values: 3 to 14. Unit: days.
	//
	// > 	- After advanced backup is enabled, this parameter no longer takes effect. Use the AdvancedDataPolicies parameter instead.
	//
	// example:
	//
	// 3
	DataLevel1BackupRetentionPeriod *string `json:"DataLevel1BackupRetentionPeriod,omitempty" xml:"DataLevel1BackupRetentionPeriod,omitempty"`
	// The time period during which automatic backups are performed. The time must be in the `hh:mmZ-hh:mmZ` format (UTC time) and must be on the hour with a 1-hour interval, such as `14:00Z-15:00Z`.
	//
	// > 	- This parameter is not supported by PolarDB for PostgreSQL (Compatible with Oracle) and PolarDB for PostgreSQL.
	//
	// > 	- This parameter is not supported if the region where your PolarDB for MySQL cluster resides does not support cross-region backup. For more information about the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// 15:00Z-16:00Z
	DataLevel1BackupTime *string `json:"DataLevel1BackupTime,omitempty" xml:"DataLevel1BackupTime,omitempty"`
	// The destination region for cross-region backup of level-2 backups. For more information about the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// > 	- After advanced backup is enabled, this parameter no longer takes effect. Use the AdvancedDataPolicies parameter instead.
	//
	// example:
	//
	// cn-hangzhou
	DataLevel2BackupAnotherRegionRegion *string `json:"DataLevel2BackupAnotherRegionRegion,omitempty" xml:"DataLevel2BackupAnotherRegionRegion,omitempty"`
	// The retention epoch for cross-region backup of level-2 backups. Valid values:
	//
	// - **0**: shutdown of the level-2 backup feature.
	//
	// - **30 to 7300**: the retention epoch of level-2 backups, in days.
	//
	// - **-1**: long-term retention (LTR) of level-2 backups.
	//
	//  > 	- 	- When a cluster is created, the default value is **0**, which indicates the shutdown of the cross-region backup feature for level-2 backups.
	//
	// > 	- 	- After advanced backup is enabled, this parameter no longer takes effect. Use the AdvancedDataPolicies parameter instead.
	//
	// example:
	//
	// 30
	DataLevel2BackupAnotherRegionRetentionPeriod *string `json:"DataLevel2BackupAnotherRegionRetentionPeriod,omitempty" xml:"DataLevel2BackupAnotherRegionRetentionPeriod,omitempty"`
	// The backup cycle for level-2 backups. Valid values:
	//
	// 	- **Monday**: Monday
	//
	// 	- **Tuesday**: Tuesday
	//
	// 	- **Wednesday**: Wednesday
	//
	// 	- **Thursday**: Thursday
	//
	// 	- **Friday**: Friday
	//
	// 	- **Saturday**: Saturday
	//
	// 	- **Sunday**: Sunday
	//
	// > 	- 	- You must select at least two days. Separate multiple values with commas (,).
	//
	// > 	- 	- This parameter is not supported by PolarDB for PostgreSQL (Compatible with Oracle) and PolarDB for PostgreSQL.
	//
	// > 	- 	- This parameter is not supported if the region where your PolarDB for MySQL cluster resides does not support cross-region backup. For more information about the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// > 	- 	- After advanced backup is enabled, this parameter no longer takes effect. Use the AdvancedDataPolicies parameter instead.
	//
	// example:
	//
	// Monday,Tuesday
	DataLevel2BackupPeriod *string `json:"DataLevel2BackupPeriod,omitempty" xml:"DataLevel2BackupPeriod,omitempty"`
	// The retention epoch of level-2 backups. Valid values:
	//
	// - **0**: shutdown of the level-2 backup feature.
	//
	// - **30 to 7300**: the retention epoch of level-2 backups, in days.
	//
	// - **-1**: long-term retention (LTR) of level-2 backups.
	//
	//  >	- 	- When a cluster is created, the default value is **0**, which indicates the shutdown of the level-2 backup feature.
	//
	// > 	- 	- After advanced backup is enabled, this parameter no longer takes effect. Use the AdvancedDataPolicies parameter instead.
	//
	// example:
	//
	// 0
	DataLevel2BackupRetentionPeriod *string `json:"DataLevel2BackupRetentionPeriod,omitempty" xml:"DataLevel2BackupRetentionPeriod,omitempty"`
	OwnerAccount                    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The backup cycle. Valid values:
	//
	// 	- **Monday**: Monday
	//
	// 	- **Tuesday**: Tuesday
	//
	// 	- **Wednesday**: Wednesday
	//
	// 	- **Thursday**: Thursday
	//
	// 	- **Friday**: Friday
	//
	// 	- **Saturday**: Saturday
	//
	// 	- **Sunday**: Sunday
	//
	// > 	- 	- You must select at least two days. Separate multiple values with commas (,).
	//
	// > 	- 	- This parameter is not supported if the region where your PolarDB for MySQL cluster resides supports cross-region backup. For more information about the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// > 	- 	- After advanced backup is enabled, this parameter no longer takes effect. Use the AdvancedDataPolicies parameter instead.
	//
	// example:
	//
	// Monday,Tuesday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The time period during which automatic backups are performed. The time must be in the `hh:mmZ-hh:mmZ` format (UTC time) and must be on the hour with a 1-hour interval, such as `14:00Z-15:00Z`.
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
