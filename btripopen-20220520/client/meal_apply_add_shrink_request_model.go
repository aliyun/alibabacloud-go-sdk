// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyAddShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyUserShrink(v string) *MealApplyAddShrinkRequest
	GetApplyUserShrink() *string
	SetCostCenterId(v int64) *MealApplyAddShrinkRequest
	GetCostCenterId() *int64
	SetInvoiceId(v int64) *MealApplyAddShrinkRequest
	GetInvoiceId() *int64
	SetItineraryListShrink(v string) *MealApplyAddShrinkRequest
	GetItineraryListShrink() *string
	SetMealAmount(v int64) *MealApplyAddShrinkRequest
	GetMealAmount() *int64
	SetMealCause(v string) *MealApplyAddShrinkRequest
	GetMealCause() *string
	SetProjectCode(v string) *MealApplyAddShrinkRequest
	GetProjectCode() *string
	SetProjectTitle(v string) *MealApplyAddShrinkRequest
	GetProjectTitle() *string
	SetStatus(v int32) *MealApplyAddShrinkRequest
	GetStatus() *int32
	SetThirdPartApplyId(v string) *MealApplyAddShrinkRequest
	GetThirdPartApplyId() *string
	SetThirdPartCostCenterId(v string) *MealApplyAddShrinkRequest
	GetThirdPartCostCenterId() *string
	SetThirdPartInvoiceId(v string) *MealApplyAddShrinkRequest
	GetThirdPartInvoiceId() *string
}

type MealApplyAddShrinkRequest struct {
	// This parameter is required.
	ApplyUserShrink *string `json:"apply_user,omitempty" xml:"apply_user,omitempty"`
	// example:
	//
	// 23
	CostCenterId *int64 `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	// example:
	//
	// 123
	InvoiceId *int64 `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// This parameter is required.
	ItineraryListShrink *string `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty"`
	// example:
	//
	// 1
	MealAmount *int64 `json:"meal_amount,omitempty" xml:"meal_amount,omitempty"`
	// This parameter is required.
	MealCause *string `json:"meal_cause,omitempty" xml:"meal_cause,omitempty"`
	// example:
	//
	// project123
	ProjectCode  *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
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

func (s MealApplyAddShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MealApplyAddShrinkRequest) GoString() string {
	return s.String()
}

func (s *MealApplyAddShrinkRequest) GetApplyUserShrink() *string {
	return s.ApplyUserShrink
}

func (s *MealApplyAddShrinkRequest) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *MealApplyAddShrinkRequest) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *MealApplyAddShrinkRequest) GetItineraryListShrink() *string {
	return s.ItineraryListShrink
}

func (s *MealApplyAddShrinkRequest) GetMealAmount() *int64 {
	return s.MealAmount
}

func (s *MealApplyAddShrinkRequest) GetMealCause() *string {
	return s.MealCause
}

func (s *MealApplyAddShrinkRequest) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *MealApplyAddShrinkRequest) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *MealApplyAddShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *MealApplyAddShrinkRequest) GetThirdPartApplyId() *string {
	return s.ThirdPartApplyId
}

func (s *MealApplyAddShrinkRequest) GetThirdPartCostCenterId() *string {
	return s.ThirdPartCostCenterId
}

func (s *MealApplyAddShrinkRequest) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *MealApplyAddShrinkRequest) SetApplyUserShrink(v string) *MealApplyAddShrinkRequest {
	s.ApplyUserShrink = &v
	return s
}

func (s *MealApplyAddShrinkRequest) SetCostCenterId(v int64) *MealApplyAddShrinkRequest {
	s.CostCenterId = &v
	return s
}

func (s *MealApplyAddShrinkRequest) SetInvoiceId(v int64) *MealApplyAddShrinkRequest {
	s.InvoiceId = &v
	return s
}

func (s *MealApplyAddShrinkRequest) SetItineraryListShrink(v string) *MealApplyAddShrinkRequest {
	s.ItineraryListShrink = &v
	return s
}

func (s *MealApplyAddShrinkRequest) SetMealAmount(v int64) *MealApplyAddShrinkRequest {
	s.MealAmount = &v
	return s
}

func (s *MealApplyAddShrinkRequest) SetMealCause(v string) *MealApplyAddShrinkRequest {
	s.MealCause = &v
	return s
}

func (s *MealApplyAddShrinkRequest) SetProjectCode(v string) *MealApplyAddShrinkRequest {
	s.ProjectCode = &v
	return s
}

func (s *MealApplyAddShrinkRequest) SetProjectTitle(v string) *MealApplyAddShrinkRequest {
	s.ProjectTitle = &v
	return s
}

func (s *MealApplyAddShrinkRequest) SetStatus(v int32) *MealApplyAddShrinkRequest {
	s.Status = &v
	return s
}

func (s *MealApplyAddShrinkRequest) SetThirdPartApplyId(v string) *MealApplyAddShrinkRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *MealApplyAddShrinkRequest) SetThirdPartCostCenterId(v string) *MealApplyAddShrinkRequest {
	s.ThirdPartCostCenterId = &v
	return s
}

func (s *MealApplyAddShrinkRequest) SetThirdPartInvoiceId(v string) *MealApplyAddShrinkRequest {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *MealApplyAddShrinkRequest) Validate() error {
	return dara.Validate(s)
}
