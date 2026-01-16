// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallClusterAddonsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddons(v []*InstallClusterAddonsRequestAddons) *InstallClusterAddonsRequest
	GetAddons() []*InstallClusterAddonsRequestAddons
	SetClusterId(v string) *InstallClusterAddonsRequest
	GetClusterId() *string
}

type InstallClusterAddonsRequest struct {
	// This parameter is required.
	Addons []*InstallClusterAddonsRequestAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s InstallClusterAddonsRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallClusterAddonsRequest) GoString() string {
	return s.String()
}

func (s *InstallClusterAddonsRequest) GetAddons() []*InstallClusterAddonsRequestAddons {
	return s.Addons
}

func (s *InstallClusterAddonsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *InstallClusterAddonsRequest) SetAddons(v []*InstallClusterAddonsRequestAddons) *InstallClusterAddonsRequest {
	s.Addons = v
	return s
}

func (s *InstallClusterAddonsRequest) SetClusterId(v string) *InstallClusterAddonsRequest {
	s.ClusterId = &v
	return s
}

func (s *InstallClusterAddonsRequest) Validate() error {
	if s.Addons != nil {
		for _, item := range s.Addons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InstallClusterAddonsRequestAddons struct {
	ConfigSchema []*InstallClusterAddonsRequestAddonsConfigSchema `json:"ConfigSchema,omitempty" xml:"ConfigSchema,omitempty" type:"Repeated"`
	// example:
	//
	// edge-csi-lite
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// v1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s InstallClusterAddonsRequestAddons) String() string {
	return dara.Prettify(s)
}

func (s InstallClusterAddonsRequestAddons) GoString() string {
	return s.String()
}

func (s *InstallClusterAddonsRequestAddons) GetConfigSchema() []*InstallClusterAddonsRequestAddonsConfigSchema {
	return s.ConfigSchema
}

func (s *InstallClusterAddonsRequestAddons) GetName() *string {
	return s.Name
}

func (s *InstallClusterAddonsRequestAddons) GetVersion() *string {
	return s.Version
}

func (s *InstallClusterAddonsRequestAddons) SetConfigSchema(v []*InstallClusterAddonsRequestAddonsConfigSchema) *InstallClusterAddonsRequestAddons {
	s.ConfigSchema = v
	return s
}

func (s *InstallClusterAddonsRequestAddons) SetName(v string) *InstallClusterAddonsRequestAddons {
	s.Name = &v
	return s
}

func (s *InstallClusterAddonsRequestAddons) SetVersion(v string) *InstallClusterAddonsRequestAddons {
	s.Version = &v
	return s
}

func (s *InstallClusterAddonsRequestAddons) Validate() error {
	if s.ConfigSchema != nil {
		for _, item := range s.ConfigSchema {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InstallClusterAddonsRequestAddonsConfigSchema struct {
	// example:
	//
	// 859e9d595b2974ed79c444658d1dea89
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 4155709cd12a09bdb8cbaca71bf03233
	ConfigVersion *string `json:"ConfigVersion,omitempty" xml:"ConfigVersion,omitempty"`
	// example:
	//
	// edge-csi-lite
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {"k1":"v1"}
	Params map[string]interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s InstallClusterAddonsRequestAddonsConfigSchema) String() string {
	return dara.Prettify(s)
}

func (s InstallClusterAddonsRequestAddonsConfigSchema) GoString() string {
	return s.String()
}

func (s *InstallClusterAddonsRequestAddonsConfigSchema) GetAppVersion() *string {
	return s.AppVersion
}

func (s *InstallClusterAddonsRequestAddonsConfigSchema) GetConfigVersion() *string {
	return s.ConfigVersion
}

func (s *InstallClusterAddonsRequestAddonsConfigSchema) GetName() *string {
	return s.Name
}

func (s *InstallClusterAddonsRequestAddonsConfigSchema) GetParams() map[string]interface{} {
	return s.Params
}

func (s *InstallClusterAddonsRequestAddonsConfigSchema) SetAppVersion(v string) *InstallClusterAddonsRequestAddonsConfigSchema {
	s.AppVersion = &v
	return s
}

func (s *InstallClusterAddonsRequestAddonsConfigSchema) SetConfigVersion(v string) *InstallClusterAddonsRequestAddonsConfigSchema {
	s.ConfigVersion = &v
	return s
}

func (s *InstallClusterAddonsRequestAddonsConfigSchema) SetName(v string) *InstallClusterAddonsRequestAddonsConfigSchema {
	s.Name = &v
	return s
}

func (s *InstallClusterAddonsRequestAddonsConfigSchema) SetParams(v map[string]interface{}) *InstallClusterAddonsRequestAddonsConfigSchema {
	s.Params = v
	return s
}

func (s *InstallClusterAddonsRequestAddonsConfigSchema) Validate() error {
	return dara.Validate(s)
}
