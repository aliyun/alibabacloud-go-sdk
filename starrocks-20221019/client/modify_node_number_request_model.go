// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyNodeNumberRequest
	GetAutoPay() *bool
	SetInstanceId(v string) *ModifyNodeNumberRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *ModifyNodeNumberRequest
	GetNodeGroupId() *string
	SetParallelism(v int32) *ModifyNodeNumberRequest
	GetParallelism() *int32
	SetPromotionOptionNo(v string) *ModifyNodeNumberRequest
	GetPromotionOptionNo() *string
	SetTarget(v int32) *ModifyNodeNumberRequest
	GetTarget() *int32
	SetTerminationGracePeriodSeconds(v int32) *ModifyNodeNumberRequest
	GetTerminationGracePeriodSeconds() *int32
}

type ModifyNodeNumberRequest struct {
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
	// true
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
	// The decommission concurrency for BE scale-in scenarios in compute-storage coupled mode. Default value: 1.
	//
	// example:
	//
	// 1
	Parallelism *int32 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// The target number of nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
	// The wait time for running tasks to complete before dropping nodes during CN scale-in scenarios in compute-storage decoupled mode.
	//
	// example:
	//
	// 60
	TerminationGracePeriodSeconds *int32 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
}

func (s ModifyNodeNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeNumberRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodeNumberRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyNodeNumberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyNodeNumberRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ModifyNodeNumberRequest) GetParallelism() *int32 {
	return s.Parallelism
}

func (s *ModifyNodeNumberRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *ModifyNodeNumberRequest) GetTarget() *int32 {
	return s.Target
}

func (s *ModifyNodeNumberRequest) GetTerminationGracePeriodSeconds() *int32 {
	return s.TerminationGracePeriodSeconds
}

func (s *ModifyNodeNumberRequest) SetAutoPay(v bool) *ModifyNodeNumberRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyNodeNumberRequest) SetInstanceId(v string) *ModifyNodeNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyNodeNumberRequest) SetNodeGroupId(v string) *ModifyNodeNumberRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyNodeNumberRequest) SetParallelism(v int32) *ModifyNodeNumberRequest {
	s.Parallelism = &v
	return s
}

func (s *ModifyNodeNumberRequest) SetPromotionOptionNo(v string) *ModifyNodeNumberRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *ModifyNodeNumberRequest) SetTarget(v int32) *ModifyNodeNumberRequest {
	s.Target = &v
	return s
}

func (s *ModifyNodeNumberRequest) SetTerminationGracePeriodSeconds(v int32) *ModifyNodeNumberRequest {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *ModifyNodeNumberRequest) Validate() error {
	return dara.Validate(s)
}
