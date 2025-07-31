// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocatePublicNetworkAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AllocatePublicNetworkAddressResponseBody
	GetRequestId() *string
}

type AllocatePublicNetworkAddressResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8AA45036-497F-52E7-B951-F9C7B239****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocatePublicNetworkAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocatePublicNetworkAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AllocatePublicNetworkAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocatePublicNetworkAddressResponseBody) SetRequestId(v string) *AllocatePublicNetworkAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocatePublicNetworkAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
