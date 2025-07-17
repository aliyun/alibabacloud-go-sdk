// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataBonreeSDKConfigModuleConfigVersionConfigsValue interface {
	dara.Model
	String() string
	GoString() string
	SetUseCustom(v bool) *DataBonreeSDKConfigModuleConfigVersionConfigsValue
	GetUseCustom() *bool
	SetCustomConfig(v map[string]*DataBonreeSDKConfigModuleConfigVersionConfigsValueCustomConfigValue) *DataBonreeSDKConfigModuleConfigVersionConfigsValue
	GetCustomConfig() map[string]*DataBonreeSDKConfigModuleConfigVersionConfigsValueCustomConfigValue
	SetDescription(v string) *DataBonreeSDKConfigModuleConfigVersionConfigsValue
	GetDescription() *string
	SetUpdateTime(v int64) *DataBonreeSDKConfigModuleConfigVersionConfigsValue
	GetUpdateTime() *int64
}

type DataBonreeSDKConfigModuleConfigVersionConfigsValue struct {
	// Indicates whether the custom configuration is used.
	//
	// example:
	//
	// true
	UseCustom *bool `json:"useCustom,omitempty" xml:"useCustom,omitempty"`
	// The custom configuration.
	CustomConfig map[string]*DataBonreeSDKConfigModuleConfigVersionConfigsValueCustomConfigValue `json:"customConfig,omitempty" xml:"customConfig,omitempty"`
	// The description of the version configuration.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The time when the version configuration was updated.
	//
	// example:
	//
	// 1721112372055
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s DataBonreeSDKConfigModuleConfigVersionConfigsValue) String() string {
	return dara.Prettify(s)
}

func (s DataBonreeSDKConfigModuleConfigVersionConfigsValue) GoString() string {
	return s.String()
}

func (s *DataBonreeSDKConfigModuleConfigVersionConfigsValue) GetUseCustom() *bool {
	return s.UseCustom
}

func (s *DataBonreeSDKConfigModuleConfigVersionConfigsValue) GetCustomConfig() map[string]*DataBonreeSDKConfigModuleConfigVersionConfigsValueCustomConfigValue {
	return s.CustomConfig
}

func (s *DataBonreeSDKConfigModuleConfigVersionConfigsValue) GetDescription() *string {
	return s.Description
}

func (s *DataBonreeSDKConfigModuleConfigVersionConfigsValue) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DataBonreeSDKConfigModuleConfigVersionConfigsValue) SetUseCustom(v bool) *DataBonreeSDKConfigModuleConfigVersionConfigsValue {
	s.UseCustom = &v
	return s
}

func (s *DataBonreeSDKConfigModuleConfigVersionConfigsValue) SetCustomConfig(v map[string]*DataBonreeSDKConfigModuleConfigVersionConfigsValueCustomConfigValue) *DataBonreeSDKConfigModuleConfigVersionConfigsValue {
	s.CustomConfig = v
	return s
}

func (s *DataBonreeSDKConfigModuleConfigVersionConfigsValue) SetDescription(v string) *DataBonreeSDKConfigModuleConfigVersionConfigsValue {
	s.Description = &v
	return s
}

func (s *DataBonreeSDKConfigModuleConfigVersionConfigsValue) SetUpdateTime(v int64) *DataBonreeSDKConfigModuleConfigVersionConfigsValue {
	s.UpdateTime = &v
	return s
}

func (s *DataBonreeSDKConfigModuleConfigVersionConfigsValue) Validate() error {
	return dara.Validate(s)
}
