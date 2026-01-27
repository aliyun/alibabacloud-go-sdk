// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRCImageResponseBody
	GetRequestId() *string
}

type CreateRCImageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F8D78AED-5050-113C-A46E-7B346*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRCImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRCImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRCImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRCImageResponseBody) SetRequestId(v string) *CreateRCImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRCImageResponseBody) Validate() error {
	return dara.Validate(s)
}
