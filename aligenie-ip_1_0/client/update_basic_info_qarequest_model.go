// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicInfoQARequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckInTime(v string) *UpdateBasicInfoQARequest
	GetCheckInTime() *string
	SetCheckOutTime(v string) *UpdateBasicInfoQARequest
	GetCheckOutTime() *string
	SetHotelAddress(v string) *UpdateBasicInfoQARequest
	GetHotelAddress() *string
	SetHotelId(v string) *UpdateBasicInfoQARequest
	GetHotelId() *string
	SetHotelIntroduction(v string) *UpdateBasicInfoQARequest
	GetHotelIntroduction() *string
	SetHotelMember(v string) *UpdateBasicInfoQARequest
	GetHotelMember() *string
	SetHotelService(v string) *UpdateBasicInfoQARequest
	GetHotelService() *string
	SetParkingExpenses(v string) *UpdateBasicInfoQARequest
	GetParkingExpenses() *string
	SetParkingPosition(v string) *UpdateBasicInfoQARequest
	GetParkingPosition() *string
	SetPhoneNumber(v string) *UpdateBasicInfoQARequest
	GetPhoneNumber() *string
	SetWifiName(v string) *UpdateBasicInfoQARequest
	GetWifiName() *string
	SetWifiPassword(v string) *UpdateBasicInfoQARequest
	GetWifiPassword() *string
}

type UpdateBasicInfoQARequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11:11
	CheckInTime *string `json:"CheckInTime,omitempty" xml:"CheckInTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11:11
	CheckOutTime *string `json:"CheckOutTime,omitempty" xml:"CheckOutTime,omitempty"`
	// This parameter is required.
	HotelAddress *string `json:"HotelAddress,omitempty" xml:"HotelAddress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId           *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	HotelIntroduction *string `json:"HotelIntroduction,omitempty" xml:"HotelIntroduction,omitempty"`
	HotelMember       *string `json:"HotelMember,omitempty" xml:"HotelMember,omitempty"`
	HotelService      *string `json:"HotelService,omitempty" xml:"HotelService,omitempty"`
	// This parameter is required.
	ParkingExpenses *string `json:"ParkingExpenses,omitempty" xml:"ParkingExpenses,omitempty"`
	// This parameter is required.
	ParkingPosition *string `json:"ParkingPosition,omitempty" xml:"ParkingPosition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123***
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// This parameter is required.
	WifiName *string `json:"WifiName,omitempty" xml:"WifiName,omitempty"`
	// This parameter is required.
	WifiPassword *string `json:"WifiPassword,omitempty" xml:"WifiPassword,omitempty"`
}

func (s UpdateBasicInfoQARequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicInfoQARequest) GoString() string {
	return s.String()
}

func (s *UpdateBasicInfoQARequest) GetCheckInTime() *string {
	return s.CheckInTime
}

func (s *UpdateBasicInfoQARequest) GetCheckOutTime() *string {
	return s.CheckOutTime
}

func (s *UpdateBasicInfoQARequest) GetHotelAddress() *string {
	return s.HotelAddress
}

func (s *UpdateBasicInfoQARequest) GetHotelId() *string {
	return s.HotelId
}

func (s *UpdateBasicInfoQARequest) GetHotelIntroduction() *string {
	return s.HotelIntroduction
}

func (s *UpdateBasicInfoQARequest) GetHotelMember() *string {
	return s.HotelMember
}

func (s *UpdateBasicInfoQARequest) GetHotelService() *string {
	return s.HotelService
}

func (s *UpdateBasicInfoQARequest) GetParkingExpenses() *string {
	return s.ParkingExpenses
}

func (s *UpdateBasicInfoQARequest) GetParkingPosition() *string {
	return s.ParkingPosition
}

func (s *UpdateBasicInfoQARequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *UpdateBasicInfoQARequest) GetWifiName() *string {
	return s.WifiName
}

func (s *UpdateBasicInfoQARequest) GetWifiPassword() *string {
	return s.WifiPassword
}

func (s *UpdateBasicInfoQARequest) SetCheckInTime(v string) *UpdateBasicInfoQARequest {
	s.CheckInTime = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetCheckOutTime(v string) *UpdateBasicInfoQARequest {
	s.CheckOutTime = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetHotelAddress(v string) *UpdateBasicInfoQARequest {
	s.HotelAddress = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetHotelId(v string) *UpdateBasicInfoQARequest {
	s.HotelId = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetHotelIntroduction(v string) *UpdateBasicInfoQARequest {
	s.HotelIntroduction = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetHotelMember(v string) *UpdateBasicInfoQARequest {
	s.HotelMember = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetHotelService(v string) *UpdateBasicInfoQARequest {
	s.HotelService = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetParkingExpenses(v string) *UpdateBasicInfoQARequest {
	s.ParkingExpenses = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetParkingPosition(v string) *UpdateBasicInfoQARequest {
	s.ParkingPosition = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetPhoneNumber(v string) *UpdateBasicInfoQARequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetWifiName(v string) *UpdateBasicInfoQARequest {
	s.WifiName = &v
	return s
}

func (s *UpdateBasicInfoQARequest) SetWifiPassword(v string) *UpdateBasicInfoQARequest {
	s.WifiPassword = &v
	return s
}

func (s *UpdateBasicInfoQARequest) Validate() error {
	return dara.Validate(s)
}
