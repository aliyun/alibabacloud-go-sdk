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
  // The allowlist for outbound access from the sandbox. After configuration, the sandbox can access only the destinations in the list.
  AllowOut []*string `json:"allowOut" xml:"allowOut" type:"Repeated"`
  // Specifies whether to allow Direct Internet Access to the sandbox.
  AllowPublicTraffic *bool `json:"allowPublicTraffic,omitempty" xml:"allowPublicTraffic,omitempty"`
  // The denylist for outbound access from the sandbox. After configuration, the sandbox cannot access the destinations in the list.
  DenyOut []*string `json:"denyOut" xml:"denyOut" type:"Repeated"`
  // The domain name used to rewrite the request Host header when accessing the sandbox. If this parameter is not specified, the original Host is retained.
  // 
  // example:
  // 
  // sandbox.example.com
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

