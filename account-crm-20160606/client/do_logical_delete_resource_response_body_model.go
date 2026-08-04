// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDoLogicalDeleteResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBid(v string) *DoLogicalDeleteResourceResponseBody
	GetBid() *string
	SetCountry(v string) *DoLogicalDeleteResourceResponseBody
	GetCountry() *string
	SetGmtWakeup(v string) *DoLogicalDeleteResourceResponseBody
	GetGmtWakeup() *string
	SetHid(v int64) *DoLogicalDeleteResourceResponseBody
	GetHid() *int64
	SetInterrupt(v bool) *DoLogicalDeleteResourceResponseBody
	GetInterrupt() *bool
	SetInvoker(v string) *DoLogicalDeleteResourceResponseBody
	GetInvoker() *string
	SetMessage(v string) *DoLogicalDeleteResourceResponseBody
	GetMessage() *string
	SetPk(v string) *DoLogicalDeleteResourceResponseBody
	GetPk() *string
	SetSuccess(v bool) *DoLogicalDeleteResourceResponseBody
	GetSuccess() *bool
	SetTaskExtraData(v string) *DoLogicalDeleteResourceResponseBody
	GetTaskExtraData() *string
	SetTaskIdentifier(v string) *DoLogicalDeleteResourceResponseBody
	GetTaskIdentifier() *string
}

type DoLogicalDeleteResourceResponseBody struct {
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	GmtWakeup      *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	Hid            *int64  `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Interrupt      *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
	Invoker        *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Pk             *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskExtraData  *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
}

func (s DoLogicalDeleteResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DoLogicalDeleteResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DoLogicalDeleteResourceResponseBody) GetBid() *string {
	return s.Bid
}

func (s *DoLogicalDeleteResourceResponseBody) GetCountry() *string {
	return s.Country
}

func (s *DoLogicalDeleteResourceResponseBody) GetGmtWakeup() *string {
	return s.GmtWakeup
}

func (s *DoLogicalDeleteResourceResponseBody) GetHid() *int64 {
	return s.Hid
}

func (s *DoLogicalDeleteResourceResponseBody) GetInterrupt() *bool {
	return s.Interrupt
}

func (s *DoLogicalDeleteResourceResponseBody) GetInvoker() *string {
	return s.Invoker
}

func (s *DoLogicalDeleteResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DoLogicalDeleteResourceResponseBody) GetPk() *string {
	return s.Pk
}

func (s *DoLogicalDeleteResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DoLogicalDeleteResourceResponseBody) GetTaskExtraData() *string {
	return s.TaskExtraData
}

func (s *DoLogicalDeleteResourceResponseBody) GetTaskIdentifier() *string {
	return s.TaskIdentifier
}

func (s *DoLogicalDeleteResourceResponseBody) SetBid(v string) *DoLogicalDeleteResourceResponseBody {
	s.Bid = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetCountry(v string) *DoLogicalDeleteResourceResponseBody {
	s.Country = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetGmtWakeup(v string) *DoLogicalDeleteResourceResponseBody {
	s.GmtWakeup = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetHid(v int64) *DoLogicalDeleteResourceResponseBody {
	s.Hid = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetInterrupt(v bool) *DoLogicalDeleteResourceResponseBody {
	s.Interrupt = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetInvoker(v string) *DoLogicalDeleteResourceResponseBody {
	s.Invoker = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetMessage(v string) *DoLogicalDeleteResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetPk(v string) *DoLogicalDeleteResourceResponseBody {
	s.Pk = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetSuccess(v bool) *DoLogicalDeleteResourceResponseBody {
	s.Success = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetTaskExtraData(v string) *DoLogicalDeleteResourceResponseBody {
	s.TaskExtraData = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) SetTaskIdentifier(v string) *DoLogicalDeleteResourceResponseBody {
	s.TaskIdentifier = &v
	return s
}

func (s *DoLogicalDeleteResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
