// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateEipAddressWithRCInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateEipAddressWithRCInstanceResponseBody
	GetRequestId() *string
}

type AssociateEipAddressWithRCInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EFFC7565-B3CF-5CFA-9E1F-164DD1E1F498
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateEipAddressWithRCInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateEipAddressWithRCInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateEipAddressWithRCInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateEipAddressWithRCInstanceResponseBody) SetRequestId(v string) *AssociateEipAddressWithRCInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateEipAddressWithRCInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
