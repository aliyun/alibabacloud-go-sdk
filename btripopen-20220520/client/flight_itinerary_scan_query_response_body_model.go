// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightItineraryScanQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightItineraryScanQueryResponseBody
	GetCode() *string
	SetMessage(v string) *FlightItineraryScanQueryResponseBody
	GetMessage() *string
	SetModule(v *FlightItineraryScanQueryResponseBodyModule) *FlightItineraryScanQueryResponseBody
	GetModule() *FlightItineraryScanQueryResponseBodyModule
	SetRequestId(v string) *FlightItineraryScanQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightItineraryScanQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightItineraryScanQueryResponseBody
	GetTraceId() *string
}

type FlightItineraryScanQueryResponseBody struct {
	// example:
	//
	// 200
	Code    *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                     `json:"message,omitempty" xml:"message,omitempty"`
	Module  *FlightItineraryScanQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightItineraryScanQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightItineraryScanQueryResponseBody) GoString() string {
	return s.String()
}

func (s *FlightItineraryScanQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightItineraryScanQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightItineraryScanQueryResponseBody) GetModule() *FlightItineraryScanQueryResponseBodyModule {
	return s.Module
}

func (s *FlightItineraryScanQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightItineraryScanQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightItineraryScanQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightItineraryScanQueryResponseBody) SetCode(v string) *FlightItineraryScanQueryResponseBody {
	s.Code = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBody) SetMessage(v string) *FlightItineraryScanQueryResponseBody {
	s.Message = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBody) SetModule(v *FlightItineraryScanQueryResponseBodyModule) *FlightItineraryScanQueryResponseBody {
	s.Module = v
	return s
}

func (s *FlightItineraryScanQueryResponseBody) SetRequestId(v string) *FlightItineraryScanQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBody) SetSuccess(v bool) *FlightItineraryScanQueryResponseBody {
	s.Success = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBody) SetTraceId(v string) *FlightItineraryScanQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightItineraryScanQueryResponseBodyModule struct {
	Items []*FlightItineraryScanQueryResponseBodyModuleItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int32 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 2
	TotalPage *int32 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// example:
	//
	// 30
	TotalSize *int32 `json:"total_size,omitempty" xml:"total_size,omitempty"`
}

func (s FlightItineraryScanQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightItineraryScanQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightItineraryScanQueryResponseBodyModule) GetItems() []*FlightItineraryScanQueryResponseBodyModuleItems {
	return s.Items
}

func (s *FlightItineraryScanQueryResponseBodyModule) GetPageNo() *int32 {
	return s.PageNo
}

func (s *FlightItineraryScanQueryResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *FlightItineraryScanQueryResponseBodyModule) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *FlightItineraryScanQueryResponseBodyModule) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *FlightItineraryScanQueryResponseBodyModule) SetItems(v []*FlightItineraryScanQueryResponseBodyModuleItems) *FlightItineraryScanQueryResponseBodyModule {
	s.Items = v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModule) SetPageNo(v int32) *FlightItineraryScanQueryResponseBodyModule {
	s.PageNo = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModule) SetPageSize(v int32) *FlightItineraryScanQueryResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModule) SetTotalPage(v int32) *FlightItineraryScanQueryResponseBodyModule {
	s.TotalPage = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModule) SetTotalSize(v int32) *FlightItineraryScanQueryResponseBodyModule {
	s.TotalSize = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightItineraryScanQueryResponseBodyModuleItems struct {
	// 销售单位代号
	//
	// example:
	//
	// SIA25608336893
	AgentCode *string `json:"agent_code,omitempty" xml:"agent_code,omitempty"`
	ApplyId   *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// example:
	//
	// 2022-12-01
	BillDate *string `json:"bill_date,omitempty" xml:"bill_date,omitempty"`
	// example:
	//
	// 50
	Build      *string `json:"build,omitempty" xml:"build,omitempty"`
	CostCenter *string `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	Department *string `json:"department,omitempty" xml:"department,omitempty"`
	// 机票行程明细
	Flights []*FlightItineraryScanQueryResponseBodyModuleItemsFlights `json:"flights,omitempty" xml:"flights,omitempty" type:"Repeated"`
	// example:
	//
	// 120
	FuelSurcharge *string `json:"fuel_surcharge,omitempty" xml:"fuel_surcharge,omitempty"`
	// UK
	//
	// example:
	//
	// 30
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 0
	Insurance    *string `json:"insurance,omitempty" xml:"insurance,omitempty"`
	InvoiceTitle *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	InvoiceType  *int32  `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 填开单位
	IssueCompany *string `json:"issue_company,omitempty" xml:"issue_company,omitempty"`
	// 填开日期
	//
	// example:
	//
	// 2019-02-28
	IssueDate *string `json:"issue_date,omitempty" xml:"issue_date,omitempty"`
	// example:
	//
	// 6666666666
	ItineraryNum *string `json:"itinerary_num,omitempty" xml:"itinerary_num,omitempty"`
	OfdOssUrl    *string `json:"ofd_oss_url,omitempty" xml:"ofd_oss_url,omitempty"`
	// example:
	//
	// 4801105714092
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// https://www.testurl.com
	OssUrl        *string `json:"oss_url,omitempty" xml:"oss_url,omitempty"`
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	PdfOssUrl     *string `json:"pdf_oss_url,omitempty" xml:"pdf_oss_url,omitempty"`
	Project       *string `json:"project,omitempty" xml:"project,omitempty"`
	// 提示信息
	PromptMessage  *string `json:"prompt_message,omitempty" xml:"prompt_message,omitempty"`
	PurchaserName  *string `json:"purchaser_name,omitempty" xml:"purchaser_name,omitempty"`
	PurchaserTaxNo *string `json:"purchaser_tax_no,omitempty" xml:"purchaser_tax_no,omitempty"`
	PurchaserType  *int32  `json:"purchaser_type,omitempty" xml:"purchaser_type,omitempty"`
	// example:
	//
	// 108.17
	TaxAmount *string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
	// example:
	//
	// 9%
	TaxRate *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// example:
	//
	// 784-1111111111
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// example:
	//
	// 1190
	TicketPrice *string `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// 1360
	TotalPrice *string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 验证码
	//
	// example:
	//
	// 9817
	ValidationCode *string `json:"validation_code,omitempty" xml:"validation_code,omitempty"`
}

func (s FlightItineraryScanQueryResponseBodyModuleItems) String() string {
	return dara.Prettify(s)
}

func (s FlightItineraryScanQueryResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetAgentCode() *string {
	return s.AgentCode
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetApplyId() *string {
	return s.ApplyId
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetBillDate() *string {
	return s.BillDate
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetBuild() *string {
	return s.Build
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetCostCenter() *string {
	return s.CostCenter
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetDepartment() *string {
	return s.Department
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetFlights() []*FlightItineraryScanQueryResponseBodyModuleItemsFlights {
	return s.Flights
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetFuelSurcharge() *string {
	return s.FuelSurcharge
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetId() *string {
	return s.Id
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetInsurance() *string {
	return s.Insurance
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetInvoiceType() *int32 {
	return s.InvoiceType
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetIssueCompany() *string {
	return s.IssueCompany
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetIssueDate() *string {
	return s.IssueDate
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetItineraryNum() *string {
	return s.ItineraryNum
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetOfdOssUrl() *string {
	return s.OfdOssUrl
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetOssUrl() *string {
	return s.OssUrl
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetPassengerName() *string {
	return s.PassengerName
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetPdfOssUrl() *string {
	return s.PdfOssUrl
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetProject() *string {
	return s.Project
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetPromptMessage() *string {
	return s.PromptMessage
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetPurchaserName() *string {
	return s.PurchaserName
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetPurchaserTaxNo() *string {
	return s.PurchaserTaxNo
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetPurchaserType() *int32 {
	return s.PurchaserType
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetTaxAmount() *string {
	return s.TaxAmount
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetTaxRate() *string {
	return s.TaxRate
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetTicketNo() *string {
	return s.TicketNo
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetTicketPrice() *string {
	return s.TicketPrice
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetTotalPrice() *string {
	return s.TotalPrice
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) GetValidationCode() *string {
	return s.ValidationCode
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetAgentCode(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.AgentCode = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetApplyId(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.ApplyId = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetBillDate(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.BillDate = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetBuild(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.Build = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetCostCenter(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.CostCenter = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetDepartment(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.Department = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetFlights(v []*FlightItineraryScanQueryResponseBodyModuleItemsFlights) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.Flights = v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetFuelSurcharge(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.FuelSurcharge = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetId(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.Id = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetInsurance(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.Insurance = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetInvoiceTitle(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.InvoiceTitle = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetInvoiceType(v int32) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.InvoiceType = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetIssueCompany(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.IssueCompany = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetIssueDate(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.IssueDate = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetItineraryNum(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.ItineraryNum = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetOfdOssUrl(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.OfdOssUrl = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetOrderId(v int64) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.OrderId = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetOssUrl(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.OssUrl = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetPassengerName(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.PassengerName = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetPdfOssUrl(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.PdfOssUrl = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetProject(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.Project = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetPromptMessage(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.PromptMessage = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetPurchaserName(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.PurchaserName = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetPurchaserTaxNo(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.PurchaserTaxNo = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetPurchaserType(v int32) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.PurchaserType = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetTaxAmount(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.TaxAmount = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetTaxRate(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.TaxRate = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetTicketNo(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.TicketNo = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetTicketPrice(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.TicketPrice = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetTotalPrice(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.TotalPrice = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) SetValidationCode(v string) *FlightItineraryScanQueryResponseBodyModuleItems {
	s.ValidationCode = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItems) Validate() error {
	return dara.Validate(s)
}

type FlightItineraryScanQueryResponseBodyModuleItemsFlights struct {
	// 航班至
	ArrivalStation *string `json:"arrival_station,omitempty" xml:"arrival_station,omitempty"`
	// 座位等级
	//
	// example:
	//
	// M
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// 承运人
	Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
	// 航班从
	DepartureStation *string `json:"departure_station,omitempty" xml:"departure_station,omitempty"`
	// 日期
	//
	// example:
	//
	// 2018-11-18
	FlightDate *string `json:"flight_date,omitempty" xml:"flight_date,omitempty"`
	// 航班号
	//
	// example:
	//
	// MU2271
	FlightNumber *string `json:"flight_number,omitempty" xml:"flight_number,omitempty"`
	// 时间
	//
	// example:
	//
	// 18:25
	FlightTime *string `json:"flight_time,omitempty" xml:"flight_time,omitempty"`
	// 免费行李
	//
	// example:
	//
	// 20K
	FreeBaggageAllowance *string `json:"free_baggage_allowance,omitempty" xml:"free_baggage_allowance,omitempty"`
	// 行号
	//
	// example:
	//
	// 1
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// 客票级别
	//
	// example:
	//
	// M
	SeatClass *string `json:"seat_class,omitempty" xml:"seat_class,omitempty"`
	// 客票生效日期
	//
	// example:
	//
	// 2023-01-01
	ValidFromDate *string `json:"valid_from_date,omitempty" xml:"valid_from_date,omitempty"`
	// 有效截止日期
	//
	// example:
	//
	// 2023-01-01
	ValidToDate *string `json:"valid_to_date,omitempty" xml:"valid_to_date,omitempty"`
}

func (s FlightItineraryScanQueryResponseBodyModuleItemsFlights) String() string {
	return dara.Prettify(s)
}

func (s FlightItineraryScanQueryResponseBodyModuleItemsFlights) GoString() string {
	return s.String()
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) GetArrivalStation() *string {
	return s.ArrivalStation
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) GetCarrier() *string {
	return s.Carrier
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) GetDepartureStation() *string {
	return s.DepartureStation
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) GetFlightDate() *string {
	return s.FlightDate
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) GetFlightNumber() *string {
	return s.FlightNumber
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) GetFlightTime() *string {
	return s.FlightTime
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) GetFreeBaggageAllowance() *string {
	return s.FreeBaggageAllowance
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) GetIndex() *string {
	return s.Index
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) GetSeatClass() *string {
	return s.SeatClass
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) GetValidFromDate() *string {
	return s.ValidFromDate
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) GetValidToDate() *string {
	return s.ValidToDate
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) SetArrivalStation(v string) *FlightItineraryScanQueryResponseBodyModuleItemsFlights {
	s.ArrivalStation = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) SetCabinClass(v string) *FlightItineraryScanQueryResponseBodyModuleItemsFlights {
	s.CabinClass = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) SetCarrier(v string) *FlightItineraryScanQueryResponseBodyModuleItemsFlights {
	s.Carrier = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) SetDepartureStation(v string) *FlightItineraryScanQueryResponseBodyModuleItemsFlights {
	s.DepartureStation = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) SetFlightDate(v string) *FlightItineraryScanQueryResponseBodyModuleItemsFlights {
	s.FlightDate = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) SetFlightNumber(v string) *FlightItineraryScanQueryResponseBodyModuleItemsFlights {
	s.FlightNumber = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) SetFlightTime(v string) *FlightItineraryScanQueryResponseBodyModuleItemsFlights {
	s.FlightTime = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) SetFreeBaggageAllowance(v string) *FlightItineraryScanQueryResponseBodyModuleItemsFlights {
	s.FreeBaggageAllowance = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) SetIndex(v string) *FlightItineraryScanQueryResponseBodyModuleItemsFlights {
	s.Index = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) SetSeatClass(v string) *FlightItineraryScanQueryResponseBodyModuleItemsFlights {
	s.SeatClass = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) SetValidFromDate(v string) *FlightItineraryScanQueryResponseBodyModuleItemsFlights {
	s.ValidFromDate = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) SetValidToDate(v string) *FlightItineraryScanQueryResponseBodyModuleItemsFlights {
	s.ValidToDate = &v
	return s
}

func (s *FlightItineraryScanQueryResponseBodyModuleItemsFlights) Validate() error {
	return dara.Validate(s)
}
