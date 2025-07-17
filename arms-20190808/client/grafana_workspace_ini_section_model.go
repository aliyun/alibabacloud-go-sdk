// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceIniSection interface {
	dara.Model
	String() string
	GoString() string
	SetPropertys(v []*GrafanaWorkspaceIniProperty) *GrafanaWorkspaceIniSection
	GetPropertys() []*GrafanaWorkspaceIniProperty
	SetSection(v string) *GrafanaWorkspaceIniSection
	GetSection() *string
}

type GrafanaWorkspaceIniSection struct {
	Propertys []*GrafanaWorkspaceIniProperty `json:"propertys,omitempty" xml:"propertys,omitempty" type:"Repeated"`
	Section   *string                        `json:"section,omitempty" xml:"section,omitempty"`
}

func (s GrafanaWorkspaceIniSection) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceIniSection) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceIniSection) GetPropertys() []*GrafanaWorkspaceIniProperty {
	return s.Propertys
}

func (s *GrafanaWorkspaceIniSection) GetSection() *string {
	return s.Section
}

func (s *GrafanaWorkspaceIniSection) SetPropertys(v []*GrafanaWorkspaceIniProperty) *GrafanaWorkspaceIniSection {
	s.Propertys = v
	return s
}

func (s *GrafanaWorkspaceIniSection) SetSection(v string) *GrafanaWorkspaceIniSection {
	s.Section = &v
	return s
}

func (s *GrafanaWorkspaceIniSection) Validate() error {
	return dara.Validate(s)
}
