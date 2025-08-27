// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyListingSearchV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCabinClassShrink(v string) *FlightModifyListingSearchV2ShrinkRequest
	GetCabinClassShrink() *string
	SetDepDateShrink(v string) *FlightModifyListingSearchV2ShrinkRequest
	GetDepDateShrink() *string
	SetInterfaceCallerIsSupportRetry(v bool) *FlightModifyListingSearchV2ShrinkRequest
	GetInterfaceCallerIsSupportRetry() *bool
	SetIsvName(v string) *FlightModifyListingSearchV2ShrinkRequest
	GetIsvName() *string
	SetOrderId(v int64) *FlightModifyListingSearchV2ShrinkRequest
	GetOrderId() *int64
	SetOutOrderId(v string) *FlightModifyListingSearchV2ShrinkRequest
	GetOutOrderId() *string
	SetPassengerSegmentRelationsShrink(v string) *FlightModifyListingSearchV2ShrinkRequest
	GetPassengerSegmentRelationsShrink() *string
	SetSearchMode(v int32) *FlightModifyListingSearchV2ShrinkRequest
	GetSearchMode() *int32
	SetSearchRetryToken(v string) *FlightModifyListingSearchV2ShrinkRequest
	GetSearchRetryToken() *string
	SetSelectedSegmentsShrink(v string) *FlightModifyListingSearchV2ShrinkRequest
	GetSelectedSegmentsShrink() *string
	SetSessionId(v string) *FlightModifyListingSearchV2ShrinkRequest
	GetSessionId() *string
	SetVoluntary(v bool) *FlightModifyListingSearchV2ShrinkRequest
	GetVoluntary() *bool
}

type FlightModifyListingSearchV2ShrinkRequest struct {
	CabinClassShrink              *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	DepDateShrink                 *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	InterfaceCallerIsSupportRetry *bool   `json:"interface_caller_is_support_retry,omitempty" xml:"interface_caller_is_support_retry,omitempty"`
	// example:
	//
	// name
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 1017002195370467138
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017002195370467137
	OutOrderId                      *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	PassengerSegmentRelationsShrink *string `json:"passenger_segment_relations,omitempty" xml:"passenger_segment_relations,omitempty"`
	// example:
	//
	// 2
	SearchMode             *int32  `json:"search_mode,omitempty" xml:"search_mode,omitempty"`
	SearchRetryToken       *string `json:"search_retry_token,omitempty" xml:"search_retry_token,omitempty"`
	SelectedSegmentsShrink *string `json:"selected_segments,omitempty" xml:"selected_segments,omitempty"`
	// example:
	//
	// a2ffebfe733742aab5c491d960ba3d59
	SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// example:
	//
	// true
	Voluntary *bool `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
}

func (s FlightModifyListingSearchV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ShrinkRequest) GetCabinClassShrink() *string {
	return s.CabinClassShrink
}

func (s *FlightModifyListingSearchV2ShrinkRequest) GetDepDateShrink() *string {
	return s.DepDateShrink
}

func (s *FlightModifyListingSearchV2ShrinkRequest) GetInterfaceCallerIsSupportRetry() *bool {
	return s.InterfaceCallerIsSupportRetry
}

func (s *FlightModifyListingSearchV2ShrinkRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightModifyListingSearchV2ShrinkRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightModifyListingSearchV2ShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightModifyListingSearchV2ShrinkRequest) GetPassengerSegmentRelationsShrink() *string {
	return s.PassengerSegmentRelationsShrink
}

func (s *FlightModifyListingSearchV2ShrinkRequest) GetSearchMode() *int32 {
	return s.SearchMode
}

func (s *FlightModifyListingSearchV2ShrinkRequest) GetSearchRetryToken() *string {
	return s.SearchRetryToken
}

func (s *FlightModifyListingSearchV2ShrinkRequest) GetSelectedSegmentsShrink() *string {
	return s.SelectedSegmentsShrink
}

func (s *FlightModifyListingSearchV2ShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *FlightModifyListingSearchV2ShrinkRequest) GetVoluntary() *bool {
	return s.Voluntary
}

func (s *FlightModifyListingSearchV2ShrinkRequest) SetCabinClassShrink(v string) *FlightModifyListingSearchV2ShrinkRequest {
	s.CabinClassShrink = &v
	return s
}

func (s *FlightModifyListingSearchV2ShrinkRequest) SetDepDateShrink(v string) *FlightModifyListingSearchV2ShrinkRequest {
	s.DepDateShrink = &v
	return s
}

func (s *FlightModifyListingSearchV2ShrinkRequest) SetInterfaceCallerIsSupportRetry(v bool) *FlightModifyListingSearchV2ShrinkRequest {
	s.InterfaceCallerIsSupportRetry = &v
	return s
}

func (s *FlightModifyListingSearchV2ShrinkRequest) SetIsvName(v string) *FlightModifyListingSearchV2ShrinkRequest {
	s.IsvName = &v
	return s
}

func (s *FlightModifyListingSearchV2ShrinkRequest) SetOrderId(v int64) *FlightModifyListingSearchV2ShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *FlightModifyListingSearchV2ShrinkRequest) SetOutOrderId(v string) *FlightModifyListingSearchV2ShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *FlightModifyListingSearchV2ShrinkRequest) SetPassengerSegmentRelationsShrink(v string) *FlightModifyListingSearchV2ShrinkRequest {
	s.PassengerSegmentRelationsShrink = &v
	return s
}

func (s *FlightModifyListingSearchV2ShrinkRequest) SetSearchMode(v int32) *FlightModifyListingSearchV2ShrinkRequest {
	s.SearchMode = &v
	return s
}

func (s *FlightModifyListingSearchV2ShrinkRequest) SetSearchRetryToken(v string) *FlightModifyListingSearchV2ShrinkRequest {
	s.SearchRetryToken = &v
	return s
}

func (s *FlightModifyListingSearchV2ShrinkRequest) SetSelectedSegmentsShrink(v string) *FlightModifyListingSearchV2ShrinkRequest {
	s.SelectedSegmentsShrink = &v
	return s
}

func (s *FlightModifyListingSearchV2ShrinkRequest) SetSessionId(v string) *FlightModifyListingSearchV2ShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *FlightModifyListingSearchV2ShrinkRequest) SetVoluntary(v bool) *FlightModifyListingSearchV2ShrinkRequest {
	s.Voluntary = &v
	return s
}

func (s *FlightModifyListingSearchV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
