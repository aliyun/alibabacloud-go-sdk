// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigResourceGroupModelTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelTemplateId(v string) *ConfigResourceGroupModelTemplateRequest
	GetModelTemplateId() *string
	SetResourceGroupId(v string) *ConfigResourceGroupModelTemplateRequest
	GetResourceGroupId() *string
}

type ConfigResourceGroupModelTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ConfigResourceGroupModelTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigResourceGroupModelTemplateRequest) GoString() string {
	return s.String()
}

func (s *ConfigResourceGroupModelTemplateRequest) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *ConfigResourceGroupModelTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigResourceGroupModelTemplateRequest) SetModelTemplateId(v string) *ConfigResourceGroupModelTemplateRequest {
	s.ModelTemplateId = &v
	return s
}

func (s *ConfigResourceGroupModelTemplateRequest) SetResourceGroupId(v string) *ConfigResourceGroupModelTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigResourceGroupModelTemplateRequest) Validate() error {
	return dara.Validate(s)
}
