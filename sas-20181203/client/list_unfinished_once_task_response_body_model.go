// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnfinishedOnceTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOnceTasks(v []*ListUnfinishedOnceTaskResponseBodyOnceTasks) *ListUnfinishedOnceTaskResponseBody
	GetOnceTasks() []*ListUnfinishedOnceTaskResponseBodyOnceTasks
	SetRequestId(v string) *ListUnfinishedOnceTaskResponseBody
	GetRequestId() *string
}

type ListUnfinishedOnceTaskResponseBody struct {
	// The details of the tasks.
	OnceTasks []*ListUnfinishedOnceTaskResponseBodyOnceTasks `json:"OnceTasks,omitempty" xml:"OnceTasks,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F5CF78A7-30AA-59DB-847F-13EE3AE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUnfinishedOnceTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUnfinishedOnceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListUnfinishedOnceTaskResponseBody) GetOnceTasks() []*ListUnfinishedOnceTaskResponseBodyOnceTasks {
	return s.OnceTasks
}

func (s *ListUnfinishedOnceTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUnfinishedOnceTaskResponseBody) SetOnceTasks(v []*ListUnfinishedOnceTaskResponseBodyOnceTasks) *ListUnfinishedOnceTaskResponseBody {
	s.OnceTasks = v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBody) SetRequestId(v string) *ListUnfinishedOnceTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBody) Validate() error {
	if s.OnceTasks != nil {
		for _, item := range s.OnceTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUnfinishedOnceTaskResponseBodyOnceTasks struct {
	// The time when the task ends.
	//
	// example:
	//
	// 1670307567000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the task is complete. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	Finish *int32 `json:"Finish,omitempty" xml:"Finish,omitempty"`
	// The number of assets on which the task is complete.
	//
	// example:
	//
	// 67
	FinishCount *int32 `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	// The progress percentage of the task.
	//
	// example:
	//
	// 75
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The execution duration of the task.
	//
	// example:
	//
	// 1670307567000
	RealRunTime *int64 `json:"RealRunTime,omitempty" xml:"RealRunTime,omitempty"`
	// The execution result of the task.
	//
	// example:
	//
	// TASK_NOT_SUPPORT_REGION
	ResultInfo *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	// The time when the task is started.
	//
	// example:
	//
	// 1640102400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **INIT**: The task is not started.
	//
	// 	- **START**: The task is started.
	//
	// 	- **SUCCESS**: The task is complete.
	//
	// 	- **TIMEOUT**: The task timed out.
	//
	// example:
	//
	// SUCCESS
	StatusText *string `json:"StatusText,omitempty" xml:"StatusText,omitempty"`
	// The objective of the task.
	//
	// example:
	//
	// 238cf050a7270dd6940602e70f1e5a11eeaf4e02035f445b7f613ff5e064****
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The type of the assets that are scanned. Valid values:
	//
	// 	- **IMAGE_REPO**: image repository
	//
	// 	- **IMAGE**: image
	//
	// example:
	//
	// IMAGE
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 9fb50f2af8bb67c9fdb684194c83****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The information about the image scan task.
	TaskImageInfo *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo `json:"TaskImageInfo,omitempty" xml:"TaskImageInfo,omitempty" type:"Struct"`
	// The name of the task.
	//
	// example:
	//
	// IMAGE_SCAN
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task.
	//
	// example:
	//
	// IMAGE_SCAN
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUnfinishedOnceTaskResponseBodyOnceTasks) String() string {
	return dara.Prettify(s)
}

func (s ListUnfinishedOnceTaskResponseBodyOnceTasks) GoString() string {
	return s.String()
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) GetFinish() *int32 {
	return s.Finish
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) GetFinishCount() *int32 {
	return s.FinishCount
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) GetProgress() *int64 {
	return s.Progress
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) GetRealRunTime() *int64 {
	return s.RealRunTime
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) GetResultInfo() *string {
	return s.ResultInfo
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) GetStatus() *int32 {
	return s.Status
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) GetStatusText() *string {
	return s.StatusText
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) GetTarget() *string {
	return s.Target
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) GetTargetType() *string {
	return s.TargetType
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) GetTaskImageInfo() *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo {
	return s.TaskImageInfo
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) SetEndTime(v int64) *ListUnfinishedOnceTaskResponseBodyOnceTasks {
	s.EndTime = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) SetFinish(v int32) *ListUnfinishedOnceTaskResponseBodyOnceTasks {
	s.Finish = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) SetFinishCount(v int32) *ListUnfinishedOnceTaskResponseBodyOnceTasks {
	s.FinishCount = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) SetProgress(v int64) *ListUnfinishedOnceTaskResponseBodyOnceTasks {
	s.Progress = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) SetRealRunTime(v int64) *ListUnfinishedOnceTaskResponseBodyOnceTasks {
	s.RealRunTime = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) SetResultInfo(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasks {
	s.ResultInfo = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) SetStartTime(v int64) *ListUnfinishedOnceTaskResponseBodyOnceTasks {
	s.StartTime = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) SetStatus(v int32) *ListUnfinishedOnceTaskResponseBodyOnceTasks {
	s.Status = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) SetStatusText(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasks {
	s.StatusText = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) SetTarget(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasks {
	s.Target = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) SetTargetType(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasks {
	s.TargetType = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) SetTaskId(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasks {
	s.TaskId = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) SetTaskImageInfo(v *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) *ListUnfinishedOnceTaskResponseBodyOnceTasks {
	s.TaskImageInfo = v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) SetTaskName(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasks {
	s.TaskName = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) SetTaskType(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasks {
	s.TaskType = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) SetTotalCount(v int32) *ListUnfinishedOnceTaskResponseBodyOnceTasks {
	s.TotalCount = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasks) Validate() error {
	if s.TaskImageInfo != nil {
		if err := s.TaskImageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo struct {
	// The name of the application.
	//
	// example:
	//
	// ack-jenkins-****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// a765ba1435e7f9446065370e9a41****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// ACK-test-****
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The image digest.
	//
	// example:
	//
	// default_digest
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The container image.
	//
	// example:
	//
	// ***s.com/sas_test/baseli***
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The instance ID of the node.
	//
	// example:
	//
	// i-0xi5mxvtmfw9****
	NodeInstanceId *string `json:"NodeInstanceId,omitempty" xml:"NodeInstanceId,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// 172.18.XXX.XXX
	NodeIp *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// pztest****
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The pod of the image.
	//
	// example:
	//
	// expoit-law-****
	Pod *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
	// The region ID of the server image.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-r88w2vryp8m****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// testyyy
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// bitn***
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The region ID of the image repository.
	//
	// example:
	//
	// cn-hangzhou
	RepoRegionId *string `json:"RepoRegionId,omitempty" xml:"RepoRegionId,omitempty"`
	// The image tag.
	//
	// example:
	//
	// v1.20-002-a2*****
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) GoString() string {
	return s.String()
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) GetAppName() *string {
	return s.AppName
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) GetDigest() *string {
	return s.Digest
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) GetImage() *string {
	return s.Image
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) GetNodeInstanceId() *string {
	return s.NodeInstanceId
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) GetNodeIp() *string {
	return s.NodeIp
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) GetNodeName() *string {
	return s.NodeName
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) GetPod() *string {
	return s.Pod
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) GetRepoId() *string {
	return s.RepoId
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) GetRepoName() *string {
	return s.RepoName
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) GetRepoRegionId() *string {
	return s.RepoRegionId
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) GetTag() *string {
	return s.Tag
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) SetAppName(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo {
	s.AppName = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) SetClusterId(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo {
	s.ClusterId = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) SetClusterName(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo {
	s.ClusterName = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) SetDigest(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo {
	s.Digest = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) SetImage(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo {
	s.Image = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) SetNodeInstanceId(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo {
	s.NodeInstanceId = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) SetNodeIp(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo {
	s.NodeIp = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) SetNodeName(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo {
	s.NodeName = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) SetPod(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo {
	s.Pod = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) SetRegionId(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo {
	s.RegionId = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) SetRepoId(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo {
	s.RepoId = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) SetRepoName(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo {
	s.RepoName = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) SetRepoNamespace(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo {
	s.RepoNamespace = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) SetRepoRegionId(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo {
	s.RepoRegionId = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) SetTag(v string) *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo {
	s.Tag = &v
	return s
}

func (s *ListUnfinishedOnceTaskResponseBodyOnceTasksTaskImageInfo) Validate() error {
	return dara.Validate(s)
}
