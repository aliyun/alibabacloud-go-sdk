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
	// The visibility of the workspace. The default value is `PRIVATE`.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// Configuration properties for the connection, provided as key-value pairs. The required keys depend on the connection type. For details, see the supplementary parameter information.
	//
	// This parameter is required.
	Configs map[string]*string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// The name of the connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-connection
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The type of the connection.
	//
	// example:
	//
	// DashScopeConnection
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The description of the connection.
	//
	// example:
	//
	// Open-source LLM service connection.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// A list of models. This parameter applies to model service connections.
	Models []*CreateConnectionRequestModels `json:"Models,omitempty" xml:"Models,omitempty" type:"Repeated"`
	// Resource metadata for the connection. This parameter is typically used for database connection types.
	ResourceMeta *CreateConnectionRequestResourceMeta `json:"ResourceMeta,omitempty" xml:"ResourceMeta,omitempty" type:"Struct"`
	// Sensitive connection properties that require encryption, such as database credentials or an API key for a model service.
	Secrets map[string]*string `json:"Secrets,omitempty" xml:"Secrets,omitempty"`
	// The ID of the workspace. To get this ID, call the [`ListWorkspaces`](https://help.aliyun.com/document_detail/449124.html) operation.
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
	//
	// example:
	//
	// Language model
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The context length.
	//
	// example:
	//
	// 4096
	MaxModelLength *int64 `json:"MaxModelLength,omitempty" xml:"MaxModelLength,omitempty"`
	// The model identifier. This value corresponds to the `model` parameter in an OpenAI API request.
	//
	// example:
	//
	// model_001
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The model type.
	//
	// example:
	//
	// LLM
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// Specifies whether the model supports deep reasoning and can output the reasoning process as `reasoning_content`.
	//
	// example:
	//
	// true
	SupportReasoning *bool `json:"SupportReasoning,omitempty" xml:"SupportReasoning,omitempty"`
	// Specifies whether the model supports structured output in the OpenAI API\\"s JSON Schema format.
	//
	// example:
	//
	// true
	SupportResponseSchema *bool `json:"SupportResponseSchema,omitempty" xml:"SupportResponseSchema,omitempty"`
	// Specifies whether the model supports visual understanding.
	//
	// example:
	//
	// false
	SupportVision *bool `json:"SupportVision,omitempty" xml:"SupportVision,omitempty"`
	// Specifies whether the model supports tool calling.
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

func (s *CreateConnectionRequestModels) GetMaxModelLength() *int64 {
	return s.MaxModelLength
}

func (s *CreateConnectionRequestModels) GetModel() *string {
	return s.Model
}

func (s *CreateConnectionRequestModels) GetModelType() *string {
	return s.ModelType
}

func (s *CreateConnectionRequestModels) GetSupportReasoning() *bool {
	return s.SupportReasoning
}

func (s *CreateConnectionRequestModels) GetSupportResponseSchema() *bool {
	return s.SupportResponseSchema
}

func (s *CreateConnectionRequestModels) GetSupportVision() *bool {
	return s.SupportVision
}

func (s *CreateConnectionRequestModels) GetToolCall() *bool {
	return s.ToolCall
}

func (s *CreateConnectionRequestModels) SetDisplayName(v string) *CreateConnectionRequestModels {
	s.DisplayName = &v
	return s
}

func (s *CreateConnectionRequestModels) SetMaxModelLength(v int64) *CreateConnectionRequestModels {
	s.MaxModelLength = &v
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

func (s *CreateConnectionRequestModels) SetSupportReasoning(v bool) *CreateConnectionRequestModels {
	s.SupportReasoning = &v
	return s
}

func (s *CreateConnectionRequestModels) SetSupportResponseSchema(v bool) *CreateConnectionRequestModels {
	s.SupportResponseSchema = &v
	return s
}

func (s *CreateConnectionRequestModels) SetSupportVision(v bool) *CreateConnectionRequestModels {
	s.SupportVision = &v
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
	// Additional configuration information.
	//
	// example:
	//
	// {"vpcId":"vpc-xxxx"}
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ld-uf69****9nqjjes
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// Test instance.
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
