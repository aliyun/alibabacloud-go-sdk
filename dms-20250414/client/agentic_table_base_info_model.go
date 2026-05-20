// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticTableBaseInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogType(v string) *AgenticTableBaseInfo
	GetCatalogType() *string
	SetDatabaseUuid(v string) *AgenticTableBaseInfo
	GetDatabaseUuid() *string
	SetDescription(v string) *AgenticTableBaseInfo
	GetDescription() *string
	SetEngineMeta(v *AgenticTableEngineMeta) *AgenticTableBaseInfo
	GetEngineMeta() *AgenticTableEngineMeta
	SetName(v string) *AgenticTableBaseInfo
	GetName() *string
	SetQualifiedName(v string) *AgenticTableBaseInfo
	GetQualifiedName() *string
	SetTableType(v string) *AgenticTableBaseInfo
	GetTableType() *string
	SetTableUuid(v string) *AgenticTableBaseInfo
	GetTableUuid() *string
}

type AgenticTableBaseInfo struct {
	CatalogType   *string                 `json:"CatalogType,omitempty" xml:"CatalogType,omitempty"`
	DatabaseUuid  *string                 `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	Description   *string                 `json:"Description,omitempty" xml:"Description,omitempty"`
	EngineMeta    *AgenticTableEngineMeta `json:"EngineMeta,omitempty" xml:"EngineMeta,omitempty"`
	Name          *string                 `json:"Name,omitempty" xml:"Name,omitempty"`
	QualifiedName *string                 `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
	TableType     *string                 `json:"TableType,omitempty" xml:"TableType,omitempty"`
	TableUuid     *string                 `json:"TableUuid,omitempty" xml:"TableUuid,omitempty"`
}

func (s AgenticTableBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s AgenticTableBaseInfo) GoString() string {
	return s.String()
}

func (s *AgenticTableBaseInfo) GetCatalogType() *string {
	return s.CatalogType
}

func (s *AgenticTableBaseInfo) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *AgenticTableBaseInfo) GetDescription() *string {
	return s.Description
}

func (s *AgenticTableBaseInfo) GetEngineMeta() *AgenticTableEngineMeta {
	return s.EngineMeta
}

func (s *AgenticTableBaseInfo) GetName() *string {
	return s.Name
}

func (s *AgenticTableBaseInfo) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *AgenticTableBaseInfo) GetTableType() *string {
	return s.TableType
}

func (s *AgenticTableBaseInfo) GetTableUuid() *string {
	return s.TableUuid
}

func (s *AgenticTableBaseInfo) SetCatalogType(v string) *AgenticTableBaseInfo {
	s.CatalogType = &v
	return s
}

func (s *AgenticTableBaseInfo) SetDatabaseUuid(v string) *AgenticTableBaseInfo {
	s.DatabaseUuid = &v
	return s
}

func (s *AgenticTableBaseInfo) SetDescription(v string) *AgenticTableBaseInfo {
	s.Description = &v
	return s
}

func (s *AgenticTableBaseInfo) SetEngineMeta(v *AgenticTableEngineMeta) *AgenticTableBaseInfo {
	s.EngineMeta = v
	return s
}

func (s *AgenticTableBaseInfo) SetName(v string) *AgenticTableBaseInfo {
	s.Name = &v
	return s
}

func (s *AgenticTableBaseInfo) SetQualifiedName(v string) *AgenticTableBaseInfo {
	s.QualifiedName = &v
	return s
}

func (s *AgenticTableBaseInfo) SetTableType(v string) *AgenticTableBaseInfo {
	s.TableType = &v
	return s
}

func (s *AgenticTableBaseInfo) SetTableUuid(v string) *AgenticTableBaseInfo {
	s.TableUuid = &v
	return s
}

func (s *AgenticTableBaseInfo) Validate() error {
	if s.EngineMeta != nil {
		if err := s.EngineMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}
