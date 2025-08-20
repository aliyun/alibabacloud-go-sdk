// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfigs(v []*CreateSessionClusterRequestApplicationConfigs) *CreateSessionClusterRequest
	GetApplicationConfigs() []*CreateSessionClusterRequestApplicationConfigs
	SetAutoStartConfiguration(v *CreateSessionClusterRequestAutoStartConfiguration) *CreateSessionClusterRequest
	GetAutoStartConfiguration() *CreateSessionClusterRequestAutoStartConfiguration
	SetAutoStopConfiguration(v *CreateSessionClusterRequestAutoStopConfiguration) *CreateSessionClusterRequest
	GetAutoStopConfiguration() *CreateSessionClusterRequestAutoStopConfiguration
	SetClientToken(v string) *CreateSessionClusterRequest
	GetClientToken() *string
	SetDisplayReleaseVersion(v string) *CreateSessionClusterRequest
	GetDisplayReleaseVersion() *string
	SetEnvId(v string) *CreateSessionClusterRequest
	GetEnvId() *string
	SetFusion(v bool) *CreateSessionClusterRequest
	GetFusion() *bool
	SetKind(v string) *CreateSessionClusterRequest
	GetKind() *string
	SetName(v string) *CreateSessionClusterRequest
	GetName() *string
	SetPublicEndpointEnabled(v bool) *CreateSessionClusterRequest
	GetPublicEndpointEnabled() *bool
	SetQueueName(v string) *CreateSessionClusterRequest
	GetQueueName() *string
	SetReleaseVersion(v string) *CreateSessionClusterRequest
	GetReleaseVersion() *string
	SetRegionId(v string) *CreateSessionClusterRequest
	GetRegionId() *string
}

type CreateSessionClusterRequest struct {
	// The Spark configurations.
	ApplicationConfigs []*CreateSessionClusterRequestApplicationConfigs `json:"applicationConfigs,omitempty" xml:"applicationConfigs,omitempty" type:"Repeated"`
	// Specifies whether to enable automatic startup.
	//
	// 	- true
	//
	// 	- false
	AutoStartConfiguration *CreateSessionClusterRequestAutoStartConfiguration `json:"autoStartConfiguration,omitempty" xml:"autoStartConfiguration,omitempty" type:"Struct"`
	// The automatic termination configuration.
	AutoStopConfiguration *CreateSessionClusterRequestAutoStopConfiguration `json:"autoStopConfiguration,omitempty" xml:"autoStopConfiguration,omitempty" type:"Struct"`
	ClientToken           *string                                           `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The version of the Spark engine.
	//
	// example:
	//
	// esr-3.3.1
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// The ID of the Python environment. This parameter takes effect only for notebook sessions.
	//
	// example:
	//
	// env-cpv569tlhtgndjl86t40
	EnvId *string `json:"envId,omitempty" xml:"envId,omitempty"`
	// Specifies whether to enable Fusion engine for acceleration.
	//
	// example:
	//
	// false
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// The session type.
	//
	// 	- SQL
	//
	// 	- NOTEBOOK
	//
	// example:
	//
	// SQL
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// spark_job_name
	Name                  *string `json:"name,omitempty" xml:"name,omitempty"`
	PublicEndpointEnabled *bool   `json:"publicEndpointEnabled,omitempty" xml:"publicEndpointEnabled,omitempty"`
	// The queue name.
	//
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The version number of Spark.
	//
	// example:
	//
	// esr-3.3.1
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateSessionClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterRequest) GetApplicationConfigs() []*CreateSessionClusterRequestApplicationConfigs {
	return s.ApplicationConfigs
}

func (s *CreateSessionClusterRequest) GetAutoStartConfiguration() *CreateSessionClusterRequestAutoStartConfiguration {
	return s.AutoStartConfiguration
}

func (s *CreateSessionClusterRequest) GetAutoStopConfiguration() *CreateSessionClusterRequestAutoStopConfiguration {
	return s.AutoStopConfiguration
}

func (s *CreateSessionClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSessionClusterRequest) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *CreateSessionClusterRequest) GetEnvId() *string {
	return s.EnvId
}

func (s *CreateSessionClusterRequest) GetFusion() *bool {
	return s.Fusion
}

func (s *CreateSessionClusterRequest) GetKind() *string {
	return s.Kind
}

func (s *CreateSessionClusterRequest) GetName() *string {
	return s.Name
}

func (s *CreateSessionClusterRequest) GetPublicEndpointEnabled() *bool {
	return s.PublicEndpointEnabled
}

func (s *CreateSessionClusterRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateSessionClusterRequest) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *CreateSessionClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSessionClusterRequest) SetApplicationConfigs(v []*CreateSessionClusterRequestApplicationConfigs) *CreateSessionClusterRequest {
	s.ApplicationConfigs = v
	return s
}

func (s *CreateSessionClusterRequest) SetAutoStartConfiguration(v *CreateSessionClusterRequestAutoStartConfiguration) *CreateSessionClusterRequest {
	s.AutoStartConfiguration = v
	return s
}

func (s *CreateSessionClusterRequest) SetAutoStopConfiguration(v *CreateSessionClusterRequestAutoStopConfiguration) *CreateSessionClusterRequest {
	s.AutoStopConfiguration = v
	return s
}

func (s *CreateSessionClusterRequest) SetClientToken(v string) *CreateSessionClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSessionClusterRequest) SetDisplayReleaseVersion(v string) *CreateSessionClusterRequest {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *CreateSessionClusterRequest) SetEnvId(v string) *CreateSessionClusterRequest {
	s.EnvId = &v
	return s
}

func (s *CreateSessionClusterRequest) SetFusion(v bool) *CreateSessionClusterRequest {
	s.Fusion = &v
	return s
}

func (s *CreateSessionClusterRequest) SetKind(v string) *CreateSessionClusterRequest {
	s.Kind = &v
	return s
}

func (s *CreateSessionClusterRequest) SetName(v string) *CreateSessionClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateSessionClusterRequest) SetPublicEndpointEnabled(v bool) *CreateSessionClusterRequest {
	s.PublicEndpointEnabled = &v
	return s
}

func (s *CreateSessionClusterRequest) SetQueueName(v string) *CreateSessionClusterRequest {
	s.QueueName = &v
	return s
}

func (s *CreateSessionClusterRequest) SetReleaseVersion(v string) *CreateSessionClusterRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *CreateSessionClusterRequest) SetRegionId(v string) *CreateSessionClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSessionClusterRequest) Validate() error {
	return dara.Validate(s)
}

type CreateSessionClusterRequestApplicationConfigs struct {
	// The name of the configuration file.
	//
	// example:
	//
	// spark-defaults.conf
	ConfigFileName *string `json:"configFileName,omitempty" xml:"configFileName,omitempty"`
	// The key of SparkConf.
	//
	// example:
	//
	// spark.app.name
	ConfigItemKey *string `json:"configItemKey,omitempty" xml:"configItemKey,omitempty"`
	// The value of SparkConf.
	//
	// example:
	//
	// test
	ConfigItemValue *string `json:"configItemValue,omitempty" xml:"configItemValue,omitempty"`
}

func (s CreateSessionClusterRequestApplicationConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionClusterRequestApplicationConfigs) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterRequestApplicationConfigs) GetConfigFileName() *string {
	return s.ConfigFileName
}

func (s *CreateSessionClusterRequestApplicationConfigs) GetConfigItemKey() *string {
	return s.ConfigItemKey
}

func (s *CreateSessionClusterRequestApplicationConfigs) GetConfigItemValue() *string {
	return s.ConfigItemValue
}

func (s *CreateSessionClusterRequestApplicationConfigs) SetConfigFileName(v string) *CreateSessionClusterRequestApplicationConfigs {
	s.ConfigFileName = &v
	return s
}

func (s *CreateSessionClusterRequestApplicationConfigs) SetConfigItemKey(v string) *CreateSessionClusterRequestApplicationConfigs {
	s.ConfigItemKey = &v
	return s
}

func (s *CreateSessionClusterRequestApplicationConfigs) SetConfigItemValue(v string) *CreateSessionClusterRequestApplicationConfigs {
	s.ConfigItemValue = &v
	return s
}

func (s *CreateSessionClusterRequestApplicationConfigs) Validate() error {
	return dara.Validate(s)
}

type CreateSessionClusterRequestAutoStartConfiguration struct {
	// Specifies whether to enable automatic startup.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s CreateSessionClusterRequestAutoStartConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionClusterRequestAutoStartConfiguration) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterRequestAutoStartConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *CreateSessionClusterRequestAutoStartConfiguration) SetEnable(v bool) *CreateSessionClusterRequestAutoStartConfiguration {
	s.Enable = &v
	return s
}

func (s *CreateSessionClusterRequestAutoStartConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateSessionClusterRequestAutoStopConfiguration struct {
	// Specifies whether to enable automatic termination.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The idle timeout period. The session is automatically terminated when the idle timeout period is exceeded.
	//
	// example:
	//
	// 60
	IdleTimeoutMinutes *int `json:"idleTimeoutMinutes,omitempty" xml:"idleTimeoutMinutes,omitempty"`
}

func (s CreateSessionClusterRequestAutoStopConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionClusterRequestAutoStopConfiguration) GoString() string {
	return s.String()
}

func (s *CreateSessionClusterRequestAutoStopConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *CreateSessionClusterRequestAutoStopConfiguration) GetIdleTimeoutMinutes() *int {
	return s.IdleTimeoutMinutes
}

func (s *CreateSessionClusterRequestAutoStopConfiguration) SetEnable(v bool) *CreateSessionClusterRequestAutoStopConfiguration {
	s.Enable = &v
	return s
}

func (s *CreateSessionClusterRequestAutoStopConfiguration) SetIdleTimeoutMinutes(v int) *CreateSessionClusterRequestAutoStopConfiguration {
	s.IdleTimeoutMinutes = &v
	return s
}

func (s *CreateSessionClusterRequestAutoStopConfiguration) Validate() error {
	return dara.Validate(s)
}
