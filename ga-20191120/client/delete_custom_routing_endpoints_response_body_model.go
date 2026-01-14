// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomRoutingEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomRoutingEndpointsResponseBody
	GetRequestId() *string
}

type DeleteCustomRoutingEndpointsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomRoutingEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomRoutingEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomRoutingEndpointsResponseBody) SetRequestId(v string) *DeleteCustomRoutingEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomRoutingEndpointsResponseBody) Validate() error {
	return dara.Validate(s)
}
