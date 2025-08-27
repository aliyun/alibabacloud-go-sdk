// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptNoSeat(v string) *TrainOrderCreateRequest
	GetAcceptNoSeat() *string
	SetBookTrainInfos(v []*TrainOrderCreateRequestBookTrainInfos) *TrainOrderCreateRequest
	GetBookTrainInfos() []*TrainOrderCreateRequestBookTrainInfos
	SetBtripUserId(v string) *TrainOrderCreateRequest
	GetBtripUserId() *string
	SetBtripUserName(v string) *TrainOrderCreateRequest
	GetBtripUserName() *string
	SetBusinessInfo(v *TrainOrderCreateRequestBusinessInfo) *TrainOrderCreateRequest
	GetBusinessInfo() *TrainOrderCreateRequestBusinessInfo
	SetContactInfo(v *TrainOrderCreateRequestContactInfo) *TrainOrderCreateRequest
	GetContactInfo() *TrainOrderCreateRequestContactInfo
	SetForceMatch(v string) *TrainOrderCreateRequest
	GetForceMatch() *string
	SetIsPayNow(v bool) *TrainOrderCreateRequest
	GetIsPayNow() *bool
	SetOutOrderId(v string) *TrainOrderCreateRequest
	GetOutOrderId() *string
	SetPassengerOpenInfoS(v []*TrainOrderCreateRequestPassengerOpenInfoS) *TrainOrderCreateRequest
	GetPassengerOpenInfoS() []*TrainOrderCreateRequestPassengerOpenInfoS
}

type TrainOrderCreateRequest struct {
	// example:
	//
	// 0
	AcceptNoSeat *string `json:"accept_no_seat,omitempty" xml:"accept_no_seat,omitempty"`
	// This parameter is required.
	BookTrainInfos []*TrainOrderCreateRequestBookTrainInfos `json:"book_train_infos,omitempty" xml:"book_train_infos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 12344321
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// This parameter is required.
	BtripUserName *string                              `json:"btrip_user_name,omitempty" xml:"btrip_user_name,omitempty"`
	BusinessInfo  *TrainOrderCreateRequestBusinessInfo `json:"business_info,omitempty" xml:"business_info,omitempty" type:"Struct"`
	// This parameter is required.
	ContactInfo *TrainOrderCreateRequestContactInfo `json:"contact_info,omitempty" xml:"contact_info,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ForceMatch *string `json:"force_match,omitempty" xml:"force_match,omitempty"`
	// example:
	//
	// false
	IsPayNow *bool `json:"is_pay_now,omitempty" xml:"is_pay_now,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// This parameter is required.
	PassengerOpenInfoS []*TrainOrderCreateRequestPassengerOpenInfoS `json:"passenger_open_info_s,omitempty" xml:"passenger_open_info_s,omitempty" type:"Repeated"`
}

func (s TrainOrderCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCreateRequest) GoString() string {
	return s.String()
}

func (s *TrainOrderCreateRequest) GetAcceptNoSeat() *string {
	return s.AcceptNoSeat
}

func (s *TrainOrderCreateRequest) GetBookTrainInfos() []*TrainOrderCreateRequestBookTrainInfos {
	return s.BookTrainInfos
}

func (s *TrainOrderCreateRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *TrainOrderCreateRequest) GetBtripUserName() *string {
	return s.BtripUserName
}

func (s *TrainOrderCreateRequest) GetBusinessInfo() *TrainOrderCreateRequestBusinessInfo {
	return s.BusinessInfo
}

func (s *TrainOrderCreateRequest) GetContactInfo() *TrainOrderCreateRequestContactInfo {
	return s.ContactInfo
}

func (s *TrainOrderCreateRequest) GetForceMatch() *string {
	return s.ForceMatch
}

func (s *TrainOrderCreateRequest) GetIsPayNow() *bool {
	return s.IsPayNow
}

func (s *TrainOrderCreateRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainOrderCreateRequest) GetPassengerOpenInfoS() []*TrainOrderCreateRequestPassengerOpenInfoS {
	return s.PassengerOpenInfoS
}

func (s *TrainOrderCreateRequest) SetAcceptNoSeat(v string) *TrainOrderCreateRequest {
	s.AcceptNoSeat = &v
	return s
}

func (s *TrainOrderCreateRequest) SetBookTrainInfos(v []*TrainOrderCreateRequestBookTrainInfos) *TrainOrderCreateRequest {
	s.BookTrainInfos = v
	return s
}

func (s *TrainOrderCreateRequest) SetBtripUserId(v string) *TrainOrderCreateRequest {
	s.BtripUserId = &v
	return s
}

func (s *TrainOrderCreateRequest) SetBtripUserName(v string) *TrainOrderCreateRequest {
	s.BtripUserName = &v
	return s
}

func (s *TrainOrderCreateRequest) SetBusinessInfo(v *TrainOrderCreateRequestBusinessInfo) *TrainOrderCreateRequest {
	s.BusinessInfo = v
	return s
}

func (s *TrainOrderCreateRequest) SetContactInfo(v *TrainOrderCreateRequestContactInfo) *TrainOrderCreateRequest {
	s.ContactInfo = v
	return s
}

func (s *TrainOrderCreateRequest) SetForceMatch(v string) *TrainOrderCreateRequest {
	s.ForceMatch = &v
	return s
}

func (s *TrainOrderCreateRequest) SetIsPayNow(v bool) *TrainOrderCreateRequest {
	s.IsPayNow = &v
	return s
}

func (s *TrainOrderCreateRequest) SetOutOrderId(v string) *TrainOrderCreateRequest {
	s.OutOrderId = &v
	return s
}

func (s *TrainOrderCreateRequest) SetPassengerOpenInfoS(v []*TrainOrderCreateRequestPassengerOpenInfoS) *TrainOrderCreateRequest {
	s.PassengerOpenInfoS = v
	return s
}

func (s *TrainOrderCreateRequest) Validate() error {
	return dara.Validate(s)
}

type TrainOrderCreateRequestBookTrainInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// BDC
	ArrStationCode *string `json:"arr_station_code,omitempty" xml:"arr_station_code,omitempty"`
	// This parameter is required.
	BookTicketInfos []*TrainOrderCreateRequestBookTrainInfosBookTicketInfos `json:"book_ticket_infos,omitempty" xml:"book_ticket_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 1T
	ChooseBeds *string `json:"choose_beds,omitempty" xml:"choose_beds,omitempty"`
	// example:
	//
	// 1T
	ChooseSeats *string `json:"choose_seats,omitempty" xml:"choose_seats,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BTC
	DepStationCode *string `json:"dep_station_code,omitempty" xml:"dep_station_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-05-06 15:19:01
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// K123456
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
}

func (s TrainOrderCreateRequestBookTrainInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCreateRequestBookTrainInfos) GoString() string {
	return s.String()
}

func (s *TrainOrderCreateRequestBookTrainInfos) GetArrStationCode() *string {
	return s.ArrStationCode
}

func (s *TrainOrderCreateRequestBookTrainInfos) GetBookTicketInfos() []*TrainOrderCreateRequestBookTrainInfosBookTicketInfos {
	return s.BookTicketInfos
}

func (s *TrainOrderCreateRequestBookTrainInfos) GetChooseBeds() *string {
	return s.ChooseBeds
}

func (s *TrainOrderCreateRequestBookTrainInfos) GetChooseSeats() *string {
	return s.ChooseSeats
}

func (s *TrainOrderCreateRequestBookTrainInfos) GetDepStationCode() *string {
	return s.DepStationCode
}

func (s *TrainOrderCreateRequestBookTrainInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainOrderCreateRequestBookTrainInfos) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainOrderCreateRequestBookTrainInfos) SetArrStationCode(v string) *TrainOrderCreateRequestBookTrainInfos {
	s.ArrStationCode = &v
	return s
}

func (s *TrainOrderCreateRequestBookTrainInfos) SetBookTicketInfos(v []*TrainOrderCreateRequestBookTrainInfosBookTicketInfos) *TrainOrderCreateRequestBookTrainInfos {
	s.BookTicketInfos = v
	return s
}

func (s *TrainOrderCreateRequestBookTrainInfos) SetChooseBeds(v string) *TrainOrderCreateRequestBookTrainInfos {
	s.ChooseBeds = &v
	return s
}

func (s *TrainOrderCreateRequestBookTrainInfos) SetChooseSeats(v string) *TrainOrderCreateRequestBookTrainInfos {
	s.ChooseSeats = &v
	return s
}

func (s *TrainOrderCreateRequestBookTrainInfos) SetDepStationCode(v string) *TrainOrderCreateRequestBookTrainInfos {
	s.DepStationCode = &v
	return s
}

func (s *TrainOrderCreateRequestBookTrainInfos) SetDepTime(v string) *TrainOrderCreateRequestBookTrainInfos {
	s.DepTime = &v
	return s
}

func (s *TrainOrderCreateRequestBookTrainInfos) SetTrainNo(v string) *TrainOrderCreateRequestBookTrainInfos {
	s.TrainNo = &v
	return s
}

func (s *TrainOrderCreateRequestBookTrainInfos) Validate() error {
	return dara.Validate(s)
}

type TrainOrderCreateRequestBookTrainInfosBookTicketInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	PassengerId *string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 14
	SeatType *string `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	TicketPrice *int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	TicketType *string `json:"ticket_type,omitempty" xml:"ticket_type,omitempty"`
}

func (s TrainOrderCreateRequestBookTrainInfosBookTicketInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCreateRequestBookTrainInfosBookTicketInfos) GoString() string {
	return s.String()
}

func (s *TrainOrderCreateRequestBookTrainInfosBookTicketInfos) GetPassengerId() *string {
	return s.PassengerId
}

func (s *TrainOrderCreateRequestBookTrainInfosBookTicketInfos) GetSeatType() *string {
	return s.SeatType
}

func (s *TrainOrderCreateRequestBookTrainInfosBookTicketInfos) GetTicketPrice() *int64 {
	return s.TicketPrice
}

func (s *TrainOrderCreateRequestBookTrainInfosBookTicketInfos) GetTicketType() *string {
	return s.TicketType
}

func (s *TrainOrderCreateRequestBookTrainInfosBookTicketInfos) SetPassengerId(v string) *TrainOrderCreateRequestBookTrainInfosBookTicketInfos {
	s.PassengerId = &v
	return s
}

func (s *TrainOrderCreateRequestBookTrainInfosBookTicketInfos) SetSeatType(v string) *TrainOrderCreateRequestBookTrainInfosBookTicketInfos {
	s.SeatType = &v
	return s
}

func (s *TrainOrderCreateRequestBookTrainInfosBookTicketInfos) SetTicketPrice(v int64) *TrainOrderCreateRequestBookTrainInfosBookTicketInfos {
	s.TicketPrice = &v
	return s
}

func (s *TrainOrderCreateRequestBookTrainInfosBookTicketInfos) SetTicketType(v string) *TrainOrderCreateRequestBookTrainInfosBookTicketInfos {
	s.TicketType = &v
	return s
}

func (s *TrainOrderCreateRequestBookTrainInfosBookTicketInfos) Validate() error {
	return dara.Validate(s)
}

type TrainOrderCreateRequestBusinessInfo struct {
	// example:
	//
	// 4321
	CustomerApplyId *string `json:"customer_apply_id,omitempty" xml:"customer_apply_id,omitempty"`
	// example:
	//
	// 1234
	CustomerItineraryId *string `json:"customer_itinerary_id,omitempty" xml:"customer_itinerary_id,omitempty"`
}

func (s TrainOrderCreateRequestBusinessInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCreateRequestBusinessInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderCreateRequestBusinessInfo) GetCustomerApplyId() *string {
	return s.CustomerApplyId
}

func (s *TrainOrderCreateRequestBusinessInfo) GetCustomerItineraryId() *string {
	return s.CustomerItineraryId
}

func (s *TrainOrderCreateRequestBusinessInfo) SetCustomerApplyId(v string) *TrainOrderCreateRequestBusinessInfo {
	s.CustomerApplyId = &v
	return s
}

func (s *TrainOrderCreateRequestBusinessInfo) SetCustomerItineraryId(v string) *TrainOrderCreateRequestBusinessInfo {
	s.CustomerItineraryId = &v
	return s
}

func (s *TrainOrderCreateRequestBusinessInfo) Validate() error {
	return dara.Validate(s)
}

type TrainOrderCreateRequestContactInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	PassengerId *string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b6a6fc1bdf1ba60e25c2e132b612c8819
	PassengerMobile *string `json:"passenger_mobile,omitempty" xml:"passenger_mobile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 949c9f34f677a0e5d249dfc94f5e62cc7
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
}

func (s TrainOrderCreateRequestContactInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCreateRequestContactInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderCreateRequestContactInfo) GetPassengerId() *string {
	return s.PassengerId
}

func (s *TrainOrderCreateRequestContactInfo) GetPassengerMobile() *string {
	return s.PassengerMobile
}

func (s *TrainOrderCreateRequestContactInfo) GetPassengerName() *string {
	return s.PassengerName
}

func (s *TrainOrderCreateRequestContactInfo) SetPassengerId(v string) *TrainOrderCreateRequestContactInfo {
	s.PassengerId = &v
	return s
}

func (s *TrainOrderCreateRequestContactInfo) SetPassengerMobile(v string) *TrainOrderCreateRequestContactInfo {
	s.PassengerMobile = &v
	return s
}

func (s *TrainOrderCreateRequestContactInfo) SetPassengerName(v string) *TrainOrderCreateRequestContactInfo {
	s.PassengerName = &v
	return s
}

func (s *TrainOrderCreateRequestContactInfo) Validate() error {
	return dara.Validate(s)
}

type TrainOrderCreateRequestPassengerOpenInfoS struct {
	CostCenterInfo *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo `json:"cost_center_info,omitempty" xml:"cost_center_info,omitempty" type:"Struct"`
	// example:
	//
	// 291487e553c5abde3b611aae283e2526f0d733ab55094aadc0b5ba587222a233c
	CountryCode *string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 291487e553c5abde3b611aae283e2526f0d733ab55094aadc0b5ba587222a233c
	PassengerCertNo *string `json:"passenger_cert_no,omitempty" xml:"passenger_cert_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 949c9f34f677a0e5d249dfc94f5e62cc7
	PassengerCertType *string `json:"passenger_cert_type,omitempty" xml:"passenger_cert_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	PassengerId *string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// example:
	//
	// b6a6fc1bdf1ba60e25c2e132b612c8819
	PassengerMobile *string `json:"passenger_mobile,omitempty" xml:"passenger_mobile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 949c9f34f677a0e5d249dfc94f5e62cc7
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// example:
	//
	// 291487e553c5abde3b611aae283e2526f0d733ab55094aadc0b5ba587222a233c
	ValidDateEnd *string `json:"valid_date_end,omitempty" xml:"valid_date_end,omitempty"`
}

func (s TrainOrderCreateRequestPassengerOpenInfoS) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCreateRequestPassengerOpenInfoS) GoString() string {
	return s.String()
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) GetCostCenterInfo() *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo {
	return s.CostCenterInfo
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) GetCountryCode() *string {
	return s.CountryCode
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) GetPassengerCertNo() *string {
	return s.PassengerCertNo
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) GetPassengerCertType() *string {
	return s.PassengerCertType
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) GetPassengerId() *string {
	return s.PassengerId
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) GetPassengerMobile() *string {
	return s.PassengerMobile
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) GetPassengerName() *string {
	return s.PassengerName
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) GetValidDateEnd() *string {
	return s.ValidDateEnd
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) SetCostCenterInfo(v *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) *TrainOrderCreateRequestPassengerOpenInfoS {
	s.CostCenterInfo = v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) SetCountryCode(v string) *TrainOrderCreateRequestPassengerOpenInfoS {
	s.CountryCode = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) SetPassengerCertNo(v string) *TrainOrderCreateRequestPassengerOpenInfoS {
	s.PassengerCertNo = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) SetPassengerCertType(v string) *TrainOrderCreateRequestPassengerOpenInfoS {
	s.PassengerCertType = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) SetPassengerId(v string) *TrainOrderCreateRequestPassengerOpenInfoS {
	s.PassengerId = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) SetPassengerMobile(v string) *TrainOrderCreateRequestPassengerOpenInfoS {
	s.PassengerMobile = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) SetPassengerName(v string) *TrainOrderCreateRequestPassengerOpenInfoS {
	s.PassengerName = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) SetValidDateEnd(v string) *TrainOrderCreateRequestPassengerOpenInfoS {
	s.ValidDateEnd = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoS) Validate() error {
	return dara.Validate(s)
}

type TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo struct {
	CascadeDeptName *string `json:"cascade_dept_name,omitempty" xml:"cascade_dept_name,omitempty"`
	// example:
	//
	// 123321
	CostCenterId   *string `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	// example:
	//
	// 12332112344
	CostCenterNo *string `json:"cost_center_no,omitempty" xml:"cost_center_no,omitempty"`
	// example:
	//
	// 010000009
	DepartId   *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// example:
	//
	// 111111
	InvoiceId    *string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	InvoiceTitle *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// example:
	//
	// acs
	ProjectCode  *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
}

func (s TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) GetCascadeDeptName() *string {
	return s.CascadeDeptName
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) GetCostCenterId() *string {
	return s.CostCenterId
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) GetCostCenterNo() *string {
	return s.CostCenterNo
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) GetDepartId() *string {
	return s.DepartId
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) GetDepartName() *string {
	return s.DepartName
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) GetInvoiceId() *string {
	return s.InvoiceId
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) SetCascadeDeptName(v string) *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo {
	s.CascadeDeptName = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) SetCostCenterId(v string) *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo {
	s.CostCenterId = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) SetCostCenterName(v string) *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo {
	s.CostCenterName = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) SetCostCenterNo(v string) *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo {
	s.CostCenterNo = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) SetDepartId(v string) *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo {
	s.DepartId = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) SetDepartName(v string) *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo {
	s.DepartName = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) SetInvoiceId(v string) *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo {
	s.InvoiceId = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) SetInvoiceTitle(v string) *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo {
	s.InvoiceTitle = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) SetProjectCode(v string) *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo {
	s.ProjectCode = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) SetProjectTitle(v string) *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo {
	s.ProjectTitle = &v
	return s
}

func (s *TrainOrderCreateRequestPassengerOpenInfoSCostCenterInfo) Validate() error {
	return dara.Validate(s)
}
