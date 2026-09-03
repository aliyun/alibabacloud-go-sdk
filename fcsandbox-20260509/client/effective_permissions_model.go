// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEffectivePermissions interface {
  dara.Model
  String() string
  GoString() string
  SetActions(v []*string) *EffectivePermissions
  GetActions() []*string 
  SetCapabilities(v []*string) *EffectivePermissions
  GetCapabilities() []*string 
}

type EffectivePermissions struct {
  // The actions.
  Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
  // The capabilities.
  Capabilities []*string `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Repeated"`
}

func (s EffectivePermissions) String() string {
  return dara.Prettify(s)
}

func (s EffectivePermissions) GoString() string {
  return s.String()
}

func (s *EffectivePermissions) GetActions() []*string  {
  return s.Actions
}

func (s *EffectivePermissions) GetCapabilities() []*string  {
  return s.Capabilities
}

func (s *EffectivePermissions) SetActions(v []*string) *EffectivePermissions {
  s.Actions = v
  return s
}

func (s *EffectivePermissions) SetCapabilities(v []*string) *EffectivePermissions {
  s.Capabilities = v
  return s
}

func (s *EffectivePermissions) Validate() error {
  return dara.Validate(s)
}

