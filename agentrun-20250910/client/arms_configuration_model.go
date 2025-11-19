// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArmsConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetArmsLicenseKey(v string) *ArmsConfiguration
	GetArmsLicenseKey() *string
	SetEnableArms(v bool) *ArmsConfiguration
	GetEnableArms() *bool
}

type ArmsConfiguration struct {
	// 应用实时监控服务（ARMS）的许可证密钥
	//
	// example:
	//
	// arms-license-key-123456
	ArmsLicenseKey *string `json:"armsLicenseKey,omitempty" xml:"armsLicenseKey,omitempty"`
	// 是否启用应用实时监控服务（ARMS）
	//
	// example:
	//
	// true
	EnableArms *bool `json:"enableArms,omitempty" xml:"enableArms,omitempty"`
}

func (s ArmsConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ArmsConfiguration) GoString() string {
	return s.String()
}

func (s *ArmsConfiguration) GetArmsLicenseKey() *string {
	return s.ArmsLicenseKey
}

func (s *ArmsConfiguration) GetEnableArms() *bool {
	return s.EnableArms
}

func (s *ArmsConfiguration) SetArmsLicenseKey(v string) *ArmsConfiguration {
	s.ArmsLicenseKey = &v
	return s
}

func (s *ArmsConfiguration) SetEnableArms(v bool) *ArmsConfiguration {
	s.EnableArms = &v
	return s
}

func (s *ArmsConfiguration) Validate() error {
	return dara.Validate(s)
}
