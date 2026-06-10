// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterAgentInstallRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
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
}

type ListClusterAgentInstallRecordsRequest struct {
	AgentConfigId *string `json:"agent_config_id,omitempty" xml:"agent_config_id,omitempty"`
	// Filter by cluster ID.
	//
	// > This cluster ID is not the ACK cluster ID. Instead, it refers to the `cluster_id` field in the data returned by this API or the `id` field in the data returned by the ListCluster API.
	//
	// example:
	//
	// cbd80af02b9d6454ebdc579c5e022d0c8
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// Current page number (starting from 1)
	//
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// If this parameter is specified, the response filters the installation list for the specified agent. It can be used together with the plugin_version parameter.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	PluginId *string `json:"plugin_id,omitempty" xml:"plugin_id,omitempty"`
	// This parameter cannot be used alone. It must be used together with plugin_id to filter the installation list for a specific agent version.
	//
	// example:
	//
	// 3.4.0-1
	PluginVersion *string `json:"plugin_version,omitempty" xml:"plugin_version,omitempty"`
}

func (s ListClusterAgentInstallRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterAgentInstallRecordsRequest) GoString() string {
	return s.String()
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

func (s *ListClusterAgentInstallRecordsRequest) Validate() error {
	return dara.Validate(s)
}
