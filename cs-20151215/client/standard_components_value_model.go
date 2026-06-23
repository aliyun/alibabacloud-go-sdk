// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStandardComponentsValue interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *StandardComponentsValue
	GetName() *string
	SetVersion(v string) *StandardComponentsValue
	GetVersion() *string
	SetDescription(v string) *StandardComponentsValue
	GetDescription() *string
	SetRequired(v string) *StandardComponentsValue
	GetRequired() *string
	SetDisabled(v bool) *StandardComponentsValue
	GetDisabled() *bool
}

type StandardComponentsValue struct {
	// The component name.
	//
	// example:
	//
	// ack-arena
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The component version.
	//
	// example:
	//
	// 0.5.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The description of the component features.
	//
	// example:
	//
	// ***
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Indicates whether the component is required by the cluster. Valid values:
	//
	// - `true`: The component is required and must be installed when the cluster is created.
	//
	// - `false`: The component is optional and can be installed through Component Management after the cluster is created.
	//
	// example:
	//
	// false
	Required *string `json:"required,omitempty" xml:"required,omitempty"`
	// Indicates whether default installation is disabled. When a cluster is created, in addition to the components required by the cluster, some logging or routing-related components (such as Ingress) are also installed by default. If you do not want to install them by default, you can set this field to disable default installation. Valid values:
	//
	// - `true`: Default installation is disabled.
	//
	// - `false`: Default installation is enabled.
	//
	// example:
	//
	// false
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
}

func (s StandardComponentsValue) String() string {
	return dara.Prettify(s)
}

func (s StandardComponentsValue) GoString() string {
	return s.String()
}

func (s *StandardComponentsValue) GetName() *string {
	return s.Name
}

func (s *StandardComponentsValue) GetVersion() *string {
	return s.Version
}

func (s *StandardComponentsValue) GetDescription() *string {
	return s.Description
}

func (s *StandardComponentsValue) GetRequired() *string {
	return s.Required
}

func (s *StandardComponentsValue) GetDisabled() *bool {
	return s.Disabled
}

func (s *StandardComponentsValue) SetName(v string) *StandardComponentsValue {
	s.Name = &v
	return s
}

func (s *StandardComponentsValue) SetVersion(v string) *StandardComponentsValue {
	s.Version = &v
	return s
}

func (s *StandardComponentsValue) SetDescription(v string) *StandardComponentsValue {
	s.Description = &v
	return s
}

func (s *StandardComponentsValue) SetRequired(v string) *StandardComponentsValue {
	s.Required = &v
	return s
}

func (s *StandardComponentsValue) SetDisabled(v bool) *StandardComponentsValue {
	s.Disabled = &v
	return s
}

func (s *StandardComponentsValue) Validate() error {
	return dara.Validate(s)
}
