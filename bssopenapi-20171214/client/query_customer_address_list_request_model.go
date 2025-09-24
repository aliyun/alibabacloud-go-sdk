// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomerAddressListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryCustomerAddressListRequest
	GetOwnerId() *int64
}

type QueryCustomerAddressListRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s QueryCustomerAddressListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomerAddressListRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomerAddressListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryCustomerAddressListRequest) SetOwnerId(v int64) *QueryCustomerAddressListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCustomerAddressListRequest) Validate() error {
	return dara.Validate(s)
}
