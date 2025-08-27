// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainNoInfoSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArrLocation(v string) *TrainNoInfoSearchRequest
	GetArrLocation() *string
	SetDepDate(v string) *TrainNoInfoSearchRequest
	GetDepDate() *string
	SetDepLocation(v string) *TrainNoInfoSearchRequest
	GetDepLocation() *string
	SetLineKey(v string) *TrainNoInfoSearchRequest
	GetLineKey() *string
	SetMiddleDate(v string) *TrainNoInfoSearchRequest
	GetMiddleDate() *string
	SetMiddleStation(v string) *TrainNoInfoSearchRequest
	GetMiddleStation() *string
	SetOrderId(v string) *TrainNoInfoSearchRequest
	GetOrderId() *string
	SetTrainNo(v string) *TrainNoInfoSearchRequest
	GetTrainNo() *string
}

type TrainNoInfoSearchRequest struct {
	// This parameter is required.
	ArrLocation *string `json:"arr_location,omitempty" xml:"arr_location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-08-15
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// This parameter is required.
	DepLocation *string `json:"dep_location,omitempty" xml:"dep_location,omitempty"`
	// example:
	//
	// qwwweqswxxa
	LineKey *string `json:"line_key,omitempty" xml:"line_key,omitempty"`
	// example:
	//
	// 2023-08-15
	MiddleDate    *string `json:"middle_date,omitempty" xml:"middle_date,omitempty"`
	MiddleStation *string `json:"middle_station,omitempty" xml:"middle_station,omitempty"`
	// example:
	//
	// 12342123212
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// K2345
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
}

func (s TrainNoInfoSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainNoInfoSearchRequest) GoString() string {
	return s.String()
}

func (s *TrainNoInfoSearchRequest) GetArrLocation() *string {
	return s.ArrLocation
}

func (s *TrainNoInfoSearchRequest) GetDepDate() *string {
	return s.DepDate
}

func (s *TrainNoInfoSearchRequest) GetDepLocation() *string {
	return s.DepLocation
}

func (s *TrainNoInfoSearchRequest) GetLineKey() *string {
	return s.LineKey
}

func (s *TrainNoInfoSearchRequest) GetMiddleDate() *string {
	return s.MiddleDate
}

func (s *TrainNoInfoSearchRequest) GetMiddleStation() *string {
	return s.MiddleStation
}

func (s *TrainNoInfoSearchRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainNoInfoSearchRequest) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainNoInfoSearchRequest) SetArrLocation(v string) *TrainNoInfoSearchRequest {
	s.ArrLocation = &v
	return s
}

func (s *TrainNoInfoSearchRequest) SetDepDate(v string) *TrainNoInfoSearchRequest {
	s.DepDate = &v
	return s
}

func (s *TrainNoInfoSearchRequest) SetDepLocation(v string) *TrainNoInfoSearchRequest {
	s.DepLocation = &v
	return s
}

func (s *TrainNoInfoSearchRequest) SetLineKey(v string) *TrainNoInfoSearchRequest {
	s.LineKey = &v
	return s
}

func (s *TrainNoInfoSearchRequest) SetMiddleDate(v string) *TrainNoInfoSearchRequest {
	s.MiddleDate = &v
	return s
}

func (s *TrainNoInfoSearchRequest) SetMiddleStation(v string) *TrainNoInfoSearchRequest {
	s.MiddleStation = &v
	return s
}

func (s *TrainNoInfoSearchRequest) SetOrderId(v string) *TrainNoInfoSearchRequest {
	s.OrderId = &v
	return s
}

func (s *TrainNoInfoSearchRequest) SetTrainNo(v string) *TrainNoInfoSearchRequest {
	s.TrainNo = &v
	return s
}

func (s *TrainNoInfoSearchRequest) Validate() error {
	return dara.Validate(s)
}
