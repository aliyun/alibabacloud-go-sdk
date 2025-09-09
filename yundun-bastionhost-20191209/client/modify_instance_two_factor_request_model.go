// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceTwoFactorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableTwoFactor(v string) *ModifyInstanceTwoFactorRequest
	GetEnableTwoFactor() *string
	SetInstanceId(v string) *ModifyInstanceTwoFactorRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyInstanceTwoFactorRequest
	GetRegionId() *string
	SetSkipTwoFactorTime(v string) *ModifyInstanceTwoFactorRequest
	GetSkipTwoFactorTime() *string
	SetTwoFactorMethods(v string) *ModifyInstanceTwoFactorRequest
	GetTwoFactorMethods() *string
}

type ModifyInstanceTwoFactorRequest struct {
	// Specifies whether to enable two-factor authentication. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	EnableTwoFactor *string `json:"EnableTwoFactor,omitempty" xml:"EnableTwoFactor,omitempty"`
	// The ID of the bastion host.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The duration within which two-factor authentication is not required after a user passes two-factor authentication. Valid values: 0 to 168. Unit: hours. If you set this parameter to 0, the user must pass two-factor authentication every time the user logs on to the bastion host.
	//
	// example:
	//
	// 1
	SkipTwoFactorTime *string `json:"SkipTwoFactorTime,omitempty" xml:"SkipTwoFactorTime,omitempty"`
	// The method used to send a verification code for two-factor authentication. If EnableTwoFactor is set to true, you must specify at least one method. Valid values:
	//
	// 	- **sms:*	- text message.
	//
	// 	- **email**: email.
	//
	// 	- **dingtalk**: notice in DingTalk.
	//
	// 	- **totp**: one-time password (OTP) token.
	//
	// 	- **gmusbkey**: SM-based USB key.
	//
	// example:
	//
	// ["sms"]
	TwoFactorMethods *string `json:"TwoFactorMethods,omitempty" xml:"TwoFactorMethods,omitempty"`
}

func (s ModifyInstanceTwoFactorRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceTwoFactorRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTwoFactorRequest) GetEnableTwoFactor() *string {
	return s.EnableTwoFactor
}

func (s *ModifyInstanceTwoFactorRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceTwoFactorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceTwoFactorRequest) GetSkipTwoFactorTime() *string {
	return s.SkipTwoFactorTime
}

func (s *ModifyInstanceTwoFactorRequest) GetTwoFactorMethods() *string {
	return s.TwoFactorMethods
}

func (s *ModifyInstanceTwoFactorRequest) SetEnableTwoFactor(v string) *ModifyInstanceTwoFactorRequest {
	s.EnableTwoFactor = &v
	return s
}

func (s *ModifyInstanceTwoFactorRequest) SetInstanceId(v string) *ModifyInstanceTwoFactorRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceTwoFactorRequest) SetRegionId(v string) *ModifyInstanceTwoFactorRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceTwoFactorRequest) SetSkipTwoFactorTime(v string) *ModifyInstanceTwoFactorRequest {
	s.SkipTwoFactorTime = &v
	return s
}

func (s *ModifyInstanceTwoFactorRequest) SetTwoFactorMethods(v string) *ModifyInstanceTwoFactorRequest {
	s.TwoFactorMethods = &v
	return s
}

func (s *ModifyInstanceTwoFactorRequest) Validate() error {
	return dara.Validate(s)
}
