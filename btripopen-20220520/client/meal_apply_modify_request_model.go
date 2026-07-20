// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyModifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyUser(v *MealApplyModifyRequestApplyUser) *MealApplyModifyRequest
	GetApplyUser() *MealApplyModifyRequestApplyUser
	SetCostCenterId(v int64) *MealApplyModifyRequest
	GetCostCenterId() *int64
	SetExtendField(v string) *MealApplyModifyRequest
	GetExtendField() *string
	SetInvoiceId(v int64) *MealApplyModifyRequest
	GetInvoiceId() *int64
	SetItineraryList(v []*MealApplyModifyRequestItineraryList) *MealApplyModifyRequest
	GetItineraryList() []*MealApplyModifyRequestItineraryList
	SetMealAmount(v int64) *MealApplyModifyRequest
	GetMealAmount() *int64
	SetMealCause(v string) *MealApplyModifyRequest
	GetMealCause() *string
	SetProjectCode(v string) *MealApplyModifyRequest
	GetProjectCode() *string
	SetProjectTitle(v string) *MealApplyModifyRequest
	GetProjectTitle() *string
	SetStatus(v int32) *MealApplyModifyRequest
	GetStatus() *int32
	SetThirdPartApplyId(v string) *MealApplyModifyRequest
	GetThirdPartApplyId() *string
	SetThirdPartCostCenterId(v string) *MealApplyModifyRequest
	GetThirdPartCostCenterId() *string
	SetThirdPartInvoiceId(v string) *MealApplyModifyRequest
	GetThirdPartInvoiceId() *string
}

type MealApplyModifyRequest struct {
	// This parameter is required.
	ApplyUser *MealApplyModifyRequestApplyUser `json:"apply_user,omitempty" xml:"apply_user,omitempty" type:"Struct"`
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
	InvoiceId     *int64                                 `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	ItineraryList []*MealApplyModifyRequestItineraryList `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty" type:"Repeated"`
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

func (s MealApplyModifyRequest) String() string {
	return dara.Prettify(s)
}

func (s MealApplyModifyRequest) GoString() string {
	return s.String()
}

func (s *MealApplyModifyRequest) GetApplyUser() *MealApplyModifyRequestApplyUser {
	return s.ApplyUser
}

func (s *MealApplyModifyRequest) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *MealApplyModifyRequest) GetExtendField() *string {
	return s.ExtendField
}

func (s *MealApplyModifyRequest) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *MealApplyModifyRequest) GetItineraryList() []*MealApplyModifyRequestItineraryList {
	return s.ItineraryList
}

func (s *MealApplyModifyRequest) GetMealAmount() *int64 {
	return s.MealAmount
}

func (s *MealApplyModifyRequest) GetMealCause() *string {
	return s.MealCause
}

func (s *MealApplyModifyRequest) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *MealApplyModifyRequest) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *MealApplyModifyRequest) GetStatus() *int32 {
	return s.Status
}

func (s *MealApplyModifyRequest) GetThirdPartApplyId() *string {
	return s.ThirdPartApplyId
}

func (s *MealApplyModifyRequest) GetThirdPartCostCenterId() *string {
	return s.ThirdPartCostCenterId
}

func (s *MealApplyModifyRequest) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *MealApplyModifyRequest) SetApplyUser(v *MealApplyModifyRequestApplyUser) *MealApplyModifyRequest {
	s.ApplyUser = v
	return s
}

func (s *MealApplyModifyRequest) SetCostCenterId(v int64) *MealApplyModifyRequest {
	s.CostCenterId = &v
	return s
}

func (s *MealApplyModifyRequest) SetExtendField(v string) *MealApplyModifyRequest {
	s.ExtendField = &v
	return s
}

func (s *MealApplyModifyRequest) SetInvoiceId(v int64) *MealApplyModifyRequest {
	s.InvoiceId = &v
	return s
}

func (s *MealApplyModifyRequest) SetItineraryList(v []*MealApplyModifyRequestItineraryList) *MealApplyModifyRequest {
	s.ItineraryList = v
	return s
}

func (s *MealApplyModifyRequest) SetMealAmount(v int64) *MealApplyModifyRequest {
	s.MealAmount = &v
	return s
}

func (s *MealApplyModifyRequest) SetMealCause(v string) *MealApplyModifyRequest {
	s.MealCause = &v
	return s
}

func (s *MealApplyModifyRequest) SetProjectCode(v string) *MealApplyModifyRequest {
	s.ProjectCode = &v
	return s
}

func (s *MealApplyModifyRequest) SetProjectTitle(v string) *MealApplyModifyRequest {
	s.ProjectTitle = &v
	return s
}

func (s *MealApplyModifyRequest) SetStatus(v int32) *MealApplyModifyRequest {
	s.Status = &v
	return s
}

func (s *MealApplyModifyRequest) SetThirdPartApplyId(v string) *MealApplyModifyRequest {
	s.ThirdPartApplyId = &v
	return s
}

func (s *MealApplyModifyRequest) SetThirdPartCostCenterId(v string) *MealApplyModifyRequest {
	s.ThirdPartCostCenterId = &v
	return s
}

func (s *MealApplyModifyRequest) SetThirdPartInvoiceId(v string) *MealApplyModifyRequest {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *MealApplyModifyRequest) Validate() error {
	if s.ApplyUser != nil {
		if err := s.ApplyUser.Validate(); err != nil {
			return err
		}
	}
	if s.ItineraryList != nil {
		for _, item := range s.ItineraryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MealApplyModifyRequestApplyUser struct {
	// example:
	//
	// userId1
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s MealApplyModifyRequestApplyUser) String() string {
	return dara.Prettify(s)
}

func (s MealApplyModifyRequestApplyUser) GoString() string {
	return s.String()
}

func (s *MealApplyModifyRequestApplyUser) GetUserId() *string {
	return s.UserId
}

func (s *MealApplyModifyRequestApplyUser) SetUserId(v string) *MealApplyModifyRequestApplyUser {
	s.UserId = &v
	return s
}

func (s *MealApplyModifyRequestApplyUser) Validate() error {
	return dara.Validate(s)
}

type MealApplyModifyRequestItineraryList struct {
	Cities []*MealApplyModifyRequestItineraryListCities `json:"cities,omitempty" xml:"cities,omitempty" type:"Repeated"`
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

func (s MealApplyModifyRequestItineraryList) String() string {
	return dara.Prettify(s)
}

func (s MealApplyModifyRequestItineraryList) GoString() string {
	return s.String()
}

func (s *MealApplyModifyRequestItineraryList) GetCities() []*MealApplyModifyRequestItineraryListCities {
	return s.Cities
}

func (s *MealApplyModifyRequestItineraryList) GetEndDate() *string {
	return s.EndDate
}

func (s *MealApplyModifyRequestItineraryList) GetStartDate() *string {
	return s.StartDate
}

func (s *MealApplyModifyRequestItineraryList) GetThirdpartItineraryId() *string {
	return s.ThirdpartItineraryId
}

func (s *MealApplyModifyRequestItineraryList) SetCities(v []*MealApplyModifyRequestItineraryListCities) *MealApplyModifyRequestItineraryList {
	s.Cities = v
	return s
}

func (s *MealApplyModifyRequestItineraryList) SetEndDate(v string) *MealApplyModifyRequestItineraryList {
	s.EndDate = &v
	return s
}

func (s *MealApplyModifyRequestItineraryList) SetStartDate(v string) *MealApplyModifyRequestItineraryList {
	s.StartDate = &v
	return s
}

func (s *MealApplyModifyRequestItineraryList) SetThirdpartItineraryId(v string) *MealApplyModifyRequestItineraryList {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *MealApplyModifyRequestItineraryList) Validate() error {
	if s.Cities != nil {
		for _, item := range s.Cities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MealApplyModifyRequestItineraryListCities struct {
	// example:
	//
	// 330702
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// example:
	//
	// 杭州
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
}

func (s MealApplyModifyRequestItineraryListCities) String() string {
	return dara.Prettify(s)
}

func (s MealApplyModifyRequestItineraryListCities) GoString() string {
	return s.String()
}

func (s *MealApplyModifyRequestItineraryListCities) GetCityCode() *string {
	return s.CityCode
}

func (s *MealApplyModifyRequestItineraryListCities) GetCityName() *string {
	return s.CityName
}

func (s *MealApplyModifyRequestItineraryListCities) SetCityCode(v string) *MealApplyModifyRequestItineraryListCities {
	s.CityCode = &v
	return s
}

func (s *MealApplyModifyRequestItineraryListCities) SetCityName(v string) *MealApplyModifyRequestItineraryListCities {
	s.CityName = &v
	return s
}

func (s *MealApplyModifyRequestItineraryListCities) Validate() error {
	return dara.Validate(s)
}
