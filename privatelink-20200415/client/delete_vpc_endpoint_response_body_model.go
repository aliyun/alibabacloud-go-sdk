// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVpcEndpointResponseBody
	GetRequestId() *string
}

type DeleteVpcEndpointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8D8992C1-6712-423C-BAC5-E5E817484C6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVpcEndpointResponseBody) SetRequestId(v string) *DeleteVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpcEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
