// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachResourceToVpcEndpointServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachResourceToVpcEndpointServiceResponseBody
	GetRequestId() *string
}

type AttachResourceToVpcEndpointServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachResourceToVpcEndpointServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachResourceToVpcEndpointServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AttachResourceToVpcEndpointServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachResourceToVpcEndpointServiceResponseBody) SetRequestId(v string) *AttachResourceToVpcEndpointServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachResourceToVpcEndpointServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
