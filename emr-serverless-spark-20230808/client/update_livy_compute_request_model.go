// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivyComputeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *UpdateLivyComputeRequest
	GetAuthType() *string
	SetAutoStartConfiguration(v *UpdateLivyComputeRequestAutoStartConfiguration) *UpdateLivyComputeRequest
	GetAutoStartConfiguration() *UpdateLivyComputeRequestAutoStartConfiguration
	SetAutoStopConfiguration(v *UpdateLivyComputeRequestAutoStopConfiguration) *UpdateLivyComputeRequest
	GetAutoStopConfiguration() *UpdateLivyComputeRequestAutoStopConfiguration
	SetCpuLimit(v string) *UpdateLivyComputeRequest
	GetCpuLimit() *string
	SetDisplayReleaseVersion(v string) *UpdateLivyComputeRequest
	GetDisplayReleaseVersion() *string
	SetEnablePublic(v bool) *UpdateLivyComputeRequest
	GetEnablePublic() *bool
	SetEnvironmentId(v string) *UpdateLivyComputeRequest
	GetEnvironmentId() *string
	SetFusion(v bool) *UpdateLivyComputeRequest
	GetFusion() *bool
	SetLivyServerConf(v string) *UpdateLivyComputeRequest
	GetLivyServerConf() *string
	SetLivyVersion(v string) *UpdateLivyComputeRequest
	GetLivyVersion() *string
	SetMemoryLimit(v string) *UpdateLivyComputeRequest
	GetMemoryLimit() *string
	SetName(v string) *UpdateLivyComputeRequest
	GetName() *string
	SetNetworkName(v string) *UpdateLivyComputeRequest
	GetNetworkName() *string
	SetQueueName(v string) *UpdateLivyComputeRequest
	GetQueueName() *string
	SetReleaseVersion(v string) *UpdateLivyComputeRequest
	GetReleaseVersion() *string
	SetRegionId(v string) *UpdateLivyComputeRequest
	GetRegionId() *string
}

type UpdateLivyComputeRequest struct {
	// The authentication method.
	//
	// example:
	//
	// Token
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// The auto-start configuration.
	AutoStartConfiguration *UpdateLivyComputeRequestAutoStartConfiguration `json:"autoStartConfiguration,omitempty" xml:"autoStartConfiguration,omitempty" type:"Struct"`
	// The auto-stop configuration.
	AutoStopConfiguration *UpdateLivyComputeRequestAutoStopConfiguration `json:"autoStopConfiguration,omitempty" xml:"autoStopConfiguration,omitempty" type:"Struct"`
	// The number of vCPUs for the Livy server.
	//
	// - 1
	//
	// - 2
	//
	// - 4
	//
	// example:
	//
	// 1
	CpuLimit *string `json:"cpuLimit,omitempty" xml:"cpuLimit,omitempty"`
	// The version number of the Spark engine.
	//
	// example:
	//
	// esr-4.3.0 (Spark 3.5.2, Scala 2.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// The status of the switch for the Internet endpoint.
	EnablePublic *bool `json:"enablePublic,omitempty" xml:"enablePublic,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// ev-cq146allhtgkulp5smk0
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Specifies whether to enable the Fusion engine for acceleration.
	//
	// example:
	//
	// false
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// The configurations of the Livy Gateway. This parameter is in JSON format and supports the following files:
	//
	// - sparkDefaultsConf
	//
	// - sparkBlackListConf
	//
	// - livyConf
	//
	// - livyClientConf
	//
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
	// The Livy version.
	//
	// example:
	//
	// 0.8.0
	LivyVersion *string `json:"livyVersion,omitempty" xml:"livyVersion,omitempty"`
	// The memory size of the Livy server.
	//
	// example:
	//
	// 4Gi
	MemoryLimit *string `json:"memoryLimit,omitempty" xml:"memoryLimit,omitempty"`
	// The name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the network connection.
	//
	// example:
	//
	// test
	NetworkName *string `json:"networkName,omitempty" xml:"networkName,omitempty"`
	// The queue name.
	//
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The version number of the Spark engine. This parameter is deprecated. Use displayReleaseVersion instead.
	//
	// example:
	//
	// esr-4.3.0 (Spark 3.5.2, Scala 2.12, Java Runtime)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s UpdateLivyComputeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivyComputeRequest) GoString() string {
	return s.String()
}

func (s *UpdateLivyComputeRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *UpdateLivyComputeRequest) GetAutoStartConfiguration() *UpdateLivyComputeRequestAutoStartConfiguration {
	return s.AutoStartConfiguration
}

func (s *UpdateLivyComputeRequest) GetAutoStopConfiguration() *UpdateLivyComputeRequestAutoStopConfiguration {
	return s.AutoStopConfiguration
}

func (s *UpdateLivyComputeRequest) GetCpuLimit() *string {
	return s.CpuLimit
}

func (s *UpdateLivyComputeRequest) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *UpdateLivyComputeRequest) GetEnablePublic() *bool {
	return s.EnablePublic
}

func (s *UpdateLivyComputeRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *UpdateLivyComputeRequest) GetFusion() *bool {
	return s.Fusion
}

func (s *UpdateLivyComputeRequest) GetLivyServerConf() *string {
	return s.LivyServerConf
}

func (s *UpdateLivyComputeRequest) GetLivyVersion() *string {
	return s.LivyVersion
}

func (s *UpdateLivyComputeRequest) GetMemoryLimit() *string {
	return s.MemoryLimit
}

func (s *UpdateLivyComputeRequest) GetName() *string {
	return s.Name
}

func (s *UpdateLivyComputeRequest) GetNetworkName() *string {
	return s.NetworkName
}

func (s *UpdateLivyComputeRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *UpdateLivyComputeRequest) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *UpdateLivyComputeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLivyComputeRequest) SetAuthType(v string) *UpdateLivyComputeRequest {
	s.AuthType = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetAutoStartConfiguration(v *UpdateLivyComputeRequestAutoStartConfiguration) *UpdateLivyComputeRequest {
	s.AutoStartConfiguration = v
	return s
}

func (s *UpdateLivyComputeRequest) SetAutoStopConfiguration(v *UpdateLivyComputeRequestAutoStopConfiguration) *UpdateLivyComputeRequest {
	s.AutoStopConfiguration = v
	return s
}

func (s *UpdateLivyComputeRequest) SetCpuLimit(v string) *UpdateLivyComputeRequest {
	s.CpuLimit = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetDisplayReleaseVersion(v string) *UpdateLivyComputeRequest {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetEnablePublic(v bool) *UpdateLivyComputeRequest {
	s.EnablePublic = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetEnvironmentId(v string) *UpdateLivyComputeRequest {
	s.EnvironmentId = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetFusion(v bool) *UpdateLivyComputeRequest {
	s.Fusion = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetLivyServerConf(v string) *UpdateLivyComputeRequest {
	s.LivyServerConf = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetLivyVersion(v string) *UpdateLivyComputeRequest {
	s.LivyVersion = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetMemoryLimit(v string) *UpdateLivyComputeRequest {
	s.MemoryLimit = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetName(v string) *UpdateLivyComputeRequest {
	s.Name = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetNetworkName(v string) *UpdateLivyComputeRequest {
	s.NetworkName = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetQueueName(v string) *UpdateLivyComputeRequest {
	s.QueueName = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetReleaseVersion(v string) *UpdateLivyComputeRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *UpdateLivyComputeRequest) SetRegionId(v string) *UpdateLivyComputeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLivyComputeRequest) Validate() error {
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

type UpdateLivyComputeRequestAutoStartConfiguration struct {
	// Specifies whether to automatically start the Livy Gateway after it is created.
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s UpdateLivyComputeRequestAutoStartConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivyComputeRequestAutoStartConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateLivyComputeRequestAutoStartConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateLivyComputeRequestAutoStartConfiguration) SetEnable(v bool) *UpdateLivyComputeRequestAutoStartConfiguration {
	s.Enable = &v
	return s
}

func (s *UpdateLivyComputeRequestAutoStartConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateLivyComputeRequestAutoStopConfiguration struct {
	// Specifies whether to enable auto-stop for the Livy Gateway.
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The number of minutes after which the Livy Gateway is automatically stopped if it is idle.
	//
	// example:
	//
	// 300
	IdleTimeoutMinutes *int64 `json:"idleTimeoutMinutes,omitempty" xml:"idleTimeoutMinutes,omitempty"`
}

func (s UpdateLivyComputeRequestAutoStopConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivyComputeRequestAutoStopConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateLivyComputeRequestAutoStopConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateLivyComputeRequestAutoStopConfiguration) GetIdleTimeoutMinutes() *int64 {
	return s.IdleTimeoutMinutes
}

func (s *UpdateLivyComputeRequestAutoStopConfiguration) SetEnable(v bool) *UpdateLivyComputeRequestAutoStopConfiguration {
	s.Enable = &v
	return s
}

func (s *UpdateLivyComputeRequestAutoStopConfiguration) SetIdleTimeoutMinutes(v int64) *UpdateLivyComputeRequestAutoStopConfiguration {
	s.IdleTimeoutMinutes = &v
	return s
}

func (s *UpdateLivyComputeRequestAutoStopConfiguration) Validate() error {
	return dara.Validate(s)
}
