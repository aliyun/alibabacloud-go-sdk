// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEndpointGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEndpointGroupResponseBody
	GetRequestId() *string
}

type DeleteEndpointGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEndpointGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEndpointGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEndpointGroupResponseBody) SetRequestId(v string) *DeleteEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEndpointGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
