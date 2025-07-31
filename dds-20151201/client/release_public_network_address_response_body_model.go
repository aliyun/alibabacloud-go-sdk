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
	// The request ID.
	//
	// example:
	//
	// 1D6AFE36-1AF5-4DE4-A954-672159D4CC69
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
