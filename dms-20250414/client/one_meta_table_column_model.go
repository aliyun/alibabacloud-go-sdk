// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOneMetaTableColumn interface {
	dara.Model
	String() string
	GoString() string
	SetColumnName(v string) *OneMetaTableColumn
	GetColumnName() *string
	SetColumnType(v string) *OneMetaTableColumn
	GetColumnType() *string
	SetDescription(v string) *OneMetaTableColumn
	GetDescription() *string
	SetEngineMeta(v *OneMetaTableColumnEngineMeta) *OneMetaTableColumn
	GetEngineMeta() *OneMetaTableColumnEngineMeta
	SetPosition(v int32) *OneMetaTableColumn
	GetPosition() *int32
}

type OneMetaTableColumn struct {
	ColumnName  *string                       `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	ColumnType  *string                       `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	Description *string                       `json:"Description,omitempty" xml:"Description,omitempty"`
	EngineMeta  *OneMetaTableColumnEngineMeta `json:"EngineMeta,omitempty" xml:"EngineMeta,omitempty"`
	Position    *int32                        `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s OneMetaTableColumn) String() string {
	return dara.Prettify(s)
}

func (s OneMetaTableColumn) GoString() string {
	return s.String()
}

func (s *OneMetaTableColumn) GetColumnName() *string {
	return s.ColumnName
}

func (s *OneMetaTableColumn) GetColumnType() *string {
	return s.ColumnType
}

func (s *OneMetaTableColumn) GetDescription() *string {
	return s.Description
}

func (s *OneMetaTableColumn) GetEngineMeta() *OneMetaTableColumnEngineMeta {
	return s.EngineMeta
}

func (s *OneMetaTableColumn) GetPosition() *int32 {
	return s.Position
}

func (s *OneMetaTableColumn) SetColumnName(v string) *OneMetaTableColumn {
	s.ColumnName = &v
	return s
}

func (s *OneMetaTableColumn) SetColumnType(v string) *OneMetaTableColumn {
	s.ColumnType = &v
	return s
}

func (s *OneMetaTableColumn) SetDescription(v string) *OneMetaTableColumn {
	s.Description = &v
	return s
}

func (s *OneMetaTableColumn) SetEngineMeta(v *OneMetaTableColumnEngineMeta) *OneMetaTableColumn {
	s.EngineMeta = v
	return s
}

func (s *OneMetaTableColumn) SetPosition(v int32) *OneMetaTableColumn {
	s.Position = &v
	return s
}

func (s *OneMetaTableColumn) Validate() error {
	if s.EngineMeta != nil {
		if err := s.EngineMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}
