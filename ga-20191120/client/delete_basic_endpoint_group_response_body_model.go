// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicEndpointGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBasicEndpointGroupResponseBody
	GetRequestId() *string
}

type DeleteBasicEndpointGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBasicEndpointGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBasicEndpointGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBasicEndpointGroupResponseBody) SetRequestId(v string) *DeleteBasicEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBasicEndpointGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
