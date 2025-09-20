// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElement interface {
  dara.Model
  String() string
  GoString() string
  SetElementContents(v []*ElementContent) *Element
  GetElementContents() []*ElementContent 
  SetElementRelations(v []*ElementRelation) *Element
  GetElementRelations() []*ElementRelation 
  SetElementType(v string) *Element
  GetElementType() *string 
  SetObjectId(v string) *Element
  GetObjectId() *string 
  SetSemanticSimilarity(v float32) *Element
  GetSemanticSimilarity() *float32 
}

type Element struct {
  ElementContents []*ElementContent `json:"ElementContents,omitempty" xml:"ElementContents,omitempty" type:"Repeated"`
  ElementRelations []*ElementRelation `json:"ElementRelations,omitempty" xml:"ElementRelations,omitempty" type:"Repeated"`
  ElementType *string `json:"ElementType,omitempty" xml:"ElementType,omitempty"`
  ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
  SemanticSimilarity *float32 `json:"SemanticSimilarity,omitempty" xml:"SemanticSimilarity,omitempty"`
}

func (s Element) String() string {
  return dara.Prettify(s)
}

func (s Element) GoString() string {
  return s.String()
}

func (s *Element) GetElementContents() []*ElementContent  {
  return s.ElementContents
}

func (s *Element) GetElementRelations() []*ElementRelation  {
  return s.ElementRelations
}

func (s *Element) GetElementType() *string  {
  return s.ElementType
}

func (s *Element) GetObjectId() *string  {
  return s.ObjectId
}

func (s *Element) GetSemanticSimilarity() *float32  {
  return s.SemanticSimilarity
}

func (s *Element) SetElementContents(v []*ElementContent) *Element {
  s.ElementContents = v
  return s
}

func (s *Element) SetElementRelations(v []*ElementRelation) *Element {
  s.ElementRelations = v
  return s
}

func (s *Element) SetElementType(v string) *Element {
  s.ElementType = &v
  return s
}

func (s *Element) SetObjectId(v string) *Element {
  s.ObjectId = &v
  return s
}

func (s *Element) SetSemanticSimilarity(v float32) *Element {
  s.SemanticSimilarity = &v
  return s
}

func (s *Element) Validate() error {
  return dara.Validate(s)
}

