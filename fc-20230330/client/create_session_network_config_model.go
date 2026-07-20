// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionNetworkConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAllowOut(v []*string) *CreateSessionNetworkConfig
	GetAllowOut() []*string
	SetAllowPublicTraffic(v bool) *CreateSessionNetworkConfig
	GetAllowPublicTraffic() *bool
	SetDenyOut(v []*string) *CreateSessionNetworkConfig
	GetDenyOut() []*string
	SetMaskRequestHost(v string) *CreateSessionNetworkConfig
	GetMaskRequestHost() *string
}

type CreateSessionNetworkConfig struct {
	AllowOut           []*string `json:"allowOut" xml:"allowOut" type:"Repeated"`
	AllowPublicTraffic *bool     `json:"allowPublicTraffic,omitempty" xml:"allowPublicTraffic,omitempty"`
	DenyOut            []*string `json:"denyOut" xml:"denyOut" type:"Repeated"`
	MaskRequestHost    *string   `json:"maskRequestHost,omitempty" xml:"maskRequestHost,omitempty"`
}

func (s CreateSessionNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionNetworkConfig) GoString() string {
	return s.String()
}

func (s *CreateSessionNetworkConfig) GetAllowOut() []*string {
	return s.AllowOut
}

func (s *CreateSessionNetworkConfig) GetAllowPublicTraffic() *bool {
	return s.AllowPublicTraffic
}

func (s *CreateSessionNetworkConfig) GetDenyOut() []*string {
	return s.DenyOut
}

func (s *CreateSessionNetworkConfig) GetMaskRequestHost() *string {
	return s.MaskRequestHost
}

func (s *CreateSessionNetworkConfig) SetAllowOut(v []*string) *CreateSessionNetworkConfig {
	s.AllowOut = v
	return s
}

func (s *CreateSessionNetworkConfig) SetAllowPublicTraffic(v bool) *CreateSessionNetworkConfig {
	s.AllowPublicTraffic = &v
	return s
}

func (s *CreateSessionNetworkConfig) SetDenyOut(v []*string) *CreateSessionNetworkConfig {
	s.DenyOut = v
	return s
}

func (s *CreateSessionNetworkConfig) SetMaskRequestHost(v string) *CreateSessionNetworkConfig {
	s.MaskRequestHost = &v
	return s
}

func (s *CreateSessionNetworkConfig) Validate() error {
	return dara.Validate(s)
}
