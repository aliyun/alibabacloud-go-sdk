// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLineageEntityVO interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v map[string]*string) *LineageEntityVO
	GetAttributes() map[string]*string
	SetDetailUrl(v string) *LineageEntityVO
	GetDetailUrl() *string
	SetEntityType(v string) *LineageEntityVO
	GetEntityType() *string
	SetName(v string) *LineageEntityVO
	GetName() *string
	SetOwner(v string) *LineageEntityVO
	GetOwner() *string
	SetParentName(v string) *LineageEntityVO
	GetParentName() *string
	SetQualifiedName(v string) *LineageEntityVO
	GetQualifiedName() *string
}

type LineageEntityVO struct {
	// example:
	//
	// attribute map
	Attributes map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// example:
	//
	// http://domain.test.url/entity
	DetailUrl *string `json:"DetailUrl,omitempty" xml:"DetailUrl,omitempty"`
	// example:
	//
	// maxcompute-table
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// tableName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// owner
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// dbName
	ParentName *string `json:"ParentName,omitempty" xml:"ParentName,omitempty"`
	// example:
	//
	// maxcompute-table.projectName.tablename
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
}

func (s LineageEntityVO) String() string {
	return dara.Prettify(s)
}

func (s LineageEntityVO) GoString() string {
	return s.String()
}

func (s *LineageEntityVO) GetAttributes() map[string]*string {
	return s.Attributes
}

func (s *LineageEntityVO) GetDetailUrl() *string {
	return s.DetailUrl
}

func (s *LineageEntityVO) GetEntityType() *string {
	return s.EntityType
}

func (s *LineageEntityVO) GetName() *string {
	return s.Name
}

func (s *LineageEntityVO) GetOwner() *string {
	return s.Owner
}

func (s *LineageEntityVO) GetParentName() *string {
	return s.ParentName
}

func (s *LineageEntityVO) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *LineageEntityVO) SetAttributes(v map[string]*string) *LineageEntityVO {
	s.Attributes = v
	return s
}

func (s *LineageEntityVO) SetDetailUrl(v string) *LineageEntityVO {
	s.DetailUrl = &v
	return s
}

func (s *LineageEntityVO) SetEntityType(v string) *LineageEntityVO {
	s.EntityType = &v
	return s
}

func (s *LineageEntityVO) SetName(v string) *LineageEntityVO {
	s.Name = &v
	return s
}

func (s *LineageEntityVO) SetOwner(v string) *LineageEntityVO {
	s.Owner = &v
	return s
}

func (s *LineageEntityVO) SetParentName(v string) *LineageEntityVO {
	s.ParentName = &v
	return s
}

func (s *LineageEntityVO) SetQualifiedName(v string) *LineageEntityVO {
	s.QualifiedName = &v
	return s
}

func (s *LineageEntityVO) Validate() error {
	return dara.Validate(s)
}
