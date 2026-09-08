// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLFlowSankey interface {
	dara.Model
	String() string
	GoString() string
	SetColumns(v []*RLFlowSankeyColumn) *RLFlowSankey
	GetColumns() []*RLFlowSankeyColumn
	SetExits(v []*RLFlowSankeyExit) *RLFlowSankey
	GetExits() []*RLFlowSankeyExit
}

type RLFlowSankey struct {
	// The five columns of the main chain.
	//
	// example:
	//
	// [{"Key":"traj","Label":"Generated trajectory","Count":96}]
	Columns []*RLFlowSankeyColumn `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The outflow edges of each column.
	//
	// example:
	//
	// [{"From":"run","FromIdx":1,"Label":"Not started","Count":0}]
	Exits []*RLFlowSankeyExit `json:"Exits,omitempty" xml:"Exits,omitempty" type:"Repeated"`
}

func (s RLFlowSankey) String() string {
	return dara.Prettify(s)
}

func (s RLFlowSankey) GoString() string {
	return s.String()
}

func (s *RLFlowSankey) GetColumns() []*RLFlowSankeyColumn {
	return s.Columns
}

func (s *RLFlowSankey) GetExits() []*RLFlowSankeyExit {
	return s.Exits
}

func (s *RLFlowSankey) SetColumns(v []*RLFlowSankeyColumn) *RLFlowSankey {
	s.Columns = v
	return s
}

func (s *RLFlowSankey) SetExits(v []*RLFlowSankeyExit) *RLFlowSankey {
	s.Exits = v
	return s
}

func (s *RLFlowSankey) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Exits != nil {
		for _, item := range s.Exits {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
