// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcEndpointServiceResourceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateVpcEndpointServiceResourceAttributeResponseBody
	GetRequestId() *string
}

type UpdateVpcEndpointServiceResourceAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVpcEndpointServiceResourceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcEndpointServiceResourceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointServiceResourceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVpcEndpointServiceResourceAttributeResponseBody) SetRequestId(v string) *UpdateVpcEndpointServiceResourceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
