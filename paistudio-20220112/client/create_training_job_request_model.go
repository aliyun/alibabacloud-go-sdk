// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrainingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmName(v string) *CreateTrainingJobRequest
	GetAlgorithmName() *string
	SetAlgorithmProvider(v string) *CreateTrainingJobRequest
	GetAlgorithmProvider() *string
	SetAlgorithmSpec(v *AlgorithmSpec) *CreateTrainingJobRequest
	GetAlgorithmSpec() *AlgorithmSpec
	SetAlgorithmVersion(v string) *CreateTrainingJobRequest
	GetAlgorithmVersion() *string
	SetAssignNodeSpec(v *AssignNodeSpec) *CreateTrainingJobRequest
	GetAssignNodeSpec() *AssignNodeSpec
	SetCodeDir(v *Location) *CreateTrainingJobRequest
	GetCodeDir() *Location
	SetComputeResource(v *CreateTrainingJobRequestComputeResource) *CreateTrainingJobRequest
	GetComputeResource() *CreateTrainingJobRequestComputeResource
	SetCredentialConfig(v *CredentialConfig) *CreateTrainingJobRequest
	GetCredentialConfig() *CredentialConfig
	SetEnvironments(v map[string]*string) *CreateTrainingJobRequest
	GetEnvironments() map[string]*string
	SetExperimentConfig(v *CreateTrainingJobRequestExperimentConfig) *CreateTrainingJobRequest
	GetExperimentConfig() *CreateTrainingJobRequestExperimentConfig
	SetHyperParameters(v []*CreateTrainingJobRequestHyperParameters) *CreateTrainingJobRequest
	GetHyperParameters() []*CreateTrainingJobRequestHyperParameters
	SetInputChannels(v []*CreateTrainingJobRequestInputChannels) *CreateTrainingJobRequest
	GetInputChannels() []*CreateTrainingJobRequestInputChannels
	SetLabels(v []*CreateTrainingJobRequestLabels) *CreateTrainingJobRequest
	GetLabels() []*CreateTrainingJobRequestLabels
	SetOutputChannels(v []*CreateTrainingJobRequestOutputChannels) *CreateTrainingJobRequest
	GetOutputChannels() []*CreateTrainingJobRequestOutputChannels
	SetPriority(v int32) *CreateTrainingJobRequest
	GetPriority() *int32
	SetPythonRequirements(v []*string) *CreateTrainingJobRequest
	GetPythonRequirements() []*string
	SetRoleArn(v string) *CreateTrainingJobRequest
	GetRoleArn() *string
	SetScheduler(v *CreateTrainingJobRequestScheduler) *CreateTrainingJobRequest
	GetScheduler() *CreateTrainingJobRequestScheduler
	SetSettings(v *JobSettings) *CreateTrainingJobRequest
	GetSettings() *JobSettings
	SetTrainingJobDescription(v string) *CreateTrainingJobRequest
	GetTrainingJobDescription() *string
	SetTrainingJobName(v string) *CreateTrainingJobRequest
	GetTrainingJobName() *string
	SetUserVpc(v *CreateTrainingJobRequestUserVpc) *CreateTrainingJobRequest
	GetUserVpc() *CreateTrainingJobRequestUserVpc
	SetWorkspaceId(v string) *CreateTrainingJobRequest
	GetWorkspaceId() *string
}

type CreateTrainingJobRequest struct {
	// The algorithm name.
	//
	// example:
	//
	// ev_classification
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// The algorithm provider.
	//
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// The algorithm configuration for the training job.
	AlgorithmSpec *AlgorithmSpec `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
	// The algorithm version.
	//
	// example:
	//
	// v1.0.0
	AlgorithmVersion *string         `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	AssignNodeSpec   *AssignNodeSpec `json:"AssignNodeSpec,omitempty" xml:"AssignNodeSpec,omitempty"`
	// The code directory for the training job.
	CodeDir *Location `json:"CodeDir,omitempty" xml:"CodeDir,omitempty"`
	// The compute resource configuration.
	ComputeResource  *CreateTrainingJobRequestComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	CredentialConfig *CredentialConfig                        `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The environment variables for the training job.
	Environments map[string]*string `json:"Environments,omitempty" xml:"Environments,omitempty"`
	// The experiment configuration associated with the training job.
	ExperimentConfig *CreateTrainingJobRequestExperimentConfig `json:"ExperimentConfig,omitempty" xml:"ExperimentConfig,omitempty" type:"Struct"`
	// The training hyperparameter settings.
	HyperParameters []*CreateTrainingJobRequestHyperParameters `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Repeated"`
	// The training input data configuration.
	InputChannels []*CreateTrainingJobRequestInputChannels `json:"InputChannels,omitempty" xml:"InputChannels,omitempty" type:"Repeated"`
	// The training job labels.
	Labels []*CreateTrainingJobRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The training output data configuration.
	OutputChannels []*CreateTrainingJobRequestOutputChannels `json:"OutputChannels,omitempty" xml:"OutputChannels,omitempty" type:"Repeated"`
	// The priority of the training job.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The Python package configuration for the training job.
	PythonRequirements []*string `json:"PythonRequirements,omitempty" xml:"PythonRequirements,omitempty" type:"Repeated"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role. Format: acs:ram::$accountID:role/$roleName.
	//
	// example:
	//
	// acs:ram::1157703270994901:role/aliyunserviceroleforpaiworkspace
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The training job scheduling configuration.
	Scheduler *CreateTrainingJobRequestScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
	// The additional parameter settings for the training node.
	Settings *JobSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// The description of the training job.
	//
	// example:
	//
	// qwen large language model training
	TrainingJobDescription *string `json:"TrainingJobDescription,omitempty" xml:"TrainingJobDescription,omitempty"`
	// The name of the training job.
	//
	// This parameter is required.
	//
	// example:
	//
	// qwen_llm
	TrainingJobName *string `json:"TrainingJobName,omitempty" xml:"TrainingJobName,omitempty"`
	// The VPC configuration.
	UserVpc *CreateTrainingJobRequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateTrainingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainingJobRequest) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequest) GetAlgorithmName() *string {
	return s.AlgorithmName
}

func (s *CreateTrainingJobRequest) GetAlgorithmProvider() *string {
	return s.AlgorithmProvider
}

func (s *CreateTrainingJobRequest) GetAlgorithmSpec() *AlgorithmSpec {
	return s.AlgorithmSpec
}

func (s *CreateTrainingJobRequest) GetAlgorithmVersion() *string {
	return s.AlgorithmVersion
}

func (s *CreateTrainingJobRequest) GetAssignNodeSpec() *AssignNodeSpec {
	return s.AssignNodeSpec
}

func (s *CreateTrainingJobRequest) GetCodeDir() *Location {
	return s.CodeDir
}

func (s *CreateTrainingJobRequest) GetComputeResource() *CreateTrainingJobRequestComputeResource {
	return s.ComputeResource
}

func (s *CreateTrainingJobRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *CreateTrainingJobRequest) GetEnvironments() map[string]*string {
	return s.Environments
}

func (s *CreateTrainingJobRequest) GetExperimentConfig() *CreateTrainingJobRequestExperimentConfig {
	return s.ExperimentConfig
}

func (s *CreateTrainingJobRequest) GetHyperParameters() []*CreateTrainingJobRequestHyperParameters {
	return s.HyperParameters
}

func (s *CreateTrainingJobRequest) GetInputChannels() []*CreateTrainingJobRequestInputChannels {
	return s.InputChannels
}

func (s *CreateTrainingJobRequest) GetLabels() []*CreateTrainingJobRequestLabels {
	return s.Labels
}

func (s *CreateTrainingJobRequest) GetOutputChannels() []*CreateTrainingJobRequestOutputChannels {
	return s.OutputChannels
}

func (s *CreateTrainingJobRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateTrainingJobRequest) GetPythonRequirements() []*string {
	return s.PythonRequirements
}

func (s *CreateTrainingJobRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CreateTrainingJobRequest) GetScheduler() *CreateTrainingJobRequestScheduler {
	return s.Scheduler
}

func (s *CreateTrainingJobRequest) GetSettings() *JobSettings {
	return s.Settings
}

func (s *CreateTrainingJobRequest) GetTrainingJobDescription() *string {
	return s.TrainingJobDescription
}

func (s *CreateTrainingJobRequest) GetTrainingJobName() *string {
	return s.TrainingJobName
}

func (s *CreateTrainingJobRequest) GetUserVpc() *CreateTrainingJobRequestUserVpc {
	return s.UserVpc
}

func (s *CreateTrainingJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateTrainingJobRequest) SetAlgorithmName(v string) *CreateTrainingJobRequest {
	s.AlgorithmName = &v
	return s
}

func (s *CreateTrainingJobRequest) SetAlgorithmProvider(v string) *CreateTrainingJobRequest {
	s.AlgorithmProvider = &v
	return s
}

func (s *CreateTrainingJobRequest) SetAlgorithmSpec(v *AlgorithmSpec) *CreateTrainingJobRequest {
	s.AlgorithmSpec = v
	return s
}

func (s *CreateTrainingJobRequest) SetAlgorithmVersion(v string) *CreateTrainingJobRequest {
	s.AlgorithmVersion = &v
	return s
}

func (s *CreateTrainingJobRequest) SetAssignNodeSpec(v *AssignNodeSpec) *CreateTrainingJobRequest {
	s.AssignNodeSpec = v
	return s
}

func (s *CreateTrainingJobRequest) SetCodeDir(v *Location) *CreateTrainingJobRequest {
	s.CodeDir = v
	return s
}

func (s *CreateTrainingJobRequest) SetComputeResource(v *CreateTrainingJobRequestComputeResource) *CreateTrainingJobRequest {
	s.ComputeResource = v
	return s
}

func (s *CreateTrainingJobRequest) SetCredentialConfig(v *CredentialConfig) *CreateTrainingJobRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateTrainingJobRequest) SetEnvironments(v map[string]*string) *CreateTrainingJobRequest {
	s.Environments = v
	return s
}

func (s *CreateTrainingJobRequest) SetExperimentConfig(v *CreateTrainingJobRequestExperimentConfig) *CreateTrainingJobRequest {
	s.ExperimentConfig = v
	return s
}

func (s *CreateTrainingJobRequest) SetHyperParameters(v []*CreateTrainingJobRequestHyperParameters) *CreateTrainingJobRequest {
	s.HyperParameters = v
	return s
}

func (s *CreateTrainingJobRequest) SetInputChannels(v []*CreateTrainingJobRequestInputChannels) *CreateTrainingJobRequest {
	s.InputChannels = v
	return s
}

func (s *CreateTrainingJobRequest) SetLabels(v []*CreateTrainingJobRequestLabels) *CreateTrainingJobRequest {
	s.Labels = v
	return s
}

func (s *CreateTrainingJobRequest) SetOutputChannels(v []*CreateTrainingJobRequestOutputChannels) *CreateTrainingJobRequest {
	s.OutputChannels = v
	return s
}

func (s *CreateTrainingJobRequest) SetPriority(v int32) *CreateTrainingJobRequest {
	s.Priority = &v
	return s
}

func (s *CreateTrainingJobRequest) SetPythonRequirements(v []*string) *CreateTrainingJobRequest {
	s.PythonRequirements = v
	return s
}

func (s *CreateTrainingJobRequest) SetRoleArn(v string) *CreateTrainingJobRequest {
	s.RoleArn = &v
	return s
}

func (s *CreateTrainingJobRequest) SetScheduler(v *CreateTrainingJobRequestScheduler) *CreateTrainingJobRequest {
	s.Scheduler = v
	return s
}

func (s *CreateTrainingJobRequest) SetSettings(v *JobSettings) *CreateTrainingJobRequest {
	s.Settings = v
	return s
}

func (s *CreateTrainingJobRequest) SetTrainingJobDescription(v string) *CreateTrainingJobRequest {
	s.TrainingJobDescription = &v
	return s
}

func (s *CreateTrainingJobRequest) SetTrainingJobName(v string) *CreateTrainingJobRequest {
	s.TrainingJobName = &v
	return s
}

func (s *CreateTrainingJobRequest) SetUserVpc(v *CreateTrainingJobRequestUserVpc) *CreateTrainingJobRequest {
	s.UserVpc = v
	return s
}

func (s *CreateTrainingJobRequest) SetWorkspaceId(v string) *CreateTrainingJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateTrainingJobRequest) Validate() error {
	if s.AlgorithmSpec != nil {
		if err := s.AlgorithmSpec.Validate(); err != nil {
			return err
		}
	}
	if s.AssignNodeSpec != nil {
		if err := s.AssignNodeSpec.Validate(); err != nil {
			return err
		}
	}
	if s.CodeDir != nil {
		if err := s.CodeDir.Validate(); err != nil {
			return err
		}
	}
	if s.ComputeResource != nil {
		if err := s.ComputeResource.Validate(); err != nil {
			return err
		}
	}
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ExperimentConfig != nil {
		if err := s.ExperimentConfig.Validate(); err != nil {
			return err
		}
	}
	if s.HyperParameters != nil {
		for _, item := range s.HyperParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InputChannels != nil {
		for _, item := range s.InputChannels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OutputChannels != nil {
		for _, item := range s.OutputChannels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Scheduler != nil {
		if err := s.Scheduler.Validate(); err != nil {
			return err
		}
	}
	if s.Settings != nil {
		if err := s.Settings.Validate(); err != nil {
			return err
		}
	}
	if s.UserVpc != nil {
		if err := s.UserVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTrainingJobRequestComputeResource struct {
	// The number of ECS instances.
	//
	// example:
	//
	// 1
	EcsCount *int64 `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	// The ECS instance type.
	//
	// example:
	//
	// ecs.gn5-c8g1.2xlarge
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// The number of instances used from the resource quota.
	//
	// example:
	//
	// 1
	InstanceCount *int64 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The instance specification for the resource quota.
	InstanceSpec *CreateTrainingJobRequestComputeResourceInstanceSpec `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" type:"Struct"`
	// The resource quota ID.
	//
	// example:
	//
	// quotam670lixikcs
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The spot instance configuration.
	SpotSpec *CreateTrainingJobRequestComputeResourceSpotSpec `json:"SpotSpec,omitempty" xml:"SpotSpec,omitempty" type:"Struct"`
	// Specifies whether to use spot instances.
	//
	// example:
	//
	// true
	UseSpotInstance *bool `json:"UseSpotInstance,omitempty" xml:"UseSpotInstance,omitempty"`
}

func (s CreateTrainingJobRequestComputeResource) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainingJobRequestComputeResource) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestComputeResource) GetEcsCount() *int64 {
	return s.EcsCount
}

func (s *CreateTrainingJobRequestComputeResource) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *CreateTrainingJobRequestComputeResource) GetInstanceCount() *int64 {
	return s.InstanceCount
}

func (s *CreateTrainingJobRequestComputeResource) GetInstanceSpec() *CreateTrainingJobRequestComputeResourceInstanceSpec {
	return s.InstanceSpec
}

func (s *CreateTrainingJobRequestComputeResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateTrainingJobRequestComputeResource) GetSpotSpec() *CreateTrainingJobRequestComputeResourceSpotSpec {
	return s.SpotSpec
}

func (s *CreateTrainingJobRequestComputeResource) GetUseSpotInstance() *bool {
	return s.UseSpotInstance
}

func (s *CreateTrainingJobRequestComputeResource) SetEcsCount(v int64) *CreateTrainingJobRequestComputeResource {
	s.EcsCount = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResource) SetEcsSpec(v string) *CreateTrainingJobRequestComputeResource {
	s.EcsSpec = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResource) SetInstanceCount(v int64) *CreateTrainingJobRequestComputeResource {
	s.InstanceCount = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResource) SetInstanceSpec(v *CreateTrainingJobRequestComputeResourceInstanceSpec) *CreateTrainingJobRequestComputeResource {
	s.InstanceSpec = v
	return s
}

func (s *CreateTrainingJobRequestComputeResource) SetResourceId(v string) *CreateTrainingJobRequestComputeResource {
	s.ResourceId = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResource) SetSpotSpec(v *CreateTrainingJobRequestComputeResourceSpotSpec) *CreateTrainingJobRequestComputeResource {
	s.SpotSpec = v
	return s
}

func (s *CreateTrainingJobRequestComputeResource) SetUseSpotInstance(v bool) *CreateTrainingJobRequestComputeResource {
	s.UseSpotInstance = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResource) Validate() error {
	if s.InstanceSpec != nil {
		if err := s.InstanceSpec.Validate(); err != nil {
			return err
		}
	}
	if s.SpotSpec != nil {
		if err := s.SpotSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTrainingJobRequestComputeResourceInstanceSpec struct {
	// The number of CPU cores for the instance.
	//
	// example:
	//
	// 8
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// The number of GPUs for the instance.
	//
	// example:
	//
	// 1
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// The GPU type for the instance.
	//
	// example:
	//
	// V100
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// The memory size of the instance. Unit: GiB.
	//
	// example:
	//
	// 32
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The shared memory size of the instance. Unit: GB.
	//
	// example:
	//
	// 32
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s CreateTrainingJobRequestComputeResourceInstanceSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainingJobRequestComputeResourceInstanceSpec) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestComputeResourceInstanceSpec) GetCPU() *string {
	return s.CPU
}

func (s *CreateTrainingJobRequestComputeResourceInstanceSpec) GetGPU() *string {
	return s.GPU
}

func (s *CreateTrainingJobRequestComputeResourceInstanceSpec) GetGPUType() *string {
	return s.GPUType
}

func (s *CreateTrainingJobRequestComputeResourceInstanceSpec) GetMemory() *string {
	return s.Memory
}

func (s *CreateTrainingJobRequestComputeResourceInstanceSpec) GetSharedMemory() *string {
	return s.SharedMemory
}

func (s *CreateTrainingJobRequestComputeResourceInstanceSpec) SetCPU(v string) *CreateTrainingJobRequestComputeResourceInstanceSpec {
	s.CPU = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResourceInstanceSpec) SetGPU(v string) *CreateTrainingJobRequestComputeResourceInstanceSpec {
	s.GPU = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResourceInstanceSpec) SetGPUType(v string) *CreateTrainingJobRequestComputeResourceInstanceSpec {
	s.GPUType = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResourceInstanceSpec) SetMemory(v string) *CreateTrainingJobRequestComputeResourceInstanceSpec {
	s.Memory = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResourceInstanceSpec) SetSharedMemory(v string) *CreateTrainingJobRequestComputeResourceInstanceSpec {
	s.SharedMemory = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResourceInstanceSpec) Validate() error {
	return dara.Validate(s)
}

type CreateTrainingJobRequestComputeResourceSpotSpec struct {
	// The maximum hourly price discount for the instance. This parameter takes effect only when SpotStrategy is set to SpotWithPriceLimit.
	//
	// example:
	//
	// 9
	SpotDiscountLimit *float32 `json:"SpotDiscountLimit,omitempty" xml:"SpotDiscountLimit,omitempty"`
	// The bidding strategy for the spot instance. Valid values:
	//
	// - SpotWithPriceLimit: a spot instance with a maximum price limit.
	//
	// - SpotAsPriceGo: the system automatically bids at the current market price.
	//
	// example:
	//
	// SpotWithPriceLimit
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s CreateTrainingJobRequestComputeResourceSpotSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainingJobRequestComputeResourceSpotSpec) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestComputeResourceSpotSpec) GetSpotDiscountLimit() *float32 {
	return s.SpotDiscountLimit
}

func (s *CreateTrainingJobRequestComputeResourceSpotSpec) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateTrainingJobRequestComputeResourceSpotSpec) SetSpotDiscountLimit(v float32) *CreateTrainingJobRequestComputeResourceSpotSpec {
	s.SpotDiscountLimit = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResourceSpotSpec) SetSpotStrategy(v string) *CreateTrainingJobRequestComputeResourceSpotSpec {
	s.SpotStrategy = &v
	return s
}

func (s *CreateTrainingJobRequestComputeResourceSpotSpec) Validate() error {
	return dara.Validate(s)
}

type CreateTrainingJobRequestExperimentConfig struct {
	// The experiment ID associated with the training job.
	//
	// example:
	//
	// exp-ds9aefia90v
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
}

func (s CreateTrainingJobRequestExperimentConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainingJobRequestExperimentConfig) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestExperimentConfig) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *CreateTrainingJobRequestExperimentConfig) SetExperimentId(v string) *CreateTrainingJobRequestExperimentConfig {
	s.ExperimentId = &v
	return s
}

func (s *CreateTrainingJobRequestExperimentConfig) Validate() error {
	return dara.Validate(s)
}

type CreateTrainingJobRequestHyperParameters struct {
	// The parameter name.
	//
	// example:
	//
	// learning_rate
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// 0.0001
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTrainingJobRequestHyperParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainingJobRequestHyperParameters) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestHyperParameters) GetName() *string {
	return s.Name
}

func (s *CreateTrainingJobRequestHyperParameters) GetValue() *string {
	return s.Value
}

func (s *CreateTrainingJobRequestHyperParameters) SetName(v string) *CreateTrainingJobRequestHyperParameters {
	s.Name = &v
	return s
}

func (s *CreateTrainingJobRequestHyperParameters) SetValue(v string) *CreateTrainingJobRequestHyperParameters {
	s.Value = &v
	return s
}

func (s *CreateTrainingJobRequestHyperParameters) Validate() error {
	return dara.Validate(s)
}

type CreateTrainingJobRequestInputChannels struct {
	// The dataset ID.
	//
	// example:
	//
	// d-475megosidivjfgfq6
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The input data URI.
	//
	// example:
	//
	// oss://pai-quickstart-cn-hangzhou.oss-cn-hangzhou-internal.aliyuncs.com/modelscope/models/qwen2-0.5b/main/
	InputUri *string `json:"InputUri,omitempty" xml:"InputUri,omitempty"`
	// The input data name.
	//
	// example:
	//
	// model
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The input data parameter settings.
	//
	// example:
	//
	// {"appendable": true}
	Options     *string `json:"Options,omitempty" xml:"Options,omitempty"`
	RoleArn     *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateTrainingJobRequestInputChannels) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainingJobRequestInputChannels) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestInputChannels) GetDatasetId() *string {
	return s.DatasetId
}

func (s *CreateTrainingJobRequestInputChannels) GetInputUri() *string {
	return s.InputUri
}

func (s *CreateTrainingJobRequestInputChannels) GetName() *string {
	return s.Name
}

func (s *CreateTrainingJobRequestInputChannels) GetOptions() *string {
	return s.Options
}

func (s *CreateTrainingJobRequestInputChannels) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CreateTrainingJobRequestInputChannels) GetVersionName() *string {
	return s.VersionName
}

func (s *CreateTrainingJobRequestInputChannels) SetDatasetId(v string) *CreateTrainingJobRequestInputChannels {
	s.DatasetId = &v
	return s
}

func (s *CreateTrainingJobRequestInputChannels) SetInputUri(v string) *CreateTrainingJobRequestInputChannels {
	s.InputUri = &v
	return s
}

func (s *CreateTrainingJobRequestInputChannels) SetName(v string) *CreateTrainingJobRequestInputChannels {
	s.Name = &v
	return s
}

func (s *CreateTrainingJobRequestInputChannels) SetOptions(v string) *CreateTrainingJobRequestInputChannels {
	s.Options = &v
	return s
}

func (s *CreateTrainingJobRequestInputChannels) SetRoleArn(v string) *CreateTrainingJobRequestInputChannels {
	s.RoleArn = &v
	return s
}

func (s *CreateTrainingJobRequestInputChannels) SetVersionName(v string) *CreateTrainingJobRequestInputChannels {
	s.VersionName = &v
	return s
}

func (s *CreateTrainingJobRequestInputChannels) Validate() error {
	return dara.Validate(s)
}

type CreateTrainingJobRequestLabels struct {
	// The key of the label.
	//
	// example:
	//
	// CreatedBy
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the label.
	//
	// example:
	//
	// QuickStart
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTrainingJobRequestLabels) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainingJobRequestLabels) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestLabels) GetKey() *string {
	return s.Key
}

func (s *CreateTrainingJobRequestLabels) GetValue() *string {
	return s.Value
}

func (s *CreateTrainingJobRequestLabels) SetKey(v string) *CreateTrainingJobRequestLabels {
	s.Key = &v
	return s
}

func (s *CreateTrainingJobRequestLabels) SetValue(v string) *CreateTrainingJobRequestLabels {
	s.Value = &v
	return s
}

func (s *CreateTrainingJobRequestLabels) Validate() error {
	return dara.Validate(s)
}

type CreateTrainingJobRequestOutputChannels struct {
	// The dataset ID.
	//
	// example:
	//
	// d-475megosidivjfgfq6
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The output data name.
	//
	// example:
	//
	// model
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output data URI.
	//
	// example:
	//
	// oss://pai-quickstart-cn-hangzhou.oss-cn-hangzhou-internal.aliyuncs.com/modelscope/models/qwen2-0.5b/main/
	OutputUri   *string `json:"OutputUri,omitempty" xml:"OutputUri,omitempty"`
	RoleArn     *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateTrainingJobRequestOutputChannels) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainingJobRequestOutputChannels) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestOutputChannels) GetDatasetId() *string {
	return s.DatasetId
}

func (s *CreateTrainingJobRequestOutputChannels) GetName() *string {
	return s.Name
}

func (s *CreateTrainingJobRequestOutputChannels) GetOutputUri() *string {
	return s.OutputUri
}

func (s *CreateTrainingJobRequestOutputChannels) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CreateTrainingJobRequestOutputChannels) GetVersionName() *string {
	return s.VersionName
}

func (s *CreateTrainingJobRequestOutputChannels) SetDatasetId(v string) *CreateTrainingJobRequestOutputChannels {
	s.DatasetId = &v
	return s
}

func (s *CreateTrainingJobRequestOutputChannels) SetName(v string) *CreateTrainingJobRequestOutputChannels {
	s.Name = &v
	return s
}

func (s *CreateTrainingJobRequestOutputChannels) SetOutputUri(v string) *CreateTrainingJobRequestOutputChannels {
	s.OutputUri = &v
	return s
}

func (s *CreateTrainingJobRequestOutputChannels) SetRoleArn(v string) *CreateTrainingJobRequestOutputChannels {
	s.RoleArn = &v
	return s
}

func (s *CreateTrainingJobRequestOutputChannels) SetVersionName(v string) *CreateTrainingJobRequestOutputChannels {
	s.VersionName = &v
	return s
}

func (s *CreateTrainingJobRequestOutputChannels) Validate() error {
	return dara.Validate(s)
}

type CreateTrainingJobRequestScheduler struct {
	// The maximum training runtime in minutes. A value of 0 indicates no limit on the maximum runtime.
	//
	// example:
	//
	// 0
	MaxRunningTimeInMinutes *int64 `json:"MaxRunningTimeInMinutes,omitempty" xml:"MaxRunningTimeInMinutes,omitempty"`
	// The maximum training runtime in seconds. A value of 0 indicates no limit on the maximum runtime.
	//
	// example:
	//
	// 0
	MaxRunningTimeInSeconds *int64 `json:"MaxRunningTimeInSeconds,omitempty" xml:"MaxRunningTimeInSeconds,omitempty"`
}

func (s CreateTrainingJobRequestScheduler) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainingJobRequestScheduler) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestScheduler) GetMaxRunningTimeInMinutes() *int64 {
	return s.MaxRunningTimeInMinutes
}

func (s *CreateTrainingJobRequestScheduler) GetMaxRunningTimeInSeconds() *int64 {
	return s.MaxRunningTimeInSeconds
}

func (s *CreateTrainingJobRequestScheduler) SetMaxRunningTimeInMinutes(v int64) *CreateTrainingJobRequestScheduler {
	s.MaxRunningTimeInMinutes = &v
	return s
}

func (s *CreateTrainingJobRequestScheduler) SetMaxRunningTimeInSeconds(v int64) *CreateTrainingJobRequestScheduler {
	s.MaxRunningTimeInSeconds = &v
	return s
}

func (s *CreateTrainingJobRequestScheduler) Validate() error {
	return dara.Validate(s)
}

type CreateTrainingJobRequestUserVpc struct {
	// The default route interface. eth0 indicates that the default route uses the PAI VPC. eth1 indicates that the default route uses the user VPC. Default value: eth0.
	//
	// example:
	//
	// eth0
	DefaultRoute *string `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	// The extended CIDR block configuration.
	ExtendedCIDRs []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	// The security group ID.
	//
	// example:
	//
	// sg-qdfasd13sdasf
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vs-icrc813vdsfol
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-dxiflssjx978sl
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateTrainingJobRequestUserVpc) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainingJobRequestUserVpc) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobRequestUserVpc) GetDefaultRoute() *string {
	return s.DefaultRoute
}

func (s *CreateTrainingJobRequestUserVpc) GetExtendedCIDRs() []*string {
	return s.ExtendedCIDRs
}

func (s *CreateTrainingJobRequestUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateTrainingJobRequestUserVpc) GetSwitchId() *string {
	return s.SwitchId
}

func (s *CreateTrainingJobRequestUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateTrainingJobRequestUserVpc) SetDefaultRoute(v string) *CreateTrainingJobRequestUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *CreateTrainingJobRequestUserVpc) SetExtendedCIDRs(v []*string) *CreateTrainingJobRequestUserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *CreateTrainingJobRequestUserVpc) SetSecurityGroupId(v string) *CreateTrainingJobRequestUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateTrainingJobRequestUserVpc) SetSwitchId(v string) *CreateTrainingJobRequestUserVpc {
	s.SwitchId = &v
	return s
}

func (s *CreateTrainingJobRequestUserVpc) SetVpcId(v string) *CreateTrainingJobRequestUserVpc {
	s.VpcId = &v
	return s
}

func (s *CreateTrainingJobRequestUserVpc) Validate() error {
	return dara.Validate(s)
}
