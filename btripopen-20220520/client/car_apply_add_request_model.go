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
	SetItineraryList(v []*CarApplyAddRequestItineraryList) *CarApplyAddRequest
	GetItineraryList() []*CarApplyAddRequestItineraryList
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
	ItineraryList []*CarApplyAddRequestItineraryList `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty" type:"Repeated"`
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
	TravelerStandard []*CarApplyAddRequestTravelerStandard `json:"traveler_standard,omitempty" xml:"traveler_standard,omitempty" type:"Repeated"`
	// The third-party employee ID of the user who initiates the approval.
	//
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

func (s *CarApplyAddRequest) GetItineraryList() []*CarApplyAddRequestItineraryList {
	return s.ItineraryList
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

func (s *CarApplyAddRequest) SetItineraryList(v []*CarApplyAddRequestItineraryList) *CarApplyAddRequest {
	s.ItineraryList = v
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
	if s.ItineraryList != nil {
		for _, item := range s.ItineraryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TravelerStandard != nil {
		for _, item := range s.TravelerStandard {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CarApplyAddRequestItineraryList struct {
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
	// 440600，440100
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
	// 2025-11-25 21:00:00
	FinishedDate *string `json:"finished_date,omitempty" xml:"finished_date,omitempty"`
}

func (s CarApplyAddRequestItineraryList) String() string {
	return dara.Prettify(s)
}

func (s CarApplyAddRequestItineraryList) GoString() string {
	return s.String()
}

func (s *CarApplyAddRequestItineraryList) GetCity() *string {
	return s.City
}

func (s *CarApplyAddRequestItineraryList) GetCityCodeSet() *string {
	return s.CityCodeSet
}

func (s *CarApplyAddRequestItineraryList) GetDate() *string {
	return s.Date
}

func (s *CarApplyAddRequestItineraryList) GetFinishedDate() *string {
	return s.FinishedDate
}

func (s *CarApplyAddRequestItineraryList) SetCity(v string) *CarApplyAddRequestItineraryList {
	s.City = &v
	return s
}

func (s *CarApplyAddRequestItineraryList) SetCityCodeSet(v string) *CarApplyAddRequestItineraryList {
	s.CityCodeSet = &v
	return s
}

func (s *CarApplyAddRequestItineraryList) SetDate(v string) *CarApplyAddRequestItineraryList {
	s.Date = &v
	return s
}

func (s *CarApplyAddRequestItineraryList) SetFinishedDate(v string) *CarApplyAddRequestItineraryList {
	s.FinishedDate = &v
	return s
}

func (s *CarApplyAddRequestItineraryList) Validate() error {
	return dara.Validate(s)
}

type CarApplyAddRequestTravelerStandard struct {
	// The cross-city car service rules. Optional. If specified, cross-city rules are read from the approval form data.
	CarCitySet []*CarApplyAddRequestTravelerStandardCarCitySet `json:"car_city_set,omitempty" xml:"car_city_set,omitempty" type:"Repeated"`
	// The user ID of the traveler.
	//
	// This parameter is required.
	//
	// example:
	//
	// userid
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
	if s.CarCitySet != nil {
		for _, item := range s.CarCitySet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CarApplyAddRequestTravelerStandardCarCitySet struct {
	// The cross-city city code. Only 6-digit codes are supported. Separate multiple values with Chinese commas.
	//
	// Note: A maximum of 10 cities can be specified. The values in city_code and city_name must correspond one-to-one.
	//
	// This parameter is required.
	//
	// example:
	//
	// 110100，330100
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// The cross-city city name. Separate multiple values with Chinese commas.
	//
	// Note: A maximum of 10 cities can be specified. The values in city_code and city_name must correspond one-to-one.
	//
	// This parameter is required.
	//
	// example:
	//
	// 北京，杭州
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
