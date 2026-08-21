// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterAgentInstallRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *ListClusterAgentInstallRecordsRequest
	GetXDebugId() *string
	SetAgentConfigId(v string) *ListClusterAgentInstallRecordsRequest
	GetAgentConfigId() *string
	SetClusterId(v string) *ListClusterAgentInstallRecordsRequest
	GetClusterId() *string
	SetCurrent(v int64) *ListClusterAgentInstallRecordsRequest
	GetCurrent() *int64
	SetPageSize(v int64) *ListClusterAgentInstallRecordsRequest
	GetPageSize() *int64
	SetPluginId(v string) *ListClusterAgentInstallRecordsRequest
	GetPluginId() *string
	SetPluginVersion(v string) *ListClusterAgentInstallRecordsRequest
	GetPluginVersion() *string
	SetXSysomInvokeSource(v string) *ListClusterAgentInstallRecordsRequest
	GetXSysomInvokeSource() *string
}

type ListClusterAgentInstallRecordsRequest struct {
	XDebugId      *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	AgentConfigId *string `json:"agent_config_id,omitempty" xml:"agent_config_id,omitempty"`
	// Filters by cluster ID.
	//
	// > This cluster ID is not the ACK cluster ID. It is the `cluster_id` field in the data returned by this operation, or the `id` field in the data returned by the ListCluster operation.
	//
	// example:
	//
	// cbd80af02b9d6454ebdc579c5e022d0c8
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The current page number (starting from 1).
	//
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Specifies the agent ID to filter the installation list for the specified agent. This parameter can be used together with the plugin_version parameter.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	PluginId *string `json:"plugin_id,omitempty" xml:"plugin_id,omitempty"`
	// Cannot be used alone. Use this parameter together with plugin_id to filter the installation list for a specified version of the specified agent.
	//
	// example:
	//
	// 3.4.0-1
	PluginVersion      *string `json:"plugin_version,omitempty" xml:"plugin_version,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s ListClusterAgentInstallRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterAgentInstallRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListClusterAgentInstallRecordsRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *ListClusterAgentInstallRecordsRequest) GetAgentConfigId() *string {
	return s.AgentConfigId
}

func (s *ListClusterAgentInstallRecordsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterAgentInstallRecordsRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *ListClusterAgentInstallRecordsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListClusterAgentInstallRecordsRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *ListClusterAgentInstallRecordsRequest) GetPluginVersion() *string {
	return s.PluginVersion
}

func (s *ListClusterAgentInstallRecordsRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *ListClusterAgentInstallRecordsRequest) SetXDebugId(v string) *ListClusterAgentInstallRecordsRequest {
	s.XDebugId = &v
	return s
}

func (s *ListClusterAgentInstallRecordsRequest) SetAgentConfigId(v string) *ListClusterAgentInstallRecordsRequest {
	s.AgentConfigId = &v
	return s
}

func (s *ListClusterAgentInstallRecordsRequest) SetClusterId(v string) *ListClusterAgentInstallRecordsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterAgentInstallRecordsRequest) SetCurrent(v int64) *ListClusterAgentInstallRecordsRequest {
	s.Current = &v
	return s
}

func (s *ListClusterAgentInstallRecordsRequest) SetPageSize(v int64) *ListClusterAgentInstallRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListClusterAgentInstallRecordsRequest) SetPluginId(v string) *ListClusterAgentInstallRecordsRequest {
	s.PluginId = &v
	return s
}

func (s *ListClusterAgentInstallRecordsRequest) SetPluginVersion(v string) *ListClusterAgentInstallRecordsRequest {
	s.PluginVersion = &v
	return s
}

func (s *ListClusterAgentInstallRecordsRequest) SetXSysomInvokeSource(v string) *ListClusterAgentInstallRecordsRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *ListClusterAgentInstallRecordsRequest) Validate() error {
	return dara.Validate(s)
}
