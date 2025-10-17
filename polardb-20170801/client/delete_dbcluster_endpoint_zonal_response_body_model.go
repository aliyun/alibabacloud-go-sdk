// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterEndpointZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDBClusterEndpointZonalResponseBody
	GetRequestId() *string
}

type DeleteDBClusterEndpointZonalResponseBody struct {
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBClusterEndpointZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterEndpointZonalResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterEndpointZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBClusterEndpointZonalResponseBody) SetRequestId(v string) *DeleteDBClusterEndpointZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBClusterEndpointZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
