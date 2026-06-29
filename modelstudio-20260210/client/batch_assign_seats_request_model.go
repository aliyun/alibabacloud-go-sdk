// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAssignSeatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *BatchAssignSeatsRequest
	GetAccountIds() []*string
	SetLocale(v string) *BatchAssignSeatsRequest
	GetLocale() *string
	SetSeatType(v string) *BatchAssignSeatsRequest
	GetSeatType() *string
}

type BatchAssignSeatsRequest struct {
	// The list of target member IDs.
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The language. Valid values: zh-CN and en-US.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The seat type. Valid values:
	//
	// - standard: standard seat
	//
	// - pro: pro seat
	//
	// - max: premium seat
	//
	// This parameter is required.
	//
	// example:
	//
	// standard
	SeatType *string `json:"SeatType,omitempty" xml:"SeatType,omitempty"`
}

func (s BatchAssignSeatsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchAssignSeatsRequest) GoString() string {
	return s.String()
}

func (s *BatchAssignSeatsRequest) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *BatchAssignSeatsRequest) GetLocale() *string {
	return s.Locale
}

func (s *BatchAssignSeatsRequest) GetSeatType() *string {
	return s.SeatType
}

func (s *BatchAssignSeatsRequest) SetAccountIds(v []*string) *BatchAssignSeatsRequest {
	s.AccountIds = v
	return s
}

func (s *BatchAssignSeatsRequest) SetLocale(v string) *BatchAssignSeatsRequest {
	s.Locale = &v
	return s
}

func (s *BatchAssignSeatsRequest) SetSeatType(v string) *BatchAssignSeatsRequest {
	s.SeatType = &v
	return s
}

func (s *BatchAssignSeatsRequest) Validate() error {
	return dara.Validate(s)
}
