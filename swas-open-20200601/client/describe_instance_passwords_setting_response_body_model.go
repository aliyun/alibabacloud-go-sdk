// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancePasswordsSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstancePasswordSetting(v bool) *DescribeInstancePasswordsSettingResponseBody
	GetInstancePasswordSetting() *bool
	SetRequestId(v string) *DescribeInstancePasswordsSettingResponseBody
	GetRequestId() *string
	SetVncPasswordSetting(v bool) *DescribeInstancePasswordsSettingResponseBody
	GetVncPasswordSetting() *bool
}

type DescribeInstancePasswordsSettingResponseBody struct {
	// Indicates whether a logon password is set for the simple application server.
	//
	// example:
	//
	// true
	InstancePasswordSetting *bool `json:"InstancePasswordSetting,omitempty" xml:"InstancePasswordSetting,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether a VNC connection password is set for the simple application server.
	//
	// example:
	//
	// true
	VncPasswordSetting *bool `json:"VncPasswordSetting,omitempty" xml:"VncPasswordSetting,omitempty"`
}

func (s DescribeInstancePasswordsSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancePasswordsSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancePasswordsSettingResponseBody) GetInstancePasswordSetting() *bool {
	return s.InstancePasswordSetting
}

func (s *DescribeInstancePasswordsSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancePasswordsSettingResponseBody) GetVncPasswordSetting() *bool {
	return s.VncPasswordSetting
}

func (s *DescribeInstancePasswordsSettingResponseBody) SetInstancePasswordSetting(v bool) *DescribeInstancePasswordsSettingResponseBody {
	s.InstancePasswordSetting = &v
	return s
}

func (s *DescribeInstancePasswordsSettingResponseBody) SetRequestId(v string) *DescribeInstancePasswordsSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancePasswordsSettingResponseBody) SetVncPasswordSetting(v bool) *DescribeInstancePasswordsSettingResponseBody {
	s.VncPasswordSetting = &v
	return s
}

func (s *DescribeInstancePasswordsSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
