// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddons(v []*ListAddonsResponseBodyAddons) *ListAddonsResponseBody
	GetAddons() []*ListAddonsResponseBodyAddons
}

type ListAddonsResponseBody struct {
	// The list of available components.
	Addons []*ListAddonsResponseBodyAddons `json:"addons,omitempty" xml:"addons,omitempty" type:"Repeated"`
}

func (s ListAddonsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBody) GetAddons() []*ListAddonsResponseBodyAddons {
	return s.Addons
}

func (s *ListAddonsResponseBody) SetAddons(v []*ListAddonsResponseBodyAddons) *ListAddonsResponseBody {
	s.Addons = v
	return s
}

func (s *ListAddonsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAddonsResponseBodyAddons struct {
	// Architectures supported by the component. Valid values:
	//
	// 	- amd64
	//
	// 	- arm64
	Architecture []*string `json:"architecture,omitempty" xml:"architecture,omitempty" type:"Repeated"`
	// The category of the component.
	//
	// example:
	//
	// monitor
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The schema of the custom parameters of the component.
	//
	// example:
	//
	// {}
	ConfigSchema *string `json:"config_schema,omitempty" xml:"config_schema,omitempty"`
	// Indicates whether the component is automatically installed by default.
	//
	// example:
	//
	// false
	InstallByDefault *bool `json:"install_by_default,omitempty" xml:"install_by_default,omitempty"`
	// Indicates whether the component is fully managed.
	//
	// example:
	//
	// false
	Managed *bool `json:"managed,omitempty" xml:"managed,omitempty"`
	// The component name.
	//
	// example:
	//
	// arms-prometheus
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Operations supported by the component. Valid values:
	//
	// 	- Install
	//
	// 	- Upgrade
	//
	// 	- Modify
	//
	// 	- Uninstall
	SupportedActions []*string `json:"supported_actions,omitempty" xml:"supported_actions,omitempty" type:"Repeated"`
	// The version number.
	//
	// example:
	//
	// 1.1.9
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListAddonsResponseBodyAddons) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsResponseBodyAddons) GoString() string {
	return s.String()
}

func (s *ListAddonsResponseBodyAddons) GetArchitecture() []*string {
	return s.Architecture
}

func (s *ListAddonsResponseBodyAddons) GetCategory() *string {
	return s.Category
}

func (s *ListAddonsResponseBodyAddons) GetConfigSchema() *string {
	return s.ConfigSchema
}

func (s *ListAddonsResponseBodyAddons) GetInstallByDefault() *bool {
	return s.InstallByDefault
}

func (s *ListAddonsResponseBodyAddons) GetManaged() *bool {
	return s.Managed
}

func (s *ListAddonsResponseBodyAddons) GetName() *string {
	return s.Name
}

func (s *ListAddonsResponseBodyAddons) GetSupportedActions() []*string {
	return s.SupportedActions
}

func (s *ListAddonsResponseBodyAddons) GetVersion() *string {
	return s.Version
}

func (s *ListAddonsResponseBodyAddons) SetArchitecture(v []*string) *ListAddonsResponseBodyAddons {
	s.Architecture = v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetCategory(v string) *ListAddonsResponseBodyAddons {
	s.Category = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetConfigSchema(v string) *ListAddonsResponseBodyAddons {
	s.ConfigSchema = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetInstallByDefault(v bool) *ListAddonsResponseBodyAddons {
	s.InstallByDefault = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetManaged(v bool) *ListAddonsResponseBodyAddons {
	s.Managed = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetName(v string) *ListAddonsResponseBodyAddons {
	s.Name = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetSupportedActions(v []*string) *ListAddonsResponseBodyAddons {
	s.SupportedActions = v
	return s
}

func (s *ListAddonsResponseBodyAddons) SetVersion(v string) *ListAddonsResponseBodyAddons {
	s.Version = &v
	return s
}

func (s *ListAddonsResponseBodyAddons) Validate() error {
	return dara.Validate(s)
}
