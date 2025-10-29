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
	// example:
	//
	// engineVersion
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 6.7
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The moderation results.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The monitoring type. Valid values:
	//
	// 	- checkClusterHealth: Cluster Health Status
	//
	// 	- checkConfigCompatible: Configuration Compatibility Status
	//
	// 	- checkClusterResource: resource space status
	//
	// 	- checkClusterSnapshot: Whether a snapshot exists
	//
	// example:
	//
	// false
	DryRun         *bool   `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
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
	Enable      *string `json:"enable,omitempty" xml:"enable,omitempty"`
	FileVersion *string `json:"fileVersion,omitempty" xml:"fileVersion,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Version     *string `json:"version,omitempty" xml:"version,omitempty"`
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
