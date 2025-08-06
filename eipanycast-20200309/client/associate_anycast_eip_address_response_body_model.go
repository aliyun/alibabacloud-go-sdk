// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateAnycastEipAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateAnycastEipAddressResponseBody
	GetRequestId() *string
}

type AssociateAnycastEipAddressResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FBDB18D8-E91E-4978-8D6C-6E2E3EE10133
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateAnycastEipAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateAnycastEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateAnycastEipAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateAnycastEipAddressResponseBody) SetRequestId(v string) *AssociateAnycastEipAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateAnycastEipAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
