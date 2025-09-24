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
	// 	- priorityConfig: Priority configuration. This key can be used only when CategoryName is set to DLCPriorityConfig or DSWPriorityConfig.
	//
	// 	- quotaMaximumDuration: Maximum run time of DLC jobs for a quota. This key can be used only when CategoryName is set to QuotaMaximumDuration.
	//
	// 	- predefinedTags: The predefined tags of the workspace. All created resources must have tags
	//
	// example:
	//
	// tempStoragePath
	ConfigKeys *string `json:"ConfigKeys,omitempty" xml:"ConfigKeys,omitempty"`
	// The tags used as filter conditions. Separate multiple tags with commas (,). These conditions are in an AND relationship.
	//
	// example:
	//
	// key1=value1,key2=value2
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// Specifies whether to show the tag information.
	//
	// 	- true
	//
	// 	- false
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
