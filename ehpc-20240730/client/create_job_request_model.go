// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateJobRequest
	GetClusterId() *string
	SetJobName(v string) *CreateJobRequest
	GetJobName() *string
	SetJobSpec(v *CreateJobRequestJobSpec) *CreateJobRequest
	GetJobSpec() *CreateJobRequestJobSpec
}

type CreateJobRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The job name.
	//
	// example:
	//
	// TestJob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The job configurations.
	JobSpec *CreateJobRequestJobSpec `json:"JobSpec,omitempty" xml:"JobSpec,omitempty" type:"Struct"`
}

func (s CreateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *CreateJobRequest) GetJobSpec() *CreateJobRequestJobSpec {
	return s.JobSpec
}

func (s *CreateJobRequest) SetClusterId(v string) *CreateJobRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateJobRequest) SetJobName(v string) *CreateJobRequest {
	s.JobName = &v
	return s
}

func (s *CreateJobRequest) SetJobSpec(v *CreateJobRequestJobSpec) *CreateJobRequest {
	s.JobSpec = v
	return s
}

func (s *CreateJobRequest) Validate() error {
	if s.JobSpec != nil {
		if err := s.JobSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateJobRequestJobSpec struct {
	// The jobs in the queue.
	//
	// Format: X-Y:Z. X is the minimum index value. Y is the maximum index value. Z is the step size. For example, 2-7:2 indicates that three jobs need to be run and their index values are 2, 4, and 6.
	//
	// example:
	//
	// 1-5:2
	ArrayRequest *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	// The command or script that is used to run the job. If you want to use a command, you must specify the full path of the command, for example, /bin/ping.
	//
	// If you want to use a script, you must make sure that you have the execution permissions on it. By default, the user root directory ~/ is used as the default script path on the cluster side. If your script is not in that directory, you must specify the full path in this parameter, such as /home/xxx/job.sh Note that in this mode, if requirements on resources such as CPU and memory are specified in the script, the job will be run based on the resource requirements specified in the script. In this case, do not specify resource requirements in the Resource parameter. Otherwise, the job may fail to run.
	//
	// If you want to run the job directly by using the CLI, you must specify the absolute path of the command and add two hyphens and a space (-- ) before the path, such as -- /bin/ping -c 10 localhost.
	//
	// This parameter is required.
	//
	// example:
	//
	// /home/xxx/test.job
	CommandLine *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	// The queue to which the job belongs.
	//
	// example:
	//
	// comp
	JobQueue *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	// The post-processing command of the job.
	//
	// example:
	//
	// /bin/sleep 10
	PostCmdLine *string `json:"PostCmdLine,omitempty" xml:"PostCmdLine,omitempty"`
	// The job priority.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The resource configurations of the job.
	Resources *CreateJobRequestJobSpecResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// The cluster-side user as which you want to submit the job.
	//
	// example:
	//
	// testuser
	RunasUser *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	// The password of the user specified by the RunasUser parameter.
	//
	// example:
	//
	// xxx
	RunasUserPassword *string `json:"RunasUserPassword,omitempty" xml:"RunasUserPassword,omitempty"`
	// The path of the standard error output file of the job. You need to specify the full path.
	//
	// example:
	//
	// /home/xxx/job.err
	StderrPath *string `json:"StderrPath,omitempty" xml:"StderrPath,omitempty"`
	// The path of the standard output file of the job. You need to specify the full path.
	//
	// example:
	//
	// /home/xxx/job.out
	StdoutPath *string `json:"StdoutPath,omitempty" xml:"StdoutPath,omitempty"`
	// The environment variables of the job. The value is a string in the JSON array format. Each array member is a JSON object that contains two members: Name and Value. Name indicates the name of the environment variable, and Value indicates the value of the environment variable.
	//
	// example:
	//
	// [{"Name":"x", "Value":"y"}]
	Variables *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
	// The maximum duration for which the job can be run. Format: `hour: minute: second`. For example, `01:00:00` indicates 1 hour.
	//
	// example:
	//
	// 360:48:50
	WallTime *string `json:"WallTime,omitempty" xml:"WallTime,omitempty"`
}

func (s CreateJobRequestJobSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestJobSpec) GoString() string {
	return s.String()
}

func (s *CreateJobRequestJobSpec) GetArrayRequest() *string {
	return s.ArrayRequest
}

func (s *CreateJobRequestJobSpec) GetCommandLine() *string {
	return s.CommandLine
}

func (s *CreateJobRequestJobSpec) GetJobQueue() *string {
	return s.JobQueue
}

func (s *CreateJobRequestJobSpec) GetPostCmdLine() *string {
	return s.PostCmdLine
}

func (s *CreateJobRequestJobSpec) GetPriority() *string {
	return s.Priority
}

func (s *CreateJobRequestJobSpec) GetResources() *CreateJobRequestJobSpecResources {
	return s.Resources
}

func (s *CreateJobRequestJobSpec) GetRunasUser() *string {
	return s.RunasUser
}

func (s *CreateJobRequestJobSpec) GetRunasUserPassword() *string {
	return s.RunasUserPassword
}

func (s *CreateJobRequestJobSpec) GetStderrPath() *string {
	return s.StderrPath
}

func (s *CreateJobRequestJobSpec) GetStdoutPath() *string {
	return s.StdoutPath
}

func (s *CreateJobRequestJobSpec) GetVariables() *string {
	return s.Variables
}

func (s *CreateJobRequestJobSpec) GetWallTime() *string {
	return s.WallTime
}

func (s *CreateJobRequestJobSpec) SetArrayRequest(v string) *CreateJobRequestJobSpec {
	s.ArrayRequest = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetCommandLine(v string) *CreateJobRequestJobSpec {
	s.CommandLine = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetJobQueue(v string) *CreateJobRequestJobSpec {
	s.JobQueue = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetPostCmdLine(v string) *CreateJobRequestJobSpec {
	s.PostCmdLine = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetPriority(v string) *CreateJobRequestJobSpec {
	s.Priority = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetResources(v *CreateJobRequestJobSpecResources) *CreateJobRequestJobSpec {
	s.Resources = v
	return s
}

func (s *CreateJobRequestJobSpec) SetRunasUser(v string) *CreateJobRequestJobSpec {
	s.RunasUser = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetRunasUserPassword(v string) *CreateJobRequestJobSpec {
	s.RunasUserPassword = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetStderrPath(v string) *CreateJobRequestJobSpec {
	s.StderrPath = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetStdoutPath(v string) *CreateJobRequestJobSpec {
	s.StdoutPath = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetVariables(v string) *CreateJobRequestJobSpec {
	s.Variables = &v
	return s
}

func (s *CreateJobRequestJobSpec) SetWallTime(v string) *CreateJobRequestJobSpec {
	s.WallTime = &v
	return s
}

func (s *CreateJobRequestJobSpec) Validate() error {
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateJobRequestJobSpecResources struct {
	// The number of vCPUs to be allocated to each compute node.
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The number of GPUs to be allocated to each compute node.
	//
	// example:
	//
	// 1
	Gpus *int32 `json:"Gpus,omitempty" xml:"Gpus,omitempty"`
	// The memory size to be allocated to each compute node. The memory size is in string format. Unit: MB or GB.
	//
	// example:
	//
	// 4gb
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The number of compute nodes to be allocated to the job.
	//
	// example:
	//
	// 2
	Nodes *int32 `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s CreateJobRequestJobSpecResources) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestJobSpecResources) GoString() string {
	return s.String()
}

func (s *CreateJobRequestJobSpecResources) GetCores() *int32 {
	return s.Cores
}

func (s *CreateJobRequestJobSpecResources) GetGpus() *int32 {
	return s.Gpus
}

func (s *CreateJobRequestJobSpecResources) GetMemory() *string {
	return s.Memory
}

func (s *CreateJobRequestJobSpecResources) GetNodes() *int32 {
	return s.Nodes
}

func (s *CreateJobRequestJobSpecResources) SetCores(v int32) *CreateJobRequestJobSpecResources {
	s.Cores = &v
	return s
}

func (s *CreateJobRequestJobSpecResources) SetGpus(v int32) *CreateJobRequestJobSpecResources {
	s.Gpus = &v
	return s
}

func (s *CreateJobRequestJobSpecResources) SetMemory(v string) *CreateJobRequestJobSpecResources {
	s.Memory = &v
	return s
}

func (s *CreateJobRequestJobSpecResources) SetNodes(v int32) *CreateJobRequestJobSpecResources {
	s.Nodes = &v
	return s
}

func (s *CreateJobRequestJobSpecResources) Validate() error {
	return dara.Validate(s)
}
