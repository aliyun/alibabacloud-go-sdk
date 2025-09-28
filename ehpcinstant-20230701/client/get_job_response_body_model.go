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
}

type GetJobResponseBody struct {
	JobInfo *GetJobResponseBodyJobInfo `json:"JobInfo,omitempty" xml:"JobInfo,omitempty" type:"Struct"`
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *GetJobResponseBody) SetJobInfo(v *GetJobResponseBodyJobInfo) *GetJobResponseBody {
	s.JobInfo = v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfo struct {
	AppExtraInfo *string `json:"AppExtraInfo,omitempty" xml:"AppExtraInfo,omitempty"`
	// example:
	//
	// 2024-03-05 20:00:46
	CreateTime       *string                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DependencyPolicy *GetJobResponseBodyJobInfoDependencyPolicy `json:"DependencyPolicy,omitempty" xml:"DependencyPolicy,omitempty" type:"Struct"`
	DeploymentPolicy *GetJobResponseBodyJobInfoDeploymentPolicy `json:"DeploymentPolicy,omitempty" xml:"DeploymentPolicy,omitempty" type:"Struct"`
	// example:
	//
	// 2024-03-05 20:01:48
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// Demo
	JobDescription *string `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	// example:
	//
	// job-xxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// testJob
	JobName      *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobScheduler *string `json:"JobScheduler,omitempty" xml:"JobScheduler,omitempty"`
	// example:
	//
	// 2024-03-05 20:00:48
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Succeed
	Status *string                           `json:"Status,omitempty" xml:"Status,omitempty"`
	Tasks  []*GetJobResponseBodyJobInfoTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s GetJobResponseBodyJobInfo) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfo) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfo) GetAppExtraInfo() *string {
	return s.AppExtraInfo
}

func (s *GetJobResponseBodyJobInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetJobResponseBodyJobInfo) GetDependencyPolicy() *GetJobResponseBodyJobInfoDependencyPolicy {
	return s.DependencyPolicy
}

func (s *GetJobResponseBodyJobInfo) GetDeploymentPolicy() *GetJobResponseBodyJobInfoDeploymentPolicy {
	return s.DeploymentPolicy
}

func (s *GetJobResponseBodyJobInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *GetJobResponseBodyJobInfo) GetJobDescription() *string {
	return s.JobDescription
}

func (s *GetJobResponseBodyJobInfo) GetJobId() *string {
	return s.JobId
}

func (s *GetJobResponseBodyJobInfo) GetJobName() *string {
	return s.JobName
}

func (s *GetJobResponseBodyJobInfo) GetJobScheduler() *string {
	return s.JobScheduler
}

func (s *GetJobResponseBodyJobInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *GetJobResponseBodyJobInfo) GetStatus() *string {
	return s.Status
}

func (s *GetJobResponseBodyJobInfo) GetTasks() []*GetJobResponseBodyJobInfoTasks {
	return s.Tasks
}

func (s *GetJobResponseBodyJobInfo) SetAppExtraInfo(v string) *GetJobResponseBodyJobInfo {
	s.AppExtraInfo = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetCreateTime(v string) *GetJobResponseBodyJobInfo {
	s.CreateTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetDependencyPolicy(v *GetJobResponseBodyJobInfoDependencyPolicy) *GetJobResponseBodyJobInfo {
	s.DependencyPolicy = v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetDeploymentPolicy(v *GetJobResponseBodyJobInfoDeploymentPolicy) *GetJobResponseBodyJobInfo {
	s.DeploymentPolicy = v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetEndTime(v string) *GetJobResponseBodyJobInfo {
	s.EndTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetJobDescription(v string) *GetJobResponseBodyJobInfo {
	s.JobDescription = &v
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

func (s *GetJobResponseBodyJobInfo) SetJobScheduler(v string) *GetJobResponseBodyJobInfo {
	s.JobScheduler = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetStartTime(v string) *GetJobResponseBodyJobInfo {
	s.StartTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetStatus(v string) *GetJobResponseBodyJobInfo {
	s.Status = &v
	return s
}

func (s *GetJobResponseBodyJobInfo) SetTasks(v []*GetJobResponseBodyJobInfoTasks) *GetJobResponseBodyJobInfo {
	s.Tasks = v
	return s
}

func (s *GetJobResponseBodyJobInfo) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoDependencyPolicy struct {
	JobDependency []*GetJobResponseBodyJobInfoDependencyPolicyJobDependency `json:"JobDependency,omitempty" xml:"JobDependency,omitempty" type:"Repeated"`
}

func (s GetJobResponseBodyJobInfoDependencyPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoDependencyPolicy) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoDependencyPolicy) GetJobDependency() []*GetJobResponseBodyJobInfoDependencyPolicyJobDependency {
	return s.JobDependency
}

func (s *GetJobResponseBodyJobInfoDependencyPolicy) SetJobDependency(v []*GetJobResponseBodyJobInfoDependencyPolicyJobDependency) *GetJobResponseBodyJobInfoDependencyPolicy {
	s.JobDependency = v
	return s
}

func (s *GetJobResponseBodyJobInfoDependencyPolicy) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoDependencyPolicyJobDependency struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetJobResponseBodyJobInfoDependencyPolicyJobDependency) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoDependencyPolicyJobDependency) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoDependencyPolicyJobDependency) GetJobId() *string {
	return s.JobId
}

func (s *GetJobResponseBodyJobInfoDependencyPolicyJobDependency) GetType() *string {
	return s.Type
}

func (s *GetJobResponseBodyJobInfoDependencyPolicyJobDependency) SetJobId(v string) *GetJobResponseBodyJobInfoDependencyPolicyJobDependency {
	s.JobId = &v
	return s
}

func (s *GetJobResponseBodyJobInfoDependencyPolicyJobDependency) SetType(v string) *GetJobResponseBodyJobInfoDependencyPolicyJobDependency {
	s.Type = &v
	return s
}

func (s *GetJobResponseBodyJobInfoDependencyPolicyJobDependency) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoDeploymentPolicy struct {
	// example:
	//
	// Dedicated
	AllocationSpec *string                                           `json:"AllocationSpec,omitempty" xml:"AllocationSpec,omitempty"`
	Level          *string                                           `json:"Level,omitempty" xml:"Level,omitempty"`
	Network        *GetJobResponseBodyJobInfoDeploymentPolicyNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	Tags           []*GetJobResponseBodyJobInfoDeploymentPolicyTags  `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetJobResponseBodyJobInfoDeploymentPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoDeploymentPolicy) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicy) GetAllocationSpec() *string {
	return s.AllocationSpec
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicy) GetLevel() *string {
	return s.Level
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicy) GetNetwork() *GetJobResponseBodyJobInfoDeploymentPolicyNetwork {
	return s.Network
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicy) GetTags() []*GetJobResponseBodyJobInfoDeploymentPolicyTags {
	return s.Tags
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicy) SetAllocationSpec(v string) *GetJobResponseBodyJobInfoDeploymentPolicy {
	s.AllocationSpec = &v
	return s
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicy) SetLevel(v string) *GetJobResponseBodyJobInfoDeploymentPolicy {
	s.Level = &v
	return s
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicy) SetNetwork(v *GetJobResponseBodyJobInfoDeploymentPolicyNetwork) *GetJobResponseBodyJobInfoDeploymentPolicy {
	s.Network = v
	return s
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicy) SetTags(v []*GetJobResponseBodyJobInfoDeploymentPolicyTags) *GetJobResponseBodyJobInfoDeploymentPolicy {
	s.Tags = v
	return s
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicy) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoDeploymentPolicyNetwork struct {
	EnableENIMapping        *bool     `json:"EnableENIMapping,omitempty" xml:"EnableENIMapping,omitempty"`
	EnableExternalIpAddress *bool     `json:"EnableExternalIpAddress,omitempty" xml:"EnableExternalIpAddress,omitempty"`
	Vswitch                 []*string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty" type:"Repeated"`
}

func (s GetJobResponseBodyJobInfoDeploymentPolicyNetwork) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoDeploymentPolicyNetwork) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyNetwork) GetEnableENIMapping() *bool {
	return s.EnableENIMapping
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyNetwork) GetEnableExternalIpAddress() *bool {
	return s.EnableExternalIpAddress
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyNetwork) GetVswitch() []*string {
	return s.Vswitch
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyNetwork) SetEnableENIMapping(v bool) *GetJobResponseBodyJobInfoDeploymentPolicyNetwork {
	s.EnableENIMapping = &v
	return s
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyNetwork) SetEnableExternalIpAddress(v bool) *GetJobResponseBodyJobInfoDeploymentPolicyNetwork {
	s.EnableExternalIpAddress = &v
	return s
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyNetwork) SetVswitch(v []*string) *GetJobResponseBodyJobInfoDeploymentPolicyNetwork {
	s.Vswitch = v
	return s
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyNetwork) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoDeploymentPolicyTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetJobResponseBodyJobInfoDeploymentPolicyTags) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoDeploymentPolicyTags) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyTags) SetTagKey(v string) *GetJobResponseBodyJobInfoDeploymentPolicyTags {
	s.TagKey = &v
	return s
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyTags) SetTagValue(v string) *GetJobResponseBodyJobInfoDeploymentPolicyTags {
	s.TagValue = &v
	return s
}

func (s *GetJobResponseBodyJobInfoDeploymentPolicyTags) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoTasks struct {
	ExecutorPolicy *GetJobResponseBodyJobInfoTasksExecutorPolicy   `json:"ExecutorPolicy,omitempty" xml:"ExecutorPolicy,omitempty" type:"Struct"`
	ExecutorStatus []*GetJobResponseBodyJobInfoTasksExecutorStatus `json:"ExecutorStatus,omitempty" xml:"ExecutorStatus,omitempty" type:"Repeated"`
	// example:
	//
	// task0
	TaskName *string                                 `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskSpec *GetJobResponseBodyJobInfoTasksTaskSpec `json:"TaskSpec,omitempty" xml:"TaskSpec,omitempty" type:"Struct"`
	// example:
	//
	// true
	TaskSustainable *bool `json:"TaskSustainable,omitempty" xml:"TaskSustainable,omitempty"`
}

func (s GetJobResponseBodyJobInfoTasks) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasks) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasks) GetExecutorPolicy() *GetJobResponseBodyJobInfoTasksExecutorPolicy {
	return s.ExecutorPolicy
}

func (s *GetJobResponseBodyJobInfoTasks) GetExecutorStatus() []*GetJobResponseBodyJobInfoTasksExecutorStatus {
	return s.ExecutorStatus
}

func (s *GetJobResponseBodyJobInfoTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *GetJobResponseBodyJobInfoTasks) GetTaskSpec() *GetJobResponseBodyJobInfoTasksTaskSpec {
	return s.TaskSpec
}

func (s *GetJobResponseBodyJobInfoTasks) GetTaskSustainable() *bool {
	return s.TaskSustainable
}

func (s *GetJobResponseBodyJobInfoTasks) SetExecutorPolicy(v *GetJobResponseBodyJobInfoTasksExecutorPolicy) *GetJobResponseBodyJobInfoTasks {
	s.ExecutorPolicy = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasks) SetExecutorStatus(v []*GetJobResponseBodyJobInfoTasksExecutorStatus) *GetJobResponseBodyJobInfoTasks {
	s.ExecutorStatus = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasks) SetTaskName(v string) *GetJobResponseBodyJobInfoTasks {
	s.TaskName = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasks) SetTaskSpec(v *GetJobResponseBodyJobInfoTasksTaskSpec) *GetJobResponseBodyJobInfoTasks {
	s.TaskSpec = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasks) SetTaskSustainable(v bool) *GetJobResponseBodyJobInfoTasks {
	s.TaskSustainable = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasks) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoTasksExecutorPolicy struct {
	ArraySpec *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec `json:"ArraySpec,omitempty" xml:"ArraySpec,omitempty" type:"Struct"`
	// example:
	//
	// 10
	MaxCount *int32 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
}

func (s GetJobResponseBodyJobInfoTasksExecutorPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksExecutorPolicy) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicy) GetArraySpec() *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec {
	return s.ArraySpec
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicy) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicy) SetArraySpec(v *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec) *GetJobResponseBodyJobInfoTasksExecutorPolicy {
	s.ArraySpec = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicy) SetMaxCount(v int32) *GetJobResponseBodyJobInfoTasksExecutorPolicy {
	s.MaxCount = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicy) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec struct {
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

func (s GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec) GetIndexEnd() *int32 {
	return s.IndexEnd
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec) GetIndexStart() *int32 {
	return s.IndexStart
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec) GetIndexStep() *int32 {
	return s.IndexStep
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec) SetIndexEnd(v int32) *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec {
	s.IndexEnd = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec) SetIndexStart(v int32) *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec {
	s.IndexStart = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec) SetIndexStep(v int32) *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec {
	s.IndexStep = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorPolicyArraySpec) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoTasksExecutorStatus struct {
	// example:
	//
	// 0
	ArrayId *int32 `json:"ArrayId,omitempty" xml:"ArrayId,omitempty"`
	// example:
	//
	// 2024-02-04 13:54:10
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-02-04 13:54:10
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2024-02-04 13:54:10
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Creating executor
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s GetJobResponseBodyJobInfoTasksExecutorStatus) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksExecutorStatus) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) GetArrayId() *int32 {
	return s.ArrayId
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) GetEndTime() *string {
	return s.EndTime
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) GetStartTime() *string {
	return s.StartTime
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) GetStatus() *string {
	return s.Status
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) SetArrayId(v int32) *GetJobResponseBodyJobInfoTasksExecutorStatus {
	s.ArrayId = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) SetCreateTime(v string) *GetJobResponseBodyJobInfoTasksExecutorStatus {
	s.CreateTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) SetEndTime(v string) *GetJobResponseBodyJobInfoTasksExecutorStatus {
	s.EndTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) SetStartTime(v string) *GetJobResponseBodyJobInfoTasksExecutorStatus {
	s.StartTime = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) SetStatus(v string) *GetJobResponseBodyJobInfoTasksExecutorStatus {
	s.Status = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) SetStatusReason(v string) *GetJobResponseBodyJobInfoTasksExecutorStatus {
	s.StatusReason = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksExecutorStatus) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoTasksTaskSpec struct {
	Resource     *GetJobResponseBodyJobInfoTasksTaskSpecResource       `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	RetryPolicy  *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicy    `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty" type:"Struct"`
	TaskExecutor []*GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor `json:"TaskExecutor,omitempty" xml:"TaskExecutor,omitempty" type:"Repeated"`
}

func (s GetJobResponseBodyJobInfoTasksTaskSpec) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksTaskSpec) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpec) GetResource() *GetJobResponseBodyJobInfoTasksTaskSpecResource {
	return s.Resource
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpec) GetRetryPolicy() *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicy {
	return s.RetryPolicy
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpec) GetTaskExecutor() []*GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor {
	return s.TaskExecutor
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpec) SetResource(v *GetJobResponseBodyJobInfoTasksTaskSpecResource) *GetJobResponseBodyJobInfoTasksTaskSpec {
	s.Resource = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpec) SetRetryPolicy(v *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicy) *GetJobResponseBodyJobInfoTasksTaskSpec {
	s.RetryPolicy = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpec) SetTaskExecutor(v []*GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor) *GetJobResponseBodyJobInfoTasksTaskSpec {
	s.TaskExecutor = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpec) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoTasksTaskSpecResource struct {
	// example:
	//
	// 1
	Cores          *float32                                               `json:"Cores,omitempty" xml:"Cores,omitempty"`
	Disks          []*GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	EnableHT       *bool                                                  `json:"EnableHT,omitempty" xml:"EnableHT,omitempty"`
	HostNamePrefix *string                                                `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	InstanceTypes  []*string                                              `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 4
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecResource) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecResource) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) GetCores() *float32 {
	return s.Cores
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) GetDisks() []*GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks {
	return s.Disks
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) GetEnableHT() *bool {
	return s.EnableHT
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) GetHostNamePrefix() *string {
	return s.HostNamePrefix
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) GetMemory() *int32 {
	return s.Memory
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) SetCores(v float32) *GetJobResponseBodyJobInfoTasksTaskSpecResource {
	s.Cores = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) SetDisks(v []*GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks) *GetJobResponseBodyJobInfoTasksTaskSpecResource {
	s.Disks = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) SetEnableHT(v bool) *GetJobResponseBodyJobInfoTasksTaskSpecResource {
	s.EnableHT = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) SetHostNamePrefix(v string) *GetJobResponseBodyJobInfoTasksTaskSpecResource {
	s.HostNamePrefix = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) SetInstanceTypes(v []*string) *GetJobResponseBodyJobInfoTasksTaskSpecResource {
	s.InstanceTypes = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) SetMemory(v int32) *GetJobResponseBodyJobInfoTasksTaskSpecResource {
	s.Memory = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResource) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks struct {
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks) GetSize() *int32 {
	return s.Size
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks) GetType() *string {
	return s.Type
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks) SetSize(v int32) *GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks {
	s.Size = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks) SetType(v string) *GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks {
	s.Type = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecResourceDisks) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicy struct {
	ExitCodeActions []*GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicyExitCodeActions `json:"ExitCodeActions,omitempty" xml:"ExitCodeActions,omitempty" type:"Repeated"`
	RetryCount      *int32                                                              `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicy) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicy) GetExitCodeActions() []*GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicyExitCodeActions {
	return s.ExitCodeActions
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicy) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicy) SetExitCodeActions(v []*GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicyExitCodeActions) *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicy {
	s.ExitCodeActions = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicy) SetRetryCount(v int32) *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicy {
	s.RetryCount = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicy) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicyExitCodeActions struct {
	Action   *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ExitCode *int64  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicyExitCodeActions) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicyExitCodeActions) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicyExitCodeActions) GetAction() *string {
	return s.Action
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicyExitCodeActions) GetExitCode() *int64 {
	return s.ExitCode
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicyExitCodeActions) SetAction(v string) *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicyExitCodeActions {
	s.Action = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicyExitCodeActions) SetExitCode(v int64) *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicyExitCodeActions {
	s.ExitCode = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecRetryPolicyExitCodeActions) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor struct {
	VM *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM `json:"VM,omitempty" xml:"VM,omitempty" type:"Struct"`
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor) GetVM() *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM {
	return s.VM
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor) SetVM(v *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor {
	s.VM = v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutor) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM struct {
	// example:
	//
	// m-xxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// ZWNobyAiMTIzNCIgPiBgZGF0ZSArJXNg
	PrologScript *string `json:"PrologScript,omitempty" xml:"PrologScript,omitempty"`
	// example:
	//
	// ZWNobyAiMTIzNCIgPiBgZGF0ZSArJXNg
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) GetImage() *string {
	return s.Image
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) GetPrologScript() *string {
	return s.PrologScript
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) GetScript() *string {
	return s.Script
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) SetImage(v string) *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM {
	s.Image = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) SetPrologScript(v string) *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM {
	s.PrologScript = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) SetScript(v string) *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM {
	s.Script = &v
	return s
}

func (s *GetJobResponseBodyJobInfoTasksTaskSpecTaskExecutorVM) Validate() error {
	return dara.Validate(s)
}
