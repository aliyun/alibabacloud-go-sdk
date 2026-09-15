// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iE2BNetwork interface {
  dara.Model
  String() string
  GoString() string
  SetAllowOut(v []*string) *E2BNetwork
  GetAllowOut() []*string 
  SetAllowPublicTraffic(v bool) *E2BNetwork
  GetAllowPublicTraffic() *bool 
  SetDenyOut(v []*string) *E2BNetwork
  GetDenyOut() []*string 
  SetMaskRequestHost(v string) *E2BNetwork
  GetMaskRequestHost() *string 
}

type E2BNetwork struct {
  AllowOut []*string `json:"allowOut,omitempty" xml:"allowOut,omitempty" type:"Repeated"`
  AllowPublicTraffic *bool `json:"allowPublicTraffic,omitempty" xml:"allowPublicTraffic,omitempty"`
  DenyOut []*string `json:"denyOut,omitempty" xml:"denyOut,omitempty" type:"Repeated"`
  MaskRequestHost *string `json:"maskRequestHost,omitempty" xml:"maskRequestHost,omitempty"`
}

func (s E2BNetwork) String() string {
  return dara.Prettify(s)
}

func (s E2BNetwork) GoString() string {
  return s.String()
}

func (s *E2BNetwork) GetAllowOut() []*string  {
  return s.AllowOut
}

func (s *E2BNetwork) GetAllowPublicTraffic() *bool  {
  return s.AllowPublicTraffic
}

func (s *E2BNetwork) GetDenyOut() []*string  {
  return s.DenyOut
}

func (s *E2BNetwork) GetMaskRequestHost() *string  {
  return s.MaskRequestHost
}

func (s *E2BNetwork) SetAllowOut(v []*string) *E2BNetwork {
  s.AllowOut = v
  return s
}

func (s *E2BNetwork) SetAllowPublicTraffic(v bool) *E2BNetwork {
  s.AllowPublicTraffic = &v
  return s
}

func (s *E2BNetwork) SetDenyOut(v []*string) *E2BNetwork {
  s.DenyOut = v
  return s
}

func (s *E2BNetwork) SetMaskRequestHost(v string) *E2BNetwork {
  s.MaskRequestHost = &v
  return s
}

func (s *E2BNetwork) Validate() error {
  return dara.Validate(s)
}

