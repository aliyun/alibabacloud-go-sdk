// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindExhibitorRfidPopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v int64) *BindExhibitorRfidPopRequest
	GetActivityId() *int64
	SetDeviceId(v string) *BindExhibitorRfidPopRequest
	GetDeviceId() *string
	SetGmtCreate(v string) *BindExhibitorRfidPopRequest
	GetGmtCreate() *string
	SetGmtModified(v string) *BindExhibitorRfidPopRequest
	GetGmtModified() *string
	SetGuestTicketRecordId(v int64) *BindExhibitorRfidPopRequest
	GetGuestTicketRecordId() *int64
	SetId(v int64) *BindExhibitorRfidPopRequest
	GetId() *int64
	SetRfid(v string) *BindExhibitorRfidPopRequest
	GetRfid() *string
	SetTicketCode(v string) *BindExhibitorRfidPopRequest
	GetTicketCode() *string
}

type BindExhibitorRfidPopRequest struct {
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

func (s BindExhibitorRfidPopRequest) String() string {
	return dara.Prettify(s)
}

func (s BindExhibitorRfidPopRequest) GoString() string {
	return s.String()
}

func (s *BindExhibitorRfidPopRequest) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *BindExhibitorRfidPopRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BindExhibitorRfidPopRequest) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *BindExhibitorRfidPopRequest) GetGmtModified() *string {
	return s.GmtModified
}

func (s *BindExhibitorRfidPopRequest) GetGuestTicketRecordId() *int64 {
	return s.GuestTicketRecordId
}

func (s *BindExhibitorRfidPopRequest) GetId() *int64 {
	return s.Id
}

func (s *BindExhibitorRfidPopRequest) GetRfid() *string {
	return s.Rfid
}

func (s *BindExhibitorRfidPopRequest) GetTicketCode() *string {
	return s.TicketCode
}

func (s *BindExhibitorRfidPopRequest) SetActivityId(v int64) *BindExhibitorRfidPopRequest {
	s.ActivityId = &v
	return s
}

func (s *BindExhibitorRfidPopRequest) SetDeviceId(v string) *BindExhibitorRfidPopRequest {
	s.DeviceId = &v
	return s
}

func (s *BindExhibitorRfidPopRequest) SetGmtCreate(v string) *BindExhibitorRfidPopRequest {
	s.GmtCreate = &v
	return s
}

func (s *BindExhibitorRfidPopRequest) SetGmtModified(v string) *BindExhibitorRfidPopRequest {
	s.GmtModified = &v
	return s
}

func (s *BindExhibitorRfidPopRequest) SetGuestTicketRecordId(v int64) *BindExhibitorRfidPopRequest {
	s.GuestTicketRecordId = &v
	return s
}

func (s *BindExhibitorRfidPopRequest) SetId(v int64) *BindExhibitorRfidPopRequest {
	s.Id = &v
	return s
}

func (s *BindExhibitorRfidPopRequest) SetRfid(v string) *BindExhibitorRfidPopRequest {
	s.Rfid = &v
	return s
}

func (s *BindExhibitorRfidPopRequest) SetTicketCode(v string) *BindExhibitorRfidPopRequest {
	s.TicketCode = &v
	return s
}

func (s *BindExhibitorRfidPopRequest) Validate() error {
	return dara.Validate(s)
}
