// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseIpAddressResponseBody
	GetRequestId() *string
}

type ReleaseIpAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseIpAddressResponseBody) SetRequestId(v string) *ReleaseIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseIpAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
