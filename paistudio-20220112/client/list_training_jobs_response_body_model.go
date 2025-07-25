// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTrainingJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListTrainingJobsResponseBody
	GetTotalCount() *int64
	SetTrainingJobs(v []*ListTrainingJobsResponseBodyTrainingJobs) *ListTrainingJobsResponseBody
	GetTrainingJobs() []*ListTrainingJobsResponseBodyTrainingJobs
}

type ListTrainingJobsResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount   *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TrainingJobs []*ListTrainingJobsResponseBodyTrainingJobs `json:"TrainingJobs,omitempty" xml:"TrainingJobs,omitempty" type:"Repeated"`
}

func (s ListTrainingJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTrainingJobsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTrainingJobsResponseBody) GetTrainingJobs() []*ListTrainingJobsResponseBodyTrainingJobs {
	return s.TrainingJobs
}

func (s *ListTrainingJobsResponseBody) SetRequestId(v string) *ListTrainingJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrainingJobsResponseBody) SetTotalCount(v int64) *ListTrainingJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrainingJobsResponseBody) SetTrainingJobs(v []*ListTrainingJobsResponseBodyTrainingJobs) *ListTrainingJobsResponseBody {
	s.TrainingJobs = v
	return s
}

func (s *ListTrainingJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTrainingJobsResponseBodyTrainingJobs struct {
	// example:
	//
	// llm_train
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// example:
	//
	// v0.0.1
	AlgorithmVersion *string                                                   `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	ComputeResource  *ListTrainingJobsResponseBodyTrainingJobsComputeResource  `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	DlcJobId         *string                                                   `json:"DlcJobId,omitempty" xml:"DlcJobId,omitempty"`
	Environments     map[string]*string                                        `json:"Environments,omitempty" xml:"Environments,omitempty"`
	ExperimentConfig *ListTrainingJobsResponseBodyTrainingJobsExperimentConfig `json:"ExperimentConfig,omitempty" xml:"ExperimentConfig,omitempty" type:"Struct"`
	// example:
	//
	// 2024-07-10T11:49:47Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2024-07-10T11:49:47Z
	GmtModifiedTime *string                                                    `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	HyperParameters []*ListTrainingJobsResponseBodyTrainingJobsHyperParameters `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Repeated"`
	InputChannels   []*ListTrainingJobsResponseBodyTrainingJobsInputChannels   `json:"InputChannels,omitempty" xml:"InputChannels,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IsTempAlgo         *bool                                                     `json:"IsTempAlgo,omitempty" xml:"IsTempAlgo,omitempty"`
	Labels             []*ListTrainingJobsResponseBodyTrainingJobsLabels         `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	OutputChannels     []*ListTrainingJobsResponseBodyTrainingJobsOutputChannels `json:"OutputChannels,omitempty" xml:"OutputChannels,omitempty" type:"Repeated"`
	PythonRequirements []*string                                                 `json:"PythonRequirements,omitempty" xml:"PythonRequirements,omitempty" type:"Repeated"`
	// example:
	//
	// TrainingJobSucceed
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// None
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// acs:ram::{accountID}:role/{roleName}
	RoleArn   *string                                            `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	Scheduler *ListTrainingJobsResponseBodyTrainingJobsScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
	// example:
	//
	// Running
	Status                 *string                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusTransitions      []*ListTrainingJobsResponseBodyTrainingJobsStatusTransitions `json:"StatusTransitions,omitempty" xml:"StatusTransitions,omitempty" type:"Repeated"`
	TrainingJobDescription *string                                                      `json:"TrainingJobDescription,omitempty" xml:"TrainingJobDescription,omitempty"`
	// example:
	//
	// train1layo6js8ra
	TrainingJobId *string `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
	// example:
	//
	// qwen2-7b
	TrainingJobName *string `json:"TrainingJobName,omitempty" xml:"TrainingJobName,omitempty"`
	// example:
	//
	// 123456789
	UserId  *string                                          `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserVpc *ListTrainingJobsResponseBodyTrainingJobsUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobs) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobs) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetAlgorithmName() *string {
	return s.AlgorithmName
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetAlgorithmProvider() *string {
	return s.AlgorithmProvider
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetAlgorithmVersion() *string {
	return s.AlgorithmVersion
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetComputeResource() *ListTrainingJobsResponseBodyTrainingJobsComputeResource {
	return s.ComputeResource
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetDlcJobId() *string {
	return s.DlcJobId
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetEnvironments() map[string]*string {
	return s.Environments
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetExperimentConfig() *ListTrainingJobsResponseBodyTrainingJobsExperimentConfig {
	return s.ExperimentConfig
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetHyperParameters() []*ListTrainingJobsResponseBodyTrainingJobsHyperParameters {
	return s.HyperParameters
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetInputChannels() []*ListTrainingJobsResponseBodyTrainingJobsInputChannels {
	return s.InputChannels
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetIsTempAlgo() *bool {
	return s.IsTempAlgo
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetLabels() []*ListTrainingJobsResponseBodyTrainingJobsLabels {
	return s.Labels
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetOutputChannels() []*ListTrainingJobsResponseBodyTrainingJobsOutputChannels {
	return s.OutputChannels
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetPythonRequirements() []*string {
	return s.PythonRequirements
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetScheduler() *ListTrainingJobsResponseBodyTrainingJobsScheduler {
	return s.Scheduler
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetStatus() *string {
	return s.Status
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetStatusTransitions() []*ListTrainingJobsResponseBodyTrainingJobsStatusTransitions {
	return s.StatusTransitions
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetTrainingJobDescription() *string {
	return s.TrainingJobDescription
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetTrainingJobId() *string {
	return s.TrainingJobId
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetTrainingJobName() *string {
	return s.TrainingJobName
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetUserId() *string {
	return s.UserId
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetUserVpc() *ListTrainingJobsResponseBodyTrainingJobsUserVpc {
	return s.UserVpc
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetAlgorithmName(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.AlgorithmName = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetAlgorithmProvider(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.AlgorithmProvider = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetAlgorithmVersion(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.AlgorithmVersion = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetComputeResource(v *ListTrainingJobsResponseBodyTrainingJobsComputeResource) *ListTrainingJobsResponseBodyTrainingJobs {
	s.ComputeResource = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetDlcJobId(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.DlcJobId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetEnvironments(v map[string]*string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.Environments = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetExperimentConfig(v *ListTrainingJobsResponseBodyTrainingJobsExperimentConfig) *ListTrainingJobsResponseBodyTrainingJobs {
	s.ExperimentConfig = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetGmtCreateTime(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.GmtCreateTime = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetGmtModifiedTime(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetHyperParameters(v []*ListTrainingJobsResponseBodyTrainingJobsHyperParameters) *ListTrainingJobsResponseBodyTrainingJobs {
	s.HyperParameters = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetInputChannels(v []*ListTrainingJobsResponseBodyTrainingJobsInputChannels) *ListTrainingJobsResponseBodyTrainingJobs {
	s.InputChannels = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetIsTempAlgo(v bool) *ListTrainingJobsResponseBodyTrainingJobs {
	s.IsTempAlgo = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetLabels(v []*ListTrainingJobsResponseBodyTrainingJobsLabels) *ListTrainingJobsResponseBodyTrainingJobs {
	s.Labels = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetOutputChannels(v []*ListTrainingJobsResponseBodyTrainingJobsOutputChannels) *ListTrainingJobsResponseBodyTrainingJobs {
	s.OutputChannels = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetPythonRequirements(v []*string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.PythonRequirements = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetReasonCode(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.ReasonCode = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetReasonMessage(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.ReasonMessage = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetRoleArn(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.RoleArn = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetScheduler(v *ListTrainingJobsResponseBodyTrainingJobsScheduler) *ListTrainingJobsResponseBodyTrainingJobs {
	s.Scheduler = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetStatus(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.Status = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetStatusTransitions(v []*ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) *ListTrainingJobsResponseBodyTrainingJobs {
	s.StatusTransitions = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetTrainingJobDescription(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.TrainingJobDescription = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetTrainingJobId(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.TrainingJobId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetTrainingJobName(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.TrainingJobName = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetUserId(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.UserId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetUserVpc(v *ListTrainingJobsResponseBodyTrainingJobsUserVpc) *ListTrainingJobsResponseBodyTrainingJobs {
	s.UserVpc = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) SetWorkspaceId(v string) *ListTrainingJobsResponseBodyTrainingJobs {
	s.WorkspaceId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobs) Validate() error {
	return dara.Validate(s)
}

type ListTrainingJobsResponseBodyTrainingJobsComputeResource struct {
	// example:
	//
	// 1
	EcsCount *int64 `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	// example:
	//
	// ecs.gn5-c8g1.2xlarge
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// example:
	//
	// 1
	InstanceCount *int64                                                               `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	InstanceSpec  *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" type:"Struct"`
	// example:
	//
	// quotam670lixikcl
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsComputeResource) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsComputeResource) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) GetEcsCount() *int64 {
	return s.EcsCount
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) GetInstanceCount() *int64 {
	return s.InstanceCount
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) GetInstanceSpec() *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec {
	return s.InstanceSpec
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) SetEcsCount(v int64) *ListTrainingJobsResponseBodyTrainingJobsComputeResource {
	s.EcsCount = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) SetEcsSpec(v string) *ListTrainingJobsResponseBodyTrainingJobsComputeResource {
	s.EcsSpec = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) SetInstanceCount(v int64) *ListTrainingJobsResponseBodyTrainingJobsComputeResource {
	s.InstanceCount = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) SetInstanceSpec(v *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) *ListTrainingJobsResponseBodyTrainingJobsComputeResource {
	s.InstanceSpec = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) SetResourceId(v string) *ListTrainingJobsResponseBodyTrainingJobsComputeResource {
	s.ResourceId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) SetResourceName(v string) *ListTrainingJobsResponseBodyTrainingJobsComputeResource {
	s.ResourceName = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResource) Validate() error {
	return dara.Validate(s)
}

type ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec struct {
	// example:
	//
	// 8
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 1
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// example:
	//
	// V100
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// 32
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// 32
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) GetCPU() *string {
	return s.CPU
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) GetGPU() *string {
	return s.GPU
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) GetGPUType() *string {
	return s.GPUType
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) GetMemory() *string {
	return s.Memory
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) GetSharedMemory() *string {
	return s.SharedMemory
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) SetCPU(v string) *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec {
	s.CPU = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) SetGPU(v string) *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec {
	s.GPU = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) SetGPUType(v string) *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec {
	s.GPUType = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) SetMemory(v string) *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec {
	s.Memory = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) SetSharedMemory(v string) *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec {
	s.SharedMemory = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsComputeResourceInstanceSpec) Validate() error {
	return dara.Validate(s)
}

type ListTrainingJobsResponseBodyTrainingJobsExperimentConfig struct {
	// example:
	//
	// exp-ds9aefia90v
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// large_language_model
	ExperimentName *string `json:"ExperimentName,omitempty" xml:"ExperimentName,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsExperimentConfig) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsExperimentConfig) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsExperimentConfig) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *ListTrainingJobsResponseBodyTrainingJobsExperimentConfig) GetExperimentName() *string {
	return s.ExperimentName
}

func (s *ListTrainingJobsResponseBodyTrainingJobsExperimentConfig) SetExperimentId(v string) *ListTrainingJobsResponseBodyTrainingJobsExperimentConfig {
	s.ExperimentId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsExperimentConfig) SetExperimentName(v string) *ListTrainingJobsResponseBodyTrainingJobsExperimentConfig {
	s.ExperimentName = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsExperimentConfig) Validate() error {
	return dara.Validate(s)
}

type ListTrainingJobsResponseBodyTrainingJobsHyperParameters struct {
	// example:
	//
	// learning_rate
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.001
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsHyperParameters) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsHyperParameters) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsHyperParameters) GetName() *string {
	return s.Name
}

func (s *ListTrainingJobsResponseBodyTrainingJobsHyperParameters) GetValue() *string {
	return s.Value
}

func (s *ListTrainingJobsResponseBodyTrainingJobsHyperParameters) SetName(v string) *ListTrainingJobsResponseBodyTrainingJobsHyperParameters {
	s.Name = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsHyperParameters) SetValue(v string) *ListTrainingJobsResponseBodyTrainingJobsHyperParameters {
	s.Value = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsHyperParameters) Validate() error {
	return dara.Validate(s)
}

type ListTrainingJobsResponseBodyTrainingJobsInputChannels struct {
	// example:
	//
	// d-475megosidivjfgfq6
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou-internal.aliyuncs.com/path/to/input/channel/
	InputUri *string `json:"InputUri,omitempty" xml:"InputUri,omitempty"`
	// example:
	//
	// model
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsInputChannels) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsInputChannels) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsInputChannels) GetDatasetId() *string {
	return s.DatasetId
}

func (s *ListTrainingJobsResponseBodyTrainingJobsInputChannels) GetInputUri() *string {
	return s.InputUri
}

func (s *ListTrainingJobsResponseBodyTrainingJobsInputChannels) GetName() *string {
	return s.Name
}

func (s *ListTrainingJobsResponseBodyTrainingJobsInputChannels) GetVersionName() *string {
	return s.VersionName
}

func (s *ListTrainingJobsResponseBodyTrainingJobsInputChannels) SetDatasetId(v string) *ListTrainingJobsResponseBodyTrainingJobsInputChannels {
	s.DatasetId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsInputChannels) SetInputUri(v string) *ListTrainingJobsResponseBodyTrainingJobsInputChannels {
	s.InputUri = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsInputChannels) SetName(v string) *ListTrainingJobsResponseBodyTrainingJobsInputChannels {
	s.Name = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsInputChannels) SetVersionName(v string) *ListTrainingJobsResponseBodyTrainingJobsInputChannels {
	s.VersionName = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsInputChannels) Validate() error {
	return dara.Validate(s)
}

type ListTrainingJobsResponseBodyTrainingJobsLabels struct {
	// example:
	//
	// CreatedBy
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// QuickStart
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsLabels) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsLabels) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsLabels) GetKey() *string {
	return s.Key
}

func (s *ListTrainingJobsResponseBodyTrainingJobsLabels) GetValue() *string {
	return s.Value
}

func (s *ListTrainingJobsResponseBodyTrainingJobsLabels) SetKey(v string) *ListTrainingJobsResponseBodyTrainingJobsLabels {
	s.Key = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsLabels) SetValue(v string) *ListTrainingJobsResponseBodyTrainingJobsLabels {
	s.Value = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsLabels) Validate() error {
	return dara.Validate(s)
}

type ListTrainingJobsResponseBodyTrainingJobsOutputChannels struct {
	// example:
	//
	// d-8o0hh35po15ejcdq2p
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// model
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou-internal.aliyuncs.com/path/to/output/channel/
	OutputUri   *string `json:"OutputUri,omitempty" xml:"OutputUri,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsOutputChannels) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsOutputChannels) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsOutputChannels) GetDatasetId() *string {
	return s.DatasetId
}

func (s *ListTrainingJobsResponseBodyTrainingJobsOutputChannels) GetName() *string {
	return s.Name
}

func (s *ListTrainingJobsResponseBodyTrainingJobsOutputChannels) GetOutputUri() *string {
	return s.OutputUri
}

func (s *ListTrainingJobsResponseBodyTrainingJobsOutputChannels) GetVersionName() *string {
	return s.VersionName
}

func (s *ListTrainingJobsResponseBodyTrainingJobsOutputChannels) SetDatasetId(v string) *ListTrainingJobsResponseBodyTrainingJobsOutputChannels {
	s.DatasetId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsOutputChannels) SetName(v string) *ListTrainingJobsResponseBodyTrainingJobsOutputChannels {
	s.Name = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsOutputChannels) SetOutputUri(v string) *ListTrainingJobsResponseBodyTrainingJobsOutputChannels {
	s.OutputUri = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsOutputChannels) SetVersionName(v string) *ListTrainingJobsResponseBodyTrainingJobsOutputChannels {
	s.VersionName = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsOutputChannels) Validate() error {
	return dara.Validate(s)
}

type ListTrainingJobsResponseBodyTrainingJobsScheduler struct {
	// example:
	//
	// 0
	MaxRunningTimeInSeconds *int64 `json:"MaxRunningTimeInSeconds,omitempty" xml:"MaxRunningTimeInSeconds,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsScheduler) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsScheduler) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsScheduler) GetMaxRunningTimeInSeconds() *int64 {
	return s.MaxRunningTimeInSeconds
}

func (s *ListTrainingJobsResponseBodyTrainingJobsScheduler) SetMaxRunningTimeInSeconds(v int64) *ListTrainingJobsResponseBodyTrainingJobsScheduler {
	s.MaxRunningTimeInSeconds = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsScheduler) Validate() error {
	return dara.Validate(s)
}

type ListTrainingJobsResponseBodyTrainingJobsStatusTransitions struct {
	// example:
	//
	// 2024-07-10T11:49:47Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// TrainingJobSucceed
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// KubeDL job runs successfully
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// 2024-07-10T11:49:47Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) GetStatus() *string {
	return s.Status
}

func (s *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) SetEndTime(v string) *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions {
	s.EndTime = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) SetReasonCode(v string) *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions {
	s.ReasonCode = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) SetReasonMessage(v string) *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions {
	s.ReasonMessage = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) SetStartTime(v string) *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions {
	s.StartTime = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) SetStatus(v string) *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions {
	s.Status = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsStatusTransitions) Validate() error {
	return dara.Validate(s)
}

type ListTrainingJobsResponseBodyTrainingJobsUserVpc struct {
	// example:
	//
	// eth1
	DefaultRoute  *string   `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	ExtendedCIDRs []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	// example:
	//
	// sg-abcdef****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// vs-abcdef****
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-abcdef****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListTrainingJobsResponseBodyTrainingJobsUserVpc) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobsResponseBodyTrainingJobsUserVpc) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponseBodyTrainingJobsUserVpc) GetDefaultRoute() *string {
	return s.DefaultRoute
}

func (s *ListTrainingJobsResponseBodyTrainingJobsUserVpc) GetExtendedCIDRs() []*string {
	return s.ExtendedCIDRs
}

func (s *ListTrainingJobsResponseBodyTrainingJobsUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListTrainingJobsResponseBodyTrainingJobsUserVpc) GetSwitchId() *string {
	return s.SwitchId
}

func (s *ListTrainingJobsResponseBodyTrainingJobsUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *ListTrainingJobsResponseBodyTrainingJobsUserVpc) SetDefaultRoute(v string) *ListTrainingJobsResponseBodyTrainingJobsUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsUserVpc) SetExtendedCIDRs(v []*string) *ListTrainingJobsResponseBodyTrainingJobsUserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsUserVpc) SetSecurityGroupId(v string) *ListTrainingJobsResponseBodyTrainingJobsUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsUserVpc) SetSwitchId(v string) *ListTrainingJobsResponseBodyTrainingJobsUserVpc {
	s.SwitchId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsUserVpc) SetVpcId(v string) *ListTrainingJobsResponseBodyTrainingJobsUserVpc {
	s.VpcId = &v
	return s
}

func (s *ListTrainingJobsResponseBodyTrainingJobsUserVpc) Validate() error {
	return dara.Validate(s)
}
