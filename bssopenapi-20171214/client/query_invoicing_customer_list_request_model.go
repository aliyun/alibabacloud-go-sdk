// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInvoicingCustomerListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryInvoicingCustomerListRequest
	GetOwnerId() *int64
}

type QueryInvoicingCustomerListRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s QueryInvoicingCustomerListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInvoicingCustomerListRequest) GoString() string {
	return s.String()
}

func (s *QueryInvoicingCustomerListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryInvoicingCustomerListRequest) SetOwnerId(v int64) *QueryInvoicingCustomerListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryInvoicingCustomerListRequest) Validate() error {
	return dara.Validate(s)
}
