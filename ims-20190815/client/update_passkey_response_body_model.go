// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePasskeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePasskeyResponseBody
	GetRequestId() *string
}

type UpdatePasskeyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePasskeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePasskeyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePasskeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePasskeyResponseBody) SetRequestId(v string) *UpdatePasskeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePasskeyResponseBody) Validate() error {
	return dara.Validate(s)
}
