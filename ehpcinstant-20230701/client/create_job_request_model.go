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
	DependencyPolicy *CreateJobRequestDependencyPolicy `json:"DependencyPolicy,omitempty" xml:"DependencyPolicy,omitempty" type:"Struct"`
	DeploymentPolicy *CreateJobRequestDeploymentPolicy `json:"DeploymentPolicy,omitempty" xml:"DeploymentPolicy,omitempty" type:"Struct"`
	// example:
	//
	// Demo
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testjob
	JobName        *string                         `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobScheduler   *string                         `json:"JobScheduler,omitempty" xml:"JobScheduler,omitempty"`
	SecurityPolicy *CreateJobRequestSecurityPolicy `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type CreateJobRequestDependencyPolicy struct {
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
	return dara.Validate(s)
}

type CreateJobRequestDependencyPolicyJobDependency struct {
	// This parameter is required.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// example:
	//
	// Dedicated
	AllocationSpec *string                                  `json:"AllocationSpec,omitempty" xml:"AllocationSpec,omitempty"`
	Level          *string                                  `json:"Level,omitempty" xml:"Level,omitempty"`
	Network        *CreateJobRequestDeploymentPolicyNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	Pool           *string                                  `json:"Pool,omitempty" xml:"Pool,omitempty"`
	Priority       *int32                                   `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Tag            []*CreateJobRequestDeploymentPolicyTag   `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type CreateJobRequestDeploymentPolicyNetwork struct {
	EnableExternalIpAddress *bool     `json:"EnableExternalIpAddress,omitempty" xml:"EnableExternalIpAddress,omitempty"`
	Vswitch                 []*string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty" type:"Repeated"`
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
	// This parameter is required.
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	return dara.Validate(s)
}

type CreateJobRequestSecurityPolicySecurityGroup struct {
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
	ExecutorPolicy *CreateJobRequestTasksExecutorPolicy `json:"ExecutorPolicy,omitempty" xml:"ExecutorPolicy,omitempty" type:"Struct"`
	// example:
	//
	// task0
	TaskName *string                        `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskSpec *CreateJobRequestTasksTaskSpec `json:"TaskSpec,omitempty" xml:"TaskSpec,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type CreateJobRequestTasksExecutorPolicy struct {
	ArraySpec *CreateJobRequestTasksExecutorPolicyArraySpec `json:"ArraySpec,omitempty" xml:"ArraySpec,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type CreateJobRequestTasksExecutorPolicyArraySpec struct {
	// example:
	//
	// 9
	IndexEnd *int32 `json:"IndexEnd,omitempty" xml:"IndexEnd,omitempty"`
	// example:
	//
	// 0
	IndexStart *int32 `json:"IndexStart,omitempty" xml:"IndexStart,omitempty"`
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
	Resource    *CreateJobRequestTasksTaskSpecResource    `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	RetryPolicy *CreateJobRequestTasksTaskSpecRetryPolicy `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty" type:"Struct"`
	// This parameter is required.
	TaskExecutor []*CreateJobRequestTasksTaskSpecTaskExecutor `json:"TaskExecutor,omitempty" xml:"TaskExecutor,omitempty" type:"Repeated"`
	VolumeMount  []*CreateJobRequestTasksTaskSpecVolumeMount  `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type CreateJobRequestTasksTaskSpecResource struct {
	// example:
	//
	// 2
	Cores         *float32                                      `json:"Cores,omitempty" xml:"Cores,omitempty"`
	Disks         []*CreateJobRequestTasksTaskSpecResourceDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	InstanceTypes []*string                                     `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
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

func (s *CreateJobRequestTasksTaskSpecResource) SetInstanceTypes(v []*string) *CreateJobRequestTasksTaskSpecResource {
	s.InstanceTypes = v
	return s
}

func (s *CreateJobRequestTasksTaskSpecResource) SetMemory(v float32) *CreateJobRequestTasksTaskSpecResource {
	s.Memory = &v
	return s
}

func (s *CreateJobRequestTasksTaskSpecResource) Validate() error {
	return dara.Validate(s)
}

type CreateJobRequestTasksTaskSpecResourceDisks struct {
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
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
	ExitCodeActions []*CreateJobRequestTasksTaskSpecRetryPolicyExitCodeActions `json:"ExitCodeActions,omitempty" xml:"ExitCodeActions,omitempty" type:"Repeated"`
	RetryCount      *int32                                                     `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
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
	return dara.Validate(s)
}

type CreateJobRequestTasksTaskSpecRetryPolicyExitCodeActions struct {
	// This parameter is required.
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// This parameter is required.
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
	Container *CreateJobRequestTasksTaskSpecTaskExecutorContainer `json:"Container,omitempty" xml:"Container,omitempty" type:"Struct"`
	VM        *CreateJobRequestTasksTaskSpecTaskExecutorVM        `json:"VM,omitempty" xml:"VM,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type CreateJobRequestTasksTaskSpecTaskExecutorContainer struct {
	AppId           *string                                                              `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Arg             []*string                                                            `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	Command         []*string                                                            `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	EnvironmentVars []*CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/ehpc/hpl:latest
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
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
	return dara.Validate(s)
}

type CreateJobRequestTasksTaskSpecTaskExecutorContainerEnvironmentVars struct {
	// example:
	//
	// PATH
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// m-xxxx
	Image    *string `json:"Image,omitempty" xml:"Image,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	PrologScript *string `json:"PrologScript,omitempty" xml:"PrologScript,omitempty"`
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
	// example:
	//
	// {"server":"xxxxx-xxxxx.cn-heyuan.nas.aliyuncs.com","vers":"3","path":"/data","options":"nolock,tcp,noresvport"}
	MountOptions *string `json:"MountOptions,omitempty" xml:"MountOptions,omitempty"`
	// example:
	//
	// /mnt
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	ReadOnly  *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
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
