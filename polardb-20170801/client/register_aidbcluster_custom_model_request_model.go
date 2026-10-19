// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterAIDBClusterCustomModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainerEnv(v map[string]*string) *RegisterAIDBClusterCustomModelRequest
	GetContainerEnv() map[string]*string
	SetCustomOssBucketName(v string) *RegisterAIDBClusterCustomModelRequest
	GetCustomOssBucketName() *string
	SetCustomOssBucketPath(v string) *RegisterAIDBClusterCustomModelRequest
	GetCustomOssBucketPath() *string
	SetDBInstanceClass(v string) *RegisterAIDBClusterCustomModelRequest
	GetDBInstanceClass() *string
	SetDirectCommand(v string) *RegisterAIDBClusterCustomModelRequest
	GetDirectCommand() *string
	SetDisplayModelName(v string) *RegisterAIDBClusterCustomModelRequest
	GetDisplayModelName() *string
	SetImageName(v string) *RegisterAIDBClusterCustomModelRequest
	GetImageName() *string
	SetInferenceEngine(v string) *RegisterAIDBClusterCustomModelRequest
	GetInferenceEngine() *string
	SetModelName(v string) *RegisterAIDBClusterCustomModelRequest
	GetModelName() *string
	SetModelParams(v map[string]*string) *RegisterAIDBClusterCustomModelRequest
	GetModelParams() map[string]*string
	SetRegionId(v string) *RegisterAIDBClusterCustomModelRequest
	GetRegionId() *string
	SetRegistryPassword(v string) *RegisterAIDBClusterCustomModelRequest
	GetRegistryPassword() *string
	SetRegistryUsername(v string) *RegisterAIDBClusterCustomModelRequest
	GetRegistryUsername() *string
}

type RegisterAIDBClusterCustomModelRequest struct {
	// The container startup environment variables. Variable names cannot be empty or contain whitespace. Variable values can be empty strings.
	//
	// example:
	//
	// {"HF_HOME":"/cache/hf","EMPTY_VALUE":""}
	ContainerEnv map[string]*string `json:"ContainerEnv,omitempty" xml:"ContainerEnv,omitempty"`
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
	ModelParams map[string]*string `json:"ModelParams,omitempty" xml:"ModelParams,omitempty"`
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

func (s RegisterAIDBClusterCustomModelRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterAIDBClusterCustomModelRequest) GoString() string {
	return s.String()
}

func (s *RegisterAIDBClusterCustomModelRequest) GetContainerEnv() map[string]*string {
	return s.ContainerEnv
}

func (s *RegisterAIDBClusterCustomModelRequest) GetCustomOssBucketName() *string {
	return s.CustomOssBucketName
}

func (s *RegisterAIDBClusterCustomModelRequest) GetCustomOssBucketPath() *string {
	return s.CustomOssBucketPath
}

func (s *RegisterAIDBClusterCustomModelRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *RegisterAIDBClusterCustomModelRequest) GetDirectCommand() *string {
	return s.DirectCommand
}

func (s *RegisterAIDBClusterCustomModelRequest) GetDisplayModelName() *string {
	return s.DisplayModelName
}

func (s *RegisterAIDBClusterCustomModelRequest) GetImageName() *string {
	return s.ImageName
}

func (s *RegisterAIDBClusterCustomModelRequest) GetInferenceEngine() *string {
	return s.InferenceEngine
}

func (s *RegisterAIDBClusterCustomModelRequest) GetModelName() *string {
	return s.ModelName
}

func (s *RegisterAIDBClusterCustomModelRequest) GetModelParams() map[string]*string {
	return s.ModelParams
}

func (s *RegisterAIDBClusterCustomModelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RegisterAIDBClusterCustomModelRequest) GetRegistryPassword() *string {
	return s.RegistryPassword
}

func (s *RegisterAIDBClusterCustomModelRequest) GetRegistryUsername() *string {
	return s.RegistryUsername
}

func (s *RegisterAIDBClusterCustomModelRequest) SetContainerEnv(v map[string]*string) *RegisterAIDBClusterCustomModelRequest {
	s.ContainerEnv = v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetCustomOssBucketName(v string) *RegisterAIDBClusterCustomModelRequest {
	s.CustomOssBucketName = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetCustomOssBucketPath(v string) *RegisterAIDBClusterCustomModelRequest {
	s.CustomOssBucketPath = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetDBInstanceClass(v string) *RegisterAIDBClusterCustomModelRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetDirectCommand(v string) *RegisterAIDBClusterCustomModelRequest {
	s.DirectCommand = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetDisplayModelName(v string) *RegisterAIDBClusterCustomModelRequest {
	s.DisplayModelName = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetImageName(v string) *RegisterAIDBClusterCustomModelRequest {
	s.ImageName = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetInferenceEngine(v string) *RegisterAIDBClusterCustomModelRequest {
	s.InferenceEngine = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetModelName(v string) *RegisterAIDBClusterCustomModelRequest {
	s.ModelName = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetModelParams(v map[string]*string) *RegisterAIDBClusterCustomModelRequest {
	s.ModelParams = v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetRegionId(v string) *RegisterAIDBClusterCustomModelRequest {
	s.RegionId = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetRegistryPassword(v string) *RegisterAIDBClusterCustomModelRequest {
	s.RegistryPassword = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) SetRegistryUsername(v string) *RegisterAIDBClusterCustomModelRequest {
	s.RegistryUsername = &v
	return s
}

func (s *RegisterAIDBClusterCustomModelRequest) Validate() error {
	return dara.Validate(s)
}
