// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContainerdConfig interface {
	dara.Model
	String() string
	GoString() string
	SetInsecureRegistries(v []*string) *ContainerdConfig
	GetInsecureRegistries() []*string
	SetRegistryMirrors(v []*string) *ContainerdConfig
	GetRegistryMirrors() []*string
}

type ContainerdConfig struct {
	InsecureRegistries []*string `json:"insecureRegistries,omitempty" xml:"insecureRegistries,omitempty" type:"Repeated"`
	RegistryMirrors    []*string `json:"registryMirrors,omitempty" xml:"registryMirrors,omitempty" type:"Repeated"`
}

func (s ContainerdConfig) String() string {
	return dara.Prettify(s)
}

func (s ContainerdConfig) GoString() string {
	return s.String()
}

func (s *ContainerdConfig) GetInsecureRegistries() []*string {
	return s.InsecureRegistries
}

func (s *ContainerdConfig) GetRegistryMirrors() []*string {
	return s.RegistryMirrors
}

func (s *ContainerdConfig) SetInsecureRegistries(v []*string) *ContainerdConfig {
	s.InsecureRegistries = v
	return s
}

func (s *ContainerdConfig) SetRegistryMirrors(v []*string) *ContainerdConfig {
	s.RegistryMirrors = v
	return s
}

func (s *ContainerdConfig) Validate() error {
	return dara.Validate(s)
}
