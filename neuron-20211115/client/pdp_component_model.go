// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpComponent interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *PdpComponent
	GetCompanyId() *int64
	SetConfiguration(v []*ConfigModel) *PdpComponent
	GetConfiguration() []*ConfigModel
	SetDescription(v string) *PdpComponent
	GetDescription() *string
	SetEnv(v string) *PdpComponent
	GetEnv() *string
	SetId(v int64) *PdpComponent
	GetId() *int64
	SetIsCustom(v bool) *PdpComponent
	GetIsCustom() *bool
	SetName(v string) *PdpComponent
	GetName() *string
	SetProductId(v int64) *PdpComponent
	GetProductId() *int64
	SetRequestId(v string) *PdpComponent
	GetRequestId() *string
	SetResourceId(v int64) *PdpComponent
	GetResourceId() *int64
	SetResourceType(v string) *PdpComponent
	GetResourceType() *string
	SetScope(v string) *PdpComponent
	GetScope() *string
	SetTemplateConfiguration(v []*ConfigModel) *PdpComponent
	GetTemplateConfiguration() []*ConfigModel
	SetTemplateId(v int64) *PdpComponent
	GetTemplateId() *int64
	SetTemplateVersion(v int64) *PdpComponent
	GetTemplateVersion() *int64
	SetType(v string) *PdpComponent
	GetType() *string
	SetUseScope(v string) *PdpComponent
	GetUseScope() *string
	SetUseType(v string) *PdpComponent
	GetUseType() *string
}

type PdpComponent struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// example:
	//
	// {}
	Configuration []*ConfigModel `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Repeated"`
	// example:
	//
	// 组件描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// 1
	Id       *int64 `json:"id,omitempty" xml:"id,omitempty"`
	IsCustom *bool  `json:"isCustom,omitempty" xml:"isCustom,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 组件
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ResourceId *int64 `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// State
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// 1
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateConfiguration []*ConfigModel `json:"templateConfiguration,omitempty" xml:"templateConfiguration,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateVersion *int64 `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CACHE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Inner
	UseScope *string `json:"useScope,omitempty" xml:"useScope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Share
	UseType *string `json:"useType,omitempty" xml:"useType,omitempty"`
}

func (s PdpComponent) String() string {
	return dara.Prettify(s)
}

func (s PdpComponent) GoString() string {
	return s.String()
}

func (s *PdpComponent) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PdpComponent) GetConfiguration() []*ConfigModel {
	return s.Configuration
}

func (s *PdpComponent) GetDescription() *string {
	return s.Description
}

func (s *PdpComponent) GetEnv() *string {
	return s.Env
}

func (s *PdpComponent) GetId() *int64 {
	return s.Id
}

func (s *PdpComponent) GetIsCustom() *bool {
	return s.IsCustom
}

func (s *PdpComponent) GetName() *string {
	return s.Name
}

func (s *PdpComponent) GetProductId() *int64 {
	return s.ProductId
}

func (s *PdpComponent) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpComponent) GetResourceId() *int64 {
	return s.ResourceId
}

func (s *PdpComponent) GetResourceType() *string {
	return s.ResourceType
}

func (s *PdpComponent) GetScope() *string {
	return s.Scope
}

func (s *PdpComponent) GetTemplateConfiguration() []*ConfigModel {
	return s.TemplateConfiguration
}

func (s *PdpComponent) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *PdpComponent) GetTemplateVersion() *int64 {
	return s.TemplateVersion
}

func (s *PdpComponent) GetType() *string {
	return s.Type
}

func (s *PdpComponent) GetUseScope() *string {
	return s.UseScope
}

func (s *PdpComponent) GetUseType() *string {
	return s.UseType
}

func (s *PdpComponent) SetCompanyId(v int64) *PdpComponent {
	s.CompanyId = &v
	return s
}

func (s *PdpComponent) SetConfiguration(v []*ConfigModel) *PdpComponent {
	s.Configuration = v
	return s
}

func (s *PdpComponent) SetDescription(v string) *PdpComponent {
	s.Description = &v
	return s
}

func (s *PdpComponent) SetEnv(v string) *PdpComponent {
	s.Env = &v
	return s
}

func (s *PdpComponent) SetId(v int64) *PdpComponent {
	s.Id = &v
	return s
}

func (s *PdpComponent) SetIsCustom(v bool) *PdpComponent {
	s.IsCustom = &v
	return s
}

func (s *PdpComponent) SetName(v string) *PdpComponent {
	s.Name = &v
	return s
}

func (s *PdpComponent) SetProductId(v int64) *PdpComponent {
	s.ProductId = &v
	return s
}

func (s *PdpComponent) SetRequestId(v string) *PdpComponent {
	s.RequestId = &v
	return s
}

func (s *PdpComponent) SetResourceId(v int64) *PdpComponent {
	s.ResourceId = &v
	return s
}

func (s *PdpComponent) SetResourceType(v string) *PdpComponent {
	s.ResourceType = &v
	return s
}

func (s *PdpComponent) SetScope(v string) *PdpComponent {
	s.Scope = &v
	return s
}

func (s *PdpComponent) SetTemplateConfiguration(v []*ConfigModel) *PdpComponent {
	s.TemplateConfiguration = v
	return s
}

func (s *PdpComponent) SetTemplateId(v int64) *PdpComponent {
	s.TemplateId = &v
	return s
}

func (s *PdpComponent) SetTemplateVersion(v int64) *PdpComponent {
	s.TemplateVersion = &v
	return s
}

func (s *PdpComponent) SetType(v string) *PdpComponent {
	s.Type = &v
	return s
}

func (s *PdpComponent) SetUseScope(v string) *PdpComponent {
	s.UseScope = &v
	return s
}

func (s *PdpComponent) SetUseType(v string) *PdpComponent {
	s.UseType = &v
	return s
}

func (s *PdpComponent) Validate() error {
	if s.Configuration != nil {
		for _, item := range s.Configuration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TemplateConfiguration != nil {
		for _, item := range s.TemplateConfiguration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
