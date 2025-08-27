// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainNoListSearchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArrLocation(v string) *TrainNoListSearchShrinkRequest
	GetArrLocation() *string
	SetDepDate(v string) *TrainNoListSearchShrinkRequest
	GetDepDate() *string
	SetDepLocation(v string) *TrainNoListSearchShrinkRequest
	GetDepLocation() *string
	SetOptionShrink(v string) *TrainNoListSearchShrinkRequest
	GetOptionShrink() *string
	SetOrderId(v string) *TrainNoListSearchShrinkRequest
	GetOrderId() *string
}

type TrainNoListSearchShrinkRequest struct {
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
	OptionShrink *string `json:"option,omitempty" xml:"option,omitempty"`
	// example:
	//
	// null
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s TrainNoListSearchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainNoListSearchShrinkRequest) GoString() string {
	return s.String()
}

func (s *TrainNoListSearchShrinkRequest) GetArrLocation() *string {
	return s.ArrLocation
}

func (s *TrainNoListSearchShrinkRequest) GetDepDate() *string {
	return s.DepDate
}

func (s *TrainNoListSearchShrinkRequest) GetDepLocation() *string {
	return s.DepLocation
}

func (s *TrainNoListSearchShrinkRequest) GetOptionShrink() *string {
	return s.OptionShrink
}

func (s *TrainNoListSearchShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainNoListSearchShrinkRequest) SetArrLocation(v string) *TrainNoListSearchShrinkRequest {
	s.ArrLocation = &v
	return s
}

func (s *TrainNoListSearchShrinkRequest) SetDepDate(v string) *TrainNoListSearchShrinkRequest {
	s.DepDate = &v
	return s
}

func (s *TrainNoListSearchShrinkRequest) SetDepLocation(v string) *TrainNoListSearchShrinkRequest {
	s.DepLocation = &v
	return s
}

func (s *TrainNoListSearchShrinkRequest) SetOptionShrink(v string) *TrainNoListSearchShrinkRequest {
	s.OptionShrink = &v
	return s
}

func (s *TrainNoListSearchShrinkRequest) SetOrderId(v string) *TrainNoListSearchShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *TrainNoListSearchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
