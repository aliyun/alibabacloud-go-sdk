// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryShelfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *TicketQueryShelfRequest
	GetAccountNo() *int64
	SetScenicId(v int64) *TicketQueryShelfRequest
	GetScenicId() *int64
}

type TicketQueryShelfRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ScenicId *int64 `json:"ScenicId,omitempty" xml:"ScenicId,omitempty"`
}

func (s TicketQueryShelfRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryShelfRequest) GoString() string {
	return s.String()
}

func (s *TicketQueryShelfRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *TicketQueryShelfRequest) GetScenicId() *int64 {
	return s.ScenicId
}

func (s *TicketQueryShelfRequest) SetAccountNo(v int64) *TicketQueryShelfRequest {
	s.AccountNo = &v
	return s
}

func (s *TicketQueryShelfRequest) SetScenicId(v int64) *TicketQueryShelfRequest {
	s.ScenicId = &v
	return s
}

func (s *TicketQueryShelfRequest) Validate() error {
	return dara.Validate(s)
}
