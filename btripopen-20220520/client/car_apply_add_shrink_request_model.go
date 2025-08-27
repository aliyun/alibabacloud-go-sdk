// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarApplyAddShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCause(v string) *CarApplyAddShrinkRequest
	GetCause() *string
	SetCity(v string) *CarApplyAddShrinkRequest
	GetCity() *string
	SetCityCodeSet(v string) *CarApplyAddShrinkRequest
	GetCityCodeSet() *string
	SetDate(v string) *CarApplyAddShrinkRequest
	GetDate() *string
	SetFinishedDate(v string) *CarApplyAddShrinkRequest
	GetFinishedDate() *string
	SetProjectCode(v string) *CarApplyAddShrinkRequest
	GetProjectCode() *string
	SetProjectName(v string) *CarApplyAddShrinkRequest
	GetProjectName() *string
	SetStatus(v int32) *CarApplyAddShrinkRequest
	GetStatus() *int32
	SetThirdPartApplyId(v string) *CarApplyAddShrinkRequest
	GetThirdPartApplyId() *string
	SetThirdPartCostCenterId(v string) *CarApplyAddShrinkRequest
	GetThirdPartCostCenterId() *string
	SetThirdPartInvoiceId(v string) *CarApplyAddShrinkRequest
	GetThirdPartInvoiceId() *string
	SetTimesTotal(v int32) *CarApplyAddShrinkRequest
	GetTimesTotal() *int32
	SetTimesType(v int32) *CarApplyAddShrinkRequest
	GetTimesType() *int32
	SetTimesUsed(v int32) *CarApplyAddShrinkRequest
	GetTimesUsed() *int32
	SetTitle(v string) *CarApplyAddShrinkRequest
	GetTitle() *string
	SetTravelerStandardShrink(v string) *CarApplyAddShrinkRequest
	GetTravelerStandardShrink() *string
	SetUserId(v string) *CarApplyAddShrinkRequest
	GetUserId() *string
}

type CarApplyAddShrinkRequest struct {
	// This parameter is required.
	Cause *string `json:"cause,omitempty" xml:"cause,omitempty"`
	// This parameter is required.
	City        *string `json:"city,omitempty" xml:"city,omitempty"`
	CityCodeSet *string `json:"city_code_set,omitempty" xml:"city_code_set,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-07-12 14:52:52
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// 2022-07-12 18:51:25
	FinishedDate *string `json:"finished_date,omitempty" xml:"finished_date,omitempty"`
	// example:
	//
	// project1413
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
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
	// IRGS1413
	ThirdPartApplyId *string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	// example:
	//
	// QA1411
	ThirdPartCostCenterId *string `json:"third_part_cost_center_id,omitempty" xml:"third_part_cost_center_id,omitempty"`
	// example:
	//
	// GA15131
	ThirdPartInvoiceId *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TimesTotal *int32 `json:"times_total,omitempty" xml:"times_total,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TimesType *int32 `json:"times_type,omitempty" xml:"times_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TimesUsed *int32 `json:"times_used,omitempty" xml:"times_used,omitempty"`
	// This parameter is required.
	Title                  *string `json:"title,omitempty" xml:"title,omitempty"`
	TravelerStandardShrink *string `json:"traveler_standard,omitempty" xml:"traveler_standard,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN1415614
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CarApplyAddShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CarApplyAddShrinkRequest) GoString() string {
	return s.String()
}

func (s *CarApplyAddShrinkRequest) GetCause() *string {
	return s.Cause
}

func (s *CarApplyAddShrinkRequest) GetCity() *string {
	return s.City
}

func (s *CarApplyAddShrinkRequest) GetCityCodeSet() *string {
	return s.CityCodeSet
}

func (s *CarApplyAddShrinkRequest) GetDate() *string {
	return s.Date
}

func (s *CarApplyAddShrinkRequest) GetFinishedDate() *string {
	return s.FinishedDate
}

func (s *CarApplyAddShrinkRequest) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *CarApplyAddShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CarApplyAddShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CarApplyAddShrinkRequest) GetThirdPartApplyId() *string {
	return s.ThirdPartApplyId
}

func (s *CarApplyAddShrinkRequest) GetThirdPartCostCenterId() *string {
	return s.ThirdPartCostCenterId
}

func (s *CarApplyAddShrinkRequest) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *CarApplyAddShrinkRequest) GetTimesTotal() *int32 {
	return s.TimesTotal
}

func (s *CarApplyAddShrinkRequest) GetTimesType() *int32 {
	return s.TimesType
}

func (s *CarApplyAddShrinkRequest) GetTimesUsed() *int32 {
	return s.TimesUsed
}

func (s *CarApplyAddShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CarApplyAddShrinkRequest) GetTravelerStandardShrink() *string {
	return s.TravelerStandardShrink
}

func (s *CarApplyAddShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *CarApplyAddShrinkRequest) SetCause(v string) *CarApplyAddShrinkRequest {
	s.Cause = &v
	return s
}

func (s *CarApplyAddShrinkRequest) SetCity(v string) *CarApplyAddShrinkRequest {
	s.City = &v
	return s
}

func (s *CarApplyAddShrinkRequest) SetCityCodeSet(v string) *CarApplyAddShrinkRequest {
	s.CityCodeSet = &v
	return s
}

func (s *CarApplyAddShrinkRequest) SetDate(v string) *CarApplyAddShrinkRequest {
	s.Date = &v
	return s
}

func (s *CarApplyAddShrinkRequest) SetFinishedDate(v string) *CarApplyAddShrinkRequest {
	s.FinishedDate = &v
	return s
}

func (s *CarApplyAddShrinkRequest) SetProjectCode(v string) *CarApplyAddShrinkRequest {
	s.ProjectCode = &v
	return s
}

func (s *CarApplyAddShrinkRequest) SetProjectName(v string) *CarApplyAddShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CarApplyAddShrinkRequest) SetStatus(v int32) *CarApplyAddShrinkRequest {
	s.Status = &v
	return s
}

func (s *CarApplyAddShrinkRequest) SetThirdPartApplyId(v string) *CarApplyAddShrinkRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *CarApplyAddShrinkRequest) SetThirdPartCostCenterId(v string) *CarApplyAddShrinkRequest {
	s.ThirdPartCostCenterId = &v
	return s
}

func (s *CarApplyAddShrinkRequest) SetThirdPartInvoiceId(v string) *CarApplyAddShrinkRequest {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *CarApplyAddShrinkRequest) SetTimesTotal(v int32) *CarApplyAddShrinkRequest {
	s.TimesTotal = &v
	return s
}

func (s *CarApplyAddShrinkRequest) SetTimesType(v int32) *CarApplyAddShrinkRequest {
	s.TimesType = &v
	return s
}

func (s *CarApplyAddShrinkRequest) SetTimesUsed(v int32) *CarApplyAddShrinkRequest {
	s.TimesUsed = &v
	return s
}

func (s *CarApplyAddShrinkRequest) SetTitle(v string) *CarApplyAddShrinkRequest {
	s.Title = &v
	return s
}

func (s *CarApplyAddShrinkRequest) SetTravelerStandardShrink(v string) *CarApplyAddShrinkRequest {
	s.TravelerStandardShrink = &v
	return s
}

func (s *CarApplyAddShrinkRequest) SetUserId(v string) *CarApplyAddShrinkRequest {
	s.UserId = &v
	return s
}

func (s *CarApplyAddShrinkRequest) Validate() error {
	return dara.Validate(s)
}
