// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBookShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactShrink(v string) *BookShrinkRequest
	GetContactShrink() *string
	SetOutOrderNum(v string) *BookShrinkRequest
	GetOutOrderNum() *string
	SetPassengerAncillaryPurchaseMapListShrink(v string) *BookShrinkRequest
	GetPassengerAncillaryPurchaseMapListShrink() *string
	SetPassengerListShrink(v string) *BookShrinkRequest
	GetPassengerListShrink() *string
	SetSolutionId(v string) *BookShrinkRequest
	GetSolutionId() *string
}

type BookShrinkRequest struct {
	// contact information
	//
	// This parameter is required.
	ContactShrink *string `json:"contact,omitempty" xml:"contact,omitempty"`
	// external order number(buyer customization)
	//
	// This parameter is required.
	//
	// example:
	//
	// x091-2023-0220-j-0001
	OutOrderNum *string `json:"out_order_num,omitempty" xml:"out_order_num,omitempty"`
	// passenger-ancillary purchase relationship
	PassengerAncillaryPurchaseMapListShrink *string `json:"passenger_ancillary_purchase_map_list,omitempty" xml:"passenger_ancillary_purchase_map_list,omitempty"`
	// passenger list
	//
	// This parameter is required.
	PassengerListShrink *string `json:"passenger_list,omitempty" xml:"passenger_list,omitempty"`
	// solution_id returned by Enrich
	//
	// This parameter is required.
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s BookShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BookShrinkRequest) GoString() string {
	return s.String()
}

func (s *BookShrinkRequest) GetContactShrink() *string {
	return s.ContactShrink
}

func (s *BookShrinkRequest) GetOutOrderNum() *string {
	return s.OutOrderNum
}

func (s *BookShrinkRequest) GetPassengerAncillaryPurchaseMapListShrink() *string {
	return s.PassengerAncillaryPurchaseMapListShrink
}

func (s *BookShrinkRequest) GetPassengerListShrink() *string {
	return s.PassengerListShrink
}

func (s *BookShrinkRequest) GetSolutionId() *string {
	return s.SolutionId
}

func (s *BookShrinkRequest) SetContactShrink(v string) *BookShrinkRequest {
	s.ContactShrink = &v
	return s
}

func (s *BookShrinkRequest) SetOutOrderNum(v string) *BookShrinkRequest {
	s.OutOrderNum = &v
	return s
}

func (s *BookShrinkRequest) SetPassengerAncillaryPurchaseMapListShrink(v string) *BookShrinkRequest {
	s.PassengerAncillaryPurchaseMapListShrink = &v
	return s
}

func (s *BookShrinkRequest) SetPassengerListShrink(v string) *BookShrinkRequest {
	s.PassengerListShrink = &v
	return s
}

func (s *BookShrinkRequest) SetSolutionId(v string) *BookShrinkRequest {
	s.SolutionId = &v
	return s
}

func (s *BookShrinkRequest) Validate() error {
	return dara.Validate(s)
}
