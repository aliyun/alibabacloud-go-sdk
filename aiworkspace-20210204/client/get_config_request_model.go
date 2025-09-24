// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryName(v string) *GetConfigRequest
	GetCategoryName() *string
	SetConfigKey(v string) *GetConfigRequest
	GetConfigKey() *string
	SetVerbose(v string) *GetConfigRequest
	GetVerbose() *string
}

type GetConfigRequest struct {
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
	// 	- predefinedTags: Predefined tags of the workspace. Created resources must include tags.
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
	Verbose *string `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRequest) GoString() string {
	return s.String()
}

func (s *GetConfigRequest) GetCategoryName() *string {
	return s.CategoryName
}

func (s *GetConfigRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *GetConfigRequest) GetVerbose() *string {
	return s.Verbose
}

func (s *GetConfigRequest) SetCategoryName(v string) *GetConfigRequest {
	s.CategoryName = &v
	return s
}

func (s *GetConfigRequest) SetConfigKey(v string) *GetConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *GetConfigRequest) SetVerbose(v string) *GetConfigRequest {
	s.Verbose = &v
	return s
}

func (s *GetConfigRequest) Validate() error {
	return dara.Validate(s)
}
