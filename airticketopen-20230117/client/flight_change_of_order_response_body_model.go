// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightChangeOfOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *FlightChangeOfOrderResponseBody
	GetRequestId() *string
	SetData(v []*FlightChangeOfOrderResponseBodyData) *FlightChangeOfOrderResponseBody
	GetData() []*FlightChangeOfOrderResponseBodyData
	SetErrorCode(v string) *FlightChangeOfOrderResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *FlightChangeOfOrderResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *FlightChangeOfOrderResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *FlightChangeOfOrderResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *FlightChangeOfOrderResponseBody
	GetSuccess() *bool
}

type FlightChangeOfOrderResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data returned for a successful request.
	Data []*FlightChangeOfOrderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The business error code.
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// The data returned with the error.
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// The error message.
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// The HTTP status code. The value is always 200 for successful requests.
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FlightChangeOfOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightChangeOfOrderResponseBody) GoString() string {
	return s.String()
}

func (s *FlightChangeOfOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightChangeOfOrderResponseBody) GetData() []*FlightChangeOfOrderResponseBodyData {
	return s.Data
}

func (s *FlightChangeOfOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *FlightChangeOfOrderResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *FlightChangeOfOrderResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *FlightChangeOfOrderResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *FlightChangeOfOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightChangeOfOrderResponseBody) SetRequestId(v string) *FlightChangeOfOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightChangeOfOrderResponseBody) SetData(v []*FlightChangeOfOrderResponseBodyData) *FlightChangeOfOrderResponseBody {
	s.Data = v
	return s
}

func (s *FlightChangeOfOrderResponseBody) SetErrorCode(v string) *FlightChangeOfOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *FlightChangeOfOrderResponseBody) SetErrorData(v interface{}) *FlightChangeOfOrderResponseBody {
	s.ErrorData = v
	return s
}

func (s *FlightChangeOfOrderResponseBody) SetErrorMsg(v string) *FlightChangeOfOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *FlightChangeOfOrderResponseBody) SetStatus(v int32) *FlightChangeOfOrderResponseBody {
	s.Status = &v
	return s
}

func (s *FlightChangeOfOrderResponseBody) SetSuccess(v bool) *FlightChangeOfOrderResponseBody {
	s.Success = &v
	return s
}

func (s *FlightChangeOfOrderResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type FlightChangeOfOrderResponseBodyData struct {
	// The flight change information.
	FlightChangeDetail *FlightChangeOfOrderResponseBodyDataFlightChangeDetail `json:"flight_change_detail,omitempty" xml:"flight_change_detail,omitempty" type:"Struct"`
	// The order number.
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s FlightChangeOfOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FlightChangeOfOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *FlightChangeOfOrderResponseBodyData) GetFlightChangeDetail() *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	return s.FlightChangeDetail
}

func (s *FlightChangeOfOrderResponseBodyData) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *FlightChangeOfOrderResponseBodyData) SetFlightChangeDetail(v *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) *FlightChangeOfOrderResponseBodyData {
	s.FlightChangeDetail = v
	return s
}

func (s *FlightChangeOfOrderResponseBodyData) SetOrderNum(v int64) *FlightChangeOfOrderResponseBodyData {
	s.OrderNum = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyData) Validate() error {
	if s.FlightChangeDetail != nil {
		if err := s.FlightChangeDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FlightChangeOfOrderResponseBodyDataFlightChangeDetail struct {
	// The reason for the flight change.
	//
	// example:
	//
	// 天气
	ChangeReason *string `json:"change_reason,omitempty" xml:"change_reason,omitempty"`
	// The time of the flight change in string format (yyyy-MM-dd HH:mm:ss).
	//
	// example:
	//
	// 2023-02-01 10:01:00
	ChangeTime *string `json:"change_time,omitempty" xml:"change_time,omitempty"`
	// The type of the flight change. Valid values:
	//
	// - 1: cancellation
	//
	// - 2: schedule change.
	//
	// example:
	//
	// 1
	ChangeType *int32 `json:"change_type,omitempty" xml:"change_type,omitempty"`
	// The three-letter IATA code of the new arrival airport (uppercase).
	//
	// example:
	//
	// MFM
	NewArrivalAirport *string `json:"new_arrival_airport,omitempty" xml:"new_arrival_airport,omitempty"`
	// The arrival date and time of the new flight in string format (yyyy-MM-dd HH:mm:ss).
	//
	// example:
	//
	// 2023-02-01 15:01:00
	NewArrivalTime *string `json:"new_arrival_time,omitempty" xml:"new_arrival_time,omitempty"`
	// The three-letter IATA code of the new departure airport (uppercase).
	//
	// example:
	//
	// PVG
	NewDepartureAirport *string `json:"new_departure_airport,omitempty" xml:"new_departure_airport,omitempty"`
	// The departure date and time of the new flight in string format (yyyy-MM-dd HH:mm:ss).
	//
	// example:
	//
	// 2023-02-01 13:01:00
	NewDepartureTime *string `json:"new_departure_time,omitempty" xml:"new_departure_time,omitempty"`
	// The new flight number.
	//
	// example:
	//
	// HO1295
	NewFlightNo *string `json:"new_flight_no,omitempty" xml:"new_flight_no,omitempty"`
	// The three-letter IATA code of the original arrival airport (uppercase).
	//
	// example:
	//
	// MFM
	OldArrivalAirport *string `json:"old_arrival_airport,omitempty" xml:"old_arrival_airport,omitempty"`
	// The arrival date and time of the original flight in string format (yyyy-MM-dd HH:mm:ss).
	//
	// example:
	//
	// 023-02-01 14:01:00
	OldArrivalTime *string `json:"old_arrival_time,omitempty" xml:"old_arrival_time,omitempty"`
	// The three-letter IATA code of the original departure airport (uppercase).
	//
	// example:
	//
	// PVG
	OldDepartureAirport *string `json:"old_departure_airport,omitempty" xml:"old_departure_airport,omitempty"`
	// The departure date and time of the original flight in string format (yyyy-MM-dd HH:mm:ss).
	//
	// example:
	//
	// 2023-02-01 12:01:00
	OldDepartureTime *string `json:"old_departure_time,omitempty" xml:"old_departure_time,omitempty"`
	// The original flight number.
	//
	// example:
	//
	// HO1295
	OldFlightNo *string `json:"old_flight_no,omitempty" xml:"old_flight_no,omitempty"`
}

func (s FlightChangeOfOrderResponseBodyDataFlightChangeDetail) String() string {
	return dara.Prettify(s)
}

func (s FlightChangeOfOrderResponseBodyDataFlightChangeDetail) GoString() string {
	return s.String()
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) GetChangeReason() *string {
	return s.ChangeReason
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) GetChangeTime() *string {
	return s.ChangeTime
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) GetChangeType() *int32 {
	return s.ChangeType
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) GetNewArrivalAirport() *string {
	return s.NewArrivalAirport
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) GetNewArrivalTime() *string {
	return s.NewArrivalTime
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) GetNewDepartureAirport() *string {
	return s.NewDepartureAirport
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) GetNewDepartureTime() *string {
	return s.NewDepartureTime
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) GetNewFlightNo() *string {
	return s.NewFlightNo
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) GetOldArrivalAirport() *string {
	return s.OldArrivalAirport
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) GetOldArrivalTime() *string {
	return s.OldArrivalTime
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) GetOldDepartureAirport() *string {
	return s.OldDepartureAirport
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) GetOldDepartureTime() *string {
	return s.OldDepartureTime
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) GetOldFlightNo() *string {
	return s.OldFlightNo
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetChangeReason(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.ChangeReason = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetChangeTime(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.ChangeTime = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetChangeType(v int32) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.ChangeType = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetNewArrivalAirport(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.NewArrivalAirport = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetNewArrivalTime(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.NewArrivalTime = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetNewDepartureAirport(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.NewDepartureAirport = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetNewDepartureTime(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.NewDepartureTime = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetNewFlightNo(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.NewFlightNo = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetOldArrivalAirport(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.OldArrivalAirport = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetOldArrivalTime(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.OldArrivalTime = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetOldDepartureAirport(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.OldDepartureAirport = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetOldDepartureTime(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.OldDepartureTime = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetOldFlightNo(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.OldFlightNo = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) Validate() error {
	return dara.Validate(s)
}
