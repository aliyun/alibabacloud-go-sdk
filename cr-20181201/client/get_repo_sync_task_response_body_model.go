// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoSyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRepoSyncTaskResponseBody
	GetCode() *string
	SetCrossUser(v bool) *GetRepoSyncTaskResponseBody
	GetCrossUser() *bool
	SetImageFrom(v *GetRepoSyncTaskResponseBodyImageFrom) *GetRepoSyncTaskResponseBody
	GetImageFrom() *GetRepoSyncTaskResponseBodyImageFrom
	SetImageTo(v *GetRepoSyncTaskResponseBodyImageTo) *GetRepoSyncTaskResponseBody
	GetImageTo() *GetRepoSyncTaskResponseBodyImageTo
	SetIsSuccess(v bool) *GetRepoSyncTaskResponseBody
	GetIsSuccess() *bool
	SetLayerTasks(v []*GetRepoSyncTaskResponseBodyLayerTasks) *GetRepoSyncTaskResponseBody
	GetLayerTasks() []*GetRepoSyncTaskResponseBodyLayerTasks
	SetPriority(v int32) *GetRepoSyncTaskResponseBody
	GetPriority() *int32
	SetProgress(v int64) *GetRepoSyncTaskResponseBody
	GetProgress() *int64
	SetRequestId(v string) *GetRepoSyncTaskResponseBody
	GetRequestId() *string
	SetSyncBatchTaskId(v string) *GetRepoSyncTaskResponseBody
	GetSyncBatchTaskId() *string
	SetSyncRuleId(v string) *GetRepoSyncTaskResponseBody
	GetSyncRuleId() *string
	SetSyncTaskId(v string) *GetRepoSyncTaskResponseBody
	GetSyncTaskId() *string
	SetSyncTransAccelerate(v bool) *GetRepoSyncTaskResponseBody
	GetSyncTransAccelerate() *bool
	SetSyncedSize(v int64) *GetRepoSyncTaskResponseBody
	GetSyncedSize() *int64
	SetTaskIssue(v string) *GetRepoSyncTaskResponseBody
	GetTaskIssue() *string
	SetTaskStatus(v string) *GetRepoSyncTaskResponseBody
	GetTaskStatus() *string
	SetTaskTrigger(v string) *GetRepoSyncTaskResponseBody
	GetTaskTrigger() *string
}

type GetRepoSyncTaskResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the synchronization is cross-account.
	//
	// example:
	//
	// true
	CrossUser *bool `json:"CrossUser,omitempty" xml:"CrossUser,omitempty"`
	// The source image.
	ImageFrom *GetRepoSyncTaskResponseBodyImageFrom `json:"ImageFrom,omitempty" xml:"ImageFrom,omitempty" type:"Struct"`
	// The destination image.
	ImageTo *GetRepoSyncTaskResponseBodyImageTo `json:"ImageTo,omitempty" xml:"ImageTo,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The list of image layer synchronization tasks.
	LayerTasks []*GetRepoSyncTaskResponseBodyLayerTasks `json:"LayerTasks,omitempty" xml:"LayerTasks,omitempty" type:"Repeated"`
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
	// The synchronization progress. Valid values:
	//
	// - `0`: The synchronization has just started or failed.
	//
	// - `1`: The synchronization succeeded.
	//
	// example:
	//
	// 1
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A6DEF8B0-5D45-46D6-867D-8C7FF0966B07
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The synchronization batch task ID.
	//
	// example:
	//
	// a9434731-95ef-4087-9cf4-369c8e90****
	SyncBatchTaskId *string `json:"SyncBatchTaskId,omitempty" xml:"SyncBatchTaskId,omitempty"`
	// The synchronization rule ID.
	//
	// example:
	//
	// crsr-cllro6ho3wne****
	SyncRuleId *string `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
	// The synchronization task ID.
	//
	// example:
	//
	// rst-zxjkiv5oil6f****
	SyncTaskId *string `json:"SyncTaskId,omitempty" xml:"SyncTaskId,omitempty"`
	// Indicates whether transfer acceleration is enabled for synchronization.
	//
	// example:
	//
	// true
	SyncTransAccelerate *bool `json:"SyncTransAccelerate,omitempty" xml:"SyncTransAccelerate,omitempty"`
	// The synchronized size, in bytes.
	//
	// example:
	//
	// 23655489
	SyncedSize *int64 `json:"SyncedSize,omitempty" xml:"SyncedSize,omitempty"`
	// The task failure information.
	//
	// > When the synchronization task fails, this field returns information about the failure.
	//
	// example:
	//
	// NETWORK_ERROR
	TaskIssue *string `json:"TaskIssue,omitempty" xml:"TaskIssue,omitempty"`
	// The task status. Valid values:
	//
	// `PENDING`: The synchronization is pending.
	//
	// `SYNCHRONIZING`: The synchronization is in progress.
	//
	// `SUCCESS`: The synchronization succeeded.
	//
	// `ERROR`: The synchronization failed.
	//
	// `CANCELED`: The synchronization task is canceled.
	//
	// example:
	//
	// SUCCESS
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The trigger type of the synchronization task. Valid values:
	//
	// `PASSIVE`: The synchronization task is automatically triggered.
	//
	// `INITIATIVE`: The synchronization task is manually triggered.
	//
	// example:
	//
	// PASSIVE
	TaskTrigger *string `json:"TaskTrigger,omitempty" xml:"TaskTrigger,omitempty"`
}

func (s GetRepoSyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRepoSyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepoSyncTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRepoSyncTaskResponseBody) GetCrossUser() *bool {
	return s.CrossUser
}

func (s *GetRepoSyncTaskResponseBody) GetImageFrom() *GetRepoSyncTaskResponseBodyImageFrom {
	return s.ImageFrom
}

func (s *GetRepoSyncTaskResponseBody) GetImageTo() *GetRepoSyncTaskResponseBodyImageTo {
	return s.ImageTo
}

func (s *GetRepoSyncTaskResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetRepoSyncTaskResponseBody) GetLayerTasks() []*GetRepoSyncTaskResponseBodyLayerTasks {
	return s.LayerTasks
}

func (s *GetRepoSyncTaskResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *GetRepoSyncTaskResponseBody) GetProgress() *int64 {
	return s.Progress
}

func (s *GetRepoSyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRepoSyncTaskResponseBody) GetSyncBatchTaskId() *string {
	return s.SyncBatchTaskId
}

func (s *GetRepoSyncTaskResponseBody) GetSyncRuleId() *string {
	return s.SyncRuleId
}

func (s *GetRepoSyncTaskResponseBody) GetSyncTaskId() *string {
	return s.SyncTaskId
}

func (s *GetRepoSyncTaskResponseBody) GetSyncTransAccelerate() *bool {
	return s.SyncTransAccelerate
}

func (s *GetRepoSyncTaskResponseBody) GetSyncedSize() *int64 {
	return s.SyncedSize
}

func (s *GetRepoSyncTaskResponseBody) GetTaskIssue() *string {
	return s.TaskIssue
}

func (s *GetRepoSyncTaskResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetRepoSyncTaskResponseBody) GetTaskTrigger() *string {
	return s.TaskTrigger
}

func (s *GetRepoSyncTaskResponseBody) SetCode(v string) *GetRepoSyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetCrossUser(v bool) *GetRepoSyncTaskResponseBody {
	s.CrossUser = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetImageFrom(v *GetRepoSyncTaskResponseBodyImageFrom) *GetRepoSyncTaskResponseBody {
	s.ImageFrom = v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetImageTo(v *GetRepoSyncTaskResponseBodyImageTo) *GetRepoSyncTaskResponseBody {
	s.ImageTo = v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetIsSuccess(v bool) *GetRepoSyncTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetLayerTasks(v []*GetRepoSyncTaskResponseBodyLayerTasks) *GetRepoSyncTaskResponseBody {
	s.LayerTasks = v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetPriority(v int32) *GetRepoSyncTaskResponseBody {
	s.Priority = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetProgress(v int64) *GetRepoSyncTaskResponseBody {
	s.Progress = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetRequestId(v string) *GetRepoSyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetSyncBatchTaskId(v string) *GetRepoSyncTaskResponseBody {
	s.SyncBatchTaskId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetSyncRuleId(v string) *GetRepoSyncTaskResponseBody {
	s.SyncRuleId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetSyncTaskId(v string) *GetRepoSyncTaskResponseBody {
	s.SyncTaskId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetSyncTransAccelerate(v bool) *GetRepoSyncTaskResponseBody {
	s.SyncTransAccelerate = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetSyncedSize(v int64) *GetRepoSyncTaskResponseBody {
	s.SyncedSize = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetTaskIssue(v string) *GetRepoSyncTaskResponseBody {
	s.TaskIssue = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetTaskStatus(v string) *GetRepoSyncTaskResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetTaskTrigger(v string) *GetRepoSyncTaskResponseBody {
	s.TaskTrigger = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) Validate() error {
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
	if s.LayerTasks != nil {
		for _, item := range s.LayerTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRepoSyncTaskResponseBodyImageFrom struct {
	// The image tag.
	//
	// example:
	//
	// master
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-sgedpenzw80e****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region.
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
	// The namespace name.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s GetRepoSyncTaskResponseBodyImageFrom) String() string {
	return dara.Prettify(s)
}

func (s GetRepoSyncTaskResponseBodyImageFrom) GoString() string {
	return s.String()
}

func (s *GetRepoSyncTaskResponseBodyImageFrom) GetImageTag() *string {
	return s.ImageTag
}

func (s *GetRepoSyncTaskResponseBodyImageFrom) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRepoSyncTaskResponseBodyImageFrom) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRepoSyncTaskResponseBodyImageFrom) GetRepoName() *string {
	return s.RepoName
}

func (s *GetRepoSyncTaskResponseBodyImageFrom) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *GetRepoSyncTaskResponseBodyImageFrom) SetImageTag(v string) *GetRepoSyncTaskResponseBodyImageFrom {
	s.ImageTag = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageFrom) SetInstanceId(v string) *GetRepoSyncTaskResponseBodyImageFrom {
	s.InstanceId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageFrom) SetRegionId(v string) *GetRepoSyncTaskResponseBodyImageFrom {
	s.RegionId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageFrom) SetRepoName(v string) *GetRepoSyncTaskResponseBodyImageFrom {
	s.RepoName = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageFrom) SetRepoNamespaceName(v string) *GetRepoSyncTaskResponseBodyImageFrom {
	s.RepoNamespaceName = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageFrom) Validate() error {
	return dara.Validate(s)
}

type GetRepoSyncTaskResponseBodyImageTo struct {
	// The image tag.
	//
	// example:
	//
	// master
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-leqzomz5vijc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The repository name.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace name.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s GetRepoSyncTaskResponseBodyImageTo) String() string {
	return dara.Prettify(s)
}

func (s GetRepoSyncTaskResponseBodyImageTo) GoString() string {
	return s.String()
}

func (s *GetRepoSyncTaskResponseBodyImageTo) GetImageTag() *string {
	return s.ImageTag
}

func (s *GetRepoSyncTaskResponseBodyImageTo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRepoSyncTaskResponseBodyImageTo) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRepoSyncTaskResponseBodyImageTo) GetRepoName() *string {
	return s.RepoName
}

func (s *GetRepoSyncTaskResponseBodyImageTo) GetRepoNamespaceName() *string {
	return s.RepoNamespaceName
}

func (s *GetRepoSyncTaskResponseBodyImageTo) SetImageTag(v string) *GetRepoSyncTaskResponseBodyImageTo {
	s.ImageTag = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageTo) SetInstanceId(v string) *GetRepoSyncTaskResponseBodyImageTo {
	s.InstanceId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageTo) SetRegionId(v string) *GetRepoSyncTaskResponseBodyImageTo {
	s.RegionId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageTo) SetRepoName(v string) *GetRepoSyncTaskResponseBodyImageTo {
	s.RepoName = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageTo) SetRepoNamespaceName(v string) *GetRepoSyncTaskResponseBodyImageTo {
	s.RepoNamespaceName = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageTo) Validate() error {
	return dara.Validate(s)
}

type GetRepoSyncTaskResponseBodyLayerTasks struct {
	// The digest value of the artifact.
	//
	// example:
	//
	// sha256:36fb85fcb5e919cb60e782397a6be04201868fe7b38ef7669fc01caec1c8fc4e
	ArtifactDigest *string `json:"ArtifactDigest,omitempty" xml:"ArtifactDigest,omitempty"`
	// The image digest value.
	//
	// example:
	//
	// sha256:36fb85fcb5e919cb60e782397a6be04201868fe7b38ef7669fc01caec1c8fc4e
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The size.
	//
	// example:
	//
	// 23655489
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The synchronization layer task ID.
	//
	// example:
	//
	// rslt-074x4q20fx2d****
	SyncLayerTaskId *string `json:"SyncLayerTaskId,omitempty" xml:"SyncLayerTaskId,omitempty"`
	// The synchronized size.
	//
	// example:
	//
	// 23655489
	SyncedSize *int64 `json:"SyncedSize,omitempty" xml:"SyncedSize,omitempty"`
	// The task status.
	//
	// example:
	//
	// SUCCESS
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetRepoSyncTaskResponseBodyLayerTasks) String() string {
	return dara.Prettify(s)
}

func (s GetRepoSyncTaskResponseBodyLayerTasks) GoString() string {
	return s.String()
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) GetArtifactDigest() *string {
	return s.ArtifactDigest
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) GetDigest() *string {
	return s.Digest
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) GetSize() *int64 {
	return s.Size
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) GetSyncLayerTaskId() *string {
	return s.SyncLayerTaskId
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) GetSyncedSize() *int64 {
	return s.SyncedSize
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) SetArtifactDigest(v string) *GetRepoSyncTaskResponseBodyLayerTasks {
	s.ArtifactDigest = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) SetDigest(v string) *GetRepoSyncTaskResponseBodyLayerTasks {
	s.Digest = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) SetSize(v int64) *GetRepoSyncTaskResponseBodyLayerTasks {
	s.Size = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) SetSyncLayerTaskId(v string) *GetRepoSyncTaskResponseBodyLayerTasks {
	s.SyncLayerTaskId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) SetSyncedSize(v int64) *GetRepoSyncTaskResponseBodyLayerTasks {
	s.SyncedSize = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) SetTaskStatus(v string) *GetRepoSyncTaskResponseBodyLayerTasks {
	s.TaskStatus = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) Validate() error {
	return dara.Validate(s)
}
