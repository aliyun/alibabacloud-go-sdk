// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobs(v []*ListJobsResponseBodyJobs) *ListJobsResponseBody
	GetJobs() []*ListJobsResponseBodyJobs
	SetPageNumber(v int32) *ListJobsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListJobsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListJobsResponseBody
	GetTotalCount() *int32
}

type ListJobsResponseBody struct {
	// The jobs.
	Jobs []*ListJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The page number. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EABFBD93-58BE-53F3-BBFE-8654BB2E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) GetJobs() []*ListJobsResponseBodyJobs {
	return s.Jobs
}

func (s *ListJobsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListJobsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListJobsResponseBody) SetJobs(v []*ListJobsResponseBodyJobs) *ListJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsResponseBody) SetPageNumber(v int32) *ListJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobsResponseBody) SetPageSize(v int32) *ListJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetSuccess(v bool) *ListJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalCount(v int32) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJobsResponseBody) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobsResponseBodyJobs struct {
	// The job name.
	//
	// example:
	//
	// testjob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The job configurations.
	JobSpec *ListJobsResponseBodyJobsJobSpec `json:"JobSpec,omitempty" xml:"JobSpec,omitempty" type:"Struct"`
}

func (s ListJobsResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobs) GetJobName() *string {
	return s.JobName
}

func (s *ListJobsResponseBodyJobs) GetJobSpec() *ListJobsResponseBodyJobsJobSpec {
	return s.JobSpec
}

func (s *ListJobsResponseBodyJobs) SetJobName(v string) *ListJobsResponseBodyJobs {
	s.JobName = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetJobSpec(v *ListJobsResponseBodyJobsJobSpec) *ListJobsResponseBodyJobs {
	s.JobSpec = v
	return s
}

func (s *ListJobsResponseBodyJobs) Validate() error {
	if s.JobSpec != nil {
		if err := s.JobSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobsResponseBodyJobsJobSpec struct {
	// The array job ID.
	//
	// example:
	//
	// 4
	ArrayJobId *string `json:"ArrayJobId,omitempty" xml:"ArrayJobId,omitempty"`
	// The ID of the job in the array.
	//
	// example:
	//
	// 1
	ArrayJobSubId *string `json:"ArrayJobSubId,omitempty" xml:"ArrayJobSubId,omitempty"`
	// The queue format of the job.
	//
	// 	- If the job is not in a queue, the output is empty.
	//
	// 	- The format is X-Y:Z. X indicates the first index, Y indicates the final index, and Z indicates the step size. For example, 2-7:2 indicates three sub-jobs numbered 2, 4, and 6.
	//
	// example:
	//
	// 1-5:2
	ArrayRequest *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	// The job description.
	//
	// example:
	//
	// jobDescription
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 12
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The queue name.
	//
	// example:
	//
	// comp
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// The time when the job was last updated.
	//
	// example:
	//
	// 1724123085
	LastModifyTime *string `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The compute nodes that were used to run the job.
	//
	// example:
	//
	// compute[002,005,003]
	NodeList *string `json:"NodeList,omitempty" xml:"NodeList,omitempty"`
	// The job priority. Valid values: 0 to 9. A larger value indicates a higher priority.
	//
	// example:
	//
	// 0
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The information about the resources required to run the job.
	Resources *ListJobsResponseBodyJobsJobSpecResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// Actual resource usage of the job program
	ResourcesActualOccupied *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied `json:"ResourcesActualOccupied,omitempty" xml:"ResourcesActualOccupied,omitempty" type:"Struct"`
	// The user that ran the job.
	//
	// example:
	//
	// testuser1
	RunasUser *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	// Job start time.
	//
	// example:
	//
	// 1724122486
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The job state. Valid values: (PBS cluster and Slurm cluster)
	//
	// 	- FINISHED/Completed
	//
	// 	- RUNNING/Running
	//
	// 	- QUEUED/Pending
	//
	// 	- FAILED/Failed
	//
	// example:
	//
	// Running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The error output path.
	//
	// example:
	//
	// ./Temp
	StderrPath *string `json:"StderrPath,omitempty" xml:"StderrPath,omitempty"`
	// The standard output path.
	//
	// example:
	//
	// ./Temp
	StdoutPath *string `json:"StdoutPath,omitempty" xml:"StdoutPath,omitempty"`
	// The time when the job was submitted.
	//
	// example:
	//
	// 1724122486
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// The variables of the job.
	//
	// example:
	//
	// {"PBS_O_SHELL":"/bin/bash", 	"PBS_O_HOST":"manager", 	"PBS_O_SYSTEM":"Linux", 	"PBS_O_LANG":"en_US.UTF-8", 	"PBS_O_QUEUE":"workq"}
	Variables *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s ListJobsResponseBodyJobsJobSpec) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobsJobSpec) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsJobSpec) GetArrayJobId() *string {
	return s.ArrayJobId
}

func (s *ListJobsResponseBodyJobsJobSpec) GetArrayJobSubId() *string {
	return s.ArrayJobSubId
}

func (s *ListJobsResponseBodyJobsJobSpec) GetArrayRequest() *string {
	return s.ArrayRequest
}

func (s *ListJobsResponseBodyJobsJobSpec) GetComment() *string {
	return s.Comment
}

func (s *ListJobsResponseBodyJobsJobSpec) GetId() *string {
	return s.Id
}

func (s *ListJobsResponseBodyJobsJobSpec) GetJobQueue() *string {
	return s.JobQueue
}

func (s *ListJobsResponseBodyJobsJobSpec) GetLastModifyTime() *string {
	return s.LastModifyTime
}

func (s *ListJobsResponseBodyJobsJobSpec) GetNodeList() *string {
	return s.NodeList
}

func (s *ListJobsResponseBodyJobsJobSpec) GetPriority() *string {
	return s.Priority
}

func (s *ListJobsResponseBodyJobsJobSpec) GetResources() *ListJobsResponseBodyJobsJobSpecResources {
	return s.Resources
}

func (s *ListJobsResponseBodyJobsJobSpec) GetResourcesActualOccupied() *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied {
	return s.ResourcesActualOccupied
}

func (s *ListJobsResponseBodyJobsJobSpec) GetRunasUser() *string {
	return s.RunasUser
}

func (s *ListJobsResponseBodyJobsJobSpec) GetStartTime() *string {
	return s.StartTime
}

func (s *ListJobsResponseBodyJobsJobSpec) GetState() *string {
	return s.State
}

func (s *ListJobsResponseBodyJobsJobSpec) GetStderrPath() *string {
	return s.StderrPath
}

func (s *ListJobsResponseBodyJobsJobSpec) GetStdoutPath() *string {
	return s.StdoutPath
}

func (s *ListJobsResponseBodyJobsJobSpec) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *ListJobsResponseBodyJobsJobSpec) GetVariables() *string {
	return s.Variables
}

func (s *ListJobsResponseBodyJobsJobSpec) SetArrayJobId(v string) *ListJobsResponseBodyJobsJobSpec {
	s.ArrayJobId = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetArrayJobSubId(v string) *ListJobsResponseBodyJobsJobSpec {
	s.ArrayJobSubId = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetArrayRequest(v string) *ListJobsResponseBodyJobsJobSpec {
	s.ArrayRequest = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetComment(v string) *ListJobsResponseBodyJobsJobSpec {
	s.Comment = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetId(v string) *ListJobsResponseBodyJobsJobSpec {
	s.Id = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetJobQueue(v string) *ListJobsResponseBodyJobsJobSpec {
	s.JobQueue = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetLastModifyTime(v string) *ListJobsResponseBodyJobsJobSpec {
	s.LastModifyTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetNodeList(v string) *ListJobsResponseBodyJobsJobSpec {
	s.NodeList = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetPriority(v string) *ListJobsResponseBodyJobsJobSpec {
	s.Priority = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetResources(v *ListJobsResponseBodyJobsJobSpecResources) *ListJobsResponseBodyJobsJobSpec {
	s.Resources = v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetResourcesActualOccupied(v *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) *ListJobsResponseBodyJobsJobSpec {
	s.ResourcesActualOccupied = v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetRunasUser(v string) *ListJobsResponseBodyJobsJobSpec {
	s.RunasUser = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetStartTime(v string) *ListJobsResponseBodyJobsJobSpec {
	s.StartTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetState(v string) *ListJobsResponseBodyJobsJobSpec {
	s.State = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetStderrPath(v string) *ListJobsResponseBodyJobsJobSpec {
	s.StderrPath = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetStdoutPath(v string) *ListJobsResponseBodyJobsJobSpec {
	s.StdoutPath = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetSubmitTime(v string) *ListJobsResponseBodyJobsJobSpec {
	s.SubmitTime = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) SetVariables(v string) *ListJobsResponseBodyJobsJobSpec {
	s.Variables = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpec) Validate() error {
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			return err
		}
	}
	if s.ResourcesActualOccupied != nil {
		if err := s.ResourcesActualOccupied.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobsResponseBodyJobsJobSpecResources struct {
	// The number of vCPUs that were used to run the job.
	//
	// example:
	//
	// 6
	Cores *string `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The number of GPUs that were used to run the job.
	//
	// example:
	//
	// 0
	Gpus *string `json:"Gpus,omitempty" xml:"Gpus,omitempty"`
	// The size of memory that was used to run the job.
	//
	// example:
	//
	// 1536MB
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The number of compute nodes that were used to run the job.
	//
	// example:
	//
	// 3
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s ListJobsResponseBodyJobsJobSpecResources) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobsJobSpecResources) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsJobSpecResources) GetCores() *string {
	return s.Cores
}

func (s *ListJobsResponseBodyJobsJobSpecResources) GetGpus() *string {
	return s.Gpus
}

func (s *ListJobsResponseBodyJobsJobSpecResources) GetMemory() *string {
	return s.Memory
}

func (s *ListJobsResponseBodyJobsJobSpecResources) GetNodes() *string {
	return s.Nodes
}

func (s *ListJobsResponseBodyJobsJobSpecResources) SetCores(v string) *ListJobsResponseBodyJobsJobSpecResources {
	s.Cores = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpecResources) SetGpus(v string) *ListJobsResponseBodyJobsJobSpecResources {
	s.Gpus = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpecResources) SetMemory(v string) *ListJobsResponseBodyJobsJobSpecResources {
	s.Memory = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpecResources) SetNodes(v string) *ListJobsResponseBodyJobsJobSpecResources {
	s.Nodes = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpecResources) Validate() error {
	return dara.Validate(s)
}

type ListJobsResponseBodyJobsJobSpecResourcesActualOccupied struct {
	// Number of CPU cores.
	//
	// example:
	//
	// 4
	Cores *string `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// Number of CPUs
	//
	// example:
	//
	// 0
	Gpus *string `json:"Gpus,omitempty" xml:"Gpus,omitempty"`
	// Number of memory.
	//
	// example:
	//
	// 982MB
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Number of compute nodes.
	//
	// example:
	//
	// 2
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) GetCores() *string {
	return s.Cores
}

func (s *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) GetGpus() *string {
	return s.Gpus
}

func (s *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) GetMemory() *string {
	return s.Memory
}

func (s *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) GetNodes() *string {
	return s.Nodes
}

func (s *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) SetCores(v string) *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied {
	s.Cores = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) SetGpus(v string) *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied {
	s.Gpus = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) SetMemory(v string) *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied {
	s.Memory = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) SetNodes(v string) *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied {
	s.Nodes = &v
	return s
}

func (s *ListJobsResponseBodyJobsJobSpecResourcesActualOccupied) Validate() error {
	return dara.Validate(s)
}
