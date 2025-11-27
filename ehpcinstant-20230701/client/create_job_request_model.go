// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDependencyPolicy(v *CreateJobRequestDependencyPolicy) *CreateJobRequest
	GetDependencyPolicy() *CreateJobRequestDependencyPolicy
	SetDeploymentPolicy(v *CreateJobRequestDeploymentPolicy) *CreateJobRequest
	GetDeploymentPolicy() *CreateJobRequestDeploymentPolicy
	SetJobDescription(v string) *CreateJobRequest
	GetJobDescription() *string
	SetJobName(v string) *CreateJobRequest
	GetJobName() *string
	SetJobScheduler(v string) *CreateJobRequest
	GetJobScheduler() *string
	SetSecurityPolicy(v *CreateJobRequestSecurityPolicy) *CreateJobRequest
	GetSecurityPolicy() *CreateJobRequestSecurityPolicy
	SetTasks(v []*CreateJobRequestTasks) *CreateJobRequest
	GetTasks() []*CreateJobRequestTasks
}

type CreateJobRequest struct {
	// Dependency policy.
	DependencyPolicy *CreateJobRequestDependencyPolicy `json:"DependencyPolicy,omitempty" xml:"DependencyPolicy,omitempty" type:"Struct"`
	// The resource deployment policy.
	DeploymentPolicy *CreateJobRequestDeploymentPolicy `json:"DeploymentPolicy,omitempty" xml:"DeploymentPolicy,omitempty" type:"Struct"`
	// The description of the job.
	//
	// example:
	//
	// Demo
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// The job name. The name must be 2 to 64 characters in length and can contain letters, digits, and Chinese characters. It can contain hyphens (-) and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// testjob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The type of the job scheduler.
	//
	// 	- HPC
	//
	// 	- K8S
	//
	// Default value: HPC
	//
	// example:
	//
	// HPC
	JobScheduler *string `json:"JobScheduler,omitempty" xml:"JobScheduler,omitempty"`
	// The security policy.
	SecurityPolicy *CreateJobRequestSecurityPolicy `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty" type:"Struct"`
	// The list of tasks. Only one task is supported.
	//
	// This parameter is required.
	Tasks []*CreateJobRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s CreateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) GetDependencyPolicy() *CreateJobRequestDependencyPolicy {
	return s.DependencyPolicy
}

func (s *CreateJobRequest) GetDeploymentPolicy() *CreateJobRequestDeploymentPolicy {
	return s.DeploymentPolicy
}

func (s *CreateJobRequest) GetJobDescription() *string {
	return s.JobDescription
}

func (s *CreateJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *CreateJobRequest) GetJobScheduler() *string {
	return s.JobScheduler
}

func (s *CreateJobRequest) GetSecurityPolicy() *CreateJobRequestSecurityPolicy {
	return s.SecurityPolicy
}

func (s *CreateJobRequest) GetTasks() []*CreateJobRequestTasks {
	return s.Tasks
}

func (s *CreateJobRequest) SetDependencyPolicy(v *CreateJobRequestDependencyPolicy) *CreateJobRequest {
	s.DependencyPolicy = v
	return s
}

func (s *CreateJobRequest) SetDeploymentPolicy(v *CreateJobRequestDeploymentPolicy) *CreateJobRequest {
	s.DeploymentPolicy = v
	return s
}

func (s *CreateJobRequest) SetJobDescription(v string) *CreateJobRequest {
	s.JobDescription = &v
	return s
}

func (s *CreateJobRequest) SetJobName(v string) *CreateJobRequest {
	s.JobName = &v
	return s
}

func (s *CreateJobRequest) SetJobScheduler(v string) *CreateJobRequest {
	s.JobScheduler = &v
	return s
}

func (s *CreateJobRequest) SetSecurityPolicy(v *CreateJobRequestSecurityPolicy) *CreateJobRequest {
	s.SecurityPolicy = v
	return s
}

func (s *CreateJobRequest) SetTasks(v []*CreateJobRequestTasks) *CreateJobRequest {
	s.Tasks = v
	return s
}

func (s *CreateJobRequest) Validate() error {
	if s.DependencyPolicy != nil {
		if err := s.DependencyPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.DeploymentPolicy != nil {
		if err := s.DeploymentPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityPolicy != nil {
		if err := s.SecurityPolicy.Validate(); err != nil {
			return err
		}
	}
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

type CreateJobRequestDependencyPolicy struct {
	// The job dependency. A maximum of 10 groups.
	JobDependency []*CreateJobRequestDependencyPolicyJobDependency `json:"JobDependency,omitempty" xml:"JobDependency,omitempty" type:"Repeated"`
}

func (s CreateJobRequestDependencyPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestDependencyPolicy) GoString() string {
	return s.String()
}

func (s *CreateJobRequestDependencyPolicy) GetJobDependency() []*CreateJobRequestDependencyPolicyJobDependency {
	return s.JobDependency
}

func (s *CreateJobRequestDependencyPolicy) SetJobDependency(v []*CreateJobRequestDependencyPolicyJobDependency) *CreateJobRequestDependencyPolicy {
	s.JobDependency = v
	return s
}

func (s *CreateJobRequestDependencyPolicy) Validate() error {
	if s.JobDependency != nil {
		for _, item := range s.JobDependency {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateJobRequestDependencyPolicyJobDependency struct {
	// The ID of the job.
	//
	// This parameter is required.
	//
	// example:
	//
	// job-bjxxxxxxxxxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The type of the dependency. Valid values:
	//
	// 	- AfterSucceeded: **All subtasks*	- of the dependent job or array job succeed. The exit code is 0.
	//
	// 	- AfterFailed: **All subtasks*	- of the dependent job or array job fail. The exit code is not 0.
	//
	// 	- AfterAny: The dependent job completes (succeeds or fails).
	//
	// 	- AfterCorresponding: The subtask corresponding to the dependent array job succeeds. The exit code is 0.
	//
	// Default value: AfterSucceeded.
	//
	// example:
	//
	// AfterSucceeded
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateJobRequestDependencyPolicyJobDependency) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestDependencyPolicyJobDependency) GoString() string {
	return s.String()
}

func (s *CreateJobRequestDependencyPolicyJobDependency) GetJobId() *string {
	return s.JobId
}

func (s *CreateJobRequestDependencyPolicyJobDependency) GetType() *string {
	return s.Type
}

func (s *CreateJobRequestDependencyPolicyJobDependency) SetJobId(v string) *CreateJobRequestDependencyPolicyJobDependency {
	s.JobId = &v
	return s
}

func (s *CreateJobRequestDependencyPolicyJobDependency) SetType(v string) *CreateJobRequestDependencyPolicyJobDependency {
	s.Type = &v
	return s
}

func (s *CreateJobRequestDependencyPolicyJobDependency) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestDeploymentPolicy struct {
	// The resource type,
	//
	// 	- Standard
	//
	// 	- Dedicated: You must enable a whitelist for use.
	//
	// 	- Economic: You must enable a whitelist for use.
	//
	// example:
	//
	// Dedicated
	AllocationSpec *string `json:"AllocationSpec,omitempty" xml:"AllocationSpec,omitempty"`
	// The computing power level. This value is valid only when the resource type is Economic. The following disk categories are supported:
	//
	// 	- General
	//
	// 	- Performance
	//
	// Default value: General.
	//
	// example:
	//
	// General
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The network configuration information.
	Network *CreateJobRequestDeploymentPolicyNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The resource pool of the job.
	//
	// example:
	//
	// compute
	Pool *string `json:"Pool,omitempty" xml:"Pool,omitempty"`
	// The priorities of the jobs. A larger value indicates a higher job scheduling priority. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The tag information of the job. A maximum of 20 groups.
	Tag []*CreateJobRequestDeploymentPolicyTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateJobRequestDeploymentPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestDeploymentPolicy) GoString() string {
	return s.String()
}

func (s *CreateJobRequestDeploymentPolicy) GetAllocationSpec() *string {
	return s.AllocationSpec
}

func (s *CreateJobRequestDeploymentPolicy) GetLevel() *string {
	return s.Level
}

func (s *CreateJobRequestDeploymentPolicy) GetNetwork() *CreateJobRequestDeploymentPolicyNetwork {
	return s.Network
}

func (s *CreateJobRequestDeploymentPolicy) GetPool() *string {
	return s.Pool
}

func (s *CreateJobRequestDeploymentPolicy) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateJobRequestDeploymentPolicy) GetTag() []*CreateJobRequestDeploymentPolicyTag {
	return s.Tag
}

func (s *CreateJobRequestDeploymentPolicy) SetAllocationSpec(v string) *CreateJobRequestDeploymentPolicy {
	s.AllocationSpec = &v
	return s
}

func (s *CreateJobRequestDeploymentPolicy) SetLevel(v string) *CreateJobRequestDeploymentPolicy {
	s.Level = &v
	return s
}

func (s *CreateJobRequestDeploymentPolicy) SetNetwork(v *CreateJobRequestDeploymentPolicyNetwork) *CreateJobRequestDeploymentPolicy {
	s.Network = v
	return s
}

func (s *CreateJobRequestDeploymentPolicy) SetPool(v string) *CreateJobRequestDeploymentPolicy {
	s.Pool = &v
	return s
}

func (s *CreateJobRequestDeploymentPolicy) SetPriority(v int32) *CreateJobRequestDeploymentPolicy {
	s.Priority = &v
	return s
}

func (s *CreateJobRequestDeploymentPolicy) SetTag(v []*CreateJobRequestDeploymentPolicyTag) *CreateJobRequestDeploymentPolicy {
	s.Tag = v
	return s
}

func (s *CreateJobRequestDeploymentPolicy) Validate() error {
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateJobRequestDeploymentPolicyNetwork struct {
	// Whether the job creates a public IP address.
	//
	// 	- true: creates a public IP address.
	//
	// 	- false: does not create a public IP address.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	EnableExternalIpAddress *bool `json:"EnableExternalIpAddress,omitempty" xml:"EnableExternalIpAddress,omitempty"`
	// The VSwitch array.
	Vswitch []*string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty" type:"Repeated"`
}

func (s CreateJobRequestDeploymentPolicyNetwork) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestDeploymentPolicyNetwork) GoString() string {
	return s.String()
}

func (s *CreateJobRequestDeploymentPolicyNetwork) GetEnableExternalIpAddress() *bool {
	return s.EnableExternalIpAddress
}

func (s *CreateJobRequestDeploymentPolicyNetwork) GetVswitch() []*string {
	return s.Vswitch
}

func (s *CreateJobRequestDeploymentPolicyNetwork) SetEnableExternalIpAddress(v bool) *CreateJobRequestDeploymentPolicyNetwork {
	s.EnableExternalIpAddress = &v
	return s
}

func (s *CreateJobRequestDeploymentPolicyNetwork) SetVswitch(v []*string) *CreateJobRequestDeploymentPolicyNetwork {
	s.Vswitch = v
	return s
}

func (s *CreateJobRequestDeploymentPolicyNetwork) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestDeploymentPolicyTag struct {
	// The key of the job tag. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the job tag. You can specify empty strings as tag values. The tag value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateJobRequestDeploymentPolicyTag) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestDeploymentPolicyTag) GoString() string {
	return s.String()
}

func (s *CreateJobRequestDeploymentPolicyTag) GetKey() *string {
	return s.Key
}

func (s *CreateJobRequestDeploymentPolicyTag) GetValue() *string {
	return s.Value
}

func (s *CreateJobRequestDeploymentPolicyTag) SetKey(v string) *CreateJobRequestDeploymentPolicyTag {
	s.Key = &v
	return s
}

func (s *CreateJobRequestDeploymentPolicyTag) SetValue(v string) *CreateJobRequestDeploymentPolicyTag {
	s.Value = &v
	return s
}

func (s *CreateJobRequestDeploymentPolicyTag) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestSecurityPolicy struct {
	// The security group ID.
	SecurityGroup *CreateJobRequestSecurityPolicySecurityGroup `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty" type:"Struct"`
}

func (s CreateJobRequestSecurityPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestSecurityPolicy) GoString() string {
	return s.String()
}

func (s *CreateJobRequestSecurityPolicy) GetSecurityGroup() *CreateJobRequestSecurityPolicySecurityGroup {
	return s.SecurityGroup
}

func (s *CreateJobRequestSecurityPolicy) SetSecurityGroup(v *CreateJobRequestSecurityPolicySecurityGroup) *CreateJobRequestSecurityPolicy {
	s.SecurityGroup = v
	return s
}

func (s *CreateJobRequestSecurityPolicy) Validate() error {
	if s.SecurityGroup != nil {
		if err := s.SecurityGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateJobRequestSecurityPolicySecurityGroup struct {
	// The array of security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s CreateJobRequestSecurityPolicySecurityGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestSecurityPolicySecurityGroup) GoString() string {
	return s.String()
}

func (s *CreateJobRequestSecurityPolicySecurityGroup) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateJobRequestSecurityPolicySecurityGroup) SetSecurityGroupIds(v []*string) *CreateJobRequestSecurityPolicySecurityGroup {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateJobRequestSecurityPolicySecurityGroup) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestTasks struct {
	// The task execution policy.
	ExecutorPolicy *CreateJobRequestTasksExecutorPolicy `json:"ExecutorPolicy,omitempty" xml:"ExecutorPolicy,omitempty" type:"Struct"`
	// The job name. It must be 2 to 32 characters in length and can contain letters, digits, and Chinese characters. It can contain hyphens (-) and underscores (_).
	//
	// example:
	//
	// task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The details of the task specification.
	TaskSpec *CreateJobRequestTasksTaskSpec `json:"TaskSpec,omitempty" xml:"TaskSpec,omitempty" type:"Struct"`
	// Indicate whether the job is a long-running job.
	//
	// 	- true: background service the job.
	//
	// 	- false: batch jobs.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	TaskSustainable *bool `json:"TaskSustainable,omitempty" xml:"TaskSustainable,omitempty"`
}

func (s CreateJobRequestTasks) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestTasks) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasks) GetExecutorPolicy() *CreateJobRequestTasksExecutorPolicy {
	return s.ExecutorPolicy
}

func (s *CreateJobRequestTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateJobRequestTasks) GetTaskSpec() *CreateJobRequestTasksTaskSpec {
	return s.TaskSpec
}

func (s *CreateJobRequestTasks) GetTaskSustainable() *bool {
	return s.TaskSustainable
}

func (s *CreateJobRequestTasks) SetExecutorPolicy(v *CreateJobRequestTasksExecutorPolicy) *CreateJobRequestTasks {
	s.ExecutorPolicy = v
	return s
}

func (s *CreateJobRequestTasks) SetTaskName(v string) *CreateJobRequestTasks {
	s.TaskName = &v
	return s
}

func (s *CreateJobRequestTasks) SetTaskSpec(v *CreateJobRequestTasksTaskSpec) *CreateJobRequestTasks {
	s.TaskSpec = v
	return s
}

func (s *CreateJobRequestTasks) SetTaskSustainable(v bool) *CreateJobRequestTasks {
	s.TaskSustainable = &v
	return s
}

func (s *CreateJobRequestTasks) Validate() error {
	if s.ExecutorPolicy != nil {
		if err := s.ExecutorPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.TaskSpec != nil {
		if err := s.TaskSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateJobRequestTasksExecutorPolicy struct {
	// The details of the array job. The index value of the sub-job is passed to the running environment through environment variables to support user business program reference. Environment variables include:
	//
	// 	- EHPC_JOB_NAME: the name of the job. This parameter corresponds to the JobName parameter.
	//
	// 	- EHPC_JOB_ID: The ID of the job.
	//
	// 	- EHPC_TASK_NAME: the name of the task. This parameter corresponds to the TaskName parameter.
	//
	// 	- EHPC_EXECUTOR_ID: The ID of the execution unit.
	//
	// 	- EHPC_ARRAY_TASK_ID: the sub-job index value.
	//
	// 	- EHPC_ARRAY_TASK_COUNT: the total number of sub-jobs.
	//
	// 	- EHPC_ARRAY_TASK_MAX: the maximum sub-job index, which corresponds to the IndexStart parameter.
	//
	// 	- EHPC_ARRAY_TASK_MIN: the minimum value of the sub-job index, which corresponds to the IndexEnd parameter.
	//
	// 	- EHPC_ARRAY_TASK_STEP: the index step size of the sub-job, which corresponds to the IndexStep parameter.
	ArraySpec *CreateJobRequestTasksExecutorPolicyArraySpec `json:"ArraySpec,omitempty" xml:"ArraySpec,omitempty" type:"Struct"`
	// The maximum number of nodes to run the job.
	//
	// > Follow the calculation formula: `MaxCount = (IndexEnd - IndexStart) / IndexStep +1`
	//
	// example:
	//
	// 1
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
}

func (s CreateJobRequestTasksExecutorPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestTasksExecutorPolicy) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksExecutorPolicy) GetArraySpec() *CreateJobRequestTasksExecutorPolicyArraySpec {
	return s.ArraySpec
}

func (s *CreateJobRequestTasksExecutorPolicy) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *CreateJobRequestTasksExecutorPolicy) SetArraySpec(v *CreateJobRequestTasksExecutorPolicyArraySpec) *CreateJobRequestTasksExecutorPolicy {
	s.ArraySpec = v
	return s
}

func (s *CreateJobRequestTasksExecutorPolicy) SetMaxCount(v int32) *CreateJobRequestTasksExecutorPolicy {
	s.MaxCount = &v
	return s
}

func (s *CreateJobRequestTasksExecutorPolicy) Validate() error {
	if s.ArraySpec != nil {
		if err := s.ArraySpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateJobRequestTasksExecutorPolicyArraySpec struct {
	// The end value of the array job index. Valid values: 0 to 4999. The value must be greater than or equal to the value of IndexStart.
	//
	// example:
	//
	// 9
	IndexEnd *int32 `json:"IndexEnd,omitempty" xml:"IndexEnd,omitempty"`
	// The starting value of the array job index. Valid values: 0 to 4999.
	//
	// example:
	//
	// 0
	IndexStart *int32 `json:"IndexStart,omitempty" xml:"IndexStart,omitempty"`
	// The interval of the array job index.
	//
	// > If the array job property is IndexStart=1,IndexEnd=5, and IndexStep=2, the array job contains three sub-jobs. The index values of the sub-jobs are 1,3, and 5. You can access the sub-jobs by using environment variables.
	//
	// example:
	//
	// 1
	IndexStep *int32 `json:"IndexStep,omitempty" xml:"IndexStep,omitempty"`
}

func (s CreateJobRequestTasksExecutorPolicyArraySpec) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestTasksExecutorPolicyArraySpec) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksExecutorPolicyArraySpec) GetIndexEnd() *int32 {
	return s.IndexEnd
}

func (s *CreateJobRequestTasksExecutorPolicyArraySpec) GetIndexStart() *int32 {
	return s.IndexStart
}

func (s *CreateJobRequestTasksExecutorPolicyArraySpec) GetIndexStep() *int32 {
	return s.IndexStep
}

func (s *CreateJobRequestTasksExecutorPolicyArraySpec) SetIndexEnd(v int32) *CreateJobRequestTasksExecutorPolicyArraySpec {
	s.IndexEnd = &v
	return s
}

func (s *CreateJobRequestTasksExecutorPolicyArraySpec) SetIndexStart(v int32) *CreateJobRequestTasksExecutorPolicyArraySpec {
	s.IndexStart = &v
	return s
}

func (s *CreateJobRequestTasksExecutorPolicyArraySpec) SetIndexStep(v int32) *CreateJobRequestTasksExecutorPolicyArraySpec {
	s.IndexStep = &v
	return s
}

func (s *CreateJobRequestTasksExecutorPolicyArraySpec) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestTasksTaskSpec struct {
	// The resource information of the running environment.
	Resource *CreateJobRequestTasksTaskSpecResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// Task retry policy.
	RetryPolicy *CreateJobRequestTasksTaskSpecRetryPolicy `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty" type:"Struct"`
	// The task execution configurations.
	//
	// This parameter is required.
	TaskExecutor []*CreateJobRequestTasksTaskSpecTaskExecutor `json:"TaskExecutor,omitempty" xml:"TaskExecutor,omitempty" type:"Repeated"`
	// The list of data volumes mounted to the task. A maximum of 10 groups.
	VolumeMount []*CreateJobRequestTasksTaskSpecVolumeMount `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
}

func (s CreateJobRequestTasksTaskSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpec) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpec) GetResource() *CreateJobRequestTasksTaskSpecResource {
	return s.Resource
}

func (s *CreateJobRequestTasksTaskSpec) GetRetryPolicy() *CreateJobRequestTasksTaskSpecRetryPolicy {
	return s.RetryPolicy
}

func (s *CreateJobRequestTasksTaskSpec) GetTaskExecutor() []*CreateJobRequestTasksTaskSpecTaskExecutor {
	return s.TaskExecutor
}

func (s *CreateJobRequestTasksTaskSpec) GetVolumeMount() []*CreateJobRequestTasksTaskSpecVolumeMount {
	return s.VolumeMount
}

func (s *CreateJobRequestTasksTaskSpec) SetResource(v *CreateJobRequestTasksTaskSpecResource) *CreateJobRequestTasksTaskSpec {
	s.Resource = v
	return s
}

func (s *CreateJobRequestTasksTaskSpec) SetRetryPolicy(v *CreateJobRequestTasksTaskSpecRetryPolicy) *CreateJobRequestTasksTaskSpec {
	s.RetryPolicy = v
	return s
}

func (s *CreateJobRequestTasksTaskSpec) SetTaskExecutor(v []*CreateJobRequestTasksTaskSpecTaskExecutor) *CreateJobRequestTasksTaskSpec {
	s.TaskExecutor = v
	return s
}

func (s *CreateJobRequestTasksTaskSpec) SetVolumeMount(v []*CreateJobRequestTasksTaskSpecVolumeMount) *CreateJobRequestTasksTaskSpec {
	s.VolumeMount = v
	return s
}

func (s *CreateJobRequestTasksTaskSpec) Validate() error {
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			return err
		}
	}
	if s.RetryPolicy != nil {
		if err := s.RetryPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.TaskExecutor != nil {
		for _, item := range s.TaskExecutor {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VolumeMount != nil {
		for _, item := range s.VolumeMount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateJobRequestTasksTaskSpecResource struct {
	// The number of CPUs in the running environment.
	//
	// example:
	//
	// 2
	Cores *float32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The array of the disks.
	Disks          []*CreateJobRequestTasksTaskSpecResourceDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	EnableHT       *bool                                         `json:"EnableHT,omitempty" xml:"EnableHT,omitempty"`
	HostNamePrefix *string                                       `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	// The instance type of the running environment. A maximum of 5 groups.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The memory size of the running environment. Unit: GiB.
	//
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s CreateJobRequestTasksTaskSpecResource) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecResource) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecResource) GetCores() *float32 {
	return s.Cores
}

func (s *CreateJobRequestTasksTaskSpecResource) GetDisks() []*CreateJobRequestTasksTaskSpecResourceDisks {
	return s.Disks
}

func (s *CreateJobRequestTasksTaskSpecResource) GetEnableHT() *bool {
	return s.EnableHT
}

func (s *CreateJobRequestTasksTaskSpecResource) GetHostNamePrefix() *string {
	return s.HostNamePrefix
}

func (s *CreateJobRequestTasksTaskSpecResource) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *CreateJobRequestTasksTaskSpecResource) GetMemory() *float32 {
	return s.Memory
}

func (s *CreateJobRequestTasksTaskSpecResource) SetCores(v float32) *CreateJobRequestTasksTaskSpecResource {
	s.Cores = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecResource) SetDisks(v []*CreateJobRequestTasksTaskSpecResourceDisks) *CreateJobRequestTasksTaskSpecResource {
	s.Disks = v
	return s
}

func (s *CreateJobRequestTasksTaskSpecResource) SetEnableHT(v bool) *CreateJobRequestTasksTaskSpecResource {
	s.EnableHT = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecResource) SetHostNamePrefix(v string) *CreateJobRequestTasksTaskSpecResource {
	s.HostNamePrefix = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecResource) SetInstanceTypes(v []*string) *CreateJobRequestTasksTaskSpecResource {
	s.InstanceTypes = v
	return s
}

func (s *CreateJobRequestTasksTaskSpecResource) SetMemory(v float32) *CreateJobRequestTasksTaskSpecResource {
	s.Memory = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecResource) Validate() error {
	if s.Disks != nil {
		for _, item := range s.Disks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateJobRequestTasksTaskSpecResourceDisks struct {
	// The size of the disk. Unit: GiB.
	//
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The type of the disk. Currently, only System is supported, which indicates the system disk.
	//
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateJobRequestTasksTaskSpecResourceDisks) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecResourceDisks) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecResourceDisks) GetSize() *int32 {
	return s.Size
}

func (s *CreateJobRequestTasksTaskSpecResourceDisks) GetType() *string {
	return s.Type
}

func (s *CreateJobRequestTasksTaskSpecResourceDisks) SetSize(v int32) *CreateJobRequestTasksTaskSpecResourceDisks {
	s.Size = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecResourceDisks) SetType(v string) *CreateJobRequestTasksTaskSpecResourceDisks {
	s.Type = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecResourceDisks) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestTasksTaskSpecRetryPolicy struct {
	// The retry rule. A maximum of 10 groups.
	ExitCodeActions []*CreateJobRequestTasksTaskSpecRetryPolicyExitCodeActions `json:"ExitCodeActions,omitempty" xml:"ExitCodeActions,omitempty" type:"Repeated"`
	// The maximum number of retries. Valid values: 1 to 10. Default value: 3.
	//
	// example:
	//
	// 5
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
}

func (s CreateJobRequestTasksTaskSpecRetryPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecRetryPolicy) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecRetryPolicy) GetExitCodeActions() []*CreateJobRequestTasksTaskSpecRetryPolicyExitCodeActions {
	return s.ExitCodeActions
}

func (s *CreateJobRequestTasksTaskSpecRetryPolicy) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *CreateJobRequestTasksTaskSpecRetryPolicy) SetExitCodeActions(v []*CreateJobRequestTasksTaskSpecRetryPolicyExitCodeActions) *CreateJobRequestTasksTaskSpecRetryPolicy {
	s.ExitCodeActions = v
	return s
}

func (s *CreateJobRequestTasksTaskSpecRetryPolicy) SetRetryCount(v int32) *CreateJobRequestTasksTaskSpecRetryPolicy {
	s.RetryCount = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecRetryPolicy) Validate() error {
	if s.ExitCodeActions != nil {
		for _, item := range s.ExitCodeActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateJobRequestTasksTaskSpecRetryPolicyExitCodeActions struct {
	// The next step behavior of the task.
	//
	// 	- Retry: The job starts a retry when a specific exit code is hit.
	//
	// 	- Exit: The job exits when a specific exit code is hit.
	//
	// This parameter is required.
	//
	// example:
	//
	// Retry
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The task exit code, which is used together with the action to form a job retry rule. Valid values: 0 to 255.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ExitCode *int64 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
}

func (s CreateJobRequestTasksTaskSpecRetryPolicyExitCodeActions) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecRetryPolicyExitCodeActions) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecRetryPolicyExitCodeActions) GetAction() *string {
	return s.Action
}

func (s *CreateJobRequestTasksTaskSpecRetryPolicyExitCodeActions) GetExitCode() *int64 {
	return s.ExitCode
}

func (s *CreateJobRequestTasksTaskSpecRetryPolicyExitCodeActions) SetAction(v string) *CreateJobRequestTasksTaskSpecRetryPolicyExitCodeActions {
	s.Action = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecRetryPolicyExitCodeActions) SetExitCode(v int64) *CreateJobRequestTasksTaskSpecRetryPolicyExitCodeActions {
	s.ExitCode = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecRetryPolicyExitCodeActions) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestTasksTaskSpecTaskExecutor struct {
	// Use the container environment.
	Container *CreateJobRequestTasksTaskSpecTaskExecutorContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// Use a virtual machine environment.
	VM *CreateJobRequestTasksTaskSpecTaskExecutorVM `json:"VM,omitempty" xml:"VM,omitempty" type:"Struct"`
}

func (s CreateJobRequestTasksTaskSpecTaskExecutor) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecTaskExecutor) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutor) GetContainer() *CreateJobRequestTasksTaskSpecTaskExecutorContainer {
	return s.Container
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutor) GetVM() *CreateJobRequestTasksTaskSpecTaskExecutorVM {
	return s.VM
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutor) SetContainer(v *CreateJobRequestTasksTaskSpecTaskExecutorContainer) *CreateJobRequestTasksTaskSpecTaskExecutor {
	s.Container = v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutor) SetVM(v *CreateJobRequestTasksTaskSpecTaskExecutorVM) *CreateJobRequestTasksTaskSpecTaskExecutor {
	s.VM = v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutor) Validate() error {
	if s.Container != nil {
		if err := s.Container.Validate(); err != nil {
			return err
		}
	}
	if s.VM != nil {
		if err := s.VM.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateJobRequestTasksTaskSpecTaskExecutorContainer struct {
	// The application ID.
	//
	// example:
	//
	// ci-vm-32k6LXAi3cOG
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The startup argument of the init container. A maximum of 10 groups.
	Arg []*string `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	// The container startup commands. You can specify up to 20 commands. Each command can be up to 256 characters in length.
	//
	// >
	//
	// 	- If the start command contains spaces (for example, `sleep 60s` ), the input JSON format parameter is `["sleep", "60s"]`.
	//
	// 	- If the startup command is complex, the parameter format may be a combination of `Command: ["/bin/bash"]` and `Arg:["-c", "<customized command>"]`. The `<customized command>` is a user-defined combination of commands and can contain characters such as spaces.
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// The environment variables of the container. A maximum of 20 groups.
	EnvironmentVars []*CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The image of the container.
	//
	// This parameter is required.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/ehpc/hpl:latest
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The working directory of the container.
	//
	// example:
	//
	// /usr/local/
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s CreateJobRequestTasksTaskSpecTaskExecutorContainer) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecTaskExecutorContainer) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) GetAppId() *string {
	return s.AppId
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) GetArg() []*string {
	return s.Arg
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) GetCommand() []*string {
	return s.Command
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) GetEnvironmentVars() []*CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars {
	return s.EnvironmentVars
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) GetImage() *string {
	return s.Image
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) SetAppId(v string) *CreateJobRequestTasksTaskSpecTaskExecutorContainer {
	s.AppId = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) SetArg(v []*string) *CreateJobRequestTasksTaskSpecTaskExecutorContainer {
	s.Arg = v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) SetCommand(v []*string) *CreateJobRequestTasksTaskSpecTaskExecutorContainer {
	s.Command = v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) SetEnvironmentVars(v []*CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars) *CreateJobRequestTasksTaskSpecTaskExecutorContainer {
	s.EnvironmentVars = v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) SetImage(v string) *CreateJobRequestTasksTaskSpecTaskExecutorContainer {
	s.Image = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) SetWorkingDir(v string) *CreateJobRequestTasksTaskSpecTaskExecutorContainer {
	s.WorkingDir = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) Validate() error {
	if s.EnvironmentVars != nil {
		for _, item := range s.EnvironmentVars {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars struct {
	// The name of the environment variable for the container. It can be 1 to 128 characters in length. Format requirement: [0-9a-zA-Z], and underscores, cannot start with a number.
	//
	// example:
	//
	// PATH
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the environment variable for the container. The value must be 0 to 256 bits in length.
	//
	// example:
	//
	// /usr/local/bin
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars) GetName() *string {
	return s.Name
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars) GetValue() *string {
	return s.Value
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars) SetName(v string) *CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars {
	s.Name = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars) SetValue(v string) *CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars {
	s.Value = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestTasksTaskSpecTaskExecutorVM struct {
	// The ID of the virtual machine application.
	//
	// example:
	//
	// ci-vm-9jc58Pm5Leky
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-xxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The logon password of the virtual machine environment. The password must be 8 to 30 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported:
	//
	// ()\\`~!@#$%^&\\*-_+=|{}[]:;\\"<>,.?/ In Windows, the password cannot contain a forward slash (/) as the first character.
	//
	// > We recommend that you use HTTPS to send requests if you specify Password to avoid password leakage.
	//
	// example:
	//
	// EHPC@1234
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The pre-processing script. Base64 encoding is required.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	PrologScript *string `json:"PrologScript,omitempty" xml:"PrologScript,omitempty"`
	// The running-job script. Base64 encoding is required.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s CreateJobRequestTasksTaskSpecTaskExecutorVM) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecTaskExecutorVM) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) GetAppId() *string {
	return s.AppId
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) GetImage() *string {
	return s.Image
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) GetPassword() *string {
	return s.Password
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) GetPrologScript() *string {
	return s.PrologScript
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) GetScript() *string {
	return s.Script
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) SetAppId(v string) *CreateJobRequestTasksTaskSpecTaskExecutorVM {
	s.AppId = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) SetImage(v string) *CreateJobRequestTasksTaskSpecTaskExecutorVM {
	s.Image = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) SetPassword(v string) *CreateJobRequestTasksTaskSpecTaskExecutorVM {
	s.Password = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) SetPrologScript(v string) *CreateJobRequestTasksTaskSpecTaskExecutorVM {
	s.PrologScript = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) SetScript(v string) *CreateJobRequestTasksTaskSpecTaskExecutorVM {
	s.Script = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestTasksTaskSpecVolumeMount struct {
	// The list of data volume mount parameters. Each option is a key-value pair in a JSON string.
	//
	// 	- Format for mounting a NAS file system:{"server":"xxxxx-xxxxx.cn-heyuan.nas.aliyuncs.com","vers":"3","path":"/data","options":"nolock,tcp,noresvport"}
	//
	// > server indicates the address of the mount point of the NAS file system. path indicates the subdirectory of the NAS file system. The subdirectory must start with a (/) and must already exist. vers indicates the version number of the NFS protocol used to mount the file system. We recommend that you use v3. options indicates the custom parameters in the format of "xxx,xxx,xxx".
	//
	// 	- OSS mount format:{"bucket":"xxxxx", "url":"oss-cn-heyuan-internal.aliyuncs.com","path":"/data","akId":"xxxxx","akSecret":"xxxxx"}
	//
	// > bucket indicates the name of the OSS bucket. url indicates the endpoint of the OSS bucket. You can log on to the OSS console and obtain the endpoint on the Overview page of the destination bucket. path indicates the directory structure of the root file of the bucket. The default value is /, which requires that the directory already exists. akId indicates the AccessKey ID. akSecret indicates the AccessKey secret.
	//
	// example:
	//
	// {"server":"xxxxx-xxxxx.cn-heyuan.nas.aliyuncs.com","vers":"3","path":"/data","options":"nolock,tcp,noresvport"}
	MountOptions *string `json:"MountOptions,omitempty" xml:"MountOptions,omitempty"`
	// The directory where the task mounts the data volume.
	//
	// > The content of the mounted directory is overwritten by the content of the volume. Exercise caution when you use the directory.
	//
	// example:
	//
	// /mnt
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// Specifies whether the volume is read-only. Default value: false.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// Currently supported data volume types.
	//
	// 	- alicloud/nas: mounts NAS.
	//
	// 	- alicloud/oss: mounts OSS.
	//
	// example:
	//
	// alicloud/nas
	VolumeDriver *string `json:"VolumeDriver,omitempty" xml:"VolumeDriver,omitempty"`
}

func (s CreateJobRequestTasksTaskSpecVolumeMount) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecVolumeMount) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecVolumeMount) GetMountOptions() *string {
	return s.MountOptions
}

func (s *CreateJobRequestTasksTaskSpecVolumeMount) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateJobRequestTasksTaskSpecVolumeMount) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateJobRequestTasksTaskSpecVolumeMount) GetVolumeDriver() *string {
	return s.VolumeDriver
}

func (s *CreateJobRequestTasksTaskSpecVolumeMount) SetMountOptions(v string) *CreateJobRequestTasksTaskSpecVolumeMount {
	s.MountOptions = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecVolumeMount) SetMountPath(v string) *CreateJobRequestTasksTaskSpecVolumeMount {
	s.MountPath = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecVolumeMount) SetReadOnly(v bool) *CreateJobRequestTasksTaskSpecVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecVolumeMount) SetVolumeDriver(v string) *CreateJobRequestTasksTaskSpecVolumeMount {
	s.VolumeDriver = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecVolumeMount) Validate() error {
	return dara.Validate(s)
}
