// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateConnectionResponseBody
	GetRequestId() *string
}

type UpdateConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A0F049F0-8D69-5BAC-8F10-B******A34C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConnectionResponseBody) SetRequestId(v string) *UpdateConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
