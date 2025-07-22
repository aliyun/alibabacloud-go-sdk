// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateIpAddressResponseBody
	GetRequestId() *string
}

type AssociateIpAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateIpAddressResponseBody) SetRequestId(v string) *AssociateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateIpAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
