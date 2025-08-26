// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeTrainTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeTrainTicketResponseBodyData) *RecognizeTrainTicketResponseBody
	GetData() *RecognizeTrainTicketResponseBodyData
	SetRequestId(v string) *RecognizeTrainTicketResponseBody
	GetRequestId() *string
}

type RecognizeTrainTicketResponseBody struct {
	Data *RecognizeTrainTicketResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// BE4B73EA-30A0-4573-A548-3A101B34641A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTrainTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTrainTicketResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTrainTicketResponseBody) GetData() *RecognizeTrainTicketResponseBodyData {
	return s.Data
}

func (s *RecognizeTrainTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeTrainTicketResponseBody) SetData(v *RecognizeTrainTicketResponseBodyData) *RecognizeTrainTicketResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeTrainTicketResponseBody) SetRequestId(v string) *RecognizeTrainTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeTrainTicketResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecognizeTrainTicketResponseBodyData struct {
	Date             *string `json:"Date,omitempty" xml:"Date,omitempty"`
	DepartureStation *string `json:"DepartureStation,omitempty" xml:"DepartureStation,omitempty"`
	Destination      *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	Level            *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// G7350
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// 104.5
	Price *float32 `json:"Price,omitempty" xml:"Price,omitempty"`
	Seat  *string  `json:"Seat,omitempty" xml:"Seat,omitempty"`
}

func (s RecognizeTrainTicketResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTrainTicketResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeTrainTicketResponseBodyData) GetDate() *string {
	return s.Date
}

func (s *RecognizeTrainTicketResponseBodyData) GetDepartureStation() *string {
	return s.DepartureStation
}

func (s *RecognizeTrainTicketResponseBodyData) GetDestination() *string {
	return s.Destination
}

func (s *RecognizeTrainTicketResponseBodyData) GetLevel() *string {
	return s.Level
}

func (s *RecognizeTrainTicketResponseBodyData) GetName() *string {
	return s.Name
}

func (s *RecognizeTrainTicketResponseBodyData) GetNumber() *string {
	return s.Number
}

func (s *RecognizeTrainTicketResponseBodyData) GetPrice() *float32 {
	return s.Price
}

func (s *RecognizeTrainTicketResponseBodyData) GetSeat() *string {
	return s.Seat
}

func (s *RecognizeTrainTicketResponseBodyData) SetDate(v string) *RecognizeTrainTicketResponseBodyData {
	s.Date = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetDepartureStation(v string) *RecognizeTrainTicketResponseBodyData {
	s.DepartureStation = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetDestination(v string) *RecognizeTrainTicketResponseBodyData {
	s.Destination = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetLevel(v string) *RecognizeTrainTicketResponseBodyData {
	s.Level = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetName(v string) *RecognizeTrainTicketResponseBodyData {
	s.Name = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetNumber(v string) *RecognizeTrainTicketResponseBodyData {
	s.Number = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetPrice(v float32) *RecognizeTrainTicketResponseBodyData {
	s.Price = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetSeat(v string) *RecognizeTrainTicketResponseBodyData {
	s.Seat = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) Validate() error {
	return dara.Validate(s)
}
