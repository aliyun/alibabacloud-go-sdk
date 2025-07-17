// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataBonreeSDKConfigModuleConfigDefaultConfigValue interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *DataBonreeSDKConfigModuleConfigDefaultConfigValue
	GetEnable() *bool
}

type DataBonreeSDKConfigModuleConfigDefaultConfigValue struct {
	// Indicates whether the configuration is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s DataBonreeSDKConfigModuleConfigDefaultConfigValue) String() string {
	return dara.Prettify(s)
}

func (s DataBonreeSDKConfigModuleConfigDefaultConfigValue) GoString() string {
	return s.String()
}

func (s *DataBonreeSDKConfigModuleConfigDefaultConfigValue) GetEnable() *bool {
	return s.Enable
}

func (s *DataBonreeSDKConfigModuleConfigDefaultConfigValue) SetEnable(v bool) *DataBonreeSDKConfigModuleConfigDefaultConfigValue {
	s.Enable = &v
	return s
}

func (s *DataBonreeSDKConfigModuleConfigDefaultConfigValue) Validate() error {
	return dara.Validate(s)
}
