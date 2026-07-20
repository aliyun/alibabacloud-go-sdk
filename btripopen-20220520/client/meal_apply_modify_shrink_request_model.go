// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyModifyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyUserShrink(v string) *MealApplyModifyShrinkRequest
	GetApplyUserShrink() *string
	SetCostCenterId(v int64) *MealApplyModifyShrinkRequest
	GetCostCenterId() *int64
	SetExtendField(v string) *MealApplyModifyShrinkRequest
	GetExtendField() *string
	SetInvoiceId(v int64) *MealApplyModifyShrinkRequest
	GetInvoiceId() *int64
	SetItineraryListShrink(v string) *MealApplyModifyShrinkRequest
	GetItineraryListShrink() *string
	SetMealAmount(v int64) *MealApplyModifyShrinkRequest
	GetMealAmount() *int64
	SetMealCause(v string) *MealApplyModifyShrinkRequest
	GetMealCause() *string
	SetProjectCode(v string) *MealApplyModifyShrinkRequest
	GetProjectCode() *string
	SetProjectTitle(v string) *MealApplyModifyShrinkRequest
	GetProjectTitle() *string
	SetStatus(v int32) *MealApplyModifyShrinkRequest
	GetStatus() *int32
	SetThirdPartApplyId(v string) *MealApplyModifyShrinkRequest
	GetThirdPartApplyId() *string
	SetThirdPartCostCenterId(v string) *MealApplyModifyShrinkRequest
	GetThirdPartCostCenterId() *string
	SetThirdPartInvoiceId(v string) *MealApplyModifyShrinkRequest
	GetThirdPartInvoiceId() *string
}

type MealApplyModifyShrinkRequest struct {
	// This parameter is required.
	ApplyUserShrink *string `json:"apply_user,omitempty" xml:"apply_user,omitempty"`
	// example:
	//
	// 23
	CostCenterId *int64 `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	// example:
	//
	// {"extend_key":"extend_value"}
	ExtendField *string `json:"extend_field,omitempty" xml:"extend_field,omitempty"`
	// example:
	//
	// 123
	InvoiceId           *int64  `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	ItineraryListShrink *string `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty"`
	// example:
	//
	// 1
	MealAmount *int64 `json:"meal_amount,omitempty" xml:"meal_amount,omitempty"`
	// example:
	//
	// 测试
	MealCause *string `json:"meal_cause,omitempty" xml:"meal_cause,omitempty"`
	// example:
	//
	// project123
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// example:
	//
	// 项目1
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1234
	ThirdPartApplyId *string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	// example:
	//
	// 1200F00010
	ThirdPartCostCenterId *string `json:"third_part_cost_center_id,omitempty" xml:"third_part_cost_center_id,omitempty"`
	// example:
	//
	// GA15131
	ThirdPartInvoiceId *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
}

func (s MealApplyModifyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MealApplyModifyShrinkRequest) GoString() string {
	return s.String()
}

func (s *MealApplyModifyShrinkRequest) GetApplyUserShrink() *string {
	return s.ApplyUserShrink
}

func (s *MealApplyModifyShrinkRequest) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *MealApplyModifyShrinkRequest) GetExtendField() *string {
	return s.ExtendField
}

func (s *MealApplyModifyShrinkRequest) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *MealApplyModifyShrinkRequest) GetItineraryListShrink() *string {
	return s.ItineraryListShrink
}

func (s *MealApplyModifyShrinkRequest) GetMealAmount() *int64 {
	return s.MealAmount
}

func (s *MealApplyModifyShrinkRequest) GetMealCause() *string {
	return s.MealCause
}

func (s *MealApplyModifyShrinkRequest) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *MealApplyModifyShrinkRequest) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *MealApplyModifyShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *MealApplyModifyShrinkRequest) GetThirdPartApplyId() *string {
	return s.ThirdPartApplyId
}

func (s *MealApplyModifyShrinkRequest) GetThirdPartCostCenterId() *string {
	return s.ThirdPartCostCenterId
}

func (s *MealApplyModifyShrinkRequest) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *MealApplyModifyShrinkRequest) SetApplyUserShrink(v string) *MealApplyModifyShrinkRequest {
	s.ApplyUserShrink = &v
	return s
}

func (s *MealApplyModifyShrinkRequest) SetCostCenterId(v int64) *MealApplyModifyShrinkRequest {
	s.CostCenterId = &v
	return s
}

func (s *MealApplyModifyShrinkRequest) SetExtendField(v string) *MealApplyModifyShrinkRequest {
	s.ExtendField = &v
	return s
}

func (s *MealApplyModifyShrinkRequest) SetInvoiceId(v int64) *MealApplyModifyShrinkRequest {
	s.InvoiceId = &v
	return s
}

func (s *MealApplyModifyShrinkRequest) SetItineraryListShrink(v string) *MealApplyModifyShrinkRequest {
	s.ItineraryListShrink = &v
	return s
}

func (s *MealApplyModifyShrinkRequest) SetMealAmount(v int64) *MealApplyModifyShrinkRequest {
	s.MealAmount = &v
	return s
}

func (s *MealApplyModifyShrinkRequest) SetMealCause(v string) *MealApplyModifyShrinkRequest {
	s.MealCause = &v
	return s
}

func (s *MealApplyModifyShrinkRequest) SetProjectCode(v string) *MealApplyModifyShrinkRequest {
	s.ProjectCode = &v
	return s
}

func (s *MealApplyModifyShrinkRequest) SetProjectTitle(v string) *MealApplyModifyShrinkRequest {
	s.ProjectTitle = &v
	return s
}

func (s *MealApplyModifyShrinkRequest) SetStatus(v int32) *MealApplyModifyShrinkRequest {
	s.Status = &v
	return s
}

func (s *MealApplyModifyShrinkRequest) SetThirdPartApplyId(v string) *MealApplyModifyShrinkRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *MealApplyModifyShrinkRequest) SetThirdPartCostCenterId(v string) *MealApplyModifyShrinkRequest {
	s.ThirdPartCostCenterId = &v
	return s
}

func (s *MealApplyModifyShrinkRequest) SetThirdPartInvoiceId(v string) *MealApplyModifyShrinkRequest {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *MealApplyModifyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
