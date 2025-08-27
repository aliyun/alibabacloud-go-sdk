// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainStopoverSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArrStation(v string) *TrainStopoverSearchRequest
	GetArrStation() *string
	SetDepStation(v string) *TrainStopoverSearchRequest
	GetDepStation() *string
	SetTrainDate(v string) *TrainStopoverSearchRequest
	GetTrainDate() *string
	SetTrainNo(v string) *TrainStopoverSearchRequest
	GetTrainNo() *string
}

type TrainStopoverSearchRequest struct {
	// This parameter is required.
	ArrStation *string `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	// This parameter is required.
	DepStation *string `json:"dep_station,omitempty" xml:"dep_station,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-05-08
	TrainDate *string `json:"train_date,omitempty" xml:"train_date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// k2345
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
}

func (s TrainStopoverSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainStopoverSearchRequest) GoString() string {
	return s.String()
}

func (s *TrainStopoverSearchRequest) GetArrStation() *string {
	return s.ArrStation
}

func (s *TrainStopoverSearchRequest) GetDepStation() *string {
	return s.DepStation
}

func (s *TrainStopoverSearchRequest) GetTrainDate() *string {
	return s.TrainDate
}

func (s *TrainStopoverSearchRequest) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainStopoverSearchRequest) SetArrStation(v string) *TrainStopoverSearchRequest {
	s.ArrStation = &v
	return s
}

func (s *TrainStopoverSearchRequest) SetDepStation(v string) *TrainStopoverSearchRequest {
	s.DepStation = &v
	return s
}

func (s *TrainStopoverSearchRequest) SetTrainDate(v string) *TrainStopoverSearchRequest {
	s.TrainDate = &v
	return s
}

func (s *TrainStopoverSearchRequest) SetTrainNo(v string) *TrainStopoverSearchRequest {
	s.TrainNo = &v
	return s
}

func (s *TrainStopoverSearchRequest) Validate() error {
	return dara.Validate(s)
}
