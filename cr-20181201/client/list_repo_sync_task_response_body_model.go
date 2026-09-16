// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoSyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRepoSyncTaskResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListRepoSyncTaskResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListRepoSyncTaskResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListRepoSyncTaskResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRepoSyncTaskResponseBody
	GetRequestId() *string
	SetSyncTasks(v []*ListRepoSyncTaskResponseBodySyncTasks) *ListRepoSyncTaskResponseBody
	GetSyncTasks() []*ListRepoSyncTaskResponseBodySyncTasks
	SetTotalCount(v string) *ListRepoSyncTaskResponseBody
	GetTotalCount() *string
}

type ListRepoSyncTaskResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The page size.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7640819A-FB5B-4E25-A227-97717F62****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of synchronization tasks.
	SyncTasks []*ListRepoSyncTaskResponseBodySyncTasks `json:"SyncTasks,omitempty" xml:"SyncTasks,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRepoSyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRepoSyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepoSyncTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRepoSyncTaskResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListRepoSyncTaskResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListRepoSyncTaskResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRepoSyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRepoSyncTaskResponseBody) GetSyncTasks() []*ListRepoSyncTaskResponseBodySyncTasks {
	return s.SyncTasks
}

func (s *ListRepoSyncTaskResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListRepoSyncTaskResponseBody) SetCode(v string) *ListRepoSyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepoSyncTaskResponseBody) SetIsSuccess(v bool) *ListRepoSyncTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepoSyncTaskResponseBody) SetPageNo(v int32) *ListRepoSyncTaskResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRepoSyncTaskResponseBody) SetPageSize(v int32) *ListRepoSyncTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepoSyncTaskResponseBody) SetRequestId(v string) *ListRepoSyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBody) SetSyncTasks(v []*ListRepoSyncTaskResponseBodySyncTasks) *ListRepoSyncTaskResponseBody {
	s.SyncTasks = v
	return s
}

func (s *ListRepoSyncTaskResponseBody) SetTotalCount(v string) *ListRepoSyncTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRepoSyncTaskResponseBody) Validate() error {
	if s.SyncTasks != nil {
		for _, item := range s.SyncTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRepoSyncTaskResponseBodySyncTasks struct {
	// The creation time.
	//
	// example:
	//
	// 1572839126000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the image is synchronized across accounts. Valid values:
	//
	// - `true`: The image is synchronized across accounts.
	//
	// - `false`: The image is synchronized within the same account.
	//
	// Default value: `false`
	//
	// example:
	//
	// true
	CrossUser *bool `json:"CrossUser,omitempty" xml:"CrossUser,omitempty"`
	// Indicates whether a custom synchronization link is used.
	//
	// example:
	//
	// true
	CustomLink *bool `json:"CustomLink,omitempty" xml:"CustomLink,omitempty"`
	// The source image.
	ImageFrom *ListRepoSyncTaskResponseBodySyncTasksImageFrom `json:"ImageFrom,omitempty" xml:"ImageFrom,omitempty" type:"Struct"`
	// The destination image.
	ImageTo *ListRepoSyncTaskResponseBodySyncTasksImageTo `json:"ImageTo,omitempty" xml:"ImageTo,omitempty" type:"Struct"`
	// The custom synchronization link ID.
	//
	// example:
	//
	// stl-b3fpik5nq6oy7***
	LinkId *string `json:"LinkId,omitempty" xml:"LinkId,omitempty"`
	// Deprecated
	//
	// The modification time.
	//
	// example:
	//
	// 1572839133000
	ModifedTime *int64 `json:"ModifedTime,omitempty" xml:"ModifedTime,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1572839133000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The execution priority of the synchronization task. Synchronization tasks are executed in descending order of priority. Tasks with the same priority are executed in random order.
	//
	// Valid values: 1 to 5.
	//
	// Default value: 3.
	//
	// example:
	//
	// 3
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The batch synchronization task ID for images, which corresponds to the SyncRecordId (synchronization task record ID) in the request parameters.
	//
	// > When an image matches multiple synchronization rules and generates multiple synchronization tasks, these tasks share the same SyncBatchTaskId.
	//
	// example:
	//
	// 9d8ac4f6-8138-4c15-a2e3-60624ad3****
	SyncBatchTaskId *string `json:"SyncBatchTaskId,omitempty" xml:"SyncBatchTaskId,omitempty"`
	// The synchronization rule ID.
	//
	// example:
	//
	// crsr-7lph66uloi6h****
	SyncRuleId *string `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
	// The synchronization task ID.
	//
	// example:
	//
	// rst-4kfd7fk6pohk****
	SyncTaskId *string `json:"SyncTaskId,omitempty" xml:"SyncTaskId,omitempty"`
	// The synchronization transfer acceleration status.
	//
	// example:
	//
	// true
	SyncTransAccelerate *bool `json:"SyncTransAccelerate,omitempty" xml:"SyncTransAccelerate,omitempty"`
	// The task failure information.
	//
	// > When a synchronization task fails, this field returns information about the failure.
	//
	// example:
	//
	// NETWORK_ERROR
	TaskIssue *string `json:"TaskIssue,omitempty" xml:"TaskIssue,omitempty"`
	// The task status.
	//
	// example:
	//
	// ERROR
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The trigger policy. Valid values:
	//
	// - `PASSIVE`: Synchronization is automatically triggered.
	//
	// - `INITIATIVE`: Synchronization is manually triggered.
	//
	// Default value: `PASSIVE`
	//
	// example:
	//
	// PASSIVE
	TaskTrigger *string `json:"TaskTrigger,omitempty" xml:"TaskTrigger,omitempty"`
}

func (s ListRepoSyncTaskResponseBodySyncTasks) String() string {
	return dara.Prettify(s)
}

func (s ListRepoSyncTaskResponseBodySyncTasks) GoString() string {
	return s.String()
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) GetCrossUser() *bool {
	return s.CrossUser
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) GetCustomLink() *bool {
	return s.CustomLink
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) GetImageFrom() *ListRepoSyncTaskResponseBodySyncTasksImageFrom {
	return s.ImageFrom
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) GetImageTo() *ListRepoSyncTaskResponseBodySyncTasksImageTo {
	return s.ImageTo
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) GetLinkId() *string {
	return s.LinkId
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) GetModifedTime() *int64 {
	return s.ModifedTime
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) GetPriority() *int32 {
	return s.Priority
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) GetSyncBatchTaskId() *string {
	return s.SyncBatchTaskId
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) GetSyncRuleId() *string {
	return s.SyncRuleId
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) GetSyncTaskId() *string {
	return s.SyncTaskId
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) GetSyncTransAccelerate() *bool {
	return s.SyncTransAccelerate
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) GetTaskIssue() *string {
	return s.TaskIssue
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) GetTaskTrigger() *string {
	return s.TaskTrigger
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetCreateTime(v int64) *ListRepoSyncTaskResponseBodySyncTasks {
	s.CreateTime = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetCrossUser(v bool) *ListRepoSyncTaskResponseBodySyncTasks {
	s.CrossUser = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetCustomLink(v bool) *ListRepoSyncTaskResponseBodySyncTasks {
	s.CustomLink = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetImageFrom(v *ListRepoSyncTaskResponseBodySyncTasksImageFrom) *ListRepoSyncTaskResponseBodySyncTasks {
	s.ImageFrom = v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetImageTo(v *ListRepoSyncTaskResponseBodySyncTasksImageTo) *ListRepoSyncTaskResponseBodySyncTasks {
	s.ImageTo = v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetLinkId(v string) *ListRepoSyncTaskResponseBodySyncTasks {
	s.LinkId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetModifedTime(v int64) *ListRepoSyncTaskResponseBodySyncTasks {
	s.ModifedTime = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetModifiedTime(v int64) *ListRepoSyncTaskResponseBodySyncTasks {
	s.ModifiedTime = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetPriority(v int32) *ListRepoSyncTaskResponseBodySyncTasks {
	s.Priority = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetSyncBatchTaskId(v string) *ListRepoSyncTaskResponseBodySyncTasks {
	s.SyncBatchTaskId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetSyncRuleId(v string) *ListRepoSyncTaskResponseBodySyncTasks {
	s.SyncRuleId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetSyncTaskId(v string) *ListRepoSyncTaskResponseBodySyncTasks {
	s.SyncTaskId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetSyncTransAccelerate(v bool) *ListRepoSyncTaskResponseBodySyncTasks {
	s.SyncTransAccelerate = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetTaskIssue(v string) *ListRepoSyncTaskResponseBodySyncTasks {
	s.TaskIssue = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetTaskStatus(v string) *ListRepoSyncTaskResponseBodySyncTasks {
	s.TaskStatus = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetTaskTrigger(v string) *ListRepoSyncTaskResponseBodySyncTasks {
	s.TaskTrigger = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) Validate() error {
	if s.ImageFrom != nil {
		if err := s.ImageFrom.Validate(); err != nil {
			return err
		}
	}
	if s.ImageTo != nil {
		if err := s.ImageTo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRepoSyncTaskResponseBodySyncTasksImageFrom struct {
	// The image tag.
	//
	// example:
	//
	// v0.1
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The repository name.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The repository namespace.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s ListRepoSyncTaskResponseBodySyncTasksImageFrom) String() string {
	return dara.Prettify(s)
}

func (s ListRepoSyncTaskResponseBodySyncTasksImageFrom) GoString() string {
	return s.String()
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageFrom) GetImageTag() *string {
	return s.ImageTag
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageFrom) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageFrom) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageFrom) GetRepoName() *string {
	return s.RepoName
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageFrom) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageFrom) SetImageTag(v string) *ListRepoSyncTaskResponseBodySyncTasksImageFrom {
	s.ImageTag = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageFrom) SetInstanceId(v string) *ListRepoSyncTaskResponseBodySyncTasksImageFrom {
	s.InstanceId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageFrom) SetRegionId(v string) *ListRepoSyncTaskResponseBodySyncTasksImageFrom {
	s.RegionId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageFrom) SetRepoName(v string) *ListRepoSyncTaskResponseBodySyncTasksImageFrom {
	s.RepoName = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageFrom) SetRepoNamespaceName(v string) *ListRepoSyncTaskResponseBodySyncTasksImageFrom {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageFrom) Validate() error {
	return dara.Validate(s)
}

type ListRepoSyncTaskResponseBodySyncTasksImageTo struct {
	// The image tag.
	//
	// example:
	//
	// v0.1
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-k77rd2eo9zttneqo
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The repository name.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The repository namespace.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s ListRepoSyncTaskResponseBodySyncTasksImageTo) String() string {
	return dara.Prettify(s)
}

func (s ListRepoSyncTaskResponseBodySyncTasksImageTo) GoString() string {
	return s.String()
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageTo) GetImageTag() *string {
	return s.ImageTag
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageTo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageTo) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageTo) GetRepoName() *string {
	return s.RepoName
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageTo) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageTo) SetImageTag(v string) *ListRepoSyncTaskResponseBodySyncTasksImageTo {
	s.ImageTag = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageTo) SetInstanceId(v string) *ListRepoSyncTaskResponseBodySyncTasksImageTo {
	s.InstanceId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageTo) SetRegionId(v string) *ListRepoSyncTaskResponseBodySyncTasksImageTo {
	s.RegionId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageTo) SetRepoName(v string) *ListRepoSyncTaskResponseBodySyncTasksImageTo {
	s.RepoName = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageTo) SetRepoNamespaceName(v string) *ListRepoSyncTaskResponseBodySyncTasksImageTo {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageTo) Validate() error {
	return dara.Validate(s)
}
