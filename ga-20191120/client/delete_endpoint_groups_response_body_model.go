// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEndpointGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEndpointGroupsResponseBody
	GetRequestId() *string
}

type DeleteEndpointGroupsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEndpointGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEndpointGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEndpointGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEndpointGroupsResponseBody) SetRequestId(v string) *DeleteEndpointGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEndpointGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
