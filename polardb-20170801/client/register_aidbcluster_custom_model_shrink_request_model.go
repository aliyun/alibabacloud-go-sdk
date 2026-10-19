// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterAIDBClusterCustomModelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainerEnvShrink(v string) *RegisterAIDBClusterCustomModelShrinkRequest
	GetContainerEnvShrink() *string
	SetCustomOssBucketName(v string) *RegisterAIDBClusterCustomModelShrinkRequest
	GetCustomOssBucketName() *string
	SetCustomOssBucketPath(v string) *RegisterAIDBClusterCustomModelShrinkRequest
	GetCustomOssBucketPath() *string
	SetDBInstanceClass(v string) *RegisterAIDBClusterCustomModelShrinkRequest
	GetDBInstanceClass() *string
	SetDirectCommand(v string) *RegisterAIDBClusterCustomModelShrinkRequest
	GetDirectCommand() *string
	SetDisplayModelName(v string) *RegisterAIDBClusterCustomModelShrinkRequest
	GetDisplayModelName() *string
	SetImageName(v string) *RegisterAIDBClusterCustomModelShrinkRequest
	GetImageName() *string
	SetInferenceEngine(v string) *RegisterAIDBClusterCustomModelShrinkRequest
	GetInferenceEngine() *string
	SetModelName(v string) *RegisterAIDBClusterCustomModelShrinkRequest
	GetModelName() *string
	SetModelParamsShrink(v string) *RegisterAIDBClusterCustomModelShrinkRequest
	GetModelParamsShrink() *string
	SetRegionId(v string) *RegisterAIDBClusterCustomModelShrinkRequest
	GetRegionId() *string
	SetRegistryPassword(v string) *RegisterAIDBClusterCustomModelShrinkRequest
	GetRegistryPassword() *string
	SetRegistryUsername(v string) *RegisterAIDBClusterCustomModelShrinkRequest
	GetRegistryUsername() *string
}

type RegisterAIDBClusterCustomModelShrinkRequest struct {
	// The container startup environment variables. Variable names cannot be empty or contain whitespace. Variable values can be empty strings.
	//
	// example:
	//
	// {"HF_HOME":"/cache/hf","EMPTY_VALUE":""}
	ContainerEnvShrink *string `json:"ContainerEnv,omitempty" xml:"ContainerEnv,omitempty"`
	// The name of the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-model-bucket
	CustomOssBucketName *string `json:"CustomOssBucketName,omitempty" xml:"CustomOssBucketName,omitempty"`
	// The model path in the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// models/qwen3
	CustomOssBucketPath *string `json:"CustomOssBucketPath,omitempty" xml:"CustomOssBucketPath,omitempty"`
	// The GPU instance type used for model deployment. Required when configuring deployment parameters.
	//
	// example:
	//
	// polardb.gpu.gn8is.4xlarge
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The direct startup command for the model container.
	//
	// example:
	//
	// python -m vllm.entrypoints.openai.api_server
	DirectCommand *string `json:"DirectCommand,omitempty" xml:"DirectCommand,omitempty"`
	// The display name in the candidate list and the initial invocation name for clients.
	//
	// example:
	//
	// my-qwen3
	DisplayModelName *string `json:"DisplayModelName,omitempty" xml:"DisplayModelName,omitempty"`
	// The full reference of the customer-owned inference image, in the format of registry/repository:tag.
	//
	// example:
	//
	// registry.example.com/customer/model:v1
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The inference engine. Required when configuring deployment parameters. Valid values: vllm and sglang.
	//
	// example:
	//
	// vllm
	InferenceEngine *string `json:"InferenceEngine,omitempty" xml:"InferenceEngine,omitempty"`
	// The custom model registration key and model directory name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Qwen3-32B
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The model startup parameters. Parameter keys must start with --.
	//
	// example:
	//
	// {"--max-model-len":"32768"}
	ModelParamsShrink *string `json:"ModelParams,omitempty" xml:"ModelParams,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The password or temporary token for the customer image repository. Required when ImageName is specified.
	//
	// example:
	//
	// temporary-token
	RegistryPassword *string `json:"RegistryPassword,omitempty" xml:"RegistryPassword,omitempty"`
	// The username for the customer image repository. Required when ImageName is specified.
	//
	// example:
	//
	// temporary-user
	RegistryUsername *string `json:"RegistryUsername,omitempty" xml:"RegistryUsername,omitempty"`
}

func (s RegisterAIDBClusterCustomModelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterAIDBClusterCustomModelShrinkRequest) GoString() string {
	return s.String()
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) GetContainerEnvShrink() *string {
	return s.ContainerEnvShrink
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) GetCustomOssBucketName() *string {
	return s.CustomOssBucketName
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) GetCustomOssBucketPath() *string {
	return s.CustomOssBucketPath
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) GetDirectCommand() *string {
	return s.DirectCommand
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) GetDisplayModelName() *string {
	return s.DisplayModelName
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) GetImageName() *string {
	return s.ImageName
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) GetInferenceEngine() *string {
	return s.InferenceEngine
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) GetModelName() *string {
	return s.ModelName
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) GetModelParamsShrink() *string {
	return s.ModelParamsShrink
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) GetRegistryPassword() *string {
	return s.RegistryPassword
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) GetRegistryUsername() *string {
	return s.RegistryUsername
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) SetContainerEnvShrink(v string) *RegisterAIDBClusterCustomModelShrinkRequest {
	s.ContainerEnvShrink = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) SetCustomOssBucketName(v string) *RegisterAIDBClusterCustomModelShrinkRequest {
	s.CustomOssBucketName = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) SetCustomOssBucketPath(v string) *RegisterAIDBClusterCustomModelShrinkRequest {
	s.CustomOssBucketPath = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) SetDBInstanceClass(v string) *RegisterAIDBClusterCustomModelShrinkRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) SetDirectCommand(v string) *RegisterAIDBClusterCustomModelShrinkRequest {
	s.DirectCommand = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) SetDisplayModelName(v string) *RegisterAIDBClusterCustomModelShrinkRequest {
	s.DisplayModelName = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) SetImageName(v string) *RegisterAIDBClusterCustomModelShrinkRequest {
	s.ImageName = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) SetInferenceEngine(v string) *RegisterAIDBClusterCustomModelShrinkRequest {
	s.InferenceEngine = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) SetModelName(v string) *RegisterAIDBClusterCustomModelShrinkRequest {
	s.ModelName = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) SetModelParamsShrink(v string) *RegisterAIDBClusterCustomModelShrinkRequest {
	s.ModelParamsShrink = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) SetRegionId(v string) *RegisterAIDBClusterCustomModelShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) SetRegistryPassword(v string) *RegisterAIDBClusterCustomModelShrinkRequest {
	s.RegistryPassword = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) SetRegistryUsername(v string) *RegisterAIDBClusterCustomModelShrinkRequest {
	s.RegistryUsername = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
