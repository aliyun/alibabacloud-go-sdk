// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOneMetaTableBaseInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogType(v string) *OneMetaTableBaseInfo
	GetCatalogType() *string
	SetDatabaseUuid(v string) *OneMetaTableBaseInfo
	GetDatabaseUuid() *string
	SetDescription(v string) *OneMetaTableBaseInfo
	GetDescription() *string
	SetEngineMeta(v *OneMetaTableEngineMeta) *OneMetaTableBaseInfo
	GetEngineMeta() *OneMetaTableEngineMeta
	SetName(v string) *OneMetaTableBaseInfo
	GetName() *string
	SetQualifiedName(v string) *OneMetaTableBaseInfo
	GetQualifiedName() *string
	SetTableType(v string) *OneMetaTableBaseInfo
	GetTableType() *string
	SetTableUuid(v string) *OneMetaTableBaseInfo
	GetTableUuid() *string
}

type OneMetaTableBaseInfo struct {
	CatalogType   *string                 `json:"CatalogType,omitempty" xml:"CatalogType,omitempty"`
	DatabaseUuid  *string                 `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	Description   *string                 `json:"Description,omitempty" xml:"Description,omitempty"`
	EngineMeta    *OneMetaTableEngineMeta `json:"EngineMeta,omitempty" xml:"EngineMeta,omitempty"`
	Name          *string                 `json:"Name,omitempty" xml:"Name,omitempty"`
	QualifiedName *string                 `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
	TableType     *string                 `json:"TableType,omitempty" xml:"TableType,omitempty"`
	TableUuid     *string                 `json:"TableUuid,omitempty" xml:"TableUuid,omitempty"`
}

func (s OneMetaTableBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s OneMetaTableBaseInfo) GoString() string {
	return s.String()
}

func (s *OneMetaTableBaseInfo) GetCatalogType() *string {
	return s.CatalogType
}

func (s *OneMetaTableBaseInfo) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *OneMetaTableBaseInfo) GetDescription() *string {
	return s.Description
}

func (s *OneMetaTableBaseInfo) GetEngineMeta() *OneMetaTableEngineMeta {
	return s.EngineMeta
}

func (s *OneMetaTableBaseInfo) GetName() *string {
	return s.Name
}

func (s *OneMetaTableBaseInfo) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *OneMetaTableBaseInfo) GetTableType() *string {
	return s.TableType
}

func (s *OneMetaTableBaseInfo) GetTableUuid() *string {
	return s.TableUuid
}

func (s *OneMetaTableBaseInfo) SetCatalogType(v string) *OneMetaTableBaseInfo {
	s.CatalogType = &v
	return s
}

func (s *OneMetaTableBaseInfo) SetDatabaseUuid(v string) *OneMetaTableBaseInfo {
	s.DatabaseUuid = &v
	return s
}

func (s *OneMetaTableBaseInfo) SetDescription(v string) *OneMetaTableBaseInfo {
	s.Description = &v
	return s
}

func (s *OneMetaTableBaseInfo) SetEngineMeta(v *OneMetaTableEngineMeta) *OneMetaTableBaseInfo {
	s.EngineMeta = v
	return s
}

func (s *OneMetaTableBaseInfo) SetName(v string) *OneMetaTableBaseInfo {
	s.Name = &v
	return s
}

func (s *OneMetaTableBaseInfo) SetQualifiedName(v string) *OneMetaTableBaseInfo {
	s.QualifiedName = &v
	return s
}

func (s *OneMetaTableBaseInfo) SetTableType(v string) *OneMetaTableBaseInfo {
	s.TableType = &v
	return s
}

func (s *OneMetaTableBaseInfo) SetTableUuid(v string) *OneMetaTableBaseInfo {
	s.TableUuid = &v
	return s
}

func (s *OneMetaTableBaseInfo) Validate() error {
	if s.EngineMeta != nil {
		if err := s.EngineMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}
