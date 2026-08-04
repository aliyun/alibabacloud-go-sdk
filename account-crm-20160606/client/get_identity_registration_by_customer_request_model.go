// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityRegistrationByCustomerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerId(v string) *GetIdentityRegistrationByCustomerRequest
	GetCustomerId() *string
}

type GetIdentityRegistrationByCustomerRequest struct {
	// This parameter is required.
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
}

func (s GetIdentityRegistrationByCustomerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityRegistrationByCustomerRequest) GoString() string {
	return s.String()
}

func (s *GetIdentityRegistrationByCustomerRequest) GetCustomerId() *string {
	return s.CustomerId
}

func (s *GetIdentityRegistrationByCustomerRequest) SetCustomerId(v string) *GetIdentityRegistrationByCustomerRequest {
	s.CustomerId = &v
	return s
}

func (s *GetIdentityRegistrationByCustomerRequest) Validate() error {
	return dara.Validate(s)
}
