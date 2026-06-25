// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*UpdateConfigsRequestConfigs) *UpdateConfigsRequest
	GetConfigs() []*UpdateConfigsRequestConfigs
}

type UpdateConfigsRequest struct {
	// A list of workspace configurations to update or add.
	Configs []*UpdateConfigsRequestConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
}

func (s UpdateConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigsRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigsRequest) GetConfigs() []*UpdateConfigsRequestConfigs {
	return s.Configs
}

func (s *UpdateConfigsRequest) SetConfigs(v []*UpdateConfigsRequestConfigs) *UpdateConfigsRequest {
	s.Configs = v
	return s
}

func (s *UpdateConfigsRequest) Validate() error {
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

type UpdateConfigsRequestConfigs struct {
	// The category of the configuration item. The following categories are supported:
	//
	// - CommonResourceConfig: General resource configuration.
	//
	// - DLCAutoRecycle: DLC automatic recycling.
	//
	// - DLCPriorityConfig: DLC priority settings.
	//
	// - DSWPriorityConfig: DSW priority settings.
	//
	// - QuotaMaximumDuration: Configuration for the maximum runtime of a DLC job within a quota.
	//
	// - CommonTagConfig: Tag settings.
	//
	// example:
	//
	// CommonResourceConfig
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The key of the configuration item. The following keys are supported:
	//
	// - tempStoragePath: The path for temporary storage. This key is valid only when CategoryName is set to CommonResourceConfig.
	//
	// - isAutoRecycle: The configuration for automatic resource recycling. This key is valid only when CategoryName is set to DLCAutoRecycle.
	//
	// - priorityConfig: The priority configuration. This key is valid only when CategoryName is set to DLCPriorityConfig or DSWPriorityConfig.
	//
	// - quotaMaximumDuration: The maximum runtime configuration for a DLC job within a quota. This key is valid only when CategoryName is set to QuotaMaximumDuration.
	//
	// - predefinedTags: The predefined tags for the workspace. Created resources must have these tags.
	//
	// example:
	//
	// tempStoragePath
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// The value of the configuration item.
	//
	// - If ConfigKey is set to predefinedTags, the format of ConfigValue is [{"Type":"Tag","Key":"Key1","Value":"{\\\\"Products\\\\":\\\\"DLC,DSW,EAS\\\\",\\\\"Values\\\\":\\\\"value1,value2,value3\\\\"}"}]. The Products field specifies which products use the predefined tags.
	//
	// example:
	//
	// oss://test/s/
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// A list of tags for the configuration item.
	Labels []*UpdateConfigsRequestConfigsLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s UpdateConfigsRequestConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigsRequestConfigs) GoString() string {
	return s.String()
}

func (s *UpdateConfigsRequestConfigs) GetCategoryName() *string {
	return s.CategoryName
}

func (s *UpdateConfigsRequestConfigs) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *UpdateConfigsRequestConfigs) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *UpdateConfigsRequestConfigs) GetLabels() []*UpdateConfigsRequestConfigsLabels {
	return s.Labels
}

func (s *UpdateConfigsRequestConfigs) SetCategoryName(v string) *UpdateConfigsRequestConfigs {
	s.CategoryName = &v
	return s
}

func (s *UpdateConfigsRequestConfigs) SetConfigKey(v string) *UpdateConfigsRequestConfigs {
	s.ConfigKey = &v
	return s
}

func (s *UpdateConfigsRequestConfigs) SetConfigValue(v string) *UpdateConfigsRequestConfigs {
	s.ConfigValue = &v
	return s
}

func (s *UpdateConfigsRequestConfigs) SetLabels(v []*UpdateConfigsRequestConfigsLabels) *UpdateConfigsRequestConfigs {
	s.Labels = v
	return s
}

func (s *UpdateConfigsRequestConfigs) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateConfigsRequestConfigsLabels struct {
	// The key of the tag.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateConfigsRequestConfigsLabels) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigsRequestConfigsLabels) GoString() string {
	return s.String()
}

func (s *UpdateConfigsRequestConfigsLabels) GetKey() *string {
	return s.Key
}

func (s *UpdateConfigsRequestConfigsLabels) GetValue() *string {
	return s.Value
}

func (s *UpdateConfigsRequestConfigsLabels) SetKey(v string) *UpdateConfigsRequestConfigsLabels {
	s.Key = &v
	return s
}

func (s *UpdateConfigsRequestConfigsLabels) SetValue(v string) *UpdateConfigsRequestConfigsLabels {
	s.Value = &v
	return s
}

func (s *UpdateConfigsRequestConfigsLabels) Validate() error {
	return dara.Validate(s)
}
