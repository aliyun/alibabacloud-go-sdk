// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseNodePrivateNetworkAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseNodePrivateNetworkAddressResponseBody
	GetRequestId() *string
}

type ReleaseNodePrivateNetworkAddressResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0FDDC511-7252-4A4A-ADDA-5CB1BF63873D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseNodePrivateNetworkAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseNodePrivateNetworkAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseNodePrivateNetworkAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseNodePrivateNetworkAddressResponseBody) SetRequestId(v string) *ReleaseNodePrivateNetworkAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
