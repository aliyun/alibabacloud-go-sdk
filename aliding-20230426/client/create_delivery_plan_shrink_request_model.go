// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeliveryPlanShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentShrink(v string) *CreateDeliveryPlanShrinkRequest
	GetContentShrink() *string
	SetEndTime(v int64) *CreateDeliveryPlanShrinkRequest
	GetEndTime() *int64
	SetResId(v string) *CreateDeliveryPlanShrinkRequest
	GetResId() *string
	SetStartTime(v int64) *CreateDeliveryPlanShrinkRequest
	GetStartTime() *int64
	SetTenantContextShrink(v string) *CreateDeliveryPlanShrinkRequest
	GetTenantContextShrink() *string
	SetUserIdListShrink(v string) *CreateDeliveryPlanShrinkRequest
	GetUserIdListShrink() *string
}

type CreateDeliveryPlanShrinkRequest struct {
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1699265024987
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1028
	ResId *string `json:"ResId,omitempty" xml:"ResId,omitempty"`
	// example:
	//
	// 1699265024987
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	UserIdListShrink    *string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty"`
}

func (s CreateDeliveryPlanShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryPlanShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDeliveryPlanShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *CreateDeliveryPlanShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateDeliveryPlanShrinkRequest) GetResId() *string {
	return s.ResId
}

func (s *CreateDeliveryPlanShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateDeliveryPlanShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateDeliveryPlanShrinkRequest) GetUserIdListShrink() *string {
	return s.UserIdListShrink
}

func (s *CreateDeliveryPlanShrinkRequest) SetContentShrink(v string) *CreateDeliveryPlanShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *CreateDeliveryPlanShrinkRequest) SetEndTime(v int64) *CreateDeliveryPlanShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *CreateDeliveryPlanShrinkRequest) SetResId(v string) *CreateDeliveryPlanShrinkRequest {
	s.ResId = &v
	return s
}

func (s *CreateDeliveryPlanShrinkRequest) SetStartTime(v int64) *CreateDeliveryPlanShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *CreateDeliveryPlanShrinkRequest) SetTenantContextShrink(v string) *CreateDeliveryPlanShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateDeliveryPlanShrinkRequest) SetUserIdListShrink(v string) *CreateDeliveryPlanShrinkRequest {
	s.UserIdListShrink = &v
	return s
}

func (s *CreateDeliveryPlanShrinkRequest) Validate() error {
	return dara.Validate(s)
}
