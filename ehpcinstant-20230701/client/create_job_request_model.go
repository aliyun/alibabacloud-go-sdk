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
	SetJobTemplateId(v string) *CreateJobRequest
	GetJobTemplateId() *string
	SetSecurityPolicy(v *CreateJobRequestSecurityPolicy) *CreateJobRequest
	GetSecurityPolicy() *CreateJobRequestSecurityPolicy
	SetTasks(v []*CreateJobRequestTasks) *CreateJobRequest
	GetTasks() []*CreateJobRequestTasks
}

type CreateJobRequest struct {
	// The dependency policy.
	DependencyPolicy *CreateJobRequestDependencyPolicy `json:"DependencyPolicy,omitempty" xml:"DependencyPolicy,omitempty" type:"Struct"`
	// The resource deployment policy.
	DeploymentPolicy *CreateJobRequestDeploymentPolicy `json:"DeploymentPolicy,omitempty" xml:"DeploymentPolicy,omitempty" type:"Struct"`
	// The job description.
	//
	// example:
	//
	// Demo
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// The job name. The name must be 2 to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// testjob
	JobName      *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobScheduler *string `json:"JobScheduler,omitempty" xml:"JobScheduler,omitempty"`
	// The job template ID.
	//
	// example:
	//
	// jt-xxxx
	JobTemplateId *string `json:"JobTemplateId,omitempty" xml:"JobTemplateId,omitempty"`
	// The security policy.
	SecurityPolicy *CreateJobRequestSecurityPolicy `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty" type:"Struct"`
	// The task list. Currently, only one task is supported.
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

func (s *CreateJobRequest) GetJobTemplateId() *string {
	return s.JobTemplateId
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

func (s *CreateJobRequest) SetJobTemplateId(v string) *CreateJobRequest {
	s.JobTemplateId = &v
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
	// The job dependencies. A maximum of 10 groups are supported.
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
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// job-bjxxxxxxxxxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The dependency type. Valid values:
	//
	// - AfterSucceeded: **All tasks*	- in the dependent job or array job run successfully (exit code 0).
	//
	// - AfterFailed: **Any task*	- in the dependent job or array job fails (exit code is not 0).
	//
	// - AfterAny: The dependent job finishes running (succeeded or failed).
	//
	// - AfterCorresponding: The corresponding task in the dependent array job runs successfully (exit code 0).
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
	// The resource type.
	//
	// example:
	//
	// Dedicated
	AllocationSpec *string `json:"AllocationSpec,omitempty" xml:"AllocationSpec,omitempty"`
	// The computing power level. This parameter is valid only when the resource type is economy. Valid values:
	//
	// - General: general-purpose.
	//
	// - Performance: compute-optimized.
	//
	// Default value: General
	//
	// example:
	//
	// General
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The network configuration.
	Network *CreateJobRequestDeploymentPolicyNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The job resource pool.
	//
	// example:
	//
	// compute
	Pool *string `json:"Pool,omitempty" xml:"Pool,omitempty"`
	// The job priority. A larger value indicates a higher scheduling priority. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The job tag information. A maximum of 20 tags are supported.
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
	// Specifies whether to create a public IP address for the job.
	//
	// example:
	//
	// true
	EnableExternalIpAddress *bool `json:"EnableExternalIpAddress,omitempty" xml:"EnableExternalIpAddress,omitempty"`
	// The vSwitch array.
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
	// The tag key of the job. If you specify this parameter, the value cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the job. If you specify this parameter, the value can be an empty string. The tag value can be up to 128 characters in length and cannot contain http:// or https://.
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
	// The security group.
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
	// The task name. The name must be 2 to 32 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task specification details.
	TaskSpec *CreateJobRequestTasksTaskSpec `json:"TaskSpec,omitempty" xml:"TaskSpec,omitempty" type:"Struct"`
	// Specifies whether the job is a long-running job. Valid values:
	//
	// - true: background service job.
	//
	// - false: batch job.
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
	// The array job details. Sub-job index values are passed to the runtime environment through environment variables, which can be referenced by user applications. The environment variables include:
	//
	// - EHPC_JOB_NAME: the job name, corresponding to the JobName parameter.
	//
	// - EHPC_JOB_ID: the job ID.
	//
	// - EHPC_TASK_NAME: the task name, corresponding to the TaskName parameter.
	//
	// - EHPC_EXECUTOR_ID: the executor ID.
	//
	// - EHPC_ARRAY_TASK_ID: the sub-job index value.
	//
	// - EHPC_ARRAY_TASK_COUNT: the total number of sub-jobs.
	//
	// - EHPC_ARRAY_TASK_MAX: the maximum sub-job index value, corresponding to the IndexStart parameter.
	//
	// - EHPC_ARRAY_TASK_MIN: the minimum sub-job index value, corresponding to the IndexEnd parameter.
	//
	// - EHPC_ARRAY_TASK_STEP: the sub-job index step, corresponding to the IndexStep parameter.
	ArraySpec *CreateJobRequestTasksExecutorPolicyArraySpec `json:"ArraySpec,omitempty" xml:"ArraySpec,omitempty" type:"Struct"`
	// The maximum number of nodes for the job.
	//
	// > The value must comply with the following formula: `MaxCount = (IndexEnd - IndexStart) / IndexStep + 1`
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
	// The end value of the array job index. Valid values: 0 to 4999. The value must be greater than or equal to IndexStart.
	//
	// example:
	//
	// 9
	IndexEnd *int32 `json:"IndexEnd,omitempty" xml:"IndexEnd,omitempty"`
	// The start value of the array job index. Valid values: 0 to 4999.
	//
	// example:
	//
	// 0
	IndexStart *int32 `json:"IndexStart,omitempty" xml:"IndexStart,omitempty"`
	// The interval between indexes in an array job.
	//
	// > If the array job has the properties IndexStart=1, IndexEnd=5, and IndexStep=2, the array job contains three sub-jobs with indexes 1, 3, and 5. Your application can access these indexes through environment variables.
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
	// The resource information of the runtime environment.
	Resource *CreateJobRequestTasksTaskSpecResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// The task retry policy.
	RetryPolicy *CreateJobRequestTasksTaskSpecRetryPolicy `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty" type:"Struct"`
	// The task execution configuration.
	//
	// This parameter is required.
	TaskExecutor []*CreateJobRequestTasksTaskSpecTaskExecutor `json:"TaskExecutor,omitempty" xml:"TaskExecutor,omitempty" type:"Repeated"`
	// The list of data volumes mounted to the task. A maximum of 10 data volumes are supported.
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
	// The number of CPUs in the runtime environment.
	//
	// example:
	//
	// 2
	Cores *float32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The cloud disk array.
	Disks []*CreateJobRequestTasksTaskSpecResourceDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// Specifies whether hyper-threading is enabled in the runtime environment. Default value: true.
	//
	// example:
	//
	// true
	EnableHT *bool `json:"EnableHT,omitempty" xml:"EnableHT,omitempty"`
	// The hostname prefix of the runtime environment. The following limits apply:
	//
	// - A period (.) and a hyphen (-) cannot be used as the first or last character, or consecutively.
	//
	// - Windows environment: The value can be up to 10 characters in length, cannot contain periods (.), and cannot consist of digits only. Uppercase and lowercase letters, digits, and hyphens (-) are allowed.
	//
	// - Linux environment: The value can be up to 32 characters in length and can contain multiple periods (.). The hostname is divided into segments by periods. Each segment can contain uppercase and lowercase letters, digits, and hyphens (-).
	//
	// example:
	//
	// compute
	HostNamePrefix *string `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	// The instance types of the runtime environment. A maximum of 5 instance types are supported.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The memory size of the runtime environment. Unit: GiB.
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
	// The cloud disk size. Unit: GiB.
	//
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The cloud disk type. Currently, only System is supported, which indicates a system cloud disk.
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
	// The retry rules. A maximum of 10 rules are supported.
	ExitCodeActions []*CreateJobRequestTasksTaskSpecRetryPolicyExitCodeActions `json:"ExitCodeActions,omitempty" xml:"ExitCodeActions,omitempty" type:"Repeated"`
	// The number of retries. Valid values: 1 to 10. Default value: 3.
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
	// The next action for the node. Valid values:
	//
	// - Retry: When a specific exit code is matched, the job starts a new retry.
	//
	// - Exit: When a specific exit code is matched, the job exits.
	//
	// This parameter is required.
	//
	// example:
	//
	// Retry
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The task exit code, which is used together with Action to form a job retry rule. Valid values: 0 to 255.
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
	// The container environment settings.
	Container *CreateJobRequestTasksTaskSpecTaskExecutorContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// The virtual machine environment settings.
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
	// The container application ID.
	//
	// example:
	//
	// ci-vm-32k6LXAi3cOG
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The arguments for the container startup command. A maximum of 10 arguments are supported.
	Arg []*string `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	// The list of container startup commands. A maximum of 20 commands are supported. Each command can contain up to 256 characters.
	//
	// > 1. If a startup command contains spaces (for example, `sleep 60s`), pass the JSON parameter as `["sleep", "60s"]`.
	//
	// > 2. If a startup command is complex, use a combination of `Command: ["/bin/bash"]` and `Arg:["-c", "<customized command>"]`, where `<customized command>` is a user-defined command that can contain spaces and other characters.
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// The environment variables of the container. A maximum of 20 environment variables are supported.
	EnvironmentVars []*CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The container image.
	//
	// This parameter is required.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/ehpc/hpl:latest
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The list of mount parameters for a self-managed image registry. The parameters are in key-value format and passed as a JSON string.
	//
	// - Reference format: {"ImageRegistryType":"https","ImageRegistryServer":"xxx","ImageRegistryUserName":"xxx","ImageRegistryPassword":"xxx"}
	//
	// example:
	//
	// {"ImageRegistryType":"https","ImageRegistryServer":"xxx","ImageRegistryUserName":"xxx","ImageRegistryPassword":"xxx"}
	ImageRegistryOptions *string `json:"ImageRegistryOptions,omitempty" xml:"ImageRegistryOptions,omitempty"`
	// The container working directory.
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

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) GetImageRegistryOptions() *string {
	return s.ImageRegistryOptions
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

func (s *CreateJobRequestTasksTaskSpecTaskExecutorContainer) SetImageRegistryOptions(v string) *CreateJobRequestTasksTaskSpecTaskExecutorContainer {
	s.ImageRegistryOptions = &v
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
	// The environment variable name. The name must be 1 to 128 characters in length. The format is [0-9a-zA-Z] and underscores. The name cannot start with a digit.
	//
	// example:
	//
	// PATH
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The environment variable value. The value can be 0 to 256 characters in length.
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
	// The virtual machine application ID.
	//
	// example:
	//
	// ci-vm-9jc58Pm5Leky
	AppId           *string                                                       `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EnvironmentVars []*CreateJobRequestTasksTaskSpecTaskExecutorVMEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-xxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The logon password for the virtual machine environment. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Supported special characters are:
	//
	// ()`~!@#$%^&*-_+=|{}[]:;\\"<>,.?/
	//
	// For Windows environments, the password cannot start with a forward slash (/).
	//
	// > If you specify the Password parameter, use HTTPS to send the request to prevent password leakage.
	//
	// example:
	//
	// EHPC@1234
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The pre-processing script. The script must be Base64-encoded.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	PrologScript *string `json:"PrologScript,omitempty" xml:"PrologScript,omitempty"`
	// The job execution script. The script must be Base64-encoded.
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

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) GetEnvironmentVars() []*CreateJobRequestTasksTaskSpecTaskExecutorVMEnvironmentVars {
	return s.EnvironmentVars
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

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVM) SetEnvironmentVars(v []*CreateJobRequestTasksTaskSpecTaskExecutorVMEnvironmentVars) *CreateJobRequestTasksTaskSpecTaskExecutorVM {
	s.EnvironmentVars = v
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

type CreateJobRequestTasksTaskSpecTaskExecutorVMEnvironmentVars struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateJobRequestTasksTaskSpecTaskExecutorVMEnvironmentVars) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequestTasksTaskSpecTaskExecutorVMEnvironmentVars) GoString() string {
	return s.String()
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVMEnvironmentVars) GetName() *string {
	return s.Name
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVMEnvironmentVars) GetValue() *string {
	return s.Value
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVMEnvironmentVars) SetName(v string) *CreateJobRequestTasksTaskSpecTaskExecutorVMEnvironmentVars {
	s.Name = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVMEnvironmentVars) SetValue(v string) *CreateJobRequestTasksTaskSpecTaskExecutorVMEnvironmentVars {
	s.Value = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecTaskExecutorVMEnvironmentVars) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestTasksTaskSpecVolumeMount struct {
	// The list of volume mount parameters. Passed as key-value pairs in JSON format.
	//
	// - Reference format for mounting NAS: {"server":"xxxxx-xxxxx.cn-heyuan.nas.aliyuncs.com","vers":"3","path":"/data","options":"nolock,tcp,noresvport"}
	//
	// > server specifies the mount target address of the NAS file system. path specifies a subdirectory under the NAS path, starting with /, and the directory must already exist. vers specifies the NFS protocol version for mounting NAS. Version 3 is recommended. options specifies custom parameters for mounting NAS, in the format "xxx,xxx,xxx".
	//
	// - Reference format for mounting OSS: {"bucket":"xxxxx", "url":"oss-cn-heyuan-internal.aliyuncs.com","path":"/data","akId":"xxxxx","akSecret":"xxxxx"}
	//
	// > bucket specifies the name of the OSS bucket. url specifies the endpoint of the OSS bucket. You can log on to the OSS console and obtain the endpoint on the overview page of the target bucket. path specifies the directory structure relative to the root of the bucket when mounting. The default value is /. The directory must already exist. akId specifies the AccessKey ID used for direct authorization with an AccessKey pair. akSecret specifies the AccessKey secret used for direct authorization with an AccessKey pair.
	//
	// example:
	//
	// {"server":"xxxxx-xxxxx.cn-heyuan.nas.aliyuncs.com","vers":"3","path":"/data","options":"nolock,tcp,noresvport"}
	MountOptions *string `json:"MountOptions,omitempty" xml:"MountOptions,omitempty"`
	// The directory where the data volume is mounted to the task.
	//
	// example:
	//
	// /mnt
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// Specifies whether the data volume is read-only. Default value: false.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The supported data volume type. Valid values:
	//
	// - alicloud/nas: mounts a NAS file system.
	//
	// - alicloud/oss: mounts an OSS bucket.
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
