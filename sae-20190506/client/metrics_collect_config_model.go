// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricsCollectConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnablePushToUserSLS(v bool) *MetricsCollectConfig
	GetEnablePushToUserSLS() *bool
	SetLogstoreName(v string) *MetricsCollectConfig
	GetLogstoreName() *string
	SetProjectName(v string) *MetricsCollectConfig
	GetProjectName() *string
}

type MetricsCollectConfig struct {
	// Specifies whether to push metric data to Simple Log Service. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnablePushToUserSLS *bool `json:"EnablePushToUserSLS,omitempty" xml:"EnablePushToUserSLS,omitempty"`
	// The name of the Logstore in Simple Log Service to which the metric data is pushed.
	//
	// example:
	//
	// my-sls-logstorename
	LogstoreName *string `json:"LogstoreName,omitempty" xml:"LogstoreName,omitempty"`
	// The name of the project in Simple Log Service to which the metric data is pushed.
	//
	// example:
	//
	// my-sls-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s MetricsCollectConfig) String() string {
	return dara.Prettify(s)
}

func (s MetricsCollectConfig) GoString() string {
	return s.String()
}

func (s *MetricsCollectConfig) GetEnablePushToUserSLS() *bool {
	return s.EnablePushToUserSLS
}

func (s *MetricsCollectConfig) GetLogstoreName() *string {
	return s.LogstoreName
}

func (s *MetricsCollectConfig) GetProjectName() *string {
	return s.ProjectName
}

func (s *MetricsCollectConfig) SetEnablePushToUserSLS(v bool) *MetricsCollectConfig {
	s.EnablePushToUserSLS = &v
	return s
}

func (s *MetricsCollectConfig) SetLogstoreName(v string) *MetricsCollectConfig {
	s.LogstoreName = &v
	return s
}

func (s *MetricsCollectConfig) SetProjectName(v string) *MetricsCollectConfig {
	s.ProjectName = &v
	return s
}

func (s *MetricsCollectConfig) Validate() error {
	return dara.Validate(s)
}
