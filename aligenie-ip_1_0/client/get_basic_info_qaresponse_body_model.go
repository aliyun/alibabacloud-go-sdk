// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicInfoQAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *GetBasicInfoQAResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBasicInfoQAResponseBody
	GetRequestId() *string
	SetResult(v *GetBasicInfoQAResponseBodyResult) *GetBasicInfoQAResponseBody
	GetResult() *GetBasicInfoQAResponseBodyResult
	SetStatusCode(v int32) *GetBasicInfoQAResponseBody
	GetStatusCode() *int32
}

type GetBasicInfoQAResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7***726E
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetBasicInfoQAResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetBasicInfoQAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBasicInfoQAResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicInfoQAResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBasicInfoQAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBasicInfoQAResponseBody) GetResult() *GetBasicInfoQAResponseBodyResult {
	return s.Result
}

func (s *GetBasicInfoQAResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBasicInfoQAResponseBody) SetMessage(v string) *GetBasicInfoQAResponseBody {
	s.Message = &v
	return s
}

func (s *GetBasicInfoQAResponseBody) SetRequestId(v string) *GetBasicInfoQAResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicInfoQAResponseBody) SetResult(v *GetBasicInfoQAResponseBodyResult) *GetBasicInfoQAResponseBody {
	s.Result = v
	return s
}

func (s *GetBasicInfoQAResponseBody) SetStatusCode(v int32) *GetBasicInfoQAResponseBody {
	s.StatusCode = &v
	return s
}

func (s *GetBasicInfoQAResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBasicInfoQAResponseBodyResult struct {
	// example:
	//
	// 11:11
	CheckInTime *string `json:"CheckInTime,omitempty" xml:"CheckInTime,omitempty"`
	// example:
	//
	// 11:11
	CheckOutTime      *string `json:"CheckOutTime,omitempty" xml:"CheckOutTime,omitempty"`
	HotelAddress      *string `json:"HotelAddress,omitempty" xml:"HotelAddress,omitempty"`
	HotelIntroduction *string `json:"HotelIntroduction,omitempty" xml:"HotelIntroduction,omitempty"`
	HotelMember       *string `json:"HotelMember,omitempty" xml:"HotelMember,omitempty"`
	HotelService      *string `json:"HotelService,omitempty" xml:"HotelService,omitempty"`
	ParkingExpenses   *string `json:"ParkingExpenses,omitempty" xml:"ParkingExpenses,omitempty"`
	ParkingPosition   *string `json:"ParkingPosition,omitempty" xml:"ParkingPosition,omitempty"`
	// example:
	//
	// 123***
	PhoneNumber  *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	WifiName     *string `json:"WifiName,omitempty" xml:"WifiName,omitempty"`
	WifiPassword *string `json:"WifiPassword,omitempty" xml:"WifiPassword,omitempty"`
}

func (s GetBasicInfoQAResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetBasicInfoQAResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetBasicInfoQAResponseBodyResult) GetCheckInTime() *string {
	return s.CheckInTime
}

func (s *GetBasicInfoQAResponseBodyResult) GetCheckOutTime() *string {
	return s.CheckOutTime
}

func (s *GetBasicInfoQAResponseBodyResult) GetHotelAddress() *string {
	return s.HotelAddress
}

func (s *GetBasicInfoQAResponseBodyResult) GetHotelIntroduction() *string {
	return s.HotelIntroduction
}

func (s *GetBasicInfoQAResponseBodyResult) GetHotelMember() *string {
	return s.HotelMember
}

func (s *GetBasicInfoQAResponseBodyResult) GetHotelService() *string {
	return s.HotelService
}

func (s *GetBasicInfoQAResponseBodyResult) GetParkingExpenses() *string {
	return s.ParkingExpenses
}

func (s *GetBasicInfoQAResponseBodyResult) GetParkingPosition() *string {
	return s.ParkingPosition
}

func (s *GetBasicInfoQAResponseBodyResult) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetBasicInfoQAResponseBodyResult) GetWifiName() *string {
	return s.WifiName
}

func (s *GetBasicInfoQAResponseBodyResult) GetWifiPassword() *string {
	return s.WifiPassword
}

func (s *GetBasicInfoQAResponseBodyResult) SetCheckInTime(v string) *GetBasicInfoQAResponseBodyResult {
	s.CheckInTime = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetCheckOutTime(v string) *GetBasicInfoQAResponseBodyResult {
	s.CheckOutTime = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetHotelAddress(v string) *GetBasicInfoQAResponseBodyResult {
	s.HotelAddress = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetHotelIntroduction(v string) *GetBasicInfoQAResponseBodyResult {
	s.HotelIntroduction = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetHotelMember(v string) *GetBasicInfoQAResponseBodyResult {
	s.HotelMember = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetHotelService(v string) *GetBasicInfoQAResponseBodyResult {
	s.HotelService = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetParkingExpenses(v string) *GetBasicInfoQAResponseBodyResult {
	s.ParkingExpenses = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetParkingPosition(v string) *GetBasicInfoQAResponseBodyResult {
	s.ParkingPosition = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetPhoneNumber(v string) *GetBasicInfoQAResponseBodyResult {
	s.PhoneNumber = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetWifiName(v string) *GetBasicInfoQAResponseBodyResult {
	s.WifiName = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) SetWifiPassword(v string) *GetBasicInfoQAResponseBodyResult {
	s.WifiPassword = &v
	return s
}

func (s *GetBasicInfoQAResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
