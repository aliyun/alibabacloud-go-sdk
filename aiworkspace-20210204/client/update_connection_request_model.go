// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v map[string]*string) *UpdateConnectionRequest
	GetConfigs() map[string]*string
	SetDescription(v string) *UpdateConnectionRequest
	GetDescription() *string
	SetModels(v []*UpdateConnectionRequestModels) *UpdateConnectionRequest
	GetModels() []*UpdateConnectionRequestModels
	SetSecrets(v map[string]*string) *UpdateConnectionRequest
	GetSecrets() map[string]*string
}

type UpdateConnectionRequest struct {
	// The configuration of the connection, specified as key-value pairs. The keys in the Configs parameter vary based on the connection type. For more information, see the request parameters in the CreateConnection topic.
	Configs map[string]*string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// The description of the connection.
	//
	// example:
	//
	// Connection for data labeling.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// A list of model information.
	Models []*UpdateConnectionRequestModels `json:"Models,omitempty" xml:"Models,omitempty" type:"Repeated"`
	// Key-value pairs that require encryption, such as database logon passwords and keys for model connections.
	Secrets map[string]*string `json:"Secrets,omitempty" xml:"Secrets,omitempty"`
}

func (s UpdateConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequest) GetConfigs() map[string]*string {
	return s.Configs
}

func (s *UpdateConnectionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateConnectionRequest) GetModels() []*UpdateConnectionRequestModels {
	return s.Models
}

func (s *UpdateConnectionRequest) GetSecrets() map[string]*string {
	return s.Secrets
}

func (s *UpdateConnectionRequest) SetConfigs(v map[string]*string) *UpdateConnectionRequest {
	s.Configs = v
	return s
}

func (s *UpdateConnectionRequest) SetDescription(v string) *UpdateConnectionRequest {
	s.Description = &v
	return s
}

func (s *UpdateConnectionRequest) SetModels(v []*UpdateConnectionRequestModels) *UpdateConnectionRequest {
	s.Models = v
	return s
}

func (s *UpdateConnectionRequest) SetSecrets(v map[string]*string) *UpdateConnectionRequest {
	s.Secrets = v
	return s
}

func (s *UpdateConnectionRequest) Validate() error {
	if s.Models != nil {
		for _, item := range s.Models {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateConnectionRequestModels struct {
	// The display name of the model.
	//
	// example:
	//
	// car_tag
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The model identifier.
	//
	// example:
	//
	// model_001
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The model type. Valid values:
	//
	// - LLM
	//
	// - Embedding
	//
	// - ReRank
	//
	// example:
	//
	// LLM
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// Indicates whether tool calling is supported. Valid values:
	//
	// - true: Tool calling is supported.
	//
	// - false: Tool calling is not supported.
	//
	// example:
	//
	// true
	ToolCall *bool `json:"ToolCall,omitempty" xml:"ToolCall,omitempty"`
}

func (s UpdateConnectionRequestModels) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionRequestModels) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequestModels) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateConnectionRequestModels) GetModel() *string {
	return s.Model
}

func (s *UpdateConnectionRequestModels) GetModelType() *string {
	return s.ModelType
}

func (s *UpdateConnectionRequestModels) GetToolCall() *bool {
	return s.ToolCall
}

func (s *UpdateConnectionRequestModels) SetDisplayName(v string) *UpdateConnectionRequestModels {
	s.DisplayName = &v
	return s
}

func (s *UpdateConnectionRequestModels) SetModel(v string) *UpdateConnectionRequestModels {
	s.Model = &v
	return s
}

func (s *UpdateConnectionRequestModels) SetModelType(v string) *UpdateConnectionRequestModels {
	s.ModelType = &v
	return s
}

func (s *UpdateConnectionRequestModels) SetToolCall(v bool) *UpdateConnectionRequestModels {
	s.ToolCall = &v
	return s
}

func (s *UpdateConnectionRequestModels) Validate() error {
	return dara.Validate(s)
}
