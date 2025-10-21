// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRow interface {
	dara.Model
	String() string
	GoString() string
	SetCells(v []*Cell) *Row
	GetCells() []*Cell
}

type Row struct {
	Cells []*Cell `json:"cells,omitempty" xml:"cells,omitempty" type:"Repeated"`
}

func (s Row) String() string {
	return dara.Prettify(s)
}

func (s Row) GoString() string {
	return s.String()
}

func (s *Row) GetCells() []*Cell {
	return s.Cells
}

func (s *Row) SetCells(v []*Cell) *Row {
	s.Cells = v
	return s
}

func (s *Row) Validate() error {
	if s.Cells != nil {
		for _, item := range s.Cells {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
