// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeApplyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangePassengerList(v []*ChangeApplyRequestChangePassengerList) *ChangeApplyRequest
	GetChangePassengerList() []*ChangeApplyRequestChangePassengerList
	SetChangedJourneys(v []*ChangeApplyRequestChangedJourneys) *ChangeApplyRequest
	GetChangedJourneys() []*ChangeApplyRequestChangedJourneys
	SetContact(v *ChangeApplyRequestContact) *ChangeApplyRequest
	GetContact() *ChangeApplyRequestContact
	SetOrderNum(v int64) *ChangeApplyRequest
	GetOrderNum() *int64
	SetRemark(v string) *ChangeApplyRequest
	GetRemark() *string
	SetType(v int32) *ChangeApplyRequest
	GetType() *int32
}

type ChangeApplyRequest struct {
	// This parameter is required.
	ChangePassengerList []*ChangeApplyRequestChangePassengerList `json:"change_passenger_list,omitempty" xml:"change_passenger_list,omitempty" type:"Repeated"`
	// This parameter is required.
	ChangedJourneys []*ChangeApplyRequestChangedJourneys `json:"changed_journeys,omitempty" xml:"changed_journeys,omitempty" type:"Repeated"`
	// This parameter is required.
	Contact *ChangeApplyRequestContact `json:"contact,omitempty" xml:"contact,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 4988430***950
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// example:
	//
	// remark desc
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ChangeApplyRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeApplyRequest) GoString() string {
	return s.String()
}

func (s *ChangeApplyRequest) GetChangePassengerList() []*ChangeApplyRequestChangePassengerList {
	return s.ChangePassengerList
}

func (s *ChangeApplyRequest) GetChangedJourneys() []*ChangeApplyRequestChangedJourneys {
	return s.ChangedJourneys
}

func (s *ChangeApplyRequest) GetContact() *ChangeApplyRequestContact {
	return s.Contact
}

func (s *ChangeApplyRequest) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *ChangeApplyRequest) GetRemark() *string {
	return s.Remark
}

func (s *ChangeApplyRequest) GetType() *int32 {
	return s.Type
}

func (s *ChangeApplyRequest) SetChangePassengerList(v []*ChangeApplyRequestChangePassengerList) *ChangeApplyRequest {
	s.ChangePassengerList = v
	return s
}

func (s *ChangeApplyRequest) SetChangedJourneys(v []*ChangeApplyRequestChangedJourneys) *ChangeApplyRequest {
	s.ChangedJourneys = v
	return s
}

func (s *ChangeApplyRequest) SetContact(v *ChangeApplyRequestContact) *ChangeApplyRequest {
	s.Contact = v
	return s
}

func (s *ChangeApplyRequest) SetOrderNum(v int64) *ChangeApplyRequest {
	s.OrderNum = &v
	return s
}

func (s *ChangeApplyRequest) SetRemark(v string) *ChangeApplyRequest {
	s.Remark = &v
	return s
}

func (s *ChangeApplyRequest) SetType(v int32) *ChangeApplyRequest {
	s.Type = &v
	return s
}

func (s *ChangeApplyRequest) Validate() error {
	if s.ChangePassengerList != nil {
		for _, item := range s.ChangePassengerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ChangedJourneys != nil {
		for _, item := range s.ChangedJourneys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Contact != nil {
		if err := s.Contact.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeApplyRequestChangePassengerList struct {
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s ChangeApplyRequestChangePassengerList) String() string {
	return dara.Prettify(s)
}

func (s ChangeApplyRequestChangePassengerList) GoString() string {
	return s.String()
}

func (s *ChangeApplyRequestChangePassengerList) GetDocument() *string {
	return s.Document
}

func (s *ChangeApplyRequestChangePassengerList) GetFirstName() *string {
	return s.FirstName
}

func (s *ChangeApplyRequestChangePassengerList) GetLastName() *string {
	return s.LastName
}

func (s *ChangeApplyRequestChangePassengerList) SetDocument(v string) *ChangeApplyRequestChangePassengerList {
	s.Document = &v
	return s
}

func (s *ChangeApplyRequestChangePassengerList) SetFirstName(v string) *ChangeApplyRequestChangePassengerList {
	s.FirstName = &v
	return s
}

func (s *ChangeApplyRequestChangePassengerList) SetLastName(v string) *ChangeApplyRequestChangePassengerList {
	s.LastName = &v
	return s
}

func (s *ChangeApplyRequestChangePassengerList) Validate() error {
	return dara.Validate(s)
}

type ChangeApplyRequestChangedJourneys struct {
	SegmentList []*ChangeApplyRequestChangedJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
}

func (s ChangeApplyRequestChangedJourneys) String() string {
	return dara.Prettify(s)
}

func (s ChangeApplyRequestChangedJourneys) GoString() string {
	return s.String()
}

func (s *ChangeApplyRequestChangedJourneys) GetSegmentList() []*ChangeApplyRequestChangedJourneysSegmentList {
	return s.SegmentList
}

func (s *ChangeApplyRequestChangedJourneys) SetSegmentList(v []*ChangeApplyRequestChangedJourneysSegmentList) *ChangeApplyRequestChangedJourneys {
	s.SegmentList = v
	return s
}

func (s *ChangeApplyRequestChangedJourneys) Validate() error {
	if s.SegmentList != nil {
		for _, item := range s.SegmentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChangeApplyRequestChangedJourneysSegmentList struct {
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// example:
	//
	// T1
	ArriveTerminal *string `json:"arrive_terminal,omitempty" xml:"arrive_terminal,omitempty"`
	// example:
	//
	// 1677232999000
	ArriveTime    *int64  `json:"arrive_time,omitempty" xml:"arrive_time,omitempty"`
	ArriveTimeStr *string `json:"arrive_time_str,omitempty" xml:"arrive_time_str,omitempty"`
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20230320
	DepartureDate *string `json:"departure_date,omitempty" xml:"departure_date,omitempty"`
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// example:
	//
	// 1677232998000
	DepartureTime    *int64  `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	DepartureTimeStr *string `json:"departure_time_str,omitempty" xml:"departure_time_str,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
}

func (s ChangeApplyRequestChangedJourneysSegmentList) String() string {
	return dara.Prettify(s)
}

func (s ChangeApplyRequestChangedJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) GetArrivalAirport() *string {
	return s.ArrivalAirport
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) GetArrivalCity() *string {
	return s.ArrivalCity
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) GetArriveTerminal() *string {
	return s.ArriveTerminal
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) GetArriveTime() *int64 {
	return s.ArriveTime
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) GetArriveTimeStr() *string {
	return s.ArriveTimeStr
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) GetCodeShare() *bool {
	return s.CodeShare
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) GetDepartureAirport() *string {
	return s.DepartureAirport
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) GetDepartureCity() *string {
	return s.DepartureCity
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) GetDepartureDate() *string {
	return s.DepartureDate
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) GetDepartureTerminal() *string {
	return s.DepartureTerminal
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) GetDepartureTime() *int64 {
	return s.DepartureTime
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) GetDepartureTimeStr() *string {
	return s.DepartureTimeStr
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) GetMarketingFlightNo() *string {
	return s.MarketingFlightNo
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetArrivalAirport(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetArrivalCity(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetArriveTerminal(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.ArriveTerminal = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetArriveTime(v int64) *ChangeApplyRequestChangedJourneysSegmentList {
	s.ArriveTime = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetArriveTimeStr(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.ArriveTimeStr = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetCodeShare(v bool) *ChangeApplyRequestChangedJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetDepartureAirport(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetDepartureCity(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetDepartureDate(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.DepartureDate = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetDepartureTerminal(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetDepartureTime(v int64) *ChangeApplyRequestChangedJourneysSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetDepartureTimeStr(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.DepartureTimeStr = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetMarketingFlightNo(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetOperatingFlightNo(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) Validate() error {
	return dara.Validate(s)
}

type ChangeApplyRequestContact struct {
	// example:
	//
	// gao******@gmail.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// example:
	//
	// 183*****92
	MobilePhoneNum *string `json:"mobile_phone_num,omitempty" xml:"mobile_phone_num,omitempty"`
}

func (s ChangeApplyRequestContact) String() string {
	return dara.Prettify(s)
}

func (s ChangeApplyRequestContact) GoString() string {
	return s.String()
}

func (s *ChangeApplyRequestContact) GetEmail() *string {
	return s.Email
}

func (s *ChangeApplyRequestContact) GetMobileCountryCode() *string {
	return s.MobileCountryCode
}

func (s *ChangeApplyRequestContact) GetMobilePhoneNum() *string {
	return s.MobilePhoneNum
}

func (s *ChangeApplyRequestContact) SetEmail(v string) *ChangeApplyRequestContact {
	s.Email = &v
	return s
}

func (s *ChangeApplyRequestContact) SetMobileCountryCode(v string) *ChangeApplyRequestContact {
	s.MobileCountryCode = &v
	return s
}

func (s *ChangeApplyRequestContact) SetMobilePhoneNum(v string) *ChangeApplyRequestContact {
	s.MobilePhoneNum = &v
	return s
}

func (s *ChangeApplyRequestContact) Validate() error {
	return dara.Validate(s)
}
