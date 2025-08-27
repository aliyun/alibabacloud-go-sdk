// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainNoListSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArrLocation(v string) *TrainNoListSearchRequest
	GetArrLocation() *string
	SetDepDate(v string) *TrainNoListSearchRequest
	GetDepDate() *string
	SetDepLocation(v string) *TrainNoListSearchRequest
	GetDepLocation() *string
	SetOption(v *TrainNoListSearchRequestOption) *TrainNoListSearchRequest
	GetOption() *TrainNoListSearchRequestOption
	SetOrderId(v string) *TrainNoListSearchRequest
	GetOrderId() *string
}

type TrainNoListSearchRequest struct {
	// This parameter is required.
	ArrLocation *string `json:"arr_location,omitempty" xml:"arr_location,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-05-16
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// This parameter is required.
	DepLocation *string `json:"dep_location,omitempty" xml:"dep_location,omitempty"`
	// This parameter is required.
	Option *TrainNoListSearchRequestOption `json:"option,omitempty" xml:"option,omitempty" type:"Struct"`
	// example:
	//
	// null
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s TrainNoListSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainNoListSearchRequest) GoString() string {
	return s.String()
}

func (s *TrainNoListSearchRequest) GetArrLocation() *string {
	return s.ArrLocation
}

func (s *TrainNoListSearchRequest) GetDepDate() *string {
	return s.DepDate
}

func (s *TrainNoListSearchRequest) GetDepLocation() *string {
	return s.DepLocation
}

func (s *TrainNoListSearchRequest) GetOption() *TrainNoListSearchRequestOption {
	return s.Option
}

func (s *TrainNoListSearchRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainNoListSearchRequest) SetArrLocation(v string) *TrainNoListSearchRequest {
	s.ArrLocation = &v
	return s
}

func (s *TrainNoListSearchRequest) SetDepDate(v string) *TrainNoListSearchRequest {
	s.DepDate = &v
	return s
}

func (s *TrainNoListSearchRequest) SetDepLocation(v string) *TrainNoListSearchRequest {
	s.DepLocation = &v
	return s
}

func (s *TrainNoListSearchRequest) SetOption(v *TrainNoListSearchRequestOption) *TrainNoListSearchRequest {
	s.Option = v
	return s
}

func (s *TrainNoListSearchRequest) SetOrderId(v string) *TrainNoListSearchRequest {
	s.OrderId = &v
	return s
}

func (s *TrainNoListSearchRequest) Validate() error {
	return dara.Validate(s)
}

type TrainNoListSearchRequestOption struct {
	// example:
	//
	// false
	NeedTransfer *bool `json:"need_transfer,omitempty" xml:"need_transfer,omitempty"`
}

func (s TrainNoListSearchRequestOption) String() string {
	return dara.Prettify(s)
}

func (s TrainNoListSearchRequestOption) GoString() string {
	return s.String()
}

func (s *TrainNoListSearchRequestOption) GetNeedTransfer() *bool {
	return s.NeedTransfer
}

func (s *TrainNoListSearchRequestOption) SetNeedTransfer(v bool) *TrainNoListSearchRequestOption {
	s.NeedTransfer = &v
	return s
}

func (s *TrainNoListSearchRequestOption) Validate() error {
	return dara.Validate(s)
}
