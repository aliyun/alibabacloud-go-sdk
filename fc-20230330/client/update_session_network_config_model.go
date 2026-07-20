// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSessionNetworkConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAllowOut(v []*string) *UpdateSessionNetworkConfig
	GetAllowOut() []*string
	SetDenyOut(v []*string) *UpdateSessionNetworkConfig
	GetDenyOut() []*string
}

type UpdateSessionNetworkConfig struct {
	AllowOut []*string `json:"allowOut" xml:"allowOut" type:"Repeated"`
	DenyOut  []*string `json:"denyOut" xml:"denyOut" type:"Repeated"`
}

func (s UpdateSessionNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateSessionNetworkConfig) GoString() string {
	return s.String()
}

func (s *UpdateSessionNetworkConfig) GetAllowOut() []*string {
	return s.AllowOut
}

func (s *UpdateSessionNetworkConfig) GetDenyOut() []*string {
	return s.DenyOut
}

func (s *UpdateSessionNetworkConfig) SetAllowOut(v []*string) *UpdateSessionNetworkConfig {
	s.AllowOut = v
	return s
}

func (s *UpdateSessionNetworkConfig) SetDenyOut(v []*string) *UpdateSessionNetworkConfig {
	s.DenyOut = v
	return s
}

func (s *UpdateSessionNetworkConfig) Validate() error {
	return dara.Validate(s)
}
