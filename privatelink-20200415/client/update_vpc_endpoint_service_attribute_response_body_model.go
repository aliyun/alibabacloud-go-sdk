// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcEndpointServiceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateVpcEndpointServiceAttributeResponseBody
	GetRequestId() *string
}

type UpdateVpcEndpointServiceAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8D8992C1-6712-423C-BAC5-E5E817484C6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVpcEndpointServiceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcEndpointServiceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointServiceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVpcEndpointServiceAttributeResponseBody) SetRequestId(v string) *UpdateVpcEndpointServiceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
