// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOnceTaskLeafRecordPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOnceTasks(v []*DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) *DescribeOnceTaskLeafRecordPageResponseBody
	GetOnceTasks() []*DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks
	SetPageInfo(v *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo) *DescribeOnceTaskLeafRecordPageResponseBody
	GetPageInfo() *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo
	SetRequestId(v string) *DescribeOnceTaskLeafRecordPageResponseBody
	GetRequestId() *string
}

type DescribeOnceTaskLeafRecordPageResponseBody struct {
	// The details of tasks.
	OnceTasks []*DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks `json:"OnceTasks,omitempty" xml:"OnceTasks,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB393****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOnceTaskLeafRecordPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnceTaskLeafRecordPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOnceTaskLeafRecordPageResponseBody) GetOnceTasks() []*DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	return s.OnceTasks
}

func (s *DescribeOnceTaskLeafRecordPageResponseBody) GetPageInfo() *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeOnceTaskLeafRecordPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOnceTaskLeafRecordPageResponseBody) SetOnceTasks(v []*DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) *DescribeOnceTaskLeafRecordPageResponseBody {
	s.OnceTasks = v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBody) SetPageInfo(v *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo) *DescribeOnceTaskLeafRecordPageResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBody) SetRequestId(v string) *DescribeOnceTaskLeafRecordPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBody) Validate() error {
	if s.OnceTasks != nil {
		for _, item := range s.OnceTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks struct {
	// The time when the sub-task ends.
	//
	// example:
	//
	// 1670307567000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the sub-task is complete.
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	Finish *int32 `json:"Finish,omitempty" xml:"Finish,omitempty"`
	// The number of the assets that are scanned.
	//
	// example:
	//
	// 67
	FinishCount *string `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	// The progress percentage of the sub-task.
	//
	// example:
	//
	// 75
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The execution duration of the sub-task.
	//
	// example:
	//
	// 1670307567000
	RealRunTime *int64 `json:"RealRunTime,omitempty" xml:"RealRunTime,omitempty"`
	// The execution result.
	//
	// example:
	//
	// TASK_NOT_SUPPORT_REGION
	ResultInfo *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	// The time when the sub-task starts.
	//
	// example:
	//
	// 1640102400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status.
	//
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The sub-task status. Valid values:
	//
	// 	- **INIT**: The sub-task is not started.
	//
	// 	- **START**: The sub-task is started.
	//
	// 	- **SUCCESS**: The sub-task is complete.
	//
	// 	- **TIMEOUT**: The sub-task timed out.
	//
	// example:
	//
	// SUCCESS
	StatusText *string `json:"StatusText,omitempty" xml:"StatusText,omitempty"`
	// The objective of the sub-task.
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
	// The sub-task ID.
	//
	// example:
	//
	// 9fb50f2af8bb67c9fdb684194c83****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The information about the image scan.
	TaskImageInfo *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo `json:"TaskImageInfo,omitempty" xml:"TaskImageInfo,omitempty" type:"Struct"`
	// The name of the sub-task.
	//
	// example:
	//
	// IMAGE_SCAN
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the sub-task.
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
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GoString() string {
	return s.String()
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GetFinish() *int32 {
	return s.Finish
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GetFinishCount() *string {
	return s.FinishCount
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GetProgress() *int64 {
	return s.Progress
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GetRealRunTime() *int64 {
	return s.RealRunTime
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GetResultInfo() *string {
	return s.ResultInfo
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GetStatus() *string {
	return s.Status
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GetStatusText() *string {
	return s.StatusText
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GetTarget() *string {
	return s.Target
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GetTaskImageInfo() *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	return s.TaskImageInfo
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) SetEndTime(v int64) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	s.EndTime = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) SetFinish(v int32) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	s.Finish = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) SetFinishCount(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	s.FinishCount = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) SetProgress(v int64) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	s.Progress = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) SetRealRunTime(v int64) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	s.RealRunTime = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) SetResultInfo(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	s.ResultInfo = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) SetStartTime(v int64) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	s.StartTime = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) SetStatus(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	s.Status = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) SetStatusText(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	s.StatusText = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) SetTarget(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	s.Target = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) SetTargetType(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	s.TargetType = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) SetTaskId(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) SetTaskImageInfo(v *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	s.TaskImageInfo = v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) SetTaskName(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	s.TaskName = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) SetTaskType(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	s.TaskType = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) SetTotalCount(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks {
	s.TotalCount = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasks) Validate() error {
	if s.TaskImageInfo != nil {
		if err := s.TaskImageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo struct {
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
	// The cluster name.
	//
	// example:
	//
	// ACK-test-****
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The time consumed. The value is in the JSON format. The end time of each item is displayed.
	//
	// 	- **vul**: system vulnerabilities
	//
	// 	- **scaVul**: application vulnerabilities
	//
	// 	- **scaVul**: baseline
	//
	// 	- **binary**: binary
	//
	// 	- **forbiddenPackageInfo**: information about the prohibited package
	//
	// 	- **identificationInfo**: identity authentication
	//
	// 	- **script**: malicious scripts
	//
	// 	- **sensitiveFile**: sensitive files
	//
	// 	- **sensitiveInfo**: AccessKey pair leaks
	//
	// 	- **webshell**: website scripts
	//
	// example:
	//
	// {"scaVul":"2023-09-04 09:37:21","identificationInfo":"2023-09-04 09:37:30","forbiddenPackageInfo":"2023-09-04 09:37:16","binary":"2023-09-04 09:37:25","baseline":"2023-09-04 09:37:19","sensitiveFile":"2023-09-04 09:38:34","vul":"2023-09-04 09:37:31","webshell":"2023-09-04 09:38:27","sensitiveInfo":"2023-09-04 09:37:16","script":"2023-09-04 09:39:44"}
	CostTimeInfo *string `json:"CostTimeInfo,omitempty" xml:"CostTimeInfo,omitempty"`
	// The digest of the image.
	//
	// example:
	//
	// 9e0dc29d872d2e386cc5c0c92b529a84e3acfade16f5cb1d054a2ee3c99****
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The image of the container.
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
	// The type of the image repository. Valid values:
	//
	// 	- **acr**
	//
	// 	- **harbor**
	//
	// 	- **quay**
	//
	// 	- **CI/CD**
	//
	// example:
	//
	// acr
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
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
	// The name of the namespace to which the image repository belongs.
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
	// The tag that is added to the image.
	//
	// example:
	//
	// v1.20-002-a2*****
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GoString() string {
	return s.String()
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetAppName() *string {
	return s.AppName
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetCostTimeInfo() *string {
	return s.CostTimeInfo
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetDigest() *string {
	return s.Digest
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetImage() *string {
	return s.Image
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetNodeInstanceId() *string {
	return s.NodeInstanceId
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetNodeIp() *string {
	return s.NodeIp
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetPod() *string {
	return s.Pod
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetRegistryType() *string {
	return s.RegistryType
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetRepoId() *string {
	return s.RepoId
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetRepoRegionId() *string {
	return s.RepoRegionId
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) GetTag() *string {
	return s.Tag
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetAppName(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.AppName = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetClusterId(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.ClusterId = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetClusterName(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.ClusterName = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetCostTimeInfo(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.CostTimeInfo = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetDigest(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.Digest = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetImage(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.Image = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetNodeInstanceId(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.NodeInstanceId = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetNodeIp(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.NodeIp = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetNodeName(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.NodeName = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetPod(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.Pod = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetRegionId(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetRegistryType(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.RegistryType = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetRepoId(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.RepoId = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetRepoName(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.RepoName = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetRepoNamespace(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetRepoRegionId(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.RepoRegionId = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) SetTag(v string) *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo {
	s.Tag = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyOnceTasksTaskImageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeOnceTaskLeafRecordPageResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 6
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOnceTaskLeafRecordPageResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnceTaskLeafRecordPageResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo) SetCount(v int32) *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo) SetPageSize(v int32) *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo) SetTotalCount(v int32) *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeOnceTaskLeafRecordPageResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
