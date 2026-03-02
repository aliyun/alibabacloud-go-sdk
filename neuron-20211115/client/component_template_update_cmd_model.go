// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iComponentTemplateUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetConfiguration(v []*ConfigModel) *ComponentTemplateUpdateCmd
	GetConfiguration() []*ConfigModel
	SetDescription(v string) *ComponentTemplateUpdateCmd
	GetDescription() *string
	SetId(v int64) *ComponentTemplateUpdateCmd
	GetId() *int64
	SetUseScope(v string) *ComponentTemplateUpdateCmd
	GetUseScope() *string
	SetUseType(v string) *ComponentTemplateUpdateCmd
	GetUseType() *string
}

type ComponentTemplateUpdateCmd struct {
	Configuration []*ConfigModel `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Repeated"`
	// example:
	//
	// 用于统一管理资源
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// Inner
	UseScope *string `json:"useScope,omitempty" xml:"useScope,omitempty"`
	// example:
	//
	// Share
	UseType *string `json:"useType,omitempty" xml:"useType,omitempty"`
}

func (s ComponentTemplateUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s ComponentTemplateUpdateCmd) GoString() string {
	return s.String()
}

func (s *ComponentTemplateUpdateCmd) GetConfiguration() []*ConfigModel {
	return s.Configuration
}

func (s *ComponentTemplateUpdateCmd) GetDescription() *string {
	return s.Description
}

func (s *ComponentTemplateUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *ComponentTemplateUpdateCmd) GetUseScope() *string {
	return s.UseScope
}

func (s *ComponentTemplateUpdateCmd) GetUseType() *string {
	return s.UseType
}

func (s *ComponentTemplateUpdateCmd) SetConfiguration(v []*ConfigModel) *ComponentTemplateUpdateCmd {
	s.Configuration = v
	return s
}

func (s *ComponentTemplateUpdateCmd) SetDescription(v string) *ComponentTemplateUpdateCmd {
	s.Description = &v
	return s
}

func (s *ComponentTemplateUpdateCmd) SetId(v int64) *ComponentTemplateUpdateCmd {
	s.Id = &v
	return s
}

func (s *ComponentTemplateUpdateCmd) SetUseScope(v string) *ComponentTemplateUpdateCmd {
	s.UseScope = &v
	return s
}

func (s *ComponentTemplateUpdateCmd) SetUseType(v string) *ComponentTemplateUpdateCmd {
	s.UseType = &v
	return s
}

func (s *ComponentTemplateUpdateCmd) Validate() error {
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
