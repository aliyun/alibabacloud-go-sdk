// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryName(v string) *ListConfigsRequest
	GetCategoryName() *string
	SetConfigKeys(v string) *ListConfigsRequest
	GetConfigKeys() *string
	SetLabels(v string) *ListConfigsRequest
	GetLabels() *string
	SetVerbose(v string) *ListConfigsRequest
	GetVerbose() *string
}

type ListConfigsRequest struct {
	// The category of the configuration item. The following categories are supported:
	//
	// - CommonResourceConfig: common resource configuration.
	//
	// - DLCAutoRecycle: DLC automatic recycling.
	//
	// - DLCPriorityConfig: DLC priority settings.
	//
	// - DSWPriorityConfig: DSW priority settings.
	//
	// - QuotaMaximumDuration: maximum runtime duration configuration for DLC jobs in a quota.
	//
	// - CommonTagConfig: tag settings.
	//
	// example:
	//
	// CommonResourceConfig
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The keys of the configuration items. The following keys are supported:
	//
	// - tempStoragePath: the temporary storage path. This ConfigKey can be used only when CategoryName is set to CommonResourceConfig.
	//
	// - isAutoRecycle: the automatic recycling configuration. This ConfigKey can be used only when CategoryName is set to DLCAutoRecycle.
	//
	// - priorityConfig: the priority configuration. This ConfigKey can be used only when CategoryName is set to DLCPriorityConfig or DSWPriorityConfig.
	//
	// - quotaMaximumDuration: the maximum runtime duration configuration for DLC jobs in a quota. This ConfigKey can be used only when CategoryName is set to QuotaMaximumDuration.
	//
	// - predefinedTags: the preset tags for the workspace. Resources that are created must include these tags.
	//
	// example:
	//
	// tempStoragePath
	ConfigKeys *string `json:"ConfigKeys,omitempty" xml:"ConfigKeys,omitempty"`
	// The labels used as filter conditions. Separate multiple conditions with commas. These conditions are evaluated using a logical AND.
	//
	// example:
	//
	// key1=value1,key2=value2
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// Specifies whether to display label information. Valid values:
	//
	// - true: Display label information.
	//
	// - false: Do not display label information.
	//
	// example:
	//
	// true
	Verbose *string `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s ListConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListConfigsRequest) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ListConfigsRequest) GetConfigKeys() *string {
	return s.ConfigKeys
}

func (s *ListConfigsRequest) GetLabels() *string {
	return s.Labels
}

func (s *ListConfigsRequest) GetVerbose() *string {
	return s.Verbose
}

func (s *ListConfigsRequest) SetCategoryName(v string) *ListConfigsRequest {
	s.CategoryName = &v
	return s
}

func (s *ListConfigsRequest) SetConfigKeys(v string) *ListConfigsRequest {
	s.ConfigKeys = &v
	return s
}

func (s *ListConfigsRequest) SetLabels(v string) *ListConfigsRequest {
	s.Labels = &v
	return s
}

func (s *ListConfigsRequest) SetVerbose(v string) *ListConfigsRequest {
	s.Verbose = &v
	return s
}

func (s *ListConfigsRequest) Validate() error {
	return dara.Validate(s)
}
