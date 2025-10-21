// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditableNamespace interface {
  dara.Model
  String() string
  GoString() string
  SetNamespace(v string) *EditableNamespace
  GetNamespace() *string 
  SetRole(v string) *EditableNamespace
  GetRole() *string 
  SetWorkspaceId(v string) *EditableNamespace
  GetWorkspaceId() *string 
}

type EditableNamespace struct {
  Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
  Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s EditableNamespace) String() string {
  return dara.Prettify(s)
}

func (s EditableNamespace) GoString() string {
  return s.String()
}

func (s *EditableNamespace) GetNamespace() *string  {
  return s.Namespace
}

func (s *EditableNamespace) GetRole() *string  {
  return s.Role
}

func (s *EditableNamespace) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *EditableNamespace) SetNamespace(v string) *EditableNamespace {
  s.Namespace = &v
  return s
}

func (s *EditableNamespace) SetRole(v string) *EditableNamespace {
  s.Role = &v
  return s
}

func (s *EditableNamespace) SetWorkspaceId(v string) *EditableNamespace {
  s.WorkspaceId = &v
  return s
}

func (s *EditableNamespace) Validate() error {
  return dara.Validate(s)
}

