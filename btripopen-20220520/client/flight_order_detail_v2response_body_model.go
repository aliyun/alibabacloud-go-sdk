// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderDetailV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightOrderDetailV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FlightOrderDetailV2ResponseBody
	GetMessage() *string
	SetModule(v *FlightOrderDetailV2ResponseBodyModule) *FlightOrderDetailV2ResponseBody
	GetModule() *FlightOrderDetailV2ResponseBodyModule
	SetRequestId(v string) *FlightOrderDetailV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightOrderDetailV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightOrderDetailV2ResponseBody
	GetTraceId() *string
}

type FlightOrderDetailV2ResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *FlightOrderDetailV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// E5F4ACF5-5677-1515-9999-ABBB5E668032
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210bc60a16917251281873772dac41
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightOrderDetailV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightOrderDetailV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightOrderDetailV2ResponseBody) GetModule() *FlightOrderDetailV2ResponseBodyModule {
	return s.Module
}

func (s *FlightOrderDetailV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightOrderDetailV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightOrderDetailV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightOrderDetailV2ResponseBody) SetCode(v string) *FlightOrderDetailV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBody) SetMessage(v string) *FlightOrderDetailV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBody) SetModule(v *FlightOrderDetailV2ResponseBodyModule) *FlightOrderDetailV2ResponseBody {
	s.Module = v
	return s
}

func (s *FlightOrderDetailV2ResponseBody) SetRequestId(v string) *FlightOrderDetailV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBody) SetSuccess(v bool) *FlightOrderDetailV2ResponseBody {
	s.Success = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBody) SetTraceId(v string) *FlightOrderDetailV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightOrderDetailV2ResponseBodyModule struct {
	B2gVipCode *string `json:"b2g_vip_code,omitempty" xml:"b2g_vip_code,omitempty"`
	// example:
	//
	// 2023-06-29 15:28:44
	BookSuccTime *string `json:"book_succ_time,omitempty" xml:"book_succ_time,omitempty"`
	// example:
	//
	// qingg1234
	BookUserId   *string `json:"book_user_id,omitempty" xml:"book_user_id,omitempty"`
	BookUserName *string `json:"book_user_name,omitempty" xml:"book_user_name,omitempty"`
	// example:
	//
	// 0
	BuildPrice     *int64                                               `json:"build_price,omitempty" xml:"build_price,omitempty"`
	ContactInfoDTO *FlightOrderDetailV2ResponseBodyModuleContactInfoDTO `json:"contact_info_d_t_o,omitempty" xml:"contact_info_d_t_o,omitempty" type:"Struct"`
	// example:
	//
	// 2023-06-29 15:03:11
	CreateTime *string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// example:
	//
	// 126000
	Facevalue         *int64                                                  `json:"facevalue,omitempty" xml:"facevalue,omitempty"`
	FlightTaleInfoDTO *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO `json:"flight_tale_info_d_t_o,omitempty" xml:"flight_tale_info_d_t_o,omitempty" type:"Struct"`
	// example:
	//
	// false
	IsProtocol *bool `json:"is_protocol,omitempty" xml:"is_protocol,omitempty"`
	// example:
	//
	// false
	Isemergency *bool `json:"isemergency,omitempty" xml:"isemergency,omitempty"`
	// example:
	//
	// true
	Issendmessage *bool `json:"issendmessage,omitempty" xml:"issendmessage,omitempty"`
	// example:
	//
	// 4000
	OilPrice *int64 `json:"oil_price,omitempty" xml:"oil_price,omitempty"`
	// example:
	//
	// 1017002195370467200
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 8500
	OrderPrice *int64 `json:"order_price,omitempty" xml:"order_price,omitempty"`
	// example:
	//
	// 1017002195370467137
	OutOrderId    *string                                               `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	PassengerList []*FlightOrderDetailV2ResponseBodyModulePassengerList `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	// key :passengerId
	//
	// value :segmentId
	PassengerSegmentMap map[string]*string `json:"passenger_segment_map,omitempty" xml:"passenger_segment_map,omitempty"`
	// example:
	//
	// 2023-06-29 15:03:59
	PayTime *string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// example:
	//
	// 4500
	Saleprice *int64 `json:"saleprice,omitempty" xml:"saleprice,omitempty"`
	// example:
	//
	// true
	Sendcpsms *bool `json:"sendcpsms,omitempty" xml:"sendcpsms,omitempty"`
	// example:
	//
	// 5
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 3
	TotalServiceFeePrice *int64 `json:"total_service_fee_price,omitempty" xml:"total_service_fee_price,omitempty"`
}

func (s FlightOrderDetailV2ResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailV2ResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetB2gVipCode() *string {
	return s.B2gVipCode
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetBookSuccTime() *string {
	return s.BookSuccTime
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetBookUserId() *string {
	return s.BookUserId
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetBookUserName() *string {
	return s.BookUserName
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetBuildPrice() *int64 {
	return s.BuildPrice
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetContactInfoDTO() *FlightOrderDetailV2ResponseBodyModuleContactInfoDTO {
	return s.ContactInfoDTO
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetCreateTime() *string {
	return s.CreateTime
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetFacevalue() *int64 {
	return s.Facevalue
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetFlightTaleInfoDTO() *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO {
	return s.FlightTaleInfoDTO
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetIsProtocol() *bool {
	return s.IsProtocol
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetIsemergency() *bool {
	return s.Isemergency
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetIssendmessage() *bool {
	return s.Issendmessage
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetOilPrice() *int64 {
	return s.OilPrice
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetOrderPrice() *int64 {
	return s.OrderPrice
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetPassengerList() []*FlightOrderDetailV2ResponseBodyModulePassengerList {
	return s.PassengerList
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetPassengerSegmentMap() map[string]*string {
	return s.PassengerSegmentMap
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetPayTime() *string {
	return s.PayTime
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetSaleprice() *int64 {
	return s.Saleprice
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetSendcpsms() *bool {
	return s.Sendcpsms
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *FlightOrderDetailV2ResponseBodyModule) GetTotalServiceFeePrice() *int64 {
	return s.TotalServiceFeePrice
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetB2gVipCode(v string) *FlightOrderDetailV2ResponseBodyModule {
	s.B2gVipCode = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetBookSuccTime(v string) *FlightOrderDetailV2ResponseBodyModule {
	s.BookSuccTime = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetBookUserId(v string) *FlightOrderDetailV2ResponseBodyModule {
	s.BookUserId = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetBookUserName(v string) *FlightOrderDetailV2ResponseBodyModule {
	s.BookUserName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetBuildPrice(v int64) *FlightOrderDetailV2ResponseBodyModule {
	s.BuildPrice = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetContactInfoDTO(v *FlightOrderDetailV2ResponseBodyModuleContactInfoDTO) *FlightOrderDetailV2ResponseBodyModule {
	s.ContactInfoDTO = v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetCreateTime(v string) *FlightOrderDetailV2ResponseBodyModule {
	s.CreateTime = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetFacevalue(v int64) *FlightOrderDetailV2ResponseBodyModule {
	s.Facevalue = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetFlightTaleInfoDTO(v *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO) *FlightOrderDetailV2ResponseBodyModule {
	s.FlightTaleInfoDTO = v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetIsProtocol(v bool) *FlightOrderDetailV2ResponseBodyModule {
	s.IsProtocol = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetIsemergency(v bool) *FlightOrderDetailV2ResponseBodyModule {
	s.Isemergency = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetIssendmessage(v bool) *FlightOrderDetailV2ResponseBodyModule {
	s.Issendmessage = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetOilPrice(v int64) *FlightOrderDetailV2ResponseBodyModule {
	s.OilPrice = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetOrderId(v int64) *FlightOrderDetailV2ResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetOrderPrice(v int64) *FlightOrderDetailV2ResponseBodyModule {
	s.OrderPrice = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetOutOrderId(v string) *FlightOrderDetailV2ResponseBodyModule {
	s.OutOrderId = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetPassengerList(v []*FlightOrderDetailV2ResponseBodyModulePassengerList) *FlightOrderDetailV2ResponseBodyModule {
	s.PassengerList = v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetPassengerSegmentMap(v map[string]*string) *FlightOrderDetailV2ResponseBodyModule {
	s.PassengerSegmentMap = v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetPayTime(v string) *FlightOrderDetailV2ResponseBodyModule {
	s.PayTime = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetSaleprice(v int64) *FlightOrderDetailV2ResponseBodyModule {
	s.Saleprice = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetSendcpsms(v bool) *FlightOrderDetailV2ResponseBodyModule {
	s.Sendcpsms = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetStatus(v int32) *FlightOrderDetailV2ResponseBodyModule {
	s.Status = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) SetTotalServiceFeePrice(v int64) *FlightOrderDetailV2ResponseBodyModule {
	s.TotalServiceFeePrice = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightOrderDetailV2ResponseBodyModuleContactInfoDTO struct {
	// example:
	//
	// 178169630111@163.com
	ContactEmail *string `json:"contact_email,omitempty" xml:"contact_email,omitempty"`
	ContactName  *string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// example:
	//
	// 178169630111
	ContactPhone *string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// example:
	//
	// true
	SendMsgToPassenger *bool `json:"send_msg_to_passenger,omitempty" xml:"send_msg_to_passenger,omitempty"`
}

func (s FlightOrderDetailV2ResponseBodyModuleContactInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailV2ResponseBodyModuleContactInfoDTO) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailV2ResponseBodyModuleContactInfoDTO) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *FlightOrderDetailV2ResponseBodyModuleContactInfoDTO) GetContactName() *string {
	return s.ContactName
}

func (s *FlightOrderDetailV2ResponseBodyModuleContactInfoDTO) GetContactPhone() *string {
	return s.ContactPhone
}

func (s *FlightOrderDetailV2ResponseBodyModuleContactInfoDTO) GetSendMsgToPassenger() *bool {
	return s.SendMsgToPassenger
}

func (s *FlightOrderDetailV2ResponseBodyModuleContactInfoDTO) SetContactEmail(v string) *FlightOrderDetailV2ResponseBodyModuleContactInfoDTO {
	s.ContactEmail = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleContactInfoDTO) SetContactName(v string) *FlightOrderDetailV2ResponseBodyModuleContactInfoDTO {
	s.ContactName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleContactInfoDTO) SetContactPhone(v string) *FlightOrderDetailV2ResponseBodyModuleContactInfoDTO {
	s.ContactPhone = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleContactInfoDTO) SetSendMsgToPassenger(v bool) *FlightOrderDetailV2ResponseBodyModuleContactInfoDTO {
	s.SendMsgToPassenger = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleContactInfoDTO) Validate() error {
	return dara.Validate(s)
}

type FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO struct {
	Journeys   []*FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys `json:"journeys,omitempty" xml:"journeys,omitempty" type:"Repeated"`
	NoticeTips *string                                                           `json:"notice_tips,omitempty" xml:"notice_tips,omitempty"`
	TripType   *string                                                           `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	// example:
	//
	// 2
	TripTypeCode *int32 `json:"trip_type_code,omitempty" xml:"trip_type_code,omitempty"`
}

func (s FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO) GetJourneys() []*FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	return s.Journeys
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO) GetNoticeTips() *string {
	return s.NoticeTips
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO) GetTripType() *string {
	return s.TripType
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO) GetTripTypeCode() *int32 {
	return s.TripTypeCode
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO) SetJourneys(v []*FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO {
	s.Journeys = v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO) SetNoticeTips(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO {
	s.NoticeTips = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO) SetTripType(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO {
	s.TripType = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO) SetTripTypeCode(v int32) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO {
	s.TripTypeCode = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTO) Validate() error {
	return dara.Validate(s)
}

type FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys struct {
	// example:
	//
	// 85
	AllFlyDuration *int64 `json:"all_fly_duration,omitempty" xml:"all_fly_duration,omitempty"`
	// example:
	//
	// 85
	AllFlyDurationAfterChange *int64 `json:"all_fly_duration_after_change,omitempty" xml:"all_fly_duration_after_change,omitempty"`
	ApplyId                   *int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// example:
	//
	// XIL
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2023-07-20 08:25:00
	ArrTime        *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	BaggageDetails *string `json:"baggage_details,omitempty" xml:"baggage_details,omitempty"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// example:
	//
	// 2023-07-20 07:00:00
	DepTime      *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	FlightStatus *string `json:"flight_status,omitempty" xml:"flight_status,omitempty"`
	// iata_no
	//
	// example:
	//
	// iata_no
	IataNo              *string                                                                      `json:"iata_no,omitempty" xml:"iata_no,omitempty"`
	IsReshopJourney     *bool                                                                        `json:"is_reshop_journey,omitempty" xml:"is_reshop_journey,omitempty"`
	IsTransfer          *bool                                                                        `json:"is_transfer,omitempty" xml:"is_transfer,omitempty"`
	JourneyTitle        *string                                                                      `json:"journey_title,omitempty" xml:"journey_title,omitempty"`
	RefundChangeDetails *string                                                                      `json:"refund_change_details,omitempty" xml:"refund_change_details,omitempty"`
	SegmentList         []*FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
}

func (s FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetAllFlyDuration() *int64 {
	return s.AllFlyDuration
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetAllFlyDurationAfterChange() *int64 {
	return s.AllFlyDurationAfterChange
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetBaggageDetails() *string {
	return s.BaggageDetails
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetDepCityName() *string {
	return s.DepCityName
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetFlightStatus() *string {
	return s.FlightStatus
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetIataNo() *string {
	return s.IataNo
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetIsReshopJourney() *bool {
	return s.IsReshopJourney
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetIsTransfer() *bool {
	return s.IsTransfer
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetJourneyTitle() *string {
	return s.JourneyTitle
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetRefundChangeDetails() *string {
	return s.RefundChangeDetails
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) GetSegmentList() []*FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	return s.SegmentList
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetAllFlyDuration(v int64) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.AllFlyDuration = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetAllFlyDurationAfterChange(v int64) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.AllFlyDurationAfterChange = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetApplyId(v int64) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.ApplyId = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetArrCityCode(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.ArrCityCode = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetArrCityName(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.ArrCityName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetArrTime(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.ArrTime = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetBaggageDetails(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.BaggageDetails = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetDepCityCode(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.DepCityCode = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetDepCityName(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.DepCityName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetDepTime(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.DepTime = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetFlightStatus(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.FlightStatus = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetIataNo(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.IataNo = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetIsReshopJourney(v bool) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.IsReshopJourney = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetIsTransfer(v bool) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.IsTransfer = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetJourneyTitle(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.JourneyTitle = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetRefundChangeDetails(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.RefundChangeDetails = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) SetSegmentList(v []*FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys {
	s.SegmentList = v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneys) Validate() error {
	return dara.Validate(s)
}

type FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList struct {
	// example:
	//
	// CA
	AirLineCode *string `json:"air_line_code,omitempty" xml:"air_line_code,omitempty"`
	// example:
	//
	// Air China Limited
	AirLineEnglishName *string `json:"air_line_english_name,omitempty" xml:"air_line_english_name,omitempty"`
	AirLineName        *string `json:"air_line_name,omitempty" xml:"air_line_name,omitempty"`
	// example:
	//
	// 95583
	AirLinePhone *string `json:"air_line_phone,omitempty" xml:"air_line_phone,omitempty"`
	// example:
	//
	// https://gw.alicdn.com/tfs/TB12fJAFHr1gK0jSZR0XXbP8XXa-450-450.png
	AirlineIconUrl   *string `json:"airline_icon_url,omitempty" xml:"airline_icon_url,omitempty"`
	AirlineShortName *string `json:"airline_short_name,omitempty" xml:"airline_short_name,omitempty"`
	// example:
	//
	// XIL
	ArrAirportCode *string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	ArrAirportName *string `json:"arr_airport_name,omitempty" xml:"arr_airport_name,omitempty"`
	// example:
	//
	// XIL
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2023-07-20 08:25:00
	ArrTime        *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	ArriveTerminal *string `json:"arrive_terminal,omitempty" xml:"arrive_terminal,omitempty"`
	// cabin
	Cabin            *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	CabinAndDiscount *string `json:"cabin_and_discount,omitempty" xml:"cabin_and_discount,omitempty"`
	// cabin_class
	//
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// cabin_class_name
	CabinClassName *string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	// example:
	//
	// false
	CodeShare    *bool   `json:"code_share,omitempty" xml:"code_share,omitempty"`
	DeadlineText *string `json:"deadline_text,omitempty" xml:"deadline_text,omitempty"`
	// example:
	//
	// PEK
	DepAirportCode *string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	DepAirportName *string `json:"dep_airport_name,omitempty" xml:"dep_airport_name,omitempty"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// example:
	//
	// 2023-07-29
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// example:
	//
	// 2023-07-20 07:00:00
	DepTime        *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	DepartTerminal *string `json:"depart_terminal,omitempty" xml:"depart_terminal,omitempty"`
	// example:
	//
	// 0.4
	Discount     *float64                                                                               `json:"discount,omitempty" xml:"discount,omitempty"`
	FlightChange *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange `json:"flight_change,omitempty" xml:"flight_change,omitempty" type:"Struct"`
	// example:
	//
	// CA1110
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// ARJ
	FlightType                  *string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	FlyDuration                 *int32  `json:"fly_duration,omitempty" xml:"fly_duration,omitempty"`
	Manufacturer                *string `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	MealDesc                    *string `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	OnTimeRate                  *string `json:"on_time_rate,omitempty" xml:"on_time_rate,omitempty"`
	OperatingAirShortName       *string `json:"operating_air_short_name,omitempty" xml:"operating_air_short_name,omitempty"`
	OperatingAirlineCode        *string `json:"operating_airline_code,omitempty" xml:"operating_airline_code,omitempty"`
	OperatingAirlineEnglishName *string `json:"operating_airline_english_name,omitempty" xml:"operating_airline_english_name,omitempty"`
	OperatingAirlineIconUrl     *string `json:"operating_airline_icon_url,omitempty" xml:"operating_airline_icon_url,omitempty"`
	OperatingAirlineName        *string `json:"operating_airline_name,omitempty" xml:"operating_airline_name,omitempty"`
	OperatingAirlinePhone       *string `json:"operating_airline_phone,omitempty" xml:"operating_airline_phone,omitempty"`
	OperatingFlightNo           *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	PlaneType                   *string `json:"plane_type,omitempty" xml:"plane_type,omitempty"`
	RaisePrice                  *int64  `json:"raise_price,omitempty" xml:"raise_price,omitempty"`
	SegmentId                   *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// segmentIndex
	//
	// example:
	//
	// 0
	SegmentIndex    *int32                                                                                    `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	SegmentPosition *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
	StopAirport     *string                                                                                   `json:"stop_airport,omitempty" xml:"stop_airport,omitempty"`
	StopArrTime     *string                                                                                   `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	StopCity        *string                                                                                   `json:"stop_city,omitempty" xml:"stop_city,omitempty"`
	StopCityName    *string                                                                                   `json:"stop_city_name,omitempty" xml:"stop_city_name,omitempty"`
	StopDepTime     *string                                                                                   `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
	StopQuantity    *int32                                                                                    `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetAirLineCode() *string {
	return s.AirLineCode
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetAirLineEnglishName() *string {
	return s.AirLineEnglishName
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetAirLineName() *string {
	return s.AirLineName
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetAirLinePhone() *string {
	return s.AirLinePhone
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetAirlineIconUrl() *string {
	return s.AirlineIconUrl
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetAirlineShortName() *string {
	return s.AirlineShortName
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetArrAirportName() *string {
	return s.ArrAirportName
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetArriveTerminal() *string {
	return s.ArriveTerminal
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetCabin() *string {
	return s.Cabin
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetCabinAndDiscount() *string {
	return s.CabinAndDiscount
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetCodeShare() *bool {
	return s.CodeShare
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetDeadlineText() *string {
	return s.DeadlineText
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetDepAirportName() *string {
	return s.DepAirportName
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetDepCityName() *string {
	return s.DepCityName
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetDepDate() *string {
	return s.DepDate
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetDepartTerminal() *string {
	return s.DepartTerminal
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetDiscount() *float64 {
	return s.Discount
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetFlightChange() *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange {
	return s.FlightChange
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetFlightType() *string {
	return s.FlightType
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetFlyDuration() *int32 {
	return s.FlyDuration
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetMealDesc() *string {
	return s.MealDesc
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetOnTimeRate() *string {
	return s.OnTimeRate
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetOperatingAirShortName() *string {
	return s.OperatingAirShortName
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetOperatingAirlineCode() *string {
	return s.OperatingAirlineCode
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetOperatingAirlineEnglishName() *string {
	return s.OperatingAirlineEnglishName
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetOperatingAirlineIconUrl() *string {
	return s.OperatingAirlineIconUrl
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetOperatingAirlineName() *string {
	return s.OperatingAirlineName
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetOperatingAirlinePhone() *string {
	return s.OperatingAirlinePhone
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetPlaneType() *string {
	return s.PlaneType
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetRaisePrice() *int64 {
	return s.RaisePrice
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetSegmentId() *string {
	return s.SegmentId
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetSegmentPosition() *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListSegmentPosition {
	return s.SegmentPosition
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetStopAirport() *string {
	return s.StopAirport
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetStopCity() *string {
	return s.StopCity
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetStopCityName() *string {
	return s.StopCityName
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) GetStopQuantity() *int32 {
	return s.StopQuantity
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetAirLineCode(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.AirLineCode = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetAirLineEnglishName(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.AirLineEnglishName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetAirLineName(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.AirLineName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetAirLinePhone(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.AirLinePhone = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetAirlineIconUrl(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.AirlineIconUrl = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetAirlineShortName(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.AirlineShortName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetArrAirportCode(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.ArrAirportCode = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetArrAirportName(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.ArrAirportName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetArrCityCode(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.ArrCityCode = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetArrCityName(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.ArrCityName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetArrTime(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.ArrTime = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetArriveTerminal(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.ArriveTerminal = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetCabin(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.Cabin = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetCabinAndDiscount(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.CabinAndDiscount = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetCabinClass(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.CabinClass = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetCabinClassName(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.CabinClassName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetCodeShare(v bool) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetDeadlineText(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.DeadlineText = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetDepAirportCode(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.DepAirportCode = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetDepAirportName(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.DepAirportName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetDepCityCode(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.DepCityCode = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetDepCityName(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.DepCityName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetDepDate(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.DepDate = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetDepTime(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.DepTime = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetDepartTerminal(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.DepartTerminal = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetDiscount(v float64) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.Discount = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetFlightChange(v *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.FlightChange = v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetFlightNo(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.FlightNo = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetFlightType(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.FlightType = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetFlyDuration(v int32) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.FlyDuration = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetManufacturer(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.Manufacturer = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetMealDesc(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.MealDesc = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetOnTimeRate(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.OnTimeRate = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetOperatingAirShortName(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.OperatingAirShortName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetOperatingAirlineCode(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.OperatingAirlineCode = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetOperatingAirlineEnglishName(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.OperatingAirlineEnglishName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetOperatingAirlineIconUrl(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.OperatingAirlineIconUrl = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetOperatingAirlineName(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.OperatingAirlineName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetOperatingAirlinePhone(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.OperatingAirlinePhone = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetOperatingFlightNo(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetPlaneType(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.PlaneType = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetRaisePrice(v int64) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.RaisePrice = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetSegmentId(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.SegmentId = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetSegmentIndex(v int32) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.SegmentIndex = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetSegmentPosition(v *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListSegmentPosition) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.SegmentPosition = v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetStopAirport(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.StopAirport = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetStopArrTime(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.StopArrTime = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetStopCity(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.StopCity = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetStopCityName(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.StopCityName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetStopDepTime(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.StopDepTime = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) SetStopQuantity(v int32) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList {
	s.StopQuantity = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange struct {
	ChangeDesc       *string     `json:"change_desc,omitempty" xml:"change_desc,omitempty"`
	ChangeStatus     *string     `json:"change_status,omitempty" xml:"change_status,omitempty"`
	ChangeStatusCode *string     `json:"change_status_code,omitempty" xml:"change_status_code,omitempty"`
	NewSegment       interface{} `json:"new_segment,omitempty" xml:"new_segment,omitempty"`
	PassengerNames   []*string   `json:"passenger_names,omitempty" xml:"passenger_names,omitempty" type:"Repeated"`
}

func (s FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange) GetChangeDesc() *string {
	return s.ChangeDesc
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange) GetChangeStatus() *string {
	return s.ChangeStatus
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange) GetChangeStatusCode() *string {
	return s.ChangeStatusCode
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange) GetNewSegment() interface{} {
	return s.NewSegment
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange) GetPassengerNames() []*string {
	return s.PassengerNames
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange) SetChangeDesc(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange {
	s.ChangeDesc = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange) SetChangeStatus(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange {
	s.ChangeStatus = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange) SetChangeStatusCode(v string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange {
	s.ChangeStatusCode = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange) SetNewSegment(v interface{}) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange {
	s.NewSegment = v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange) SetPassengerNames(v []*string) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange {
	s.PassengerNames = v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListFlightChange) Validate() error {
	return dara.Validate(s)
}

type FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListSegmentPosition struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListSegmentPosition) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListSegmentPosition) SetJourneyIndex(v int32) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListSegmentPosition) SetSegmentIndex(v int32) *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModuleFlightTaleInfoDTOJourneysSegmentListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type FlightOrderDetailV2ResponseBodyModulePassengerList struct {
	// example:
	//
	// 1991-06-21 00:00:00
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// example:
	//
	// 12172819047252004460056
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	Code        *int32  `json:"code,omitempty" xml:"code,omitempty"`
	Country     *string `json:"country,omitempty" xml:"country,omitempty"`
	// example:
	//
	// CN
	CountryCode *string                                                          `json:"country_code,omitempty" xml:"country_code,omitempty"`
	Credential  *FlightOrderDetailV2ResponseBodyModulePassengerListCredential    `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	Credentials []*FlightOrderDetailV2ResponseBodyModulePassengerListCredentials `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
	// example:
	//
	// 17800000001@163.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// san
	EnFirstName *string `json:"en_first_name,omitempty" xml:"en_first_name,omitempty"`
	// example:
	//
	// zhang
	EnLastName *string `json:"en_last_name,omitempty" xml:"en_last_name,omitempty"`
	// example:
	//
	// zhangsan
	EnglishName *string `json:"english_name,omitempty" xml:"english_name,omitempty"`
	// example:
	//
	// 1
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// 3243028
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	IsComplete *bool `json:"is_complete,omitempty" xml:"is_complete,omitempty"`
	// example:
	//
	// false
	IsFrequently *bool   `json:"is_frequently,omitempty" xml:"is_frequently,omitempty"`
	Memo         *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// example:
	//
	// 17800000001
	MobilePhoneNumber *string `json:"mobile_phone_number,omitempty" xml:"mobile_phone_number,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// true
	Oneself   *bool   `json:"oneself,omitempty" xml:"oneself,omitempty"`
	OrderName *string `json:"order_name,omitempty" xml:"order_name,omitempty"`
	// example:
	//
	// 1111
	OutPassengerId *string `json:"out_passenger_id,omitempty" xml:"out_passenger_id,omitempty"`
	// example:
	//
	// 17800000001
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// example:
	//
	// 111
	ShengPiPinyin *string                                                      `json:"sheng_pi_pinyin,omitempty" xml:"sheng_pi_pinyin,omitempty"`
	TicketNos     []*string                                                    `json:"ticket_nos,omitempty" xml:"ticket_nos,omitempty" type:"Repeated"`
	Tickets       []*FlightOrderDetailV2ResponseBodyModulePassengerListTickets `json:"tickets,omitempty" xml:"tickets,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 312312
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightOrderDetailV2ResponseBodyModulePassengerList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailV2ResponseBodyModulePassengerList) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetBirthday() *string {
	return s.Birthday
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetCode() *int32 {
	return s.Code
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetCountry() *string {
	return s.Country
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetCountryCode() *string {
	return s.CountryCode
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetCredential() *FlightOrderDetailV2ResponseBodyModulePassengerListCredential {
	return s.Credential
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetCredentials() []*FlightOrderDetailV2ResponseBodyModulePassengerListCredentials {
	return s.Credentials
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetEmail() *string {
	return s.Email
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetEnFirstName() *string {
	return s.EnFirstName
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetEnLastName() *string {
	return s.EnLastName
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetEnglishName() *string {
	return s.EnglishName
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetGender() *int32 {
	return s.Gender
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetId() *string {
	return s.Id
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetIsComplete() *bool {
	return s.IsComplete
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetIsFrequently() *bool {
	return s.IsFrequently
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetMemo() *string {
	return s.Memo
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetMobileCountryCode() *string {
	return s.MobileCountryCode
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetMobilePhoneNumber() *string {
	return s.MobilePhoneNumber
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetName() *string {
	return s.Name
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetOneself() *bool {
	return s.Oneself
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetOrderName() *string {
	return s.OrderName
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetOutPassengerId() *string {
	return s.OutPassengerId
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetPhone() *string {
	return s.Phone
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetShengPiPinyin() *string {
	return s.ShengPiPinyin
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetTicketNos() []*string {
	return s.TicketNos
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetTickets() []*FlightOrderDetailV2ResponseBodyModulePassengerListTickets {
	return s.Tickets
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetType() *int32 {
	return s.Type
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetBirthday(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.Birthday = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetBtripUserId(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.BtripUserId = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetCode(v int32) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.Code = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetCountry(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.Country = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetCountryCode(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.CountryCode = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetCredential(v *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.Credential = v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetCredentials(v []*FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.Credentials = v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetEmail(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.Email = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetEnFirstName(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.EnFirstName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetEnLastName(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.EnLastName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetEnglishName(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.EnglishName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetGender(v int32) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.Gender = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetId(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.Id = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetIsComplete(v bool) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.IsComplete = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetIsFrequently(v bool) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.IsFrequently = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetMemo(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.Memo = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetMobileCountryCode(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.MobileCountryCode = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetMobilePhoneNumber(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.MobilePhoneNumber = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetName(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.Name = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetOneself(v bool) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.Oneself = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetOrderName(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.OrderName = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetOutPassengerId(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.OutPassengerId = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetPhone(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.Phone = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetShengPiPinyin(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.ShengPiPinyin = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetTicketNos(v []*string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.TicketNos = v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetTickets(v []*FlightOrderDetailV2ResponseBodyModulePassengerListTickets) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.Tickets = v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetType(v int32) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.Type = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) SetUserId(v string) *FlightOrderDetailV2ResponseBodyModulePassengerList {
	s.UserId = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderDetailV2ResponseBodyModulePassengerListCredential struct {
	// example:
	//
	// 1991-06-21 00:00:00
	BirthDate *string `json:"birth_date,omitempty" xml:"birth_date,omitempty"`
	// example:
	//
	// 1991-06-21 00:00:00
	CertIssueDate     *string `json:"cert_issue_date,omitempty" xml:"cert_issue_date,omitempty"`
	CertIssuePlace    *string `json:"cert_issue_place,omitempty" xml:"cert_issue_place,omitempty"`
	CredentialNo      *string `json:"credential_no,omitempty" xml:"credential_no,omitempty"`
	DriveLicenceFirst *string `json:"drive_licence_first,omitempty" xml:"drive_licence_first,omitempty"`
	DriveLicenceType  *string `json:"drive_licence_type,omitempty" xml:"drive_licence_type,omitempty"`
	// example:
	//
	// 1991-06-21 00:00:00
	ExpireDate        *string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	HolderNationality *string `json:"holder_nationality,omitempty" xml:"holder_nationality,omitempty"`
	// example:
	//
	// 131332
	Id           *int64  `json:"id,omitempty" xml:"id,omitempty"`
	IdCheckCode  *string `json:"id_check_code,omitempty" xml:"id_check_code,omitempty"`
	IssueCountry *string `json:"issue_country,omitempty" xml:"issue_country,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOrderDetailV2ResponseBodyModulePassengerListCredential) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailV2ResponseBodyModulePassengerListCredential) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) GetBirthDate() *string {
	return s.BirthDate
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) GetCertIssueDate() *string {
	return s.CertIssueDate
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) GetCertIssuePlace() *string {
	return s.CertIssuePlace
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) GetCredentialNo() *string {
	return s.CredentialNo
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) GetDriveLicenceFirst() *string {
	return s.DriveLicenceFirst
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) GetDriveLicenceType() *string {
	return s.DriveLicenceType
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) GetHolderNationality() *string {
	return s.HolderNationality
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) GetId() *int64 {
	return s.Id
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) GetIdCheckCode() *string {
	return s.IdCheckCode
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) GetIssueCountry() *string {
	return s.IssueCountry
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) GetType() *int32 {
	return s.Type
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) SetBirthDate(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredential {
	s.BirthDate = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) SetCertIssueDate(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredential {
	s.CertIssueDate = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) SetCertIssuePlace(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredential {
	s.CertIssuePlace = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) SetCredentialNo(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredential {
	s.CredentialNo = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) SetDriveLicenceFirst(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredential {
	s.DriveLicenceFirst = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) SetDriveLicenceType(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredential {
	s.DriveLicenceType = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) SetExpireDate(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredential {
	s.ExpireDate = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) SetHolderNationality(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredential {
	s.HolderNationality = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) SetId(v int64) *FlightOrderDetailV2ResponseBodyModulePassengerListCredential {
	s.Id = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) SetIdCheckCode(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredential {
	s.IdCheckCode = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) SetIssueCountry(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredential {
	s.IssueCountry = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) SetType(v int32) *FlightOrderDetailV2ResponseBodyModulePassengerListCredential {
	s.Type = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredential) Validate() error {
	return dara.Validate(s)
}

type FlightOrderDetailV2ResponseBodyModulePassengerListCredentials struct {
	BirthDate         *string `json:"birth_date,omitempty" xml:"birth_date,omitempty"`
	CertIssueDate     *string `json:"cert_issue_date,omitempty" xml:"cert_issue_date,omitempty"`
	CertIssuePlace    *string `json:"cert_issue_place,omitempty" xml:"cert_issue_place,omitempty"`
	CredentialNo      *string `json:"credential_no,omitempty" xml:"credential_no,omitempty"`
	DriveLicenceFirst *string `json:"drive_licence_first,omitempty" xml:"drive_licence_first,omitempty"`
	DriveLicenceType  *string `json:"drive_licence_type,omitempty" xml:"drive_licence_type,omitempty"`
	ExpireDate        *string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	HolderNationality *string `json:"holder_nationality,omitempty" xml:"holder_nationality,omitempty"`
	// example:
	//
	// 131332
	Id           *int64  `json:"id,omitempty" xml:"id,omitempty"`
	IdCheckCode  *string `json:"id_check_code,omitempty" xml:"id_check_code,omitempty"`
	IssueCountry *string `json:"issue_country,omitempty" xml:"issue_country,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) GetBirthDate() *string {
	return s.BirthDate
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) GetCertIssueDate() *string {
	return s.CertIssueDate
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) GetCertIssuePlace() *string {
	return s.CertIssuePlace
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) GetCredentialNo() *string {
	return s.CredentialNo
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) GetDriveLicenceFirst() *string {
	return s.DriveLicenceFirst
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) GetDriveLicenceType() *string {
	return s.DriveLicenceType
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) GetHolderNationality() *string {
	return s.HolderNationality
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) GetId() *int64 {
	return s.Id
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) GetIdCheckCode() *string {
	return s.IdCheckCode
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) GetIssueCountry() *string {
	return s.IssueCountry
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) GetType() *int32 {
	return s.Type
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) SetBirthDate(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials {
	s.BirthDate = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) SetCertIssueDate(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials {
	s.CertIssueDate = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) SetCertIssuePlace(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials {
	s.CertIssuePlace = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) SetCredentialNo(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials {
	s.CredentialNo = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) SetDriveLicenceFirst(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials {
	s.DriveLicenceFirst = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) SetDriveLicenceType(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials {
	s.DriveLicenceType = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) SetExpireDate(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials {
	s.ExpireDate = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) SetHolderNationality(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials {
	s.HolderNationality = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) SetId(v int64) *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials {
	s.Id = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) SetIdCheckCode(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials {
	s.IdCheckCode = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) SetIssueCountry(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials {
	s.IssueCountry = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) SetType(v int32) *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials {
	s.Type = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListCredentials) Validate() error {
	return dara.Validate(s)
}

type FlightOrderDetailV2ResponseBodyModulePassengerListTickets struct {
	Channel      *string `json:"channel,omitempty" xml:"channel,omitempty"`
	JourneyTitle *string `json:"journey_title,omitempty" xml:"journey_title,omitempty"`
	// example:
	//
	// OPEN_FOR_USE
	OpenTicketStatus *string `json:"open_ticket_status,omitempty" xml:"open_ticket_status,omitempty"`
	// pcc/office
	//
	// example:
	//
	// pcc/office
	Pcc                   *string                                                                           `json:"pcc,omitempty" xml:"pcc,omitempty"`
	SegmentOpenTicketList []*FlightOrderDetailV2ResponseBodyModulePassengerListTicketsSegmentOpenTicketList `json:"segment_open_ticket_list,omitempty" xml:"segment_open_ticket_list,omitempty" type:"Repeated"`
	TicketAuthMemo        *string                                                                           `json:"ticket_auth_memo,omitempty" xml:"ticket_auth_memo,omitempty"`
	// example:
	//
	// 2
	TicketAuthStatus *int32 `json:"ticket_auth_status,omitempty" xml:"ticket_auth_status,omitempty"`
	// example:
	//
	// 444-2023062999
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// example:
	//
	// 4500
	TicketPrice  *int64  `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	TicketStatus *string `json:"ticket_status,omitempty" xml:"ticket_status,omitempty"`
}

func (s FlightOrderDetailV2ResponseBodyModulePassengerListTickets) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailV2ResponseBodyModulePassengerListTickets) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) GetChannel() *string {
	return s.Channel
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) GetJourneyTitle() *string {
	return s.JourneyTitle
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) GetOpenTicketStatus() *string {
	return s.OpenTicketStatus
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) GetPcc() *string {
	return s.Pcc
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) GetSegmentOpenTicketList() []*FlightOrderDetailV2ResponseBodyModulePassengerListTicketsSegmentOpenTicketList {
	return s.SegmentOpenTicketList
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) GetTicketAuthMemo() *string {
	return s.TicketAuthMemo
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) GetTicketAuthStatus() *int32 {
	return s.TicketAuthStatus
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) GetTicketNo() *string {
	return s.TicketNo
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) GetTicketPrice() *int64 {
	return s.TicketPrice
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) GetTicketStatus() *string {
	return s.TicketStatus
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) SetChannel(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListTickets {
	s.Channel = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) SetJourneyTitle(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListTickets {
	s.JourneyTitle = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) SetOpenTicketStatus(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListTickets {
	s.OpenTicketStatus = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) SetPcc(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListTickets {
	s.Pcc = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) SetSegmentOpenTicketList(v []*FlightOrderDetailV2ResponseBodyModulePassengerListTicketsSegmentOpenTicketList) *FlightOrderDetailV2ResponseBodyModulePassengerListTickets {
	s.SegmentOpenTicketList = v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) SetTicketAuthMemo(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListTickets {
	s.TicketAuthMemo = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) SetTicketAuthStatus(v int32) *FlightOrderDetailV2ResponseBodyModulePassengerListTickets {
	s.TicketAuthStatus = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) SetTicketNo(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListTickets {
	s.TicketNo = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) SetTicketPrice(v int64) *FlightOrderDetailV2ResponseBodyModulePassengerListTickets {
	s.TicketPrice = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) SetTicketStatus(v string) *FlightOrderDetailV2ResponseBodyModulePassengerListTickets {
	s.TicketStatus = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTickets) Validate() error {
	return dara.Validate(s)
}

type FlightOrderDetailV2ResponseBodyModulePassengerListTicketsSegmentOpenTicketList struct {
	JourneyIndex     *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	OpenTicketStatus *int32 `json:"open_ticket_status,omitempty" xml:"open_ticket_status,omitempty"`
	SegmentIndex     *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s FlightOrderDetailV2ResponseBodyModulePassengerListTicketsSegmentOpenTicketList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailV2ResponseBodyModulePassengerListTicketsSegmentOpenTicketList) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTicketsSegmentOpenTicketList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTicketsSegmentOpenTicketList) GetOpenTicketStatus() *int32 {
	return s.OpenTicketStatus
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTicketsSegmentOpenTicketList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTicketsSegmentOpenTicketList) SetJourneyIndex(v int32) *FlightOrderDetailV2ResponseBodyModulePassengerListTicketsSegmentOpenTicketList {
	s.JourneyIndex = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTicketsSegmentOpenTicketList) SetOpenTicketStatus(v int32) *FlightOrderDetailV2ResponseBodyModulePassengerListTicketsSegmentOpenTicketList {
	s.OpenTicketStatus = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTicketsSegmentOpenTicketList) SetSegmentIndex(v int32) *FlightOrderDetailV2ResponseBodyModulePassengerListTicketsSegmentOpenTicketList {
	s.SegmentIndex = &v
	return s
}

func (s *FlightOrderDetailV2ResponseBodyModulePassengerListTicketsSegmentOpenTicketList) Validate() error {
	return dara.Validate(s)
}
