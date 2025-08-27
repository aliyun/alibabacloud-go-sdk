// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainTicketScanQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainTicketScanQueryResponseBody
	GetCode() *string
	SetMessage(v string) *TrainTicketScanQueryResponseBody
	GetMessage() *string
	SetModule(v *TrainTicketScanQueryResponseBodyModule) *TrainTicketScanQueryResponseBody
	GetModule() *TrainTicketScanQueryResponseBodyModule
	SetRequestId(v string) *TrainTicketScanQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainTicketScanQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainTicketScanQueryResponseBody
	GetTraceId() *string
}

type TrainTicketScanQueryResponseBody struct {
	// example:
	//
	// 200
	Code    *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	Module  *TrainTicketScanQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
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

func (s TrainTicketScanQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainTicketScanQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TrainTicketScanQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainTicketScanQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainTicketScanQueryResponseBody) GetModule() *TrainTicketScanQueryResponseBodyModule {
	return s.Module
}

func (s *TrainTicketScanQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainTicketScanQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainTicketScanQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainTicketScanQueryResponseBody) SetCode(v string) *TrainTicketScanQueryResponseBody {
	s.Code = &v
	return s
}

func (s *TrainTicketScanQueryResponseBody) SetMessage(v string) *TrainTicketScanQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TrainTicketScanQueryResponseBody) SetModule(v *TrainTicketScanQueryResponseBodyModule) *TrainTicketScanQueryResponseBody {
	s.Module = v
	return s
}

func (s *TrainTicketScanQueryResponseBody) SetRequestId(v string) *TrainTicketScanQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainTicketScanQueryResponseBody) SetSuccess(v bool) *TrainTicketScanQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TrainTicketScanQueryResponseBody) SetTraceId(v string) *TrainTicketScanQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainTicketScanQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainTicketScanQueryResponseBodyModule struct {
	Items []*TrainTicketScanQueryResponseBodyModuleItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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

func (s TrainTicketScanQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainTicketScanQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainTicketScanQueryResponseBodyModule) GetItems() []*TrainTicketScanQueryResponseBodyModuleItems {
	return s.Items
}

func (s *TrainTicketScanQueryResponseBodyModule) GetPageNo() *int32 {
	return s.PageNo
}

func (s *TrainTicketScanQueryResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *TrainTicketScanQueryResponseBodyModule) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *TrainTicketScanQueryResponseBodyModule) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *TrainTicketScanQueryResponseBodyModule) SetItems(v []*TrainTicketScanQueryResponseBodyModuleItems) *TrainTicketScanQueryResponseBodyModule {
	s.Items = v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModule) SetPageNo(v int32) *TrainTicketScanQueryResponseBodyModule {
	s.PageNo = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModule) SetPageSize(v int32) *TrainTicketScanQueryResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModule) SetTotalPage(v int32) *TrainTicketScanQueryResponseBodyModule {
	s.TotalPage = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModule) SetTotalSize(v int32) *TrainTicketScanQueryResponseBodyModule {
	s.TotalSize = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TrainTicketScanQueryResponseBodyModuleItems struct {
	ApplyId    *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ArrStation *string `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	// example:
	//
	// 2022-12-01
	BillDate   *string `json:"bill_date,omitempty" xml:"bill_date,omitempty"`
	CoachName  *string `json:"coach_name,omitempty" xml:"coach_name,omitempty"`
	CostCenter *string `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	DepStation *string `json:"dep_station,omitempty" xml:"dep_station,omitempty"`
	// example:
	//
	// 2023-01-12 10:00:00
	DepTime         *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	Department      *string `json:"department,omitempty" xml:"department,omitempty"`
	FeeTypeShowCode *int32  `json:"fee_type_show_code,omitempty" xml:"fee_type_show_code,omitempty"`
	// example:
	//
	// 71
	Id              *string `json:"id,omitempty" xml:"id,omitempty"`
	InvoiceDate     *string `json:"invoice_date,omitempty" xml:"invoice_date,omitempty"`
	InvoiceMaterial *int32  `json:"invoice_material,omitempty" xml:"invoice_material,omitempty"`
	InvoiceTitle    *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	OfdUrl          *string `json:"ofd_url,omitempty" xml:"ofd_url,omitempty"`
	// example:
	//
	// 3137168772101111000
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// https://www.testurl.com
	OssUrl    *string `json:"oss_url,omitempty" xml:"oss_url,omitempty"`
	Passenger *string `json:"passenger,omitempty" xml:"passenger,omitempty"`
	PdfUrl    *string `json:"pdf_url,omitempty" xml:"pdf_url,omitempty"`
	// example:
	//
	// 100
	Price          *string `json:"price,omitempty" xml:"price,omitempty"`
	Project        *string `json:"project,omitempty" xml:"project,omitempty"`
	PurchaserName  *string `json:"purchaser_name,omitempty" xml:"purchaser_name,omitempty"`
	PurchaserTaxNo *string `json:"purchaser_tax_no,omitempty" xml:"purchaser_tax_no,omitempty"`
	Seat           *string `json:"seat,omitempty" xml:"seat,omitempty"`
	SeatNo         *string `json:"seat_no,omitempty" xml:"seat_no,omitempty"`
	// example:
	//
	// 30671211200127U123456
	SerialNumber *string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
	// example:
	//
	// 8.26
	TaxAmount *string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
	// example:
	//
	// 9%
	TaxRate *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 取票号
	//
	// example:
	//
	// 784-1111111111
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 车次
	//
	// example:
	//
	// G99
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
}

func (s TrainTicketScanQueryResponseBodyModuleItems) String() string {
	return dara.Prettify(s)
}

func (s TrainTicketScanQueryResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetApplyId() *string {
	return s.ApplyId
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetArrStation() *string {
	return s.ArrStation
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetBillDate() *string {
	return s.BillDate
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetCoachName() *string {
	return s.CoachName
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetCostCenter() *string {
	return s.CostCenter
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetDepStation() *string {
	return s.DepStation
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetDepartment() *string {
	return s.Department
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetFeeTypeShowCode() *int32 {
	return s.FeeTypeShowCode
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetId() *string {
	return s.Id
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetInvoiceDate() *string {
	return s.InvoiceDate
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetInvoiceMaterial() *int32 {
	return s.InvoiceMaterial
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetOfdUrl() *string {
	return s.OfdUrl
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetOrderId() *int64 {
	return s.OrderId
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetOssUrl() *string {
	return s.OssUrl
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetPassenger() *string {
	return s.Passenger
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetPdfUrl() *string {
	return s.PdfUrl
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetPrice() *string {
	return s.Price
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetProject() *string {
	return s.Project
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetPurchaserName() *string {
	return s.PurchaserName
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetPurchaserTaxNo() *string {
	return s.PurchaserTaxNo
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetSeat() *string {
	return s.Seat
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetSeatNo() *string {
	return s.SeatNo
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetTaxAmount() *string {
	return s.TaxAmount
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetTaxRate() *string {
	return s.TaxRate
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetTicketNo() *string {
	return s.TicketNo
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetApplyId(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.ApplyId = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetArrStation(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.ArrStation = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetBillDate(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.BillDate = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetCoachName(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.CoachName = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetCostCenter(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.CostCenter = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetDepStation(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.DepStation = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetDepTime(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.DepTime = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetDepartment(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.Department = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetFeeTypeShowCode(v int32) *TrainTicketScanQueryResponseBodyModuleItems {
	s.FeeTypeShowCode = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetId(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.Id = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetInvoiceDate(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.InvoiceDate = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetInvoiceMaterial(v int32) *TrainTicketScanQueryResponseBodyModuleItems {
	s.InvoiceMaterial = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetInvoiceTitle(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.InvoiceTitle = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetOfdUrl(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.OfdUrl = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetOrderId(v int64) *TrainTicketScanQueryResponseBodyModuleItems {
	s.OrderId = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetOssUrl(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.OssUrl = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetPassenger(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.Passenger = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetPdfUrl(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.PdfUrl = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetPrice(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.Price = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetProject(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.Project = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetPurchaserName(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.PurchaserName = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetPurchaserTaxNo(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.PurchaserTaxNo = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetSeat(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.Seat = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetSeatNo(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.SeatNo = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetSerialNumber(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.SerialNumber = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetTaxAmount(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.TaxAmount = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetTaxRate(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.TaxRate = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetTicketNo(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.TicketNo = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) SetTrainNo(v string) *TrainTicketScanQueryResponseBodyModuleItems {
	s.TrainNo = &v
	return s
}

func (s *TrainTicketScanQueryResponseBodyModuleItems) Validate() error {
	return dara.Validate(s)
}
