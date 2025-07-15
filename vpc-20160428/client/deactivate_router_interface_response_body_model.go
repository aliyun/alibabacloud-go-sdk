// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateRouterInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeactivateRouterInterfaceResponseBody
	GetRequestId() *string
}

type DeactivateRouterInterfaceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// BE7EB53A-99AB-4DA8-AEDE-75FA90D046A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeactivateRouterInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeactivateRouterInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeactivateRouterInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeactivateRouterInterfaceResponseBody) SetRequestId(v string) *DeactivateRouterInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeactivateRouterInterfaceResponseBody) Validate() error {
	return dara.Validate(s)
}
