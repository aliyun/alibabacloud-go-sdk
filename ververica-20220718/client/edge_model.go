// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEdge interface {
  dara.Model
  String() string
  GoString() string
  SetColumnLineage(v []*Relation) *Edge
  GetColumnLineage() []*Relation 
  SetTableLineage(v []*Relation) *Edge
  GetTableLineage() []*Relation 
}

type Edge struct {
  ColumnLineage []*Relation `json:"columnLineage,omitempty" xml:"columnLineage,omitempty" type:"Repeated"`
  TableLineage []*Relation `json:"tableLineage,omitempty" xml:"tableLineage,omitempty" type:"Repeated"`
}

func (s Edge) String() string {
  return dara.Prettify(s)
}

func (s Edge) GoString() string {
  return s.String()
}

func (s *Edge) GetColumnLineage() []*Relation  {
  return s.ColumnLineage
}

func (s *Edge) GetTableLineage() []*Relation  {
  return s.TableLineage
}

func (s *Edge) SetColumnLineage(v []*Relation) *Edge {
  s.ColumnLineage = v
  return s
}

func (s *Edge) SetTableLineage(v []*Relation) *Edge {
  s.TableLineage = v
  return s
}

func (s *Edge) Validate() error {
  if s.ColumnLineage != nil {
    for _, item := range s.ColumnLineage {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.TableLineage != nil {
    for _, item := range s.TableLineage {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

