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
	SetItineraryListShrink(v string) *CarApplyAddShrinkRequest
	GetItineraryListShrink() *string
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
	// The reason for the business trip.
	//
	// This parameter is required.
	//
	// example:
	//
	// 访问客户
	Cause *string `json:"cause,omitempty" xml:"cause,omitempty"`
	// The car service cities. Separate multiple cities with Chinese commas (，).
	//
	// Note: A maximum of 10 cities can be specified. The values in city and city_code_set must correspond one-to-one.
	//
	// example:
	//
	// 北京，杭州
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// The set of city codes for intra-city car service. Separate multiple cities with Chinese commas (，).
	//
	// Note: 1) Either city_code_set or city is required. If both are specified, city_code_set takes precedence.
	//
	// A maximum of 10 cities can be specified.
	//
	// example:
	//
	// 110100，330100
	CityCodeSet *string `json:"city_code_set,omitempty" xml:"city_code_set,omitempty"`
	// The car service date. Access is controlled on a daily basis. For example, a value of 2021-03-18 20:26:56 indicates that the car service is available on 2021-03-18. For cross-day scenarios, use this parameter together with the finished_date parameter. The time parameter must be in the yyyy-MM-dd HH:mm:ss string format.
	//
	// example:
	//
	// 2022-07-12 14:52:52
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The car service end date. Access is controlled on a daily basis. For example, if date is set to 2021-03-18 20:26:56 and finished_date is set to 2021-03-30 20:26:56, the car service is available from 2021-03-18 (inclusive) to 2021-03-30 (inclusive). If this parameter is not specified, the value of date is used as the end date. The time parameter must be in the yyyy-MM-dd HH:mm:ss string format.
	//
	// example:
	//
	// 2022-07-12 18:51:25
	FinishedDate *string `json:"finished_date,omitempty" xml:"finished_date,omitempty"`
	// The intra-city car service itinerary.
	ItineraryListShrink *string `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty"`
	// The project code associated with the approval form.
	//
	// example:
	//
	// project1413
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// The project name associated with the approval form.
	//
	// example:
	//
	// 项目1
	ProjectName *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// The approval status.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the third-party approval form.
	//
	// This parameter is required.
	//
	// example:
	//
	// IRGS1413
	ThirdPartApplyId *string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	// The ID of the third-party cost center associated with the approval form.
	//
	// 	Warning: This field is required. To configure it as optional, contact operations.
	//
	// example:
	//
	// QA1411
	ThirdPartCostCenterId *string `json:"third_part_cost_center_id,omitempty" xml:"third_part_cost_center_id,omitempty"`
	// The ID of the third-party invoice header associated with the approval form.
	//
	// 	Warning: This field is required. To configure it as optional, contact operations.
	//
	// example:
	//
	// GA15131
	ThirdPartInvoiceId *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	// The total available count for the approval form.
	//
	// example:
	//
	// 1
	TimesTotal *int32 `json:"times_total,omitempty" xml:"times_total,omitempty"`
	// The type of available usage count for the approval form. If the enterprise does not need to limit the number of times the approval form can be used, set this parameter to 1 (unlimited) and set both times_total and times_used to 0.
	//
	// Valid values:
	//
	// - 1: Unlimited.
	//
	// - 2: User-specified count.
	//
	// - 3: Admin-limited count.
	//
	// example:
	//
	// 1
	TimesType *int32 `json:"times_type,omitempty" xml:"times_type,omitempty"`
	// The number of times the approval form has been used.
	//
	// example:
	//
	// 1
	TimesUsed *int32 `json:"times_used,omitempty" xml:"times_used,omitempty"`
	// The title of the approval form.
	//
	// This parameter is required.
	//
	// example:
	//
	// 访问客户
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The intra-city car service rules.
	TravelerStandardShrink *string `json:"traveler_standard,omitempty" xml:"traveler_standard,omitempty"`
	// The third-party employee ID of the user who initiates the approval.
	//
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

func (s *CarApplyAddShrinkRequest) GetItineraryListShrink() *string {
	return s.ItineraryListShrink
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

func (s *CarApplyAddShrinkRequest) SetItineraryListShrink(v string) *CarApplyAddShrinkRequest {
	s.ItineraryListShrink = &v
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
