// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRowUpdate interface {
	dara.Model
	String() string
	GoString() string
	SetRow(v *Row) *RowUpdate
	GetRow() *Row
	SetRowKind(v string) *RowUpdate
	GetRowKind() *string
}

type RowUpdate struct {
	Row     *Row    `json:"row,omitempty" xml:"row,omitempty"`
	RowKind *string `json:"rowKind,omitempty" xml:"rowKind,omitempty"`
}

func (s RowUpdate) String() string {
	return dara.Prettify(s)
}

func (s RowUpdate) GoString() string {
	return s.String()
}

func (s *RowUpdate) GetRow() *Row {
	return s.Row
}

func (s *RowUpdate) GetRowKind() *string {
	return s.RowKind
}

func (s *RowUpdate) SetRow(v *Row) *RowUpdate {
	s.Row = v
	return s
}

func (s *RowUpdate) SetRowKind(v string) *RowUpdate {
	s.RowKind = &v
	return s
}

func (s *RowUpdate) Validate() error {
	if s.Row != nil {
		if err := s.Row.Validate(); err != nil {
			return err
		}
	}
	return nil
}
