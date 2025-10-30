// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArchitecture(v []*string) *DescribeAddonResponseBody
	GetArchitecture() []*string
	SetCategory(v string) *DescribeAddonResponseBody
	GetCategory() *string
	SetConfigSchema(v string) *DescribeAddonResponseBody
	GetConfigSchema() *string
	SetInstallByDefault(v bool) *DescribeAddonResponseBody
	GetInstallByDefault() *bool
	SetManaged(v bool) *DescribeAddonResponseBody
	GetManaged() *bool
	SetName(v string) *DescribeAddonResponseBody
	GetName() *string
	SetNewerVersions(v []*DescribeAddonResponseBodyNewerVersions) *DescribeAddonResponseBody
	GetNewerVersions() []*DescribeAddonResponseBodyNewerVersions
	SetSupportedActions(v []*string) *DescribeAddonResponseBody
	GetSupportedActions() []*string
	SetVersion(v string) *DescribeAddonResponseBody
	GetVersion() *string
}

type DescribeAddonResponseBody struct {
	// The CPU architecture supported by the component.
	Architecture []*string `json:"architecture,omitempty" xml:"architecture,omitempty" type:"Repeated"`
	// The category of the component.
	//
	// example:
	//
	// network
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The custom parameter schema of the component.
	//
	// example:
	//
	// {}
	ConfigSchema *string `json:"config_schema,omitempty" xml:"config_schema,omitempty"`
	// Indicates whether the component is automatically installed by default.
	//
	// example:
	//
	// true
	InstallByDefault *bool `json:"install_by_default,omitempty" xml:"install_by_default,omitempty"`
	// Indicates whether the component is fully managed.
	//
	// example:
	//
	// false
	Managed *bool `json:"managed,omitempty" xml:"managed,omitempty"`
	// The name of the component.
	//
	// example:
	//
	// coredns
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The latest version information of the component.
	NewerVersions []*DescribeAddonResponseBodyNewerVersions `json:"newer_versions,omitempty" xml:"newer_versions,omitempty" type:"Repeated"`
	// The operations supported by the component.
	SupportedActions []*string `json:"supported_actions,omitempty" xml:"supported_actions,omitempty" type:"Repeated"`
	// The version of the component.
	//
	// example:
	//
	// v1.9.3.6-32932850-aliyun
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeAddonResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAddonResponseBody) GetArchitecture() []*string {
	return s.Architecture
}

func (s *DescribeAddonResponseBody) GetCategory() *string {
	return s.Category
}

func (s *DescribeAddonResponseBody) GetConfigSchema() *string {
	return s.ConfigSchema
}

func (s *DescribeAddonResponseBody) GetInstallByDefault() *bool {
	return s.InstallByDefault
}

func (s *DescribeAddonResponseBody) GetManaged() *bool {
	return s.Managed
}

func (s *DescribeAddonResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeAddonResponseBody) GetNewerVersions() []*DescribeAddonResponseBodyNewerVersions {
	return s.NewerVersions
}

func (s *DescribeAddonResponseBody) GetSupportedActions() []*string {
	return s.SupportedActions
}

func (s *DescribeAddonResponseBody) GetVersion() *string {
	return s.Version
}

func (s *DescribeAddonResponseBody) SetArchitecture(v []*string) *DescribeAddonResponseBody {
	s.Architecture = v
	return s
}

func (s *DescribeAddonResponseBody) SetCategory(v string) *DescribeAddonResponseBody {
	s.Category = &v
	return s
}

func (s *DescribeAddonResponseBody) SetConfigSchema(v string) *DescribeAddonResponseBody {
	s.ConfigSchema = &v
	return s
}

func (s *DescribeAddonResponseBody) SetInstallByDefault(v bool) *DescribeAddonResponseBody {
	s.InstallByDefault = &v
	return s
}

func (s *DescribeAddonResponseBody) SetManaged(v bool) *DescribeAddonResponseBody {
	s.Managed = &v
	return s
}

func (s *DescribeAddonResponseBody) SetName(v string) *DescribeAddonResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeAddonResponseBody) SetNewerVersions(v []*DescribeAddonResponseBodyNewerVersions) *DescribeAddonResponseBody {
	s.NewerVersions = v
	return s
}

func (s *DescribeAddonResponseBody) SetSupportedActions(v []*string) *DescribeAddonResponseBody {
	s.SupportedActions = v
	return s
}

func (s *DescribeAddonResponseBody) SetVersion(v string) *DescribeAddonResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeAddonResponseBody) Validate() error {
	if s.NewerVersions != nil {
		for _, item := range s.NewerVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAddonResponseBodyNewerVersions struct {
	// The minimum cluster version required by the component version.
	//
	// example:
	//
	// 1.20.4
	MinimumClusterVersion *string `json:"minimum_cluster_version,omitempty" xml:"minimum_cluster_version,omitempty"`
	// Indicates whether the component can be updated to the version.
	//
	// 	- true: yes
	//
	// 	- false: no
	//
	// example:
	//
	// true
	Upgradable *bool `json:"upgradable,omitempty" xml:"upgradable,omitempty"`
	// The latest version number of the component.
	//
	// example:
	//
	// v1.9.3.10-7dfca203-aliyun
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeAddonResponseBodyNewerVersions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonResponseBodyNewerVersions) GoString() string {
	return s.String()
}

func (s *DescribeAddonResponseBodyNewerVersions) GetMinimumClusterVersion() *string {
	return s.MinimumClusterVersion
}

func (s *DescribeAddonResponseBodyNewerVersions) GetUpgradable() *bool {
	return s.Upgradable
}

func (s *DescribeAddonResponseBodyNewerVersions) GetVersion() *string {
	return s.Version
}

func (s *DescribeAddonResponseBodyNewerVersions) SetMinimumClusterVersion(v string) *DescribeAddonResponseBodyNewerVersions {
	s.MinimumClusterVersion = &v
	return s
}

func (s *DescribeAddonResponseBodyNewerVersions) SetUpgradable(v bool) *DescribeAddonResponseBodyNewerVersions {
	s.Upgradable = &v
	return s
}

func (s *DescribeAddonResponseBodyNewerVersions) SetVersion(v string) *DescribeAddonResponseBodyNewerVersions {
	s.Version = &v
	return s
}

func (s *DescribeAddonResponseBodyNewerVersions) Validate() error {
	return dara.Validate(s)
}
