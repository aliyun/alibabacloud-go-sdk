// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *TicketQueryProductRequest
	GetAccountNo() *int64
	SetProductId(v string) *TicketQueryProductRequest
	GetProductId() *string
}

type TicketQueryProductRequest struct {
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
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s TicketQueryProductRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryProductRequest) GoString() string {
	return s.String()
}

func (s *TicketQueryProductRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *TicketQueryProductRequest) GetProductId() *string {
	return s.ProductId
}

func (s *TicketQueryProductRequest) SetAccountNo(v int64) *TicketQueryProductRequest {
	s.AccountNo = &v
	return s
}

func (s *TicketQueryProductRequest) SetProductId(v string) *TicketQueryProductRequest {
	s.ProductId = &v
	return s
}

func (s *TicketQueryProductRequest) Validate() error {
	return dara.Validate(s)
}
