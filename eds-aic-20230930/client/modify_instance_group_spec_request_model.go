// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceGroupSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyInstanceGroupSpecRequest
	GetAutoPay() *bool
	SetInstanceGroupIds(v []*string) *ModifyInstanceGroupSpecRequest
	GetInstanceGroupIds() []*string
	SetInstanceGroupSpec(v string) *ModifyInstanceGroupSpecRequest
	GetInstanceGroupSpec() *string
	SetPromotionId(v string) *ModifyInstanceGroupSpecRequest
	GetPromotionId() *string
}

type ModifyInstanceGroupSpecRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **true**: Automatic payment is enabled. Make sure that your account balance is sufficient.
	//
	// - **false*	- (default): Only generates an order without deducting fees.
	//
	//
	//
	//
	// > If your payment method balance is insufficient, set this parameter to false. An unpaid order is generated, and you can log on to the Cloud Phone console to complete the payment.
	//
	// >
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The list of instance group IDs.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The instance group specification. You can call [DescribeSpec](~~DescribeSpec~~) to query the specifications available for purchase for cloud phones.
	//
	// This parameter is required.
	//
	// example:
	//
	// acp.basic.small
	InstanceGroupSpec *string `json:"InstanceGroupSpec,omitempty" xml:"InstanceGroupSpec,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 50003308011****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
}

func (s ModifyInstanceGroupSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceGroupSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceGroupSpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyInstanceGroupSpecRequest) GetInstanceGroupIds() []*string {
	return s.InstanceGroupIds
}

func (s *ModifyInstanceGroupSpecRequest) GetInstanceGroupSpec() *string {
	return s.InstanceGroupSpec
}

func (s *ModifyInstanceGroupSpecRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *ModifyInstanceGroupSpecRequest) SetAutoPay(v bool) *ModifyInstanceGroupSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyInstanceGroupSpecRequest) SetInstanceGroupIds(v []*string) *ModifyInstanceGroupSpecRequest {
	s.InstanceGroupIds = v
	return s
}

func (s *ModifyInstanceGroupSpecRequest) SetInstanceGroupSpec(v string) *ModifyInstanceGroupSpecRequest {
	s.InstanceGroupSpec = &v
	return s
}

func (s *ModifyInstanceGroupSpecRequest) SetPromotionId(v string) *ModifyInstanceGroupSpecRequest {
	s.PromotionId = &v
	return s
}

func (s *ModifyInstanceGroupSpecRequest) Validate() error {
	return dara.Validate(s)
}
