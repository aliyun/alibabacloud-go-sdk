// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactSubscriptionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactType(v string) *GetArtifactSubscriptionTaskResponseBody
	GetArtifactType() *string
	SetCode(v string) *GetArtifactSubscriptionTaskResponseBody
	GetCode() *string
	SetEndTime(v int64) *GetArtifactSubscriptionTaskResponseBody
	GetEndTime() *int64
	SetInstanceId(v string) *GetArtifactSubscriptionTaskResponseBody
	GetInstanceId() *string
	SetIsSuccess(v bool) *GetArtifactSubscriptionTaskResponseBody
	GetIsSuccess() *bool
	SetMessage(v string) *GetArtifactSubscriptionTaskResponseBody
	GetMessage() *string
	SetNamespaceName(v string) *GetArtifactSubscriptionTaskResponseBody
	GetNamespaceName() *string
	SetRepoName(v string) *GetArtifactSubscriptionTaskResponseBody
	GetRepoName() *string
	SetRequestId(v string) *GetArtifactSubscriptionTaskResponseBody
	GetRequestId() *string
	SetSourceNamespaceName(v string) *GetArtifactSubscriptionTaskResponseBody
	GetSourceNamespaceName() *string
	SetSourceProvider(v string) *GetArtifactSubscriptionTaskResponseBody
	GetSourceProvider() *string
	SetSourceRepoName(v string) *GetArtifactSubscriptionTaskResponseBody
	GetSourceRepoName() *string
	SetSourceRepoType(v string) *GetArtifactSubscriptionTaskResponseBody
	GetSourceRepoType() *string
	SetStartTime(v int64) *GetArtifactSubscriptionTaskResponseBody
	GetStartTime() *int64
	SetTagSubscriptionCount(v int64) *GetArtifactSubscriptionTaskResponseBody
	GetTagSubscriptionCount() *int64
	SetTagTotalCount(v int64) *GetArtifactSubscriptionTaskResponseBody
	GetTagTotalCount() *int64
	SetTaskId(v string) *GetArtifactSubscriptionTaskResponseBody
	GetTaskId() *string
	SetTaskResult(v string) *GetArtifactSubscriptionTaskResponseBody
	GetTaskResult() *string
	SetTaskStatus(v string) *GetArtifactSubscriptionTaskResponseBody
	GetTaskStatus() *string
	SetTaskType(v string) *GetArtifactSubscriptionTaskResponseBody
	GetTaskType() *string
}

type GetArtifactSubscriptionTaskResponseBody struct {
	// The artifact type.
	//
	// example:
	//
	// IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The end time of the artifact subscription task.
	//
	// example:
	//
	// 1691979202000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the Container Registry namespace.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The name of the Container Registry repository.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 12589EF7-96E2-4554-AAD7-F7209E88CAD3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the source namespace.
	//
	// example:
	//
	// library
	SourceNamespaceName *string `json:"SourceNamespaceName,omitempty" xml:"SourceNamespaceName,omitempty"`
	// The artifact source.
	//
	// example:
	//
	// DOCKER_HUB
	SourceProvider *string `json:"SourceProvider,omitempty" xml:"SourceProvider,omitempty"`
	// The name of the source repository.
	//
	// example:
	//
	// nginx
	SourceRepoName *string `json:"SourceRepoName,omitempty" xml:"SourceRepoName,omitempty"`
	// The type of the source repository.
	//
	// example:
	//
	// PUBLIC
	SourceRepoType *string `json:"SourceRepoType,omitempty" xml:"SourceRepoType,omitempty"`
	// The start time of the artifact subscription task.
	//
	// example:
	//
	// 1568718468000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total subscribed tags.
	//
	// example:
	//
	// 1
	TagSubscriptionCount *int64 `json:"TagSubscriptionCount,omitempty" xml:"TagSubscriptionCount,omitempty"`
	// The total number of tags.
	//
	// example:
	//
	// 6
	TagTotalCount *int64 `json:"TagTotalCount,omitempty" xml:"TagTotalCount,omitempty"`
	// The task ID.
	//
	// example:
	//
	// crast-40le4es9yh0p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task results.
	//
	// example:
	//
	// SUCCESS
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// RUNNING
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The type of the task. Valid values:
	//
	// example:
	//
	// AUTO
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetArtifactSubscriptionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactSubscriptionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetRepoName() *string {
	return s.RepoName
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetSourceNamespaceName() *string {
	return s.SourceNamespaceName
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetSourceProvider() *string {
	return s.SourceProvider
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetSourceRepoName() *string {
	return s.SourceRepoName
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetSourceRepoType() *string {
	return s.SourceRepoType
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetTagSubscriptionCount() *int64 {
	return s.TagSubscriptionCount
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetTagTotalCount() *int64 {
	return s.TagTotalCount
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetTaskResult() *string {
	return s.TaskResult
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetArtifactSubscriptionTaskResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetArtifactType(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetCode(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetEndTime(v int64) *GetArtifactSubscriptionTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetInstanceId(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetIsSuccess(v bool) *GetArtifactSubscriptionTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetMessage(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetNamespaceName(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.NamespaceName = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetRepoName(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.RepoName = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetRequestId(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetSourceNamespaceName(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.SourceNamespaceName = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetSourceProvider(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.SourceProvider = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetSourceRepoName(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.SourceRepoName = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetSourceRepoType(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.SourceRepoType = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetStartTime(v int64) *GetArtifactSubscriptionTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetTagSubscriptionCount(v int64) *GetArtifactSubscriptionTaskResponseBody {
	s.TagSubscriptionCount = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetTagTotalCount(v int64) *GetArtifactSubscriptionTaskResponseBody {
	s.TagTotalCount = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetTaskId(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetTaskResult(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.TaskResult = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetTaskStatus(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetTaskType(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
