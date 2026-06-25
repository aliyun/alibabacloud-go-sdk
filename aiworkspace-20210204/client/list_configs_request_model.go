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
	// - CommonResourceConfig: The common resource configuration.
	//
	// - DLCAutoRecycle: The automatic recycling configuration for DLC.
	//
	// - DLCPriorityConfig: The priority configuration for DLC.
	//
	// - DSWPriorityConfig: The priority configuration for DSW.
	//
	// - QuotaMaximumDuration: The configuration for the maximum runtime of a DLC task in a quota.
	//
	// - CommonTagConfig: The label configuration.
	//
	// example:
	//
	// CommonResourceConfig
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The key of the configuration item. The following keys are supported:
	//
	// - tempStoragePath: The path for temporary storage. This key is valid only when CategoryName is set to CommonResourceConfig.
	//
	// - isAutoRecycle: The automatic recycling configuration. This key is valid only when CategoryName is set to DLCAutoRecycle.
	//
	// - priorityConfig: The priority configuration. This key is valid only when CategoryName is set to DLCPriorityConfig or DSWPriorityConfig.
	//
	// - quotaMaximumDuration: The configuration for the maximum runtime of a DLC task in a quota. This key is valid only when CategoryName is set to QuotaMaximumDuration.
	//
	// - predefinedTags: The predefined labels for the workspace. Resources that you create must have these labels.
	//
	// example:
	//
	// tempStoragePath
	ConfigKeys *string `json:"ConfigKeys,omitempty" xml:"ConfigKeys,omitempty"`
	// The labels to use as filter conditions. Separate multiple labels with commas. A logical AND operation is performed on these labels.
	//
	// example:
	//
	// key1=value1,key2=value2
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// Specifies whether to return label information.
	//
	// - true: Returns label information.
	//
	// - false: Does not return label information.
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
