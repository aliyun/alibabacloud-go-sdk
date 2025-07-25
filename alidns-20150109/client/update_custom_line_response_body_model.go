// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCustomLineResponseBody
	GetRequestId() *string
}

type UpdateCustomLineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomLineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomLineResponseBody) SetRequestId(v string) *UpdateCustomLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomLineResponseBody) Validate() error {
	return dara.Validate(s)
}
