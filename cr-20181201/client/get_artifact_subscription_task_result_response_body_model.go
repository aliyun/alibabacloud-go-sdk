// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactSubscriptionTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetArtifactSubscriptionTaskResultResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *GetArtifactSubscriptionTaskResultResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *GetArtifactSubscriptionTaskResultResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *GetArtifactSubscriptionTaskResultResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetArtifactSubscriptionTaskResultResponseBody
	GetRequestId() *string
	SetTaskResults(v []*GetArtifactSubscriptionTaskResultResponseBodyTaskResults) *GetArtifactSubscriptionTaskResultResponseBody
	GetTaskResults() []*GetArtifactSubscriptionTaskResultResponseBodyTaskResults
	SetTotalCount(v int32) *GetArtifactSubscriptionTaskResultResponseBody
	GetTotalCount() *int32
}

type GetArtifactSubscriptionTaskResultResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
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
	// 0A8768F6-9B47-5127-A075-9CFB9F79181F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the artifact subscription task.
	TaskResults []*GetArtifactSubscriptionTaskResultResponseBodyTaskResults `json:"TaskResults,omitempty" xml:"TaskResults,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetArtifactSubscriptionTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactSubscriptionTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) GetTaskResults() []*GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	return s.TaskResults
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) SetCode(v string) *GetArtifactSubscriptionTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) SetIsSuccess(v bool) *GetArtifactSubscriptionTaskResultResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) SetPageNo(v int32) *GetArtifactSubscriptionTaskResultResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) SetPageSize(v int32) *GetArtifactSubscriptionTaskResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) SetRequestId(v string) *GetArtifactSubscriptionTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) SetTaskResults(v []*GetArtifactSubscriptionTaskResultResponseBodyTaskResults) *GetArtifactSubscriptionTaskResultResponseBody {
	s.TaskResults = v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) SetTotalCount(v int32) *GetArtifactSubscriptionTaskResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetArtifactSubscriptionTaskResultResponseBodyTaskResults struct {
	// The end time of the subscription task.
	//
	// example:
	//
	// 1692756630000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-isj2wgaw4z9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// test-reop
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The result of the task.
	//
	// example:
	//
	// SUCCESS
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The start time of the subscription task.
	//
	// example:
	//
	// 1691719501000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The image tag.
	//
	// example:
	//
	// v2.0
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The task ID.
	//
	// example:
	//
	// crast-wkpfwqozjiq****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetArtifactSubscriptionTaskResultResponseBodyTaskResults) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactSubscriptionTaskResultResponseBodyTaskResults) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) GetRepoName() *string {
	return s.RepoName
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) GetResult() *string {
	return s.Result
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) GetStatus() *string {
	return s.Status
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) GetTag() *string {
	return s.Tag
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) GetTaskId() *string {
	return s.TaskId
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetEndTime(v int64) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.EndTime = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetInstanceId(v string) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetNamespaceName(v string) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.NamespaceName = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetRepoName(v string) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.RepoName = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetResult(v string) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.Result = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetStartTime(v int64) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.StartTime = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetStatus(v string) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.Status = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetTag(v string) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.Tag = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetTaskId(v string) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.TaskId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) Validate() error {
	return dara.Validate(s)
}
