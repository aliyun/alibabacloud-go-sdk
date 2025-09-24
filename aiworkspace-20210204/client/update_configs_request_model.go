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
	// The list of workspace configurations to update or add.
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
	return dara.Validate(s)
}

type UpdateConfigsRequestConfigs struct {
	// The category of the configuration item. Supported categories:
	//
	// 	- CommonResourceConfig
	//
	// 	- DLCAutoRecycle
	//
	// 	- DLCPriorityConfig
	//
	// 	- DSWPriorityConfig
	//
	// 	- QuotaMaximumDuration
	//
	// 	- CommonTagConfig
	//
	// example:
	//
	// CommonResourceConfig
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The key of the configuration item. Supported keys:
	//
	// 	- tempStoragePath: Temporary storage path. This key can be used only when CategoryName is set to CommonResourceConfig.
	//
	// 	- isAutoRecycle: Automatic recycle configuration. This key can be used only when CategoryName is set to DLCAutoRecycle.
	//
	// 	- tempStoragePath: Temporary storage path. This key can be used only when CategoryName is set to CommonResourceConfig.
	//
	// 	- quotaMaximumDuration: Maximum run time of DLC jobs for a quota. This key can be used only when CategoryName is set to QuotaMaximumDuration.
	//
	// 	- predefinedTags: The predefined tags of the workspace. All created resources must have tags.
	//
	// example:
	//
	// tempStoragePath
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// The value of the configuration item.
	//
	// 	- When ConfigKey is predefinedTags, the ConfigValue follows this format: [{"Type":"Tag","Key":"Key1","Value":"{"Products":"DLC,DSW,EAS","Values":"value1,value2,value3"}"}]. "Products" indicates the products that use the predefined tags.
	//
	// example:
	//
	// oss://test/s/
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The tags of the configuration item.
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
	return dara.Validate(s)
}

type UpdateConfigsRequestConfigsLabels struct {
	// The tag key.
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
