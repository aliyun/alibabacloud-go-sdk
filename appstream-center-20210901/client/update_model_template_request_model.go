// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *UpdateModelTemplateRequest
	GetConfig() *string
	SetDescription(v string) *UpdateModelTemplateRequest
	GetDescription() *string
	SetModelTemplateId(v string) *UpdateModelTemplateRequest
	GetModelTemplateId() *string
	SetName(v string) *UpdateModelTemplateRequest
	GetName() *string
}

type UpdateModelTemplateRequest struct {
	// example:
	//
	// {
	//
	// 	"defaults": {
	//
	// 		"model": {
	//
	// 			"primary": "bailian/qwen3.5-plus"
	//
	// 		}
	//
	// 	}
	//
	// }
	Config      *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// example:
	//
	// model-template-001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateModelTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelTemplateRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateModelTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelTemplateRequest) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *UpdateModelTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateModelTemplateRequest) SetConfig(v string) *UpdateModelTemplateRequest {
	s.Config = &v
	return s
}

func (s *UpdateModelTemplateRequest) SetDescription(v string) *UpdateModelTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateModelTemplateRequest) SetModelTemplateId(v string) *UpdateModelTemplateRequest {
	s.ModelTemplateId = &v
	return s
}

func (s *UpdateModelTemplateRequest) SetName(v string) *UpdateModelTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateModelTemplateRequest) Validate() error {
	return dara.Validate(s)
}
