// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobTemplateConfig(v *UpdateJobTemplateRequestJobTemplateConfig) *UpdateJobTemplateRequest
	GetJobTemplateConfig() *UpdateJobTemplateRequestJobTemplateConfig
	SetJobTemplateDescription(v string) *UpdateJobTemplateRequest
	GetJobTemplateDescription() *string
	SetJobTemplateId(v string) *UpdateJobTemplateRequest
	GetJobTemplateId() *string
	SetJobTemplateName(v string) *UpdateJobTemplateRequest
	GetJobTemplateName() *string
}

type UpdateJobTemplateRequest struct {
	// The template configuration.
	JobTemplateConfig *UpdateJobTemplateRequestJobTemplateConfig `json:"JobTemplateConfig,omitempty" xml:"JobTemplateConfig,omitempty" type:"Struct"`
	// The template description.
	//
	// example:
	//
	// Demo
	JobTemplateDescription *string `json:"JobTemplateDescription,omitempty" xml:"JobTemplateDescription,omitempty"`
	// The job template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// jt-xxxx
	JobTemplateId *string `json:"JobTemplateId,omitempty" xml:"JobTemplateId,omitempty"`
	// The template name.
	//
	// example:
	//
	// my-template
	JobTemplateName *string `json:"JobTemplateName,omitempty" xml:"JobTemplateName,omitempty"`
}

func (s UpdateJobTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequest) GetJobTemplateConfig() *UpdateJobTemplateRequestJobTemplateConfig {
	return s.JobTemplateConfig
}

func (s *UpdateJobTemplateRequest) GetJobTemplateDescription() *string {
	return s.JobTemplateDescription
}

func (s *UpdateJobTemplateRequest) GetJobTemplateId() *string {
	return s.JobTemplateId
}

func (s *UpdateJobTemplateRequest) GetJobTemplateName() *string {
	return s.JobTemplateName
}

func (s *UpdateJobTemplateRequest) SetJobTemplateConfig(v *UpdateJobTemplateRequestJobTemplateConfig) *UpdateJobTemplateRequest {
	s.JobTemplateConfig = v
	return s
}

func (s *UpdateJobTemplateRequest) SetJobTemplateDescription(v string) *UpdateJobTemplateRequest {
	s.JobTemplateDescription = &v
	return s
}

func (s *UpdateJobTemplateRequest) SetJobTemplateId(v string) *UpdateJobTemplateRequest {
	s.JobTemplateId = &v
	return s
}

func (s *UpdateJobTemplateRequest) SetJobTemplateName(v string) *UpdateJobTemplateRequest {
	s.JobTemplateName = &v
	return s
}

func (s *UpdateJobTemplateRequest) Validate() error {
	if s.JobTemplateConfig != nil {
		if err := s.JobTemplateConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateJobTemplateRequestJobTemplateConfig struct {
	// The resource deployment policy.
	DeploymentPolicy *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy `json:"DeploymentPolicy,omitempty" xml:"DeploymentPolicy,omitempty" type:"Struct"`
	// The job description.
	//
	// example:
	//
	// Demo
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// The job name.
	//
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The job scheduler type.
	//
	// example:
	//
	// HPC
	JobScheduler *string `json:"JobScheduler,omitempty" xml:"JobScheduler,omitempty"`
	// The security policy.
	SecurityPolicy *UpdateJobTemplateRequestJobTemplateConfigSecurityPolicy `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty" type:"Struct"`
	// The task list. A maximum of five tasks are supported.
	Tasks []*UpdateJobTemplateRequestJobTemplateConfigTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s UpdateJobTemplateRequestJobTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfig) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfig) GetDeploymentPolicy() *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy {
	return s.DeploymentPolicy
}

func (s *UpdateJobTemplateRequestJobTemplateConfig) GetJobDescription() *string {
	return s.JobDescription
}

func (s *UpdateJobTemplateRequestJobTemplateConfig) GetJobName() *string {
	return s.JobName
}

func (s *UpdateJobTemplateRequestJobTemplateConfig) GetJobScheduler() *string {
	return s.JobScheduler
}

func (s *UpdateJobTemplateRequestJobTemplateConfig) GetSecurityPolicy() *UpdateJobTemplateRequestJobTemplateConfigSecurityPolicy {
	return s.SecurityPolicy
}

func (s *UpdateJobTemplateRequestJobTemplateConfig) GetTasks() []*UpdateJobTemplateRequestJobTemplateConfigTasks {
	return s.Tasks
}

func (s *UpdateJobTemplateRequestJobTemplateConfig) SetDeploymentPolicy(v *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy) *UpdateJobTemplateRequestJobTemplateConfig {
	s.DeploymentPolicy = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfig) SetJobDescription(v string) *UpdateJobTemplateRequestJobTemplateConfig {
	s.JobDescription = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfig) SetJobName(v string) *UpdateJobTemplateRequestJobTemplateConfig {
	s.JobName = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfig) SetJobScheduler(v string) *UpdateJobTemplateRequestJobTemplateConfig {
	s.JobScheduler = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfig) SetSecurityPolicy(v *UpdateJobTemplateRequestJobTemplateConfigSecurityPolicy) *UpdateJobTemplateRequestJobTemplateConfig {
	s.SecurityPolicy = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfig) SetTasks(v []*UpdateJobTemplateRequestJobTemplateConfigTasks) *UpdateJobTemplateRequestJobTemplateConfig {
	s.Tasks = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfig) Validate() error {
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

type UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy struct {
	// The resource type.
	//
	// example:
	//
	// Dedicated
	AllocationSpec *string `json:"AllocationSpec,omitempty" xml:"AllocationSpec,omitempty"`
	// The computing power level. This parameter is valid only when the resource type is economy. Valid values: General: general-purpose. Performance: performance-enhanced.
	//
	// example:
	//
	// General
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The network configuration.
	Network *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The job resource pool.
	//
	// example:
	//
	// my-pool
	Pool *string `json:"Pool,omitempty" xml:"Pool,omitempty"`
	// The job priority. A larger value indicates a higher scheduling priority for the job.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The job tags.
	Tags []*UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy) GetAllocationSpec() *string {
	return s.AllocationSpec
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy) GetLevel() *string {
	return s.Level
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy) GetNetwork() *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork {
	return s.Network
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy) GetPool() *string {
	return s.Pool
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy) GetTags() []*UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags {
	return s.Tags
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy) SetAllocationSpec(v string) *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy {
	s.AllocationSpec = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy) SetLevel(v string) *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy {
	s.Level = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy) SetNetwork(v *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork) *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy {
	s.Network = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy) SetPool(v string) *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy {
	s.Pool = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy) SetPriority(v int32) *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy {
	s.Priority = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy) SetTags(v []*UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags) *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy {
	s.Tags = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicy) Validate() error {
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

type UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork struct {
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

func (s UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork) GetEnableExternalIpAddress() *bool {
	return s.EnableExternalIpAddress
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork) GetVswitch() []*string {
	return s.Vswitch
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork) SetEnableExternalIpAddress(v bool) *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork {
	s.EnableExternalIpAddress = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork) SetVswitch(v []*string) *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork {
	s.Vswitch = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyNetwork) Validate() error {
	return dara.Validate(s)
}

type UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags struct {
	// The job tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The job tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags) GetKey() *string {
	return s.Key
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags) GetValue() *string {
	return s.Value
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags) SetKey(v string) *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags {
	s.Key = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags) SetValue(v string) *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags {
	s.Value = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigDeploymentPolicyTags) Validate() error {
	return dara.Validate(s)
}

type UpdateJobTemplateRequestJobTemplateConfigSecurityPolicy struct {
	// The security group.
	SecurityGroup *UpdateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty" type:"Struct"`
}

func (s UpdateJobTemplateRequestJobTemplateConfigSecurityPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigSecurityPolicy) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigSecurityPolicy) GetSecurityGroup() *UpdateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup {
	return s.SecurityGroup
}

func (s *UpdateJobTemplateRequestJobTemplateConfigSecurityPolicy) SetSecurityGroup(v *UpdateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup) *UpdateJobTemplateRequestJobTemplateConfigSecurityPolicy {
	s.SecurityGroup = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigSecurityPolicy) Validate() error {
	if s.SecurityGroup != nil {
		if err := s.SecurityGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup struct {
	// The security group IDs.
	//
	// example:
	//
	// sg-xxxx
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s UpdateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *UpdateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup) SetSecurityGroupIds(v []*string) *UpdateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup {
	s.SecurityGroupIds = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigSecurityPolicySecurityGroup) Validate() error {
	return dara.Validate(s)
}

type UpdateJobTemplateRequestJobTemplateConfigTasks struct {
	// The task execution policy.
	ExecutorPolicy *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy `json:"ExecutorPolicy,omitempty" xml:"ExecutorPolicy,omitempty" type:"Struct"`
	// The task name.
	//
	// example:
	//
	// task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task specification details.
	TaskSpec *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec `json:"TaskSpec,omitempty" xml:"TaskSpec,omitempty" type:"Struct"`
	// Specifies whether the job is long-running. Valid values: true: Background service job. false: Batch job. Default value: false.
	//
	// example:
	//
	// false
	TaskSustainable *bool `json:"TaskSustainable,omitempty" xml:"TaskSustainable,omitempty"`
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasks) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasks) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasks) GetExecutorPolicy() *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy {
	return s.ExecutorPolicy
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasks) GetTaskSpec() *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec {
	return s.TaskSpec
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasks) GetTaskSustainable() *bool {
	return s.TaskSustainable
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasks) SetExecutorPolicy(v *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy) *UpdateJobTemplateRequestJobTemplateConfigTasks {
	s.ExecutorPolicy = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasks) SetTaskName(v string) *UpdateJobTemplateRequestJobTemplateConfigTasks {
	s.TaskName = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasks) SetTaskSpec(v *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec) *UpdateJobTemplateRequestJobTemplateConfigTasks {
	s.TaskSpec = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasks) SetTaskSustainable(v bool) *UpdateJobTemplateRequestJobTemplateConfigTasks {
	s.TaskSustainable = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasks) Validate() error {
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

type UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy struct {
	// The array job details.
	ArraySpec *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec `json:"ArraySpec,omitempty" xml:"ArraySpec,omitempty" type:"Struct"`
	// The maximum number of nodes for job execution.
	//
	// example:
	//
	// 1
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy) GetArraySpec() *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec {
	return s.ArraySpec
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy) SetArraySpec(v *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy {
	s.ArraySpec = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy) SetMaxCount(v int32) *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy {
	s.MaxCount = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicy) Validate() error {
	if s.ArraySpec != nil {
		if err := s.ArraySpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec struct {
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

func (s UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) GetIndexEnd() *int32 {
	return s.IndexEnd
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) GetIndexStart() *int32 {
	return s.IndexStart
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) GetIndexStep() *int32 {
	return s.IndexStep
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) SetIndexEnd(v int32) *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec {
	s.IndexEnd = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) SetIndexStart(v int32) *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec {
	s.IndexStart = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) SetIndexStep(v int32) *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec {
	s.IndexStep = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksExecutorPolicyArraySpec) Validate() error {
	return dara.Validate(s)
}

type UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec struct {
	// The resource information of the runtime environment.
	Resource *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// The task retry policy.
	RetryPolicy *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty" type:"Struct"`
	// The task execution configuration.
	TaskExecutor []*UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor `json:"TaskExecutor,omitempty" xml:"TaskExecutor,omitempty" type:"Repeated"`
	// The list of data volumes mounted to the task. You can mount up to 10 data volumes.
	VolumeMount []*UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec) GetResource() *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource {
	return s.Resource
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec) GetRetryPolicy() *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy {
	return s.RetryPolicy
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec) GetTaskExecutor() []*UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor {
	return s.TaskExecutor
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec) GetVolumeMount() []*UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount {
	return s.VolumeMount
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec) SetResource(v *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec {
	s.Resource = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec) SetRetryPolicy(v *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec {
	s.RetryPolicy = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec) SetTaskExecutor(v []*UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec {
	s.TaskExecutor = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec) SetVolumeMount(v []*UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec {
	s.VolumeMount = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpec) Validate() error {
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

type UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource struct {
	// The number of CPUs in the runtime environment.
	//
	// example:
	//
	// 2
	Cores *float32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The cloud disk array.
	Disks []*UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
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

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) GetCores() *float32 {
	return s.Cores
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) GetDisks() []*UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks {
	return s.Disks
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) GetEnableHT() *bool {
	return s.EnableHT
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) GetHostNamePrefix() *string {
	return s.HostNamePrefix
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) GetMemory() *float32 {
	return s.Memory
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) SetCores(v float32) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource {
	s.Cores = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) SetDisks(v []*UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource {
	s.Disks = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) SetEnableHT(v bool) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource {
	s.EnableHT = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) SetHostNamePrefix(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource {
	s.HostNamePrefix = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) SetInstanceTypes(v []*string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource {
	s.InstanceTypes = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) SetMemory(v float32) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource {
	s.Memory = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResource) Validate() error {
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

type UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks struct {
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

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks) GetSize() *int32 {
	return s.Size
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks) GetType() *string {
	return s.Type
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks) SetSize(v int32) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks {
	s.Size = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks) SetType(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks {
	s.Type = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecResourceDisks) Validate() error {
	return dara.Validate(s)
}

type UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy struct {
	// The retry rules.
	ExitCodeActions []*UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions `json:"ExitCodeActions,omitempty" xml:"ExitCodeActions,omitempty" type:"Repeated"`
	// The number of retries.
	//
	// example:
	//
	// 5
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy) GetExitCodeActions() []*UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions {
	return s.ExitCodeActions
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy) SetExitCodeActions(v []*UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy {
	s.ExitCodeActions = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy) SetRetryCount(v int32) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy {
	s.RetryCount = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicy) Validate() error {
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

type UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions struct {
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
	// The task exit code. This parameter is used together with Action to form the job retry rule. Valid values: 0 to 255.
	//
	// example:
	//
	// 1
	ExitCode *int64 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) GetAction() *string {
	return s.Action
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) GetExitCode() *int64 {
	return s.ExitCode
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) SetAction(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions {
	s.Action = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) SetExitCode(v int64) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions {
	s.ExitCode = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) Validate() error {
	return dara.Validate(s)
}

type UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor struct {
	// The container environment.
	Container *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// The virtual machine environment.
	VM *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM `json:"VM,omitempty" xml:"VM,omitempty" type:"Struct"`
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor) GetContainer() *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	return s.Container
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor) GetVM() *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	return s.VM
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor) SetContainer(v *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor {
	s.Container = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor) SetVM(v *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor {
	s.VM = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutor) Validate() error {
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

type UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer struct {
	// The ID of the container application.
	//
	// example:
	//
	// ci-vm-32k6LXAi3cOG
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The arguments for the container startup command.
	//
	// example:
	//
	// 60s
	Arg []*string `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	// The list of container startup commands.
	//
	// example:
	//
	// sleep
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// The environment variables of the container.
	EnvironmentVars []*UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
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
	// The list of mount parameters for the self-managed image repository. The parameters are in key-value pair format and passed in JSON.
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

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetAppId() *string {
	return s.AppId
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetArg() []*string {
	return s.Arg
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetCommand() []*string {
	return s.Command
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetEnvironmentVars() []*UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars {
	return s.EnvironmentVars
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetGpuDriverVersion() *string {
	return s.GpuDriverVersion
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetImage() *string {
	return s.Image
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetImageRegistryOptions() *string {
	return s.ImageRegistryOptions
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetPrivileged() *bool {
	return s.Privileged
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetAppId(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.AppId = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetArg(v []*string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.Arg = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetCommand(v []*string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.Command = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetEnvironmentVars(v []*UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.EnvironmentVars = v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetGpuDriverVersion(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.GpuDriverVersion = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetImage(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.Image = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetImageRegistryOptions(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.ImageRegistryOptions = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetPrivileged(v bool) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.Privileged = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetWorkingDir(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.WorkingDir = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainer) Validate() error {
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

type UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars struct {
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

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) GetName() *string {
	return s.Name
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) GetValue() *string {
	return s.Value
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) SetName(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars {
	s.Name = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) SetValue(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars {
	s.Value = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) Validate() error {
	return dara.Validate(s)
}

type UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM struct {
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
	// The preprocessing script. The script must be Base64-encoded.
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

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) GetAppId() *string {
	return s.AppId
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) GetImage() *string {
	return s.Image
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) GetPassword() *string {
	return s.Password
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) GetPrologScript() *string {
	return s.PrologScript
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) GetScript() *string {
	return s.Script
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) SetAppId(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	s.AppId = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) SetImage(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	s.Image = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) SetPassword(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	s.Password = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) SetPrologScript(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	s.PrologScript = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) SetScript(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	s.Script = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecTaskExecutorVM) Validate() error {
	return dara.Validate(s)
}

type UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount struct {
	// The list of mount options for the data volume. The options are passed as key-value pairs in JSON format.
	//
	// example:
	//
	// {"server":"xxxxx-xxxxx.cn-heyuan.nas.aliyuncs.com","vers":"3","path":"/data","options":"nolock,tcp,noresvport"}
	MountOptions *string `json:"MountOptions,omitempty" xml:"MountOptions,omitempty"`
	// The directory where the data volume is mounted for the task.
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
	// The supported data volume type.
	//
	// example:
	//
	// alicloud/nas
	VolumeDriver *string `json:"VolumeDriver,omitempty" xml:"VolumeDriver,omitempty"`
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) GoString() string {
	return s.String()
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) GetMountOptions() *string {
	return s.MountOptions
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) GetMountPath() *string {
	return s.MountPath
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) GetVolumeDriver() *string {
	return s.VolumeDriver
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) SetMountOptions(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount {
	s.MountOptions = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) SetMountPath(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount {
	s.MountPath = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) SetReadOnly(v bool) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) SetVolumeDriver(v string) *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount {
	s.VolumeDriver = &v
	return s
}

func (s *UpdateJobTemplateRequestJobTemplateConfigTasksTaskSpecVolumeMount) Validate() error {
	return dara.Validate(s)
}
