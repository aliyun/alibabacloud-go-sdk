// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyAddRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyUser(v *MealApplyAddRequestApplyUser) *MealApplyAddRequest
	GetApplyUser() *MealApplyAddRequestApplyUser
	SetCostCenterId(v int64) *MealApplyAddRequest
	GetCostCenterId() *int64
	SetInvoiceId(v int64) *MealApplyAddRequest
	GetInvoiceId() *int64
	SetItineraryList(v []*MealApplyAddRequestItineraryList) *MealApplyAddRequest
	GetItineraryList() []*MealApplyAddRequestItineraryList
	SetMealAmount(v int64) *MealApplyAddRequest
	GetMealAmount() *int64
	SetMealCause(v string) *MealApplyAddRequest
	GetMealCause() *string
	SetProjectCode(v string) *MealApplyAddRequest
	GetProjectCode() *string
	SetProjectTitle(v string) *MealApplyAddRequest
	GetProjectTitle() *string
	SetStatus(v int32) *MealApplyAddRequest
	GetStatus() *int32
	SetThirdPartApplyId(v string) *MealApplyAddRequest
	GetThirdPartApplyId() *string
	SetThirdPartCostCenterId(v string) *MealApplyAddRequest
	GetThirdPartCostCenterId() *string
	SetThirdPartInvoiceId(v string) *MealApplyAddRequest
	GetThirdPartInvoiceId() *string
}

type MealApplyAddRequest struct {
	// This parameter is required.
	ApplyUser *MealApplyAddRequestApplyUser `json:"apply_user,omitempty" xml:"apply_user,omitempty" type:"Struct"`
	// example:
	//
	// 23
	CostCenterId *int64 `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	// example:
	//
	// 123
	InvoiceId *int64 `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// This parameter is required.
	ItineraryList []*MealApplyAddRequestItineraryList `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty" type:"Repeated"`
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

func (s MealApplyAddRequest) String() string {
	return dara.Prettify(s)
}

func (s MealApplyAddRequest) GoString() string {
	return s.String()
}

func (s *MealApplyAddRequest) GetApplyUser() *MealApplyAddRequestApplyUser {
	return s.ApplyUser
}

func (s *MealApplyAddRequest) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *MealApplyAddRequest) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *MealApplyAddRequest) GetItineraryList() []*MealApplyAddRequestItineraryList {
	return s.ItineraryList
}

func (s *MealApplyAddRequest) GetMealAmount() *int64 {
	return s.MealAmount
}

func (s *MealApplyAddRequest) GetMealCause() *string {
	return s.MealCause
}

func (s *MealApplyAddRequest) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *MealApplyAddRequest) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *MealApplyAddRequest) GetStatus() *int32 {
	return s.Status
}

func (s *MealApplyAddRequest) GetThirdPartApplyId() *string {
	return s.ThirdPartApplyId
}

func (s *MealApplyAddRequest) GetThirdPartCostCenterId() *string {
	return s.ThirdPartCostCenterId
}

func (s *MealApplyAddRequest) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *MealApplyAddRequest) SetApplyUser(v *MealApplyAddRequestApplyUser) *MealApplyAddRequest {
	s.ApplyUser = v
	return s
}

func (s *MealApplyAddRequest) SetCostCenterId(v int64) *MealApplyAddRequest {
	s.CostCenterId = &v
	return s
}

func (s *MealApplyAddRequest) SetInvoiceId(v int64) *MealApplyAddRequest {
	s.InvoiceId = &v
	return s
}

func (s *MealApplyAddRequest) SetItineraryList(v []*MealApplyAddRequestItineraryList) *MealApplyAddRequest {
	s.ItineraryList = v
	return s
}

func (s *MealApplyAddRequest) SetMealAmount(v int64) *MealApplyAddRequest {
	s.MealAmount = &v
	return s
}

func (s *MealApplyAddRequest) SetMealCause(v string) *MealApplyAddRequest {
	s.MealCause = &v
	return s
}

func (s *MealApplyAddRequest) SetProjectCode(v string) *MealApplyAddRequest {
	s.ProjectCode = &v
	return s
}

func (s *MealApplyAddRequest) SetProjectTitle(v string) *MealApplyAddRequest {
	s.ProjectTitle = &v
	return s
}

func (s *MealApplyAddRequest) SetStatus(v int32) *MealApplyAddRequest {
	s.Status = &v
	return s
}

func (s *MealApplyAddRequest) SetThirdPartApplyId(v string) *MealApplyAddRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *MealApplyAddRequest) SetThirdPartCostCenterId(v string) *MealApplyAddRequest {
	s.ThirdPartCostCenterId = &v
	return s
}

func (s *MealApplyAddRequest) SetThirdPartInvoiceId(v string) *MealApplyAddRequest {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *MealApplyAddRequest) Validate() error {
	return dara.Validate(s)
}

type MealApplyAddRequestApplyUser struct {
	// This parameter is required.
	//
	// example:
	//
	// userId1
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s MealApplyAddRequestApplyUser) String() string {
	return dara.Prettify(s)
}

func (s MealApplyAddRequestApplyUser) GoString() string {
	return s.String()
}

func (s *MealApplyAddRequestApplyUser) GetUserId() *string {
	return s.UserId
}

func (s *MealApplyAddRequestApplyUser) SetUserId(v string) *MealApplyAddRequestApplyUser {
	s.UserId = &v
	return s
}

func (s *MealApplyAddRequestApplyUser) Validate() error {
	return dara.Validate(s)
}

type MealApplyAddRequestItineraryList struct {
	// This parameter is required.
	Cities []*MealApplyAddRequestItineraryListCities `json:"cities,omitempty" xml:"cities,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-02-05 00:00:00
	EndDate *string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-02-05 00:00:00
	StartDate *string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// example:
	//
	// 2134
	ThirdpartItineraryId *string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
}

func (s MealApplyAddRequestItineraryList) String() string {
	return dara.Prettify(s)
}

func (s MealApplyAddRequestItineraryList) GoString() string {
	return s.String()
}

func (s *MealApplyAddRequestItineraryList) GetCities() []*MealApplyAddRequestItineraryListCities {
	return s.Cities
}

func (s *MealApplyAddRequestItineraryList) GetEndDate() *string {
	return s.EndDate
}

func (s *MealApplyAddRequestItineraryList) GetStartDate() *string {
	return s.StartDate
}

func (s *MealApplyAddRequestItineraryList) GetThirdpartItineraryId() *string {
	return s.ThirdpartItineraryId
}

func (s *MealApplyAddRequestItineraryList) SetCities(v []*MealApplyAddRequestItineraryListCities) *MealApplyAddRequestItineraryList {
	s.Cities = v
	return s
}

func (s *MealApplyAddRequestItineraryList) SetEndDate(v string) *MealApplyAddRequestItineraryList {
	s.EndDate = &v
	return s
}

func (s *MealApplyAddRequestItineraryList) SetStartDate(v string) *MealApplyAddRequestItineraryList {
	s.StartDate = &v
	return s
}

func (s *MealApplyAddRequestItineraryList) SetThirdpartItineraryId(v string) *MealApplyAddRequestItineraryList {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *MealApplyAddRequestItineraryList) Validate() error {
	return dara.Validate(s)
}

type MealApplyAddRequestItineraryListCities struct {
	// example:
	//
	// 330702
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
}

func (s MealApplyAddRequestItineraryListCities) String() string {
	return dara.Prettify(s)
}

func (s MealApplyAddRequestItineraryListCities) GoString() string {
	return s.String()
}

func (s *MealApplyAddRequestItineraryListCities) GetCityCode() *string {
	return s.CityCode
}

func (s *MealApplyAddRequestItineraryListCities) GetCityName() *string {
	return s.CityName
}

func (s *MealApplyAddRequestItineraryListCities) SetCityCode(v string) *MealApplyAddRequestItineraryListCities {
	s.CityCode = &v
	return s
}

func (s *MealApplyAddRequestItineraryListCities) SetCityName(v string) *MealApplyAddRequestItineraryListCities {
	s.CityName = &v
	return s
}

func (s *MealApplyAddRequestItineraryListCities) Validate() error {
	return dara.Validate(s)
}
