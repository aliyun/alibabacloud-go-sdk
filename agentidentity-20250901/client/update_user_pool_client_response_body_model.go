// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserPoolClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserPoolClientResponseBody
	GetRequestId() *string
}

type UpdateUserPoolClientResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserPoolClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserPoolClientResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserPoolClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserPoolClientResponseBody) SetRequestId(v string) *UpdateUserPoolClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserPoolClientResponseBody) Validate() error {
	return dara.Validate(s)
}
