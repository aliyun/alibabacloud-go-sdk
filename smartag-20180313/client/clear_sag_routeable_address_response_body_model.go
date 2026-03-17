// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearSagRouteableAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ClearSagRouteableAddressResponseBody
	GetRequestId() *string
}

type ClearSagRouteableAddressResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AEEC8A5A-360E-4865-82D4-38CDE46445FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearSagRouteableAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearSagRouteableAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ClearSagRouteableAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearSagRouteableAddressResponseBody) SetRequestId(v string) *ClearSagRouteableAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearSagRouteableAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
