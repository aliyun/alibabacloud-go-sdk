// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePublicNetworkAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleasePublicNetworkAddressResponseBody
	GetRequestId() *string
}

type ReleasePublicNetworkAddressResponseBody struct {
	// example:
	//
	// B18D4390-A968-4444-B323-4360B8E5DA3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleasePublicNetworkAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleasePublicNetworkAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleasePublicNetworkAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleasePublicNetworkAddressResponseBody) SetRequestId(v string) *ReleasePublicNetworkAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleasePublicNetworkAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
