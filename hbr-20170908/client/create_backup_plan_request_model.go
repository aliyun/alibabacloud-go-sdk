// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupType(v string) *CreateBackupPlanRequest
	GetBackupType() *string
	SetBucket(v string) *CreateBackupPlanRequest
	GetBucket() *string
	SetChangeListPath(v string) *CreateBackupPlanRequest
	GetChangeListPath() *string
	SetClusterId(v string) *CreateBackupPlanRequest
	GetClusterId() *string
	SetCreateTime(v int64) *CreateBackupPlanRequest
	GetCreateTime() *int64
	SetCrossAccountRoleName(v string) *CreateBackupPlanRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *CreateBackupPlanRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *CreateBackupPlanRequest
	GetCrossAccountUserId() *int64
	SetDataSourceId(v string) *CreateBackupPlanRequest
	GetDataSourceId() *string
	SetDestDataSourceDetail(v map[string]interface{}) *CreateBackupPlanRequest
	GetDestDataSourceDetail() map[string]interface{}
	SetDestDataSourceId(v string) *CreateBackupPlanRequest
	GetDestDataSourceId() *string
	SetDestSourceType(v string) *CreateBackupPlanRequest
	GetDestSourceType() *string
	SetDetail(v map[string]interface{}) *CreateBackupPlanRequest
	GetDetail() map[string]interface{}
	SetDisabled(v bool) *CreateBackupPlanRequest
	GetDisabled() *bool
	SetExclude(v string) *CreateBackupPlanRequest
	GetExclude() *string
	SetFileSystemId(v string) *CreateBackupPlanRequest
	GetFileSystemId() *string
	SetInclude(v string) *CreateBackupPlanRequest
	GetInclude() *string
	SetInstanceId(v string) *CreateBackupPlanRequest
	GetInstanceId() *string
	SetInstanceName(v string) *CreateBackupPlanRequest
	GetInstanceName() *string
	SetKeepLatestSnapshots(v int64) *CreateBackupPlanRequest
	GetKeepLatestSnapshots() *int64
	SetOptions(v string) *CreateBackupPlanRequest
	GetOptions() *string
	SetOtsDetail(v *OtsDetail) *CreateBackupPlanRequest
	GetOtsDetail() *OtsDetail
	SetPath(v []*string) *CreateBackupPlanRequest
	GetPath() []*string
	SetPlanName(v string) *CreateBackupPlanRequest
	GetPlanName() *string
	SetPrefix(v string) *CreateBackupPlanRequest
	GetPrefix() *string
	SetRetention(v int64) *CreateBackupPlanRequest
	GetRetention() *int64
	SetRule(v []*CreateBackupPlanRequestRule) *CreateBackupPlanRequest
	GetRule() []*CreateBackupPlanRequestRule
	SetSchedule(v string) *CreateBackupPlanRequest
	GetSchedule() *string
	SetSourceType(v string) *CreateBackupPlanRequest
	GetSourceType() *string
	SetSpeedLimit(v string) *CreateBackupPlanRequest
	GetSpeedLimit() *string
	SetUdmRegionId(v string) *CreateBackupPlanRequest
	GetUdmRegionId() *string
	SetVaultId(v string) *CreateBackupPlanRequest
	GetVaultId() *string
}

type CreateBackupPlanRequest struct {
	// Backup type. Value: **COMPLETE**, indicating a full backup.
	//
	// This parameter is required.
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
	DestDataSourceDetail map[string]interface{} `json:"DestDataSourceDetail,omitempty" xml:"DestDataSourceDetail,omitempty"`
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
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
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
	OtsDetail *OtsDetail `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty"`
	// Backup paths.
	Path []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
	// Name of the backup plan. 1 to 64 characters. The name must be unique for each data source type within a single backup vault.
	//
	// This parameter is required.
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
	Rule []*CreateBackupPlanRequestRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
	// Backup policy. Optional format: `I|{startTime}|{interval}`. This indicates that a backup task will be executed every `{interval}` starting from `{startTime}`. It does not compensate for missed backup tasks due to past time. If the previous backup task has not been completed, the next backup task will not be triggered. For example, `I|1631685600|P1D` means a backup is performed every day starting from 2021-09-15 14:00:00.
	//
	// - **startTime**: Start time of the backup, in UNIX timestamp, in seconds.
	//
	// - **interval**: ISO8601 time interval. For example, PT1H indicates an interval of one hour, and P1D indicates an interval of one day.
	//
	// This parameter is required.
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

func (s CreateBackupPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanRequest) GetBackupType() *string {
	return s.BackupType
}

func (s *CreateBackupPlanRequest) GetBucket() *string {
	return s.Bucket
}

func (s *CreateBackupPlanRequest) GetChangeListPath() *string {
	return s.ChangeListPath
}

func (s *CreateBackupPlanRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateBackupPlanRequest) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateBackupPlanRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *CreateBackupPlanRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *CreateBackupPlanRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *CreateBackupPlanRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateBackupPlanRequest) GetDestDataSourceDetail() map[string]interface{} {
	return s.DestDataSourceDetail
}

func (s *CreateBackupPlanRequest) GetDestDataSourceId() *string {
	return s.DestDataSourceId
}

func (s *CreateBackupPlanRequest) GetDestSourceType() *string {
	return s.DestSourceType
}

func (s *CreateBackupPlanRequest) GetDetail() map[string]interface{} {
	return s.Detail
}

func (s *CreateBackupPlanRequest) GetDisabled() *bool {
	return s.Disabled
}

func (s *CreateBackupPlanRequest) GetExclude() *string {
	return s.Exclude
}

func (s *CreateBackupPlanRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateBackupPlanRequest) GetInclude() *string {
	return s.Include
}

func (s *CreateBackupPlanRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateBackupPlanRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateBackupPlanRequest) GetKeepLatestSnapshots() *int64 {
	return s.KeepLatestSnapshots
}

func (s *CreateBackupPlanRequest) GetOptions() *string {
	return s.Options
}

func (s *CreateBackupPlanRequest) GetOtsDetail() *OtsDetail {
	return s.OtsDetail
}

func (s *CreateBackupPlanRequest) GetPath() []*string {
	return s.Path
}

func (s *CreateBackupPlanRequest) GetPlanName() *string {
	return s.PlanName
}

func (s *CreateBackupPlanRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *CreateBackupPlanRequest) GetRetention() *int64 {
	return s.Retention
}

func (s *CreateBackupPlanRequest) GetRule() []*CreateBackupPlanRequestRule {
	return s.Rule
}

func (s *CreateBackupPlanRequest) GetSchedule() *string {
	return s.Schedule
}

func (s *CreateBackupPlanRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateBackupPlanRequest) GetSpeedLimit() *string {
	return s.SpeedLimit
}

func (s *CreateBackupPlanRequest) GetUdmRegionId() *string {
	return s.UdmRegionId
}

func (s *CreateBackupPlanRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *CreateBackupPlanRequest) SetBackupType(v string) *CreateBackupPlanRequest {
	s.BackupType = &v
	return s
}

func (s *CreateBackupPlanRequest) SetBucket(v string) *CreateBackupPlanRequest {
	s.Bucket = &v
	return s
}

func (s *CreateBackupPlanRequest) SetChangeListPath(v string) *CreateBackupPlanRequest {
	s.ChangeListPath = &v
	return s
}

func (s *CreateBackupPlanRequest) SetClusterId(v string) *CreateBackupPlanRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateBackupPlanRequest) SetCreateTime(v int64) *CreateBackupPlanRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateBackupPlanRequest) SetCrossAccountRoleName(v string) *CreateBackupPlanRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *CreateBackupPlanRequest) SetCrossAccountType(v string) *CreateBackupPlanRequest {
	s.CrossAccountType = &v
	return s
}

func (s *CreateBackupPlanRequest) SetCrossAccountUserId(v int64) *CreateBackupPlanRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *CreateBackupPlanRequest) SetDataSourceId(v string) *CreateBackupPlanRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateBackupPlanRequest) SetDestDataSourceDetail(v map[string]interface{}) *CreateBackupPlanRequest {
	s.DestDataSourceDetail = v
	return s
}

func (s *CreateBackupPlanRequest) SetDestDataSourceId(v string) *CreateBackupPlanRequest {
	s.DestDataSourceId = &v
	return s
}

func (s *CreateBackupPlanRequest) SetDestSourceType(v string) *CreateBackupPlanRequest {
	s.DestSourceType = &v
	return s
}

func (s *CreateBackupPlanRequest) SetDetail(v map[string]interface{}) *CreateBackupPlanRequest {
	s.Detail = v
	return s
}

func (s *CreateBackupPlanRequest) SetDisabled(v bool) *CreateBackupPlanRequest {
	s.Disabled = &v
	return s
}

func (s *CreateBackupPlanRequest) SetExclude(v string) *CreateBackupPlanRequest {
	s.Exclude = &v
	return s
}

func (s *CreateBackupPlanRequest) SetFileSystemId(v string) *CreateBackupPlanRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateBackupPlanRequest) SetInclude(v string) *CreateBackupPlanRequest {
	s.Include = &v
	return s
}

func (s *CreateBackupPlanRequest) SetInstanceId(v string) *CreateBackupPlanRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBackupPlanRequest) SetInstanceName(v string) *CreateBackupPlanRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateBackupPlanRequest) SetKeepLatestSnapshots(v int64) *CreateBackupPlanRequest {
	s.KeepLatestSnapshots = &v
	return s
}

func (s *CreateBackupPlanRequest) SetOptions(v string) *CreateBackupPlanRequest {
	s.Options = &v
	return s
}

func (s *CreateBackupPlanRequest) SetOtsDetail(v *OtsDetail) *CreateBackupPlanRequest {
	s.OtsDetail = v
	return s
}

func (s *CreateBackupPlanRequest) SetPath(v []*string) *CreateBackupPlanRequest {
	s.Path = v
	return s
}

func (s *CreateBackupPlanRequest) SetPlanName(v string) *CreateBackupPlanRequest {
	s.PlanName = &v
	return s
}

func (s *CreateBackupPlanRequest) SetPrefix(v string) *CreateBackupPlanRequest {
	s.Prefix = &v
	return s
}

func (s *CreateBackupPlanRequest) SetRetention(v int64) *CreateBackupPlanRequest {
	s.Retention = &v
	return s
}

func (s *CreateBackupPlanRequest) SetRule(v []*CreateBackupPlanRequestRule) *CreateBackupPlanRequest {
	s.Rule = v
	return s
}

func (s *CreateBackupPlanRequest) SetSchedule(v string) *CreateBackupPlanRequest {
	s.Schedule = &v
	return s
}

func (s *CreateBackupPlanRequest) SetSourceType(v string) *CreateBackupPlanRequest {
	s.SourceType = &v
	return s
}

func (s *CreateBackupPlanRequest) SetSpeedLimit(v string) *CreateBackupPlanRequest {
	s.SpeedLimit = &v
	return s
}

func (s *CreateBackupPlanRequest) SetUdmRegionId(v string) *CreateBackupPlanRequest {
	s.UdmRegionId = &v
	return s
}

func (s *CreateBackupPlanRequest) SetVaultId(v string) *CreateBackupPlanRequest {
	s.VaultId = &v
	return s
}

func (s *CreateBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}

type CreateBackupPlanRequestRule struct {
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

func (s CreateBackupPlanRequestRule) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPlanRequestRule) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanRequestRule) GetBackupType() *string {
	return s.BackupType
}

func (s *CreateBackupPlanRequestRule) GetDestinationRegionId() *string {
	return s.DestinationRegionId
}

func (s *CreateBackupPlanRequestRule) GetDestinationRetention() *int64 {
	return s.DestinationRetention
}

func (s *CreateBackupPlanRequestRule) GetDisabled() *bool {
	return s.Disabled
}

func (s *CreateBackupPlanRequestRule) GetDoCopy() *bool {
	return s.DoCopy
}

func (s *CreateBackupPlanRequestRule) GetRetention() *int64 {
	return s.Retention
}

func (s *CreateBackupPlanRequestRule) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateBackupPlanRequestRule) GetSchedule() *string {
	return s.Schedule
}

func (s *CreateBackupPlanRequestRule) SetBackupType(v string) *CreateBackupPlanRequestRule {
	s.BackupType = &v
	return s
}

func (s *CreateBackupPlanRequestRule) SetDestinationRegionId(v string) *CreateBackupPlanRequestRule {
	s.DestinationRegionId = &v
	return s
}

func (s *CreateBackupPlanRequestRule) SetDestinationRetention(v int64) *CreateBackupPlanRequestRule {
	s.DestinationRetention = &v
	return s
}

func (s *CreateBackupPlanRequestRule) SetDisabled(v bool) *CreateBackupPlanRequestRule {
	s.Disabled = &v
	return s
}

func (s *CreateBackupPlanRequestRule) SetDoCopy(v bool) *CreateBackupPlanRequestRule {
	s.DoCopy = &v
	return s
}

func (s *CreateBackupPlanRequestRule) SetRetention(v int64) *CreateBackupPlanRequestRule {
	s.Retention = &v
	return s
}

func (s *CreateBackupPlanRequestRule) SetRuleName(v string) *CreateBackupPlanRequestRule {
	s.RuleName = &v
	return s
}

func (s *CreateBackupPlanRequestRule) SetSchedule(v string) *CreateBackupPlanRequestRule {
	s.Schedule = &v
	return s
}

func (s *CreateBackupPlanRequestRule) Validate() error {
	return dara.Validate(s)
}
