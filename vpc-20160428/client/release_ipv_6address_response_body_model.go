// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseIpv6AddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseIpv6AddressResponseBody
	GetRequestId() *string
}

type ReleaseIpv6AddressResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C44F62BE-9CE7-4277-B117-69243F3988BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseIpv6AddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseIpv6AddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseIpv6AddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseIpv6AddressResponseBody) SetRequestId(v string) *ReleaseIpv6AddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseIpv6AddressResponseBody) Validate() error {
	return dara.Validate(s)
}
