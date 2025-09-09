// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelServiceRegistrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelServiceRegistrationResponseBody
	GetRequestId() *string
}

type CancelServiceRegistrationResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// C4A145D8-6F6C-532A-9001-9730CDA27578
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelServiceRegistrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelServiceRegistrationResponseBody) GoString() string {
	return s.String()
}

func (s *CancelServiceRegistrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelServiceRegistrationResponseBody) SetRequestId(v string) *CancelServiceRegistrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelServiceRegistrationResponseBody) Validate() error {
	return dara.Validate(s)
}
