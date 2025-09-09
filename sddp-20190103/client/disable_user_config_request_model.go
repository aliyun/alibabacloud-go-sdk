// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableUserConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableUserConfigRequest
	GetCode() *string
	SetFeatureType(v int32) *DisableUserConfigRequest
	GetFeatureType() *int32
	SetLang(v string) *DisableUserConfigRequest
	GetLang() *string
}

type DisableUserConfigRequest struct {
	// The code of the configuration item. You can call the [DescribeConfigs](~~DescribeConfigs~~) operation to obtain the code of the configuration item.
	//
	// example:
	//
	// access_failed_cnt
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh_cn**: Chinese (default)
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DisableUserConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableUserConfigRequest) GoString() string {
	return s.String()
}

func (s *DisableUserConfigRequest) GetCode() *string {
	return s.Code
}

func (s *DisableUserConfigRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DisableUserConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DisableUserConfigRequest) SetCode(v string) *DisableUserConfigRequest {
	s.Code = &v
	return s
}

func (s *DisableUserConfigRequest) SetFeatureType(v int32) *DisableUserConfigRequest {
	s.FeatureType = &v
	return s
}

func (s *DisableUserConfigRequest) SetLang(v string) *DisableUserConfigRequest {
	s.Lang = &v
	return s
}

func (s *DisableUserConfigRequest) Validate() error {
	return dara.Validate(s)
}
