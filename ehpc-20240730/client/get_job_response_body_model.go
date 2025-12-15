// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobInfo(v *GetJobResponseBodyJobInfo) *GetJobResponseBody
	GetJobInfo() *GetJobResponseBodyJobInfo
	SetRequestId(v string) *GetJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetJobResponseBody
	GetSuccess() *string
}

type GetJobResponseBody struct {
	// The job details.
	JobInfo *GetJobResponseBodyJobInfo `json:"JobInfo,omitempty" xml:"JobInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0****-1335-****-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) GetJobInfo() *GetJobResponseBodyJobInfo {
	return s.JobInfo
}

func (s *GetJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetJobResponseBody) SetJobInfo(v *GetJobResponseBodyJobInfo) *GetJobResponseBody {
	s.JobInfo = v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResponseBody) SetSuccess(v string) *GetJobResponseBody {
	s.Success = &v
	return s
}

func (s *GetJobResponseBody) Validate() error {
	if s.JobInfo != nil {
		if err := s.JobInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobResponseBodyJobInfo struct {
	// The parent job ID. If the return value is a non-empty string, the job is an array job.
	//
	// example:
	//
	// 1
	ArrayJobId *string `json:"ArrayJobId,omitempty" xml:"ArrayJobId,omitempty"`
	// The sub-job ID. This parameter is valid when the ArrayJobId parameter is a non-empty string.
	//
	// example:
	//
	// 3
	ArrayJobSubId *string `json:"ArrayJobSubId,omitempty" xml:"ArrayJobSubId,omitempty"`
	// The job queue. If the job is not in a queue, the output is empty.
	//
	// The format is X-Y:Z. X indicates the first index, Y indicates the final index, and Z indicates the step size. For example, 2-7:2 indicates three sub-jobs numbered 2, 4, and 6.
	//
	// example:
	//
	// 1-5:2
	ArrayRequest *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	// The command that is used to run the job.
	//
	// example:
	//
	// /home/huangsf/ehpc/job_meta.pbs
	CommandLine *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	// The time when the job was submitted.
	//
	// example:
	//
	// 2024-08-16T10:52:48
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error log file of the job.
	//
	// example:
	//
	// /home/xxx/STDIN.e1
	ErrorLog *string `json:"ErrorLog,omitempty" xml:"ErrorLog,omitempty"`
	// Additional information.
	//
	// example:
	//
	// {}
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 1.manager
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job name.
	//
	// example:
	//
	// testJob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The queue to which the job belongs.
	//
	// example:
	//
	// workq
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// The time when the job was last modified.
	//
	// example:
	//
	// 2024-08-16T10:52:48
	LastModifyTime *string `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The compute nodes that are used to run the job.
	//
	// example:
	//
	// compute000
	NodeList *string `json:"NodeList,omitempty" xml:"NodeList,omitempty"`
	// The standard output log file of the job.
	//
	// example:
	//
	// /home/xxx/STDIN.o1
	OutputLog *string `json:"OutputLog,omitempty" xml:"OutputLog,omitempty"`
	// The priority of the job.
	//
	// example:
	//
	// 0
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The resources that were requested when the job was submitted.
	Resources *GetJobResponseBodyJobInfoResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// The resources that are actually used by the job.
	ResourcesUsed *GetJobResponseBodyJobInfoResourcesUsed `json:"ResourcesUsed,omitempty" xml:"ResourcesUsed,omitempty" type:"Struct"`
	// The user to which the job belongs or that is used to submit the job. This user is a cluster-side user.
	//
	// example:
	//
	// testuser
	RunasUser *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	// The time when the job was started.
	//
	// example:
	//
	// 2024-08-16T10:52:48
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The job state.
	//
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The variables of the job.
	Variables []*GetJobResponseBodyJobInfoVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s GetJobResponseBodyJobInfo) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfo) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfo) GetArrayJobId() *string {
	return s.ArrayJobId
}

func (s *GetJobResponseBodyJobInfo) GetArrayJobSubId() *string {
	return s.ArrayJobSubId
}

func (s *GetJobResponseBodyJobInfo) GetArrayRequest() *string {
	return s.ArrayRequest
}

func (s *GetJobResponseBodyJobInfo) GetCommandLine() *string {
	return s.CommandLine
}

func (s *GetJobResponseBodyJobInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetJobResponseBodyJobInfo) GetErrorLog() *string {
	return s.ErrorLog
}

func (s *GetJobResponseBodyJobInfo) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *GetJobResponseBodyJobInfo) GetJobId() *string {
	return s.JobId
}

func (s *GetJobResponseBodyJobInfo) GetJobName() *string {
	return s.JobName
}

func (s *GetJobResponseBodyJobInfo) GetJobQueue() *string {
	return s.JobQueue
}

func (s *GetJobResponseBodyJobInfo) GetLastModifyTime() *string {
	return s.LastModifyTime
}

func (s *GetJobResponseBodyJobInfo) GetNodeList() *string {
	return s.NodeList
}

func (s *GetJobResponseBodyJobInfo) GetOutputLog() *string {
	return s.OutputLog
}

func (s *GetJobResponseBodyJobInfo) GetPriority() *string {
	return s.Priority
}

func (s *GetJobResponseBodyJobInfo) GetResources() *GetJobResponseBodyJobInfoResources {
	return s.Resources
}

func (s *GetJobResponseBodyJobInfo) GetResourcesUsed() *GetJobResponseBodyJobInfoResourcesUsed {
	return s.ResourcesUsed
}

func (s *GetJobResponseBodyJobInfo) GetRunasUser() *string {
	return s.RunasUser
}

func (s *GetJobResponseBodyJobInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *GetJobResponseBodyJobInfo) GetState() *string {
	return s.State
}

func (s *GetJobResponseBodyJobInfo) GetVariables() []*GetJobResponseBodyJobInfoVariables {
	return s.Variables
}

func (s *GetJobResponseBodyJobInfo) SetArrayJobId(v string) *GetJobResponseBodyJobInfo {
	s.ArrayJobId = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetArrayJobSubId(v string) *GetJobResponseBodyJobInfo {
	s.ArrayJobSubId = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetArrayRequest(v string) *GetJobResponseBodyJobInfo {
	s.ArrayRequest = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetCommandLine(v string) *GetJobResponseBodyJobInfo {
	s.CommandLine = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetCreateTime(v string) *GetJobResponseBodyJobInfo {
	s.CreateTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetErrorLog(v string) *GetJobResponseBodyJobInfo {
	s.ErrorLog = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetExtraInfo(v string) *GetJobResponseBodyJobInfo {
	s.ExtraInfo = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetJobId(v string) *GetJobResponseBodyJobInfo {
	s.JobId = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetJobName(v string) *GetJobResponseBodyJobInfo {
	s.JobName = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetJobQueue(v string) *GetJobResponseBodyJobInfo {
	s.JobQueue = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetLastModifyTime(v string) *GetJobResponseBodyJobInfo {
	s.LastModifyTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetNodeList(v string) *GetJobResponseBodyJobInfo {
	s.NodeList = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetOutputLog(v string) *GetJobResponseBodyJobInfo {
	s.OutputLog = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetPriority(v string) *GetJobResponseBodyJobInfo {
	s.Priority = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetResources(v *GetJobResponseBodyJobInfoResources) *GetJobResponseBodyJobInfo {
	s.Resources = v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetResourcesUsed(v *GetJobResponseBodyJobInfoResourcesUsed) *GetJobResponseBodyJobInfo {
	s.ResourcesUsed = v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetRunasUser(v string) *GetJobResponseBodyJobInfo {
	s.RunasUser = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetStartTime(v string) *GetJobResponseBodyJobInfo {
	s.StartTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetState(v string) *GetJobResponseBodyJobInfo {
	s.State = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetVariables(v []*GetJobResponseBodyJobInfoVariables) *GetJobResponseBodyJobInfo {
	s.Variables = v
	return s
}

func (s *GetJobResponseBodyJobInfo) Validate() error {
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			return err
		}
	}
	if s.ResourcesUsed != nil {
		if err := s.ResourcesUsed.Validate(); err != nil {
			return err
		}
	}
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetJobResponseBodyJobInfoResources struct {
	// The number of vCPUs used by the job on each node.
	//
	// example:
	//
	// 2
	Cores *string `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The number of GPUs used by the job on each node.
	//
	// example:
	//
	// 1
	Gpus *string `json:"Gpus,omitempty" xml:"Gpus,omitempty"`
	// The memory size used by the job on each node.
	//
	// example:
	//
	// 1gb
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The number of nodes that are used to run the job.
	//
	// example:
	//
	// 1
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s GetJobResponseBodyJobInfoResources) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoResources) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoResources) GetCores() *string {
	return s.Cores
}

func (s *GetJobResponseBodyJobInfoResources) GetGpus() *string {
	return s.Gpus
}

func (s *GetJobResponseBodyJobInfoResources) GetMemory() *string {
	return s.Memory
}

func (s *GetJobResponseBodyJobInfoResources) GetNodes() *string {
	return s.Nodes
}

func (s *GetJobResponseBodyJobInfoResources) SetCores(v string) *GetJobResponseBodyJobInfoResources {
	s.Cores = &v
	return s
}

func (s *GetJobResponseBodyJobInfoResources) SetGpus(v string) *GetJobResponseBodyJobInfoResources {
	s.Gpus = &v
	return s
}

func (s *GetJobResponseBodyJobInfoResources) SetMemory(v string) *GetJobResponseBodyJobInfoResources {
	s.Memory = &v
	return s
}

func (s *GetJobResponseBodyJobInfoResources) SetNodes(v string) *GetJobResponseBodyJobInfoResources {
	s.Nodes = &v
	return s
}

func (s *GetJobResponseBodyJobInfoResources) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoResourcesUsed struct {
	// The number of vCPUs used by the job on each node.
	//
	// example:
	//
	// 2
	Cores *string `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The memory size used by the job on each node.
	//
	// example:
	//
	// 512mb
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The number of nodes that are used to run the job.
	//
	// example:
	//
	// 2
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s GetJobResponseBodyJobInfoResourcesUsed) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoResourcesUsed) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoResourcesUsed) GetCores() *string {
	return s.Cores
}

func (s *GetJobResponseBodyJobInfoResourcesUsed) GetMemory() *string {
	return s.Memory
}

func (s *GetJobResponseBodyJobInfoResourcesUsed) GetNodes() *string {
	return s.Nodes
}

func (s *GetJobResponseBodyJobInfoResourcesUsed) SetCores(v string) *GetJobResponseBodyJobInfoResourcesUsed {
	s.Cores = &v
	return s
}

func (s *GetJobResponseBodyJobInfoResourcesUsed) SetMemory(v string) *GetJobResponseBodyJobInfoResourcesUsed {
	s.Memory = &v
	return s
}

func (s *GetJobResponseBodyJobInfoResourcesUsed) SetNodes(v string) *GetJobResponseBodyJobInfoResourcesUsed {
	s.Nodes = &v
	return s
}

func (s *GetJobResponseBodyJobInfoResourcesUsed) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoVariables struct {
	// The name of the environment variable.
	//
	// example:
	//
	// ProxyIP
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the environment variable.
	//
	// example:
	//
	// 10.x.x.x
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetJobResponseBodyJobInfoVariables) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoVariables) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoVariables) GetName() *string {
	return s.Name
}

func (s *GetJobResponseBodyJobInfoVariables) GetValue() *string {
	return s.Value
}

func (s *GetJobResponseBodyJobInfoVariables) SetName(v string) *GetJobResponseBodyJobInfoVariables {
	s.Name = &v
	return s
}

func (s *GetJobResponseBodyJobInfoVariables) SetValue(v string) *GetJobResponseBodyJobInfoVariables {
	s.Value = &v
	return s
}

func (s *GetJobResponseBodyJobInfoVariables) Validate() error {
	return dara.Validate(s)
}
