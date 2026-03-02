// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iComponentCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ComponentCreateCmd
	GetCompanyId() *int64
	SetConfiguration(v []*ConfigModel) *ComponentCreateCmd
	GetConfiguration() []*ConfigModel
	SetDescription(v string) *ComponentCreateCmd
	GetDescription() *string
	SetEnv(v string) *ComponentCreateCmd
	GetEnv() *string
	SetIsCustom(v bool) *ComponentCreateCmd
	GetIsCustom() *bool
	SetName(v string) *ComponentCreateCmd
	GetName() *string
	SetProductId(v int64) *ComponentCreateCmd
	GetProductId() *int64
	SetResourceId(v int64) *ComponentCreateCmd
	GetResourceId() *int64
	SetScopeServiceIds(v string) *ComponentCreateCmd
	GetScopeServiceIds() *string
	SetTemplateId(v int64) *ComponentCreateCmd
	GetTemplateId() *int64
	SetType(v string) *ComponentCreateCmd
	GetType() *string
	SetUseScope(v string) *ComponentCreateCmd
	GetUseScope() *string
	SetUseType(v string) *ComponentCreateCmd
	GetUseType() *string
}

type ComponentCreateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CompanyId     *int64         `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Configuration []*ConfigModel `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Repeated"`
	// example:
	//
	// 用于统一管理资源
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// false
	IsCustom *bool `json:"isCustom,omitempty" xml:"isCustom,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Redis
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// 1
	ResourceId *int64 `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// 1
	ScopeServiceIds *string `json:"scopeServiceIds,omitempty" xml:"scopeServiceIds,omitempty"`
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
	// State
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// INNER
	UseScope *string `json:"useScope,omitempty" xml:"useScope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SHARESHARE
	UseType *string `json:"useType,omitempty" xml:"useType,omitempty"`
}

func (s ComponentCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s ComponentCreateCmd) GoString() string {
	return s.String()
}

func (s *ComponentCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ComponentCreateCmd) GetConfiguration() []*ConfigModel {
	return s.Configuration
}

func (s *ComponentCreateCmd) GetDescription() *string {
	return s.Description
}

func (s *ComponentCreateCmd) GetEnv() *string {
	return s.Env
}

func (s *ComponentCreateCmd) GetIsCustom() *bool {
	return s.IsCustom
}

func (s *ComponentCreateCmd) GetName() *string {
	return s.Name
}

func (s *ComponentCreateCmd) GetProductId() *int64 {
	return s.ProductId
}

func (s *ComponentCreateCmd) GetResourceId() *int64 {
	return s.ResourceId
}

func (s *ComponentCreateCmd) GetScopeServiceIds() *string {
	return s.ScopeServiceIds
}

func (s *ComponentCreateCmd) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ComponentCreateCmd) GetType() *string {
	return s.Type
}

func (s *ComponentCreateCmd) GetUseScope() *string {
	return s.UseScope
}

func (s *ComponentCreateCmd) GetUseType() *string {
	return s.UseType
}

func (s *ComponentCreateCmd) SetCompanyId(v int64) *ComponentCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *ComponentCreateCmd) SetConfiguration(v []*ConfigModel) *ComponentCreateCmd {
	s.Configuration = v
	return s
}

func (s *ComponentCreateCmd) SetDescription(v string) *ComponentCreateCmd {
	s.Description = &v
	return s
}

func (s *ComponentCreateCmd) SetEnv(v string) *ComponentCreateCmd {
	s.Env = &v
	return s
}

func (s *ComponentCreateCmd) SetIsCustom(v bool) *ComponentCreateCmd {
	s.IsCustom = &v
	return s
}

func (s *ComponentCreateCmd) SetName(v string) *ComponentCreateCmd {
	s.Name = &v
	return s
}

func (s *ComponentCreateCmd) SetProductId(v int64) *ComponentCreateCmd {
	s.ProductId = &v
	return s
}

func (s *ComponentCreateCmd) SetResourceId(v int64) *ComponentCreateCmd {
	s.ResourceId = &v
	return s
}

func (s *ComponentCreateCmd) SetScopeServiceIds(v string) *ComponentCreateCmd {
	s.ScopeServiceIds = &v
	return s
}

func (s *ComponentCreateCmd) SetTemplateId(v int64) *ComponentCreateCmd {
	s.TemplateId = &v
	return s
}

func (s *ComponentCreateCmd) SetType(v string) *ComponentCreateCmd {
	s.Type = &v
	return s
}

func (s *ComponentCreateCmd) SetUseScope(v string) *ComponentCreateCmd {
	s.UseScope = &v
	return s
}

func (s *ComponentCreateCmd) SetUseType(v string) *ComponentCreateCmd {
	s.UseType = &v
	return s
}

func (s *ComponentCreateCmd) Validate() error {
	if s.Configuration != nil {
		for _, item := range s.Configuration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
