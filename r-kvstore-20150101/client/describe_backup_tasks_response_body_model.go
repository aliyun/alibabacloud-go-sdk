// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeBackupTasksResponseBodyAccessDeniedDetail) *DescribeBackupTasksResponseBody
	GetAccessDeniedDetail() *DescribeBackupTasksResponseBodyAccessDeniedDetail
	SetBackupJobs(v []*DescribeBackupTasksResponseBodyBackupJobs) *DescribeBackupTasksResponseBody
	GetBackupJobs() []*DescribeBackupTasksResponseBodyBackupJobs
	SetInstanceId(v string) *DescribeBackupTasksResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *DescribeBackupTasksResponseBody
	GetRequestId() *string
}

type DescribeBackupTasksResponseBody struct {
	// The following parameters are no longer used. Ignore the parameters.
	AccessDeniedDetail *DescribeBackupTasksResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The details of the backup tasks.
	BackupJobs []*DescribeBackupTasksResponseBodyBackupJobs `json:"BackupJobs,omitempty" xml:"BackupJobs,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BB73740C-23E2-4392-9DA4-2660C74C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponseBody) GetAccessDeniedDetail() *DescribeBackupTasksResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeBackupTasksResponseBody) GetBackupJobs() []*DescribeBackupTasksResponseBodyBackupJobs {
	return s.BackupJobs
}

func (s *DescribeBackupTasksResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBackupTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupTasksResponseBody) SetAccessDeniedDetail(v *DescribeBackupTasksResponseBodyAccessDeniedDetail) *DescribeBackupTasksResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeBackupTasksResponseBody) SetBackupJobs(v []*DescribeBackupTasksResponseBodyBackupJobs) *DescribeBackupTasksResponseBody {
	s.BackupJobs = v
	return s
}

func (s *DescribeBackupTasksResponseBody) SetInstanceId(v string) *DescribeBackupTasksResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackupTasksResponseBody) SetRequestId(v string) *DescribeBackupTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupTasksResponseBodyAccessDeniedDetail struct {
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeBackupTasksResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTasksResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeBackupTasksResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeBackupTasksResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeBackupTasksResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeBackupTasksResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeBackupTasksResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeBackupTasksResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeBackupTasksResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeBackupTasksResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeBackupTasksResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeBackupTasksResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeBackupTasksResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeBackupTasksResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeBackupTasksResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeBackupTasksResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupTasksResponseBodyBackupJobs struct {
	// The ID of the backup task.
	//
	// example:
	//
	// 8491111
	BackupJobID *int64 `json:"BackupJobID,omitempty" xml:"BackupJobID,omitempty"`
	// The state of the backup task. Valid values:
	//
	// 	- **NoStart**: The backup task is not started.
	//
	// 	- **Preparing**: The backup task is being prepared.
	//
	// 	- **Waiting**: The backup task is pending.
	//
	// 	- **Uploading**: The system is uploading the backup file.
	//
	// 	- **Checking**: The system is checking the uploaded backup file.
	//
	// 	- **Finished**: The backup task is completed.
	//
	// example:
	//
	// Automated
	BackupProgressStatus *string `json:"BackupProgressStatus,omitempty" xml:"BackupProgressStatus,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- **Automated**: automatic backup
	//
	// 	- **Manual**: manual backup
	//
	// example:
	//
	// Manual
	JobMode *string `json:"JobMode,omitempty" xml:"JobMode,omitempty"`
	// The ID of the data node.
	//
	// example:
	//
	// ****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The progress of the backup task in percentage.
	//
	// example:
	//
	// 0
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// The backup progress.
	//
	// example:
	//
	// 27
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The start time of the backup task. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-01-05T19:24:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the backup task. Valid values:
	//
	// 	- **TempBackupTask**: The backup task was manually performed.
	//
	// 	- **NormalBackupTask**: The backup task was automatically performed.
	//
	// example:
	//
	// NormalBackupTask
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
}

func (s DescribeBackupTasksResponseBodyBackupJobs) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTasksResponseBodyBackupJobs) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) GetBackupJobID() *int64 {
	return s.BackupJobID
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) GetBackupProgressStatus() *string {
	return s.BackupProgressStatus
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) GetJobMode() *string {
	return s.JobMode
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) GetProcess() *string {
	return s.Process
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) GetProgress() *string {
	return s.Progress
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) GetTaskAction() *string {
	return s.TaskAction
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetBackupJobID(v int64) *DescribeBackupTasksResponseBodyBackupJobs {
	s.BackupJobID = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetBackupProgressStatus(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.BackupProgressStatus = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetJobMode(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.JobMode = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetNodeId(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.NodeId = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetProcess(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.Process = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetProgress(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.Progress = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetStartTime(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) SetTaskAction(v string) *DescribeBackupTasksResponseBodyBackupJobs {
	s.TaskAction = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyBackupJobs) Validate() error {
	return dara.Validate(s)
}
