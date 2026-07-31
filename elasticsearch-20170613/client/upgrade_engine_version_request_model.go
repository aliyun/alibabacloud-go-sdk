// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeEngineVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlugins(v []*UpgradeEngineVersionRequestPlugins) *UpgradeEngineVersionRequest
	GetPlugins() []*UpgradeEngineVersionRequestPlugins
	SetType(v string) *UpgradeEngineVersionRequest
	GetType() *string
	SetVersion(v string) *UpgradeEngineVersionRequest
	GetVersion() *string
	SetClientToken(v string) *UpgradeEngineVersionRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpgradeEngineVersionRequest
	GetDryRun() *bool
	SetUpdateStrategy(v string) *UpgradeEngineVersionRequest
	GetUpdateStrategy() *string
}

type UpgradeEngineVersionRequest struct {
	Plugins []*UpgradeEngineVersionRequestPlugins `json:"plugins,omitempty" xml:"plugins,omitempty" type:"Repeated"`
	// The upgrade type. Valid values:
	//
	// - engineVersion (default): major engine version upgrade.
	//
	// - aliVersion: kernel version upgrade.
	//
	// example:
	//
	// engineVersion
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version after the upgrade. If type is set to engineVersion, the value is the instance version, such as 6.7. If type is set to aliVersion, the value is the kernel version, such as ali1.2.0.
	//
	// example:
	//
	// 6.7
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// Specifies whether to perform a pre-upgrade check. Valid values:
	//
	// - true: performs a check.
	//
	// - false (default): does not perform a check.
	//
	// 	Warning:  The version upgrade check involves checks on cluster YML, plug-in configurations, cluster status, indexes, and resources. Perform a pre-upgrade check before upgrading. Otherwise, upgrade issues may occur.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The update strategy. Valid values:
	//
	// - blue_green: blue-green deployment.
	//
	// - normal: in-place update.
	//
	// - intelligent: intelligent update.
	UpdateStrategy *string `json:"updateStrategy,omitempty" xml:"updateStrategy,omitempty"`
}

func (s UpgradeEngineVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeEngineVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeEngineVersionRequest) GetPlugins() []*UpgradeEngineVersionRequestPlugins {
	return s.Plugins
}

func (s *UpgradeEngineVersionRequest) GetType() *string {
	return s.Type
}

func (s *UpgradeEngineVersionRequest) GetVersion() *string {
	return s.Version
}

func (s *UpgradeEngineVersionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpgradeEngineVersionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpgradeEngineVersionRequest) GetUpdateStrategy() *string {
	return s.UpdateStrategy
}

func (s *UpgradeEngineVersionRequest) SetPlugins(v []*UpgradeEngineVersionRequestPlugins) *UpgradeEngineVersionRequest {
	s.Plugins = v
	return s
}

func (s *UpgradeEngineVersionRequest) SetType(v string) *UpgradeEngineVersionRequest {
	s.Type = &v
	return s
}

func (s *UpgradeEngineVersionRequest) SetVersion(v string) *UpgradeEngineVersionRequest {
	s.Version = &v
	return s
}

func (s *UpgradeEngineVersionRequest) SetClientToken(v string) *UpgradeEngineVersionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeEngineVersionRequest) SetDryRun(v bool) *UpgradeEngineVersionRequest {
	s.DryRun = &v
	return s
}

func (s *UpgradeEngineVersionRequest) SetUpdateStrategy(v string) *UpgradeEngineVersionRequest {
	s.UpdateStrategy = &v
	return s
}

func (s *UpgradeEngineVersionRequest) Validate() error {
	if s.Plugins != nil {
		for _, item := range s.Plugins {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpgradeEngineVersionRequestPlugins struct {
	// Specifies whether to install or uninstall the plug-in. A value of true indicates install, and a value of false indicates uninstall.
	Enable *string `json:"enable,omitempty" xml:"enable,omitempty"`
	// The fileVersion of the plug-in. Refer to the response of ListUserPlugin.
	FileVersion *string `json:"fileVersion,omitempty" xml:"fileVersion,omitempty"`
	// The plug-in name.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Elasticsearch version for the plug-in, such as 7.16.2.
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpgradeEngineVersionRequestPlugins) String() string {
	return dara.Prettify(s)
}

func (s UpgradeEngineVersionRequestPlugins) GoString() string {
	return s.String()
}

func (s *UpgradeEngineVersionRequestPlugins) GetEnable() *string {
	return s.Enable
}

func (s *UpgradeEngineVersionRequestPlugins) GetFileVersion() *string {
	return s.FileVersion
}

func (s *UpgradeEngineVersionRequestPlugins) GetName() *string {
	return s.Name
}

func (s *UpgradeEngineVersionRequestPlugins) GetVersion() *string {
	return s.Version
}

func (s *UpgradeEngineVersionRequestPlugins) SetEnable(v string) *UpgradeEngineVersionRequestPlugins {
	s.Enable = &v
	return s
}

func (s *UpgradeEngineVersionRequestPlugins) SetFileVersion(v string) *UpgradeEngineVersionRequestPlugins {
	s.FileVersion = &v
	return s
}

func (s *UpgradeEngineVersionRequestPlugins) SetName(v string) *UpgradeEngineVersionRequestPlugins {
	s.Name = &v
	return s
}

func (s *UpgradeEngineVersionRequestPlugins) SetVersion(v string) *UpgradeEngineVersionRequestPlugins {
	s.Version = &v
	return s
}

func (s *UpgradeEngineVersionRequestPlugins) Validate() error {
	return dara.Validate(s)
}
