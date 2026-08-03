// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *CreateOrderRequest
	GetAccountNo() *int64
	SetContact(v *CreateOrderRequestContact) *CreateOrderRequest
	GetContact() *CreateOrderRequestContact
	SetExternalOrderNo(v string) *CreateOrderRequest
	GetExternalOrderNo() *string
	SetGuests(v [][]*CreateOrderRequestGuests) *CreateOrderRequest
	GetGuests() [][]*CreateOrderRequestGuests
	SetItemOfferId(v string) *CreateOrderRequest
	GetItemOfferId() *string
	SetRoomCount(v int32) *CreateOrderRequest
	GetRoomCount() *int32
	SetTracerId(v string) *CreateOrderRequest
	GetTracerId() *string
}

type CreateOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// This parameter is required.
	Contact *CreateOrderRequestContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
	// example:
	//
	// EXT_ORDER_001
	ExternalOrderNo *string `json:"ExternalOrderNo,omitempty" xml:"ExternalOrderNo,omitempty"`
	// This parameter is required.
	Guests [][]*CreateOrderRequestGuests `json:"Guests,omitempty" xml:"Guests,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// itemOffer_123
	ItemOfferId *string `json:"ItemOfferId,omitempty" xml:"ItemOfferId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RoomCount *int32 `json:"RoomCount,omitempty" xml:"RoomCount,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s CreateOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *CreateOrderRequest) GetContact() *CreateOrderRequestContact {
	return s.Contact
}

func (s *CreateOrderRequest) GetExternalOrderNo() *string {
	return s.ExternalOrderNo
}

func (s *CreateOrderRequest) GetGuests() [][]*CreateOrderRequestGuests {
	return s.Guests
}

func (s *CreateOrderRequest) GetItemOfferId() *string {
	return s.ItemOfferId
}

func (s *CreateOrderRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *CreateOrderRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *CreateOrderRequest) SetAccountNo(v int64) *CreateOrderRequest {
	s.AccountNo = &v
	return s
}

func (s *CreateOrderRequest) SetContact(v *CreateOrderRequestContact) *CreateOrderRequest {
	s.Contact = v
	return s
}

func (s *CreateOrderRequest) SetExternalOrderNo(v string) *CreateOrderRequest {
	s.ExternalOrderNo = &v
	return s
}

func (s *CreateOrderRequest) SetGuests(v [][]*CreateOrderRequestGuests) *CreateOrderRequest {
	s.Guests = v
	return s
}

func (s *CreateOrderRequest) SetItemOfferId(v string) *CreateOrderRequest {
	s.ItemOfferId = &v
	return s
}

func (s *CreateOrderRequest) SetRoomCount(v int32) *CreateOrderRequest {
	s.RoomCount = &v
	return s
}

func (s *CreateOrderRequest) SetTracerId(v string) *CreateOrderRequest {
	s.TracerId = &v
	return s
}

func (s *CreateOrderRequest) Validate() error {
	if s.Contact != nil {
		if err := s.Contact.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrderRequestContact struct {
	// example:
	//
	// john@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// John
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// example:
	//
	// Doe
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// example:
	//
	// +86-13800138000
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s CreateOrderRequestContact) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderRequestContact) GoString() string {
	return s.String()
}

func (s *CreateOrderRequestContact) GetEmail() *string {
	return s.Email
}

func (s *CreateOrderRequestContact) GetFirstName() *string {
	return s.FirstName
}

func (s *CreateOrderRequestContact) GetLastName() *string {
	return s.LastName
}

func (s *CreateOrderRequestContact) GetPhone() *string {
	return s.Phone
}

func (s *CreateOrderRequestContact) GetTracerId() *string {
	return s.TracerId
}

func (s *CreateOrderRequestContact) SetEmail(v string) *CreateOrderRequestContact {
	s.Email = &v
	return s
}

func (s *CreateOrderRequestContact) SetFirstName(v string) *CreateOrderRequestContact {
	s.FirstName = &v
	return s
}

func (s *CreateOrderRequestContact) SetLastName(v string) *CreateOrderRequestContact {
	s.LastName = &v
	return s
}

func (s *CreateOrderRequestContact) SetPhone(v string) *CreateOrderRequestContact {
	s.Phone = &v
	return s
}

func (s *CreateOrderRequestContact) SetTracerId(v string) *CreateOrderRequestContact {
	s.TracerId = &v
	return s
}

func (s *CreateOrderRequestContact) Validate() error {
	return dara.Validate(s)
}

type CreateOrderRequestGuests struct {
	// example:
	//
	// John
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// example:
	//
	// Doe
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s CreateOrderRequestGuests) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderRequestGuests) GoString() string {
	return s.String()
}

func (s *CreateOrderRequestGuests) GetFirstName() *string {
	return s.FirstName
}

func (s *CreateOrderRequestGuests) GetLastName() *string {
	return s.LastName
}

func (s *CreateOrderRequestGuests) GetTracerId() *string {
	return s.TracerId
}

func (s *CreateOrderRequestGuests) SetFirstName(v string) *CreateOrderRequestGuests {
	s.FirstName = &v
	return s
}

func (s *CreateOrderRequestGuests) SetLastName(v string) *CreateOrderRequestGuests {
	s.LastName = &v
	return s
}

func (s *CreateOrderRequestGuests) SetTracerId(v string) *CreateOrderRequestGuests {
	s.TracerId = &v
	return s
}

func (s *CreateOrderRequestGuests) Validate() error {
	return dara.Validate(s)
}
