// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterAgentInstallRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListClusterAgentInstallRecordsResponseBody
	GetRequestId() *string
	SetCode(v string) *ListClusterAgentInstallRecordsResponseBody
	GetCode() *string
	SetData(v []*ListClusterAgentInstallRecordsResponseBodyData) *ListClusterAgentInstallRecordsResponseBody
	GetData() []*ListClusterAgentInstallRecordsResponseBodyData
	SetMessage(v string) *ListClusterAgentInstallRecordsResponseBody
	GetMessage() *string
	SetTotal(v int64) *ListClusterAgentInstallRecordsResponseBody
	GetTotal() *int64
}

type ListClusterAgentInstallRecordsResponseBody struct {
	// example:
	//
	// B149FD9C-ED5C-5765-B3AD-05AA4A4D64D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                                           `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListClusterAgentInstallRecordsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 42
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListClusterAgentInstallRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterAgentInstallRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterAgentInstallRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterAgentInstallRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListClusterAgentInstallRecordsResponseBody) GetData() []*ListClusterAgentInstallRecordsResponseBodyData {
	return s.Data
}

func (s *ListClusterAgentInstallRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListClusterAgentInstallRecordsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListClusterAgentInstallRecordsResponseBody) SetRequestId(v string) *ListClusterAgentInstallRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBody) SetCode(v string) *ListClusterAgentInstallRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBody) SetData(v []*ListClusterAgentInstallRecordsResponseBodyData) *ListClusterAgentInstallRecordsResponseBody {
	s.Data = v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBody) SetMessage(v string) *ListClusterAgentInstallRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBody) SetTotal(v int64) *ListClusterAgentInstallRecordsResponseBody {
	s.Total = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClusterAgentInstallRecordsResponseBodyData struct {
	AgentConfigId   *string `json:"agent_config_id,omitempty" xml:"agent_config_id,omitempty"`
	AgentConfigName *string `json:"agent_config_name,omitempty" xml:"agent_config_name,omitempty"`
	// example:
	//
	// cbf7a37bc905d4682a3338b3744810269
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// example:
	//
	// 2024-12-25T15:08:19
	CreatedAt       *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	GrayscaleConfig *string `json:"grayscale_config,omitempty" xml:"grayscale_config,omitempty"`
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	PluginId *string `json:"plugin_id,omitempty" xml:"plugin_id,omitempty"`
	// example:
	//
	// 3.4.0-1
	PluginVersion *string `json:"plugin_version,omitempty" xml:"plugin_version,omitempty"`
	// example:
	//
	// 2024-12-25T15:08:19
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s ListClusterAgentInstallRecordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListClusterAgentInstallRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) GetAgentConfigId() *string {
	return s.AgentConfigId
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) GetAgentConfigName() *string {
	return s.AgentConfigName
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) GetGrayscaleConfig() *string {
	return s.GrayscaleConfig
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) GetPluginId() *string {
	return s.PluginId
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) GetPluginVersion() *string {
	return s.PluginVersion
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) SetAgentConfigId(v string) *ListClusterAgentInstallRecordsResponseBodyData {
	s.AgentConfigId = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) SetAgentConfigName(v string) *ListClusterAgentInstallRecordsResponseBodyData {
	s.AgentConfigName = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) SetClusterId(v string) *ListClusterAgentInstallRecordsResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) SetCreatedAt(v string) *ListClusterAgentInstallRecordsResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) SetGrayscaleConfig(v string) *ListClusterAgentInstallRecordsResponseBodyData {
	s.GrayscaleConfig = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) SetPluginId(v string) *ListClusterAgentInstallRecordsResponseBodyData {
	s.PluginId = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) SetPluginVersion(v string) *ListClusterAgentInstallRecordsResponseBodyData {
	s.PluginVersion = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) SetUpdatedAt(v string) *ListClusterAgentInstallRecordsResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
