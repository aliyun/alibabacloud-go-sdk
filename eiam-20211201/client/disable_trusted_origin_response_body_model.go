// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableTrustedOriginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableTrustedOriginResponseBody
	GetRequestId() *string
}

type DisableTrustedOriginResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-example
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableTrustedOriginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableTrustedOriginResponseBody) GoString() string {
	return s.String()
}

func (s *DisableTrustedOriginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableTrustedOriginResponseBody) SetRequestId(v string) *DisableTrustedOriginResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableTrustedOriginResponseBody) Validate() error {
	return dara.Validate(s)
}
