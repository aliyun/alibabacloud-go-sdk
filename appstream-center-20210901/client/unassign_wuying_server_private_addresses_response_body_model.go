// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassignWuyingServerPrivateAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnassignWuyingServerPrivateAddressesResponseBody
	GetRequestId() *string
}

type UnassignWuyingServerPrivateAddressesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassignWuyingServerPrivateAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnassignWuyingServerPrivateAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *UnassignWuyingServerPrivateAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnassignWuyingServerPrivateAddressesResponseBody) SetRequestId(v string) *UnassignWuyingServerPrivateAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassignWuyingServerPrivateAddressesResponseBody) Validate() error {
	return dara.Validate(s)
}
