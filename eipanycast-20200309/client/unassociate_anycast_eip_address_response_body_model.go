// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateAnycastEipAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnassociateAnycastEipAddressResponseBody
	GetRequestId() *string
}

type UnassociateAnycastEipAddressResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FBDB18D8-E91E-4978-8D6C-6E2E3EE10133
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassociateAnycastEipAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnassociateAnycastEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *UnassociateAnycastEipAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnassociateAnycastEipAddressResponseBody) SetRequestId(v string) *UnassociateAnycastEipAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassociateAnycastEipAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
