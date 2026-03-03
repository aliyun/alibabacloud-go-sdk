// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRow interface {
	dara.Model
	String() string
	GoString() string
	SetColumns(v []*Column) *Row
	GetColumns() []*Column
}

type Row struct {
	Columns []*Column `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
}

func (s Row) String() string {
	return dara.Prettify(s)
}

func (s Row) GoString() string {
	return s.String()
}

func (s *Row) GetColumns() []*Column {
	return s.Columns
}

func (s *Row) SetColumns(v []*Column) *Row {
	s.Columns = v
	return s
}

func (s *Row) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
