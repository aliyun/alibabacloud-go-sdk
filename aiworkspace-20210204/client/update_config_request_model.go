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
	// The classification of the configuration item. The following classifications are supported:
	//
	// - CommonResourceConfig: The common resource configuration.
	//
	// - DLCAutoRecycle: The DLC auto-recycle configuration.
	//
	// - DLCPriorityConfig: The DLC priority settings.
	//
	// - DSWPriorityConfig: The DSW priority settings.
	//
	// - QuotaMaximumDuration: The maximum runtime of a DLC task for a quota.
	//
	// - CommonTagConfig: The tag settings.
	//
	// example:
	//
	// CommonResourceConfig
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The key of the configuration item. The following keys are supported:
	//
	// - tempStoragePath: The path for temporary storage. This key is valid only when CategoryName is set to CommonResourceConfig.
	//
	// - isAutoRecycle: The auto-recycle configuration. This key is valid only when CategoryName is set to DLCAutoRecycle.
	//
	// - priorityConfig: The priority configuration. This key is valid only when CategoryName is set to DLCPriorityConfig or DSWPriorityConfig.
	//
	// - quotaMaximumDuration: The maximum runtime of a DLC task for a quota. This key is valid only when CategoryName is set to QuotaMaximumDuration.
	//
	// - predefinedTags: The predefined tags for the workspace. Created resources must have these tags.
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
	// The list of labels for the configuration item.
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

type UpdateConfigRequestLabels struct {
	// The key of the label.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the label.
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
