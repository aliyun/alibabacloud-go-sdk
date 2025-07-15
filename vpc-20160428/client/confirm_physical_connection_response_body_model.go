// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmPhysicalConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfirmPhysicalConnectionResponseBody
	GetRequestId() *string
}

type ConfirmPhysicalConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// API-20365164-5b0d-460a-83c2-2189972b****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfirmPhysicalConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfirmPhysicalConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmPhysicalConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmPhysicalConnectionResponseBody) SetRequestId(v string) *ConfirmPhysicalConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmPhysicalConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
