// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrustedOriginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTrustedOriginResponseBody
	GetRequestId() *string
}

type DeleteTrustedOriginResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-example
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrustedOriginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrustedOriginResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrustedOriginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTrustedOriginResponseBody) SetRequestId(v string) *DeleteTrustedOriginResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTrustedOriginResponseBody) Validate() error {
	return dara.Validate(s)
}
