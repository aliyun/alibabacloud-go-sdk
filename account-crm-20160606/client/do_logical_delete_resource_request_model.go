// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDoLogicalDeleteResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBid(v string) *DoLogicalDeleteResourceRequest
	GetBid() *string
	SetCountry(v string) *DoLogicalDeleteResourceRequest
	GetCountry() *string
	SetGmtWakeup(v string) *DoLogicalDeleteResourceRequest
	GetGmtWakeup() *string
	SetHid(v int64) *DoLogicalDeleteResourceRequest
	GetHid() *int64
	SetInterrupt(v bool) *DoLogicalDeleteResourceRequest
	GetInterrupt() *bool
	SetInvoker(v string) *DoLogicalDeleteResourceRequest
	GetInvoker() *string
	SetMessage(v string) *DoLogicalDeleteResourceRequest
	GetMessage() *string
	SetPk(v string) *DoLogicalDeleteResourceRequest
	GetPk() *string
	SetSuccess(v bool) *DoLogicalDeleteResourceRequest
	GetSuccess() *bool
	SetTaskExtraData(v string) *DoLogicalDeleteResourceRequest
	GetTaskExtraData() *string
	SetTaskIdentifier(v string) *DoLogicalDeleteResourceRequest
	GetTaskIdentifier() *string
}

type DoLogicalDeleteResourceRequest struct {
	// This parameter is required.
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// This parameter is required.
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// This parameter is required.
	GmtWakeup *string `json:"GmtWakeup,omitempty" xml:"GmtWakeup,omitempty"`
	// This parameter is required.
	Hid       *int64 `json:"Hid,omitempty" xml:"Hid,omitempty"`
	Interrupt *bool  `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
	// This parameter is required.
	Invoker *string `json:"Invoker,omitempty" xml:"Invoker,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	Pk      *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	// This parameter is required.
	TaskExtraData *string `json:"TaskExtraData,omitempty" xml:"TaskExtraData,omitempty"`
	// This parameter is required.
	TaskIdentifier *string `json:"TaskIdentifier,omitempty" xml:"TaskIdentifier,omitempty"`
}

func (s DoLogicalDeleteResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DoLogicalDeleteResourceRequest) GoString() string {
	return s.String()
}

func (s *DoLogicalDeleteResourceRequest) GetBid() *string {
	return s.Bid
}

func (s *DoLogicalDeleteResourceRequest) GetCountry() *string {
	return s.Country
}

func (s *DoLogicalDeleteResourceRequest) GetGmtWakeup() *string {
	return s.GmtWakeup
}

func (s *DoLogicalDeleteResourceRequest) GetHid() *int64 {
	return s.Hid
}

func (s *DoLogicalDeleteResourceRequest) GetInterrupt() *bool {
	return s.Interrupt
}

func (s *DoLogicalDeleteResourceRequest) GetInvoker() *string {
	return s.Invoker
}

func (s *DoLogicalDeleteResourceRequest) GetMessage() *string {
	return s.Message
}

func (s *DoLogicalDeleteResourceRequest) GetPk() *string {
	return s.Pk
}

func (s *DoLogicalDeleteResourceRequest) GetSuccess() *bool {
	return s.Success
}

func (s *DoLogicalDeleteResourceRequest) GetTaskExtraData() *string {
	return s.TaskExtraData
}

func (s *DoLogicalDeleteResourceRequest) GetTaskIdentifier() *string {
	return s.TaskIdentifier
}

func (s *DoLogicalDeleteResourceRequest) SetBid(v string) *DoLogicalDeleteResourceRequest {
	s.Bid = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetCountry(v string) *DoLogicalDeleteResourceRequest {
	s.Country = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetGmtWakeup(v string) *DoLogicalDeleteResourceRequest {
	s.GmtWakeup = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetHid(v int64) *DoLogicalDeleteResourceRequest {
	s.Hid = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetInterrupt(v bool) *DoLogicalDeleteResourceRequest {
	s.Interrupt = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetInvoker(v string) *DoLogicalDeleteResourceRequest {
	s.Invoker = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetMessage(v string) *DoLogicalDeleteResourceRequest {
	s.Message = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetPk(v string) *DoLogicalDeleteResourceRequest {
	s.Pk = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetSuccess(v bool) *DoLogicalDeleteResourceRequest {
	s.Success = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetTaskExtraData(v string) *DoLogicalDeleteResourceRequest {
	s.TaskExtraData = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) SetTaskIdentifier(v string) *DoLogicalDeleteResourceRequest {
	s.TaskIdentifier = &v
	return s
}

func (s *DoLogicalDeleteResourceRequest) Validate() error {
	return dara.Validate(s)
}
