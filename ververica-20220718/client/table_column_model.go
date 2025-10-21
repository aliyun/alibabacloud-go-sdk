// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTableColumn interface {
	dara.Model
	String() string
	GoString() string
	SetExpression(v string) *TableColumn
	GetExpression() *string
	SetLogicalType(v string) *TableColumn
	GetLogicalType() *string
	SetMetadataInfo(v *MetadataInfo) *TableColumn
	GetMetadataInfo() *MetadataInfo
	SetName(v string) *TableColumn
	GetName() *string
	SetNullable(v bool) *TableColumn
	GetNullable() *bool
	SetType(v string) *TableColumn
	GetType() *string
}

type TableColumn struct {
	Expression   *string       `json:"expression,omitempty" xml:"expression,omitempty"`
	LogicalType  *string       `json:"logicalType,omitempty" xml:"logicalType,omitempty"`
	MetadataInfo *MetadataInfo `json:"metadataInfo,omitempty" xml:"metadataInfo,omitempty"`
	// This parameter is required.
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Nullable *bool   `json:"nullable,omitempty" xml:"nullable,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TableColumn) String() string {
	return dara.Prettify(s)
}

func (s TableColumn) GoString() string {
	return s.String()
}

func (s *TableColumn) GetExpression() *string {
	return s.Expression
}

func (s *TableColumn) GetLogicalType() *string {
	return s.LogicalType
}

func (s *TableColumn) GetMetadataInfo() *MetadataInfo {
	return s.MetadataInfo
}

func (s *TableColumn) GetName() *string {
	return s.Name
}

func (s *TableColumn) GetNullable() *bool {
	return s.Nullable
}

func (s *TableColumn) GetType() *string {
	return s.Type
}

func (s *TableColumn) SetExpression(v string) *TableColumn {
	s.Expression = &v
	return s
}

func (s *TableColumn) SetLogicalType(v string) *TableColumn {
	s.LogicalType = &v
	return s
}

func (s *TableColumn) SetMetadataInfo(v *MetadataInfo) *TableColumn {
	s.MetadataInfo = v
	return s
}

func (s *TableColumn) SetName(v string) *TableColumn {
	s.Name = &v
	return s
}

func (s *TableColumn) SetNullable(v bool) *TableColumn {
	s.Nullable = &v
	return s
}

func (s *TableColumn) SetType(v string) *TableColumn {
	s.Type = &v
	return s
}

func (s *TableColumn) Validate() error {
	if s.MetadataInfo != nil {
		if err := s.MetadataInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
