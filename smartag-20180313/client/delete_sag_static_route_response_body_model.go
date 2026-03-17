// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSagStaticRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSagStaticRouteResponseBody
	GetRequestId() *string
}

type DeleteSagStaticRouteResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A53F71B6-7577-492A-A0CD-C7D3DFFE2D0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSagStaticRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSagStaticRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSagStaticRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSagStaticRouteResponseBody) SetRequestId(v string) *DeleteSagStaticRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSagStaticRouteResponseBody) Validate() error {
	return dara.Validate(s)
}
