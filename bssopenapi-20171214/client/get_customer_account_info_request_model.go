// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerAccountInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetCustomerAccountInfoRequest
	GetOwnerId() *int64
}

type GetCustomerAccountInfoRequest struct {
	// This parameter is required.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s GetCustomerAccountInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerAccountInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetCustomerAccountInfoRequest) SetOwnerId(v int64) *GetCustomerAccountInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCustomerAccountInfoRequest) Validate() error {
	return dara.Validate(s)
}
