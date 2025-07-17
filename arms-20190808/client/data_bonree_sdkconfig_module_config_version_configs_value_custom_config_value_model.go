// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataBonreeSDKConfigModuleConfigVersionConfigsValueCustomConfigValue interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *DataBonreeSDKConfigModuleConfigVersionConfigsValueCustomConfigValue
	GetEnable() *bool
}

type DataBonreeSDKConfigModuleConfigVersionConfigsValueCustomConfigValue struct {
	// Indicates whether the configuration is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s DataBonreeSDKConfigModuleConfigVersionConfigsValueCustomConfigValue) String() string {
	return dara.Prettify(s)
}

func (s DataBonreeSDKConfigModuleConfigVersionConfigsValueCustomConfigValue) GoString() string {
	return s.String()
}

func (s *DataBonreeSDKConfigModuleConfigVersionConfigsValueCustomConfigValue) GetEnable() *bool {
	return s.Enable
}

func (s *DataBonreeSDKConfigModuleConfigVersionConfigsValueCustomConfigValue) SetEnable(v bool) *DataBonreeSDKConfigModuleConfigVersionConfigsValueCustomConfigValue {
	s.Enable = &v
	return s
}

func (s *DataBonreeSDKConfigModuleConfigVersionConfigsValueCustomConfigValue) Validate() error {
	return dara.Validate(s)
}
