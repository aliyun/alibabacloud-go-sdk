// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRunConfigurationResponseBody
	GetRequestId() *string
	SetRunConfiguration(v *GetRunConfigurationResponseBodyRunConfiguration) *GetRunConfigurationResponseBody
	GetRunConfiguration() *GetRunConfigurationResponseBodyRunConfiguration
}

type GetRunConfigurationResponseBody struct {
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId        *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	RunConfiguration *GetRunConfigurationResponseBodyRunConfiguration `json:"runConfiguration,omitempty" xml:"runConfiguration,omitempty" type:"Struct"`
}

func (s GetRunConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRunConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetRunConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRunConfigurationResponseBody) GetRunConfiguration() *GetRunConfigurationResponseBodyRunConfiguration {
	return s.RunConfiguration
}

func (s *GetRunConfigurationResponseBody) SetRequestId(v string) *GetRunConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRunConfigurationResponseBody) SetRunConfiguration(v *GetRunConfigurationResponseBodyRunConfiguration) *GetRunConfigurationResponseBody {
	s.RunConfiguration = v
	return s
}

func (s *GetRunConfigurationResponseBody) Validate() error {
	if s.RunConfiguration != nil {
		if err := s.RunConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRunConfigurationResponseBodyRunConfiguration struct {
	// 应用配置项
	ApplicationConfigs []*GetRunConfigurationResponseBodyRunConfigurationApplicationConfigs `json:"applicationConfigs,omitempty" xml:"applicationConfigs,omitempty" type:"Repeated"`
	LogConfig          *GetRunConfigurationResponseBodyRunConfigurationLogConfig            `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	// 运行配置。
	RuntimeConfigs []*Tag `json:"runtimeConfigs,omitempty" xml:"runtimeConfigs,omitempty" type:"Repeated"`
}

func (s GetRunConfigurationResponseBodyRunConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetRunConfigurationResponseBodyRunConfiguration) GoString() string {
	return s.String()
}

func (s *GetRunConfigurationResponseBodyRunConfiguration) GetApplicationConfigs() []*GetRunConfigurationResponseBodyRunConfigurationApplicationConfigs {
	return s.ApplicationConfigs
}

func (s *GetRunConfigurationResponseBodyRunConfiguration) GetLogConfig() *GetRunConfigurationResponseBodyRunConfigurationLogConfig {
	return s.LogConfig
}

func (s *GetRunConfigurationResponseBodyRunConfiguration) GetRuntimeConfigs() []*Tag {
	return s.RuntimeConfigs
}

func (s *GetRunConfigurationResponseBodyRunConfiguration) SetApplicationConfigs(v []*GetRunConfigurationResponseBodyRunConfigurationApplicationConfigs) *GetRunConfigurationResponseBodyRunConfiguration {
	s.ApplicationConfigs = v
	return s
}

func (s *GetRunConfigurationResponseBodyRunConfiguration) SetLogConfig(v *GetRunConfigurationResponseBodyRunConfigurationLogConfig) *GetRunConfigurationResponseBodyRunConfiguration {
	s.LogConfig = v
	return s
}

func (s *GetRunConfigurationResponseBodyRunConfiguration) SetRuntimeConfigs(v []*Tag) *GetRunConfigurationResponseBodyRunConfiguration {
	s.RuntimeConfigs = v
	return s
}

func (s *GetRunConfigurationResponseBodyRunConfiguration) Validate() error {
	if s.ApplicationConfigs != nil {
		for _, item := range s.ApplicationConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LogConfig != nil {
		if err := s.LogConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RuntimeConfigs != nil {
		for _, item := range s.RuntimeConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRunConfigurationResponseBodyRunConfigurationApplicationConfigs struct {
	// 应用配置文件名。 应用配置文件名。 ```spark-defaults.conf```
	//
	// example:
	//
	// spark-defaults.conf
	ConfigFileName *string `json:"configFileName,omitempty" xml:"configFileName,omitempty"`
	// 配置项键。 配置项键。 ```dfs.namenode.checkpoint.period```
	//
	// example:
	//
	// spark.driver.cores
	ConfigItemKey *string `json:"configItemKey,omitempty" xml:"configItemKey,omitempty"`
	// 配置项值。 配置项值。 ```3600s```
	//
	// example:
	//
	// 2
	ConfigItemValue *string `json:"configItemValue,omitempty" xml:"configItemValue,omitempty"`
}

func (s GetRunConfigurationResponseBodyRunConfigurationApplicationConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetRunConfigurationResponseBodyRunConfigurationApplicationConfigs) GoString() string {
	return s.String()
}

func (s *GetRunConfigurationResponseBodyRunConfigurationApplicationConfigs) GetConfigFileName() *string {
	return s.ConfigFileName
}

func (s *GetRunConfigurationResponseBodyRunConfigurationApplicationConfigs) GetConfigItemKey() *string {
	return s.ConfigItemKey
}

func (s *GetRunConfigurationResponseBodyRunConfigurationApplicationConfigs) GetConfigItemValue() *string {
	return s.ConfigItemValue
}

func (s *GetRunConfigurationResponseBodyRunConfigurationApplicationConfigs) SetConfigFileName(v string) *GetRunConfigurationResponseBodyRunConfigurationApplicationConfigs {
	s.ConfigFileName = &v
	return s
}

func (s *GetRunConfigurationResponseBodyRunConfigurationApplicationConfigs) SetConfigItemKey(v string) *GetRunConfigurationResponseBodyRunConfigurationApplicationConfigs {
	s.ConfigItemKey = &v
	return s
}

func (s *GetRunConfigurationResponseBodyRunConfigurationApplicationConfigs) SetConfigItemValue(v string) *GetRunConfigurationResponseBodyRunConfigurationApplicationConfigs {
	s.ConfigItemValue = &v
	return s
}

func (s *GetRunConfigurationResponseBodyRunConfigurationApplicationConfigs) Validate() error {
	return dara.Validate(s)
}

type GetRunConfigurationResponseBodyRunConfigurationLogConfig struct {
	// example:
	//
	// INFO
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// example:
	//
	// oss://test
	LogPath *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
}

func (s GetRunConfigurationResponseBodyRunConfigurationLogConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRunConfigurationResponseBodyRunConfigurationLogConfig) GoString() string {
	return s.String()
}

func (s *GetRunConfigurationResponseBodyRunConfigurationLogConfig) GetLogLevel() *string {
	return s.LogLevel
}

func (s *GetRunConfigurationResponseBodyRunConfigurationLogConfig) GetLogPath() *string {
	return s.LogPath
}

func (s *GetRunConfigurationResponseBodyRunConfigurationLogConfig) SetLogLevel(v string) *GetRunConfigurationResponseBodyRunConfigurationLogConfig {
	s.LogLevel = &v
	return s
}

func (s *GetRunConfigurationResponseBodyRunConfigurationLogConfig) SetLogPath(v string) *GetRunConfigurationResponseBodyRunConfigurationLogConfig {
	s.LogPath = &v
	return s
}

func (s *GetRunConfigurationResponseBodyRunConfigurationLogConfig) Validate() error {
	return dara.Validate(s)
}
