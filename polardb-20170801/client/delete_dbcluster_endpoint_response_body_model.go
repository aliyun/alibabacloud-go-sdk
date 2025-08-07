// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDBClusterEndpointResponseBody
	GetRequestId() *string
}

type DeleteDBClusterEndpointResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBClusterEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBClusterEndpointResponseBody) SetRequestId(v string) *DeleteDBClusterEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBClusterEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
