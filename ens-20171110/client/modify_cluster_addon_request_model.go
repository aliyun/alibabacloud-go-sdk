// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterAddonRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddon(v *ModifyClusterAddonRequestAddon) *ModifyClusterAddonRequest
	GetAddon() *ModifyClusterAddonRequestAddon
	SetClusterId(v string) *ModifyClusterAddonRequest
	GetClusterId() *string
	SetComponentName(v string) *ModifyClusterAddonRequest
	GetComponentName() *string
}

type ModifyClusterAddonRequest struct {
	Addon *ModifyClusterAddonRequestAddon `json:"Addon,omitempty" xml:"Addon,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// edge-csi-lite
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
}

func (s ModifyClusterAddonRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterAddonRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterAddonRequest) GetAddon() *ModifyClusterAddonRequestAddon {
	return s.Addon
}

func (s *ModifyClusterAddonRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyClusterAddonRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *ModifyClusterAddonRequest) SetAddon(v *ModifyClusterAddonRequestAddon) *ModifyClusterAddonRequest {
	s.Addon = v
	return s
}

func (s *ModifyClusterAddonRequest) SetClusterId(v string) *ModifyClusterAddonRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterAddonRequest) SetComponentName(v string) *ModifyClusterAddonRequest {
	s.ComponentName = &v
	return s
}

func (s *ModifyClusterAddonRequest) Validate() error {
	if s.Addon != nil {
		if err := s.Addon.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyClusterAddonRequestAddon struct {
	ConfigSchema []*ModifyClusterAddonRequestAddonConfigSchema `json:"ConfigSchema,omitempty" xml:"ConfigSchema,omitempty" type:"Repeated"`
	// example:
	//
	// edge-csi-lite
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// v1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ModifyClusterAddonRequestAddon) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterAddonRequestAddon) GoString() string {
	return s.String()
}

func (s *ModifyClusterAddonRequestAddon) GetConfigSchema() []*ModifyClusterAddonRequestAddonConfigSchema {
	return s.ConfigSchema
}

func (s *ModifyClusterAddonRequestAddon) GetName() *string {
	return s.Name
}

func (s *ModifyClusterAddonRequestAddon) GetVersion() *string {
	return s.Version
}

func (s *ModifyClusterAddonRequestAddon) SetConfigSchema(v []*ModifyClusterAddonRequestAddonConfigSchema) *ModifyClusterAddonRequestAddon {
	s.ConfigSchema = v
	return s
}

func (s *ModifyClusterAddonRequestAddon) SetName(v string) *ModifyClusterAddonRequestAddon {
	s.Name = &v
	return s
}

func (s *ModifyClusterAddonRequestAddon) SetVersion(v string) *ModifyClusterAddonRequestAddon {
	s.Version = &v
	return s
}

func (s *ModifyClusterAddonRequestAddon) Validate() error {
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

type ModifyClusterAddonRequestAddonConfigSchema struct {
	// example:
	//
	// d0ead1f4e28de0f9e3c86588409a88a4
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// d0ead1f4e28de0f9e3c86588409a88a4
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

func (s ModifyClusterAddonRequestAddonConfigSchema) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterAddonRequestAddonConfigSchema) GoString() string {
	return s.String()
}

func (s *ModifyClusterAddonRequestAddonConfigSchema) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ModifyClusterAddonRequestAddonConfigSchema) GetConfigVersion() *string {
	return s.ConfigVersion
}

func (s *ModifyClusterAddonRequestAddonConfigSchema) GetName() *string {
	return s.Name
}

func (s *ModifyClusterAddonRequestAddonConfigSchema) GetParams() map[string]interface{} {
	return s.Params
}

func (s *ModifyClusterAddonRequestAddonConfigSchema) SetAppVersion(v string) *ModifyClusterAddonRequestAddonConfigSchema {
	s.AppVersion = &v
	return s
}

func (s *ModifyClusterAddonRequestAddonConfigSchema) SetConfigVersion(v string) *ModifyClusterAddonRequestAddonConfigSchema {
	s.ConfigVersion = &v
	return s
}

func (s *ModifyClusterAddonRequestAddonConfigSchema) SetName(v string) *ModifyClusterAddonRequestAddonConfigSchema {
	s.Name = &v
	return s
}

func (s *ModifyClusterAddonRequestAddonConfigSchema) SetParams(v map[string]interface{}) *ModifyClusterAddonRequestAddonConfigSchema {
	s.Params = v
	return s
}

func (s *ModifyClusterAddonRequestAddonConfigSchema) Validate() error {
	return dara.Validate(s)
}
