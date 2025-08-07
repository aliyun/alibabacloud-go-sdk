// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDBClusterEndpointResponseBody
	GetRequestId() *string
}

type CreateDBClusterEndpointResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CD35F3-F3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBClusterEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBClusterEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBClusterEndpointResponseBody) SetRequestId(v string) *CreateDBClusterEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBClusterEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
