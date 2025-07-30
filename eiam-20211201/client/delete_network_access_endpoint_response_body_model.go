// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkAccessEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNetworkAccessEndpointResponseBody
	GetRequestId() *string
}

type DeleteNetworkAccessEndpointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkAccessEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkAccessEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkAccessEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetworkAccessEndpointResponseBody) SetRequestId(v string) *DeleteNetworkAccessEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkAccessEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
