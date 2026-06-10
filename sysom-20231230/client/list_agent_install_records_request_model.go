// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentInstallRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int64) *ListAgentInstallRecordsRequest
	GetCurrent() *int64
	SetInstanceId(v string) *ListAgentInstallRecordsRequest
	GetInstanceId() *string
	SetPageSize(v int64) *ListAgentInstallRecordsRequest
	GetPageSize() *int64
	SetPluginId(v string) *ListAgentInstallRecordsRequest
	GetPluginId() *string
	SetPluginVersion(v string) *ListAgentInstallRecordsRequest
	GetPluginVersion() *string
	SetRegion(v string) *ListAgentInstallRecordsRequest
	GetRegion() *string
	SetStatus(v string) *ListAgentInstallRecordsRequest
	GetStatus() *string
}

type ListAgentInstallRecordsRequest struct {
	// Current page number (starting from 1)
	//
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// If this field is specified, the response filters the Agent installation status for the specified instance.
	//
	// example:
	//
	// i-bp118piqcio9tiwgh84b
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// Page size
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// If this parameter is specified, the response filters the installation list for the specified Agent. It can be used together with the plugin_version parameter.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	PluginId *string `json:"plugin_id,omitempty" xml:"plugin_id,omitempty"`
	// This parameter cannot be used alone. It must be used together with plugin_id to filter the installation list for a specified version of the specified Agent.
	//
	// example:
	//
	// 3.4.0-1
	PluginVersion *string `json:"plugin_version,omitempty" xml:"plugin_version,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// Filter widget installation records by status
	//
	// example:
	//
	// Installed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListAgentInstallRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentInstallRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentInstallRecordsRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *ListAgentInstallRecordsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAgentInstallRecordsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAgentInstallRecordsRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *ListAgentInstallRecordsRequest) GetPluginVersion() *string {
	return s.PluginVersion
}

func (s *ListAgentInstallRecordsRequest) GetRegion() *string {
	return s.Region
}

func (s *ListAgentInstallRecordsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAgentInstallRecordsRequest) SetCurrent(v int64) *ListAgentInstallRecordsRequest {
	s.Current = &v
	return s
}

func (s *ListAgentInstallRecordsRequest) SetInstanceId(v string) *ListAgentInstallRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAgentInstallRecordsRequest) SetPageSize(v int64) *ListAgentInstallRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentInstallRecordsRequest) SetPluginId(v string) *ListAgentInstallRecordsRequest {
	s.PluginId = &v
	return s
}

func (s *ListAgentInstallRecordsRequest) SetPluginVersion(v string) *ListAgentInstallRecordsRequest {
	s.PluginVersion = &v
	return s
}

func (s *ListAgentInstallRecordsRequest) SetRegion(v string) *ListAgentInstallRecordsRequest {
	s.Region = &v
	return s
}

func (s *ListAgentInstallRecordsRequest) SetStatus(v string) *ListAgentInstallRecordsRequest {
	s.Status = &v
	return s
}

func (s *ListAgentInstallRecordsRequest) Validate() error {
	return dara.Validate(s)
}
