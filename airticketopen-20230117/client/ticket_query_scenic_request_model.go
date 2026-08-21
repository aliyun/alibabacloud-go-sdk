// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryScenicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *TicketQueryScenicRequest
	GetAccountNo() *int64
	SetScenicId(v int64) *TicketQueryScenicRequest
	GetScenicId() *int64
}

type TicketQueryScenicRequest struct {
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

func (s TicketQueryScenicRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryScenicRequest) GoString() string {
	return s.String()
}

func (s *TicketQueryScenicRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *TicketQueryScenicRequest) GetScenicId() *int64 {
	return s.ScenicId
}

func (s *TicketQueryScenicRequest) SetAccountNo(v int64) *TicketQueryScenicRequest {
	s.AccountNo = &v
	return s
}

func (s *TicketQueryScenicRequest) SetScenicId(v int64) *TicketQueryScenicRequest {
	s.ScenicId = &v
	return s
}

func (s *TicketQueryScenicRequest) Validate() error {
	return dara.Validate(s)
}
