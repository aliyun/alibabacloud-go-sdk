// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetJobTemplateResponseBody
	GetCreateTime() *string
	SetJobTemplateConfig(v *GetJobTemplateResponseBodyJobTemplateConfig) *GetJobTemplateResponseBody
	GetJobTemplateConfig() *GetJobTemplateResponseBodyJobTemplateConfig
	SetJobTemplateDescription(v string) *GetJobTemplateResponseBody
	GetJobTemplateDescription() *string
	SetJobTemplateId(v string) *GetJobTemplateResponseBody
	GetJobTemplateId() *string
	SetJobTemplateName(v string) *GetJobTemplateResponseBody
	GetJobTemplateName() *string
	SetRequestId(v string) *GetJobTemplateResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetJobTemplateResponseBody
	GetStatus() *string
	SetUpdateTime(v string) *GetJobTemplateResponseBody
	GetUpdateTime() *string
}

type GetJobTemplateResponseBody struct {
	// The creation time. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2026-04-20 11:09:59
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The template configuration.
	JobTemplateConfig *GetJobTemplateResponseBodyJobTemplateConfig `json:"JobTemplateConfig,omitempty" xml:"JobTemplateConfig,omitempty" type:"Struct"`
	// The template description.
	//
	// example:
	//
	// Demo
	JobTemplateDescription *string `json:"JobTemplateDescription,omitempty" xml:"JobTemplateDescription,omitempty"`
	// The job template ID.
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
	// The request ID.
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the job template.
	//
	// example:
	//
	// Working
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The update time. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2026-04-20 11:09:59
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetJobTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetJobTemplateResponseBody) GetJobTemplateConfig() *GetJobTemplateResponseBodyJobTemplateConfig {
	return s.JobTemplateConfig
}

func (s *GetJobTemplateResponseBody) GetJobTemplateDescription() *string {
	return s.JobTemplateDescription
}

func (s *GetJobTemplateResponseBody) GetJobTemplateId() *string {
	return s.JobTemplateId
}

func (s *GetJobTemplateResponseBody) GetJobTemplateName() *string {
	return s.JobTemplateName
}

func (s *GetJobTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobTemplateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetJobTemplateResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetJobTemplateResponseBody) SetCreateTime(v string) *GetJobTemplateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetJobTemplateConfig(v *GetJobTemplateResponseBodyJobTemplateConfig) *GetJobTemplateResponseBody {
	s.JobTemplateConfig = v
	return s
}

func (s *GetJobTemplateResponseBody) SetJobTemplateDescription(v string) *GetJobTemplateResponseBody {
	s.JobTemplateDescription = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetJobTemplateId(v string) *GetJobTemplateResponseBody {
	s.JobTemplateId = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetJobTemplateName(v string) *GetJobTemplateResponseBody {
	s.JobTemplateName = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetRequestId(v string) *GetJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetStatus(v string) *GetJobTemplateResponseBody {
	s.Status = &v
	return s
}

func (s *GetJobTemplateResponseBody) SetUpdateTime(v string) *GetJobTemplateResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetJobTemplateResponseBody) Validate() error {
	if s.JobTemplateConfig != nil {
		if err := s.JobTemplateConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobTemplateResponseBodyJobTemplateConfig struct {
	// The resource deployment policy.
	DeploymentPolicy *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy `json:"DeploymentPolicy,omitempty" xml:"DeploymentPolicy,omitempty" type:"Struct"`
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
	SecurityPolicy *GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicy `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty" type:"Struct"`
	// The tasks.
	Tasks []*GetJobTemplateResponseBodyJobTemplateConfigTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s GetJobTemplateResponseBodyJobTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfig) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfig) GetDeploymentPolicy() *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy {
	return s.DeploymentPolicy
}

func (s *GetJobTemplateResponseBodyJobTemplateConfig) GetJobDescription() *string {
	return s.JobDescription
}

func (s *GetJobTemplateResponseBodyJobTemplateConfig) GetJobName() *string {
	return s.JobName
}

func (s *GetJobTemplateResponseBodyJobTemplateConfig) GetJobScheduler() *string {
	return s.JobScheduler
}

func (s *GetJobTemplateResponseBodyJobTemplateConfig) GetSecurityPolicy() *GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicy {
	return s.SecurityPolicy
}

func (s *GetJobTemplateResponseBodyJobTemplateConfig) GetTasks() []*GetJobTemplateResponseBodyJobTemplateConfigTasks {
	return s.Tasks
}

func (s *GetJobTemplateResponseBodyJobTemplateConfig) SetDeploymentPolicy(v *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy) *GetJobTemplateResponseBodyJobTemplateConfig {
	s.DeploymentPolicy = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfig) SetJobDescription(v string) *GetJobTemplateResponseBodyJobTemplateConfig {
	s.JobDescription = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfig) SetJobName(v string) *GetJobTemplateResponseBodyJobTemplateConfig {
	s.JobName = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfig) SetJobScheduler(v string) *GetJobTemplateResponseBodyJobTemplateConfig {
	s.JobScheduler = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfig) SetSecurityPolicy(v *GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicy) *GetJobTemplateResponseBodyJobTemplateConfig {
	s.SecurityPolicy = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfig) SetTasks(v []*GetJobTemplateResponseBodyJobTemplateConfigTasks) *GetJobTemplateResponseBodyJobTemplateConfig {
	s.Tasks = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfig) Validate() error {
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

type GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy struct {
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
	// example:
	//
	// General
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The network configuration.
	Network *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
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
	// The tags of the job.
	Tags []*GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy) GetAllocationSpec() *string {
	return s.AllocationSpec
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy) GetLevel() *string {
	return s.Level
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy) GetNetwork() *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyNetwork {
	return s.Network
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy) GetPool() *string {
	return s.Pool
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy) GetPriority() *int32 {
	return s.Priority
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy) GetTags() []*GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyTags {
	return s.Tags
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy) SetAllocationSpec(v string) *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy {
	s.AllocationSpec = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy) SetLevel(v string) *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy {
	s.Level = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy) SetNetwork(v *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyNetwork) *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy {
	s.Network = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy) SetPool(v string) *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy {
	s.Pool = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy) SetPriority(v int32) *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy {
	s.Priority = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy) SetTags(v []*GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyTags) *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy {
	s.Tags = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicy) Validate() error {
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

type GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyNetwork struct {
	// Specifies whether to create a public IP address for the job. Default value: false.
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

func (s GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyNetwork) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyNetwork) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyNetwork) GetEnableExternalIpAddress() *bool {
	return s.EnableExternalIpAddress
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyNetwork) GetVswitch() []*string {
	return s.Vswitch
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyNetwork) SetEnableExternalIpAddress(v bool) *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyNetwork {
	s.EnableExternalIpAddress = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyNetwork) SetVswitch(v []*string) *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyNetwork {
	s.Vswitch = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyNetwork) Validate() error {
	return dara.Validate(s)
}

type GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyTags) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyTags) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyTags) GetKey() *string {
	return s.Key
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyTags) GetValue() *string {
	return s.Value
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyTags) SetKey(v string) *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyTags {
	s.Key = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyTags) SetValue(v string) *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyTags {
	s.Value = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigDeploymentPolicyTags) Validate() error {
	return dara.Validate(s)
}

type GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicy struct {
	// The security group.
	SecurityGroup *GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicySecurityGroup `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty" type:"Struct"`
}

func (s GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicy) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicy) GetSecurityGroup() *GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicySecurityGroup {
	return s.SecurityGroup
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicy) SetSecurityGroup(v *GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicySecurityGroup) *GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicy {
	s.SecurityGroup = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicy) Validate() error {
	if s.SecurityGroup != nil {
		if err := s.SecurityGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicySecurityGroup struct {
	// The security group IDs.
	//
	// example:
	//
	// sg-xxxx
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicySecurityGroup) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicySecurityGroup) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicySecurityGroup) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicySecurityGroup) SetSecurityGroupIds(v []*string) *GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicySecurityGroup {
	s.SecurityGroupIds = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigSecurityPolicySecurityGroup) Validate() error {
	return dara.Validate(s)
}

type GetJobTemplateResponseBodyJobTemplateConfigTasks struct {
	// The execution policy of the task.
	ExecutorPolicy *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicy `json:"ExecutorPolicy,omitempty" xml:"ExecutorPolicy,omitempty" type:"Struct"`
	// The task name.
	//
	// example:
	//
	// task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task specifications.
	TaskSpec *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec `json:"TaskSpec,omitempty" xml:"TaskSpec,omitempty" type:"Struct"`
	// Indicates whether the job is a long-running job. Valid values: true: The job is a background service job. false: The job is a batch job. Default value: false.
	//
	// example:
	//
	// false
	TaskSustainable *bool `json:"TaskSustainable,omitempty" xml:"TaskSustainable,omitempty"`
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasks) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasks) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasks) GetExecutorPolicy() *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicy {
	return s.ExecutorPolicy
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasks) GetTaskSpec() *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec {
	return s.TaskSpec
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasks) GetTaskSustainable() *bool {
	return s.TaskSustainable
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasks) SetExecutorPolicy(v *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicy) *GetJobTemplateResponseBodyJobTemplateConfigTasks {
	s.ExecutorPolicy = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasks) SetTaskName(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasks {
	s.TaskName = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasks) SetTaskSpec(v *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec) *GetJobTemplateResponseBodyJobTemplateConfigTasks {
	s.TaskSpec = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasks) SetTaskSustainable(v bool) *GetJobTemplateResponseBodyJobTemplateConfigTasks {
	s.TaskSustainable = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasks) Validate() error {
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

type GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicy struct {
	// The details of the array job.
	ArraySpec *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicyArraySpec `json:"ArraySpec,omitempty" xml:"ArraySpec,omitempty" type:"Struct"`
	// The maximum number of nodes for the job execution.
	//
	// example:
	//
	// 1
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicy) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicy) GetArraySpec() *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicyArraySpec {
	return s.ArraySpec
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicy) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicy) SetArraySpec(v *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicyArraySpec) *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicy {
	s.ArraySpec = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicy) SetMaxCount(v int32) *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicy {
	s.MaxCount = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicy) Validate() error {
	if s.ArraySpec != nil {
		if err := s.ArraySpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicyArraySpec struct {
	// The end index of the array job. Valid values: 0 to 4999. The value must be greater than or equal to the value of IndexStart.
	//
	// example:
	//
	// 9
	IndexEnd *int32 `json:"IndexEnd,omitempty" xml:"IndexEnd,omitempty"`
	// The start index of the array job. Valid values: 0 to 4999.
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

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicyArraySpec) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicyArraySpec) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicyArraySpec) GetIndexEnd() *int32 {
	return s.IndexEnd
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicyArraySpec) GetIndexStart() *int32 {
	return s.IndexStart
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicyArraySpec) GetIndexStep() *int32 {
	return s.IndexStep
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicyArraySpec) SetIndexEnd(v int32) *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicyArraySpec {
	s.IndexEnd = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicyArraySpec) SetIndexStart(v int32) *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicyArraySpec {
	s.IndexStart = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicyArraySpec) SetIndexStep(v int32) *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicyArraySpec {
	s.IndexStep = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksExecutorPolicyArraySpec) Validate() error {
	return dara.Validate(s)
}

type GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec struct {
	// The resource information of the runtime environment.
	Resource *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// The retry policy for the task.
	RetryPolicy *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicy `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty" type:"Struct"`
	// The task execution configuration.
	TaskExecutor []*GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutor `json:"TaskExecutor,omitempty" xml:"TaskExecutor,omitempty" type:"Repeated"`
	// The list of data volumes mounted to the task. A maximum of 10 volumes can be mounted.
	VolumeMount []*GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec) GetResource() *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource {
	return s.Resource
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec) GetRetryPolicy() *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicy {
	return s.RetryPolicy
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec) GetTaskExecutor() []*GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutor {
	return s.TaskExecutor
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec) GetVolumeMount() []*GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount {
	return s.VolumeMount
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec) SetResource(v *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec {
	s.Resource = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec) SetRetryPolicy(v *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicy) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec {
	s.RetryPolicy = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec) SetTaskExecutor(v []*GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutor) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec {
	s.TaskExecutor = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec) SetVolumeMount(v []*GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec {
	s.VolumeMount = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpec) Validate() error {
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

type GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource struct {
	// The number of CPU cores in the runtime environment.
	//
	// example:
	//
	// 2
	Cores *float32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The array of cloud disks.
	Disks []*GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResourceDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// Indicates whether hyper-threading is enabled in the runtime environment.
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

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource) GetCores() *float32 {
	return s.Cores
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource) GetDisks() []*GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResourceDisks {
	return s.Disks
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource) GetEnableHT() *bool {
	return s.EnableHT
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource) GetHostNamePrefix() *string {
	return s.HostNamePrefix
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource) GetMemory() *float32 {
	return s.Memory
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource) SetCores(v float32) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource {
	s.Cores = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource) SetDisks(v []*GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResourceDisks) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource {
	s.Disks = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource) SetEnableHT(v bool) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource {
	s.EnableHT = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource) SetHostNamePrefix(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource {
	s.HostNamePrefix = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource) SetInstanceTypes(v []*string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource {
	s.InstanceTypes = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource) SetMemory(v float32) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource {
	s.Memory = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResource) Validate() error {
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

type GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResourceDisks struct {
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

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResourceDisks) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResourceDisks) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResourceDisks) GetSize() *int32 {
	return s.Size
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResourceDisks) GetType() *string {
	return s.Type
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResourceDisks) SetSize(v int32) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResourceDisks {
	s.Size = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResourceDisks) SetType(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResourceDisks {
	s.Type = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecResourceDisks) Validate() error {
	return dara.Validate(s)
}

type GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicy struct {
	// The retry rules.
	ExitCodeActions []*GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions `json:"ExitCodeActions,omitempty" xml:"ExitCodeActions,omitempty" type:"Repeated"`
	// The number of retries.
	//
	// example:
	//
	// 5
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicy) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicy) GetExitCodeActions() []*GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions {
	return s.ExitCodeActions
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicy) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicy) SetExitCodeActions(v []*GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicy {
	s.ExitCodeActions = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicy) SetRetryCount(v int32) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicy {
	s.RetryCount = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicy) Validate() error {
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

type GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions struct {
	// The next action of the task. Valid values: Retry: retries the task. Exit: exits the task.
	//
	// example:
	//
	// Retry
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The exit code of the task. It is used together with the Action parameter to form the job retry rule. Valid values: 0 to 255.
	//
	// example:
	//
	// 1
	ExitCode *int64 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) GetAction() *string {
	return s.Action
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) GetExitCode() *int64 {
	return s.ExitCode
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) SetAction(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions {
	s.Action = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) SetExitCode(v int64) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions {
	s.ExitCode = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecRetryPolicyExitCodeActions) Validate() error {
	return dara.Validate(s)
}

type GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutor struct {
	// The container environment.
	Container *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	// The virtual machine environment.
	VM *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM `json:"VM,omitempty" xml:"VM,omitempty" type:"Struct"`
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutor) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutor) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutor) GetContainer() *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	return s.Container
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutor) GetVM() *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	return s.VM
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutor) SetContainer(v *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutor {
	s.Container = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutor) SetVM(v *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutor {
	s.VM = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutor) Validate() error {
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

type GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer struct {
	// The application ID of the container.
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
	// The startup commands of the container.
	//
	// example:
	//
	// sleep
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// The environment variables of the container.
	EnvironmentVars []*GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
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
	// The mount parameters for the self-managed image registry. The parameters are in key-value pairs and passed in JSON format.
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

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetAppId() *string {
	return s.AppId
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetArg() []*string {
	return s.Arg
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetCommand() []*string {
	return s.Command
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetEnvironmentVars() []*GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars {
	return s.EnvironmentVars
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetGpuDriverVersion() *string {
	return s.GpuDriverVersion
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetImage() *string {
	return s.Image
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetImageRegistryOptions() *string {
	return s.ImageRegistryOptions
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetPrivileged() *bool {
	return s.Privileged
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetAppId(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.AppId = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetArg(v []*string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.Arg = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetCommand(v []*string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.Command = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetEnvironmentVars(v []*GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.EnvironmentVars = v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetGpuDriverVersion(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.GpuDriverVersion = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetImage(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.Image = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetImageRegistryOptions(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.ImageRegistryOptions = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetPrivileged(v bool) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.Privileged = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) SetWorkingDir(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer {
	s.WorkingDir = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainer) Validate() error {
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

type GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars struct {
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

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) GetName() *string {
	return s.Name
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) GetValue() *string {
	return s.Value
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) SetName(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars {
	s.Name = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) SetValue(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars {
	s.Value = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorContainerEnvironmentVars) Validate() error {
	return dara.Validate(s)
}

type GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM struct {
	// The VM application ID.
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
	// The logon password for the VM environment.
	//
	// example:
	//
	// EHPC@1234
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The prolog script. It must be Base64-encoded.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	PrologScript *string `json:"PrologScript,omitempty" xml:"PrologScript,omitempty"`
	// The job script. It must be Base64-encoded.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM) GetAppId() *string {
	return s.AppId
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM) GetImage() *string {
	return s.Image
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM) GetPassword() *string {
	return s.Password
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM) GetPrologScript() *string {
	return s.PrologScript
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM) GetScript() *string {
	return s.Script
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM) SetAppId(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	s.AppId = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM) SetImage(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	s.Image = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM) SetPassword(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	s.Password = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM) SetPrologScript(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	s.PrologScript = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM) SetScript(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM {
	s.Script = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecTaskExecutorVM) Validate() error {
	return dara.Validate(s)
}

type GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount struct {
	// The list of mount options for the data volume. The options are passed in JSON format as key-value pairs.
	//
	// example:
	//
	// {"server":"xxxxx-xxxxx.cn-heyuan.nas.aliyuncs.com","vers":"3","path":"/data","options":"nolock,tcp,noresvport"}
	MountOptions *string `json:"MountOptions,omitempty" xml:"MountOptions,omitempty"`
	// The mount path of the data volume for the task.
	//
	// example:
	//
	// /mnt
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// Indicates whether the data volume is read-only.
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

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount) GetMountOptions() *string {
	return s.MountOptions
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount) GetMountPath() *string {
	return s.MountPath
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount) GetVolumeDriver() *string {
	return s.VolumeDriver
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount) SetMountOptions(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount {
	s.MountOptions = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount) SetMountPath(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount {
	s.MountPath = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount) SetReadOnly(v bool) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount) SetVolumeDriver(v string) *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount {
	s.VolumeDriver = &v
	return s
}

func (s *GetJobTemplateResponseBodyJobTemplateConfigTasksTaskSpecVolumeMount) Validate() error {
	return dara.Validate(s)
}
