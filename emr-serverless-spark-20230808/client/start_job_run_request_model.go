// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartJobRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartJobRunRequest
	GetClientToken() *string
	SetCodeType(v string) *StartJobRunRequest
	GetCodeType() *string
	SetConfigurationOverrides(v *StartJobRunRequestConfigurationOverrides) *StartJobRunRequest
	GetConfigurationOverrides() *StartJobRunRequestConfigurationOverrides
	SetDisplayReleaseVersion(v string) *StartJobRunRequest
	GetDisplayReleaseVersion() *string
	SetExecutionTimeoutSeconds(v int32) *StartJobRunRequest
	GetExecutionTimeoutSeconds() *int32
	SetFusion(v bool) *StartJobRunRequest
	GetFusion() *bool
	SetJobDriver(v *JobDriver) *StartJobRunRequest
	GetJobDriver() *JobDriver
	SetJobId(v string) *StartJobRunRequest
	GetJobId() *string
	SetName(v string) *StartJobRunRequest
	GetName() *string
	SetReleaseVersion(v string) *StartJobRunRequest
	GetReleaseVersion() *string
	SetResourceQueueId(v string) *StartJobRunRequest
	GetResourceQueueId() *string
	SetTags(v []*Tag) *StartJobRunRequest
	GetTags() []*Tag
	SetRegionId(v string) *StartJobRunRequest
	GetRegionId() *string
}

type StartJobRunRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 8e6aae2810c8f67229ca70bb31cd6028
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The code type of the job. Valid values:
	//
	// 	- SQL
	//
	// 	- JAR
	//
	// 	- PYTHON
	//
	// example:
	//
	// SQL
	CodeType *string `json:"codeType,omitempty" xml:"codeType,omitempty"`
	// The advanced configurations of Spark.
	ConfigurationOverrides *StartJobRunRequestConfigurationOverrides `json:"configurationOverrides,omitempty" xml:"configurationOverrides,omitempty" type:"Struct"`
	// The version of the Spark engine.
	//
	// example:
	//
	// esr-3.3.1
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// The timeout period of the job.
	//
	// example:
	//
	// 100
	ExecutionTimeoutSeconds *int32 `json:"executionTimeoutSeconds,omitempty" xml:"executionTimeoutSeconds,omitempty"`
	// Specifies whether to enable Fusion engine for acceleration.
	//
	// example:
	//
	// false
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// The information about Spark Driver.
	JobDriver *JobDriver `json:"jobDriver,omitempty" xml:"jobDriver,omitempty"`
	// The job ID.
	//
	// example:
	//
	// jr-12345
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// spark_job_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version number of Spark.
	//
	// example:
	//
	// esr-3.3.1
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The name of the resource queue on which the Spark job runs.
	//
	// example:
	//
	// dev_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	// The tags of the job.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s StartJobRunRequest) String() string {
	return dara.Prettify(s)
}

func (s StartJobRunRequest) GoString() string {
	return s.String()
}

func (s *StartJobRunRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartJobRunRequest) GetCodeType() *string {
	return s.CodeType
}

func (s *StartJobRunRequest) GetConfigurationOverrides() *StartJobRunRequestConfigurationOverrides {
	return s.ConfigurationOverrides
}

func (s *StartJobRunRequest) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *StartJobRunRequest) GetExecutionTimeoutSeconds() *int32 {
	return s.ExecutionTimeoutSeconds
}

func (s *StartJobRunRequest) GetFusion() *bool {
	return s.Fusion
}

func (s *StartJobRunRequest) GetJobDriver() *JobDriver {
	return s.JobDriver
}

func (s *StartJobRunRequest) GetJobId() *string {
	return s.JobId
}

func (s *StartJobRunRequest) GetName() *string {
	return s.Name
}

func (s *StartJobRunRequest) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *StartJobRunRequest) GetResourceQueueId() *string {
	return s.ResourceQueueId
}

func (s *StartJobRunRequest) GetTags() []*Tag {
	return s.Tags
}

func (s *StartJobRunRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartJobRunRequest) SetClientToken(v string) *StartJobRunRequest {
	s.ClientToken = &v
	return s
}

func (s *StartJobRunRequest) SetCodeType(v string) *StartJobRunRequest {
	s.CodeType = &v
	return s
}

func (s *StartJobRunRequest) SetConfigurationOverrides(v *StartJobRunRequestConfigurationOverrides) *StartJobRunRequest {
	s.ConfigurationOverrides = v
	return s
}

func (s *StartJobRunRequest) SetDisplayReleaseVersion(v string) *StartJobRunRequest {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *StartJobRunRequest) SetExecutionTimeoutSeconds(v int32) *StartJobRunRequest {
	s.ExecutionTimeoutSeconds = &v
	return s
}

func (s *StartJobRunRequest) SetFusion(v bool) *StartJobRunRequest {
	s.Fusion = &v
	return s
}

func (s *StartJobRunRequest) SetJobDriver(v *JobDriver) *StartJobRunRequest {
	s.JobDriver = v
	return s
}

func (s *StartJobRunRequest) SetJobId(v string) *StartJobRunRequest {
	s.JobId = &v
	return s
}

func (s *StartJobRunRequest) SetName(v string) *StartJobRunRequest {
	s.Name = &v
	return s
}

func (s *StartJobRunRequest) SetReleaseVersion(v string) *StartJobRunRequest {
	s.ReleaseVersion = &v
	return s
}

func (s *StartJobRunRequest) SetResourceQueueId(v string) *StartJobRunRequest {
	s.ResourceQueueId = &v
	return s
}

func (s *StartJobRunRequest) SetTags(v []*Tag) *StartJobRunRequest {
	s.Tags = v
	return s
}

func (s *StartJobRunRequest) SetRegionId(v string) *StartJobRunRequest {
	s.RegionId = &v
	return s
}

func (s *StartJobRunRequest) Validate() error {
	if s.ConfigurationOverrides != nil {
		if err := s.ConfigurationOverrides.Validate(); err != nil {
			return err
		}
	}
	if s.JobDriver != nil {
		if err := s.JobDriver.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type StartJobRunRequestConfigurationOverrides struct {
	// The SparkConf objects.
	Configurations []*StartJobRunRequestConfigurationOverridesConfigurations `json:"configurations,omitempty" xml:"configurations,omitempty" type:"Repeated"`
}

func (s StartJobRunRequestConfigurationOverrides) String() string {
	return dara.Prettify(s)
}

func (s StartJobRunRequestConfigurationOverrides) GoString() string {
	return s.String()
}

func (s *StartJobRunRequestConfigurationOverrides) GetConfigurations() []*StartJobRunRequestConfigurationOverridesConfigurations {
	return s.Configurations
}

func (s *StartJobRunRequestConfigurationOverrides) SetConfigurations(v []*StartJobRunRequestConfigurationOverridesConfigurations) *StartJobRunRequestConfigurationOverrides {
	s.Configurations = v
	return s
}

func (s *StartJobRunRequestConfigurationOverrides) Validate() error {
	if s.Configurations != nil {
		for _, item := range s.Configurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type StartJobRunRequestConfigurationOverridesConfigurations struct {
	// The configuration file of SparkConf.
	//
	// example:
	//
	// spark-default.conf
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
	// test_app
	ConfigItemValue *string `json:"configItemValue,omitempty" xml:"configItemValue,omitempty"`
}

func (s StartJobRunRequestConfigurationOverridesConfigurations) String() string {
	return dara.Prettify(s)
}

func (s StartJobRunRequestConfigurationOverridesConfigurations) GoString() string {
	return s.String()
}

func (s *StartJobRunRequestConfigurationOverridesConfigurations) GetConfigFileName() *string {
	return s.ConfigFileName
}

func (s *StartJobRunRequestConfigurationOverridesConfigurations) GetConfigItemKey() *string {
	return s.ConfigItemKey
}

func (s *StartJobRunRequestConfigurationOverridesConfigurations) GetConfigItemValue() *string {
	return s.ConfigItemValue
}

func (s *StartJobRunRequestConfigurationOverridesConfigurations) SetConfigFileName(v string) *StartJobRunRequestConfigurationOverridesConfigurations {
	s.ConfigFileName = &v
	return s
}

func (s *StartJobRunRequestConfigurationOverridesConfigurations) SetConfigItemKey(v string) *StartJobRunRequestConfigurationOverridesConfigurations {
	s.ConfigItemKey = &v
	return s
}

func (s *StartJobRunRequestConfigurationOverridesConfigurations) SetConfigItemValue(v string) *StartJobRunRequestConfigurationOverridesConfigurations {
	s.ConfigItemValue = &v
	return s
}

func (s *StartJobRunRequestConfigurationOverridesConfigurations) Validate() error {
	return dara.Validate(s)
}
