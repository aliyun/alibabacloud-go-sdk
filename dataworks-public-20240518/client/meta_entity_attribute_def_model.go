// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaEntityAttributeDef interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedValues(v []*string) *MetaEntityAttributeDef
	GetAllowedValues() []*string
	SetDescription(v string) *MetaEntityAttributeDef
	GetDescription() *string
	SetDisplayEnabled(v bool) *MetaEntityAttributeDef
	GetDisplayEnabled() *bool
	SetDisplayName(v string) *MetaEntityAttributeDef
	GetDisplayName() *string
	SetIsOptional(v bool) *MetaEntityAttributeDef
	GetIsOptional() *bool
	SetName(v string) *MetaEntityAttributeDef
	GetName() *string
	SetSearchFilterEnabled(v bool) *MetaEntityAttributeDef
	GetSearchFilterEnabled() *bool
	SetType(v string) *MetaEntityAttributeDef
	GetType() *string
}

type MetaEntityAttributeDef struct {
	AllowedValues []*string `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty" type:"Repeated"`
	Description   *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DisplayEnabled *bool `json:"DisplayEnabled,omitempty" xml:"DisplayEnabled,omitempty"`
	// example:
	//
	// API编码
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// true
	IsOptional *bool `json:"IsOptional,omitempty" xml:"IsOptional,omitempty"`
	// example:
	//
	// apiCode
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	SearchFilterEnabled *bool `json:"SearchFilterEnabled,omitempty" xml:"SearchFilterEnabled,omitempty"`
	// example:
	//
	// STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MetaEntityAttributeDef) String() string {
	return dara.Prettify(s)
}

func (s MetaEntityAttributeDef) GoString() string {
	return s.String()
}

func (s *MetaEntityAttributeDef) GetAllowedValues() []*string {
	return s.AllowedValues
}

func (s *MetaEntityAttributeDef) GetDescription() *string {
	return s.Description
}

func (s *MetaEntityAttributeDef) GetDisplayEnabled() *bool {
	return s.DisplayEnabled
}

func (s *MetaEntityAttributeDef) GetDisplayName() *string {
	return s.DisplayName
}

func (s *MetaEntityAttributeDef) GetIsOptional() *bool {
	return s.IsOptional
}

func (s *MetaEntityAttributeDef) GetName() *string {
	return s.Name
}

func (s *MetaEntityAttributeDef) GetSearchFilterEnabled() *bool {
	return s.SearchFilterEnabled
}

func (s *MetaEntityAttributeDef) GetType() *string {
	return s.Type
}

func (s *MetaEntityAttributeDef) SetAllowedValues(v []*string) *MetaEntityAttributeDef {
	s.AllowedValues = v
	return s
}

func (s *MetaEntityAttributeDef) SetDescription(v string) *MetaEntityAttributeDef {
	s.Description = &v
	return s
}

func (s *MetaEntityAttributeDef) SetDisplayEnabled(v bool) *MetaEntityAttributeDef {
	s.DisplayEnabled = &v
	return s
}

func (s *MetaEntityAttributeDef) SetDisplayName(v string) *MetaEntityAttributeDef {
	s.DisplayName = &v
	return s
}

func (s *MetaEntityAttributeDef) SetIsOptional(v bool) *MetaEntityAttributeDef {
	s.IsOptional = &v
	return s
}

func (s *MetaEntityAttributeDef) SetName(v string) *MetaEntityAttributeDef {
	s.Name = &v
	return s
}

func (s *MetaEntityAttributeDef) SetSearchFilterEnabled(v bool) *MetaEntityAttributeDef {
	s.SearchFilterEnabled = &v
	return s
}

func (s *MetaEntityAttributeDef) SetType(v string) *MetaEntityAttributeDef {
	s.Type = &v
	return s
}

func (s *MetaEntityAttributeDef) Validate() error {
	return dara.Validate(s)
}
