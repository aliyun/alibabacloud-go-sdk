// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndPayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *CreateAndPayRequest
	GetAccountNo() *int64
	SetContact(v *CreateAndPayRequestContact) *CreateAndPayRequest
	GetContact() *CreateAndPayRequestContact
	SetExternalOrderNo(v string) *CreateAndPayRequest
	GetExternalOrderNo() *string
	SetGuests(v [][]*CreateAndPayRequestGuests) *CreateAndPayRequest
	GetGuests() [][]*CreateAndPayRequestGuests
	SetItemOfferId(v string) *CreateAndPayRequest
	GetItemOfferId() *string
	SetRoomCount(v int32) *CreateAndPayRequest
	GetRoomCount() *int32
	SetTracerId(v string) *CreateAndPayRequest
	GetTracerId() *string
}

type CreateAndPayRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// This parameter is required.
	Contact *CreateAndPayRequestContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Struct"`
	// example:
	//
	// EXT_ORDER_001
	ExternalOrderNo *string `json:"ExternalOrderNo,omitempty" xml:"ExternalOrderNo,omitempty"`
	// This parameter is required.
	Guests [][]*CreateAndPayRequestGuests `json:"Guests,omitempty" xml:"Guests,omitempty" type:"Repeated"`
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
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s CreateAndPayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPayRequest) GoString() string {
	return s.String()
}

func (s *CreateAndPayRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *CreateAndPayRequest) GetContact() *CreateAndPayRequestContact {
	return s.Contact
}

func (s *CreateAndPayRequest) GetExternalOrderNo() *string {
	return s.ExternalOrderNo
}

func (s *CreateAndPayRequest) GetGuests() [][]*CreateAndPayRequestGuests {
	return s.Guests
}

func (s *CreateAndPayRequest) GetItemOfferId() *string {
	return s.ItemOfferId
}

func (s *CreateAndPayRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *CreateAndPayRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *CreateAndPayRequest) SetAccountNo(v int64) *CreateAndPayRequest {
	s.AccountNo = &v
	return s
}

func (s *CreateAndPayRequest) SetContact(v *CreateAndPayRequestContact) *CreateAndPayRequest {
	s.Contact = v
	return s
}

func (s *CreateAndPayRequest) SetExternalOrderNo(v string) *CreateAndPayRequest {
	s.ExternalOrderNo = &v
	return s
}

func (s *CreateAndPayRequest) SetGuests(v [][]*CreateAndPayRequestGuests) *CreateAndPayRequest {
	s.Guests = v
	return s
}

func (s *CreateAndPayRequest) SetItemOfferId(v string) *CreateAndPayRequest {
	s.ItemOfferId = &v
	return s
}

func (s *CreateAndPayRequest) SetRoomCount(v int32) *CreateAndPayRequest {
	s.RoomCount = &v
	return s
}

func (s *CreateAndPayRequest) SetTracerId(v string) *CreateAndPayRequest {
	s.TracerId = &v
	return s
}

func (s *CreateAndPayRequest) Validate() error {
	if s.Contact != nil {
		if err := s.Contact.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAndPayRequestContact struct {
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

func (s CreateAndPayRequestContact) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPayRequestContact) GoString() string {
	return s.String()
}

func (s *CreateAndPayRequestContact) GetEmail() *string {
	return s.Email
}

func (s *CreateAndPayRequestContact) GetFirstName() *string {
	return s.FirstName
}

func (s *CreateAndPayRequestContact) GetLastName() *string {
	return s.LastName
}

func (s *CreateAndPayRequestContact) GetPhone() *string {
	return s.Phone
}

func (s *CreateAndPayRequestContact) GetTracerId() *string {
	return s.TracerId
}

func (s *CreateAndPayRequestContact) SetEmail(v string) *CreateAndPayRequestContact {
	s.Email = &v
	return s
}

func (s *CreateAndPayRequestContact) SetFirstName(v string) *CreateAndPayRequestContact {
	s.FirstName = &v
	return s
}

func (s *CreateAndPayRequestContact) SetLastName(v string) *CreateAndPayRequestContact {
	s.LastName = &v
	return s
}

func (s *CreateAndPayRequestContact) SetPhone(v string) *CreateAndPayRequestContact {
	s.Phone = &v
	return s
}

func (s *CreateAndPayRequestContact) SetTracerId(v string) *CreateAndPayRequestContact {
	s.TracerId = &v
	return s
}

func (s *CreateAndPayRequestContact) Validate() error {
	return dara.Validate(s)
}

type CreateAndPayRequestGuests struct {
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
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s CreateAndPayRequestGuests) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPayRequestGuests) GoString() string {
	return s.String()
}

func (s *CreateAndPayRequestGuests) GetFirstName() *string {
	return s.FirstName
}

func (s *CreateAndPayRequestGuests) GetLastName() *string {
	return s.LastName
}

func (s *CreateAndPayRequestGuests) GetTracerId() *string {
	return s.TracerId
}

func (s *CreateAndPayRequestGuests) SetFirstName(v string) *CreateAndPayRequestGuests {
	s.FirstName = &v
	return s
}

func (s *CreateAndPayRequestGuests) SetLastName(v string) *CreateAndPayRequestGuests {
	s.LastName = &v
	return s
}

func (s *CreateAndPayRequestGuests) SetTracerId(v string) *CreateAndPayRequestGuests {
	s.TracerId = &v
	return s
}

func (s *CreateAndPayRequestGuests) Validate() error {
	return dara.Validate(s)
}
