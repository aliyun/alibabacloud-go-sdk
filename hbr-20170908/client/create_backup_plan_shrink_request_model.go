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
	SetEdition(v string) *CreateBackupPlanShrinkRequest
	GetEdition() *string
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
	// The backup type. Set the value to **COMPLETE**, which indicates full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **OSS**. The name of the OSS bucket.
	//
	// example:
	//
	// hbr-backup-oss
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The configuration of the incremental file synchronization list. This parameter is required only for data synchronization.
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
	// This parameter is required only when **SourceType*	- is set to **NAS**. The time when the file system was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1607436917
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the RAM role created in the source account for cross-account backup.
	//
	// example:
	//
	// BackupRole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// The cross-account backup type. Valid values:
	//
	// - SELF_ACCOUNT: backup within the same account.
	//
	// - CROSS_ACCOUNT: cross-account backup.
	//
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the source account for cross-account backup.
	//
	// example:
	//
	// 15897534xxxx4625
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The ID of the source data source. This parameter is required only for data synchronization.
	//
	// example:
	//
	// ds-****************
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The details of the destination data source. This parameter is required only for data synchronization.
	//
	// example:
	//
	// {\\"prefix\\":\\"/\\"}
	DestDataSourceDetailShrink *string `json:"DestDataSourceDetail,omitempty" xml:"DestDataSourceDetail,omitempty"`
	// The ID of the destination data source. This parameter is required only for data synchronization.
	//
	// example:
	//
	// ds-*********************
	DestDataSourceId *string `json:"DestDataSourceId,omitempty" xml:"DestDataSourceId,omitempty"`
	// The type of the destination data source. This parameter is required only for data synchronization.
	//
	// example:
	//
	// OSS
	DestSourceType *string `json:"DestSourceType,omitempty" xml:"DestSourceType,omitempty"`
	// The details of the full-copy backup. The value is a JSON string.
	//
	// 	- snapshotGroup: specifies whether to use a consistent snapshot group. This parameter is valid only when all cloud disks of the instance are ESSDs.
	//
	// 	- appConsistent: specifies whether to use application consistency. This parameter must be used together with the preScriptPath and postScriptPath parameters.
	//
	// 	- preScriptPath: the path of the pre-freeze script.
	//
	// 	- postScriptPath: the path of the post-thaw script.
	//
	// example:
	//
	// {\\"EnableFsFreeze\\":true,\\"appConsistent\\":false,\\"postScriptPath\\":\\"\\",\\"preScriptPath\\":\\"\\",\\"snapshotGroup\\":true,\\"timeoutInSeconds\\":60}
	DetailShrink *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// Specifies whether the plan is disabled by default.
	//
	// example:
	//
	// true
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The edition type. Valid values: BASIC and STANDARD. Default value: STANDARD.
	//
	// example:
	//
	// STANDARD
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **ECS_FILE**. The path to exclude from the backup. All files in this path are not backed up. The value can be up to 255 characters in length.
	//
	// example:
	//
	// ["/var", "/proc"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **NAS**. The file system ID.
	//
	// example:
	//
	// 005494
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **ECS_FILE**. The path to include in the backup. All files in this path are backed up. The value can be up to 255 characters in length.
	//
	// example:
	//
	// ["/home/alice/*.pdf", "/home/bob/*.txt"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **ECS_FILE**. The ECS instance ID.
	//
	// example:
	//
	// i-m5e*****6q
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Tablestore instance.
	//
	// example:
	//
	// instancename
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Specifies whether to retain at least one backup version. Valid values:
	//
	// - 0: does not retain.
	//
	// - 1: retains.
	//
	// example:
	//
	// 1
	KeepLatestSnapshots *int64 `json:"KeepLatestSnapshots,omitempty" xml:"KeepLatestSnapshots,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **ECS_FILE**. Specifies whether to use Windows Volume Shadow Copy Service (VSS) to define the source path.
	//
	// - This feature is supported only for Windows ECS instances.
	//
	// - If the backup source contains data changes and you need to ensure consistency between the backup data and the source data, set this parameter to `["UseVSS":true]`.
	//
	// - After VSS is enabled, multiple file folders cannot be backed up simultaneously.
	//
	// example:
	//
	// {"UseVSS":false}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The details of the Tablestore instance.
	OtsDetailShrink *string `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty"`
	// The source paths.
	Path []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
	// The name of the backup plan. The name must be 1 to 64 characters in length. The backup plan name must be unique for each data source type within a single vault.
	//
	// example:
	//
	// planname
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **OSS**. The backup prefix. If specified, only objects that match the prefix are backed up.
	//
	// example:
	//
	// oss-prefix
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The retention period of the backup data. Minimum value: 1. Unit: days.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The backup plan rules.
	Rule []*CreateBackupPlanShrinkRequestRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
	// The backup policy. Format: `I|{startTime}|{interval}`. This indicates that a backup job is executed at every `{interval}` starting from `{startTime}`. Backup jobs for past time periods are not compensated. If the previous backup job is not completed, the next backup job is not triggered. Example: `I|1631685600|P1D` indicates that a backup is performed once a day starting from 2021-09-15 14:00:00.
	//
	// - **startTime**: the start time of the backup. The value is a UNIX timestamp. Unit: seconds.
	//
	// - **interval**: the ISO 8601 time interval. Example: PT1H indicates an interval of one hour. P1D indicates an interval of one day.
	//
	// example:
	//
	// I|1602673264|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The type of the data source. Valid values:
	//
	// - **ECS_FILE**: backs up ECS files.
	//
	// - **OSS**: backs up Alibaba Cloud OSS.
	//
	// - **NAS**: backs up Alibaba Cloud NAS.
	//
	// - **OTS**: backs up Alibaba Cloud OTS.
	//
	// - **UDM_ECS**: backs up an entire ECS instance.
	//
	// - **SYNC**: data synchronization.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **ECS_FILE**. The backup traffic control. Format: `{start}:{end}:{bandwidth}`. Separate multiple traffic control configurations with vertical bars (|). The time ranges of the configurations cannot overlap.
	//
	// - **start**: the start hour.
	//
	// - **end**: the end hour.
	//
	// - **bandwidth**: the rate limit. Unit: KB/s.
	//
	// example:
	//
	// 0:24:5120
	SpeedLimit *string `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	// The region where the ECS instance for full-copy backup resides.
	//
	// example:
	//
	// cn-shanghai
	UdmRegionId *string `json:"UdmRegionId,omitempty" xml:"UdmRegionId,omitempty"`
	// The vault ID.
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

func (s *CreateBackupPlanShrinkRequest) GetEdition() *string {
	return s.Edition
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

func (s *CreateBackupPlanShrinkRequest) SetEdition(v string) *CreateBackupPlanShrinkRequest {
	s.Edition = &v
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
	// The backup type.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The ID of the destination region for cross-region replication.
	//
	// example:
	//
	// cn-hangzhou
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// The retention period of the geo-redundancy backup. Unit: days.
	//
	// example:
	//
	// 7
	DestinationRetention *int64 `json:"DestinationRetention,omitempty" xml:"DestinationRetention,omitempty"`
	// Specifies whether the rule is disabled.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// Specifies whether to enable cross-region replication.
	//
	// example:
	//
	// false
	DoCopy *bool `json:"DoCopy,omitempty" xml:"DoCopy,omitempty"`
	// The retention period of the backup.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule-test-name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The backup policy. Format: I|{startTime}|{interval}. This indicates that a backup job is executed at every {interval} starting from {startTime}. Backup jobs for past time periods are not executed. If the previous backup job is not completed, the next backup job is not triggered. Example: I|1631685600|P1D indicates that a backup is performed once a day starting from 2021-09-15 14:00:00.
	//
	// startTime: the start time of the backup. The value is a UNIX timestamp. Unit: seconds.
	//
	// interval: the ISO 8601 time interval. Example: PT1H indicates an interval of one hour. P1D indicates an interval of one day.
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
