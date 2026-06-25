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
	// The classification of the configuration item. The following classifications are supported:
	//
	// - CommonResourceConfig: common resource configurations
	//
	// - DLCAutoRecycle: automatic DLC resource recycling
	//
	// - DLCPriorityConfig: DLC priority settings
	//
	// - DSWPriorityConfig: DSW priority settings
	//
	// - QuotaMaximumDuration: the maximum runtime of a DLC task for a quota
	//
	// - CommonTagConfig: tag settings
	//
	// example:
	//
	// CommonResourceConfig
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The key of the configuration item. The following keys are supported:
	//
	// - tempStoragePath: the temporary storage path. This key applies only when CategoryName is set to CommonResourceConfig.
	//
	// - isAutoRecycle: the automatic recycling configuration. This key applies only when CategoryName is set to DLCAutoRecycle.
	//
	// - priorityConfig: the priority configuration. This key applies only when CategoryName is set to DLCPriorityConfig or DSWPriorityConfig.
	//
	// - quotaMaximumDuration: the maximum runtime of a DLC task for a quota. This key applies only when CategoryName is set to QuotaMaximumDuration.
	//
	// - predefinedTags: the predefined tags for the workspace. Resources that you create must have tags.
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
