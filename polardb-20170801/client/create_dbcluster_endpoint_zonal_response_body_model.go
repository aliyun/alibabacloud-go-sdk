// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterEndpointZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDBClusterEndpointZonalResponseBody
	GetRequestId() *string
}

type CreateDBClusterEndpointZonalResponseBody struct {
	// example:
	//
	// CD35F3-F3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBClusterEndpointZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterEndpointZonalResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBClusterEndpointZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBClusterEndpointZonalResponseBody) SetRequestId(v string) *CreateDBClusterEndpointZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBClusterEndpointZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
