// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactSubscriptionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListArtifactSubscriptionTaskResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *ListArtifactSubscriptionTaskResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListArtifactSubscriptionTaskResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListArtifactSubscriptionTaskResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListArtifactSubscriptionTaskResponseBody
	GetRequestId() *string
	SetTasks(v []*ListArtifactSubscriptionTaskResponseBodyTasks) *ListArtifactSubscriptionTaskResponseBody
	GetTasks() []*ListArtifactSubscriptionTaskResponseBodyTasks
	SetTotalCount(v int32) *ListArtifactSubscriptionTaskResponseBody
	GetTotalCount() *int32
}

type ListArtifactSubscriptionTaskResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 81E7A039-A4EF-57D9-A100-88E5DCEF9D56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried artifact subscription tasks.
	Tasks []*ListArtifactSubscriptionTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListArtifactSubscriptionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactSubscriptionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactSubscriptionTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListArtifactSubscriptionTaskResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListArtifactSubscriptionTaskResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListArtifactSubscriptionTaskResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListArtifactSubscriptionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListArtifactSubscriptionTaskResponseBody) GetTasks() []*ListArtifactSubscriptionTaskResponseBodyTasks {
	return s.Tasks
}

func (s *ListArtifactSubscriptionTaskResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListArtifactSubscriptionTaskResponseBody) SetCode(v string) *ListArtifactSubscriptionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBody) SetIsSuccess(v bool) *ListArtifactSubscriptionTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBody) SetPageNo(v int32) *ListArtifactSubscriptionTaskResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBody) SetPageSize(v int32) *ListArtifactSubscriptionTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBody) SetRequestId(v string) *ListArtifactSubscriptionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBody) SetTasks(v []*ListArtifactSubscriptionTaskResponseBodyTasks) *ListArtifactSubscriptionTaskResponseBody {
	s.Tasks = v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBody) SetTotalCount(v int32) *ListArtifactSubscriptionTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListArtifactSubscriptionTaskResponseBodyTasks struct {
	// The type of the artifact.
	//
	// example:
	//
	// IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The end time of the artifact subscription task.
	//
	// example:
	//
	// 1692756630000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-7pd01myak****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test-ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
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
	// The type of the source artifact.
	//
	// example:
	//
	// PUBLIC
	SourceRepoType *string `json:"SourceRepoType,omitempty" xml:"SourceRepoType,omitempty"`
	// The start time of the artifact subscription task.
	//
	// example:
	//
	// 1695348301000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of subscribed tags.
	//
	// example:
	//
	// 3
	TagSubscriptionCount *int64 `json:"TagSubscriptionCount,omitempty" xml:"TagSubscriptionCount,omitempty"`
	// The total number of tags.
	//
	// example:
	//
	// 311
	TagTotalCount *int64 `json:"TagTotalCount,omitempty" xml:"TagTotalCount,omitempty"`
	// The task ID.
	//
	// example:
	//
	// crast-40le4es9yh0p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task result.
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
	// The type of the task.
	//
	// example:
	//
	// AUTO
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListArtifactSubscriptionTaskResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactSubscriptionTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetEndTime() *string {
	return s.EndTime
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetMessage() *string {
	return s.Message
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetRepoName() *string {
	return s.RepoName
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetSourceNamespaceName() *string {
	return s.SourceNamespaceName
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetSourceProvider() *string {
	return s.SourceProvider
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetSourceRepoName() *string {
	return s.SourceRepoName
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetSourceRepoType() *string {
	return s.SourceRepoType
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetStartTime() *string {
	return s.StartTime
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetTagSubscriptionCount() *int64 {
	return s.TagSubscriptionCount
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetTagTotalCount() *int64 {
	return s.TagTotalCount
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetTaskResult() *string {
	return s.TaskResult
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetArtifactType(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.ArtifactType = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetEndTime(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.EndTime = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetInstanceId(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.InstanceId = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetMessage(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.Message = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetNamespaceName(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.NamespaceName = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetRepoName(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.RepoName = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetSourceNamespaceName(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.SourceNamespaceName = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetSourceProvider(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.SourceProvider = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetSourceRepoName(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.SourceRepoName = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetSourceRepoType(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.SourceRepoType = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetStartTime(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.StartTime = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetTagSubscriptionCount(v int64) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.TagSubscriptionCount = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetTagTotalCount(v int64) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.TagTotalCount = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetTaskId(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetTaskResult(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.TaskResult = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetTaskStatus(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.TaskStatus = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetTaskType(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.TaskType = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
