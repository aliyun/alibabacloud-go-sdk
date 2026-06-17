// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedDataPolicies(v []*ModifyBackupPolicyRequestAdvancedDataPolicies) *ModifyBackupPolicyRequest
	GetAdvancedDataPolicies() []*ModifyBackupPolicyRequestAdvancedDataPolicies
	SetBackupFrequency(v string) *ModifyBackupPolicyRequest
	GetBackupFrequency() *string
	SetBackupPolicyLevel(v string) *ModifyBackupPolicyRequest
	GetBackupPolicyLevel() *string
	SetBackupRetentionPolicyOnClusterDeletion(v string) *ModifyBackupPolicyRequest
	GetBackupRetentionPolicyOnClusterDeletion() *string
	SetDBClusterId(v string) *ModifyBackupPolicyRequest
	GetDBClusterId() *string
	SetDataLevel1BackupFrequency(v string) *ModifyBackupPolicyRequest
	GetDataLevel1BackupFrequency() *string
	SetDataLevel1BackupPeriod(v string) *ModifyBackupPolicyRequest
	GetDataLevel1BackupPeriod() *string
	SetDataLevel1BackupRetentionPeriod(v string) *ModifyBackupPolicyRequest
	GetDataLevel1BackupRetentionPeriod() *string
	SetDataLevel1BackupTime(v string) *ModifyBackupPolicyRequest
	GetDataLevel1BackupTime() *string
	SetDataLevel2BackupAnotherRegionRegion(v string) *ModifyBackupPolicyRequest
	GetDataLevel2BackupAnotherRegionRegion() *string
	SetDataLevel2BackupAnotherRegionRetentionPeriod(v string) *ModifyBackupPolicyRequest
	GetDataLevel2BackupAnotherRegionRetentionPeriod() *string
	SetDataLevel2BackupPeriod(v string) *ModifyBackupPolicyRequest
	GetDataLevel2BackupPeriod() *string
	SetDataLevel2BackupRetentionPeriod(v string) *ModifyBackupPolicyRequest
	GetDataLevel2BackupRetentionPeriod() *string
	SetOwnerAccount(v string) *ModifyBackupPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyBackupPolicyRequest
	GetOwnerId() *int64
	SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupTime() *string
	SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest
	GetResourceOwnerId() *int64
}

type ModifyBackupPolicyRequest struct {
	// The advanced backup policies.
	//
	// > - - This parameter is not supported for PolarDB for PostgreSQL (compatible with Oracle) and PolarDB for PostgreSQL.
	//
	// >
	//
	// > - - This parameter is supported only for clusters for which `BackupPolicyLevel` is set to `Advanced`.
	AdvancedDataPolicies []*ModifyBackupPolicyRequestAdvancedDataPolicies `json:"AdvancedDataPolicies,omitempty" xml:"AdvancedDataPolicies,omitempty" type:"Repeated"`
	// The backup frequency. Valid values:
	//
	// - **Normal*	- (default): standard backup. The cluster is backed up once a day.
	//
	// - **2/24H**: high-frequency backup. The cluster is backed up every 2 hours.
	//
	// - **3/24H**: high-frequency backup. The cluster is backed up every 3 hours.
	//
	// - **4/24H**: high-frequency backup. The cluster is backed up every 4 hours.
	//
	// > 	- 	- If you enable high-frequency backup, all backups completed within the last 24 hours are retained. For backups older than 24 hours, the system retains only the first backup completed after 00:00 each day and deletes the rest.
	//
	// >
	//
	// > 	- - If you enable high-frequency backup, the **PreferredBackupPeriod*	- parameter is automatically set to all days of the week (from Monday to Sunday).
	//
	// >
	//
	// > 	- - This parameter is not supported if your PolarDB for MySQL cluster is in a region that supports the cross-region backup feature. For more information about the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// >
	//
	// > 	- - After you enable advanced backup, this parameter is no longer effective. Use the `AdvancedDataPolicies` parameter instead.
	//
	// example:
	//
	// Normal
	BackupFrequency *string `json:"BackupFrequency,omitempty" xml:"BackupFrequency,omitempty"`
	// The level of the backup policy. Valid values:
	//
	// - **Normal**: standard backup
	//
	// - **Advanced**: advanced backup
	//
	// > 	- 	- This parameter is not supported for PolarDB for PostgreSQL (compatible with Oracle) and PolarDB for PostgreSQL.
	//
	// >
	//
	// > 	- - Check the `AdvancedPolicyOption` parameter in the response of the [DescribeBackupPolicy](https://help.aliyun.com/document_detail/2319231.html) operation to determine whether the cluster supports advanced backup. If the cluster supports advanced backup, you can request this feature in [Advanced backup settings](~611727~~).
	//
	// >
	//
	// > 	- - After you enable advanced backup, you **cannot*	- switch back to standard backup.
	//
	// example:
	//
	// Normal
	BackupPolicyLevel *string `json:"BackupPolicyLevel,omitempty" xml:"BackupPolicyLevel,omitempty"`
	// Specifies whether to retain backups when you delete the cluster. Valid values:
	//
	// - **ALL**: Permanently retains all backups.
	//
	// - **LATEST**: Permanently retains the last backup.
	//
	// - **NONE**: Does not retain backup sets.
	//
	// > The default value is `NONE`.
	//
	// example:
	//
	// NONE
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of all clusters in a specified region, including the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp13wz9586voc****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The backup frequency. Valid values:
	//
	// - **Normal*	- (default): standard backup. The cluster is backed up once a day.
	//
	// - **2/24H**: high-frequency backup. The cluster is backed up every 2 hours.
	//
	// - **3/24H**: high-frequency backup. The cluster is backed up every 3 hours.
	//
	// - **4/24H**: high-frequency backup. The cluster is backed up every 4 hours.
	//
	// > 	- 	- This parameter is not supported for PolarDB for PostgreSQL (compatible with Oracle) and PolarDB for PostgreSQL.
	//
	// >
	//
	// > 	- - This parameter is not supported if your PolarDB for MySQL cluster is in a region that supports the cross-region backup feature. For more information about the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// >
	//
	// > 	- - After you enable advanced backup, this parameter is no longer effective. Use the `AdvancedDataPolicies` parameter instead.
	//
	// example:
	//
	// Normal
	DataLevel1BackupFrequency *string `json:"DataLevel1BackupFrequency,omitempty" xml:"DataLevel1BackupFrequency,omitempty"`
	// The level-1 backup cycle. Valid values:
	//
	// - **Monday**
	//
	// - **Tuesday**
	//
	// - **Wednesday**
	//
	// - **Thursday**
	//
	// - **Friday**
	//
	// - **Saturday**
	//
	// - **Sunday**
	//
	// > 	- 	- You must select at least two days. Separate multiple values with commas.
	//
	// >
	//
	// > 	- - This parameter is not supported for PolarDB for PostgreSQL (compatible with Oracle) and PolarDB for PostgreSQL.
	//
	// >
	//
	// > 	- - This parameter is not supported if your PolarDB for MySQL cluster is in a region that supports the cross-region backup feature. For more information about the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// >
	//
	// > 	- - After you enable advanced backup, this parameter is no longer effective. Use the `AdvancedDataPolicies` parameter instead.
	//
	// example:
	//
	// Monday,Tuesday
	DataLevel1BackupPeriod *string `json:"DataLevel1BackupPeriod,omitempty" xml:"DataLevel1BackupPeriod,omitempty"`
	// The retention period for level-1 backups, in days. Valid values: 3 to 14.
	//
	// > - After you enable advanced backup, this parameter is no longer effective. Use the `AdvancedDataPolicies` parameter instead.
	//
	// example:
	//
	// 3
	DataLevel1BackupRetentionPeriod *string `json:"DataLevel1BackupRetentionPeriod,omitempty" xml:"DataLevel1BackupRetentionPeriod,omitempty"`
	// The time window for automatic backups. Specify the time in UTC and in the `hh:mmZ-hh:mmZ` format. The time window must be a one-hour period that starts on the hour. For example, `14:00Z-15:00Z`.
	//
	// > - This parameter is not supported for PolarDB for PostgreSQL (compatible with Oracle) and PolarDB for PostgreSQL.
	//
	// >
	//
	// > - This parameter is not supported if your PolarDB for MySQL cluster is in a region that supports the cross-region backup feature. For more information about the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// example:
	//
	// 15:00Z-16:00Z
	DataLevel1BackupTime *string `json:"DataLevel1BackupTime,omitempty" xml:"DataLevel1BackupTime,omitempty"`
	// The destination region for the cross-region level-2 backup. For more information about the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// > - After you enable advanced backup, this parameter is no longer effective. Use the `AdvancedDataPolicies` parameter instead.
	//
	// example:
	//
	// cn-hangzhou
	DataLevel2BackupAnotherRegionRegion *string `json:"DataLevel2BackupAnotherRegionRegion,omitempty" xml:"DataLevel2BackupAnotherRegionRegion,omitempty"`
	// The retention period of cross-region level-2 backups. Valid values:
	//
	// - **0**: Disables the cross-region level-2 backup feature.
	//
	// - **30 to 7300**: The retention period of cross-region level-2 backups, in days.
	//
	// - **-1**: Cross-region level-2 backups are permanently retained.
	//
	// > 	- 	- When you create a cluster, the default value is **0**, which disables the cross-region level-2 backup feature.
	//
	// >
	//
	// > 	- - After you enable advanced backup, this parameter is no longer effective. Use the `AdvancedDataPolicies` parameter instead.
	//
	// example:
	//
	// 30
	DataLevel2BackupAnotherRegionRetentionPeriod *string `json:"DataLevel2BackupAnotherRegionRetentionPeriod,omitempty" xml:"DataLevel2BackupAnotherRegionRetentionPeriod,omitempty"`
	// The level-2 backup cycle. Valid values:
	//
	// - **Monday**
	//
	// - **Tuesday**
	//
	// - **Wednesday**
	//
	// - **Thursday**
	//
	// - **Friday**
	//
	// - **Saturday**
	//
	// - **Sunday**
	//
	// > 	- 	- You must select at least two days. Separate multiple values with commas.
	//
	// >
	//
	// > 	- - This parameter is not supported for PolarDB for PostgreSQL (compatible with Oracle) and PolarDB for PostgreSQL.
	//
	// >
	//
	// > 	- - This parameter is not supported if your PolarDB for MySQL cluster is in a region that supports the cross-region backup feature. For more information about the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// >
	//
	// > 	- - After you enable advanced backup, this parameter is no longer effective. Use the `AdvancedDataPolicies` parameter instead.
	//
	// example:
	//
	// Monday,Tuesday
	DataLevel2BackupPeriod *string `json:"DataLevel2BackupPeriod,omitempty" xml:"DataLevel2BackupPeriod,omitempty"`
	// The retention period of level-2 backups. Valid values:
	//
	// - **0**: Disables the level-2 backup feature.
	//
	// - **30 to 7300**: The retention period of level-2 backups, in days.
	//
	// - **-1**: Level-2 backups are permanently retained.
	//
	// > 	- 	- When you create a cluster, the default value is **0**, which disables the level-2 backup feature.
	//
	// >
	//
	// > 	- - After you enable advanced backup, this parameter is no longer effective. Use the `AdvancedDataPolicies` parameter instead.
	//
	// example:
	//
	// 0
	DataLevel2BackupRetentionPeriod *string `json:"DataLevel2BackupRetentionPeriod,omitempty" xml:"DataLevel2BackupRetentionPeriod,omitempty"`
	OwnerAccount                    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The backup cycle. Valid values:
	//
	// - **Monday**
	//
	// - **Tuesday**
	//
	// - **Wednesday**
	//
	// - **Thursday**
	//
	// - **Friday**
	//
	// - **Saturday**
	//
	// - **Sunday**
	//
	// > 	- 	- You must select at least two days. Separate multiple values with commas.
	//
	// >
	//
	// > 	- - This parameter is not supported if your PolarDB for MySQL cluster is in a region that supports the cross-region backup feature. For more information about the regions that support cross-region backup, see [Overview](https://help.aliyun.com/document_detail/72672.html).
	//
	// >
	//
	// > 	- - After you enable advanced backup, this parameter is no longer effective. Use the `AdvancedDataPolicies` parameter instead.
	//
	// example:
	//
	// Monday,Tuesday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The time window for automatic backups. Specify the time in UTC and in the `hh:mmZ-hh:mmZ` format. The time window must be a one-hour period that starts on the hour. For example, `14:00Z-15:00Z`.
	//
	// example:
	//
	// 15:00Z-16:00Z
	PreferredBackupTime  *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) GetAdvancedDataPolicies() []*ModifyBackupPolicyRequestAdvancedDataPolicies {
	return s.AdvancedDataPolicies
}

func (s *ModifyBackupPolicyRequest) GetBackupFrequency() *string {
	return s.BackupFrequency
}

func (s *ModifyBackupPolicyRequest) GetBackupPolicyLevel() *string {
	return s.BackupPolicyLevel
}

func (s *ModifyBackupPolicyRequest) GetBackupRetentionPolicyOnClusterDeletion() *string {
	return s.BackupRetentionPolicyOnClusterDeletion
}

func (s *ModifyBackupPolicyRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyBackupPolicyRequest) GetDataLevel1BackupFrequency() *string {
	return s.DataLevel1BackupFrequency
}

func (s *ModifyBackupPolicyRequest) GetDataLevel1BackupPeriod() *string {
	return s.DataLevel1BackupPeriod
}

func (s *ModifyBackupPolicyRequest) GetDataLevel1BackupRetentionPeriod() *string {
	return s.DataLevel1BackupRetentionPeriod
}

func (s *ModifyBackupPolicyRequest) GetDataLevel1BackupTime() *string {
	return s.DataLevel1BackupTime
}

func (s *ModifyBackupPolicyRequest) GetDataLevel2BackupAnotherRegionRegion() *string {
	return s.DataLevel2BackupAnotherRegionRegion
}

func (s *ModifyBackupPolicyRequest) GetDataLevel2BackupAnotherRegionRetentionPeriod() *string {
	return s.DataLevel2BackupAnotherRegionRetentionPeriod
}

func (s *ModifyBackupPolicyRequest) GetDataLevel2BackupPeriod() *string {
	return s.DataLevel2BackupPeriod
}

func (s *ModifyBackupPolicyRequest) GetDataLevel2BackupRetentionPeriod() *string {
	return s.DataLevel2BackupRetentionPeriod
}

func (s *ModifyBackupPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyBackupPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *ModifyBackupPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyBackupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyBackupPolicyRequest) SetAdvancedDataPolicies(v []*ModifyBackupPolicyRequestAdvancedDataPolicies) *ModifyBackupPolicyRequest {
	s.AdvancedDataPolicies = v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupFrequency(v string) *ModifyBackupPolicyRequest {
	s.BackupFrequency = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupPolicyLevel(v string) *ModifyBackupPolicyRequest {
	s.BackupPolicyLevel = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPolicyOnClusterDeletion(v string) *ModifyBackupPolicyRequest {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDBClusterId(v string) *ModifyBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDataLevel1BackupFrequency(v string) *ModifyBackupPolicyRequest {
	s.DataLevel1BackupFrequency = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDataLevel1BackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.DataLevel1BackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDataLevel1BackupRetentionPeriod(v string) *ModifyBackupPolicyRequest {
	s.DataLevel1BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDataLevel1BackupTime(v string) *ModifyBackupPolicyRequest {
	s.DataLevel1BackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDataLevel2BackupAnotherRegionRegion(v string) *ModifyBackupPolicyRequest {
	s.DataLevel2BackupAnotherRegionRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDataLevel2BackupAnotherRegionRetentionPeriod(v string) *ModifyBackupPolicyRequest {
	s.DataLevel2BackupAnotherRegionRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDataLevel2BackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.DataLevel2BackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDataLevel2BackupRetentionPeriod(v string) *ModifyBackupPolicyRequest {
	s.DataLevel2BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) Validate() error {
	if s.AdvancedDataPolicies != nil {
		for _, item := range s.AdvancedDataPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyBackupPolicyRequestAdvancedDataPolicies struct {
	// The action type. Valid values:
	//
	// - **CREATE**: Create
	//
	// - **UPDATE**: Update
	//
	// - **DELETE**: Delete
	//
	// example:
	//
	// CREATE
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// Indicates whether the backup policy is automatically generated by the system.
	//
	// > This value is automatically generated. You do not need to specify this parameter.
	//
	// example:
	//
	// false
	AutoCreated *bool `json:"AutoCreated,omitempty" xml:"AutoCreated,omitempty"`
	// The backup type. Valid values:
	//
	// - **F**: full backup
	//
	// > This parameter cannot be modified and is fixed to `F`.
	//
	// example:
	//
	// F
	BakType *string `json:"BakType,omitempty" xml:"BakType,omitempty"`
	// The destination region for the backup policy.
	//
	// example:
	//
	// cn-beijing
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// The destination type for the backup policy. Valid values:
	//
	// - **level1**: level-1 backup
	//
	// - **level2**: level-2 backup
	//
	// - **level2Cross**: cross-region level-2 backup
	//
	// example:
	//
	// level2
	DestType *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	// The method to convert a level-1 backup to a level-2 backup. Valid values:
	//
	// - **copy**: Copy
	//
	// example:
	//
	// copy
	DumpAction *string `json:"DumpAction,omitempty" xml:"DumpAction,omitempty"`
	// The scheduling type. Valid values:
	//
	// - **dayOfWeek**: Weekly schedule
	//
	// - **dayOfMonth**: Monthly schedule
	//
	// - **dayOfYear**: Yearly schedule
	//
	// - **backupInterval**: Fixed interval schedule
	//
	// > This parameter is required only if `FilterType` is set to **crontab**.
	//
	// example:
	//
	// dayOfWeek
	FilterKey *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	// The filter type for the advanced policy. Valid values:
	//
	// - **crontab**: Recurring schedule
	//
	// - **event**: Event-based schedule
	//
	// example:
	//
	// crontab
	FilterType *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	// The backup cycle.
	//
	// example:
	//
	// 1,2,3,4,5,6,7
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	// The retention policy for backups that are more than 24 hours old. Valid values:
	//
	// - **true**: Retains only the first backup set of the day.
	//
	// - **false**: Retains all backups.
	//
	// example:
	//
	// true
	OnlyPreserveOneEachDay *bool `json:"OnlyPreserveOneEachDay,omitempty" xml:"OnlyPreserveOneEachDay,omitempty"`
	// The retention policy for hourly backups. Valid values:
	//
	// - **true**: Retains only the first backup set of each hour for backups that are more than one hour old.
	//
	// - **false**: Retains all backup sets.
	//
	// > This parameter cannot be modified and is fixed to `true`.
	//
	// example:
	//
	// true
	OnlyPreserveOneEachHour *bool `json:"OnlyPreserveOneEachHour,omitempty" xml:"OnlyPreserveOneEachHour,omitempty"`
	// The ID of the backup policy. You can call the [DescribeBackupPolicy](https://help.aliyun.com/document_detail/2319231.html) operation to view the ID.
	//
	// example:
	//
	// 71930ac2e9f15e41615e10627c******
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The retention type for the backup set. Valid values:
	//
	// - **never**: Never expires.
	//
	// - **delay**: Expires after a specified number of days.
	//
	// example:
	//
	// delay
	RetentionType *string `json:"RetentionType,omitempty" xml:"RetentionType,omitempty"`
	// The number of days to retain the backup.
	//
	// example:
	//
	// 7
	RetentionValue *string `json:"RetentionValue,omitempty" xml:"RetentionValue,omitempty"`
	// The source region for the backup policy.
	//
	// example:
	//
	// cn-shanghai
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The source type for the backup policy. Valid values:
	//
	// - **db**: database cluster
	//
	// - **level1**: level-1 backup
	//
	// - **level2**: level-2 backup
	//
	// - **level2Cross**: cross-region level-2 backup
	//
	// example:
	//
	// level1
	SrcType      *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s ModifyBackupPolicyRequestAdvancedDataPolicies) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequestAdvancedDataPolicies) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetActionType() *string {
	return s.ActionType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetAutoCreated() *bool {
	return s.AutoCreated
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetBakType() *string {
	return s.BakType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetDestType() *string {
	return s.DestType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetDumpAction() *string {
	return s.DumpAction
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetFilterKey() *string {
	return s.FilterKey
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetFilterType() *string {
	return s.FilterType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetFilterValue() *string {
	return s.FilterValue
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetOnlyPreserveOneEachDay() *bool {
	return s.OnlyPreserveOneEachDay
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetOnlyPreserveOneEachHour() *bool {
	return s.OnlyPreserveOneEachHour
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetRetentionType() *string {
	return s.RetentionType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetRetentionValue() *string {
	return s.RetentionValue
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetStorageClass() *string {
	return s.StorageClass
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetActionType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.ActionType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetAutoCreated(v bool) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.AutoCreated = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetBakType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.BakType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetDestRegion(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.DestRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetDestType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.DestType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetDumpAction(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.DumpAction = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetFilterKey(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.FilterKey = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetFilterType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.FilterType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetFilterValue(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.FilterValue = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetOnlyPreserveOneEachDay(v bool) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.OnlyPreserveOneEachDay = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetOnlyPreserveOneEachHour(v bool) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.OnlyPreserveOneEachHour = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetPolicyId(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.PolicyId = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetRetentionType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.RetentionType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetRetentionValue(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.RetentionValue = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetSrcRegion(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.SrcRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetSrcType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.SrcType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetStorageClass(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.StorageClass = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) Validate() error {
	return dara.Validate(s)
}
