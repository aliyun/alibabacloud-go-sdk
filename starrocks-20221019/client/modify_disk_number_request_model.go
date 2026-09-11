// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyDiskNumberRequest
	GetAutoPay() *bool
	SetFastMode(v bool) *ModifyDiskNumberRequest
	GetFastMode() *bool
	SetInstanceId(v string) *ModifyDiskNumberRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *ModifyDiskNumberRequest
	GetNodeGroupId() *string
	SetPromotionOptionNo(v string) *ModifyDiskNumberRequest
	GetPromotionOptionNo() *string
	SetTarget(v int32) *ModifyDiskNumberRequest
	GetTarget() *int32
}

type ModifyDiskNumberRequest struct {
	// Specifies whether to automatically purchase (pay for) all products specified in the Products parameter.
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
	// Specifies whether to use the fast restart mode. Default value: false.
	//
	// - true: Restarts compute nodes in fast restart mode. Compute nodes are restarted in multiple batches. Nodes within a batch are restarted in parallel, and batches execute sequentially.
	//
	// - false: Restarts compute nodes in rolling restart mode.
	//
	// example:
	//
	// true
	FastMode *bool `json:"FastMode,omitempty" xml:"FastMode,omitempty"`
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
	// The target number of disks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyDiskNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskNumberRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskNumberRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyDiskNumberRequest) GetFastMode() *bool {
	return s.FastMode
}

func (s *ModifyDiskNumberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDiskNumberRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ModifyDiskNumberRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *ModifyDiskNumberRequest) GetTarget() *int32 {
	return s.Target
}

func (s *ModifyDiskNumberRequest) SetAutoPay(v bool) *ModifyDiskNumberRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyDiskNumberRequest) SetFastMode(v bool) *ModifyDiskNumberRequest {
	s.FastMode = &v
	return s
}

func (s *ModifyDiskNumberRequest) SetInstanceId(v string) *ModifyDiskNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDiskNumberRequest) SetNodeGroupId(v string) *ModifyDiskNumberRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyDiskNumberRequest) SetPromotionOptionNo(v string) *ModifyDiskNumberRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *ModifyDiskNumberRequest) SetTarget(v int32) *ModifyDiskNumberRequest {
	s.Target = &v
	return s
}

func (s *ModifyDiskNumberRequest) Validate() error {
	return dara.Validate(s)
}
