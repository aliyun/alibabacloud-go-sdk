// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcEndpointAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateVpcEndpointAttributeResponseBody
	GetRequestId() *string
}

type UpdateVpcEndpointAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVpcEndpointAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcEndpointAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVpcEndpointAttributeResponseBody) SetRequestId(v string) *UpdateVpcEndpointAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVpcEndpointAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
