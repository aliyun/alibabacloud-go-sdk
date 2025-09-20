// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElementRelation interface {
  dara.Model
  String() string
  GoString() string
  SetObjectId(v string) *ElementRelation
  GetObjectId() *string 
  SetType(v string) *ElementRelation
  GetType() *string 
}

type ElementRelation struct {
  ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
  Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ElementRelation) String() string {
  return dara.Prettify(s)
}

func (s ElementRelation) GoString() string {
  return s.String()
}

func (s *ElementRelation) GetObjectId() *string  {
  return s.ObjectId
}

func (s *ElementRelation) GetType() *string  {
  return s.Type
}

func (s *ElementRelation) SetObjectId(v string) *ElementRelation {
  s.ObjectId = &v
  return s
}

func (s *ElementRelation) SetType(v string) *ElementRelation {
  s.Type = &v
  return s
}

func (s *ElementRelation) Validate() error {
  return dara.Validate(s)
}

