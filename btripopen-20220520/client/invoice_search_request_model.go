// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetThirdPartId(v string) *InvoiceSearchRequest
	GetThirdPartId() *string
	SetTitle(v string) *InvoiceSearchRequest
	GetTitle() *string
	SetUserId(v string) *InvoiceSearchRequest
	GetUserId() *string
}

type InvoiceSearchRequest struct {
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 12345
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s InvoiceSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s InvoiceSearchRequest) GoString() string {
	return s.String()
}

func (s *InvoiceSearchRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *InvoiceSearchRequest) GetTitle() *string {
	return s.Title
}

func (s *InvoiceSearchRequest) GetUserId() *string {
	return s.UserId
}

func (s *InvoiceSearchRequest) SetThirdPartId(v string) *InvoiceSearchRequest {
	s.ThirdPartId = &v
	return s
}

func (s *InvoiceSearchRequest) SetTitle(v string) *InvoiceSearchRequest {
	s.Title = &v
	return s
}

func (s *InvoiceSearchRequest) SetUserId(v string) *InvoiceSearchRequest {
	s.UserId = &v
	return s
}

func (s *InvoiceSearchRequest) Validate() error {
	return dara.Validate(s)
}
