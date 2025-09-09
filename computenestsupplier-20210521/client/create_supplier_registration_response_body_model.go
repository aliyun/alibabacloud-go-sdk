// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSupplierRegistrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSupplierRegistrationResponseBody
	GetRequestId() *string
}

type CreateSupplierRegistrationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// C4A145D8-xxxx-xxxx-xxxx-9730CDA27578
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSupplierRegistrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSupplierRegistrationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSupplierRegistrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSupplierRegistrationResponseBody) SetRequestId(v string) *CreateSupplierRegistrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSupplierRegistrationResponseBody) Validate() error {
	return dara.Validate(s)
}
