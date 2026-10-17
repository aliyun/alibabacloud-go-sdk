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
	SetOptions(v string) *ListConfigsRequest
	GetOptions() *string
	SetVerbose(v string) *ListConfigsRequest
	GetVerbose() *string
}

type ListConfigsRequest struct {
	// The categorization of the configuration item. Valid values:
	//
	// - CommonResourceConfig: general resource configuration
	//
	// - DLCAutoRecycle: automatic DLC revoke
	//
	// - DLCPriorityConfig: DLC priority settings
	//
	// - DSWPriorityConfig: DSW priority settings
	//
	// - QuotaMaximumDuration: maximum runtime configuration of DLC nodes for a quota
	//
	// - CommonTagConfig: label settings
	//
	// example:
	//
	// CommonResourceConfig
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The keys of the configuration items. Valid values:
	//
	// - tempStoragePath: Temporary storage path. This key is allowed only when CategoryName is set to CommonResourceConfig.
	//
	// - isAutoRecycle: Automatic revoke configuration. This key is allowed only when CategoryName is set to DLCAutoRecycle.
	//
	// - priorityConfig: Priority configuration. This key is allowed only when CategoryName is set to DLCPriorityConfig or DSWPriorityConfig.
	//
	// - quotaMaximumDuration: Maximum runtime configuration of DLC nodes for a quota. This key is allowed only when CategoryName is set to QuotaMaximumDuration.
	//
	// - predefinedTags: Preset labels for the workspace. Resources created must have these labels.
	//
	// example:
	//
	// tempStoragePath
	ConfigKeys *string `json:"ConfigKeys,omitempty" xml:"ConfigKeys,omitempty"`
	// The labels used as filter conditions. Separate multiple conditions with commas. These conditions are combined by using the AND operator.
	//
	// example:
	//
	// key1=value1,key2=value2
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The additional options. Currently, only ThrowExceptionOnError is supported, which indicates that an error is thrown when an exception occurs.
	//
	// example:
	//
	// ThrowExceptionOnError
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// Specifies whether to display label information. Valid values:
	//
	// - true: Displays label information.
	//
	// - false: Does not display label information.
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

func (s *ListConfigsRequest) GetOptions() *string {
	return s.Options
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

func (s *ListConfigsRequest) SetOptions(v string) *ListConfigsRequest {
	s.Options = &v
	return s
}

func (s *ListConfigsRequest) SetVerbose(v string) *ListConfigsRequest {
	s.Verbose = &v
	return s
}

func (s *ListConfigsRequest) Validate() error {
	return dara.Validate(s)
}
