// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iComponentTemplateCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ComponentTemplateCreateCmd
	GetCompanyId() *int64
	SetConfiguration(v []*ConfigModel) *ComponentTemplateCreateCmd
	GetConfiguration() []*ConfigModel
	SetDescription(v string) *ComponentTemplateCreateCmd
	GetDescription() *string
	SetIsCustom(v bool) *ComponentTemplateCreateCmd
	GetIsCustom() *bool
	SetName(v string) *ComponentTemplateCreateCmd
	GetName() *string
	SetProductId(v int64) *ComponentTemplateCreateCmd
	GetProductId() *int64
	SetType(v string) *ComponentTemplateCreateCmd
	GetType() *string
	SetUseScope(v string) *ComponentTemplateCreateCmd
	GetUseScope() *string
	SetUseType(v string) *ComponentTemplateCreateCmd
	GetUseType() *string
}

type ComponentTemplateCreateCmd struct {
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
	IsCustom    *bool   `json:"isCustom,omitempty" xml:"isCustom,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// State.Redis
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

func (s ComponentTemplateCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s ComponentTemplateCreateCmd) GoString() string {
	return s.String()
}

func (s *ComponentTemplateCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ComponentTemplateCreateCmd) GetConfiguration() []*ConfigModel {
	return s.Configuration
}

func (s *ComponentTemplateCreateCmd) GetDescription() *string {
	return s.Description
}

func (s *ComponentTemplateCreateCmd) GetIsCustom() *bool {
	return s.IsCustom
}

func (s *ComponentTemplateCreateCmd) GetName() *string {
	return s.Name
}

func (s *ComponentTemplateCreateCmd) GetProductId() *int64 {
	return s.ProductId
}

func (s *ComponentTemplateCreateCmd) GetType() *string {
	return s.Type
}

func (s *ComponentTemplateCreateCmd) GetUseScope() *string {
	return s.UseScope
}

func (s *ComponentTemplateCreateCmd) GetUseType() *string {
	return s.UseType
}

func (s *ComponentTemplateCreateCmd) SetCompanyId(v int64) *ComponentTemplateCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *ComponentTemplateCreateCmd) SetConfiguration(v []*ConfigModel) *ComponentTemplateCreateCmd {
	s.Configuration = v
	return s
}

func (s *ComponentTemplateCreateCmd) SetDescription(v string) *ComponentTemplateCreateCmd {
	s.Description = &v
	return s
}

func (s *ComponentTemplateCreateCmd) SetIsCustom(v bool) *ComponentTemplateCreateCmd {
	s.IsCustom = &v
	return s
}

func (s *ComponentTemplateCreateCmd) SetName(v string) *ComponentTemplateCreateCmd {
	s.Name = &v
	return s
}

func (s *ComponentTemplateCreateCmd) SetProductId(v int64) *ComponentTemplateCreateCmd {
	s.ProductId = &v
	return s
}

func (s *ComponentTemplateCreateCmd) SetType(v string) *ComponentTemplateCreateCmd {
	s.Type = &v
	return s
}

func (s *ComponentTemplateCreateCmd) SetUseScope(v string) *ComponentTemplateCreateCmd {
	s.UseScope = &v
	return s
}

func (s *ComponentTemplateCreateCmd) SetUseType(v string) *ComponentTemplateCreateCmd {
	s.UseType = &v
	return s
}

func (s *ComponentTemplateCreateCmd) Validate() error {
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
