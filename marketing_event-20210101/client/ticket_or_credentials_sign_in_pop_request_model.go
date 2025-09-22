// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketOrCredentialsSignInPopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v string) *TicketOrCredentialsSignInPopRequest
	GetActivityId() *string
	SetCode(v string) *TicketOrCredentialsSignInPopRequest
	GetCode() *string
	SetConferenceName(v string) *TicketOrCredentialsSignInPopRequest
	GetConferenceName() *string
	SetDeviceId(v string) *TicketOrCredentialsSignInPopRequest
	GetDeviceId() *string
	SetEntryName(v string) *TicketOrCredentialsSignInPopRequest
	GetEntryName() *string
	SetIdcard(v string) *TicketOrCredentialsSignInPopRequest
	GetIdcard() *string
	SetSignTime(v string) *TicketOrCredentialsSignInPopRequest
	GetSignTime() *string
	SetType(v int32) *TicketOrCredentialsSignInPopRequest
	GetType() *int32
}

type TicketOrCredentialsSignInPopRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4546
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// 4546-100000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 会议名称
	ConferenceName *string `json:"ConferenceName,omitempty" xml:"ConferenceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Z10
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 入口名称
	EntryName *string `json:"EntryName,omitempty" xml:"EntryName,omitempty"`
	// example:
	//
	// 429005111100000
	Idcard *string `json:"Idcard,omitempty" xml:"Idcard,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-09-25 14:11
	SignTime *string `json:"SignTime,omitempty" xml:"SignTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s TicketOrCredentialsSignInPopRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketOrCredentialsSignInPopRequest) GoString() string {
	return s.String()
}

func (s *TicketOrCredentialsSignInPopRequest) GetActivityId() *string {
	return s.ActivityId
}

func (s *TicketOrCredentialsSignInPopRequest) GetCode() *string {
	return s.Code
}

func (s *TicketOrCredentialsSignInPopRequest) GetConferenceName() *string {
	return s.ConferenceName
}

func (s *TicketOrCredentialsSignInPopRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *TicketOrCredentialsSignInPopRequest) GetEntryName() *string {
	return s.EntryName
}

func (s *TicketOrCredentialsSignInPopRequest) GetIdcard() *string {
	return s.Idcard
}

func (s *TicketOrCredentialsSignInPopRequest) GetSignTime() *string {
	return s.SignTime
}

func (s *TicketOrCredentialsSignInPopRequest) GetType() *int32 {
	return s.Type
}

func (s *TicketOrCredentialsSignInPopRequest) SetActivityId(v string) *TicketOrCredentialsSignInPopRequest {
	s.ActivityId = &v
	return s
}

func (s *TicketOrCredentialsSignInPopRequest) SetCode(v string) *TicketOrCredentialsSignInPopRequest {
	s.Code = &v
	return s
}

func (s *TicketOrCredentialsSignInPopRequest) SetConferenceName(v string) *TicketOrCredentialsSignInPopRequest {
	s.ConferenceName = &v
	return s
}

func (s *TicketOrCredentialsSignInPopRequest) SetDeviceId(v string) *TicketOrCredentialsSignInPopRequest {
	s.DeviceId = &v
	return s
}

func (s *TicketOrCredentialsSignInPopRequest) SetEntryName(v string) *TicketOrCredentialsSignInPopRequest {
	s.EntryName = &v
	return s
}

func (s *TicketOrCredentialsSignInPopRequest) SetIdcard(v string) *TicketOrCredentialsSignInPopRequest {
	s.Idcard = &v
	return s
}

func (s *TicketOrCredentialsSignInPopRequest) SetSignTime(v string) *TicketOrCredentialsSignInPopRequest {
	s.SignTime = &v
	return s
}

func (s *TicketOrCredentialsSignInPopRequest) SetType(v int32) *TicketOrCredentialsSignInPopRequest {
	s.Type = &v
	return s
}

func (s *TicketOrCredentialsSignInPopRequest) Validate() error {
	return dara.Validate(s)
}
