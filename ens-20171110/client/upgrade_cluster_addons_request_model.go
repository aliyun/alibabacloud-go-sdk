// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClusterAddonsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddons(v []*UpgradeClusterAddonsRequestAddons) *UpgradeClusterAddonsRequest
	GetAddons() []*UpgradeClusterAddonsRequestAddons
	SetClusterId(v string) *UpgradeClusterAddonsRequest
	GetClusterId() *string
}

type UpgradeClusterAddonsRequest struct {
	Addons []*UpgradeClusterAddonsRequestAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UpgradeClusterAddonsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterAddonsRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClusterAddonsRequest) GetAddons() []*UpgradeClusterAddonsRequestAddons {
	return s.Addons
}

func (s *UpgradeClusterAddonsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpgradeClusterAddonsRequest) SetAddons(v []*UpgradeClusterAddonsRequestAddons) *UpgradeClusterAddonsRequest {
	s.Addons = v
	return s
}

func (s *UpgradeClusterAddonsRequest) SetClusterId(v string) *UpgradeClusterAddonsRequest {
	s.ClusterId = &v
	return s
}

func (s *UpgradeClusterAddonsRequest) Validate() error {
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

type UpgradeClusterAddonsRequestAddons struct {
	ConfigSchema []*UpgradeClusterAddonsRequestAddonsConfigSchema `json:"ConfigSchema,omitempty" xml:"ConfigSchema,omitempty" type:"Repeated"`
	// example:
	//
	// edge-csi-lite
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// v2
	NextVersion *string `json:"NextVersion,omitempty" xml:"NextVersion,omitempty"`
}

func (s UpgradeClusterAddonsRequestAddons) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterAddonsRequestAddons) GoString() string {
	return s.String()
}

func (s *UpgradeClusterAddonsRequestAddons) GetConfigSchema() []*UpgradeClusterAddonsRequestAddonsConfigSchema {
	return s.ConfigSchema
}

func (s *UpgradeClusterAddonsRequestAddons) GetName() *string {
	return s.Name
}

func (s *UpgradeClusterAddonsRequestAddons) GetNextVersion() *string {
	return s.NextVersion
}

func (s *UpgradeClusterAddonsRequestAddons) SetConfigSchema(v []*UpgradeClusterAddonsRequestAddonsConfigSchema) *UpgradeClusterAddonsRequestAddons {
	s.ConfigSchema = v
	return s
}

func (s *UpgradeClusterAddonsRequestAddons) SetName(v string) *UpgradeClusterAddonsRequestAddons {
	s.Name = &v
	return s
}

func (s *UpgradeClusterAddonsRequestAddons) SetNextVersion(v string) *UpgradeClusterAddonsRequestAddons {
	s.NextVersion = &v
	return s
}

func (s *UpgradeClusterAddonsRequestAddons) Validate() error {
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

type UpgradeClusterAddonsRequestAddonsConfigSchema struct {
	// example:
	//
	// 4155709cd12a09bdb8cbaca71bf03233
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

func (s UpgradeClusterAddonsRequestAddonsConfigSchema) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterAddonsRequestAddonsConfigSchema) GoString() string {
	return s.String()
}

func (s *UpgradeClusterAddonsRequestAddonsConfigSchema) GetAppVersion() *string {
	return s.AppVersion
}

func (s *UpgradeClusterAddonsRequestAddonsConfigSchema) GetConfigVersion() *string {
	return s.ConfigVersion
}

func (s *UpgradeClusterAddonsRequestAddonsConfigSchema) GetName() *string {
	return s.Name
}

func (s *UpgradeClusterAddonsRequestAddonsConfigSchema) GetParams() map[string]interface{} {
	return s.Params
}

func (s *UpgradeClusterAddonsRequestAddonsConfigSchema) SetAppVersion(v string) *UpgradeClusterAddonsRequestAddonsConfigSchema {
	s.AppVersion = &v
	return s
}

func (s *UpgradeClusterAddonsRequestAddonsConfigSchema) SetConfigVersion(v string) *UpgradeClusterAddonsRequestAddonsConfigSchema {
	s.ConfigVersion = &v
	return s
}

func (s *UpgradeClusterAddonsRequestAddonsConfigSchema) SetName(v string) *UpgradeClusterAddonsRequestAddonsConfigSchema {
	s.Name = &v
	return s
}

func (s *UpgradeClusterAddonsRequestAddonsConfigSchema) SetParams(v map[string]interface{}) *UpgradeClusterAddonsRequestAddonsConfigSchema {
	s.Params = v
	return s
}

func (s *UpgradeClusterAddonsRequestAddonsConfigSchema) Validate() error {
	return dara.Validate(s)
}
