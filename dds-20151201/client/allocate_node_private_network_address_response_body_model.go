// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateNodePrivateNetworkAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AllocateNodePrivateNetworkAddressResponseBody
	GetRequestId() *string
}

type AllocateNodePrivateNetworkAddressResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 60EEBD77-227C-5B39-86EA-D89163C5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateNodePrivateNetworkAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateNodePrivateNetworkAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateNodePrivateNetworkAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateNodePrivateNetworkAddressResponseBody) SetRequestId(v string) *AllocateNodePrivateNetworkAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
