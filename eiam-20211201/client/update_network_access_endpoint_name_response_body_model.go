// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkAccessEndpointNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNetworkAccessEndpointNameResponseBody
	GetRequestId() *string
}

type UpdateNetworkAccessEndpointNameResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNetworkAccessEndpointNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkAccessEndpointNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNetworkAccessEndpointNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNetworkAccessEndpointNameResponseBody) SetRequestId(v string) *UpdateNetworkAccessEndpointNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNetworkAccessEndpointNameResponseBody) Validate() error {
	return dara.Validate(s)
}
