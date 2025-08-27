// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyInvoiceTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillDate(v string) *ApplyInvoiceTaskRequest
	GetBillDate() *string
	SetInvoiceTaskList(v []*ApplyInvoiceTaskRequestInvoiceTaskList) *ApplyInvoiceTaskRequest
	GetInvoiceTaskList() []*ApplyInvoiceTaskRequestInvoiceTaskList
}

type ApplyInvoiceTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2022-12-01
	BillDate *string `json:"bill_date,omitempty" xml:"bill_date,omitempty"`
	// This parameter is required.
	InvoiceTaskList []*ApplyInvoiceTaskRequestInvoiceTaskList `json:"invoice_task_list,omitempty" xml:"invoice_task_list,omitempty" type:"Repeated"`
}

func (s ApplyInvoiceTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyInvoiceTaskRequest) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceTaskRequest) GetBillDate() *string {
	return s.BillDate
}

func (s *ApplyInvoiceTaskRequest) GetInvoiceTaskList() []*ApplyInvoiceTaskRequestInvoiceTaskList {
	return s.InvoiceTaskList
}

func (s *ApplyInvoiceTaskRequest) SetBillDate(v string) *ApplyInvoiceTaskRequest {
	s.BillDate = &v
	return s
}

func (s *ApplyInvoiceTaskRequest) SetInvoiceTaskList(v []*ApplyInvoiceTaskRequestInvoiceTaskList) *ApplyInvoiceTaskRequest {
	s.InvoiceTaskList = v
	return s
}

func (s *ApplyInvoiceTaskRequest) Validate() error {
	return dara.Validate(s)
}

type ApplyInvoiceTaskRequestInvoiceTaskList struct {
	Contact *string `json:"contact,omitempty" xml:"contact,omitempty"`
	Email   *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 0
	FlightInvoiceFee *string `json:"flight_invoice_fee,omitempty" xml:"flight_invoice_fee,omitempty"`
	// example:
	//
	// 3.12
	FuPointInvoiceFee *string `json:"fu_point_invoice_fee,omitempty" xml:"fu_point_invoice_fee,omitempty"`
	// example:
	//
	// 0
	HotelNormalInvoiceFee *string `json:"hotel_normal_invoice_fee,omitempty" xml:"hotel_normal_invoice_fee,omitempty"`
	// example:
	//
	// 100
	HotelSpecialInvoiceFee *string `json:"hotel_special_invoice_fee,omitempty" xml:"hotel_special_invoice_fee,omitempty"`
	// example:
	//
	// 0
	InternationalFlightInvoiceFee *string `json:"international_flight_invoice_fee,omitempty" xml:"international_flight_invoice_fee,omitempty"`
	InternationalHotelInvoiceFee  *string `json:"international_hotel_invoice_fee,omitempty" xml:"international_hotel_invoice_fee,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	InvoiceThirdPartId   *string `json:"invoice_third_part_id,omitempty" xml:"invoice_third_part_id,omitempty"`
	InvoiceType          *int32  `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	MailAddress          *string `json:"mail_address,omitempty" xml:"mail_address,omitempty"`
	MailCity             *string `json:"mail_city,omitempty" xml:"mail_city,omitempty"`
	MailFullAddress      *string `json:"mail_full_address,omitempty" xml:"mail_full_address,omitempty"`
	MailProvince         *string `json:"mail_province,omitempty" xml:"mail_province,omitempty"`
	MealNormalInvoiceFee *string `json:"meal_normal_invoice_fee,omitempty" xml:"meal_normal_invoice_fee,omitempty"`
	// example:
	//
	// 0
	PenaltyFee *string `json:"penalty_fee,omitempty" xml:"penalty_fee,omitempty"`
	Remark     *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// 1.02
	ServiceFee *string `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// example:
	//
	// 1234567890
	Telephone                          *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	TrainAccelerationPackageInvoiceFee *string `json:"train_acceleration_package_invoice_fee,omitempty" xml:"train_acceleration_package_invoice_fee,omitempty"`
	// example:
	//
	// 0
	TrainInvoiceFee          *string `json:"train_invoice_fee,omitempty" xml:"train_invoice_fee,omitempty"`
	VacationNormalInvoiceFee *string `json:"vacation_normal_invoice_fee,omitempty" xml:"vacation_normal_invoice_fee,omitempty"`
	VasMallSpecialInvoiceFee *string `json:"vas_mall_special_invoice_fee,omitempty" xml:"vas_mall_special_invoice_fee,omitempty"`
	// example:
	//
	// 100
	VehicleInvoiceFee       *string `json:"vehicle_invoice_fee,omitempty" xml:"vehicle_invoice_fee,omitempty"`
	VehicleNormalInvoiceFee *string `json:"vehicle_normal_invoice_fee,omitempty" xml:"vehicle_normal_invoice_fee,omitempty"`
}

func (s ApplyInvoiceTaskRequestInvoiceTaskList) String() string {
	return dara.Prettify(s)
}

func (s ApplyInvoiceTaskRequestInvoiceTaskList) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetContact() *string {
	return s.Contact
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetEmail() *string {
	return s.Email
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetFlightInvoiceFee() *string {
	return s.FlightInvoiceFee
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetFuPointInvoiceFee() *string {
	return s.FuPointInvoiceFee
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetHotelNormalInvoiceFee() *string {
	return s.HotelNormalInvoiceFee
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetHotelSpecialInvoiceFee() *string {
	return s.HotelSpecialInvoiceFee
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetInternationalFlightInvoiceFee() *string {
	return s.InternationalFlightInvoiceFee
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetInternationalHotelInvoiceFee() *string {
	return s.InternationalHotelInvoiceFee
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetInvoiceThirdPartId() *string {
	return s.InvoiceThirdPartId
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetInvoiceType() *int32 {
	return s.InvoiceType
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetMailAddress() *string {
	return s.MailAddress
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetMailCity() *string {
	return s.MailCity
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetMailFullAddress() *string {
	return s.MailFullAddress
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetMailProvince() *string {
	return s.MailProvince
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetMealNormalInvoiceFee() *string {
	return s.MealNormalInvoiceFee
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetPenaltyFee() *string {
	return s.PenaltyFee
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetRemark() *string {
	return s.Remark
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetServiceFee() *string {
	return s.ServiceFee
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetTelephone() *string {
	return s.Telephone
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetTrainAccelerationPackageInvoiceFee() *string {
	return s.TrainAccelerationPackageInvoiceFee
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetTrainInvoiceFee() *string {
	return s.TrainInvoiceFee
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetVacationNormalInvoiceFee() *string {
	return s.VacationNormalInvoiceFee
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetVasMallSpecialInvoiceFee() *string {
	return s.VasMallSpecialInvoiceFee
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetVehicleInvoiceFee() *string {
	return s.VehicleInvoiceFee
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) GetVehicleNormalInvoiceFee() *string {
	return s.VehicleNormalInvoiceFee
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetContact(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.Contact = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetEmail(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.Email = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetFlightInvoiceFee(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.FlightInvoiceFee = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetFuPointInvoiceFee(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.FuPointInvoiceFee = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetHotelNormalInvoiceFee(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.HotelNormalInvoiceFee = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetHotelSpecialInvoiceFee(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.HotelSpecialInvoiceFee = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetInternationalFlightInvoiceFee(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.InternationalFlightInvoiceFee = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetInternationalHotelInvoiceFee(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.InternationalHotelInvoiceFee = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetInvoiceThirdPartId(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.InvoiceThirdPartId = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetInvoiceType(v int32) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.InvoiceType = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetMailAddress(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.MailAddress = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetMailCity(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.MailCity = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetMailFullAddress(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.MailFullAddress = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetMailProvince(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.MailProvince = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetMealNormalInvoiceFee(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.MealNormalInvoiceFee = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetPenaltyFee(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.PenaltyFee = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetRemark(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.Remark = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetServiceFee(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.ServiceFee = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetTelephone(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.Telephone = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetTrainAccelerationPackageInvoiceFee(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.TrainAccelerationPackageInvoiceFee = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetTrainInvoiceFee(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.TrainInvoiceFee = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetVacationNormalInvoiceFee(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.VacationNormalInvoiceFee = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetVasMallSpecialInvoiceFee(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.VasMallSpecialInvoiceFee = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetVehicleInvoiceFee(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.VehicleInvoiceFee = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) SetVehicleNormalInvoiceFee(v string) *ApplyInvoiceTaskRequestInvoiceTaskList {
	s.VehicleNormalInvoiceFee = &v
	return s
}

func (s *ApplyInvoiceTaskRequestInvoiceTaskList) Validate() error {
	return dara.Validate(s)
}
