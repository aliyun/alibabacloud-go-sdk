// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteExpressConnectResponseBody
	GetRequestId() *string
}

type DeleteExpressConnectResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 980960B0-2969-40BF-8542-EBB34FD358AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExpressConnectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExpressConnectResponseBody) SetRequestId(v string) *DeleteExpressConnectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExpressConnectResponseBody) Validate() error {
	return dara.Validate(s)
}
