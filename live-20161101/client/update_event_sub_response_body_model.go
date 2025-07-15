// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventSubResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEventSubResponseBody
	GetRequestId() *string
}

type UpdateEventSubResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AE050E24-BE9B-1E79-BB30-7EA0BBAE7F08
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEventSubResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventSubResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventSubResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEventSubResponseBody) SetRequestId(v string) *UpdateEventSubResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventSubResponseBody) Validate() error {
	return dara.Validate(s)
}
