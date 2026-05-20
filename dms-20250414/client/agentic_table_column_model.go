// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticTableColumn interface {
	dara.Model
	String() string
	GoString() string
	SetColumnName(v string) *AgenticTableColumn
	GetColumnName() *string
	SetColumnType(v string) *AgenticTableColumn
	GetColumnType() *string
	SetDescription(v string) *AgenticTableColumn
	GetDescription() *string
	SetEngineMeta(v *AgenticTableColumnEngineMeta) *AgenticTableColumn
	GetEngineMeta() *AgenticTableColumnEngineMeta
	SetPosition(v int32) *AgenticTableColumn
	GetPosition() *int32
}

type AgenticTableColumn struct {
	ColumnName  *string                       `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	ColumnType  *string                       `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	Description *string                       `json:"Description,omitempty" xml:"Description,omitempty"`
	EngineMeta  *AgenticTableColumnEngineMeta `json:"EngineMeta,omitempty" xml:"EngineMeta,omitempty"`
	Position    *int32                        `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s AgenticTableColumn) String() string {
	return dara.Prettify(s)
}

func (s AgenticTableColumn) GoString() string {
	return s.String()
}

func (s *AgenticTableColumn) GetColumnName() *string {
	return s.ColumnName
}

func (s *AgenticTableColumn) GetColumnType() *string {
	return s.ColumnType
}

func (s *AgenticTableColumn) GetDescription() *string {
	return s.Description
}

func (s *AgenticTableColumn) GetEngineMeta() *AgenticTableColumnEngineMeta {
	return s.EngineMeta
}

func (s *AgenticTableColumn) GetPosition() *int32 {
	return s.Position
}

func (s *AgenticTableColumn) SetColumnName(v string) *AgenticTableColumn {
	s.ColumnName = &v
	return s
}

func (s *AgenticTableColumn) SetColumnType(v string) *AgenticTableColumn {
	s.ColumnType = &v
	return s
}

func (s *AgenticTableColumn) SetDescription(v string) *AgenticTableColumn {
	s.Description = &v
	return s
}

func (s *AgenticTableColumn) SetEngineMeta(v *AgenticTableColumnEngineMeta) *AgenticTableColumn {
	s.EngineMeta = v
	return s
}

func (s *AgenticTableColumn) SetPosition(v int32) *AgenticTableColumn {
	s.Position = &v
	return s
}

func (s *AgenticTableColumn) Validate() error {
	if s.EngineMeta != nil {
		if err := s.EngineMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}
