// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryName(v string) *UpdateConfigRequest
	GetCategoryName() *string
	SetConfigKey(v string) *UpdateConfigRequest
	GetConfigKey() *string
	SetConfigValue(v string) *UpdateConfigRequest
	GetConfigValue() *string
	SetLabels(v []*UpdateConfigRequestLabels) *UpdateConfigRequest
	GetLabels() []*UpdateConfigRequestLabels
}

type UpdateConfigRequest struct {
	// The category of the configuration item. Valid values:
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
	// The key of the configuration item. Valid values:
	//
	// 	- tempStoragePath: Temporary storage path. This key can be used only when CategoryName is set to CommonResourceConfig.
	//
	// 	- isAutoRecycle: Automatic recycle configuration. This key can be used only when CategoryName is set to DLCAutoRecycle.
	//
	// 	- priorityConfig: Priority configuration. This key can be used only when CategoryName is set to DLCPriorityConfig or DSWPriorityConfig.
	//
	// 	- quotaMaximumDuration: Maximum run time of DLC jobs for a quota. This key can be used only when CategoryName is set to QuotaMaximumDuration.
	//
	// 	- predefinedTags: Preset tags of the workspace. Created resources must include tags.
	//
	// example:
	//
	// tempStoragePath
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// The value of the configuration item.
	//
	// example:
	//
	// oss://***
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The tags of the configuration item.
	Labels []*UpdateConfigRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s UpdateConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigRequest) GetCategoryName() *string {
	return s.CategoryName
}

func (s *UpdateConfigRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *UpdateConfigRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *UpdateConfigRequest) GetLabels() []*UpdateConfigRequestLabels {
	return s.Labels
}

func (s *UpdateConfigRequest) SetCategoryName(v string) *UpdateConfigRequest {
	s.CategoryName = &v
	return s
}

func (s *UpdateConfigRequest) SetConfigKey(v string) *UpdateConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *UpdateConfigRequest) SetConfigValue(v string) *UpdateConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *UpdateConfigRequest) SetLabels(v []*UpdateConfigRequestLabels) *UpdateConfigRequest {
	s.Labels = v
	return s
}

func (s *UpdateConfigRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateConfigRequestLabels struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateConfigRequestLabels) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigRequestLabels) GoString() string {
	return s.String()
}

func (s *UpdateConfigRequestLabels) GetKey() *string {
	return s.Key
}

func (s *UpdateConfigRequestLabels) GetValue() *string {
	return s.Value
}

func (s *UpdateConfigRequestLabels) SetKey(v string) *UpdateConfigRequestLabels {
	s.Key = &v
	return s
}

func (s *UpdateConfigRequestLabels) SetValue(v string) *UpdateConfigRequestLabels {
	s.Value = &v
	return s
}

func (s *UpdateConfigRequestLabels) Validate() error {
	return dara.Validate(s)
}
