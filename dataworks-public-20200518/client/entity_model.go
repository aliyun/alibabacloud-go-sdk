// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntity interface {
  dara.Model
  String() string
  GoString() string
  SetEntityContent(v map[string]interface{}) *Entity
  GetEntityContent() map[string]interface{} 
  SetQualifiedName(v string) *Entity
  GetQualifiedName() *string 
  SetTenantId(v int64) *Entity
  GetTenantId() *int64 
}

type Entity struct {
  EntityContent map[string]interface{} `json:"EntityContent,omitempty" xml:"EntityContent,omitempty"`
  // example:
  // 
  // maxcompute_table.563f0357118d05ef145d6bddf2966cc23e86ca8f2f013f915e565afdf09f7a23
  QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
  // example:
  // 
  // 12345
  TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s Entity) String() string {
  return dara.Prettify(s)
}

func (s Entity) GoString() string {
  return s.String()
}

func (s *Entity) GetEntityContent() map[string]interface{}  {
  return s.EntityContent
}

func (s *Entity) GetQualifiedName() *string  {
  return s.QualifiedName
}

func (s *Entity) GetTenantId() *int64  {
  return s.TenantId
}

func (s *Entity) SetEntityContent(v map[string]interface{}) *Entity {
  s.EntityContent = v
  return s
}

func (s *Entity) SetQualifiedName(v string) *Entity {
  s.QualifiedName = &v
  return s
}

func (s *Entity) SetTenantId(v int64) *Entity {
  s.TenantId = &v
  return s
}

func (s *Entity) Validate() error {
  return dara.Validate(s)
}

