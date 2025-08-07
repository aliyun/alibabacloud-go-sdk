// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterEndpointResponseBody
	GetRequestId() *string
}

type ModifyDBClusterEndpointResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterEndpointResponseBody) SetRequestId(v string) *ModifyDBClusterEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
