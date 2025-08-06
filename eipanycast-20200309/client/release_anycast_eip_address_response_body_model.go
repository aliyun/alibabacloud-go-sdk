// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseAnycastEipAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseAnycastEipAddressResponseBody
	GetRequestId() *string
}

type ReleaseAnycastEipAddressResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FBDB18D8-E91E-4978-8D6C-6E2E3EE10133
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseAnycastEipAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseAnycastEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseAnycastEipAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseAnycastEipAddressResponseBody) SetRequestId(v string) *ReleaseAnycastEipAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseAnycastEipAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
