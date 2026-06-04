// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaEntityDef interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeDefs(v []*MetaEntityAttributeDef) *MetaEntityDef
	GetAttributeDefs() []*MetaEntityAttributeDef
	SetCreateTime(v int64) *MetaEntityDef
	GetCreateTime() *int64
	SetDescription(v string) *MetaEntityDef
	GetDescription() *string
	SetDisplayName(v string) *MetaEntityDef
	GetDisplayName() *string
	SetEntityType(v string) *MetaEntityDef
	GetEntityType() *string
	SetExtend(v string) *MetaEntityDef
	GetExtend() *string
	SetModifyTime(v int64) *MetaEntityDef
	GetModifyTime() *int64
	SetName(v string) *MetaEntityDef
	GetName() *string
}

type MetaEntityDef struct {
	AttributeDefs []*MetaEntityAttributeDef `json:"AttributeDefs,omitempty" xml:"AttributeDefs,omitempty" type:"Repeated"`
	CreateTime    *int64                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string                   `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName   *string                   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// custom_entity-customer_api
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// NONE
	Extend     *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	ModifyTime *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// customer_api
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s MetaEntityDef) String() string {
	return dara.Prettify(s)
}

func (s MetaEntityDef) GoString() string {
	return s.String()
}

func (s *MetaEntityDef) GetAttributeDefs() []*MetaEntityAttributeDef {
	return s.AttributeDefs
}

func (s *MetaEntityDef) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *MetaEntityDef) GetDescription() *string {
	return s.Description
}

func (s *MetaEntityDef) GetDisplayName() *string {
	return s.DisplayName
}

func (s *MetaEntityDef) GetEntityType() *string {
	return s.EntityType
}

func (s *MetaEntityDef) GetExtend() *string {
	return s.Extend
}

func (s *MetaEntityDef) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *MetaEntityDef) GetName() *string {
	return s.Name
}

func (s *MetaEntityDef) SetAttributeDefs(v []*MetaEntityAttributeDef) *MetaEntityDef {
	s.AttributeDefs = v
	return s
}

func (s *MetaEntityDef) SetCreateTime(v int64) *MetaEntityDef {
	s.CreateTime = &v
	return s
}

func (s *MetaEntityDef) SetDescription(v string) *MetaEntityDef {
	s.Description = &v
	return s
}

func (s *MetaEntityDef) SetDisplayName(v string) *MetaEntityDef {
	s.DisplayName = &v
	return s
}

func (s *MetaEntityDef) SetEntityType(v string) *MetaEntityDef {
	s.EntityType = &v
	return s
}

func (s *MetaEntityDef) SetExtend(v string) *MetaEntityDef {
	s.Extend = &v
	return s
}

func (s *MetaEntityDef) SetModifyTime(v int64) *MetaEntityDef {
	s.ModifyTime = &v
	return s
}

func (s *MetaEntityDef) SetName(v string) *MetaEntityDef {
	s.Name = &v
	return s
}

func (s *MetaEntityDef) Validate() error {
	if s.AttributeDefs != nil {
		for _, item := range s.AttributeDefs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
