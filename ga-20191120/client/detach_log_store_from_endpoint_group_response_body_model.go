// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachLogStoreFromEndpointGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachLogStoreFromEndpointGroupResponseBody
	GetRequestId() *string
}

type DetachLogStoreFromEndpointGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 64ADAB1E-0B7F-4FD8-A404-3BECC0E9CCFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachLogStoreFromEndpointGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachLogStoreFromEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DetachLogStoreFromEndpointGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachLogStoreFromEndpointGroupResponseBody) SetRequestId(v string) *DetachLogStoreFromEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachLogStoreFromEndpointGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
