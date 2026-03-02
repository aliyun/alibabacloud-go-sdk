// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iComponentTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ComponentTemplate
	GetCompanyId() *int64
	SetConfiguration(v []*ConfigModel) *ComponentTemplate
	GetConfiguration() []*ConfigModel
	SetDescription(v string) *ComponentTemplate
	GetDescription() *string
	SetId(v int64) *ComponentTemplate
	GetId() *int64
	SetIsCustom(v bool) *ComponentTemplate
	GetIsCustom() *bool
	SetName(v string) *ComponentTemplate
	GetName() *string
	SetProductId(v int64) *ComponentTemplate
	GetProductId() *int64
	SetRequestId(v string) *ComponentTemplate
	GetRequestId() *string
	SetResourceType(v string) *ComponentTemplate
	GetResourceType() *string
	SetType(v string) *ComponentTemplate
	GetType() *string
	SetUseScope(v string) *ComponentTemplate
	GetUseScope() *string
	SetUseType(v string) *ComponentTemplate
	GetUseType() *string
	SetVersion(v int64) *ComponentTemplate
	GetVersion() *int64
}

type ComponentTemplate struct {
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
	// Redis模板描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Id       *int64 `json:"id,omitempty" xml:"id,omitempty"`
	IsCustom *bool  `json:"isCustom,omitempty" xml:"isCustom,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Redis模板
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
	// State
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
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
	// Inner
	UseScope *string `json:"useScope,omitempty" xml:"useScope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Share
	UseType *string `json:"useType,omitempty" xml:"useType,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ComponentTemplate) String() string {
	return dara.Prettify(s)
}

func (s ComponentTemplate) GoString() string {
	return s.String()
}

func (s *ComponentTemplate) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ComponentTemplate) GetConfiguration() []*ConfigModel {
	return s.Configuration
}

func (s *ComponentTemplate) GetDescription() *string {
	return s.Description
}

func (s *ComponentTemplate) GetId() *int64 {
	return s.Id
}

func (s *ComponentTemplate) GetIsCustom() *bool {
	return s.IsCustom
}

func (s *ComponentTemplate) GetName() *string {
	return s.Name
}

func (s *ComponentTemplate) GetProductId() *int64 {
	return s.ProductId
}

func (s *ComponentTemplate) GetRequestId() *string {
	return s.RequestId
}

func (s *ComponentTemplate) GetResourceType() *string {
	return s.ResourceType
}

func (s *ComponentTemplate) GetType() *string {
	return s.Type
}

func (s *ComponentTemplate) GetUseScope() *string {
	return s.UseScope
}

func (s *ComponentTemplate) GetUseType() *string {
	return s.UseType
}

func (s *ComponentTemplate) GetVersion() *int64 {
	return s.Version
}

func (s *ComponentTemplate) SetCompanyId(v int64) *ComponentTemplate {
	s.CompanyId = &v
	return s
}

func (s *ComponentTemplate) SetConfiguration(v []*ConfigModel) *ComponentTemplate {
	s.Configuration = v
	return s
}

func (s *ComponentTemplate) SetDescription(v string) *ComponentTemplate {
	s.Description = &v
	return s
}

func (s *ComponentTemplate) SetId(v int64) *ComponentTemplate {
	s.Id = &v
	return s
}

func (s *ComponentTemplate) SetIsCustom(v bool) *ComponentTemplate {
	s.IsCustom = &v
	return s
}

func (s *ComponentTemplate) SetName(v string) *ComponentTemplate {
	s.Name = &v
	return s
}

func (s *ComponentTemplate) SetProductId(v int64) *ComponentTemplate {
	s.ProductId = &v
	return s
}

func (s *ComponentTemplate) SetRequestId(v string) *ComponentTemplate {
	s.RequestId = &v
	return s
}

func (s *ComponentTemplate) SetResourceType(v string) *ComponentTemplate {
	s.ResourceType = &v
	return s
}

func (s *ComponentTemplate) SetType(v string) *ComponentTemplate {
	s.Type = &v
	return s
}

func (s *ComponentTemplate) SetUseScope(v string) *ComponentTemplate {
	s.UseScope = &v
	return s
}

func (s *ComponentTemplate) SetUseType(v string) *ComponentTemplate {
	s.UseType = &v
	return s
}

func (s *ComponentTemplate) SetVersion(v int64) *ComponentTemplate {
	s.Version = &v
	return s
}

func (s *ComponentTemplate) Validate() error {
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
