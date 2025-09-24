// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLineageEntity interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v map[string]interface{}) *LineageEntity
	GetAttributes() map[string]interface{}
	SetEntityType(v string) *LineageEntity
	GetEntityType() *string
	SetName(v string) *LineageEntity
	GetName() *string
	SetQualifiedName(v string) *LineageEntity
	GetQualifiedName() *string
}

type LineageEntity struct {
	Attributes    map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	EntityType    *string                `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Name          *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	QualifiedName *string                `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
}

func (s LineageEntity) String() string {
	return dara.Prettify(s)
}

func (s LineageEntity) GoString() string {
	return s.String()
}

func (s *LineageEntity) GetAttributes() map[string]interface{} {
	return s.Attributes
}

func (s *LineageEntity) GetEntityType() *string {
	return s.EntityType
}

func (s *LineageEntity) GetName() *string {
	return s.Name
}

func (s *LineageEntity) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *LineageEntity) SetAttributes(v map[string]interface{}) *LineageEntity {
	s.Attributes = v
	return s
}

func (s *LineageEntity) SetEntityType(v string) *LineageEntity {
	s.EntityType = &v
	return s
}

func (s *LineageEntity) SetName(v string) *LineageEntity {
	s.Name = &v
	return s
}

func (s *LineageEntity) SetQualifiedName(v string) *LineageEntity {
	s.QualifiedName = &v
	return s
}

func (s *LineageEntity) Validate() error {
	return dara.Validate(s)
}
