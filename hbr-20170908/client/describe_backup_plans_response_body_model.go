// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlans(v *DescribeBackupPlansResponseBodyBackupPlans) *DescribeBackupPlansResponseBody
	GetBackupPlans() *DescribeBackupPlansResponseBodyBackupPlans
	SetCode(v string) *DescribeBackupPlansResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeBackupPlansResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeBackupPlansResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupPlansResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBackupPlansResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackupPlansResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeBackupPlansResponseBody
	GetTotalCount() *int64
}

type DescribeBackupPlansResponseBody struct {
	// The queried backup plans.
	BackupPlans *DescribeBackupPlansResponseBodyBackupPlans `json:"BackupPlans,omitempty" xml:"BackupPlans,omitempty" type:"Struct"`
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned backup plans that meet the specified conditions.
	//
	// example:
	//
	// 8
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupPlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBody) GetBackupPlans() *DescribeBackupPlansResponseBodyBackupPlans {
	return s.BackupPlans
}

func (s *DescribeBackupPlansResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeBackupPlansResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeBackupPlansResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupPlansResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupPlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPlansResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupPlansResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeBackupPlansResponseBody) SetBackupPlans(v *DescribeBackupPlansResponseBodyBackupPlans) *DescribeBackupPlansResponseBody {
	s.BackupPlans = v
	return s
}

func (s *DescribeBackupPlansResponseBody) SetCode(v string) *DescribeBackupPlansResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBackupPlansResponseBody) SetMessage(v string) *DescribeBackupPlansResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBackupPlansResponseBody) SetPageNumber(v int32) *DescribeBackupPlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupPlansResponseBody) SetPageSize(v int32) *DescribeBackupPlansResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupPlansResponseBody) SetRequestId(v string) *DescribeBackupPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPlansResponseBody) SetSuccess(v bool) *DescribeBackupPlansResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupPlansResponseBody) SetTotalCount(v int64) *DescribeBackupPlansResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBackupPlansResponseBody) Validate() error {
	if s.BackupPlans != nil {
		if err := s.BackupPlans.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupPlansResponseBodyBackupPlans struct {
	BackupPlan []*DescribeBackupPlansResponseBodyBackupPlansBackupPlan `json:"BackupPlan,omitempty" xml:"BackupPlan,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlansResponseBodyBackupPlans) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlans) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlans) GetBackupPlan() []*DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	return s.BackupPlan
}

func (s *DescribeBackupPlansResponseBodyBackupPlans) SetBackupPlan(v []*DescribeBackupPlansResponseBodyBackupPlansBackupPlan) *DescribeBackupPlansResponseBodyBackupPlans {
	s.BackupPlan = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlans) Validate() error {
	if s.BackupPlan != nil {
		for _, item := range s.BackupPlan {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlan struct {
	// The ID of the data source group.
	//
	// example:
	//
	// System-Database
	BackupSourceGroupId *string `json:"BackupSourceGroupId,omitempty" xml:"BackupSourceGroupId,omitempty"`
	// The backup type. Valid value: **COMPLETE**, which indicates full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// This parameter is valid only when **SourceType*	- is set to **OSS**. This parameter indicates the name of the OSS bucket.
	//
	// example:
	//
	// hbr-backup-oss
	Bucket         *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The configurations of the incremental file synchronization. This parameter is returned only for data synchronization.
	//
	// example:
	//
	// {"dataSourceId": "ds-123456789", "path": "/changelist"}
	ChangeListPath *string `json:"ChangeListPath,omitempty" xml:"ChangeListPath,omitempty"`
	// The ID of the backup client.
	//
	// example:
	//
	// c-000ge4w*****1qb
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the client group.
	//
	// example:
	//
	// cl-000ht6o9******h
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is valid only when **SourceType*	- is set to **NAS**. This parameter indicates the time when the file system was created. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether a backup plan is automatically created based on tags.
	//
	// example:
	//
	// false
	CreatedByTag *bool `json:"CreatedByTag,omitempty" xml:"CreatedByTag,omitempty"`
	// The time when the backup plan was created. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The name of the Resource Access Management (RAM) role that is created within the source Alibaba Cloud account and assigned to the current Alibaba Cloud account to authorize the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// BackupRole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// Indicates whether data is backed up within the same Alibaba Cloud account or across Alibaba Cloud accounts. Valid values:
	//
	// 	- SELF_ACCOUNT
	//
	// 	- CROSS_ACCOUNT
	//
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the source Alibaba Cloud account that authorizes the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// 1841642xxxxx9795
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The ID of the data source.
	//
	// example:
	//
	// ds-000ht6o9*****w61
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The data source details at the destination. This parameter is returned only for data synchronization.
	//
	// example:
	//
	// {\\"prefix\\":\\"/\\"}
	DestDataSourceDetail *string `json:"DestDataSourceDetail,omitempty" xml:"DestDataSourceDetail,omitempty"`
	// The data source ID at the destination. This parameter is returned only for data synchronization.
	//
	// example:
	//
	// ds-*********************
	DestDataSourceId *string `json:"DestDataSourceId,omitempty" xml:"DestDataSourceId,omitempty"`
	// The data source type at the destination. This parameter is returned only for data synchronization.
	//
	// example:
	//
	// OSS
	DestSourceType *string `json:"DestSourceType,omitempty" xml:"DestSourceType,omitempty"`
	// The details about ECS instance backup.
	//
	// example:
	//
	// {\\"doCopy\\":true,\\"doBackup\\":false,\\"instanceName\\":\\"instance example\\",\\"appConsistent\\":false,\\"destinationRegionId\\":\\"cn-shanghai\\",\\"enableFsFreeze\\":true,\\"osNameEn\\":\\"Windows Server  2019 Data Center Edition 64bit Chinese Edition\\",\\"osName\\":\\"Windows Server  2019 Data Center Edition 64bit Chinese Edition\\",\\"diskIdList\\":[],\\"backupVaultId\\":\\"\\",\\"snapshotGroup\\":true,\\"destinationRetention\\":35,\\"platform\\":\\"Windows Server 2012\\",\\"timeoutInSeconds\\":60,\\"backupRetention\\":1,\\"osType\\":\\"windows\\",\\"preScriptPath\\":\\"\\",\\"postScriptPath\\":\\"\\",\\"enableWriters\\":true,\\"ecsDeleted\\":false}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// Indicates whether the backup plan is disabled. Valid values:
	//
	// 	- true: The backup plan is disabled.
	//
	// 	- false: The backup plan is enabled.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// This parameter is valid only when **SourceType*	- is set to **ECS_FILE**. This parameter indicates the paths to the files that are excluded from the backup job.
	//
	// example:
	//
	// ["/var", "/proc"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// This parameter is valid only when **SourceType*	- is set to **NAS**. This parameter indicates the ID of the NAS file system.
	//
	// example:
	//
	// 00594
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The matched tag rules.
	HitTags *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTags `json:"HitTags,omitempty" xml:"HitTags,omitempty" type:"Struct"`
	// This parameter is valid only when **SourceType*	- is set to **ECS_FILE**. This parameter indicates the paths to the files that are backed up.
	//
	// example:
	//
	// ["/home/alice/*.pdf", "/home/bob/*.txt"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// The ID of the instance group.
	//
	// example:
	//
	// i-**
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// This parameter is valid only when **SourceType*	- is set to **ECS_FILE**. This parameter indicates the ID of the ECS instance.
	//
	// example:
	//
	// i-*********************
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Tablestore instance.
	//
	// example:
	//
	// instancename
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Indicates whether the feature of keeping at least one backup version is enabled. Valid values:
	//
	// 	- 0: The feature is disabled.
	//
	// 	- 1: The feature is enabled.
	//
	// example:
	//
	// 0
	KeepLatestSnapshots *int64 `json:"KeepLatestSnapshots,omitempty" xml:"KeepLatestSnapshots,omitempty"`
	// The latest execution job id of plan.
	//
	// example:
	//
	// job-12345678
	LatestExecuteJobId *string `json:"LatestExecuteJobId,omitempty" xml:"LatestExecuteJobId,omitempty"`
	// example:
	//
	// job-00**************9khz
	LatestFinishJobId *string `json:"LatestFinishJobId,omitempty" xml:"LatestFinishJobId,omitempty"`
	// This parameter is valid only when **SourceType*	- is set to **ECS_FILE**. This parameter indicates whether Windows Volume Shadow Copy Service (VSS) is used to define a source path.
	//
	// example:
	//
	// {"UseVSS":false}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The details about the Tablestore instance.
	OtsDetail *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty" type:"Struct"`
	// The source paths. This parameter is valid only when **SourceType*	- is set to **ECS_FILE**.
	Paths *DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Struct"`
	// The ID of the backup plan.
	//
	// example:
	//
	// plan-*********************
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The name of the backup plan.
	//
	// example:
	//
	// planname
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// This parameter is valid only when **SourceType*	- is set to **OSS**. This parameter indicates the prefix of the objects that are backed up.
	//
	// example:
	//
	// oss-prefix
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The backup resources. This parameter is valid only for disk backup.
	Resources *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// The retention period of the backup data. Unit: days.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The backup policies. This parameter is valid only for disk backup.
	Rules *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	// The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified by `{startTime}` and the subsequent backup jobs at an interval that is specified by `{interval}`. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, `I|1631685600|P1D` indicates that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
	//
	// 	- **startTime**: the time at which the system starts to run a backup job. The time follows the UNIX time format. Unit: seconds.
	//
	// 	- **interval**: the interval at which the system runs a backup job. The interval follows the ISO 8601 standard. For example, PT1H indicates an interval of 1 hour. P1D indicates an interval of one day.
	//
	// example:
	//
	// I|1602673264|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **ECS_FILE**: ECS files
	//
	// 	- **OSS**: OSS buckets
	//
	// 	- **NAS**: NAS file systems
	//
	// 	- **OTS**: Tablestore instances
	//
	// 	- **UDM_ECS**: ECS instances
	//
	// 	- **SYNC**: data synchronization
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is valid only when **SourceType*	- is set to **ECS_FILE**. This parameter indicates the throttling rules. Format: `{start}|{end}|{bandwidth}`. Multiple throttling rules are separated with vertical bars (`|`). A time range cannot overlap with another one.
	//
	// 	- start: the start hour.
	//
	// 	- end: the end hour.
	//
	// 	- bandwidth: the bandwidth. Unit: KB.
	//
	// example:
	//
	// 0:24:5120
	SpeedLimit *string `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	// The free trial information.
	TrialInfo *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo `json:"TrialInfo,omitempty" xml:"TrialInfo,omitempty" type:"Struct"`
	// The time when the backup plan was updated. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-*********************
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlan) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetBackupSourceGroupId() *string {
	return s.BackupSourceGroupId
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetBucket() *string {
	return s.Bucket
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetChangeListPath() *string {
	return s.ChangeListPath
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetCreatedByTag() *bool {
	return s.CreatedByTag
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetDestDataSourceDetail() *string {
	return s.DestDataSourceDetail
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetDestDataSourceId() *string {
	return s.DestDataSourceId
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetDestSourceType() *string {
	return s.DestSourceType
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetDetail() *string {
	return s.Detail
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetExclude() *string {
	return s.Exclude
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetHitTags() *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTags {
	return s.HitTags
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetInclude() *string {
	return s.Include
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetInstanceGroupId() *string {
	return s.InstanceGroupId
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetKeepLatestSnapshots() *int64 {
	return s.KeepLatestSnapshots
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetLatestExecuteJobId() *string {
	return s.LatestExecuteJobId
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetLatestFinishJobId() *string {
	return s.LatestFinishJobId
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetOptions() *string {
	return s.Options
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetOtsDetail() *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail {
	return s.OtsDetail
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetPaths() *DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths {
	return s.Paths
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetPlanId() *string {
	return s.PlanId
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetPlanName() *string {
	return s.PlanName
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetPrefix() *string {
	return s.Prefix
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetResources() *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources {
	return s.Resources
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetRetention() *int64 {
	return s.Retention
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetRules() *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules {
	return s.Rules
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetSchedule() *string {
	return s.Schedule
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetSpeedLimit() *string {
	return s.SpeedLimit
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetTrialInfo() *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo {
	return s.TrialInfo
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetBackupSourceGroupId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.BackupSourceGroupId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetBackupType(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetBucket(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Bucket = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetBusinessStatus(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetChangeListPath(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.ChangeListPath = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetClientId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.ClientId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetClusterId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetCreateTime(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.CreateTime = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetCreatedByTag(v bool) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.CreatedByTag = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetCreatedTime(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.CreatedTime = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetCrossAccountRoleName(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetCrossAccountType(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.CrossAccountType = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetCrossAccountUserId(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetDataSourceId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.DataSourceId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetDestDataSourceDetail(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.DestDataSourceDetail = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetDestDataSourceId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.DestDataSourceId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetDestSourceType(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.DestSourceType = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetDetail(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Detail = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetDisabled(v bool) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Disabled = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetExclude(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Exclude = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetFileSystemId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.FileSystemId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetHitTags(v *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTags) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.HitTags = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetInclude(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Include = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetInstanceGroupId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.InstanceGroupId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetInstanceId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetInstanceName(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.InstanceName = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetKeepLatestSnapshots(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.KeepLatestSnapshots = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetLatestExecuteJobId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.LatestExecuteJobId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetLatestFinishJobId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.LatestFinishJobId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetOptions(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Options = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetOtsDetail(v *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.OtsDetail = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetPaths(v *DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Paths = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetPlanId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.PlanId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetPlanName(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.PlanName = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetPrefix(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Prefix = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetResources(v *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Resources = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetRetention(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Retention = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetRules(v *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Rules = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetSchedule(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.Schedule = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetSourceType(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.SourceType = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetSpeedLimit(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.SpeedLimit = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetTrialInfo(v *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.TrialInfo = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetUpdatedTime(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) SetVaultId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlan {
	s.VaultId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlan) Validate() error {
	if s.HitTags != nil {
		if err := s.HitTags.Validate(); err != nil {
			return err
		}
	}
	if s.OtsDetail != nil {
		if err := s.OtsDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Paths != nil {
		if err := s.Paths.Validate(); err != nil {
			return err
		}
	}
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		if err := s.Rules.Validate(); err != nil {
			return err
		}
	}
	if s.TrialInfo != nil {
		if err := s.TrialInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTags struct {
	HitTag []*DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTagsHitTag `json:"HitTag,omitempty" xml:"HitTag,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTags) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTags) GetHitTag() []*DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTagsHitTag {
	return s.HitTag
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTags) SetHitTag(v []*DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTagsHitTag) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTags {
	s.HitTag = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTags) Validate() error {
	if s.HitTag != nil {
		for _, item := range s.HitTag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTagsHitTag struct {
	// The tag key.
	//
	// example:
	//
	// type
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag-based matching rule. Valid values:
	//
	// 	- **EQUAL**: Both the tag key and tag value are matched.
	//
	// 	- **NOT**: The tag key is matched and the tag value is not matched.
	//
	// example:
	//
	// EQUAL
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTagsHitTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTagsHitTag) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTagsHitTag) GetKey() *string {
	return s.Key
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTagsHitTag) GetOperator() *string {
	return s.Operator
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTagsHitTag) GetValue() *string {
	return s.Value
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTagsHitTag) SetKey(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTagsHitTag {
	s.Key = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTagsHitTag) SetOperator(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTagsHitTag {
	s.Operator = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTagsHitTag) SetValue(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTagsHitTag {
	s.Value = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanHitTagsHitTag) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail struct {
	// The names of the tables in the Tablestore instance.
	TableNames *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Struct"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail) GetTableNames() *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames {
	return s.TableNames
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail) SetTableNames(v *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail {
	s.TableNames = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetail) Validate() error {
	if s.TableNames != nil {
		if err := s.TableNames.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames struct {
	TableName []*string `json:"TableName,omitempty" xml:"TableName,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames) GetTableName() []*string {
	return s.TableName
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames) SetTableName(v []*string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames {
	s.TableName = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanOtsDetailTableNames) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths struct {
	Path []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths) GetPath() []*string {
	return s.Path
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths) SetPath(v []*string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths {
	s.Path = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanPaths) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources struct {
	Resource []*DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources) GetResource() []*DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource {
	return s.Resource
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources) SetResource(v []*DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources {
	s.Resource = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResources) Validate() error {
	if s.Resource != nil {
		for _, item := range s.Resource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource struct {
	// Additional information about the data source.
	//
	// example:
	//
	// {\\"doBackup\\":false,\\"diskName\\":\\"data_disk\\",\\"size\\":100,\\"type\\":\\"data\\",\\"category\\":\\"cloud_essd\\",\\"imageId\\":\\"\\",\\"device\\":\\"/dev/xvdb\\",\\"encrypted\\":false}
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The ID of the data source.
	//
	// example:
	//
	// d-j6cgioir6m******lu4
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the data source. Valid value: **UDM_DISK**.
	//
	// example:
	//
	// UDMDISK
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource) GetExtra() *string {
	return s.Extra
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource) SetExtra(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource {
	s.Extra = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource) SetResourceId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource {
	s.ResourceId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource) SetSourceType(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource {
	s.SourceType = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanResourcesResource) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules struct {
	Rule []*DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules) GetRule() []*DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	return s.Rule
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules) SetRule(v []*DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules {
	s.Rule = v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRules) Validate() error {
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

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule struct {
	// The backup type. Valid value: **COMPLETE**, which indicates full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The ID of the region in which the remote backup vault resides.
	//
	// example:
	//
	// cn-shanghai
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// The retention period of the backup data in remote backup mode. Unit: days.
	//
	// example:
	//
	// 90
	DestinationRetention *int64 `json:"DestinationRetention,omitempty" xml:"DestinationRetention,omitempty"`
	// Indicates whether the policy is disabled.
	//
	// example:
	//
	// true
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// Indicates whether the snapshot data is backed up to the backup vault.
	//
	// example:
	//
	// false
	DoCopy *bool `json:"DoCopy,omitempty" xml:"DoCopy,omitempty"`
	// The retention period of the backup data. Unit: days.
	//
	// example:
	//
	// 90
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// rule-0008i52rf0ulpni6kn6m
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The policy name.
	//
	// example:
	//
	// Disk Golden Rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified by `{startTime}` and the subsequent backup jobs at an interval that is specified by `{interval}`. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, `I|1631685600|P1D` indicates that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
	//
	// 	- `startTime`: the time at which the system starts to run a backup job. The time follows the UNIX time format. Unit: seconds.
	//
	// 	- `interval`: the interval at which the system runs a backup job. The interval follows the ISO 8601 standard. For example, PT1H indicates an interval of 1 hour. P1D indicates an interval of one day.
	//
	// example:
	//
	// I|1631685600|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) GetDestinationRegionId() *string {
	return s.DestinationRegionId
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) GetDestinationRetention() *int64 {
	return s.DestinationRetention
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) GetDoCopy() *bool {
	return s.DoCopy
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) GetRetention() *int64 {
	return s.Retention
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) GetSchedule() *string {
	return s.Schedule
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetBackupType(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetDestinationRegionId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.DestinationRegionId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetDestinationRetention(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.DestinationRetention = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetDisabled(v bool) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.Disabled = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetDoCopy(v bool) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.DoCopy = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetRetention(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.Retention = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetRuleId(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetRuleName(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.RuleName = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) SetSchedule(v string) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule {
	s.Schedule = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanRulesRule) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo struct {
	// Indicates whether you are billed based on the pay-as-you-go method after the free trial ends.
	//
	// example:
	//
	// true
	KeepAfterTrialExpiration *bool `json:"KeepAfterTrialExpiration,omitempty" xml:"KeepAfterTrialExpiration,omitempty"`
	// The expiration time of the free trial.
	//
	// example:
	//
	// 1584597600
	TrialExpireTime *int64 `json:"TrialExpireTime,omitempty" xml:"TrialExpireTime,omitempty"`
	// The start time of the free trial.
	//
	// example:
	//
	// 1579413159
	TrialStartTime *int64 `json:"TrialStartTime,omitempty" xml:"TrialStartTime,omitempty"`
	// The time when the free-trial backup vault is released.
	//
	// example:
	//
	// 1594965600
	TrialVaultReleaseTime *int64 `json:"TrialVaultReleaseTime,omitempty" xml:"TrialVaultReleaseTime,omitempty"`
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) GetKeepAfterTrialExpiration() *bool {
	return s.KeepAfterTrialExpiration
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) GetTrialExpireTime() *int64 {
	return s.TrialExpireTime
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) GetTrialStartTime() *int64 {
	return s.TrialStartTime
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) GetTrialVaultReleaseTime() *int64 {
	return s.TrialVaultReleaseTime
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) SetKeepAfterTrialExpiration(v bool) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo {
	s.KeepAfterTrialExpiration = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) SetTrialExpireTime(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo {
	s.TrialExpireTime = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) SetTrialStartTime(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo {
	s.TrialStartTime = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) SetTrialVaultReleaseTime(v int64) *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo {
	s.TrialVaultReleaseTime = &v
	return s
}

func (s *DescribeBackupPlansResponseBodyBackupPlansBackupPlanTrialInfo) Validate() error {
	return dara.Validate(s)
}
