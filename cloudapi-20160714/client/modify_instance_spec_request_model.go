// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyInstanceSpecRequest
	GetAutoPay() *bool
	SetInstanceId(v string) *ModifyInstanceSpecRequest
	GetInstanceId() *string
	SetInstanceSpec(v string) *ModifyInstanceSpecRequest
	GetInstanceSpec() *string
	SetModifyAction(v string) *ModifyInstanceSpecRequest
	GetModifyAction() *string
	SetSkipWaitSwitch(v bool) *ModifyInstanceSpecRequest
	GetSkipWaitSwitch() *bool
	SetToken(v string) *ModifyInstanceSpecRequest
	GetToken() *string
}

type ModifyInstanceSpecRequest struct {
	// Specifies whether payment is automatically made during renewal. Valid values:
	//
	// 	- **True**: Automatic payment is enabled. Make sure that your Alibaba Cloud account has adequate balance.
	//
	// 	- **False**: Automatic payment is disabled. You have to manually pay in the console. Log on to the console. In the upper-right corner, choose **Expenses > User Center**. In the left-side navigation pane, click **Orders**. On the page that appears, find your order and complete the payment.
	//
	// Default value: **False**.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// apigateway-cn-v6419k43xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The specifications of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// api.s1.small
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// Specifies whether to upgrade or downgrade the instance. Valid values:
	//
	// 	- **UPGRADE**
	//
	// 	- **DOWNGRADE**
	//
	// Default value: **UPGRADE**.
	//
	// example:
	//
	// UPGRADE
	ModifyAction *string `json:"ModifyAction,omitempty" xml:"ModifyAction,omitempty"`
	// Specifies whether to skip the Waiting for Traffic Switchover state. During the upgrade or downgrade, a new outbound IP address may be added to the API Gateway instance. The Waiting for Traffic Switchover state is used to remind users of adding the new outbound IP address to the whitelist. If you set the SkipWaitSwitch parameter to true, the instance does not enter the Waiting for Traffic Switchover state when a new outbound IP address is available. Instead, the system sends internal messages to the user.
	//
	// example:
	//
	// false
	SkipWaitSwitch *bool `json:"SkipWaitSwitch,omitempty" xml:"SkipWaitSwitch,omitempty"`
	// The password.
	//
	// This parameter is required.
	//
	// example:
	//
	// b5845042-2f2f-4e96-bd5c-36c6e5c2a68c
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s ModifyInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyInstanceSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceSpecRequest) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *ModifyInstanceSpecRequest) GetModifyAction() *string {
	return s.ModifyAction
}

func (s *ModifyInstanceSpecRequest) GetSkipWaitSwitch() *bool {
	return s.SkipWaitSwitch
}

func (s *ModifyInstanceSpecRequest) GetToken() *string {
	return s.Token
}

func (s *ModifyInstanceSpecRequest) SetAutoPay(v bool) *ModifyInstanceSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetInstanceId(v string) *ModifyInstanceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetInstanceSpec(v string) *ModifyInstanceSpecRequest {
	s.InstanceSpec = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetModifyAction(v string) *ModifyInstanceSpecRequest {
	s.ModifyAction = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetSkipWaitSwitch(v bool) *ModifyInstanceSpecRequest {
	s.SkipWaitSwitch = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetToken(v string) *ModifyInstanceSpecRequest {
	s.Token = &v
	return s
}

func (s *ModifyInstanceSpecRequest) Validate() error {
	return dara.Validate(s)
}
