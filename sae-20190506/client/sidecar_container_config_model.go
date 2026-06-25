// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSidecarContainerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAcrInstanceId(v string) *SidecarContainerConfig
	GetAcrInstanceId() *string
	SetCommand(v string) *SidecarContainerConfig
	GetCommand() *string
	SetCommandArgs(v string) *SidecarContainerConfig
	GetCommandArgs() *string
	SetConfigMapMountDesc(v string) *SidecarContainerConfig
	GetConfigMapMountDesc() *string
	SetCpu(v int32) *SidecarContainerConfig
	GetCpu() *int32
	SetEmptyDirDesc(v string) *SidecarContainerConfig
	GetEmptyDirDesc() *string
	SetEnvs(v string) *SidecarContainerConfig
	GetEnvs() *string
	SetImageUrl(v string) *SidecarContainerConfig
	GetImageUrl() *string
	SetLiveness(v string) *SidecarContainerConfig
	GetLiveness() *string
	SetMemory(v int32) *SidecarContainerConfig
	GetMemory() *int32
	SetName(v string) *SidecarContainerConfig
	GetName() *string
	SetPostStart(v string) *SidecarContainerConfig
	GetPostStart() *string
	SetPreStop(v string) *SidecarContainerConfig
	GetPreStop() *string
	SetReadiness(v string) *SidecarContainerConfig
	GetReadiness() *string
	SetSecretMountDesc(v string) *SidecarContainerConfig
	GetSecretMountDesc() *string
}

type SidecarContainerConfig struct {
	// The instance ID of the ACR Enterprise Edition. This parameter is required if the `ImageUrl` is from an ACR Enterprise Edition repository.
	//
	// example:
	//
	// cri-xxxxxx
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The startup command for the image. This command overrides the `ENTRYPOINT` defined in the image.
	//
	// example:
	//
	// python
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The arguments for the startup command. This parameter corresponds to `CMD` in the Dockerfile.
	//
	// example:
	//
	// ["a","b"]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The settings for mounting a ConfigMap. Use this to inject configuration data into the container as files.
	//
	// example:
	//
	// [{"configMapId":16,"key":"test","mountPath":"/tmp"}]
	ConfigMapMountDesc *string `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty"`
	// The CPU resources allocated to the container, measured in millicores. For example, a value of 1000 represents 1 vCPU.
	//
	// example:
	//
	// 1000
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The configuration for an `emptyDir` volume. This creates a temporary directory that persists for the life of the application instance.
	//
	// example:
	//
	// [{\\"name\\":\\"workdir\\",\\"mountPath\\":\\"/usr/local/tomcat/webapps\\"}]
	EmptyDirDesc *string `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty"`
	// The environment variables to set in the container. Specify the variables as a JSON array of key-value pairs.
	//
	// example:
	//
	// [{"name":"TEST_ENV_KEY","value":"TEST_ENV_VAR"}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The container image URL.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/demo/nginx:latest
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The configuration for the liveness probe. The liveness probe checks if the container is running. If the probe fails, the system restarts the container.
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// The amount of memory allocated to the container, measured in MB.
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the container.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration for the postStart hook. This hook runs immediately after the container starts to perform initialization tasks.
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The configuration for the preStop hook. This hook runs immediately before the container is terminated to ensure a graceful shutdown.
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The configuration for the readiness probe. The readiness probe checks if the container is ready to handle requests. The system will not direct traffic to a container until its readiness probe succeeds.
	Readiness *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	// Specifies how to mount a Secret. This lets you securely use sensitive data, such as credentials or keys, in your application.
	SecretMountDesc *string `json:"SecretMountDesc,omitempty" xml:"SecretMountDesc,omitempty"`
}

func (s SidecarContainerConfig) String() string {
	return dara.Prettify(s)
}

func (s SidecarContainerConfig) GoString() string {
	return s.String()
}

func (s *SidecarContainerConfig) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *SidecarContainerConfig) GetCommand() *string {
	return s.Command
}

func (s *SidecarContainerConfig) GetCommandArgs() *string {
	return s.CommandArgs
}

func (s *SidecarContainerConfig) GetConfigMapMountDesc() *string {
	return s.ConfigMapMountDesc
}

func (s *SidecarContainerConfig) GetCpu() *int32 {
	return s.Cpu
}

func (s *SidecarContainerConfig) GetEmptyDirDesc() *string {
	return s.EmptyDirDesc
}

func (s *SidecarContainerConfig) GetEnvs() *string {
	return s.Envs
}

func (s *SidecarContainerConfig) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *SidecarContainerConfig) GetLiveness() *string {
	return s.Liveness
}

func (s *SidecarContainerConfig) GetMemory() *int32 {
	return s.Memory
}

func (s *SidecarContainerConfig) GetName() *string {
	return s.Name
}

func (s *SidecarContainerConfig) GetPostStart() *string {
	return s.PostStart
}

func (s *SidecarContainerConfig) GetPreStop() *string {
	return s.PreStop
}

func (s *SidecarContainerConfig) GetReadiness() *string {
	return s.Readiness
}

func (s *SidecarContainerConfig) GetSecretMountDesc() *string {
	return s.SecretMountDesc
}

func (s *SidecarContainerConfig) SetAcrInstanceId(v string) *SidecarContainerConfig {
	s.AcrInstanceId = &v
	return s
}

func (s *SidecarContainerConfig) SetCommand(v string) *SidecarContainerConfig {
	s.Command = &v
	return s
}

func (s *SidecarContainerConfig) SetCommandArgs(v string) *SidecarContainerConfig {
	s.CommandArgs = &v
	return s
}

func (s *SidecarContainerConfig) SetConfigMapMountDesc(v string) *SidecarContainerConfig {
	s.ConfigMapMountDesc = &v
	return s
}

func (s *SidecarContainerConfig) SetCpu(v int32) *SidecarContainerConfig {
	s.Cpu = &v
	return s
}

func (s *SidecarContainerConfig) SetEmptyDirDesc(v string) *SidecarContainerConfig {
	s.EmptyDirDesc = &v
	return s
}

func (s *SidecarContainerConfig) SetEnvs(v string) *SidecarContainerConfig {
	s.Envs = &v
	return s
}

func (s *SidecarContainerConfig) SetImageUrl(v string) *SidecarContainerConfig {
	s.ImageUrl = &v
	return s
}

func (s *SidecarContainerConfig) SetLiveness(v string) *SidecarContainerConfig {
	s.Liveness = &v
	return s
}

func (s *SidecarContainerConfig) SetMemory(v int32) *SidecarContainerConfig {
	s.Memory = &v
	return s
}

func (s *SidecarContainerConfig) SetName(v string) *SidecarContainerConfig {
	s.Name = &v
	return s
}

func (s *SidecarContainerConfig) SetPostStart(v string) *SidecarContainerConfig {
	s.PostStart = &v
	return s
}

func (s *SidecarContainerConfig) SetPreStop(v string) *SidecarContainerConfig {
	s.PreStop = &v
	return s
}

func (s *SidecarContainerConfig) SetReadiness(v string) *SidecarContainerConfig {
	s.Readiness = &v
	return s
}

func (s *SidecarContainerConfig) SetSecretMountDesc(v string) *SidecarContainerConfig {
	s.SecretMountDesc = &v
	return s
}

func (s *SidecarContainerConfig) Validate() error {
	return dara.Validate(s)
}
