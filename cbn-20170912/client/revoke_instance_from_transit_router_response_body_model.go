// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromTransitRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeInstanceFromTransitRouterResponseBody
	GetRequestId() *string
}

type RevokeInstanceFromTransitRouterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AA4BFFD1-5090-5896-935F-4B353557F1A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeInstanceFromTransitRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromTransitRouterResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromTransitRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeInstanceFromTransitRouterResponseBody) SetRequestId(v string) *RevokeInstanceFromTransitRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
