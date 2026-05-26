// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserPoolResponseBody
	GetRequestId() *string
}

type UpdateUserPoolResponseBody struct {
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserPoolResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserPoolResponseBody) SetRequestId(v string) *UpdateUserPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
