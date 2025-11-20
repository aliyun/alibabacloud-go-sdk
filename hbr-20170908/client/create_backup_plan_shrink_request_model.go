// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupPlanShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupType(v string) *CreateBackupPlanShrinkRequest
	GetBackupType() *string
	SetBucket(v string) *CreateBackupPlanShrinkRequest
	GetBucket() *string
	SetChangeListPath(v string) *CreateBackupPlanShrinkRequest
	GetChangeListPath() *string
	SetClusterId(v string) *CreateBackupPlanShrinkRequest
	GetClusterId() *string
	SetCreateTime(v int64) *CreateBackupPlanShrinkRequest
	GetCreateTime() *int64
	SetCrossAccountRoleName(v string) *CreateBackupPlanShrinkRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *CreateBackupPlanShrinkRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *CreateBackupPlanShrinkRequest
	GetCrossAccountUserId() *int64
	SetDataSourceId(v string) *CreateBackupPlanShrinkRequest
	GetDataSourceId() *string
	SetDestDataSourceDetailShrink(v string) *CreateBackupPlanShrinkRequest
	GetDestDataSourceDetailShrink() *string
	SetDestDataSourceId(v string) *CreateBackupPlanShrinkRequest
	GetDestDataSourceId() *string
	SetDestSourceType(v string) *CreateBackupPlanShrinkRequest
	GetDestSourceType() *string
	SetDetailShrink(v string) *CreateBackupPlanShrinkRequest
	GetDetailShrink() *string
	SetDisabled(v bool) *CreateBackupPlanShrinkRequest
	GetDisabled() *bool
	SetExclude(v string) *CreateBackupPlanShrinkRequest
	GetExclude() *string
	SetFileSystemId(v string) *CreateBackupPlanShrinkRequest
	GetFileSystemId() *string
	SetInclude(v string) *CreateBackupPlanShrinkRequest
	GetInclude() *string
	SetInstanceId(v string) *CreateBackupPlanShrinkRequest
	GetInstanceId() *string
	SetInstanceName(v string) *CreateBackupPlanShrinkRequest
	GetInstanceName() *string
	SetKeepLatestSnapshots(v int64) *CreateBackupPlanShrinkRequest
	GetKeepLatestSnapshots() *int64
	SetOptions(v string) *CreateBackupPlanShrinkRequest
	GetOptions() *string
	SetOtsDetailShrink(v string) *CreateBackupPlanShrinkRequest
	GetOtsDetailShrink() *string
	SetPath(v []*string) *CreateBackupPlanShrinkRequest
	GetPath() []*string
	SetPlanName(v string) *CreateBackupPlanShrinkRequest
	GetPlanName() *string
	SetPrefix(v string) *CreateBackupPlanShrinkRequest
	GetPrefix() *string
	SetRetention(v int64) *CreateBackupPlanShrinkRequest
	GetRetention() *int64
	SetRule(v []*CreateBackupPlanShrinkRequestRule) *CreateBackupPlanShrinkRequest
	GetRule() []*CreateBackupPlanShrinkRequestRule
	SetSchedule(v string) *CreateBackupPlanShrinkRequest
	GetSchedule() *string
	SetSourceType(v string) *CreateBackupPlanShrinkRequest
	GetSourceType() *string
	SetSpeedLimit(v string) *CreateBackupPlanShrinkRequest
	GetSpeedLimit() *string
	SetUdmRegionId(v string) *CreateBackupPlanShrinkRequest
	GetUdmRegionId() *string
	SetVaultId(v string) *CreateBackupPlanShrinkRequest
	GetVaultId() *string
}

type CreateBackupPlanShrinkRequest struct {
	// Backup type. Value: **COMPLETE**, indicating a full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// This parameter is required when **SourceType*	- is set to **OSS**. It represents the OSS bucket name.
	//
	// example:
	//
	// hbr-backup-oss
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// Configuration for the incremental file synchronization list. (Required only for synchronization)
	//
	// example:
	//
	// {"dataSourceId": "ds-123456789", "path": "/changelist"}
	ChangeListPath *string `json:"ChangeListPath,omitempty" xml:"ChangeListPath,omitempty"`
	// The ID of the client group that executes the data synchronization plan. This parameter is required only for data synchronization.
	//
	// example:
	//
	// cl-***************
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required when **SourceType*	- is set to **NAS**. It represents the creation time of the file system, in UNIX timestamp, in seconds.
	//
	// example:
	//
	// 1607436917
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The role name created in the RAM of the original account for cross-account backup.
	//
	// example:
	//
	// BackupRole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// Cross-account backup type. Supported values:
	//
	// - SELF_ACCOUNT: Backup within the same account
	//
	// - CROSS_ACCOUNT: Cross-account backup
	//
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The original account ID used for cross-account backup.
	//
	// example:
	//
	// 15897534xxxx4625
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The ID of the data source. This parameter is required only for data synchronization.
	//
	// example:
	//
	// ds-****************
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Destination data source details. (Required only for synchronization)
	//
	// example:
	//
	// {\\"prefix\\":\\"/\\"}
	DestDataSourceDetailShrink *string `json:"DestDataSourceDetail,omitempty" xml:"DestDataSourceDetail,omitempty"`
	// Destination data source ID. (Required only for synchronization)
	//
	// example:
	//
	// ds-*********************
	DestDataSourceId *string `json:"DestDataSourceId,omitempty" xml:"DestDataSourceId,omitempty"`
	// Destination data source type. (Required only for synchronization)
	//
	// example:
	//
	// OSS
	DestSourceType *string `json:"DestSourceType,omitempty" xml:"DestSourceType,omitempty"`
	// Details of the whole machine backup, in JSON string format.
	//
	// 	- snapshotGroup: Whether to use a consistent snapshot group (only valid if all instance disks are ESSD).
	//
	// 	- appConsistent: Whether to use application consistency (requires the use of preScriptPath and postScriptPath parameters).
	//
	// 	- preScriptPath: Path to the freeze script.
	//
	// 	- postScriptPath: Path to the thaw script.
	//
	// example:
	//
	// {\\"EnableFsFreeze\\":true,\\"appConsistent\\":false,\\"postScriptPath\\":\\"\\",\\"preScriptPath\\":\\"\\",\\"snapshotGroup\\":true,\\"timeoutInSeconds\\":60}
	DetailShrink *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// Is the plan disabled by default
	//
	// example:
	//
	// true
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **ECS_FILE**. It specifies the path that should not be backed up, meaning all files under this path will not be included in the backup. The maximum length is 255 characters.
	//
	// example:
	//
	// ["/var", "/proc"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// This parameter is required when **SourceType*	- is set to **NAS**. It represents the file system ID.
	//
	// example:
	//
	// 005494
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required when **SourceType*	- is set to **ECS_FILE**. It represents the path to be backed up, and all files under this path will be backed up. Supports up to 255 characters.
	//
	// example:
	//
	// ["/home/alice/*.pdf", "/home/bob/*.txt"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// This parameter is required when **SourceType*	- is set to **ECS_FILE**. It represents the ECS instance ID.
	//
	// example:
	//
	// i-m5e*****6q
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Table store instance name.
	//
	// example:
	//
	// instancename
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Whether to enable retaining at least one backup version.
	//
	// - 0 - Do not retain
	//
	// - 1 - Retain
	//
	// example:
	//
	// 1
	KeepLatestSnapshots *int64 `json:"KeepLatestSnapshots,omitempty" xml:"KeepLatestSnapshots,omitempty"`
	// This parameter is required when **SourceType*	- is set to **ECS_FILE**. It indicates whether to use the Windows system VSS to define the backup path.
	//
	// - This feature only supports Windows type ECS instances.
	//
	// - If there are data changes in the backup source and you need to ensure consistency between the backup data and the source data, you can configure it as `["UseVSS":true]`.
	//
	// - After choosing to use VSS, multiple file directories cannot be backed up simultaneously.
	//
	// example:
	//
	// {"UseVSS":false}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The details about the Tablestore instance.
	OtsDetailShrink *string `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty"`
	// Backup paths.
	Path []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
	// Name of the backup plan. 1 to 64 characters. The name must be unique for each data source type within a single backup vault.
	//
	// example:
	//
	// planname
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// This parameter is required when **SourceType*	- is set to **OSS**. It represents the backup prefix. When specified, only objects matching the prefix are backed up.
	//
	// example:
	//
	// oss-prefix
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// Number of days to retain the backup, with a minimum value of 1, in days.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// Backup plan rules.
	Rule []*CreateBackupPlanShrinkRequestRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
	// Backup policy. Optional format: `I|{startTime}|{interval}`. This indicates that a backup task will be executed every `{interval}` starting from `{startTime}`. It does not compensate for missed backup tasks due to past time. If the previous backup task has not been completed, the next backup task will not be triggered. For example, `I|1631685600|P1D` means a backup is performed every day starting from 2021-09-15 14:00:00.
	//
	// - **startTime**: Start time of the backup, in UNIX timestamp, in seconds.
	//
	// - **interval**: ISO8601 time interval. For example, PT1H indicates an interval of one hour, and P1D indicates an interval of one day.
	//
	// example:
	//
	// I|1602673264|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **ECS_FILE**: Elastic Compute Service (ECS) files
	//
	// 	- **OSS**: Object Storage Service (OSS) buckets
	//
	// 	- **NAS**: File Storage NAS (NAS) file systems
	//
	// 	- **OTS**: Tablestore instances
	//
	// 	- **UDM_ECS**: ECS instances
	//
	// 	- **SYNC**: data synchronization
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required when **SourceType*	- is set to **ECS_FILE**. It represents the backup traffic control. Format: `{start}:{end}:{bandwidth}`. Multiple traffic control configurations are separated by |, and the configured times should not overlap.
	//
	// - **start**: Start hour.
	//
	// - **end**: End hour.
	//
	// - **bandwidth**: Limit rate, in KB/s.
	//
	// example:
	//
	// 0:24:5120
	SpeedLimit *string `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	// Region where the whole machine backup instance is located.
	//
	// example:
	//
	// cn-shanghai
	UdmRegionId *string `json:"UdmRegionId,omitempty" xml:"UdmRegionId,omitempty"`
	// Backup vault ID.
	//
	// example:
	//
	// v-0006******q
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateBackupPlanShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPlanShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanShrinkRequest) GetBackupType() *string {
	return s.BackupType
}

func (s *CreateBackupPlanShrinkRequest) GetBucket() *string {
	return s.Bucket
}

func (s *CreateBackupPlanShrinkRequest) GetChangeListPath() *string {
	return s.ChangeListPath
}

func (s *CreateBackupPlanShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateBackupPlanShrinkRequest) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateBackupPlanShrinkRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *CreateBackupPlanShrinkRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *CreateBackupPlanShrinkRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *CreateBackupPlanShrinkRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateBackupPlanShrinkRequest) GetDestDataSourceDetailShrink() *string {
	return s.DestDataSourceDetailShrink
}

func (s *CreateBackupPlanShrinkRequest) GetDestDataSourceId() *string {
	return s.DestDataSourceId
}

func (s *CreateBackupPlanShrinkRequest) GetDestSourceType() *string {
	return s.DestSourceType
}

func (s *CreateBackupPlanShrinkRequest) GetDetailShrink() *string {
	return s.DetailShrink
}

func (s *CreateBackupPlanShrinkRequest) GetDisabled() *bool {
	return s.Disabled
}

func (s *CreateBackupPlanShrinkRequest) GetExclude() *string {
	return s.Exclude
}

func (s *CreateBackupPlanShrinkRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateBackupPlanShrinkRequest) GetInclude() *string {
	return s.Include
}

func (s *CreateBackupPlanShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateBackupPlanShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateBackupPlanShrinkRequest) GetKeepLatestSnapshots() *int64 {
	return s.KeepLatestSnapshots
}

func (s *CreateBackupPlanShrinkRequest) GetOptions() *string {
	return s.Options
}

func (s *CreateBackupPlanShrinkRequest) GetOtsDetailShrink() *string {
	return s.OtsDetailShrink
}

func (s *CreateBackupPlanShrinkRequest) GetPath() []*string {
	return s.Path
}

func (s *CreateBackupPlanShrinkRequest) GetPlanName() *string {
	return s.PlanName
}

func (s *CreateBackupPlanShrinkRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *CreateBackupPlanShrinkRequest) GetRetention() *int64 {
	return s.Retention
}

func (s *CreateBackupPlanShrinkRequest) GetRule() []*CreateBackupPlanShrinkRequestRule {
	return s.Rule
}

func (s *CreateBackupPlanShrinkRequest) GetSchedule() *string {
	return s.Schedule
}

func (s *CreateBackupPlanShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateBackupPlanShrinkRequest) GetSpeedLimit() *string {
	return s.SpeedLimit
}

func (s *CreateBackupPlanShrinkRequest) GetUdmRegionId() *string {
	return s.UdmRegionId
}

func (s *CreateBackupPlanShrinkRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *CreateBackupPlanShrinkRequest) SetBackupType(v string) *CreateBackupPlanShrinkRequest {
	s.BackupType = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetBucket(v string) *CreateBackupPlanShrinkRequest {
	s.Bucket = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetChangeListPath(v string) *CreateBackupPlanShrinkRequest {
	s.ChangeListPath = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetClusterId(v string) *CreateBackupPlanShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetCreateTime(v int64) *CreateBackupPlanShrinkRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetCrossAccountRoleName(v string) *CreateBackupPlanShrinkRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetCrossAccountType(v string) *CreateBackupPlanShrinkRequest {
	s.CrossAccountType = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetCrossAccountUserId(v int64) *CreateBackupPlanShrinkRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetDataSourceId(v string) *CreateBackupPlanShrinkRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetDestDataSourceDetailShrink(v string) *CreateBackupPlanShrinkRequest {
	s.DestDataSourceDetailShrink = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetDestDataSourceId(v string) *CreateBackupPlanShrinkRequest {
	s.DestDataSourceId = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetDestSourceType(v string) *CreateBackupPlanShrinkRequest {
	s.DestSourceType = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetDetailShrink(v string) *CreateBackupPlanShrinkRequest {
	s.DetailShrink = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetDisabled(v bool) *CreateBackupPlanShrinkRequest {
	s.Disabled = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetExclude(v string) *CreateBackupPlanShrinkRequest {
	s.Exclude = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetFileSystemId(v string) *CreateBackupPlanShrinkRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetInclude(v string) *CreateBackupPlanShrinkRequest {
	s.Include = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetInstanceId(v string) *CreateBackupPlanShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetInstanceName(v string) *CreateBackupPlanShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetKeepLatestSnapshots(v int64) *CreateBackupPlanShrinkRequest {
	s.KeepLatestSnapshots = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetOptions(v string) *CreateBackupPlanShrinkRequest {
	s.Options = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetOtsDetailShrink(v string) *CreateBackupPlanShrinkRequest {
	s.OtsDetailShrink = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetPath(v []*string) *CreateBackupPlanShrinkRequest {
	s.Path = v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetPlanName(v string) *CreateBackupPlanShrinkRequest {
	s.PlanName = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetPrefix(v string) *CreateBackupPlanShrinkRequest {
	s.Prefix = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetRetention(v int64) *CreateBackupPlanShrinkRequest {
	s.Retention = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetRule(v []*CreateBackupPlanShrinkRequestRule) *CreateBackupPlanShrinkRequest {
	s.Rule = v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetSchedule(v string) *CreateBackupPlanShrinkRequest {
	s.Schedule = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetSourceType(v string) *CreateBackupPlanShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetSpeedLimit(v string) *CreateBackupPlanShrinkRequest {
	s.SpeedLimit = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetUdmRegionId(v string) *CreateBackupPlanShrinkRequest {
	s.UdmRegionId = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) SetVaultId(v string) *CreateBackupPlanShrinkRequest {
	s.VaultId = &v
	return s
}

func (s *CreateBackupPlanShrinkRequest) Validate() error {
	if s.Rule != nil {
		for _, item := range s.Rule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateBackupPlanShrinkRequestRule struct {
	// Backup type.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// ID of the region for offsite replication.
	//
	// example:
	//
	// cn-hangzhou
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// Number of days to retain offsite backups.
	//
	// example:
	//
	// 7
	DestinationRetention *int64 `json:"DestinationRetention,omitempty" xml:"DestinationRetention,omitempty"`
	// Whether the rule is enabled.
	//
	// example:
	//
	// true
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// Whether to enable offsite replication.
	//
	// example:
	//
	// true
	DoCopy *bool `json:"DoCopy,omitempty" xml:"DoCopy,omitempty"`
	// Backup retention period.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// Rule name.
	//
	// example:
	//
	// rule-test-name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Backup strategy. Optional format: I|{startTime}|{interval}. This means that a backup task is executed every {interval} starting from {startTime}. Backup tasks for past times will not be executed. If the previous backup task has not been completed, the next backup task will not be triggered. For example, I|1631685600|P1D means a backup is performed every day starting from 2021-09-15 14:00:00.
	//
	// - startTime: The start time of the backup, in UNIX time, in seconds.
	//
	// - interval: ISO8601 time interval. For example, PT1H means an interval of one hour. P1D means an interval of one day.
	//
	// example:
	//
	// I|1602673264|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s CreateBackupPlanShrinkRequestRule) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPlanShrinkRequestRule) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanShrinkRequestRule) GetBackupType() *string {
	return s.BackupType
}

func (s *CreateBackupPlanShrinkRequestRule) GetDestinationRegionId() *string {
	return s.DestinationRegionId
}

func (s *CreateBackupPlanShrinkRequestRule) GetDestinationRetention() *int64 {
	return s.DestinationRetention
}

func (s *CreateBackupPlanShrinkRequestRule) GetDisabled() *bool {
	return s.Disabled
}

func (s *CreateBackupPlanShrinkRequestRule) GetDoCopy() *bool {
	return s.DoCopy
}

func (s *CreateBackupPlanShrinkRequestRule) GetRetention() *int64 {
	return s.Retention
}

func (s *CreateBackupPlanShrinkRequestRule) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateBackupPlanShrinkRequestRule) GetSchedule() *string {
	return s.Schedule
}

func (s *CreateBackupPlanShrinkRequestRule) SetBackupType(v string) *CreateBackupPlanShrinkRequestRule {
	s.BackupType = &v
	return s
}

func (s *CreateBackupPlanShrinkRequestRule) SetDestinationRegionId(v string) *CreateBackupPlanShrinkRequestRule {
	s.DestinationRegionId = &v
	return s
}

func (s *CreateBackupPlanShrinkRequestRule) SetDestinationRetention(v int64) *CreateBackupPlanShrinkRequestRule {
	s.DestinationRetention = &v
	return s
}

func (s *CreateBackupPlanShrinkRequestRule) SetDisabled(v bool) *CreateBackupPlanShrinkRequestRule {
	s.Disabled = &v
	return s
}

func (s *CreateBackupPlanShrinkRequestRule) SetDoCopy(v bool) *CreateBackupPlanShrinkRequestRule {
	s.DoCopy = &v
	return s
}

func (s *CreateBackupPlanShrinkRequestRule) SetRetention(v int64) *CreateBackupPlanShrinkRequestRule {
	s.Retention = &v
	return s
}

func (s *CreateBackupPlanShrinkRequestRule) SetRuleName(v string) *CreateBackupPlanShrinkRequestRule {
	s.RuleName = &v
	return s
}

func (s *CreateBackupPlanShrinkRequestRule) SetSchedule(v string) *CreateBackupPlanShrinkRequestRule {
	s.Schedule = &v
	return s
}

func (s *CreateBackupPlanShrinkRequestRule) Validate() error {
	return dara.Validate(s)
}
