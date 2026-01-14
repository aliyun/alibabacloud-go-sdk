// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachLogStoreToEndpointGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachLogStoreToEndpointGroupResponseBody
	GetRequestId() *string
}

type AttachLogStoreToEndpointGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachLogStoreToEndpointGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachLogStoreToEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AttachLogStoreToEndpointGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachLogStoreToEndpointGroupResponseBody) SetRequestId(v string) *AttachLogStoreToEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachLogStoreToEndpointGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
