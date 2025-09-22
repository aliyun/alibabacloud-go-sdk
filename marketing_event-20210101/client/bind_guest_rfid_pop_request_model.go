// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindGuestRfidPopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v int64) *BindGuestRfidPopRequest
	GetActivityId() *int64
	SetDeviceId(v string) *BindGuestRfidPopRequest
	GetDeviceId() *string
	SetGmtCreate(v string) *BindGuestRfidPopRequest
	GetGmtCreate() *string
	SetGmtModified(v string) *BindGuestRfidPopRequest
	GetGmtModified() *string
	SetGuestTicketRecordId(v int64) *BindGuestRfidPopRequest
	GetGuestTicketRecordId() *int64
	SetId(v int64) *BindGuestRfidPopRequest
	GetId() *int64
	SetRfid(v string) *BindGuestRfidPopRequest
	GetRfid() *string
	SetTicketCode(v string) *BindGuestRfidPopRequest
	GetTicketCode() *string
}

type BindGuestRfidPopRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4546
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Z10
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 451246
	GuestTicketRecordId *int64 `json:"GuestTicketRecordId,omitempty" xml:"GuestTicketRecordId,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// asdojzopf
	Rfid *string `json:"Rfid,omitempty" xml:"Rfid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4546-100000
	TicketCode *string `json:"TicketCode,omitempty" xml:"TicketCode,omitempty"`
}

func (s BindGuestRfidPopRequest) String() string {
	return dara.Prettify(s)
}

func (s BindGuestRfidPopRequest) GoString() string {
	return s.String()
}

func (s *BindGuestRfidPopRequest) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *BindGuestRfidPopRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BindGuestRfidPopRequest) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *BindGuestRfidPopRequest) GetGmtModified() *string {
	return s.GmtModified
}

func (s *BindGuestRfidPopRequest) GetGuestTicketRecordId() *int64 {
	return s.GuestTicketRecordId
}

func (s *BindGuestRfidPopRequest) GetId() *int64 {
	return s.Id
}

func (s *BindGuestRfidPopRequest) GetRfid() *string {
	return s.Rfid
}

func (s *BindGuestRfidPopRequest) GetTicketCode() *string {
	return s.TicketCode
}

func (s *BindGuestRfidPopRequest) SetActivityId(v int64) *BindGuestRfidPopRequest {
	s.ActivityId = &v
	return s
}

func (s *BindGuestRfidPopRequest) SetDeviceId(v string) *BindGuestRfidPopRequest {
	s.DeviceId = &v
	return s
}

func (s *BindGuestRfidPopRequest) SetGmtCreate(v string) *BindGuestRfidPopRequest {
	s.GmtCreate = &v
	return s
}

func (s *BindGuestRfidPopRequest) SetGmtModified(v string) *BindGuestRfidPopRequest {
	s.GmtModified = &v
	return s
}

func (s *BindGuestRfidPopRequest) SetGuestTicketRecordId(v int64) *BindGuestRfidPopRequest {
	s.GuestTicketRecordId = &v
	return s
}

func (s *BindGuestRfidPopRequest) SetId(v int64) *BindGuestRfidPopRequest {
	s.Id = &v
	return s
}

func (s *BindGuestRfidPopRequest) SetRfid(v string) *BindGuestRfidPopRequest {
	s.Rfid = &v
	return s
}

func (s *BindGuestRfidPopRequest) SetTicketCode(v string) *BindGuestRfidPopRequest {
	s.TicketCode = &v
	return s
}

func (s *BindGuestRfidPopRequest) Validate() error {
	return dara.Validate(s)
}
