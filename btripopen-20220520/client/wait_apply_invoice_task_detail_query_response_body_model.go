// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWaitApplyInvoiceTaskDetailQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *WaitApplyInvoiceTaskDetailQueryResponseBody
	GetCode() *string
	SetMessage(v string) *WaitApplyInvoiceTaskDetailQueryResponseBody
	GetMessage() *string
	SetModule(v []*WaitApplyInvoiceTaskDetailQueryResponseBodyModule) *WaitApplyInvoiceTaskDetailQueryResponseBody
	GetModule() []*WaitApplyInvoiceTaskDetailQueryResponseBodyModule
	SetRequestId(v string) *WaitApplyInvoiceTaskDetailQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *WaitApplyInvoiceTaskDetailQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *WaitApplyInvoiceTaskDetailQueryResponseBody
	GetTraceId() *string
}

type WaitApplyInvoiceTaskDetailQueryResponseBody struct {
	// example:
	//
	// 200
	Code    *string                                              `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                              `json:"message,omitempty" xml:"message,omitempty"`
	Module  []*WaitApplyInvoiceTaskDetailQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210f079e16603757182131635d866a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s WaitApplyInvoiceTaskDetailQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WaitApplyInvoiceTaskDetailQueryResponseBody) GoString() string {
	return s.String()
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBody) GetModule() []*WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	return s.Module
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBody) SetCode(v string) *WaitApplyInvoiceTaskDetailQueryResponseBody {
	s.Code = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBody) SetMessage(v string) *WaitApplyInvoiceTaskDetailQueryResponseBody {
	s.Message = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBody) SetModule(v []*WaitApplyInvoiceTaskDetailQueryResponseBodyModule) *WaitApplyInvoiceTaskDetailQueryResponseBody {
	s.Module = v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBody) SetRequestId(v string) *WaitApplyInvoiceTaskDetailQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBody) SetSuccess(v bool) *WaitApplyInvoiceTaskDetailQueryResponseBody {
	s.Success = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBody) SetTraceId(v string) *WaitApplyInvoiceTaskDetailQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type WaitApplyInvoiceTaskDetailQueryResponseBodyModule struct {
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
	// example:
	//
	// 123
	InvoiceThirdPartId   *string `json:"invoice_third_part_id,omitempty" xml:"invoice_third_part_id,omitempty"`
	InvoiceTitle         *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
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

func (s WaitApplyInvoiceTaskDetailQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetContact() *string {
	return s.Contact
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetEmail() *string {
	return s.Email
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetFlightInvoiceFee() *string {
	return s.FlightInvoiceFee
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetFuPointInvoiceFee() *string {
	return s.FuPointInvoiceFee
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetHotelNormalInvoiceFee() *string {
	return s.HotelNormalInvoiceFee
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetHotelSpecialInvoiceFee() *string {
	return s.HotelSpecialInvoiceFee
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetInternationalFlightInvoiceFee() *string {
	return s.InternationalFlightInvoiceFee
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetInternationalHotelInvoiceFee() *string {
	return s.InternationalHotelInvoiceFee
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetInvoiceThirdPartId() *string {
	return s.InvoiceThirdPartId
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetMailAddress() *string {
	return s.MailAddress
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetMailCity() *string {
	return s.MailCity
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetMailFullAddress() *string {
	return s.MailFullAddress
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetMailProvince() *string {
	return s.MailProvince
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetMealNormalInvoiceFee() *string {
	return s.MealNormalInvoiceFee
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetPenaltyFee() *string {
	return s.PenaltyFee
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetRemark() *string {
	return s.Remark
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetServiceFee() *string {
	return s.ServiceFee
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetTelephone() *string {
	return s.Telephone
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetTrainAccelerationPackageInvoiceFee() *string {
	return s.TrainAccelerationPackageInvoiceFee
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetTrainInvoiceFee() *string {
	return s.TrainInvoiceFee
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetVacationNormalInvoiceFee() *string {
	return s.VacationNormalInvoiceFee
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetVasMallSpecialInvoiceFee() *string {
	return s.VasMallSpecialInvoiceFee
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetVehicleInvoiceFee() *string {
	return s.VehicleInvoiceFee
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) GetVehicleNormalInvoiceFee() *string {
	return s.VehicleNormalInvoiceFee
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetContact(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.Contact = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetEmail(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.Email = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetFlightInvoiceFee(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.FlightInvoiceFee = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetFuPointInvoiceFee(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.FuPointInvoiceFee = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetHotelNormalInvoiceFee(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.HotelNormalInvoiceFee = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetHotelSpecialInvoiceFee(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.HotelSpecialInvoiceFee = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetInternationalFlightInvoiceFee(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.InternationalFlightInvoiceFee = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetInternationalHotelInvoiceFee(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.InternationalHotelInvoiceFee = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetInvoiceThirdPartId(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.InvoiceThirdPartId = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetInvoiceTitle(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.InvoiceTitle = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetMailAddress(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.MailAddress = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetMailCity(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.MailCity = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetMailFullAddress(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.MailFullAddress = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetMailProvince(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.MailProvince = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetMealNormalInvoiceFee(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.MealNormalInvoiceFee = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetPenaltyFee(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.PenaltyFee = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetRemark(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.Remark = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetServiceFee(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.ServiceFee = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetTelephone(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.Telephone = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetTrainAccelerationPackageInvoiceFee(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.TrainAccelerationPackageInvoiceFee = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetTrainInvoiceFee(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.TrainInvoiceFee = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetVacationNormalInvoiceFee(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.VacationNormalInvoiceFee = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetVasMallSpecialInvoiceFee(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.VasMallSpecialInvoiceFee = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetVehicleInvoiceFee(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.VehicleInvoiceFee = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) SetVehicleNormalInvoiceFee(v string) *WaitApplyInvoiceTaskDetailQueryResponseBodyModule {
	s.VehicleNormalInvoiceFee = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
