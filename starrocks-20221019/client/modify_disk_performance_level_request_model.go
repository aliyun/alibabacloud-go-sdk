// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskPerformanceLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyDiskPerformanceLevelRequest
	GetAutoPay() *bool
	SetInstanceId(v string) *ModifyDiskPerformanceLevelRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *ModifyDiskPerformanceLevelRequest
	GetNodeGroupId() *string
	SetPromotionOptionNo(v string) *ModifyDiskPerformanceLevelRequest
	GetPromotionOptionNo() *string
	SetTarget(v string) *ModifyDiskPerformanceLevelRequest
	GetTarget() *string
}

type ModifyDiskPerformanceLevelRequest struct {
	// Specifies whether to automatically purchase (pay for) all products specified in the Products parameter. Valid values:
	//
	// - true: Automatic payment.
	//
	// - false: No automatic payment.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The compute group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// The target disk performance level (PL).
	//
	// This parameter is required.
	//
	// example:
	//
	// pl2
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyDiskPerformanceLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskPerformanceLevelRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskPerformanceLevelRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyDiskPerformanceLevelRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDiskPerformanceLevelRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ModifyDiskPerformanceLevelRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *ModifyDiskPerformanceLevelRequest) GetTarget() *string {
	return s.Target
}

func (s *ModifyDiskPerformanceLevelRequest) SetAutoPay(v bool) *ModifyDiskPerformanceLevelRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyDiskPerformanceLevelRequest) SetInstanceId(v string) *ModifyDiskPerformanceLevelRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDiskPerformanceLevelRequest) SetNodeGroupId(v string) *ModifyDiskPerformanceLevelRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyDiskPerformanceLevelRequest) SetPromotionOptionNo(v string) *ModifyDiskPerformanceLevelRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *ModifyDiskPerformanceLevelRequest) SetTarget(v string) *ModifyDiskPerformanceLevelRequest {
	s.Target = &v
	return s
}

func (s *ModifyDiskPerformanceLevelRequest) Validate() error {
	return dara.Validate(s)
}
