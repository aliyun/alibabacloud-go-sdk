// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundApplyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderNum(v int64) *RefundApplyRequest
	GetOrderNum() *int64
	SetRefundJourneys(v []*RefundApplyRequestRefundJourneys) *RefundApplyRequest
	GetRefundJourneys() []*RefundApplyRequestRefundJourneys
	SetRefundPassengerList(v []*RefundApplyRequestRefundPassengerList) *RefundApplyRequest
	GetRefundPassengerList() []*RefundApplyRequestRefundPassengerList
	SetRefundType(v *RefundApplyRequestRefundType) *RefundApplyRequest
	GetRefundType() *RefundApplyRequestRefundType
}

type RefundApplyRequest struct {
	// The order number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// The journeys for the refund application.
	//
	// This parameter is required.
	RefundJourneys []*RefundApplyRequestRefundJourneys `json:"refund_journeys,omitempty" xml:"refund_journeys,omitempty" type:"Repeated"`
	// The list of passengers for the refund application.
	//
	// This parameter is required.
	RefundPassengerList []*RefundApplyRequestRefundPassengerList `json:"refund_passenger_list,omitempty" xml:"refund_passenger_list,omitempty" type:"Repeated"`
	// The refund type. Attachments are required for involuntary refund applications.
	//
	// This parameter is required.
	RefundType *RefundApplyRequestRefundType `json:"refund_type,omitempty" xml:"refund_type,omitempty" type:"Struct"`
}

func (s RefundApplyRequest) String() string {
	return dara.Prettify(s)
}

func (s RefundApplyRequest) GoString() string {
	return s.String()
}

func (s *RefundApplyRequest) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *RefundApplyRequest) GetRefundJourneys() []*RefundApplyRequestRefundJourneys {
	return s.RefundJourneys
}

func (s *RefundApplyRequest) GetRefundPassengerList() []*RefundApplyRequestRefundPassengerList {
	return s.RefundPassengerList
}

func (s *RefundApplyRequest) GetRefundType() *RefundApplyRequestRefundType {
	return s.RefundType
}

func (s *RefundApplyRequest) SetOrderNum(v int64) *RefundApplyRequest {
	s.OrderNum = &v
	return s
}

func (s *RefundApplyRequest) SetRefundJourneys(v []*RefundApplyRequestRefundJourneys) *RefundApplyRequest {
	s.RefundJourneys = v
	return s
}

func (s *RefundApplyRequest) SetRefundPassengerList(v []*RefundApplyRequestRefundPassengerList) *RefundApplyRequest {
	s.RefundPassengerList = v
	return s
}

func (s *RefundApplyRequest) SetRefundType(v *RefundApplyRequestRefundType) *RefundApplyRequest {
	s.RefundType = v
	return s
}

func (s *RefundApplyRequest) Validate() error {
	if s.RefundJourneys != nil {
		for _, item := range s.RefundJourneys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RefundPassengerList != nil {
		for _, item := range s.RefundPassengerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RefundType != nil {
		if err := s.RefundType.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RefundApplyRequestRefundJourneys struct {
	// The segment information.
	//
	// This parameter is required.
	SegmentList []*RefundApplyRequestRefundJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
}

func (s RefundApplyRequestRefundJourneys) String() string {
	return dara.Prettify(s)
}

func (s RefundApplyRequestRefundJourneys) GoString() string {
	return s.String()
}

func (s *RefundApplyRequestRefundJourneys) GetSegmentList() []*RefundApplyRequestRefundJourneysSegmentList {
	return s.SegmentList
}

func (s *RefundApplyRequestRefundJourneys) SetSegmentList(v []*RefundApplyRequestRefundJourneysSegmentList) *RefundApplyRequestRefundJourneys {
	s.SegmentList = v
	return s
}

func (s *RefundApplyRequestRefundJourneys) Validate() error {
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

type RefundApplyRequestRefundJourneysSegmentList struct {
	// The three-letter IATA code of the arrival airport (uppercase).
	//
	// This parameter is required.
	//
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// The three-letter IATA code of the arrival city (uppercase).
	//
	// This parameter is required.
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// The three-letter IATA code of the departure airport (uppercase).
	//
	// This parameter is required.
	//
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// The three-letter IATA code of the departure city (uppercase).
	//
	// This parameter is required.
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
}

func (s RefundApplyRequestRefundJourneysSegmentList) String() string {
	return dara.Prettify(s)
}

func (s RefundApplyRequestRefundJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *RefundApplyRequestRefundJourneysSegmentList) GetArrivalAirport() *string {
	return s.ArrivalAirport
}

func (s *RefundApplyRequestRefundJourneysSegmentList) GetArrivalCity() *string {
	return s.ArrivalCity
}

func (s *RefundApplyRequestRefundJourneysSegmentList) GetDepartureAirport() *string {
	return s.DepartureAirport
}

func (s *RefundApplyRequestRefundJourneysSegmentList) GetDepartureCity() *string {
	return s.DepartureCity
}

func (s *RefundApplyRequestRefundJourneysSegmentList) SetArrivalAirport(v string) *RefundApplyRequestRefundJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *RefundApplyRequestRefundJourneysSegmentList) SetArrivalCity(v string) *RefundApplyRequestRefundJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *RefundApplyRequestRefundJourneysSegmentList) SetDepartureAirport(v string) *RefundApplyRequestRefundJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *RefundApplyRequestRefundJourneysSegmentList) SetDepartureCity(v string) *RefundApplyRequestRefundJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *RefundApplyRequestRefundJourneysSegmentList) Validate() error {
	return dara.Validate(s)
}

type RefundApplyRequestRefundPassengerList struct {
	// The document number of the passenger.
	//
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// The first name of the passenger.
	//
	// This parameter is required.
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// The last name of the passenger.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s RefundApplyRequestRefundPassengerList) String() string {
	return dara.Prettify(s)
}

func (s RefundApplyRequestRefundPassengerList) GoString() string {
	return s.String()
}

func (s *RefundApplyRequestRefundPassengerList) GetDocument() *string {
	return s.Document
}

func (s *RefundApplyRequestRefundPassengerList) GetFirstName() *string {
	return s.FirstName
}

func (s *RefundApplyRequestRefundPassengerList) GetLastName() *string {
	return s.LastName
}

func (s *RefundApplyRequestRefundPassengerList) SetDocument(v string) *RefundApplyRequestRefundPassengerList {
	s.Document = &v
	return s
}

func (s *RefundApplyRequestRefundPassengerList) SetFirstName(v string) *RefundApplyRequestRefundPassengerList {
	s.FirstName = &v
	return s
}

func (s *RefundApplyRequestRefundPassengerList) SetLastName(v string) *RefundApplyRequestRefundPassengerList {
	s.LastName = &v
	return s
}

func (s *RefundApplyRequestRefundPassengerList) Validate() error {
	return dara.Validate(s)
}

type RefundApplyRequestRefundType struct {
	// The array of attachment file URLs. Upload files first by using the dedicated file upload operation to obtain the file URLs.
	//
	// example:
	//
	// [xxx,yyy]
	File []*string `json:"file,omitempty" xml:"file,omitempty" type:"Repeated"`
	// The refund type. Valid values:
	//
	// - 2: Voluntary refund (change of travel plans or decision not to fly).
	//
	// - 5: Involuntary refund due to airline reasons such as flight delay, cancellation, or schedule change.
	//
	// - 6: Involuntary refund due to medical reasons with a certificate from a Grade II Class A hospital or above.
	//
	// Note: Attachments are not mandatory, but providing attachments for involuntary refunds can improve the success rate of the refund application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	RefundTypeId *int32 `json:"refund_type_id,omitempty" xml:"refund_type_id,omitempty"`
	// The remarks.
	//
	// example:
	//
	// remark desc
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s RefundApplyRequestRefundType) String() string {
	return dara.Prettify(s)
}

func (s RefundApplyRequestRefundType) GoString() string {
	return s.String()
}

func (s *RefundApplyRequestRefundType) GetFile() []*string {
	return s.File
}

func (s *RefundApplyRequestRefundType) GetRefundTypeId() *int32 {
	return s.RefundTypeId
}

func (s *RefundApplyRequestRefundType) GetRemark() *string {
	return s.Remark
}

func (s *RefundApplyRequestRefundType) SetFile(v []*string) *RefundApplyRequestRefundType {
	s.File = v
	return s
}

func (s *RefundApplyRequestRefundType) SetRefundTypeId(v int32) *RefundApplyRequestRefundType {
	s.RefundTypeId = &v
	return s
}

func (s *RefundApplyRequestRefundType) SetRemark(v string) *RefundApplyRequestRefundType {
	s.Remark = &v
	return s
}

func (s *RefundApplyRequestRefundType) Validate() error {
	return dara.Validate(s)
}
