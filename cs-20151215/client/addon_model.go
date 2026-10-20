// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddon interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *Addon
	GetConfig() *string
	SetDisabled(v bool) *Addon
	GetDisabled() *bool
	SetName(v string) *Addon
	GetName() *string
	SetVersion(v string) *Addon
	GetVersion() *string
}

type Addon struct {
	// The configuration of the component.
	//
	// example:
	//
	// {\\"IngressSlbNetworkType\\":\\"internet\\"}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// Specifies whether to disable default installation. When a cluster is created, required components and additional components such as log components are installed by default. You can disable this default behavior and install these components later by calling the component installation API or by using the console. Valid values:
	//
	// - `true`: Disables default installation.
	//
	// - `false`: Allows default installation.
	//
	// example:
	//
	// false
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// The name of the component.
	//
	// example:
	//
	// nginx-ingress-controller
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version of the component.
	//
	// example:
	//
	// v1.9.3-aliyun.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s Addon) String() string {
	return dara.Prettify(s)
}

func (s Addon) GoString() string {
	return s.String()
}

func (s *Addon) GetConfig() *string {
	return s.Config
}

func (s *Addon) GetDisabled() *bool {
	return s.Disabled
}

func (s *Addon) GetName() *string {
	return s.Name
}

func (s *Addon) GetVersion() *string {
	return s.Version
}

func (s *Addon) SetConfig(v string) *Addon {
	s.Config = &v
	return s
}

func (s *Addon) SetDisabled(v bool) *Addon {
	s.Disabled = &v
	return s
}

func (s *Addon) SetName(v string) *Addon {
	s.Name = &v
	return s
}

func (s *Addon) SetVersion(v string) *Addon {
	s.Version = &v
	return s
}

func (s *Addon) Validate() error {
	return dara.Validate(s)
}
