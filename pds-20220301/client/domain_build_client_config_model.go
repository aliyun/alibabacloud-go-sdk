// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDomainBuildClientConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCopyright(v string) *DomainBuildClientConfig
	GetCopyright() *string
	SetName(v string) *DomainBuildClientConfig
	GetName() *string
}

type DomainBuildClientConfig struct {
	Copyright *string `json:"copyright,omitempty" xml:"copyright,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DomainBuildClientConfig) String() string {
	return dara.Prettify(s)
}

func (s DomainBuildClientConfig) GoString() string {
	return s.String()
}

func (s *DomainBuildClientConfig) GetCopyright() *string {
	return s.Copyright
}

func (s *DomainBuildClientConfig) GetName() *string {
	return s.Name
}

func (s *DomainBuildClientConfig) SetCopyright(v string) *DomainBuildClientConfig {
	s.Copyright = &v
	return s
}

func (s *DomainBuildClientConfig) SetName(v string) *DomainBuildClientConfig {
	s.Name = &v
	return s
}

func (s *DomainBuildClientConfig) Validate() error {
	return dara.Validate(s)
}
