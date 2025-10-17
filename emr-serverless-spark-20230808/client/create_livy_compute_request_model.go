// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivyComputeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *CreateLivyComputeRequest
	GetAuthType() *string
	SetAutoStartConfiguration(v *CreateLivyComputeRequestAutoStartConfiguration) *CreateLivyComputeRequest
	GetAutoStartConfiguration() *CreateLivyComputeRequestAutoStartConfiguration
	SetAutoStopConfiguration(v *CreateLivyComputeRequestAutoStopConfiguration) *CreateLivyComputeRequest
	GetAutoStopConfiguration() *CreateLivyComputeRequestAutoStopConfiguration
	SetCpuLimit(v string) *CreateLivyComputeRequest
	GetCpuLimit() *string
	SetDisplayReleaseVersion(v string) *CreateLivyComputeRequest
	GetDisplayReleaseVersion() *string
	SetEnablePublic(v bool) *CreateLivyComputeRequest
	GetEnablePublic() *bool
	SetEnvironmentId(v string) *CreateLivyComputeRequest
	GetEnvironmentId() *string
	SetFusion(v bool) *CreateLivyComputeRequest
	GetFusion() *bool
	SetLivyServerConf(v string) *CreateLivyComputeRequest
	GetLivyServerConf() *string
	SetLivyVersion(v string) *CreateLivyComputeRequest
	GetLivyVersion() *string
	SetMemoryLimit(v string) *CreateLivyComputeRequest
	GetMemoryLimit() *string
	SetName(v string) *CreateLivyComputeRequest
	GetName() *string
	SetNetworkName(v string) *CreateLivyComputeRequest
	GetNetworkName() *string
	SetQueueName(v string) *CreateLivyComputeRequest
	GetQueueName() *string
	SetReleaseVersion(v string) *CreateLivyComputeRequest
	GetReleaseVersion() *string
	SetRegionId(v string) *CreateLivyComputeRequest
	GetRegionId() *string
}

type CreateLivyComputeRequest struct {
	// example:
	//
	// Token
	AuthType               *string                                         `json:"authType,omitempty" xml:"authType,omitempty"`
	AutoStartConfiguration *CreateLivyComputeRequestAutoStartConfiguration `json:"autoStartConfiguration,omitempty" xml:"autoStartConfiguration,omitempty" type:"Struct"`
	AutoStopConfiguration  *CreateLivyComputeRequestAutoStopConfiguration  `json:"autoStopConfiguration,omitempty" xml:"autoStopConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// 1
	CpuLimit *string `json:"cpuLimit,omitempty" xml:"cpuLimit,omitempty"`
	// example:
	//
	// esr-4.3.0 (Spark 3.5.2, Scala 2.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// example:
	//
	// true
	EnablePublic *bool `json:"enablePublic,omitempty" xml:"enablePublic,omitempty"`
	// example:
	//
	// ev-ctfq0fem1hkhgv4hapng
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// false
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// example:
	//
	// {
	//
	//   "sparkDefaultsConf": "spark.driver.cores     1\\nspark.driver.memory    4g\\nspark.executor.cores   1\\nspark.executor.memory  4g\\n",
	//
	//   "sparkBlackListConf": "spark.driver.cores\\nspark.driver.memory",
	//
	//   "livyConf": "livy.server.session.timeout  1h\\n",
	//
	//   "livyClientConf": "livy.rsc.sql.num-rows  1000\\n"
	//
	// }
	LivyServerConf *string `json:"livyServerConf,omitempty" xml:"livyServerConf,omitempty"`
	// example:
	//
	// 0.8.0
	LivyVersion *string `json:"livyVersion,omitempty" xml:"livyVersion,omitempty"`
	// example:
	//
	// 4Gi
	MemoryLimit *string `json:"memoryLimit,omitempty" xml:"memoryLimit,omitempty"`
	// example:
	//
	// testGateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// test
	NetworkName *string `json:"networkName,omitempty" xml:"networkName,omitempty"`
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// example:
	//
	// esr-4.3.0 (Spark 3.5.2, Scala 2.12, Java Runtime)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateLivyComputeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLivyComputeRequest) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *CreateLivyComputeRequest) GetAutoStartConfiguration() *CreateLivyComputeRequestAutoStartConfiguration {
	return s.AutoStartConfiguration
}

func (s *CreateLivyComputeRequest) GetAutoStopConfiguration() *CreateLivyComputeRequestAutoStopConfiguration {
	return s.AutoStopConfiguration
}

func (s *CreateLivyComputeRequest) GetCpuLimit() *string {
	return s.CpuLimit
}

func (s *CreateLivyComputeRequest) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *CreateLivyComputeRequest) GetEnablePublic() *bool {
	return s.EnablePublic
}

func (s *CreateLivyComputeRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *CreateLivyComputeRequest) GetFusion() *bool {
	return s.Fusion
}

func (s *CreateLivyComputeRequest) GetLivyServerConf() *string {
	return s.LivyServerConf
}

func (s *CreateLivyComputeRequest) GetLivyVersion() *string {
	return s.LivyVersion
}

func (s *CreateLivyComputeRequest) GetMemoryLimit() *string {
	return s.MemoryLimit
}

func (s *CreateLivyComputeRequest) GetName() *string {
	return s.Name
}

func (s *CreateLivyComputeRequest) GetNetworkName() *string {
	return s.NetworkName
}

func (s *CreateLivyComputeRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateLivyComputeRequest) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *CreateLivyComputeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLivyComputeRequest) SetAuthType(v string) *CreateLivyComputeRequest {
	s.AuthType = &v
	return s
}

func (s *CreateLivyComputeRequest) SetAutoStartConfiguration(v *CreateLivyComputeRequestAutoStartConfiguration) *CreateLivyComputeRequest {
	s.AutoStartConfiguration = v
	return s
}

func (s *CreateLivyComputeRequest) SetAutoStopConfiguration(v *CreateLivyComputeRequestAutoStopConfiguration) *CreateLivyComputeRequest {
	s.AutoStopConfiguration = v
	return s
}

func (s *CreateLivyComputeRequest) SetCpuLimit(v string) *CreateLivyComputeRequest {
	s.CpuLimit = &v
	return s
}

func (s *CreateLivyComputeRequest) SetDisplayReleaseVersion(v string) *CreateLivyComputeRequest {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *CreateLivyComputeRequest) SetEnablePublic(v bool) *CreateLivyComputeRequest {
	s.EnablePublic = &v
	return s
}

func (s *CreateLivyComputeRequest) SetEnvironmentId(v string) *CreateLivyComputeRequest {
	s.EnvironmentId = &v
	return s
}

func (s *CreateLivyComputeRequest) SetFusion(v bool) *CreateLivyComputeRequest {
	s.Fusion = &v
	return s
}

func (s *CreateLivyComputeRequest) SetLivyServerConf(v string) *CreateLivyComputeRequest {
	s.LivyServerConf = &v
	return s
}

func (s *CreateLivyComputeRequest) SetLivyVersion(v string) *CreateLivyComputeRequest {
	s.LivyVersion = &v
	return s
}

func (s *CreateLivyComputeRequest) SetMemoryLimit(v string) *CreateLivyComputeRequest {
	s.MemoryLimit = &v
	return s
}

func (s *CreateLivyComputeRequest) SetName(v string) *CreateLivyComputeRequest {
	s.Name = &v
	return s
}

func (s *CreateLivyComputeRequest) SetNetworkName(v string) *CreateLivyComputeRequest {
	s.NetworkName = &v
	return s
}

func (s *CreateLivyComputeRequest) SetQueueName(v string) *CreateLivyComputeRequest {
	s.QueueName = &v
	return s
}

func (s *CreateLivyComputeRequest) SetReleaseVersion(v string) *CreateLivyComputeRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *CreateLivyComputeRequest) SetRegionId(v string) *CreateLivyComputeRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLivyComputeRequest) Validate() error {
	if s.AutoStartConfiguration != nil {
		if err := s.AutoStartConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.AutoStopConfiguration != nil {
		if err := s.AutoStopConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLivyComputeRequestAutoStartConfiguration struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s CreateLivyComputeRequestAutoStartConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateLivyComputeRequestAutoStartConfiguration) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeRequestAutoStartConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *CreateLivyComputeRequestAutoStartConfiguration) SetEnable(v bool) *CreateLivyComputeRequestAutoStartConfiguration {
	s.Enable = &v
	return s
}

func (s *CreateLivyComputeRequestAutoStartConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateLivyComputeRequestAutoStopConfiguration struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 300
	IdleTimeoutMinutes *int64 `json:"idleTimeoutMinutes,omitempty" xml:"idleTimeoutMinutes,omitempty"`
}

func (s CreateLivyComputeRequestAutoStopConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateLivyComputeRequestAutoStopConfiguration) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeRequestAutoStopConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *CreateLivyComputeRequestAutoStopConfiguration) GetIdleTimeoutMinutes() *int64 {
	return s.IdleTimeoutMinutes
}

func (s *CreateLivyComputeRequestAutoStopConfiguration) SetEnable(v bool) *CreateLivyComputeRequestAutoStopConfiguration {
	s.Enable = &v
	return s
}

func (s *CreateLivyComputeRequestAutoStopConfiguration) SetIdleTimeoutMinutes(v int64) *CreateLivyComputeRequestAutoStopConfiguration {
	s.IdleTimeoutMinutes = &v
	return s
}

func (s *CreateLivyComputeRequestAutoStopConfiguration) Validate() error {
	return dara.Validate(s)
}
