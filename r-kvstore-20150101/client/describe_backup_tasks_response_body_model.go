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
	// **[Deprecated]*	- This group of parameters is deprecated. Ignore these parameters.
	AccessDeniedDetail *DescribeBackupTasksResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The details of the backup tasks.
	BackupJobs []*DescribeBackupTasksResponseBodyBackupJobs `json:"BackupJobs,omitempty" xml:"BackupJobs,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request.
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
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.BackupJobs != nil {
		for _, item := range s.BackupJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupTasksResponseBodyAccessDeniedDetail struct {
	// Same as above.
	//
	// example:
	//
	// -
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// Same as above.
	//
	// example:
	//
	// -
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// Same as above.
	//
	// example:
	//
	// -
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// Same as above.
	//
	// example:
	//
	// -
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// Same as above.
	//
	// example:
	//
	// -
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// Same as above.
	//
	// example:
	//
	// -
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// Same as above.
	//
	// example:
	//
	// -
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
	// 29****10
	BackupJobID *int64 `json:"BackupJobID,omitempty" xml:"BackupJobID,omitempty"`
	// The status of the backup task. Valid values:
	//
	// 	- **NoStart**: not started
	//
	// 	- **Preparing**: preparing
	//
	// 	- **Waiting**: waiting
	//
	// 	- **Uploading**: uploading
	//
	// 	- **Checking**: checking
	//
	// 	- **Finished**: completed
	//
	// example:
	//
	// NoStart
	BackupProgressStatus *string `json:"BackupProgressStatus,omitempty" xml:"BackupProgressStatus,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- **Automated**: automatic backup.
	//
	// 	- **Manual**: manual backup.
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
	// The progress percentage of the backup task. This parameter is not recommended. Refer to the **Progress*	- parameter.
	//
	// example:
	//
	// 0
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// The progress percentage of the backup. Valid values: 0 to 100.
	//
	// example:
	//
	// 27
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The start time of the backup task. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-01-05T19:24:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the backup task. Valid values:
	//
	// 	- **TempBackupTask**: temporary backup task, which indicates a manual backup.
	//
	// 	- **NormalBackupTask**: normal backup task, which indicates an automatic backup.
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
