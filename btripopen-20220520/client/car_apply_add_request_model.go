// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarApplyAddRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCause(v string) *CarApplyAddRequest
	GetCause() *string
	SetCity(v string) *CarApplyAddRequest
	GetCity() *string
	SetCityCodeSet(v string) *CarApplyAddRequest
	GetCityCodeSet() *string
	SetDate(v string) *CarApplyAddRequest
	GetDate() *string
	SetFinishedDate(v string) *CarApplyAddRequest
	GetFinishedDate() *string
	SetProjectCode(v string) *CarApplyAddRequest
	GetProjectCode() *string
	SetProjectName(v string) *CarApplyAddRequest
	GetProjectName() *string
	SetStatus(v int32) *CarApplyAddRequest
	GetStatus() *int32
	SetThirdPartApplyId(v string) *CarApplyAddRequest
	GetThirdPartApplyId() *string
	SetThirdPartCostCenterId(v string) *CarApplyAddRequest
	GetThirdPartCostCenterId() *string
	SetThirdPartInvoiceId(v string) *CarApplyAddRequest
	GetThirdPartInvoiceId() *string
	SetTimesTotal(v int32) *CarApplyAddRequest
	GetTimesTotal() *int32
	SetTimesType(v int32) *CarApplyAddRequest
	GetTimesType() *int32
	SetTimesUsed(v int32) *CarApplyAddRequest
	GetTimesUsed() *int32
	SetTitle(v string) *CarApplyAddRequest
	GetTitle() *string
	SetTravelerStandard(v []*CarApplyAddRequestTravelerStandard) *CarApplyAddRequest
	GetTravelerStandard() []*CarApplyAddRequestTravelerStandard
	SetUserId(v string) *CarApplyAddRequest
	GetUserId() *string
}

type CarApplyAddRequest struct {
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
	Title            *string                               `json:"title,omitempty" xml:"title,omitempty"`
	TravelerStandard []*CarApplyAddRequestTravelerStandard `json:"traveler_standard,omitempty" xml:"traveler_standard,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN1415614
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CarApplyAddRequest) String() string {
	return dara.Prettify(s)
}

func (s CarApplyAddRequest) GoString() string {
	return s.String()
}

func (s *CarApplyAddRequest) GetCause() *string {
	return s.Cause
}

func (s *CarApplyAddRequest) GetCity() *string {
	return s.City
}

func (s *CarApplyAddRequest) GetCityCodeSet() *string {
	return s.CityCodeSet
}

func (s *CarApplyAddRequest) GetDate() *string {
	return s.Date
}

func (s *CarApplyAddRequest) GetFinishedDate() *string {
	return s.FinishedDate
}

func (s *CarApplyAddRequest) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *CarApplyAddRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CarApplyAddRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CarApplyAddRequest) GetThirdPartApplyId() *string {
	return s.ThirdPartApplyId
}

func (s *CarApplyAddRequest) GetThirdPartCostCenterId() *string {
	return s.ThirdPartCostCenterId
}

func (s *CarApplyAddRequest) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *CarApplyAddRequest) GetTimesTotal() *int32 {
	return s.TimesTotal
}

func (s *CarApplyAddRequest) GetTimesType() *int32 {
	return s.TimesType
}

func (s *CarApplyAddRequest) GetTimesUsed() *int32 {
	return s.TimesUsed
}

func (s *CarApplyAddRequest) GetTitle() *string {
	return s.Title
}

func (s *CarApplyAddRequest) GetTravelerStandard() []*CarApplyAddRequestTravelerStandard {
	return s.TravelerStandard
}

func (s *CarApplyAddRequest) GetUserId() *string {
	return s.UserId
}

func (s *CarApplyAddRequest) SetCause(v string) *CarApplyAddRequest {
	s.Cause = &v
	return s
}

func (s *CarApplyAddRequest) SetCity(v string) *CarApplyAddRequest {
	s.City = &v
	return s
}

func (s *CarApplyAddRequest) SetCityCodeSet(v string) *CarApplyAddRequest {
	s.CityCodeSet = &v
	return s
}

func (s *CarApplyAddRequest) SetDate(v string) *CarApplyAddRequest {
	s.Date = &v
	return s
}

func (s *CarApplyAddRequest) SetFinishedDate(v string) *CarApplyAddRequest {
	s.FinishedDate = &v
	return s
}

func (s *CarApplyAddRequest) SetProjectCode(v string) *CarApplyAddRequest {
	s.ProjectCode = &v
	return s
}

func (s *CarApplyAddRequest) SetProjectName(v string) *CarApplyAddRequest {
	s.ProjectName = &v
	return s
}

func (s *CarApplyAddRequest) SetStatus(v int32) *CarApplyAddRequest {
	s.Status = &v
	return s
}

func (s *CarApplyAddRequest) SetThirdPartApplyId(v string) *CarApplyAddRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *CarApplyAddRequest) SetThirdPartCostCenterId(v string) *CarApplyAddRequest {
	s.ThirdPartCostCenterId = &v
	return s
}

func (s *CarApplyAddRequest) SetThirdPartInvoiceId(v string) *CarApplyAddRequest {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *CarApplyAddRequest) SetTimesTotal(v int32) *CarApplyAddRequest {
	s.TimesTotal = &v
	return s
}

func (s *CarApplyAddRequest) SetTimesType(v int32) *CarApplyAddRequest {
	s.TimesType = &v
	return s
}

func (s *CarApplyAddRequest) SetTimesUsed(v int32) *CarApplyAddRequest {
	s.TimesUsed = &v
	return s
}

func (s *CarApplyAddRequest) SetTitle(v string) *CarApplyAddRequest {
	s.Title = &v
	return s
}

func (s *CarApplyAddRequest) SetTravelerStandard(v []*CarApplyAddRequestTravelerStandard) *CarApplyAddRequest {
	s.TravelerStandard = v
	return s
}

func (s *CarApplyAddRequest) SetUserId(v string) *CarApplyAddRequest {
	s.UserId = &v
	return s
}

func (s *CarApplyAddRequest) Validate() error {
	return dara.Validate(s)
}

type CarApplyAddRequestTravelerStandard struct {
	CarCitySet []*CarApplyAddRequestTravelerStandardCarCitySet `json:"car_city_set,omitempty" xml:"car_city_set,omitempty" type:"Repeated"`
	// This parameter is required.
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CarApplyAddRequestTravelerStandard) String() string {
	return dara.Prettify(s)
}

func (s CarApplyAddRequestTravelerStandard) GoString() string {
	return s.String()
}

func (s *CarApplyAddRequestTravelerStandard) GetCarCitySet() []*CarApplyAddRequestTravelerStandardCarCitySet {
	return s.CarCitySet
}

func (s *CarApplyAddRequestTravelerStandard) GetUserId() *string {
	return s.UserId
}

func (s *CarApplyAddRequestTravelerStandard) SetCarCitySet(v []*CarApplyAddRequestTravelerStandardCarCitySet) *CarApplyAddRequestTravelerStandard {
	s.CarCitySet = v
	return s
}

func (s *CarApplyAddRequestTravelerStandard) SetUserId(v string) *CarApplyAddRequestTravelerStandard {
	s.UserId = &v
	return s
}

func (s *CarApplyAddRequestTravelerStandard) Validate() error {
	return dara.Validate(s)
}

type CarApplyAddRequestTravelerStandardCarCitySet struct {
	// This parameter is required.
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// This parameter is required.
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
}

func (s CarApplyAddRequestTravelerStandardCarCitySet) String() string {
	return dara.Prettify(s)
}

func (s CarApplyAddRequestTravelerStandardCarCitySet) GoString() string {
	return s.String()
}

func (s *CarApplyAddRequestTravelerStandardCarCitySet) GetCityCode() *string {
	return s.CityCode
}

func (s *CarApplyAddRequestTravelerStandardCarCitySet) GetCityName() *string {
	return s.CityName
}

func (s *CarApplyAddRequestTravelerStandardCarCitySet) SetCityCode(v string) *CarApplyAddRequestTravelerStandardCarCitySet {
	s.CityCode = &v
	return s
}

func (s *CarApplyAddRequestTravelerStandardCarCitySet) SetCityName(v string) *CarApplyAddRequestTravelerStandardCarCitySet {
	s.CityName = &v
	return s
}

func (s *CarApplyAddRequestTravelerStandardCarCitySet) Validate() error {
	return dara.Validate(s)
}
