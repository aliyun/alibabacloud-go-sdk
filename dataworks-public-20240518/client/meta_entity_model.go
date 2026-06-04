// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaEntity interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v map[string]*string) *MetaEntity
	GetAttributes() map[string]*string
	SetComment(v string) *MetaEntity
	GetComment() *string
	SetCreateTime(v int64) *MetaEntity
	GetCreateTime() *int64
	SetCustomAttributes(v map[string][]*string) *MetaEntity
	GetCustomAttributes() map[string][]*string
	SetEntityType(v string) *MetaEntity
	GetEntityType() *string
	SetId(v string) *MetaEntity
	GetId() *string
	SetMetaEntityDef(v *MetaEntityDef) *MetaEntity
	GetMetaEntityDef() *MetaEntityDef
	SetModifyTime(v int64) *MetaEntity
	GetModifyTime() *int64
	SetName(v string) *MetaEntity
	GetName() *string
	SetOwnerId(v string) *MetaEntity
	GetOwnerId() *string
}

type MetaEntity struct {
	Attributes       map[string]*string   `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	Comment          *string              `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreateTime       *int64               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomAttributes map[string][]*string `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty"`
	// example:
	//
	// custom_entity-customer_api
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// custom_entity-customer_api:api_001
	Id            *string        `json:"Id,omitempty" xml:"Id,omitempty"`
	MetaEntityDef *MetaEntityDef `json:"MetaEntityDef,omitempty" xml:"MetaEntityDef,omitempty"`
	ModifyTime    *int64         `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// api_001
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s MetaEntity) String() string {
	return dara.Prettify(s)
}

func (s MetaEntity) GoString() string {
	return s.String()
}

func (s *MetaEntity) GetAttributes() map[string]*string {
	return s.Attributes
}

func (s *MetaEntity) GetComment() *string {
	return s.Comment
}

func (s *MetaEntity) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *MetaEntity) GetCustomAttributes() map[string][]*string {
	return s.CustomAttributes
}

func (s *MetaEntity) GetEntityType() *string {
	return s.EntityType
}

func (s *MetaEntity) GetId() *string {
	return s.Id
}

func (s *MetaEntity) GetMetaEntityDef() *MetaEntityDef {
	return s.MetaEntityDef
}

func (s *MetaEntity) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *MetaEntity) GetName() *string {
	return s.Name
}

func (s *MetaEntity) GetOwnerId() *string {
	return s.OwnerId
}

func (s *MetaEntity) SetAttributes(v map[string]*string) *MetaEntity {
	s.Attributes = v
	return s
}

func (s *MetaEntity) SetComment(v string) *MetaEntity {
	s.Comment = &v
	return s
}

func (s *MetaEntity) SetCreateTime(v int64) *MetaEntity {
	s.CreateTime = &v
	return s
}

func (s *MetaEntity) SetCustomAttributes(v map[string][]*string) *MetaEntity {
	s.CustomAttributes = v
	return s
}

func (s *MetaEntity) SetEntityType(v string) *MetaEntity {
	s.EntityType = &v
	return s
}

func (s *MetaEntity) SetId(v string) *MetaEntity {
	s.Id = &v
	return s
}

func (s *MetaEntity) SetMetaEntityDef(v *MetaEntityDef) *MetaEntity {
	s.MetaEntityDef = v
	return s
}

func (s *MetaEntity) SetModifyTime(v int64) *MetaEntity {
	s.ModifyTime = &v
	return s
}

func (s *MetaEntity) SetName(v string) *MetaEntity {
	s.Name = &v
	return s
}

func (s *MetaEntity) SetOwnerId(v string) *MetaEntity {
	s.OwnerId = &v
	return s
}

func (s *MetaEntity) Validate() error {
	if s.MetaEntityDef != nil {
		if err := s.MetaEntityDef.Validate(); err != nil {
			return err
		}
	}
	return nil
}
