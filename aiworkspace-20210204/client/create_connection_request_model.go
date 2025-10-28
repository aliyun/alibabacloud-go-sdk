// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CreateConnectionRequest
	GetAccessibility() *string
	SetConfigs(v map[string]*string) *CreateConnectionRequest
	GetConfigs() map[string]*string
	SetConnectionName(v string) *CreateConnectionRequest
	GetConnectionName() *string
	SetConnectionType(v string) *CreateConnectionRequest
	GetConnectionType() *string
	SetDescription(v string) *CreateConnectionRequest
	GetDescription() *string
	SetModels(v []*CreateConnectionRequestModels) *CreateConnectionRequest
	GetModels() []*CreateConnectionRequestModels
	SetResourceMeta(v *CreateConnectionRequestResourceMeta) *CreateConnectionRequest
	GetResourceMeta() *CreateConnectionRequestResourceMeta
	SetSecrets(v map[string]*string) *CreateConnectionRequest
	GetSecrets() map[string]*string
	SetWorkspaceId(v string) *CreateConnectionRequest
	GetWorkspaceId() *string
}

type CreateConnectionRequest struct {
	// The accessibility of the workspace. Valid values:
	//
	// 	- PRIVATE: The workspace is accessible only to you and the administrator of the workspace. This is the default value.
	//
	// 	- PUBLIC: The workspace is accessible to all users in the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The connection configurations, in key-value pairs. The key varies based on the connection type. For more information, see the supplementary notes below the request parameters.
	//
	// This parameter is required.
	Configs map[string]*string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// The connection name.
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-connection
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The connection type. Valid values:
	//
	// 	- DashScopeConnection: Alibaba Cloud Model Studio connection
	//
	// 	- OpenLLMConnection: open source model connection
	//
	// 	- MilvusConnection: Milvus connection
	//
	// 	- OpenSearchConnection: OpenSearch connection
	//
	// 	- LindormConnection: Lindorm connection
	//
	// 	- ElasticsearchConnection: Elasticsearch connection
	//
	// 	- HologresConnection: Hologres connection
	//
	// 	- RDSConnection: RDS connection
	//
	// 	- CustomConnection: custom connection
	//
	// example:
	//
	// DashScopeConnection
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The connection description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The models, which apply to model service connections.
	Models []*CreateConnectionRequestModels `json:"Models,omitempty" xml:"Models,omitempty" type:"Repeated"`
	// The instance resource information of the connection, which applies to database connections.
	ResourceMeta *CreateConnectionRequestResourceMeta `json:"ResourceMeta,omitempty" xml:"ResourceMeta,omitempty" type:"Struct"`
	// The configuration to be encrypted. Examples: the database logon account and password and the key of the model service.
	Secrets map[string]*string `json:"Secrets,omitempty" xml:"Secrets,omitempty"`
	// The workspace ID. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// example:
	//
	// 123**45
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateConnectionRequest) GetConfigs() map[string]*string {
	return s.Configs
}

func (s *CreateConnectionRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *CreateConnectionRequest) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *CreateConnectionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateConnectionRequest) GetModels() []*CreateConnectionRequestModels {
	return s.Models
}

func (s *CreateConnectionRequest) GetResourceMeta() *CreateConnectionRequestResourceMeta {
	return s.ResourceMeta
}

func (s *CreateConnectionRequest) GetSecrets() map[string]*string {
	return s.Secrets
}

func (s *CreateConnectionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateConnectionRequest) SetAccessibility(v string) *CreateConnectionRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateConnectionRequest) SetConfigs(v map[string]*string) *CreateConnectionRequest {
	s.Configs = v
	return s
}

func (s *CreateConnectionRequest) SetConnectionName(v string) *CreateConnectionRequest {
	s.ConnectionName = &v
	return s
}

func (s *CreateConnectionRequest) SetConnectionType(v string) *CreateConnectionRequest {
	s.ConnectionType = &v
	return s
}

func (s *CreateConnectionRequest) SetDescription(v string) *CreateConnectionRequest {
	s.Description = &v
	return s
}

func (s *CreateConnectionRequest) SetModels(v []*CreateConnectionRequestModels) *CreateConnectionRequest {
	s.Models = v
	return s
}

func (s *CreateConnectionRequest) SetResourceMeta(v *CreateConnectionRequestResourceMeta) *CreateConnectionRequest {
	s.ResourceMeta = v
	return s
}

func (s *CreateConnectionRequest) SetSecrets(v map[string]*string) *CreateConnectionRequest {
	s.Secrets = v
	return s
}

func (s *CreateConnectionRequest) SetWorkspaceId(v string) *CreateConnectionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateConnectionRequest) Validate() error {
	if s.Models != nil {
		for _, item := range s.Models {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceMeta != nil {
		if err := s.ResourceMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateConnectionRequestModels struct {
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
	// Specifies whether a tool can be called by using ToolCall. Valid values:
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

func (s CreateConnectionRequestModels) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionRequestModels) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequestModels) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateConnectionRequestModels) GetModel() *string {
	return s.Model
}

func (s *CreateConnectionRequestModels) GetModelType() *string {
	return s.ModelType
}

func (s *CreateConnectionRequestModels) GetToolCall() *bool {
	return s.ToolCall
}

func (s *CreateConnectionRequestModels) SetDisplayName(v string) *CreateConnectionRequestModels {
	s.DisplayName = &v
	return s
}

func (s *CreateConnectionRequestModels) SetModel(v string) *CreateConnectionRequestModels {
	s.Model = &v
	return s
}

func (s *CreateConnectionRequestModels) SetModelType(v string) *CreateConnectionRequestModels {
	s.ModelType = &v
	return s
}

func (s *CreateConnectionRequestModels) SetToolCall(v bool) *CreateConnectionRequestModels {
	s.ToolCall = &v
	return s
}

func (s *CreateConnectionRequestModels) Validate() error {
	return dara.Validate(s)
}

type CreateConnectionRequestResourceMeta struct {
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ld-uf69****9nqjjes
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s CreateConnectionRequestResourceMeta) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionRequestResourceMeta) GoString() string {
	return s.String()
}

func (s *CreateConnectionRequestResourceMeta) GetExtra() *string {
	return s.Extra
}

func (s *CreateConnectionRequestResourceMeta) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateConnectionRequestResourceMeta) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateConnectionRequestResourceMeta) SetExtra(v string) *CreateConnectionRequestResourceMeta {
	s.Extra = &v
	return s
}

func (s *CreateConnectionRequestResourceMeta) SetInstanceId(v string) *CreateConnectionRequestResourceMeta {
	s.InstanceId = &v
	return s
}

func (s *CreateConnectionRequestResourceMeta) SetInstanceName(v string) *CreateConnectionRequestResourceMeta {
	s.InstanceName = &v
	return s
}

func (s *CreateConnectionRequestResourceMeta) Validate() error {
	return dara.Validate(s)
}
