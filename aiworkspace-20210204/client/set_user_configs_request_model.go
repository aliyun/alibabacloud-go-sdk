// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*SetUserConfigsRequestConfigs) *SetUserConfigsRequest
	GetConfigs() []*SetUserConfigsRequestConfigs
}

type SetUserConfigsRequest struct {
	// The configurations list.
	Configs []*SetUserConfigsRequestConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
}

func (s SetUserConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s SetUserConfigsRequest) GoString() string {
	return s.String()
}

func (s *SetUserConfigsRequest) GetConfigs() []*SetUserConfigsRequestConfigs {
	return s.Configs
}

func (s *SetUserConfigsRequest) SetConfigs(v []*SetUserConfigsRequestConfigs) *SetUserConfigsRequest {
	s.Configs = v
	return s
}

func (s *SetUserConfigsRequest) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SetUserConfigsRequestConfigs struct {
	// The category. Only DataPrivacyConfig is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// DataPrivacyConfig
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The key of the configuration item.
	//
	// This parameter is required.
	//
	// example:
	//
	// customizePAIAssumedRole
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// The value of the configuration item.
	//
	// This parameter is required.
	//
	// example:
	//
	// role
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The scope. Valid values: subUser and owner.
	//
	// This parameter is required.
	//
	// example:
	//
	// owner
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s SetUserConfigsRequestConfigs) String() string {
	return dara.Prettify(s)
}

func (s SetUserConfigsRequestConfigs) GoString() string {
	return s.String()
}

func (s *SetUserConfigsRequestConfigs) GetCategoryName() *string {
	return s.CategoryName
}

func (s *SetUserConfigsRequestConfigs) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *SetUserConfigsRequestConfigs) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *SetUserConfigsRequestConfigs) GetScope() *string {
	return s.Scope
}

func (s *SetUserConfigsRequestConfigs) SetCategoryName(v string) *SetUserConfigsRequestConfigs {
	s.CategoryName = &v
	return s
}

func (s *SetUserConfigsRequestConfigs) SetConfigKey(v string) *SetUserConfigsRequestConfigs {
	s.ConfigKey = &v
	return s
}

func (s *SetUserConfigsRequestConfigs) SetConfigValue(v string) *SetUserConfigsRequestConfigs {
	s.ConfigValue = &v
	return s
}

func (s *SetUserConfigsRequestConfigs) SetScope(v string) *SetUserConfigsRequestConfigs {
	s.Scope = &v
	return s
}

func (s *SetUserConfigsRequestConfigs) Validate() error {
	return dara.Validate(s)
}
