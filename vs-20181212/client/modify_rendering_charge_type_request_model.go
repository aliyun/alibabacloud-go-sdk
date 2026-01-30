// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRenderingChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *ModifyRenderingChargeTypeRequest
	GetAutoRenew() *bool
	SetInstanceBillingCycle(v string) *ModifyRenderingChargeTypeRequest
	GetInstanceBillingCycle() *string
	SetInstanceChargeType(v string) *ModifyRenderingChargeTypeRequest
	GetInstanceChargeType() *string
	SetPeriod(v string) *ModifyRenderingChargeTypeRequest
	GetPeriod() *string
	SetRenderingInstanceId(v string) *ModifyRenderingChargeTypeRequest
	GetRenderingInstanceId() *string
}

type ModifyRenderingChargeTypeRequest struct {
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// Hour
	InstanceBillingCycle *string `json:"InstanceBillingCycle,omitempty" xml:"InstanceBillingCycle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// example:
	//
	// 1
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s ModifyRenderingChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRenderingChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyRenderingChargeTypeRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ModifyRenderingChargeTypeRequest) GetInstanceBillingCycle() *string {
	return s.InstanceBillingCycle
}

func (s *ModifyRenderingChargeTypeRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ModifyRenderingChargeTypeRequest) GetPeriod() *string {
	return s.Period
}

func (s *ModifyRenderingChargeTypeRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ModifyRenderingChargeTypeRequest) SetAutoRenew(v bool) *ModifyRenderingChargeTypeRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyRenderingChargeTypeRequest) SetInstanceBillingCycle(v string) *ModifyRenderingChargeTypeRequest {
	s.InstanceBillingCycle = &v
	return s
}

func (s *ModifyRenderingChargeTypeRequest) SetInstanceChargeType(v string) *ModifyRenderingChargeTypeRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ModifyRenderingChargeTypeRequest) SetPeriod(v string) *ModifyRenderingChargeTypeRequest {
	s.Period = &v
	return s
}

func (s *ModifyRenderingChargeTypeRequest) SetRenderingInstanceId(v string) *ModifyRenderingChargeTypeRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *ModifyRenderingChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
