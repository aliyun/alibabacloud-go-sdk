// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobTemplateConfig(v *CreateJobTemplateRequestJobTemplateConfig) *CreateJobTemplateRequest
	GetJobTemplateConfig() *CreateJobTemplateRequestJobTemplateConfig
	SetJobTemplateDescription(v string) *CreateJobTemplateRequest
	GetJobTemplateDescription() *string
	SetJobTemplateName(v string) *CreateJobTemplateRequest
	GetJobTemplateName() *string
}

type CreateJobTemplateRequest struct {
	// The configuration of the template.
	JobTemplateConfig *CreateJobTemplateRequestJobTemplateConfig `json:"JobTemplateConfig,omitempty" xml:"JobTemplateConfig,omitempty" type:"Struct"`
	// The description of the template.
	//
	// example:
	//
	// Demo
	JobTemplateDescription *string `json:"JobTemplateDescription,omitempty" xml:"JobTemplateDescription,omitempty"`
	// The name of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-template
	JobTemplateName *string `json:"JobTemplateName,omitempty" xml:"JobTemplateName,omitempty"`
}

func (s CreateJobTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequest) GetJobTemplateConfig() *CreateJobTemplateRequestJobTemplateConfig {
	return s.JobTemplateConfig
}

func (s *CreateJobTemplateRequest) GetJobTemplateDescription() *string {
	return s.JobTemplateDescription
}

func (s *CreateJobTemplateRequest) GetJobTemplateName() *string {
	return s.JobTemplateName
}

func (s *CreateJobTemplateRequest) SetJobTemplateConfig(v *CreateJobTemplateRequestJobTemplateConfig) *CreateJobTemplateRequest {
	s.JobTemplateConfig = v
	return s
}

func (s *CreateJobTemplateRequest) SetJobTemplateDescription(v string) *CreateJobTemplateRequest {
	s.JobTemplateDescription = &v
	return s
}

func (s *CreateJobTemplateRequest) SetJobTemplateName(v string) *CreateJobTemplateRequest {
	s.JobTemplateName = &v
	return s
}

func (s *CreateJobTemplateRequest) Validate() error {
	if s.JobTemplateConfig != nil {
		if err := s.JobTemplateConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateJobTemplateRequestJobTemplateConfig struct {
	// The resource deployment policy.
	DeploymentPolicy *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy `json:"DeploymentPolicy,omitempty" xml:"DeploymentPolicy,omitempty" type:"Struct"`
	// The description of the job.
	//
	// example:
	//
	// Demo
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// HPC
	JobScheduler *string `json:"JobScheduler,omitempty" xml:"JobScheduler,omitempty"`
	// The security policy.
	SecurityPolicy *CreateJobTemplateRequestJobTemplateConfigSecurityPolicy `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty" type:"Struct"`
	// The list of tasks. A maximum of 5 tasks are supported.
	Tasks []*CreateJobTemplateRequestJobTemplateConfigTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s CreateJobTemplateRequestJobTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfig) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfig) GetDeploymentPolicy() *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy {
	return s.DeploymentPolicy
}

func (s *CreateJobTemplateRequestJobTemplateConfig) GetJobDescription() *string {
	return s.JobDescription
}

func (s *CreateJobTemplateRequestJobTemplateConfig) GetJobName() *string {
	return s.JobName
}

func (s *CreateJobTemplateRequestJobTemplateConfig) GetJobScheduler() *string {
	return s.JobScheduler
}

func (s *CreateJobTemplateRequestJobTemplateConfig) GetSecurityPolicy() *CreateJobTemplateRequestJobTemplateConfigSecurityPolicy {
	return s.SecurityPolicy
}

func (s *CreateJobTemplateRequestJobTemplateConfig) GetTasks() []*CreateJobTemplateRequestJobTemplateConfigTasks {
	return s.Tasks
}

func (s *CreateJobTemplateRequestJobTemplateConfig) SetDeploymentPolicy(v *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy) *CreateJobTemplateRequestJobTemplateConfig {
	s.DeploymentPolicy = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfig) SetJobDescription(v string) *CreateJobTemplateRequestJobTemplateConfig {
	s.JobDescription = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfig) SetJobName(v string) *CreateJobTemplateRequestJobTemplateConfig {
	s.JobName = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfig) SetJobScheduler(v string) *CreateJobTemplateRequestJobTemplateConfig {
	s.JobScheduler = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfig) SetSecurityPolicy(v *CreateJobTemplateRequestJobTemplateConfigSecurityPolicy) *CreateJobTemplateRequestJobTemplateConfig {
	s.SecurityPolicy = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfig) SetTasks(v []*CreateJobTemplateRequestJobTemplateConfigTasks) *CreateJobTemplateRequestJobTemplateConfig {
	s.Tasks = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfig) Validate() error {
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

type CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy struct {
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
	// - Performance: performance-enhanced.
	//
	// example:
	//
	// General
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The network configuration.
	Network *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The resource pool of the job.
	//
	// example:
	//
	// my-pool
	Pool *string `json:"Pool,omitempty" xml:"Pool,omitempty"`
	// The priority of the job. A larger value indicates a higher scheduling priority for the job.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The tags of the job.
	Tags []*CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy) GetAllocationSpec() *string {
	return s.AllocationSpec
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy) GetLevel() *string {
	return s.Level
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy) GetNetwork() *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork {
	return s.Network
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy) GetPool() *string {
	return s.Pool
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy) GetTags() []*CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags {
	return s.Tags
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy) SetAllocationSpec(v string) *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy {
	s.AllocationSpec = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy) SetLevel(v string) *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy {
	s.Level = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy) SetNetwork(v *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork) *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy {
	s.Network = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy) SetPool(v string) *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy {
	s.Pool = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy) SetPriority(v int32) *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy {
	s.Priority = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy) SetTags(v []*CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags) *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy {
	s.Tags = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicy) Validate() error {
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork struct {
	// Specifies whether to assign a public IP address to the job. Default value: false.
	//
	// example:
	//
	// false
	EnableExternalIpAddress *bool `json:"EnableExternalIpAddress,omitempty" xml:"EnableExternalIpAddress,omitempty"`
	// The vSwitches.
	//
	// example:
	//
	// vsw-xxxx
	Vswitch []*string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty" type:"Repeated"`
}

func (s CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork) GetEnableExternalIpAddress() *bool {
	return s.EnableExternalIpAddress
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork) GetVswitch() []*string {
	return s.Vswitch
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork) SetEnableExternalIpAddress(v bool) *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork {
	s.EnableExternalIpAddress = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork) SetVswitch(v []*string) *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork {
	s.Vswitch = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork) Validate() error {
	return dara.Validate(s)
}

type CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags struct {
	// The key of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags) GetKey() *string {
	return s.Key
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags) GetValue() *string {
	return s.Value
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags) SetKey(v string) *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags {
	s.Key = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags) SetValue(v string) *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags {
	s.Value = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags) Validate() error {
	return dara.Validate(s)
}

type CreateJobTemplateRequestJobTemplateConfigSecurityPolicy struct {
	// The security group.
	SecurityGroup *CreateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty" type:"Struct"`
}

func (s CreateJobTemplateRequestJobTemplateConfigSecurityPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigSecurityPolicy) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigSecurityPolicy) GetSecurityGroup() *CreateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup {
	return s.SecurityGroup
}

func (s *CreateJobTemplateRequestJobTemplateConfigSecurityPolicy) SetSecurityGroup(v *CreateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup) *CreateJobTemplateRequestJobTemplateConfigSecurityPolicy {
	s.SecurityGroup = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigSecurityPolicy) Validate() error {
	if s.SecurityGroup != nil {
		if err := s.SecurityGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup struct {
	// The security group IDs.
	//
	// example:
	//
	// sg-xxxx
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s CreateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup) SetSecurityGroupIds(v []*string) *CreateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup) Validate() error {
	return dara.Validate(s)
}

type CreateJobTemplateRequestJobTemplateConfigTasks struct {
	// The task execution policy.
	ExecutorPolicy *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy `json:"ExecutorPolicy,omitempty" xml:"ExecutorPolicy,omitempty" type:"Struct"`
	// The name of the task.
	//
	// example:
	//
	// task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The specification details of the task.
	TaskSpec *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec `json:"TaskSpec,omitempty" xml:"TaskSpec,omitempty" type:"Struct"`
	// Specifies whether the job is a long-running job. Valid values:
	//
	// - true: The job is a background service job.
	//
	// - false: The job is a batch job.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	TaskSustainable *bool `json:"TaskSustainable,omitempty" xml:"TaskSustainable,omitempty"`
}

func (s CreateJobTemplateRequestJobTemplateConfigTasks) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigTasks) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasks) GetExecutorPolicy() *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy {
	return s.ExecutorPolicy
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasks) GetTaskSpec() *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec {
	return s.TaskSpec
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasks) GetTaskSustainable() *bool {
	return s.TaskSustainable
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasks) SetExecutorPolicy(v *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy) *CreateJobTemplateRequestJobTemplateConfigTasks {
	s.ExecutorPolicy = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasks) SetTaskName(v string) *CreateJobTemplateRequestJobTemplateConfigTasks {
	s.TaskName = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasks) SetTaskSpec(v *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec) *CreateJobTemplateRequestJobTemplateConfigTasks {
	s.TaskSpec = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasks) SetTaskSustainable(v bool) *CreateJobTemplateRequestJobTemplateConfigTasks {
	s.TaskSustainable = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasks) Validate() error {
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

type CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy struct {
	// The details of the array job.
	ArraySpec *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec `json:"ArraySpec,omitempty" xml:"ArraySpec,omitempty" type:"Struct"`
	// The maximum number of nodes on which the job can run.
	//
	// example:
	//
	// 1
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy) GetArraySpec() *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec {
	return s.ArraySpec
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy) SetArraySpec(v *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy {
	s.ArraySpec = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy) SetMaxCount(v int32) *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy {
	s.MaxCount = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy) Validate() error {
	if s.ArraySpec != nil {
		if err := s.ArraySpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec struct {
	// The end value of the array job index. Valid values: 0 to 4999. The value must be greater than or equal to the value of IndexStart.
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
	// The step size of the array job index.
	//
	// example:
	//
	// 1
	IndexStep *int32 `json:"IndexStep,omitempty" xml:"IndexStep,omitempty"`
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) GetIndexEnd() *int32 {
	return s.IndexEnd
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) GetIndexStart() *int32 {
	return s.IndexStart
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) GetIndexStep() *int32 {
	return s.IndexStep
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) SetIndexEnd(v int32) *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec {
	s.IndexEnd = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) SetIndexStart(v int32) *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec {
	s.IndexStart = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) SetIndexStep(v int32) *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec {
	s.IndexStep = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) Validate() error {
	return dara.Validate(s)
}

type CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec struct {
	// The resource information of the runtime environment.
	Resource *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// The retry policy of the task.
	RetryPolicy *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty" type:"Struct"`
	// The execution configuration of the task.
	TaskExecutor []*CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor `json:"TaskExecutor,omitempty" xml:"TaskExecutor,omitempty" type:"Repeated"`
	// The list of data volumes mounted to the task. A maximum of 10 data volumes are supported.
	VolumeMount []*CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec) GetResource() *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource {
	return s.Resource
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec) GetRetryPolicy() *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy {
	return s.RetryPolicy
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec) GetTaskExecutor() []*CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor {
	return s.TaskExecutor
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec) GetVolumeMount() []*CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount {
	return s.VolumeMount
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec) SetResource(v *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec {
	s.Resource = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec) SetRetryPolicy(v *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec {
	s.RetryPolicy = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec) SetTaskExecutor(v []*CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec {
	s.TaskExecutor = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec) SetVolumeMount(v []*CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec {
	s.VolumeMount = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpec) Validate() error {
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

type CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource struct {
	// The number of CPU cores in the runtime environment.
	//
	// example:
	//
	// 2
	Cores *float32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The array of cloud disks.
	Disks []*CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// Specifies whether hyper-threading is enabled in the runtime environment.
	//
	// example:
	//
	// true
	EnableHT *bool `json:"EnableHT,omitempty" xml:"EnableHT,omitempty"`
	// The hostname prefix of the runtime environment.
	//
	// example:
	//
	// compute
	HostNamePrefix *string `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	// The instance types of the runtime environment.
	//
	// example:
	//
	// ecs.c7a.8xlarge
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The memory size of the runtime environment. Unit: GiB.
	//
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) GetCores() *float32 {
	return s.Cores
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) GetDisks() []*CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks {
	return s.Disks
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) GetEnableHT() *bool {
	return s.EnableHT
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) GetHostNamePrefix() *string {
	return s.HostNamePrefix
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) GetMemory() *float32 {
	return s.Memory
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) SetCores(v float32) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource {
	s.Cores = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) SetDisks(v []*CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource {
	s.Disks = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) SetEnableHT(v bool) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource {
	s.EnableHT = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) SetHostNamePrefix(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource {
	s.HostNamePrefix = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) SetInstanceTypes(v []*string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource {
	s.InstanceTypes = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) SetMemory(v float32) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource {
	s.Memory = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) Validate() error {
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

type CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks struct {
	// The size of the cloud disk. Unit: GiB.
	//
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The type of the cloud disk.
	//
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks) GetSize() *int32 {
	return s.Size
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks) GetType() *string {
	return s.Type
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks) SetSize(v int32) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks {
	s.Size = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks) SetType(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks {
	s.Type = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks) Validate() error {
	return dara.Validate(s)
}

type CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy struct {
	// The retry rules.
	ExitCodeActions []*CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions `json:"ExitCodeActions,omitempty" xml:"ExitCodeActions,omitempty" type:"Repeated"`
	// The number of retries.
	//
	// example:
	//
	// 5
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy) GetExitCodeActions() []*CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions {
	return s.ExitCodeActions
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy) SetExitCodeActions(v []*CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy {
	s.ExitCodeActions = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy) SetRetryCount(v int32) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy {
	s.RetryCount = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy) Validate() error {
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

type CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions struct {
	// The next action of the task. Valid values:
	//
	// - Retry: Retries the task.
	//
	// - Exit: Exits the task.
	//
	// example:
	//
	// Retry
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The exit code of the task. This code is used together with the Action parameter to form the retry rules for the job. Valid values: 0 to 255.
	//
	// example:
	//
	// 1
	ExitCode *int64 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) GetAction() *string {
	return s.Action
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) GetExitCode() *int64 {
	return s.ExitCode
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) SetAction(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions {
	s.Action = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) SetExitCode(v int64) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions {
	s.ExitCode = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) Validate() error {
	return dara.Validate(s)
}

type CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor struct {
	// The container environment configuration.
	Container *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// The virtual machine environment configuration.
	VM *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM `json:"VM,omitempty" xml:"VM,omitempty" type:"Struct"`
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor) GetContainer() *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	return s.Container
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor) GetVM() *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	return s.VM
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor) SetContainer(v *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor {
	s.Container = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor) SetVM(v *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor {
	s.VM = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor) Validate() error {
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

type CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer struct {
	// The ID of the container application.
	//
	// example:
	//
	// ci-vm-32k6LXAi3cOG
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The arguments for the container startup commands.
	//
	// example:
	//
	// 60s
	Arg []*string `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	// The list of startup commands for the container.
	//
	// example:
	//
	// sleep
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// The environment variables of the container.
	EnvironmentVars []*CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The GPU driver version.
	//
	// example:
	//
	// tesla=535
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// The container image.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/ehpc/hpl:latest
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The list of mount parameters for the custom image registry. The parameters are in key-value pair format and passed in JSON.
	//
	// example:
	//
	// {}
	ImageRegistryOptions *string `json:"ImageRegistryOptions,omitempty" xml:"ImageRegistryOptions,omitempty"`
	// Specifies whether the container runs in privileged mode. Default value: false.
	//
	// example:
	//
	// true
	Privileged *bool `json:"Privileged,omitempty" xml:"Privileged,omitempty"`
	// The working directory of the container.
	//
	// example:
	//
	// /usr/local/
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetAppId() *string {
	return s.AppId
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetArg() []*string {
	return s.Arg
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetCommand() []*string {
	return s.Command
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetEnvironmentVars() []*CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars {
	return s.EnvironmentVars
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetGpuDriverVersion() *string {
	return s.GpuDriverVersion
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetImage() *string {
	return s.Image
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetImageRegistryOptions() *string {
	return s.ImageRegistryOptions
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetPrivileged() *bool {
	return s.Privileged
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetAppId(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.AppId = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetArg(v []*string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.Arg = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetCommand(v []*string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.Command = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetEnvironmentVars(v []*CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.EnvironmentVars = v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetGpuDriverVersion(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.GpuDriverVersion = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetImage(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.Image = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetImageRegistryOptions(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.ImageRegistryOptions = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetPrivileged(v bool) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.Privileged = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetWorkingDir(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.WorkingDir = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) Validate() error {
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

type CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars struct {
	// The name of the environment variable.
	//
	// example:
	//
	// PATH
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the environment variable.
	//
	// example:
	//
	// /usr/local/bin
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) GetName() *string {
	return s.Name
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) GetValue() *string {
	return s.Value
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) SetName(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars {
	s.Name = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) SetValue(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars {
	s.Value = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) Validate() error {
	return dara.Validate(s)
}

type CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM struct {
	// The ID of the virtual machine application.
	//
	// example:
	//
	// ci-vm-9jc58Pm5Leky
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The image ID.
	//
	// example:
	//
	// m-xxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The logon password for the virtual machine environment.
	//
	// example:
	//
	// EHPC@1234
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The prolog script. The script must be Base64-encoded.
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

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) GetAppId() *string {
	return s.AppId
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) GetImage() *string {
	return s.Image
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) GetPassword() *string {
	return s.Password
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) GetPrologScript() *string {
	return s.PrologScript
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) GetScript() *string {
	return s.Script
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) SetAppId(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	s.AppId = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) SetImage(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	s.Image = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) SetPassword(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	s.Password = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) SetPrologScript(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	s.PrologScript = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) SetScript(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	s.Script = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) Validate() error {
	return dara.Validate(s)
}

type CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount struct {
	// The list of mount parameters for the data volume. The parameters are in key-value pair format and passed in JSON.
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
	// Specifies whether the data volume is read-only.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The type of the data volume.
	//
	// example:
	//
	// alicloud/nas
	VolumeDriver *string `json:"VolumeDriver,omitempty" xml:"VolumeDriver,omitempty"`
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) GetMountOptions() *string {
	return s.MountOptions
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) GetVolumeDriver() *string {
	return s.VolumeDriver
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) SetMountOptions(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount {
	s.MountOptions = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) SetMountPath(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount {
	s.MountPath = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) SetReadOnly(v bool) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) SetVolumeDriver(v string) *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount {
	s.VolumeDriver = &v
	return s
}

func (s *CreateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) Validate() error {
	return dara.Validate(s)
}
