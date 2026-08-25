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
	SetEdition(v string) *CreateBackupPlanRequest
	GetEdition() *string
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
	DestDataSourceDetail map[string]interface{} `json:"DestDataSourceDetail,omitempty" xml:"DestDataSourceDetail,omitempty"`
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
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
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
	OtsDetail *OtsDetail `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty"`
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
	Rule []*CreateBackupPlanRequestRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
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

func (s *CreateBackupPlanRequest) GetEdition() *string {
	return s.Edition
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

func (s *CreateBackupPlanRequest) SetEdition(v string) *CreateBackupPlanRequest {
	s.Edition = &v
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
	if s.OtsDetail != nil {
		if err := s.OtsDetail.Validate(); err != nil {
			return err
		}
	}
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

type CreateBackupPlanRequestRule struct {
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
