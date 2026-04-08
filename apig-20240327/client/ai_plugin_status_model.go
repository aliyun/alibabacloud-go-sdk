// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiPluginStatus interface {
	dara.Model
	String() string
	GoString() string
	SetErrorLogs(v map[string]*string) *AiPluginStatus
	GetErrorLogs() map[string]*string
	SetPluginId(v string) *AiPluginStatus
	GetPluginId() *string
	SetServiceHealthy(v bool) *AiPluginStatus
	GetServiceHealthy() *bool
}

type AiPluginStatus struct {
	ErrorLogs      map[string]*string `json:"errorLogs,omitempty" xml:"errorLogs,omitempty"`
	PluginId       *string            `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	ServiceHealthy *bool              `json:"serviceHealthy,omitempty" xml:"serviceHealthy,omitempty"`
}

func (s AiPluginStatus) String() string {
	return dara.Prettify(s)
}

func (s AiPluginStatus) GoString() string {
	return s.String()
}

func (s *AiPluginStatus) GetErrorLogs() map[string]*string {
	return s.ErrorLogs
}

func (s *AiPluginStatus) GetPluginId() *string {
	return s.PluginId
}

func (s *AiPluginStatus) GetServiceHealthy() *bool {
	return s.ServiceHealthy
}

func (s *AiPluginStatus) SetErrorLogs(v map[string]*string) *AiPluginStatus {
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
