// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentInstallRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *ListAgentInstallRecordsRequest
	GetXDebugId() *string
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
	SetXSysomInvokeSource(v string) *ListAgentInstallRecordsRequest
	GetXSysomInvokeSource() *string
}

type ListAgentInstallRecordsRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The current page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// The ID of the instance. If you specify this parameter, only the Agent installation records for the specified instance are returned.
	//
	// example:
	//
	// i-bp118piqcio9tiwgh84b
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the Agent. If you specify this parameter, only the installation records for the specified Agent are returned. You can use this parameter together with the plugin_version parameter.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	PluginId *string `json:"plugin_id,omitempty" xml:"plugin_id,omitempty"`
	// The version of the Agent. This parameter cannot be used alone. Use this parameter together with the plugin_id parameter to filter installation records for a specific version of the specified Agent.
	//
	// example:
	//
	// 3.4.0-1
	PluginVersion *string `json:"plugin_version,omitempty" xml:"plugin_version,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// Filters component installation records by status.
	//
	// example:
	//
	// Installed
	Status             *string `json:"status,omitempty" xml:"status,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s ListAgentInstallRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentInstallRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentInstallRecordsRequest) GetXDebugId() *string {
	return s.XDebugId
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

func (s *ListAgentInstallRecordsRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *ListAgentInstallRecordsRequest) SetXDebugId(v string) *ListAgentInstallRecordsRequest {
	s.XDebugId = &v
	return s
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

func (s *ListAgentInstallRecordsRequest) SetXSysomInvokeSource(v string) *ListAgentInstallRecordsRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *ListAgentInstallRecordsRequest) Validate() error {
	return dara.Validate(s)
}
