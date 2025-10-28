// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListUserConfigsResponseBodyConfigs) *ListUserConfigsResponseBody
	GetConfigs() []*ListUserConfigsResponseBodyConfigs
	SetRequestId(v string) *ListUserConfigsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListUserConfigsResponseBody
	GetTotalCount() *int64
}

type ListUserConfigsResponseBody struct {
	// The configurations list.
	Configs []*ListUserConfigsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// sdjksdk-******-dsfds
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of items returned.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserConfigsResponseBody) GetConfigs() []*ListUserConfigsResponseBodyConfigs {
	return s.Configs
}

func (s *ListUserConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserConfigsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUserConfigsResponseBody) SetConfigs(v []*ListUserConfigsResponseBodyConfigs) *ListUserConfigsResponseBody {
	s.Configs = v
	return s
}

func (s *ListUserConfigsResponseBody) SetRequestId(v string) *ListUserConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserConfigsResponseBody) SetTotalCount(v int64) *ListUserConfigsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserConfigsResponseBody) Validate() error {
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

type ListUserConfigsResponseBodyConfigs struct {
	// The category. Currently, only DataPrivacyConfig is supported.
	//
	// example:
	//
	// DataPrivacyConfig
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The key of the configuration item.
	//
	// example:
	//
	// customizePAIAssumedRole
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// The value of the configuration item.
	//
	// example:
	//
	// role
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The scope. Currently, subUser and owner are supported.
	//
	// example:
	//
	// subUser
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ListUserConfigsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListUserConfigsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListUserConfigsResponseBodyConfigs) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ListUserConfigsResponseBodyConfigs) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *ListUserConfigsResponseBodyConfigs) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *ListUserConfigsResponseBodyConfigs) GetScope() *string {
	return s.Scope
}

func (s *ListUserConfigsResponseBodyConfigs) SetCategoryName(v string) *ListUserConfigsResponseBodyConfigs {
	s.CategoryName = &v
	return s
}

func (s *ListUserConfigsResponseBodyConfigs) SetConfigKey(v string) *ListUserConfigsResponseBodyConfigs {
	s.ConfigKey = &v
	return s
}

func (s *ListUserConfigsResponseBodyConfigs) SetConfigValue(v string) *ListUserConfigsResponseBodyConfigs {
	s.ConfigValue = &v
	return s
}

func (s *ListUserConfigsResponseBodyConfigs) SetScope(v string) *ListUserConfigsResponseBodyConfigs {
	s.Scope = &v
	return s
}

func (s *ListUserConfigsResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
