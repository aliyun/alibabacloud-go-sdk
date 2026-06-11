// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiPluginStatus interface {
	dara.Model
	String() string
	GoString() string
	SetErrorLogs(v []map[string]interface{}) *AiPluginStatus
	GetErrorLogs() []map[string]interface{}
	SetPluginId(v string) *AiPluginStatus
	GetPluginId() *string
	SetServiceHealthy(v bool) *AiPluginStatus
	GetServiceHealthy() *bool
}

type AiPluginStatus struct {
	// An object containing error logs, where each key is a string identifying an error and the value is a string describing that error.
	ErrorLogs []map[string]interface{} `json:"errorLogs,omitempty" xml:"errorLogs,omitempty" type:"Repeated"`
	// The unique identifier of the plugin.
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	// Indicates if the service is healthy. `true` indicates normal operation, and `false` indicates an issue.
	ServiceHealthy *bool `json:"serviceHealthy,omitempty" xml:"serviceHealthy,omitempty"`
}

func (s AiPluginStatus) String() string {
	return dara.Prettify(s)
}

func (s AiPluginStatus) GoString() string {
	return s.String()
}

func (s *AiPluginStatus) GetErrorLogs() []map[string]interface{} {
	return s.ErrorLogs
}

func (s *AiPluginStatus) GetPluginId() *string {
	return s.PluginId
}

func (s *AiPluginStatus) GetServiceHealthy() *bool {
	return s.ServiceHealthy
}

func (s *AiPluginStatus) SetErrorLogs(v []map[string]interface{}) *AiPluginStatus {
	s.ErrorLogs = v
	return s
}

func (s *AiPluginStatus) SetPluginId(v string) *AiPluginStatus {
	s.PluginId = &v
	return s
}

func (s *AiPluginStatus) SetServiceHealthy(v bool) *AiPluginStatus {
	s.ServiceHealthy = &v
	return s
}

func (s *AiPluginStatus) Validate() error {
	return dara.Validate(s)
}
