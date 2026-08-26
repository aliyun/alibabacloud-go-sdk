// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrustedOriginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTrustedOriginResponseBody
	GetRequestId() *string
}

type UpdateTrustedOriginResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-example
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTrustedOriginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrustedOriginResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrustedOriginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTrustedOriginResponseBody) SetRequestId(v string) *UpdateTrustedOriginResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTrustedOriginResponseBody) Validate() error {
	return dara.Validate(s)
}
