// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrainingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmId(v string) *GetTrainingJobResponseBody
	GetAlgorithmId() *string
	SetAlgorithmName(v string) *GetTrainingJobResponseBody
	GetAlgorithmName() *string
	SetAlgorithmProvider(v string) *GetTrainingJobResponseBody
	GetAlgorithmProvider() *string
	SetAlgorithmSpec(v *AlgorithmSpec) *GetTrainingJobResponseBody
	GetAlgorithmSpec() *AlgorithmSpec
	SetAlgorithmVersion(v string) *GetTrainingJobResponseBody
	GetAlgorithmVersion() *string
	SetAssignNodeSpec(v *AssignNodeSpec) *GetTrainingJobResponseBody
	GetAssignNodeSpec() *AssignNodeSpec
	SetComputeResource(v *GetTrainingJobResponseBodyComputeResource) *GetTrainingJobResponseBody
	GetComputeResource() *GetTrainingJobResponseBodyComputeResource
	SetCredentialConfig(v *CredentialConfig) *GetTrainingJobResponseBody
	GetCredentialConfig() *CredentialConfig
	SetDuration(v int64) *GetTrainingJobResponseBody
	GetDuration() *int64
	SetEnvironments(v map[string]*string) *GetTrainingJobResponseBody
	GetEnvironments() map[string]*string
	SetExperimentConfig(v *GetTrainingJobResponseBodyExperimentConfig) *GetTrainingJobResponseBody
	GetExperimentConfig() *GetTrainingJobResponseBodyExperimentConfig
	SetGmtCreateTime(v string) *GetTrainingJobResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetTrainingJobResponseBody
	GetGmtModifiedTime() *string
	SetHyperParameters(v []*GetTrainingJobResponseBodyHyperParameters) *GetTrainingJobResponseBody
	GetHyperParameters() []*GetTrainingJobResponseBodyHyperParameters
	SetInputChannels(v []*GetTrainingJobResponseBodyInputChannels) *GetTrainingJobResponseBody
	GetInputChannels() []*GetTrainingJobResponseBodyInputChannels
	SetInstances(v []*GetTrainingJobResponseBodyInstances) *GetTrainingJobResponseBody
	GetInstances() []*GetTrainingJobResponseBodyInstances
	SetIsTempAlgo(v bool) *GetTrainingJobResponseBody
	GetIsTempAlgo() *bool
	SetLabels(v []*GetTrainingJobResponseBodyLabels) *GetTrainingJobResponseBody
	GetLabels() []*GetTrainingJobResponseBodyLabels
	SetLatestMetrics(v []*GetTrainingJobResponseBodyLatestMetrics) *GetTrainingJobResponseBody
	GetLatestMetrics() []*GetTrainingJobResponseBodyLatestMetrics
	SetLatestProgress(v *GetTrainingJobResponseBodyLatestProgress) *GetTrainingJobResponseBody
	GetLatestProgress() *GetTrainingJobResponseBodyLatestProgress
	SetOutputChannels(v []*GetTrainingJobResponseBodyOutputChannels) *GetTrainingJobResponseBody
	GetOutputChannels() []*GetTrainingJobResponseBodyOutputChannels
	SetOutputModel(v *GetTrainingJobResponseBodyOutputModel) *GetTrainingJobResponseBody
	GetOutputModel() *GetTrainingJobResponseBodyOutputModel
	SetPriority(v int32) *GetTrainingJobResponseBody
	GetPriority() *int32
	SetPythonRequirements(v []*string) *GetTrainingJobResponseBody
	GetPythonRequirements() []*string
	SetReasonCode(v string) *GetTrainingJobResponseBody
	GetReasonCode() *string
	SetReasonMessage(v string) *GetTrainingJobResponseBody
	GetReasonMessage() *string
	SetRequestId(v string) *GetTrainingJobResponseBody
	GetRequestId() *string
	SetRoleArn(v string) *GetTrainingJobResponseBody
	GetRoleArn() *string
	SetScheduler(v *GetTrainingJobResponseBodyScheduler) *GetTrainingJobResponseBody
	GetScheduler() *GetTrainingJobResponseBodyScheduler
	SetSettings(v *JobSettings) *GetTrainingJobResponseBody
	GetSettings() *JobSettings
	SetStatus(v string) *GetTrainingJobResponseBody
	GetStatus() *string
	SetStatusTransitions(v []*GetTrainingJobResponseBodyStatusTransitions) *GetTrainingJobResponseBody
	GetStatusTransitions() []*GetTrainingJobResponseBodyStatusTransitions
	SetTrainingJobDescription(v string) *GetTrainingJobResponseBody
	GetTrainingJobDescription() *string
	SetTrainingJobId(v string) *GetTrainingJobResponseBody
	GetTrainingJobId() *string
	SetTrainingJobName(v string) *GetTrainingJobResponseBody
	GetTrainingJobName() *string
	SetTrainingJobUrl(v string) *GetTrainingJobResponseBody
	GetTrainingJobUrl() *string
	SetUserId(v string) *GetTrainingJobResponseBody
	GetUserId() *string
	SetUserVpc(v *GetTrainingJobResponseBodyUserVpc) *GetTrainingJobResponseBody
	GetUserVpc() *GetTrainingJobResponseBodyUserVpc
	SetWorkspaceId(v string) *GetTrainingJobResponseBody
	GetWorkspaceId() *string
}

type GetTrainingJobResponseBody struct {
	// The training algorithm ID.
	//
	// example:
	//
	// algo-xsldfvu1334
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// The algorithm name.
	//
	// example:
	//
	// llm_training
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	// The algorithm provider.
	//
	// example:
	//
	// pai
	AlgorithmProvider *string `json:"AlgorithmProvider,omitempty" xml:"AlgorithmProvider,omitempty"`
	// The temporary algorithm definition.
	AlgorithmSpec *AlgorithmSpec `json:"AlgorithmSpec,omitempty" xml:"AlgorithmSpec,omitempty"`
	// The algorithm version.
	//
	// example:
	//
	// v0.0.1
	AlgorithmVersion *string         `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
	AssignNodeSpec   *AssignNodeSpec `json:"AssignNodeSpec,omitempty" xml:"AssignNodeSpec,omitempty"`
	// The compute resource configuration.
	ComputeResource  *GetTrainingJobResponseBodyComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	CredentialConfig *CredentialConfig                          `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The running duration of the training job. Unit: seconds.
	//
	// example:
	//
	// 7200
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The environment variables of the training job.
	Environments map[string]*string `json:"Environments,omitempty" xml:"Environments,omitempty"`
	// The experiment configuration associated with the training job.
	ExperimentConfig *GetTrainingJobResponseBodyExperimentConfig `json:"ExperimentConfig,omitempty" xml:"ExperimentConfig,omitempty" type:"Struct"`
	// The time when the training job was created.
	//
	// example:
	//
	// 2024-07-10T11:49:47Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the training job status was last updated.
	//
	// example:
	//
	// 2024-07-10T11:49:47Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The training hyperparameter settings.
	HyperParameters []*GetTrainingJobResponseBodyHyperParameters `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Repeated"`
	// The training input data configurations.
	InputChannels []*GetTrainingJobResponseBodyInputChannels `json:"InputChannels,omitempty" xml:"InputChannels,omitempty" type:"Repeated"`
	// The list of training job instances.
	Instances []*GetTrainingJobResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// Indicates whether a temporary algorithm is used.
	//
	// example:
	//
	// true
	IsTempAlgo *bool `json:"IsTempAlgo,omitempty" xml:"IsTempAlgo,omitempty"`
	// The list of training job labels.
	Labels []*GetTrainingJobResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The list of training job metrics.
	LatestMetrics []*GetTrainingJobResponseBodyLatestMetrics `json:"LatestMetrics,omitempty" xml:"LatestMetrics,omitempty" type:"Repeated"`
	// The latest progress of the training job.
	LatestProgress *GetTrainingJobResponseBodyLatestProgress `json:"LatestProgress,omitempty" xml:"LatestProgress,omitempty" type:"Struct"`
	// The training output data configurations.
	OutputChannels []*GetTrainingJobResponseBodyOutputChannels `json:"OutputChannels,omitempty" xml:"OutputChannels,omitempty" type:"Repeated"`
	// The model produced by the training job.
	OutputModel *GetTrainingJobResponseBodyOutputModel `json:"OutputModel,omitempty" xml:"OutputModel,omitempty" type:"Struct"`
	// The job priority.
	//
	// example:
	//
	// 0
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The Python package configuration for the training job.
	PythonRequirements []*string `json:"PythonRequirements,omitempty" xml:"PythonRequirements,omitempty" type:"Repeated"`
	// The status code of the training job.
	//
	// example:
	//
	// TrainingJobSucceed
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The error message of the training job.
	//
	// example:
	//
	// None
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ARN of the RAM role used for proxy authorization.
	//
	// example:
	//
	// acs:ram::{accountID}:role/{roleName}
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The scheduling configuration of the training job.
	Scheduler *GetTrainingJobResponseBodyScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
	// The additional parameter settings for the training node.
	Settings *JobSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// The task status.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of training job status transitions.
	StatusTransitions []*GetTrainingJobResponseBodyStatusTransitions `json:"StatusTransitions,omitempty" xml:"StatusTransitions,omitempty" type:"Repeated"`
	// The description of the training job.
	//
	// example:
	//
	// Qwen2 large language model training.
	TrainingJobDescription *string `json:"TrainingJobDescription,omitempty" xml:"TrainingJobDescription,omitempty"`
	// The training job ID.
	//
	// example:
	//
	// traini6hhxiq69eo
	TrainingJobId *string `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
	// The name of the training job.
	//
	// example:
	//
	// qwen_llm
	TrainingJobName *string `json:"TrainingJobName,omitempty" xml:"TrainingJobName,omitempty"`
	// The URL of the training job details page.
	//
	// example:
	//
	// https://pai.console.aliyun.com/?regionId=cn-hangzhou&workspaceId=1234#/training/jobs/train1ouyadsl8n4
	TrainingJobUrl *string `json:"TrainingJobUrl,omitempty" xml:"TrainingJobUrl,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The user VPC configuration.
	UserVpc *GetTrainingJobResponseBodyUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// The workspace ID.
	//
	// example:
	//
	// 86995
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetTrainingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBody) GetAlgorithmId() *string {
	return s.AlgorithmId
}

func (s *GetTrainingJobResponseBody) GetAlgorithmName() *string {
	return s.AlgorithmName
}

func (s *GetTrainingJobResponseBody) GetAlgorithmProvider() *string {
	return s.AlgorithmProvider
}

func (s *GetTrainingJobResponseBody) GetAlgorithmSpec() *AlgorithmSpec {
	return s.AlgorithmSpec
}

func (s *GetTrainingJobResponseBody) GetAlgorithmVersion() *string {
	return s.AlgorithmVersion
}

func (s *GetTrainingJobResponseBody) GetAssignNodeSpec() *AssignNodeSpec {
	return s.AssignNodeSpec
}

func (s *GetTrainingJobResponseBody) GetComputeResource() *GetTrainingJobResponseBodyComputeResource {
	return s.ComputeResource
}

func (s *GetTrainingJobResponseBody) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *GetTrainingJobResponseBody) GetDuration() *int64 {
	return s.Duration
}

func (s *GetTrainingJobResponseBody) GetEnvironments() map[string]*string {
	return s.Environments
}

func (s *GetTrainingJobResponseBody) GetExperimentConfig() *GetTrainingJobResponseBodyExperimentConfig {
	return s.ExperimentConfig
}

func (s *GetTrainingJobResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetTrainingJobResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetTrainingJobResponseBody) GetHyperParameters() []*GetTrainingJobResponseBodyHyperParameters {
	return s.HyperParameters
}

func (s *GetTrainingJobResponseBody) GetInputChannels() []*GetTrainingJobResponseBodyInputChannels {
	return s.InputChannels
}

func (s *GetTrainingJobResponseBody) GetInstances() []*GetTrainingJobResponseBodyInstances {
	return s.Instances
}

func (s *GetTrainingJobResponseBody) GetIsTempAlgo() *bool {
	return s.IsTempAlgo
}

func (s *GetTrainingJobResponseBody) GetLabels() []*GetTrainingJobResponseBodyLabels {
	return s.Labels
}

func (s *GetTrainingJobResponseBody) GetLatestMetrics() []*GetTrainingJobResponseBodyLatestMetrics {
	return s.LatestMetrics
}

func (s *GetTrainingJobResponseBody) GetLatestProgress() *GetTrainingJobResponseBodyLatestProgress {
	return s.LatestProgress
}

func (s *GetTrainingJobResponseBody) GetOutputChannels() []*GetTrainingJobResponseBodyOutputChannels {
	return s.OutputChannels
}

func (s *GetTrainingJobResponseBody) GetOutputModel() *GetTrainingJobResponseBodyOutputModel {
	return s.OutputModel
}

func (s *GetTrainingJobResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *GetTrainingJobResponseBody) GetPythonRequirements() []*string {
	return s.PythonRequirements
}

func (s *GetTrainingJobResponseBody) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *GetTrainingJobResponseBody) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *GetTrainingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTrainingJobResponseBody) GetRoleArn() *string {
	return s.RoleArn
}

func (s *GetTrainingJobResponseBody) GetScheduler() *GetTrainingJobResponseBodyScheduler {
	return s.Scheduler
}

func (s *GetTrainingJobResponseBody) GetSettings() *JobSettings {
	return s.Settings
}

func (s *GetTrainingJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetTrainingJobResponseBody) GetStatusTransitions() []*GetTrainingJobResponseBodyStatusTransitions {
	return s.StatusTransitions
}

func (s *GetTrainingJobResponseBody) GetTrainingJobDescription() *string {
	return s.TrainingJobDescription
}

func (s *GetTrainingJobResponseBody) GetTrainingJobId() *string {
	return s.TrainingJobId
}

func (s *GetTrainingJobResponseBody) GetTrainingJobName() *string {
	return s.TrainingJobName
}

func (s *GetTrainingJobResponseBody) GetTrainingJobUrl() *string {
	return s.TrainingJobUrl
}

func (s *GetTrainingJobResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetTrainingJobResponseBody) GetUserVpc() *GetTrainingJobResponseBodyUserVpc {
	return s.UserVpc
}

func (s *GetTrainingJobResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetTrainingJobResponseBody) SetAlgorithmId(v string) *GetTrainingJobResponseBody {
	s.AlgorithmId = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetAlgorithmName(v string) *GetTrainingJobResponseBody {
	s.AlgorithmName = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetAlgorithmProvider(v string) *GetTrainingJobResponseBody {
	s.AlgorithmProvider = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetAlgorithmSpec(v *AlgorithmSpec) *GetTrainingJobResponseBody {
	s.AlgorithmSpec = v
	return s
}

func (s *GetTrainingJobResponseBody) SetAlgorithmVersion(v string) *GetTrainingJobResponseBody {
	s.AlgorithmVersion = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetAssignNodeSpec(v *AssignNodeSpec) *GetTrainingJobResponseBody {
	s.AssignNodeSpec = v
	return s
}

func (s *GetTrainingJobResponseBody) SetComputeResource(v *GetTrainingJobResponseBodyComputeResource) *GetTrainingJobResponseBody {
	s.ComputeResource = v
	return s
}

func (s *GetTrainingJobResponseBody) SetCredentialConfig(v *CredentialConfig) *GetTrainingJobResponseBody {
	s.CredentialConfig = v
	return s
}

func (s *GetTrainingJobResponseBody) SetDuration(v int64) *GetTrainingJobResponseBody {
	s.Duration = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetEnvironments(v map[string]*string) *GetTrainingJobResponseBody {
	s.Environments = v
	return s
}

func (s *GetTrainingJobResponseBody) SetExperimentConfig(v *GetTrainingJobResponseBodyExperimentConfig) *GetTrainingJobResponseBody {
	s.ExperimentConfig = v
	return s
}

func (s *GetTrainingJobResponseBody) SetGmtCreateTime(v string) *GetTrainingJobResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetGmtModifiedTime(v string) *GetTrainingJobResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetHyperParameters(v []*GetTrainingJobResponseBodyHyperParameters) *GetTrainingJobResponseBody {
	s.HyperParameters = v
	return s
}

func (s *GetTrainingJobResponseBody) SetInputChannels(v []*GetTrainingJobResponseBodyInputChannels) *GetTrainingJobResponseBody {
	s.InputChannels = v
	return s
}

func (s *GetTrainingJobResponseBody) SetInstances(v []*GetTrainingJobResponseBodyInstances) *GetTrainingJobResponseBody {
	s.Instances = v
	return s
}

func (s *GetTrainingJobResponseBody) SetIsTempAlgo(v bool) *GetTrainingJobResponseBody {
	s.IsTempAlgo = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetLabels(v []*GetTrainingJobResponseBodyLabels) *GetTrainingJobResponseBody {
	s.Labels = v
	return s
}

func (s *GetTrainingJobResponseBody) SetLatestMetrics(v []*GetTrainingJobResponseBodyLatestMetrics) *GetTrainingJobResponseBody {
	s.LatestMetrics = v
	return s
}

func (s *GetTrainingJobResponseBody) SetLatestProgress(v *GetTrainingJobResponseBodyLatestProgress) *GetTrainingJobResponseBody {
	s.LatestProgress = v
	return s
}

func (s *GetTrainingJobResponseBody) SetOutputChannels(v []*GetTrainingJobResponseBodyOutputChannels) *GetTrainingJobResponseBody {
	s.OutputChannels = v
	return s
}

func (s *GetTrainingJobResponseBody) SetOutputModel(v *GetTrainingJobResponseBodyOutputModel) *GetTrainingJobResponseBody {
	s.OutputModel = v
	return s
}

func (s *GetTrainingJobResponseBody) SetPriority(v int32) *GetTrainingJobResponseBody {
	s.Priority = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetPythonRequirements(v []*string) *GetTrainingJobResponseBody {
	s.PythonRequirements = v
	return s
}

func (s *GetTrainingJobResponseBody) SetReasonCode(v string) *GetTrainingJobResponseBody {
	s.ReasonCode = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetReasonMessage(v string) *GetTrainingJobResponseBody {
	s.ReasonMessage = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetRequestId(v string) *GetTrainingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetRoleArn(v string) *GetTrainingJobResponseBody {
	s.RoleArn = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetScheduler(v *GetTrainingJobResponseBodyScheduler) *GetTrainingJobResponseBody {
	s.Scheduler = v
	return s
}

func (s *GetTrainingJobResponseBody) SetSettings(v *JobSettings) *GetTrainingJobResponseBody {
	s.Settings = v
	return s
}

func (s *GetTrainingJobResponseBody) SetStatus(v string) *GetTrainingJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetStatusTransitions(v []*GetTrainingJobResponseBodyStatusTransitions) *GetTrainingJobResponseBody {
	s.StatusTransitions = v
	return s
}

func (s *GetTrainingJobResponseBody) SetTrainingJobDescription(v string) *GetTrainingJobResponseBody {
	s.TrainingJobDescription = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetTrainingJobId(v string) *GetTrainingJobResponseBody {
	s.TrainingJobId = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetTrainingJobName(v string) *GetTrainingJobResponseBody {
	s.TrainingJobName = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetTrainingJobUrl(v string) *GetTrainingJobResponseBody {
	s.TrainingJobUrl = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetUserId(v string) *GetTrainingJobResponseBody {
	s.UserId = &v
	return s
}

func (s *GetTrainingJobResponseBody) SetUserVpc(v *GetTrainingJobResponseBodyUserVpc) *GetTrainingJobResponseBody {
	s.UserVpc = v
	return s
}

func (s *GetTrainingJobResponseBody) SetWorkspaceId(v string) *GetTrainingJobResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetTrainingJobResponseBody) Validate() error {
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
	if s.Instances != nil {
		for _, item := range s.Instances {
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
	if s.LatestMetrics != nil {
		for _, item := range s.LatestMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LatestProgress != nil {
		if err := s.LatestProgress.Validate(); err != nil {
			return err
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
	if s.OutputModel != nil {
		if err := s.OutputModel.Validate(); err != nil {
			return err
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
	if s.StatusTransitions != nil {
		for _, item := range s.StatusTransitions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserVpc != nil {
		if err := s.UserVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTrainingJobResponseBodyComputeResource struct {
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
	// The number of instances used by the resource quota.
	//
	// example:
	//
	// 1
	InstanceCount *int64 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The instance specification of the resource quota.
	InstanceSpec *GetTrainingJobResponseBodyComputeResourceInstanceSpec `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" type:"Struct"`
	// The resource quota ID.
	//
	// example:
	//
	// quotam670lixikcl
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource quota name.
	//
	// example:
	//
	// quota
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The spot instance configuration.
	SpotSpec *GetTrainingJobResponseBodyComputeResourceSpotSpec `json:"SpotSpec,omitempty" xml:"SpotSpec,omitempty" type:"Struct"`
	// Indicates whether spot instances are used.
	//
	// example:
	//
	// true
	UseSpotInstance *bool `json:"UseSpotInstance,omitempty" xml:"UseSpotInstance,omitempty"`
}

func (s GetTrainingJobResponseBodyComputeResource) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyComputeResource) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyComputeResource) GetEcsCount() *int64 {
	return s.EcsCount
}

func (s *GetTrainingJobResponseBodyComputeResource) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *GetTrainingJobResponseBodyComputeResource) GetInstanceCount() *int64 {
	return s.InstanceCount
}

func (s *GetTrainingJobResponseBodyComputeResource) GetInstanceSpec() *GetTrainingJobResponseBodyComputeResourceInstanceSpec {
	return s.InstanceSpec
}

func (s *GetTrainingJobResponseBodyComputeResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetTrainingJobResponseBodyComputeResource) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetTrainingJobResponseBodyComputeResource) GetSpotSpec() *GetTrainingJobResponseBodyComputeResourceSpotSpec {
	return s.SpotSpec
}

func (s *GetTrainingJobResponseBodyComputeResource) GetUseSpotInstance() *bool {
	return s.UseSpotInstance
}

func (s *GetTrainingJobResponseBodyComputeResource) SetEcsCount(v int64) *GetTrainingJobResponseBodyComputeResource {
	s.EcsCount = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResource) SetEcsSpec(v string) *GetTrainingJobResponseBodyComputeResource {
	s.EcsSpec = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResource) SetInstanceCount(v int64) *GetTrainingJobResponseBodyComputeResource {
	s.InstanceCount = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResource) SetInstanceSpec(v *GetTrainingJobResponseBodyComputeResourceInstanceSpec) *GetTrainingJobResponseBodyComputeResource {
	s.InstanceSpec = v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResource) SetResourceId(v string) *GetTrainingJobResponseBodyComputeResource {
	s.ResourceId = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResource) SetResourceName(v string) *GetTrainingJobResponseBodyComputeResource {
	s.ResourceName = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResource) SetSpotSpec(v *GetTrainingJobResponseBodyComputeResourceSpotSpec) *GetTrainingJobResponseBodyComputeResource {
	s.SpotSpec = v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResource) SetUseSpotInstance(v bool) *GetTrainingJobResponseBodyComputeResource {
	s.UseSpotInstance = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResource) Validate() error {
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

type GetTrainingJobResponseBodyComputeResourceInstanceSpec struct {
	// The number of CPU cores of the instance.
	//
	// example:
	//
	// 8
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// The number of GPUs of the instance.
	//
	// example:
	//
	// 1
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// The GPU type of the instance.
	//
	// example:
	//
	// V100
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// The memory size of the instance, in GiB.
	//
	// example:
	//
	// 32
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The shared memory size of the instance, in GiB.
	//
	// example:
	//
	// 32
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s GetTrainingJobResponseBodyComputeResourceInstanceSpec) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyComputeResourceInstanceSpec) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyComputeResourceInstanceSpec) GetCPU() *string {
	return s.CPU
}

func (s *GetTrainingJobResponseBodyComputeResourceInstanceSpec) GetGPU() *string {
	return s.GPU
}

func (s *GetTrainingJobResponseBodyComputeResourceInstanceSpec) GetGPUType() *string {
	return s.GPUType
}

func (s *GetTrainingJobResponseBodyComputeResourceInstanceSpec) GetMemory() *string {
	return s.Memory
}

func (s *GetTrainingJobResponseBodyComputeResourceInstanceSpec) GetSharedMemory() *string {
	return s.SharedMemory
}

func (s *GetTrainingJobResponseBodyComputeResourceInstanceSpec) SetCPU(v string) *GetTrainingJobResponseBodyComputeResourceInstanceSpec {
	s.CPU = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResourceInstanceSpec) SetGPU(v string) *GetTrainingJobResponseBodyComputeResourceInstanceSpec {
	s.GPU = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResourceInstanceSpec) SetGPUType(v string) *GetTrainingJobResponseBodyComputeResourceInstanceSpec {
	s.GPUType = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResourceInstanceSpec) SetMemory(v string) *GetTrainingJobResponseBodyComputeResourceInstanceSpec {
	s.Memory = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResourceInstanceSpec) SetSharedMemory(v string) *GetTrainingJobResponseBodyComputeResourceInstanceSpec {
	s.SharedMemory = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResourceInstanceSpec) Validate() error {
	return dara.Validate(s)
}

type GetTrainingJobResponseBodyComputeResourceSpotSpec struct {
	// The maximum hourly price discount for the instance. This parameter takes effect only when SpotStrategy is set to SpotWithPriceLimit.
	//
	// example:
	//
	// 0.9
	SpotDiscountLimit *float32 `json:"SpotDiscountLimit,omitempty" xml:"SpotDiscountLimit,omitempty"`
	// SpotStrategy: The bidding policy of the instance. Valid values:
	//
	// example:
	//
	// SpotWithPriceLimit
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s GetTrainingJobResponseBodyComputeResourceSpotSpec) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyComputeResourceSpotSpec) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyComputeResourceSpotSpec) GetSpotDiscountLimit() *float32 {
	return s.SpotDiscountLimit
}

func (s *GetTrainingJobResponseBodyComputeResourceSpotSpec) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *GetTrainingJobResponseBodyComputeResourceSpotSpec) SetSpotDiscountLimit(v float32) *GetTrainingJobResponseBodyComputeResourceSpotSpec {
	s.SpotDiscountLimit = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResourceSpotSpec) SetSpotStrategy(v string) *GetTrainingJobResponseBodyComputeResourceSpotSpec {
	s.SpotStrategy = &v
	return s
}

func (s *GetTrainingJobResponseBodyComputeResourceSpotSpec) Validate() error {
	return dara.Validate(s)
}

type GetTrainingJobResponseBodyExperimentConfig struct {
	// The experiment ID associated with the training job.
	//
	// example:
	//
	// exp-ds9aefia90v
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// The experiment name associated with the training job.
	//
	// example:
	//
	// large_language_model_train
	ExperimentName *string `json:"ExperimentName,omitempty" xml:"ExperimentName,omitempty"`
}

func (s GetTrainingJobResponseBodyExperimentConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyExperimentConfig) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyExperimentConfig) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *GetTrainingJobResponseBodyExperimentConfig) GetExperimentName() *string {
	return s.ExperimentName
}

func (s *GetTrainingJobResponseBodyExperimentConfig) SetExperimentId(v string) *GetTrainingJobResponseBodyExperimentConfig {
	s.ExperimentId = &v
	return s
}

func (s *GetTrainingJobResponseBodyExperimentConfig) SetExperimentName(v string) *GetTrainingJobResponseBodyExperimentConfig {
	s.ExperimentName = &v
	return s
}

func (s *GetTrainingJobResponseBodyExperimentConfig) Validate() error {
	return dara.Validate(s)
}

type GetTrainingJobResponseBodyHyperParameters struct {
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

func (s GetTrainingJobResponseBodyHyperParameters) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyHyperParameters) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyHyperParameters) GetName() *string {
	return s.Name
}

func (s *GetTrainingJobResponseBodyHyperParameters) GetValue() *string {
	return s.Value
}

func (s *GetTrainingJobResponseBodyHyperParameters) SetName(v string) *GetTrainingJobResponseBodyHyperParameters {
	s.Name = &v
	return s
}

func (s *GetTrainingJobResponseBodyHyperParameters) SetValue(v string) *GetTrainingJobResponseBodyHyperParameters {
	s.Value = &v
	return s
}

func (s *GetTrainingJobResponseBodyHyperParameters) Validate() error {
	return dara.Validate(s)
}

type GetTrainingJobResponseBodyInputChannels struct {
	// The dataset ID.
	//
	// example:
	//
	// d-475megosidivjfgfq6
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The URI of the input data.
	//
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou-internal.aliyuncs.com/path/to/input/model/
	InputUri *string `json:"InputUri,omitempty" xml:"InputUri,omitempty"`
	// The name of the input data.
	//
	// example:
	//
	// model
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The file system parameters of the input data.
	//
	// example:
	//
	// ossAppendable=true
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The dataset version.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetTrainingJobResponseBodyInputChannels) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyInputChannels) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyInputChannels) GetDatasetId() *string {
	return s.DatasetId
}

func (s *GetTrainingJobResponseBodyInputChannels) GetInputUri() *string {
	return s.InputUri
}

func (s *GetTrainingJobResponseBodyInputChannels) GetName() *string {
	return s.Name
}

func (s *GetTrainingJobResponseBodyInputChannels) GetOptions() *string {
	return s.Options
}

func (s *GetTrainingJobResponseBodyInputChannels) GetRoleArn() *string {
	return s.RoleArn
}

func (s *GetTrainingJobResponseBodyInputChannels) GetVersionName() *string {
	return s.VersionName
}

func (s *GetTrainingJobResponseBodyInputChannels) SetDatasetId(v string) *GetTrainingJobResponseBodyInputChannels {
	s.DatasetId = &v
	return s
}

func (s *GetTrainingJobResponseBodyInputChannels) SetInputUri(v string) *GetTrainingJobResponseBodyInputChannels {
	s.InputUri = &v
	return s
}

func (s *GetTrainingJobResponseBodyInputChannels) SetName(v string) *GetTrainingJobResponseBodyInputChannels {
	s.Name = &v
	return s
}

func (s *GetTrainingJobResponseBodyInputChannels) SetOptions(v string) *GetTrainingJobResponseBodyInputChannels {
	s.Options = &v
	return s
}

func (s *GetTrainingJobResponseBodyInputChannels) SetRoleArn(v string) *GetTrainingJobResponseBodyInputChannels {
	s.RoleArn = &v
	return s
}

func (s *GetTrainingJobResponseBodyInputChannels) SetVersionName(v string) *GetTrainingJobResponseBodyInputChannels {
	s.VersionName = &v
	return s
}

func (s *GetTrainingJobResponseBodyInputChannels) Validate() error {
	return dara.Validate(s)
}

type GetTrainingJobResponseBodyInstances struct {
	// The instance name.
	//
	// example:
	//
	// train1oug3yehan4-master-0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The instance role.
	//
	// example:
	//
	// master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The instance status.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTrainingJobResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyInstances) GetName() *string {
	return s.Name
}

func (s *GetTrainingJobResponseBodyInstances) GetRole() *string {
	return s.Role
}

func (s *GetTrainingJobResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *GetTrainingJobResponseBodyInstances) SetName(v string) *GetTrainingJobResponseBodyInstances {
	s.Name = &v
	return s
}

func (s *GetTrainingJobResponseBodyInstances) SetRole(v string) *GetTrainingJobResponseBodyInstances {
	s.Role = &v
	return s
}

func (s *GetTrainingJobResponseBodyInstances) SetStatus(v string) *GetTrainingJobResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *GetTrainingJobResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

type GetTrainingJobResponseBodyLabels struct {
	// The label name.
	//
	// example:
	//
	// CreatedBy
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The label value.
	//
	// example:
	//
	// QuickStart
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTrainingJobResponseBodyLabels) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyLabels) GetKey() *string {
	return s.Key
}

func (s *GetTrainingJobResponseBodyLabels) GetValue() *string {
	return s.Value
}

func (s *GetTrainingJobResponseBodyLabels) SetKey(v string) *GetTrainingJobResponseBodyLabels {
	s.Key = &v
	return s
}

func (s *GetTrainingJobResponseBodyLabels) SetValue(v string) *GetTrainingJobResponseBodyLabels {
	s.Value = &v
	return s
}

func (s *GetTrainingJobResponseBodyLabels) Validate() error {
	return dara.Validate(s)
}

type GetTrainingJobResponseBodyLatestMetrics struct {
	// The metric name.
	//
	// example:
	//
	// loss
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the metric was collected.
	//
	// example:
	//
	// 2024-07-10T11:49:47Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The metric value.
	//
	// example:
	//
	// 0.11
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTrainingJobResponseBodyLatestMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyLatestMetrics) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyLatestMetrics) GetName() *string {
	return s.Name
}

func (s *GetTrainingJobResponseBodyLatestMetrics) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetTrainingJobResponseBodyLatestMetrics) GetValue() *float64 {
	return s.Value
}

func (s *GetTrainingJobResponseBodyLatestMetrics) SetName(v string) *GetTrainingJobResponseBodyLatestMetrics {
	s.Name = &v
	return s
}

func (s *GetTrainingJobResponseBodyLatestMetrics) SetTimestamp(v string) *GetTrainingJobResponseBodyLatestMetrics {
	s.Timestamp = &v
	return s
}

func (s *GetTrainingJobResponseBodyLatestMetrics) SetValue(v float64) *GetTrainingJobResponseBodyLatestMetrics {
	s.Value = &v
	return s
}

func (s *GetTrainingJobResponseBodyLatestMetrics) Validate() error {
	return dara.Validate(s)
}

type GetTrainingJobResponseBodyLatestProgress struct {
	// The overall progress of the training job execution.
	OverallProgress *GetTrainingJobResponseBodyLatestProgressOverallProgress `json:"OverallProgress,omitempty" xml:"OverallProgress,omitempty" type:"Struct"`
	// The estimated remaining time for the training job execution, in seconds.
	RemainingTime *GetTrainingJobResponseBodyLatestProgressRemainingTime `json:"RemainingTime,omitempty" xml:"RemainingTime,omitempty" type:"Struct"`
}

func (s GetTrainingJobResponseBodyLatestProgress) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyLatestProgress) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyLatestProgress) GetOverallProgress() *GetTrainingJobResponseBodyLatestProgressOverallProgress {
	return s.OverallProgress
}

func (s *GetTrainingJobResponseBodyLatestProgress) GetRemainingTime() *GetTrainingJobResponseBodyLatestProgressRemainingTime {
	return s.RemainingTime
}

func (s *GetTrainingJobResponseBodyLatestProgress) SetOverallProgress(v *GetTrainingJobResponseBodyLatestProgressOverallProgress) *GetTrainingJobResponseBodyLatestProgress {
	s.OverallProgress = v
	return s
}

func (s *GetTrainingJobResponseBodyLatestProgress) SetRemainingTime(v *GetTrainingJobResponseBodyLatestProgressRemainingTime) *GetTrainingJobResponseBodyLatestProgress {
	s.RemainingTime = v
	return s
}

func (s *GetTrainingJobResponseBodyLatestProgress) Validate() error {
	if s.OverallProgress != nil {
		if err := s.OverallProgress.Validate(); err != nil {
			return err
		}
	}
	if s.RemainingTime != nil {
		if err := s.RemainingTime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTrainingJobResponseBodyLatestProgressOverallProgress struct {
	// The progress timestamp.
	//
	// example:
	//
	// 2023-07-04T13:20:18Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The progress value.
	//
	// example:
	//
	// 0.75
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTrainingJobResponseBodyLatestProgressOverallProgress) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyLatestProgressOverallProgress) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyLatestProgressOverallProgress) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetTrainingJobResponseBodyLatestProgressOverallProgress) GetValue() *float32 {
	return s.Value
}

func (s *GetTrainingJobResponseBodyLatestProgressOverallProgress) SetTimestamp(v string) *GetTrainingJobResponseBodyLatestProgressOverallProgress {
	s.Timestamp = &v
	return s
}

func (s *GetTrainingJobResponseBodyLatestProgressOverallProgress) SetValue(v float32) *GetTrainingJobResponseBodyLatestProgressOverallProgress {
	s.Value = &v
	return s
}

func (s *GetTrainingJobResponseBodyLatestProgressOverallProgress) Validate() error {
	return dara.Validate(s)
}

type GetTrainingJobResponseBodyLatestProgressRemainingTime struct {
	// The progress timestamp.
	//
	// example:
	//
	// 2023-07-04T13:20:18Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The remaining time, in seconds.
	//
	// example:
	//
	// 3600
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTrainingJobResponseBodyLatestProgressRemainingTime) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyLatestProgressRemainingTime) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyLatestProgressRemainingTime) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetTrainingJobResponseBodyLatestProgressRemainingTime) GetValue() *int64 {
	return s.Value
}

func (s *GetTrainingJobResponseBodyLatestProgressRemainingTime) SetTimestamp(v string) *GetTrainingJobResponseBodyLatestProgressRemainingTime {
	s.Timestamp = &v
	return s
}

func (s *GetTrainingJobResponseBodyLatestProgressRemainingTime) SetValue(v int64) *GetTrainingJobResponseBodyLatestProgressRemainingTime {
	s.Value = &v
	return s
}

func (s *GetTrainingJobResponseBodyLatestProgressRemainingTime) Validate() error {
	return dara.Validate(s)
}

type GetTrainingJobResponseBodyOutputChannels struct {
	// The dataset ID.
	//
	// example:
	//
	// d-8o0hh35po15ejcdq2p
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The name of the output data.
	//
	// example:
	//
	// model
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The URI of the output data.
	//
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou-internal.aliyuncs.com/path/to/output/model/
	OutputUri *string `json:"OutputUri,omitempty" xml:"OutputUri,omitempty"`
	RoleArn   *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The dataset version.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetTrainingJobResponseBodyOutputChannels) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyOutputChannels) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyOutputChannels) GetDatasetId() *string {
	return s.DatasetId
}

func (s *GetTrainingJobResponseBodyOutputChannels) GetName() *string {
	return s.Name
}

func (s *GetTrainingJobResponseBodyOutputChannels) GetOutputUri() *string {
	return s.OutputUri
}

func (s *GetTrainingJobResponseBodyOutputChannels) GetRoleArn() *string {
	return s.RoleArn
}

func (s *GetTrainingJobResponseBodyOutputChannels) GetVersionName() *string {
	return s.VersionName
}

func (s *GetTrainingJobResponseBodyOutputChannels) SetDatasetId(v string) *GetTrainingJobResponseBodyOutputChannels {
	s.DatasetId = &v
	return s
}

func (s *GetTrainingJobResponseBodyOutputChannels) SetName(v string) *GetTrainingJobResponseBodyOutputChannels {
	s.Name = &v
	return s
}

func (s *GetTrainingJobResponseBodyOutputChannels) SetOutputUri(v string) *GetTrainingJobResponseBodyOutputChannels {
	s.OutputUri = &v
	return s
}

func (s *GetTrainingJobResponseBodyOutputChannels) SetRoleArn(v string) *GetTrainingJobResponseBodyOutputChannels {
	s.RoleArn = &v
	return s
}

func (s *GetTrainingJobResponseBodyOutputChannels) SetVersionName(v string) *GetTrainingJobResponseBodyOutputChannels {
	s.VersionName = &v
	return s
}

func (s *GetTrainingJobResponseBodyOutputChannels) Validate() error {
	return dara.Validate(s)
}

type GetTrainingJobResponseBodyOutputModel struct {
	// The OutputChannel name corresponding to the model.
	//
	// example:
	//
	// model
	OutputChannelName *string `json:"OutputChannelName,omitempty" xml:"OutputChannelName,omitempty"`
	// The model URI.
	//
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou-internal.aliyuncs.com/path/to/model/output/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s GetTrainingJobResponseBodyOutputModel) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyOutputModel) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyOutputModel) GetOutputChannelName() *string {
	return s.OutputChannelName
}

func (s *GetTrainingJobResponseBodyOutputModel) GetUri() *string {
	return s.Uri
}

func (s *GetTrainingJobResponseBodyOutputModel) SetOutputChannelName(v string) *GetTrainingJobResponseBodyOutputModel {
	s.OutputChannelName = &v
	return s
}

func (s *GetTrainingJobResponseBodyOutputModel) SetUri(v string) *GetTrainingJobResponseBodyOutputModel {
	s.Uri = &v
	return s
}

func (s *GetTrainingJobResponseBodyOutputModel) Validate() error {
	return dara.Validate(s)
}

type GetTrainingJobResponseBodyScheduler struct {
	// The maximum runtime in minutes.
	//
	// example:
	//
	// 100
	MaxRunningTimeInMinutes *string `json:"MaxRunningTimeInMinutes,omitempty" xml:"MaxRunningTimeInMinutes,omitempty"`
	// The maximum training runtime in seconds. A value of 0 indicates no limit on the maximum runtime.
	//
	// example:
	//
	// 0
	MaxRunningTimeInSeconds *string `json:"MaxRunningTimeInSeconds,omitempty" xml:"MaxRunningTimeInSeconds,omitempty"`
}

func (s GetTrainingJobResponseBodyScheduler) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyScheduler) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyScheduler) GetMaxRunningTimeInMinutes() *string {
	return s.MaxRunningTimeInMinutes
}

func (s *GetTrainingJobResponseBodyScheduler) GetMaxRunningTimeInSeconds() *string {
	return s.MaxRunningTimeInSeconds
}

func (s *GetTrainingJobResponseBodyScheduler) SetMaxRunningTimeInMinutes(v string) *GetTrainingJobResponseBodyScheduler {
	s.MaxRunningTimeInMinutes = &v
	return s
}

func (s *GetTrainingJobResponseBodyScheduler) SetMaxRunningTimeInSeconds(v string) *GetTrainingJobResponseBodyScheduler {
	s.MaxRunningTimeInSeconds = &v
	return s
}

func (s *GetTrainingJobResponseBodyScheduler) Validate() error {
	return dara.Validate(s)
}

type GetTrainingJobResponseBodyStatusTransitions struct {
	// The end time of the status.
	//
	// example:
	//
	// 2024-07-10T11:49:47Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The status code.
	//
	// example:
	//
	// TrainingJobSucceed
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The status update message.
	//
	// example:
	//
	// KubeDL job runs successfully
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// The start time of the status.
	//
	// example:
	//
	// 2024-07-10T11:49:47Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The training job status.
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTrainingJobResponseBodyStatusTransitions) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyStatusTransitions) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyStatusTransitions) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTrainingJobResponseBodyStatusTransitions) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *GetTrainingJobResponseBodyStatusTransitions) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *GetTrainingJobResponseBodyStatusTransitions) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTrainingJobResponseBodyStatusTransitions) GetStatus() *string {
	return s.Status
}

func (s *GetTrainingJobResponseBodyStatusTransitions) SetEndTime(v string) *GetTrainingJobResponseBodyStatusTransitions {
	s.EndTime = &v
	return s
}

func (s *GetTrainingJobResponseBodyStatusTransitions) SetReasonCode(v string) *GetTrainingJobResponseBodyStatusTransitions {
	s.ReasonCode = &v
	return s
}

func (s *GetTrainingJobResponseBodyStatusTransitions) SetReasonMessage(v string) *GetTrainingJobResponseBodyStatusTransitions {
	s.ReasonMessage = &v
	return s
}

func (s *GetTrainingJobResponseBodyStatusTransitions) SetStartTime(v string) *GetTrainingJobResponseBodyStatusTransitions {
	s.StartTime = &v
	return s
}

func (s *GetTrainingJobResponseBodyStatusTransitions) SetStatus(v string) *GetTrainingJobResponseBodyStatusTransitions {
	s.Status = &v
	return s
}

func (s *GetTrainingJobResponseBodyStatusTransitions) Validate() error {
	return dara.Validate(s)
}

type GetTrainingJobResponseBodyUserVpc struct {
	// The extended CIDR block configuration.
	ExtendedCIDRs []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	// The security group ID.
	//
	// example:
	//
	// sg-abcdef****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID.
	//
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

func (s GetTrainingJobResponseBodyUserVpc) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobResponseBodyUserVpc) GoString() string {
	return s.String()
}

func (s *GetTrainingJobResponseBodyUserVpc) GetExtendedCIDRs() []*string {
	return s.ExtendedCIDRs
}

func (s *GetTrainingJobResponseBodyUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetTrainingJobResponseBodyUserVpc) GetSwitchId() *string {
	return s.SwitchId
}

func (s *GetTrainingJobResponseBodyUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *GetTrainingJobResponseBodyUserVpc) SetExtendedCIDRs(v []*string) *GetTrainingJobResponseBodyUserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *GetTrainingJobResponseBodyUserVpc) SetSecurityGroupId(v string) *GetTrainingJobResponseBodyUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *GetTrainingJobResponseBodyUserVpc) SetSwitchId(v string) *GetTrainingJobResponseBodyUserVpc {
	s.SwitchId = &v
	return s
}

func (s *GetTrainingJobResponseBodyUserVpc) SetVpcId(v string) *GetTrainingJobResponseBodyUserVpc {
	s.VpcId = &v
	return s
}

func (s *GetTrainingJobResponseBodyUserVpc) Validate() error {
	return dara.Validate(s)
}
