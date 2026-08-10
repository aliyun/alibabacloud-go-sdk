// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelCreateOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelCreateOrderRequest
	GetAccountNo() *int64
	SetContact(v *GlobalHotelCreateOrderRequestContact) *GlobalHotelCreateOrderRequest
	GetContact() *GlobalHotelCreateOrderRequestContact
	SetExternalOrderNo(v string) *GlobalHotelCreateOrderRequest
	GetExternalOrderNo() *string
	SetGuests(v [][]*GlobalHotelCreateOrderRequestGuests) *GlobalHotelCreateOrderRequest
	GetGuests() [][]*GlobalHotelCreateOrderRequestGuests
	SetItemOfferId(v string) *GlobalHotelCreateOrderRequest
	GetItemOfferId() *string
	SetRoomCount(v int32) *GlobalHotelCreateOrderRequest
	GetRoomCount() *int32
	SetTracerId(v string) *GlobalHotelCreateOrderRequest
	GetTracerId() *string
}

type GlobalHotelCreateOrderRequest struct {
	// The distributor account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// The contact information.
	//
	// This parameter is required.
	Contact *GlobalHotelCreateOrderRequestContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
	// The external order number.
	//
	// This parameter is required.
	//
	// example:
	//
	// EXT_ORDER_001
	ExternalOrderNo *string `json:"ExternalOrderNo,omitempty" xml:"ExternalOrderNo,omitempty"`
	// The guests grouped by room.
	//
	// This parameter is required.
	Guests [][]*GlobalHotelCreateOrderRequestGuests `json:"Guests,omitempty" xml:"Guests,omitempty" type:"Repeated"`
	// The offer ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// itemOffer_123
	ItemOfferId *string `json:"ItemOfferId,omitempty" xml:"ItemOfferId,omitempty"`
	// The number of rooms.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RoomCount *int32 `json:"RoomCount,omitempty" xml:"RoomCount,omitempty"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelCreateOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCreateOrderRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelCreateOrderRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelCreateOrderRequest) GetContact() *GlobalHotelCreateOrderRequestContact {
	return s.Contact
}

func (s *GlobalHotelCreateOrderRequest) GetExternalOrderNo() *string {
	return s.ExternalOrderNo
}

func (s *GlobalHotelCreateOrderRequest) GetGuests() [][]*GlobalHotelCreateOrderRequestGuests {
	return s.Guests
}

func (s *GlobalHotelCreateOrderRequest) GetItemOfferId() *string {
	return s.ItemOfferId
}

func (s *GlobalHotelCreateOrderRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *GlobalHotelCreateOrderRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCreateOrderRequest) SetAccountNo(v int64) *GlobalHotelCreateOrderRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelCreateOrderRequest) SetContact(v *GlobalHotelCreateOrderRequestContact) *GlobalHotelCreateOrderRequest {
	s.Contact = v
	return s
}

func (s *GlobalHotelCreateOrderRequest) SetExternalOrderNo(v string) *GlobalHotelCreateOrderRequest {
	s.ExternalOrderNo = &v
	return s
}

func (s *GlobalHotelCreateOrderRequest) SetGuests(v [][]*GlobalHotelCreateOrderRequestGuests) *GlobalHotelCreateOrderRequest {
	s.Guests = v
	return s
}

func (s *GlobalHotelCreateOrderRequest) SetItemOfferId(v string) *GlobalHotelCreateOrderRequest {
	s.ItemOfferId = &v
	return s
}

func (s *GlobalHotelCreateOrderRequest) SetRoomCount(v int32) *GlobalHotelCreateOrderRequest {
	s.RoomCount = &v
	return s
}

func (s *GlobalHotelCreateOrderRequest) SetTracerId(v string) *GlobalHotelCreateOrderRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCreateOrderRequest) Validate() error {
	if s.Contact != nil {
		if err := s.Contact.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelCreateOrderRequestContact struct {
	// The email address of the contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// john@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The first name of the contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// John
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// The last name of the contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// Doe
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// The phone number of the contact.
	//
	// example:
	//
	// +86-13800138000
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelCreateOrderRequestContact) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCreateOrderRequestContact) GoString() string {
	return s.String()
}

func (s *GlobalHotelCreateOrderRequestContact) GetEmail() *string {
	return s.Email
}

func (s *GlobalHotelCreateOrderRequestContact) GetFirstName() *string {
	return s.FirstName
}

func (s *GlobalHotelCreateOrderRequestContact) GetLastName() *string {
	return s.LastName
}

func (s *GlobalHotelCreateOrderRequestContact) GetPhone() *string {
	return s.Phone
}

func (s *GlobalHotelCreateOrderRequestContact) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCreateOrderRequestContact) SetEmail(v string) *GlobalHotelCreateOrderRequestContact {
	s.Email = &v
	return s
}

func (s *GlobalHotelCreateOrderRequestContact) SetFirstName(v string) *GlobalHotelCreateOrderRequestContact {
	s.FirstName = &v
	return s
}

func (s *GlobalHotelCreateOrderRequestContact) SetLastName(v string) *GlobalHotelCreateOrderRequestContact {
	s.LastName = &v
	return s
}

func (s *GlobalHotelCreateOrderRequestContact) SetPhone(v string) *GlobalHotelCreateOrderRequestContact {
	s.Phone = &v
	return s
}

func (s *GlobalHotelCreateOrderRequestContact) SetTracerId(v string) *GlobalHotelCreateOrderRequestContact {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCreateOrderRequestContact) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelCreateOrderRequestGuests struct {
	// The first name.
	//
	// This parameter is required.
	//
	// example:
	//
	// John
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// The last name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Doe
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelCreateOrderRequestGuests) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCreateOrderRequestGuests) GoString() string {
	return s.String()
}

func (s *GlobalHotelCreateOrderRequestGuests) GetFirstName() *string {
	return s.FirstName
}

func (s *GlobalHotelCreateOrderRequestGuests) GetLastName() *string {
	return s.LastName
}

func (s *GlobalHotelCreateOrderRequestGuests) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCreateOrderRequestGuests) SetFirstName(v string) *GlobalHotelCreateOrderRequestGuests {
	s.FirstName = &v
	return s
}

func (s *GlobalHotelCreateOrderRequestGuests) SetLastName(v string) *GlobalHotelCreateOrderRequestGuests {
	s.LastName = &v
	return s
}

func (s *GlobalHotelCreateOrderRequestGuests) SetTracerId(v string) *GlobalHotelCreateOrderRequestGuests {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCreateOrderRequestGuests) Validate() error {
	return dara.Validate(s)
}
