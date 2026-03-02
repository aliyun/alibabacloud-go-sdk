// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iComponentUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetConfiguration(v []*ConfigModel) *ComponentUpdateCmd
	GetConfiguration() []*ConfigModel
	SetDescription(v string) *ComponentUpdateCmd
	GetDescription() *string
	SetId(v int64) *ComponentUpdateCmd
	GetId() *int64
	SetScope(v string) *ComponentUpdateCmd
	GetScope() *string
	SetUseScope(v string) *ComponentUpdateCmd
	GetUseScope() *string
	SetUseType(v string) *ComponentUpdateCmd
	GetUseType() *string
}

type ComponentUpdateCmd struct {
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
	// 1,2
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// example:
	//
	// Inner
	UseScope *string `json:"useScope,omitempty" xml:"useScope,omitempty"`
	// example:
	//
	// Share
	UseType *string `json:"useType,omitempty" xml:"useType,omitempty"`
}

func (s ComponentUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s ComponentUpdateCmd) GoString() string {
	return s.String()
}

func (s *ComponentUpdateCmd) GetConfiguration() []*ConfigModel {
	return s.Configuration
}

func (s *ComponentUpdateCmd) GetDescription() *string {
	return s.Description
}

func (s *ComponentUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *ComponentUpdateCmd) GetScope() *string {
	return s.Scope
}

func (s *ComponentUpdateCmd) GetUseScope() *string {
	return s.UseScope
}

func (s *ComponentUpdateCmd) GetUseType() *string {
	return s.UseType
}

func (s *ComponentUpdateCmd) SetConfiguration(v []*ConfigModel) *ComponentUpdateCmd {
	s.Configuration = v
	return s
}

func (s *ComponentUpdateCmd) SetDescription(v string) *ComponentUpdateCmd {
	s.Description = &v
	return s
}

func (s *ComponentUpdateCmd) SetId(v int64) *ComponentUpdateCmd {
	s.Id = &v
	return s
}

func (s *ComponentUpdateCmd) SetScope(v string) *ComponentUpdateCmd {
	s.Scope = &v
	return s
}

func (s *ComponentUpdateCmd) SetUseScope(v string) *ComponentUpdateCmd {
	s.UseScope = &v
	return s
}

func (s *ComponentUpdateCmd) SetUseType(v string) *ComponentUpdateCmd {
	s.UseType = &v
	return s
}

func (s *ComponentUpdateCmd) Validate() error {
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
