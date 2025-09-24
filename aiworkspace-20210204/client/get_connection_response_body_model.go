// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *GetConnectionResponseBody
	GetAccessibility() *string
	SetConfigs(v map[string]*string) *GetConnectionResponseBody
	GetConfigs() map[string]*string
	SetConnectionId(v string) *GetConnectionResponseBody
	GetConnectionId() *string
	SetConnectionName(v string) *GetConnectionResponseBody
	GetConnectionName() *string
	SetConnectionType(v string) *GetConnectionResponseBody
	GetConnectionType() *string
	SetCreator(v string) *GetConnectionResponseBody
	GetCreator() *string
	SetDescription(v string) *GetConnectionResponseBody
	GetDescription() *string
	SetGmtCreateTime(v string) *GetConnectionResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetConnectionResponseBody
	GetGmtModifiedTime() *string
	SetModels(v []*GetConnectionResponseBodyModels) *GetConnectionResponseBody
	GetModels() []*GetConnectionResponseBodyModels
	SetRequestId(v string) *GetConnectionResponseBody
	GetRequestId() *string
	SetResourceMeta(v *GetConnectionResponseBodyResourceMeta) *GetConnectionResponseBody
	GetResourceMeta() *GetConnectionResponseBodyResourceMeta
	SetSecrets(v map[string]*string) *GetConnectionResponseBody
	GetSecrets() map[string]*string
	SetWorkspaceId(v string) *GetConnectionResponseBody
	GetWorkspaceId() *string
}

type GetConnectionResponseBody struct {
	// The resource accessibility. Valid values:
	//
	// 	- PUBLIC: All members in the workspace can access the workspace.
	//
	// 	- PRIVATE: Only the creator can access the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The connection configuration.
	Configs map[string]*string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// The connection ID.
	//
	// example:
	//
	// conn-pai9m***mi47
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// The connection name.
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The type of the connection. Valid values:
	//
	// 	- DashScopeConnection: Alibaba Cloud Model Studio connection.
	//
	// 	- OpenLLMConnection: Open source model connection.
	//
	// 	- MilvusConnection: Milvus connection.
	//
	// 	- OpenSearchConnection: OpenSearch connection.
	//
	// 	- LindormConnection: Lindorm connection.
	//
	// 	- ElasticsearchConnection: Elasticsearch connection.
	//
	// 	- HologresConnection: Hologres connection.
	//
	// 	- RDSConnection: RDS connection.
	//
	// 	- CustomConnection: Custom connection.
	//
	// example:
	//
	// OpenSearchConnection
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The creator of the connection.
	//
	// example:
	//
	// 28632***898231
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The connection description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the connection is created, in UTC. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2025-03-07T07:54:56Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the connection is modified, in UTC. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2025-03-07T07:54:56Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The models, which apply to model service connections.
	Models []*GetConnectionResponseBodyModels `json:"Models,omitempty" xml:"Models,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The instance resource information of the connection, which applies to database connections.
	ResourceMeta *GetConnectionResponseBodyResourceMeta `json:"ResourceMeta,omitempty" xml:"ResourceMeta,omitempty" type:"Struct"`
	// The encrypted configuration, in key-value pairs. Examples: the database logon password and the key of the model connection.
	Secrets map[string]*string `json:"Secrets,omitempty" xml:"Secrets,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 11**43
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBody) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetConnectionResponseBody) GetConfigs() map[string]*string {
	return s.Configs
}

func (s *GetConnectionResponseBody) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *GetConnectionResponseBody) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *GetConnectionResponseBody) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *GetConnectionResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetConnectionResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetConnectionResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetConnectionResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetConnectionResponseBody) GetModels() []*GetConnectionResponseBodyModels {
	return s.Models
}

func (s *GetConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConnectionResponseBody) GetResourceMeta() *GetConnectionResponseBodyResourceMeta {
	return s.ResourceMeta
}

func (s *GetConnectionResponseBody) GetSecrets() map[string]*string {
	return s.Secrets
}

func (s *GetConnectionResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetConnectionResponseBody) SetAccessibility(v string) *GetConnectionResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetConnectionResponseBody) SetConfigs(v map[string]*string) *GetConnectionResponseBody {
	s.Configs = v
	return s
}

func (s *GetConnectionResponseBody) SetConnectionId(v string) *GetConnectionResponseBody {
	s.ConnectionId = &v
	return s
}

func (s *GetConnectionResponseBody) SetConnectionName(v string) *GetConnectionResponseBody {
	s.ConnectionName = &v
	return s
}

func (s *GetConnectionResponseBody) SetConnectionType(v string) *GetConnectionResponseBody {
	s.ConnectionType = &v
	return s
}

func (s *GetConnectionResponseBody) SetCreator(v string) *GetConnectionResponseBody {
	s.Creator = &v
	return s
}

func (s *GetConnectionResponseBody) SetDescription(v string) *GetConnectionResponseBody {
	s.Description = &v
	return s
}

func (s *GetConnectionResponseBody) SetGmtCreateTime(v string) *GetConnectionResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetConnectionResponseBody) SetGmtModifiedTime(v string) *GetConnectionResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetConnectionResponseBody) SetModels(v []*GetConnectionResponseBodyModels) *GetConnectionResponseBody {
	s.Models = v
	return s
}

func (s *GetConnectionResponseBody) SetRequestId(v string) *GetConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectionResponseBody) SetResourceMeta(v *GetConnectionResponseBodyResourceMeta) *GetConnectionResponseBody {
	s.ResourceMeta = v
	return s
}

func (s *GetConnectionResponseBody) SetSecrets(v map[string]*string) *GetConnectionResponseBody {
	s.Secrets = v
	return s
}

func (s *GetConnectionResponseBody) SetWorkspaceId(v string) *GetConnectionResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetConnectionResponseBodyModels struct {
	// The display name of the model.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The model identifier.
	//
	// example:
	//
	// model_001
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The model type. Valid values:
	//
	// 	- LLM
	//
	// 	- Embedding
	//
	// 	- ReRank
	//
	// example:
	//
	// LLM
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// Indicates whether a tool can be called by using ToolCall. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ToolCall *bool `json:"ToolCall,omitempty" xml:"ToolCall,omitempty"`
}

func (s GetConnectionResponseBodyModels) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionResponseBodyModels) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyModels) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetConnectionResponseBodyModels) GetModel() *string {
	return s.Model
}

func (s *GetConnectionResponseBodyModels) GetModelType() *string {
	return s.ModelType
}

func (s *GetConnectionResponseBodyModels) GetToolCall() *bool {
	return s.ToolCall
}

func (s *GetConnectionResponseBodyModels) SetDisplayName(v string) *GetConnectionResponseBodyModels {
	s.DisplayName = &v
	return s
}

func (s *GetConnectionResponseBodyModels) SetModel(v string) *GetConnectionResponseBodyModels {
	s.Model = &v
	return s
}

func (s *GetConnectionResponseBodyModels) SetModelType(v string) *GetConnectionResponseBodyModels {
	s.ModelType = &v
	return s
}

func (s *GetConnectionResponseBodyModels) SetToolCall(v bool) *GetConnectionResponseBodyModels {
	s.ToolCall = &v
	return s
}

func (s *GetConnectionResponseBodyModels) Validate() error {
	return dara.Validate(s)
}

type GetConnectionResponseBodyResourceMeta struct {
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ld-2vc1***v1zaqgzol
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s GetConnectionResponseBodyResourceMeta) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionResponseBodyResourceMeta) GoString() string {
	return s.String()
}

func (s *GetConnectionResponseBodyResourceMeta) GetExtra() *string {
	return s.Extra
}

func (s *GetConnectionResponseBodyResourceMeta) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetConnectionResponseBodyResourceMeta) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetConnectionResponseBodyResourceMeta) SetExtra(v string) *GetConnectionResponseBodyResourceMeta {
	s.Extra = &v
	return s
}

func (s *GetConnectionResponseBodyResourceMeta) SetInstanceId(v string) *GetConnectionResponseBodyResourceMeta {
	s.InstanceId = &v
	return s
}

func (s *GetConnectionResponseBodyResourceMeta) SetInstanceName(v string) *GetConnectionResponseBodyResourceMeta {
	s.InstanceName = &v
	return s
}

func (s *GetConnectionResponseBodyResourceMeta) Validate() error {
	return dara.Validate(s)
}
