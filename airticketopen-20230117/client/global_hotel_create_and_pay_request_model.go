// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelCreateAndPayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelCreateAndPayRequest
	GetAccountNo() *int64
	SetContact(v *GlobalHotelCreateAndPayRequestContact) *GlobalHotelCreateAndPayRequest
	GetContact() *GlobalHotelCreateAndPayRequestContact
	SetExternalOrderNo(v string) *GlobalHotelCreateAndPayRequest
	GetExternalOrderNo() *string
	SetGuests(v [][]*GlobalHotelCreateAndPayRequestGuests) *GlobalHotelCreateAndPayRequest
	GetGuests() [][]*GlobalHotelCreateAndPayRequestGuests
	SetItemOfferId(v string) *GlobalHotelCreateAndPayRequest
	GetItemOfferId() *string
	SetRoomCount(v int32) *GlobalHotelCreateAndPayRequest
	GetRoomCount() *int32
	SetTracerId(v string) *GlobalHotelCreateAndPayRequest
	GetTracerId() *string
}

type GlobalHotelCreateAndPayRequest struct {
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
	Contact *GlobalHotelCreateAndPayRequestContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
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
	Guests [][]*GlobalHotelCreateAndPayRequestGuests `json:"Guests,omitempty" xml:"Guests,omitempty" type:"Repeated"`
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

func (s GlobalHotelCreateAndPayRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCreateAndPayRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelCreateAndPayRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelCreateAndPayRequest) GetContact() *GlobalHotelCreateAndPayRequestContact {
	return s.Contact
}

func (s *GlobalHotelCreateAndPayRequest) GetExternalOrderNo() *string {
	return s.ExternalOrderNo
}

func (s *GlobalHotelCreateAndPayRequest) GetGuests() [][]*GlobalHotelCreateAndPayRequestGuests {
	return s.Guests
}

func (s *GlobalHotelCreateAndPayRequest) GetItemOfferId() *string {
	return s.ItemOfferId
}

func (s *GlobalHotelCreateAndPayRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *GlobalHotelCreateAndPayRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCreateAndPayRequest) SetAccountNo(v int64) *GlobalHotelCreateAndPayRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelCreateAndPayRequest) SetContact(v *GlobalHotelCreateAndPayRequestContact) *GlobalHotelCreateAndPayRequest {
	s.Contact = v
	return s
}

func (s *GlobalHotelCreateAndPayRequest) SetExternalOrderNo(v string) *GlobalHotelCreateAndPayRequest {
	s.ExternalOrderNo = &v
	return s
}

func (s *GlobalHotelCreateAndPayRequest) SetGuests(v [][]*GlobalHotelCreateAndPayRequestGuests) *GlobalHotelCreateAndPayRequest {
	s.Guests = v
	return s
}

func (s *GlobalHotelCreateAndPayRequest) SetItemOfferId(v string) *GlobalHotelCreateAndPayRequest {
	s.ItemOfferId = &v
	return s
}

func (s *GlobalHotelCreateAndPayRequest) SetRoomCount(v int32) *GlobalHotelCreateAndPayRequest {
	s.RoomCount = &v
	return s
}

func (s *GlobalHotelCreateAndPayRequest) SetTracerId(v string) *GlobalHotelCreateAndPayRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCreateAndPayRequest) Validate() error {
	if s.Contact != nil {
		if err := s.Contact.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelCreateAndPayRequestContact struct {
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

func (s GlobalHotelCreateAndPayRequestContact) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCreateAndPayRequestContact) GoString() string {
	return s.String()
}

func (s *GlobalHotelCreateAndPayRequestContact) GetEmail() *string {
	return s.Email
}

func (s *GlobalHotelCreateAndPayRequestContact) GetFirstName() *string {
	return s.FirstName
}

func (s *GlobalHotelCreateAndPayRequestContact) GetLastName() *string {
	return s.LastName
}

func (s *GlobalHotelCreateAndPayRequestContact) GetPhone() *string {
	return s.Phone
}

func (s *GlobalHotelCreateAndPayRequestContact) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCreateAndPayRequestContact) SetEmail(v string) *GlobalHotelCreateAndPayRequestContact {
	s.Email = &v
	return s
}

func (s *GlobalHotelCreateAndPayRequestContact) SetFirstName(v string) *GlobalHotelCreateAndPayRequestContact {
	s.FirstName = &v
	return s
}

func (s *GlobalHotelCreateAndPayRequestContact) SetLastName(v string) *GlobalHotelCreateAndPayRequestContact {
	s.LastName = &v
	return s
}

func (s *GlobalHotelCreateAndPayRequestContact) SetPhone(v string) *GlobalHotelCreateAndPayRequestContact {
	s.Phone = &v
	return s
}

func (s *GlobalHotelCreateAndPayRequestContact) SetTracerId(v string) *GlobalHotelCreateAndPayRequestContact {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCreateAndPayRequestContact) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelCreateAndPayRequestGuests struct {
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

func (s GlobalHotelCreateAndPayRequestGuests) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCreateAndPayRequestGuests) GoString() string {
	return s.String()
}

func (s *GlobalHotelCreateAndPayRequestGuests) GetFirstName() *string {
	return s.FirstName
}

func (s *GlobalHotelCreateAndPayRequestGuests) GetLastName() *string {
	return s.LastName
}

func (s *GlobalHotelCreateAndPayRequestGuests) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCreateAndPayRequestGuests) SetFirstName(v string) *GlobalHotelCreateAndPayRequestGuests {
	s.FirstName = &v
	return s
}

func (s *GlobalHotelCreateAndPayRequestGuests) SetLastName(v string) *GlobalHotelCreateAndPayRequestGuests {
	s.LastName = &v
	return s
}

func (s *GlobalHotelCreateAndPayRequestGuests) SetTracerId(v string) *GlobalHotelCreateAndPayRequestGuests {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCreateAndPayRequestGuests) Validate() error {
	return dara.Validate(s)
}
